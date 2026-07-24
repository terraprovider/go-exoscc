package adminapi

import (
	"errors"
	"strings"
	"testing"
)

// A real 400 body: the actionable message is nested (JSON) inside
// error.details[].message; error.message itself is the generic envelope.
const alreadyExistsBody = `{"error":{"code":"BadRequest","message":"Error executing cmdlet",` +
	`"details":[{"code":"Context","target":"",` +
	`"message":"{\"Message\":\"The object 'CN=rg,OU=x' already exists.\",\"TypeName\":\"Microsoft.Exchange.Data.Directory.ADObjectAlreadyExistsException\",\"StackTrace\":null,\"InnerError\":null}"}],` +
	`"innererror":{"message":"Error executing cmdlet","type":"ODataServiceException"}}}`

const notFoundBody = `{"error":{"code":"NotFound","message":"Error executing cmdlet",` +
	`"details":[{"code":"Client","message":"{\"Message\":\"The operation couldn't be performed because object 'rg' couldn't be found.\",\"TypeName\":\"Microsoft.Exchange.Configuration.Tasks.ManagementObjectNotFoundException\",\"InnerError\":null}"}]}}`

const noDetailsBody = `{"error":{"code":"BadRequest","message":"Error executing cmdlet"}}`

func TestParseAPIError_ExtractsRealMessage(t *testing.T) {
	err := parseAPIError(400, []byte(alreadyExistsBody))
	var ae *APIError
	if !errors.As(err, &ae) {
		t.Fatalf("want *APIError, got %T", err)
	}
	if !strings.Contains(ae.Message, "already exists") {
		t.Errorf("message not extracted: %q", ae.Message)
	}
	if ae.Type != "Microsoft.Exchange.Data.Directory.ADObjectAlreadyExistsException" {
		t.Errorf("type not extracted: %q", ae.Type)
	}
	if ae.Status != 400 || ae.Code != "BadRequest" {
		t.Errorf("status/code: %d %q", ae.Status, ae.Code)
	}
	// Error() shows the real message + exception type, not the generic envelope.
	if s := ae.Error(); strings.Contains(s, "Error executing cmdlet") || !strings.Contains(s, "already exists") {
		t.Errorf("Error() = %q", s)
	}
}

func TestParseAPIError_NotFoundMessage(t *testing.T) {
	err := parseAPIError(404, []byte(notFoundBody))
	var ae *APIError
	if !errors.As(err, &ae) || !strings.Contains(ae.Message, "couldn't be found") {
		t.Fatalf("not-found message not surfaced: %v", err)
	}
}

func TestParseAPIError_FallsBackWhenNoDetails(t *testing.T) {
	err := parseAPIError(400, []byte(noDetailsBody))
	var ae *APIError
	if !errors.As(err, &ae) || ae.Message != "Error executing cmdlet" {
		t.Fatalf("fallback failed: %v", err)
	}
}
