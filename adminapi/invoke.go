package adminapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// RateLimit is the throttling state parsed from the response headers.
type RateLimit struct {
	Limit     int
	Remaining int
	Reset     time.Time
}

// APIError is an OData error returned by the Admin API (e.g. a Forbidden cmdlet).
// Message carries the actual cmdlet error (e.g. "...The object '...' already
// exists.") extracted from error.details, not the generic "Error executing
// cmdlet" envelope; Type is the underlying exception type name when known.
type APIError struct {
	Status   int
	Code     string
	Type     string
	Message  string
	InnerRaw json.RawMessage
}

func (e *APIError) Error() string {
	code := e.Code
	if e.Type != "" {
		code = e.Type
	}
	return fmt.Sprintf("adminapi: %s (%s, http %d)", e.Message, code, e.Status)
}

// odataError is the OData "error" object of an InvokeCommand response.
type odataError struct {
	Code       string          `json:"code"`
	Message    string          `json:"message"`
	Details    []odataDetail   `json:"details"`
	InnerError json.RawMessage `json:"innererror"`
}

// odataDetail carries the real cmdlet error. Its Message is itself a JSON string
// describing the thrown exception (see cmdletException).
type odataDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// cmdletException is the JSON encoded inside odataDetail.Message.
type cmdletException struct {
	Message  string           `json:"Message"`
	TypeName string           `json:"TypeName"`
	InnerErr *cmdletException `json:"InnerError"`
}

// realError digs the actionable cmdlet message + exception type out of the
// error details, falling back to the generic envelope message.
func (e *odataError) realError() (msg, typeName string) {
	for _, d := range e.Details {
		if d.Message == "" {
			continue
		}
		var ce cmdletException
		if json.Unmarshal([]byte(d.Message), &ce) == nil && ce.Message != "" {
			return ce.Message, ce.TypeName
		}
	}
	return e.Message, ""
}

func (e *odataError) toAPIError(status int) *APIError {
	msg, typ := e.realError()
	return &APIError{Status: status, Code: e.Code, Type: typ, Message: msg, InnerRaw: e.InnerError}
}

// odataResponse is the InvokeCommand response envelope.
type odataResponse struct {
	Context  string           `json:"@odata.context"`
	NextLink string           `json:"@odata.nextLink"`
	Warnings []string         `json:"@adminapi.warnings"`
	Value    []map[string]any `json:"value"`
	Error    *odataError      `json:"error"`
}

// Result carries the objects plus per-call metadata.
type Result struct {
	Value     []map[string]any
	Warnings  []string
	RateLimit RateLimit
}

// Decode unmarshals the result Value into v (e.g. *[]models.RoleGroup). It round-
// trips through JSON, so v's fields should carry json tags matching the OData
// property names (the generated models package does).
func (r *Result) Decode(v any) error {
	b, err := json.Marshal(r.Value)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

// Invoke runs one cmdlet via InvokeCommand and returns all pages of results.
// params holds only the bound parameters (PowerShell semantics); generated
// request types build this map for you.
func (c *Client) Invoke(ctx context.Context, cmdlet string, params map[string]any) (*Result, error) {
	if params == nil {
		params = map[string]any{}
	}
	body, err := json.Marshal(map[string]any{
		"CmdletInput": map[string]any{"CmdletName": cmdlet, "Parameters": params},
	})
	if err != nil {
		return nil, err
	}

	out := &Result{}
	resp, err := c.postInvoke(ctx, cmdlet, body)
	if err != nil {
		return nil, err
	}
	out.RateLimit = parseRateLimit(resp.Header)
	raw, err := readBody(resp)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return out, parseAPIError(resp.StatusCode, raw)
	}
	var env odataResponse
	if err := json.Unmarshal(raw, &env); err != nil {
		return nil, fmt.Errorf("adminapi: decode response: %w", err)
	}
	if env.Error != nil {
		return out, env.Error.toAPIError(resp.StatusCode)
	}
	out.Value = append(out.Value, env.Value...)
	out.Warnings = append(out.Warnings, env.Warnings...)
	// Follow @odata.nextLink pages (GET) until exhausted.
	return c.followPages(ctx, cmdlet, out, env.NextLink)
}

// followPages continues GET paging from a nextLink until exhausted.
func (c *Client) followPages(ctx context.Context, cmdlet string, out *Result, nextLink string) (*Result, error) {
	for nextLink != "" {
		pg, err := c.getPage(ctx, cmdlet, normalizeURL(nextLink))
		if err != nil {
			return out, err
		}
		out.Value = append(out.Value, pg.Value...)
		out.Warnings = append(out.Warnings, pg.Warnings...)
		nextLink = pg.NextLink
	}
	return out, nil
}

type pageResult struct {
	Value     []map[string]any
	Warnings  []string
	NextLink  string
	RateLimit RateLimit
}

func (c *Client) getPage(ctx context.Context, cmdlet, url string) (*pageResult, error) {
	token, err := c.opt.Tokens.Token(ctx, c.opt.Cloud.Resource)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	c.setInvokeHeaders(req, cmdlet, token)
	req.Header.Del("Content-Type")
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	raw, err := readBody(resp)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, parseAPIError(resp.StatusCode, raw)
	}
	var env odataResponse
	if err := json.Unmarshal(raw, &env); err != nil {
		return nil, err
	}
	return &pageResult{Value: env.Value, Warnings: env.Warnings, NextLink: env.NextLink, RateLimit: parseRateLimit(resp.Header)}, nil
}

func parseAPIError(status int, raw []byte) error {
	var env odataResponse
	if json.Unmarshal(raw, &env) == nil && env.Error != nil {
		return env.Error.toAPIError(status)
	}
	msg := strings.TrimSpace(string(raw))
	if msg == "" {
		msg = "(no response body)"
	} else if len(msg) > 500 {
		msg = msg[:500]
	}
	if status == 403 {
		msg += " — the app likely lacks a directory role (assign one for app-only, e.g. Global Reader)"
	}
	return &APIError{Status: status, Code: "Unknown", Message: msg}
}

func parseRateLimit(h http.Header) RateLimit {
	rl := RateLimit{}
	rl.Limit, _ = strconv.Atoi(h.Get("Rate-Limit-Limit"))
	rl.Remaining, _ = strconv.Atoi(h.Get("Rate-Limit-Remaining"))
	if v := h.Get("Rate-Limit-Reset"); v != "" {
		if t, err := time.Parse(time.RFC3339, v); err == nil {
			rl.Reset = t
		}
	}
	return rl
}
