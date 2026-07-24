// Code generated from Purview-ExchangeOnline.psm1 by gen-go. DO NOT EDIT.

package purview

import (
	"context"

	"github.com/terraprovider/go-exoscc/adminapi"
)

// Service exposes the 418 cmdlets of Purview-ExchangeOnline.psm1 as typed methods.
type Service struct{ C *adminapi.Client }

// New wraps an *adminapi.Client.
func New(c *adminapi.Client) *Service { return &Service{C: c} }

// AddComplianceCaseMemberParams are the parameters of Add-ComplianceCaseMember.
// DefaultParameterSetName: Identity
type AddComplianceCaseMemberParams struct {
	Case   string `ps:"Case"`
	Member string `ps:"Member"`
}

func (p AddComplianceCaseMemberParams) params() map[string]any {
	m := map[string]any{}
	if p.Case != "" {
		m["Case"] = p.Case
	}
	if p.Member != "" {
		m["Member"] = p.Member
	}
	return m
}

// AddComplianceCaseMember runs the Add-ComplianceCaseMember cmdlet.
func (s *Service) AddComplianceCaseMember(ctx context.Context, p AddComplianceCaseMemberParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Add-ComplianceCaseMember", p.params())
}

// AddRoleGroupMemberParams are the parameters of Add-RoleGroupMember.
type AddRoleGroupMemberParams struct {
}

func (p AddRoleGroupMemberParams) params() map[string]any {
	m := map[string]any{}
	return m
}

// AddRoleGroupMember runs the Add-RoleGroupMember cmdlet.
func (s *Service) AddRoleGroupMember(ctx context.Context, p AddRoleGroupMemberParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Add-RoleGroupMember", p.params())
}

// CancelDlpEdmSessionParams are the parameters of Cancel-DlpEdmSession.
type CancelDlpEdmSessionParams struct {
	DataStoreName string `ps:"DataStoreName"`
	SessionId     string `ps:"SessionId"`
}

func (p CancelDlpEdmSessionParams) params() map[string]any {
	m := map[string]any{}
	if p.DataStoreName != "" {
		m["DataStoreName"] = p.DataStoreName
	}
	if p.SessionId != "" {
		m["SessionId"] = p.SessionId
	}
	return m
}

// CancelDlpEdmSession runs the Cancel-DlpEdmSession cmdlet.
func (s *Service) CancelDlpEdmSession(ctx context.Context, p CancelDlpEdmSessionParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Cancel-DlpEdmSession", p.params())
}

// CancelSensitiveInformationScanParams are the parameters of Cancel-SensitiveInformationScan.
// DefaultParameterSetName: Identity
type CancelSensitiveInformationScanParams struct {
	Identity any `ps:"Identity"`
}

func (p CancelSensitiveInformationScanParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// CancelSensitiveInformationScan runs the Cancel-SensitiveInformationScan cmdlet.
func (s *Service) CancelSensitiveInformationScan(ctx context.Context, p CancelSensitiveInformationScanParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Cancel-SensitiveInformationScan", p.params())
}

// CheckPurviewConfigParams are the parameters of Check-PurviewConfig.
type CheckPurviewConfigParams struct {
	CaseHoldPolicyNameOrId string   `ps:"CaseHoldPolicyNameOrId"`
	CaseId                 string   `ps:"CaseId"`
	Component              any      `ps:"Component"`
	DateTimeUTC            string   `ps:"DateTimeUTC"`
	DeviceName             string   `ps:"DeviceName"`
	File                   string   `ps:"File"`
	FileAsBytes            []string `ps:"FileAsBytes"`
	IncidentId             string   `ps:"IncidentId"`
	ItemId                 string   `ps:"ItemId"`
	MessageId              string   `ps:"MessageId"`
	Organization           any      `ps:"Organization"`
	PolicyName             string   `ps:"PolicyName"`
	RecipientAddress       any      `ps:"RecipientAddress"`
	RecordId               string   `ps:"RecordId"`
	RuleName               string   `ps:"RuleName"`
	SenderAddress          any      `ps:"SenderAddress"`
	SharepointItemUniqueId string   `ps:"SharepointItemUniqueId"`
	SharepointSiteId       string   `ps:"SharepointSiteId"`
	SiteUrl                string   `ps:"SiteUrl"`
	SitIdentity            string   `ps:"SitIdentity"`
	TestCases              any      `ps:"TestCases"`
	Theme                  any      `ps:"Theme"`
	UserPrincipalName      any      `ps:"UserPrincipalName"`
	Workload               string   `ps:"Workload"`
}

func (p CheckPurviewConfigParams) params() map[string]any {
	m := map[string]any{}
	if p.CaseHoldPolicyNameOrId != "" {
		m["CaseHoldPolicyNameOrId"] = p.CaseHoldPolicyNameOrId
	}
	if p.CaseId != "" {
		m["CaseId"] = p.CaseId
	}
	if p.Component != nil {
		m["Component"] = p.Component
	}
	if p.DateTimeUTC != "" {
		m["DateTimeUTC"] = p.DateTimeUTC
	}
	if p.DeviceName != "" {
		m["DeviceName"] = p.DeviceName
	}
	if p.File != "" {
		m["File"] = p.File
	}
	if len(p.FileAsBytes) > 0 {
		m["FileAsBytes"] = p.FileAsBytes
	}
	if p.IncidentId != "" {
		m["IncidentId"] = p.IncidentId
	}
	if p.ItemId != "" {
		m["ItemId"] = p.ItemId
	}
	if p.MessageId != "" {
		m["MessageId"] = p.MessageId
	}
	if p.Organization != nil {
		m["Organization"] = p.Organization
	}
	if p.PolicyName != "" {
		m["PolicyName"] = p.PolicyName
	}
	if p.RecipientAddress != nil {
		m["RecipientAddress"] = p.RecipientAddress
	}
	if p.RecordId != "" {
		m["RecordId"] = p.RecordId
	}
	if p.RuleName != "" {
		m["RuleName"] = p.RuleName
	}
	if p.SenderAddress != nil {
		m["SenderAddress"] = p.SenderAddress
	}
	if p.SharepointItemUniqueId != "" {
		m["SharepointItemUniqueId"] = p.SharepointItemUniqueId
	}
	if p.SharepointSiteId != "" {
		m["SharepointSiteId"] = p.SharepointSiteId
	}
	if p.SiteUrl != "" {
		m["SiteUrl"] = p.SiteUrl
	}
	if p.SitIdentity != "" {
		m["SitIdentity"] = p.SitIdentity
	}
	if p.TestCases != nil {
		m["TestCases"] = p.TestCases
	}
	if p.Theme != nil {
		m["Theme"] = p.Theme
	}
	if p.UserPrincipalName != nil {
		m["UserPrincipalName"] = p.UserPrincipalName
	}
	if p.Workload != "" {
		m["Workload"] = p.Workload
	}
	return m
}

// CheckPurviewConfig runs the Check-PurviewConfig cmdlet.
func (s *Service) CheckPurviewConfig(ctx context.Context, p CheckPurviewConfigParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Check-PurviewConfig", p.params())
}

// CreateFilePlanFirstRunLabelsParams are the parameters of Create-FilePlanFirstRunLabels.
type CreateFilePlanFirstRunLabelsParams struct {
	DomainController any `ps:"DomainController"`
}

func (p CreateFilePlanFirstRunLabelsParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	return m
}

// CreateFilePlanFirstRunLabels runs the Create-FilePlanFirstRunLabels cmdlet.
func (s *Service) CreateFilePlanFirstRunLabels(ctx context.Context, p CreateFilePlanFirstRunLabelsParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Create-FilePlanFirstRunLabels", p.params())
}

// DeleteQuarantineMessageParams are the parameters of Delete-QuarantineMessage.
type DeleteQuarantineMessageParams struct {
	EntityType       any      `ps:"EntityType"`
	HardDelete       bool     `ps:"HardDelete"`
	Identities       []string `ps:"Identities"`
	Identity         any      `ps:"Identity"`
	RecipientAddress []string `ps:"RecipientAddress"`
}

func (p DeleteQuarantineMessageParams) params() map[string]any {
	m := map[string]any{}
	if p.EntityType != nil {
		m["EntityType"] = p.EntityType
	}
	if p.HardDelete {
		m["HardDelete"] = true
	}
	if len(p.Identities) > 0 {
		m["Identities"] = p.Identities
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if len(p.RecipientAddress) > 0 {
		m["RecipientAddress"] = p.RecipientAddress
	}
	return m
}

// DeleteQuarantineMessage runs the Delete-QuarantineMessage cmdlet.
func (s *Service) DeleteQuarantineMessage(ctx context.Context, p DeleteQuarantineMessageParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Delete-QuarantineMessage", p.params())
}

// EnableAdaptiveScopeStorageParams are the parameters of Enable-AdaptiveScopeStorage.
type EnableAdaptiveScopeStorageParams struct {
}

func (p EnableAdaptiveScopeStorageParams) params() map[string]any {
	m := map[string]any{}
	return m
}

// EnableAdaptiveScopeStorage runs the Enable-AdaptiveScopeStorage cmdlet.
func (s *Service) EnableAdaptiveScopeStorage(ctx context.Context, p EnableAdaptiveScopeStorageParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Enable-AdaptiveScopeStorage", p.params())
}

// EnableComplianceFeatureParams are the parameters of Enable-ComplianceFeature.
type EnableComplianceFeatureParams struct {
}

func (p EnableComplianceFeatureParams) params() map[string]any {
	m := map[string]any{}
	return m
}

// EnableComplianceFeature runs the Enable-ComplianceFeature cmdlet.
func (s *Service) EnableComplianceFeature(ctx context.Context, p EnableComplianceFeatureParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Enable-ComplianceFeature", p.params())
}

// EnableComplianceTagStorageParams are the parameters of Enable-ComplianceTagStorage.
type EnableComplianceTagStorageParams struct {
	RecordsManagementSecurityGroupEmail string `ps:"RecordsManagementSecurityGroupEmail"`
}

func (p EnableComplianceTagStorageParams) params() map[string]any {
	m := map[string]any{}
	if p.RecordsManagementSecurityGroupEmail != "" {
		m["RecordsManagementSecurityGroupEmail"] = p.RecordsManagementSecurityGroupEmail
	}
	return m
}

// EnableComplianceTagStorage runs the Enable-ComplianceTagStorage cmdlet.
func (s *Service) EnableComplianceTagStorage(ctx context.Context, p EnableComplianceTagStorageParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Enable-ComplianceTagStorage", p.params())
}

// EnableServiceDomainGroupStorageParams are the parameters of Enable-ServiceDomainGroupStorage.
type EnableServiceDomainGroupStorageParams struct {
}

func (p EnableServiceDomainGroupStorageParams) params() map[string]any {
	m := map[string]any{}
	return m
}

// EnableServiceDomainGroupStorage runs the Enable-ServiceDomainGroupStorage cmdlet.
func (s *Service) EnableServiceDomainGroupStorage(ctx context.Context, p EnableServiceDomainGroupStorageParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Enable-ServiceDomainGroupStorage", p.params())
}

// ExecuteAzureAdLabelSyncParams are the parameters of Execute-AzureAdLabelSync.
type ExecuteAzureAdLabelSyncParams struct {
}

func (p ExecuteAzureAdLabelSyncParams) params() map[string]any {
	m := map[string]any{}
	return m
}

// ExecuteAzureAdLabelSync runs the Execute-AzureAdLabelSync cmdlet.
func (s *Service) ExecuteAzureAdLabelSync(ctx context.Context, p ExecuteAzureAdLabelSyncParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Execute-AzureAdLabelSync", p.params())
}

// ExecuteUnifiedPolicyCmdletBatchParams are the parameters of Execute-UnifiedPolicyCmdletBatch.
// DefaultParameterSetName: Identity
type ExecuteUnifiedPolicyCmdletBatchParams struct {
	CmdletData  []string `ps:"CmdletData"`
	ExecutionId string   `ps:"ExecutionId"`
	MigrationId string   `ps:"MigrationId"`
}

func (p ExecuteUnifiedPolicyCmdletBatchParams) params() map[string]any {
	m := map[string]any{}
	if len(p.CmdletData) > 0 {
		m["CmdletData"] = p.CmdletData
	}
	if p.ExecutionId != "" {
		m["ExecutionId"] = p.ExecutionId
	}
	if p.MigrationId != "" {
		m["MigrationId"] = p.MigrationId
	}
	return m
}

// ExecuteUnifiedPolicyCmdletBatch runs the Execute-UnifiedPolicyCmdletBatch cmdlet.
func (s *Service) ExecuteUnifiedPolicyCmdletBatch(ctx context.Context, p ExecuteUnifiedPolicyCmdletBatchParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Execute-UnifiedPolicyCmdletBatch", p.params())
}

// ExportActivityExplorerDataParams are the parameters of Export-ActivityExplorerData.
type ExportActivityExplorerDataParams struct {
	EndTime      any      `ps:"EndTime"`
	Filter1      []string `ps:"Filter1"`
	Filter2      []string `ps:"Filter2"`
	Filter3      []string `ps:"Filter3"`
	Filter4      []string `ps:"Filter4"`
	Filter5      []string `ps:"Filter5"`
	OutputFormat string   `ps:"OutputFormat"` // one of: csv, json
	PageCookie   string   `ps:"PageCookie"`
	PageSize     int      `ps:"PageSize"`
	StartTime    any      `ps:"StartTime"`
}

func (p ExportActivityExplorerDataParams) params() map[string]any {
	m := map[string]any{}
	if p.EndTime != nil {
		m["EndTime"] = p.EndTime
	}
	if len(p.Filter1) > 0 {
		m["Filter1"] = p.Filter1
	}
	if len(p.Filter2) > 0 {
		m["Filter2"] = p.Filter2
	}
	if len(p.Filter3) > 0 {
		m["Filter3"] = p.Filter3
	}
	if len(p.Filter4) > 0 {
		m["Filter4"] = p.Filter4
	}
	if len(p.Filter5) > 0 {
		m["Filter5"] = p.Filter5
	}
	if p.OutputFormat != "" {
		m["OutputFormat"] = p.OutputFormat
	}
	if p.PageCookie != "" {
		m["PageCookie"] = p.PageCookie
	}
	if p.PageSize != 0 {
		m["PageSize"] = p.PageSize
	}
	if p.StartTime != nil {
		m["StartTime"] = p.StartTime
	}
	return m
}

// ExportActivityExplorerData runs the Export-ActivityExplorerData cmdlet.
func (s *Service) ExportActivityExplorerData(ctx context.Context, p ExportActivityExplorerDataParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Export-ActivityExplorerData", p.params())
}

// ExportFilePlanPropertyParams are the parameters of Export-FilePlanProperty.
// DefaultParameterSetName: Identity
type ExportFilePlanPropertyParams struct {
	DomainController any `ps:"DomainController"`
}

func (p ExportFilePlanPropertyParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	return m
}

// ExportFilePlanProperty runs the Export-FilePlanProperty cmdlet.
func (s *Service) ExportFilePlanProperty(ctx context.Context, p ExportFilePlanPropertyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Export-FilePlanProperty", p.params())
}

// ExportPurviewConfigParams are the parameters of Export-PurviewConfig.
type ExportPurviewConfigParams struct {
	Components        any    `ps:"Components"`
	DomainController  any    `ps:"DomainController"`
	PolicyName        string `ps:"PolicyName"`
	UserPrincipalName any    `ps:"UserPrincipalName"`
}

func (p ExportPurviewConfigParams) params() map[string]any {
	m := map[string]any{}
	if p.Components != nil {
		m["Components"] = p.Components
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.PolicyName != "" {
		m["PolicyName"] = p.PolicyName
	}
	if p.UserPrincipalName != nil {
		m["UserPrincipalName"] = p.UserPrincipalName
	}
	return m
}

// ExportPurviewConfig runs the Export-PurviewConfig cmdlet.
func (s *Service) ExportPurviewConfig(ctx context.Context, p ExportPurviewConfigParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Export-PurviewConfig", p.params())
}

// ExportQuarantineMessageParams are the parameters of Export-QuarantineMessage.
type ExportQuarantineMessageParams struct {
	CompressOutput        bool     `ps:"CompressOutput"`
	EntityType            any      `ps:"EntityType"`
	ForceConversionToMime bool     `ps:"ForceConversionToMime"`
	Identities            []string `ps:"Identities"`
	Identity              any      `ps:"Identity"`
	Password              any      `ps:"Password"`
	PasswordV2            string   `ps:"PasswordV2"`
	ReasonForExport       string   `ps:"ReasonForExport"`
	RecipientAddress      string   `ps:"RecipientAddress"`
}

func (p ExportQuarantineMessageParams) params() map[string]any {
	m := map[string]any{}
	if p.CompressOutput {
		m["CompressOutput"] = true
	}
	if p.EntityType != nil {
		m["EntityType"] = p.EntityType
	}
	if p.ForceConversionToMime {
		m["ForceConversionToMime"] = true
	}
	if len(p.Identities) > 0 {
		m["Identities"] = p.Identities
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Password != nil {
		m["Password"] = p.Password
	}
	if p.PasswordV2 != "" {
		m["PasswordV2"] = p.PasswordV2
	}
	if p.ReasonForExport != "" {
		m["ReasonForExport"] = p.ReasonForExport
	}
	if p.RecipientAddress != "" {
		m["RecipientAddress"] = p.RecipientAddress
	}
	return m
}

// ExportQuarantineMessage runs the Export-QuarantineMessage cmdlet.
func (s *Service) ExportQuarantineMessage(ctx context.Context, p ExportQuarantineMessageParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Export-QuarantineMessage", p.params())
}

// ExportQuarantineMessageV1Params are the parameters of Export-QuarantineMessageV1.
type ExportQuarantineMessageV1Params struct {
	CompressOutput        bool     `ps:"CompressOutput"`
	EntityType            any      `ps:"EntityType"`
	ForceConversionToMime bool     `ps:"ForceConversionToMime"`
	Identities            []string `ps:"Identities"`
	Identity              any      `ps:"Identity"`
	Password              any      `ps:"Password"`
	PasswordV2            string   `ps:"PasswordV2"`
	ReasonForExport       string   `ps:"ReasonForExport"`
	RecipientAddress      string   `ps:"RecipientAddress"`
}

func (p ExportQuarantineMessageV1Params) params() map[string]any {
	m := map[string]any{}
	if p.CompressOutput {
		m["CompressOutput"] = true
	}
	if p.EntityType != nil {
		m["EntityType"] = p.EntityType
	}
	if p.ForceConversionToMime {
		m["ForceConversionToMime"] = true
	}
	if len(p.Identities) > 0 {
		m["Identities"] = p.Identities
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Password != nil {
		m["Password"] = p.Password
	}
	if p.PasswordV2 != "" {
		m["PasswordV2"] = p.PasswordV2
	}
	if p.ReasonForExport != "" {
		m["ReasonForExport"] = p.ReasonForExport
	}
	if p.RecipientAddress != "" {
		m["RecipientAddress"] = p.RecipientAddress
	}
	return m
}

// ExportQuarantineMessageV1 runs the Export-QuarantineMessageV1 cmdlet.
func (s *Service) ExportQuarantineMessageV1(ctx context.Context, p ExportQuarantineMessageV1Params) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Export-QuarantineMessageV1", p.params())
}

// FixPurviewConfigParams are the parameters of Fix-PurviewConfig.
type FixPurviewConfigParams struct {
	Component    any    `ps:"Component"`
	Organization any    `ps:"Organization"`
	PolicyId     string `ps:"PolicyId"`
	Scenario     any    `ps:"Scenario"`
	SiteUrl      string `ps:"SiteUrl"`
	Workload     string `ps:"Workload"`
}

func (p FixPurviewConfigParams) params() map[string]any {
	m := map[string]any{}
	if p.Component != nil {
		m["Component"] = p.Component
	}
	if p.Organization != nil {
		m["Organization"] = p.Organization
	}
	if p.PolicyId != "" {
		m["PolicyId"] = p.PolicyId
	}
	if p.Scenario != nil {
		m["Scenario"] = p.Scenario
	}
	if p.SiteUrl != "" {
		m["SiteUrl"] = p.SiteUrl
	}
	if p.Workload != "" {
		m["Workload"] = p.Workload
	}
	return m
}

// FixPurviewConfig runs the Fix-PurviewConfig cmdlet.
func (s *Service) FixPurviewConfig(ctx context.Context, p FixPurviewConfigParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Fix-PurviewConfig", p.params())
}

// GetAadProtectionLevelParams are the parameters of Get-AadProtectionLevel.
type GetAadProtectionLevelParams struct {
	IncludeUnavailableItems bool `ps:"IncludeUnavailableItems"`
}

func (p GetAadProtectionLevelParams) params() map[string]any {
	m := map[string]any{}
	if p.IncludeUnavailableItems {
		m["IncludeUnavailableItems"] = true
	}
	return m
}

// GetAadProtectionLevel runs the Get-AadProtectionLevel cmdlet.
func (s *Service) GetAadProtectionLevel(ctx context.Context, p GetAadProtectionLevelParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-AadProtectionLevel", p.params())
}

// GetActivityAlertParams are the parameters of Get-ActivityAlert.
// DefaultParameterSetName: Identity
type GetActivityAlertParams struct {
	Identity any `ps:"Identity"`
}

func (p GetActivityAlertParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetActivityAlert runs the Get-ActivityAlert cmdlet.
func (s *Service) GetActivityAlert(ctx context.Context, p GetActivityAlertParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ActivityAlert", p.params())
}

// GetAdaptiveScopeParams are the parameters of Get-AdaptiveScope.
// DefaultParameterSetName: Identity
type GetAdaptiveScopeParams struct {
	AdministrativeUnits any `ps:"AdministrativeUnits"`
	Identity            any `ps:"Identity"`
	LocationTypes       any `ps:"LocationTypes"`
}

func (p GetAdaptiveScopeParams) params() map[string]any {
	m := map[string]any{}
	if p.AdministrativeUnits != nil {
		m["AdministrativeUnits"] = p.AdministrativeUnits
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.LocationTypes != nil {
		m["LocationTypes"] = p.LocationTypes
	}
	return m
}

// GetAdaptiveScope runs the Get-AdaptiveScope cmdlet.
func (s *Service) GetAdaptiveScope(ctx context.Context, p GetAdaptiveScopeParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-AdaptiveScope", p.params())
}

// GetAdaptiveScopeMembersParams are the parameters of Get-AdaptiveScopeMembers.
// DefaultParameterSetName: Identity
type GetAdaptiveScopeMembersParams struct {
	AdaptiveReportFilters any    `ps:"AdaptiveReportFilters"`
	AdaptiveReportQuery   string `ps:"AdaptiveReportQuery"`
	EndDateTime           any    `ps:"EndDateTime"`
	Identity              any    `ps:"Identity"`
	PageCookie            string `ps:"PageCookie"`
	PageResultSize        any    `ps:"PageResultSize"`
	StartDateTime         any    `ps:"StartDateTime"`
	State                 any    `ps:"State"`
}

func (p GetAdaptiveScopeMembersParams) params() map[string]any {
	m := map[string]any{}
	if p.AdaptiveReportFilters != nil {
		m["AdaptiveReportFilters"] = p.AdaptiveReportFilters
	}
	if p.AdaptiveReportQuery != "" {
		m["AdaptiveReportQuery"] = p.AdaptiveReportQuery
	}
	if p.EndDateTime != nil {
		m["EndDateTime"] = p.EndDateTime
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.PageCookie != "" {
		m["PageCookie"] = p.PageCookie
	}
	if p.PageResultSize != nil {
		m["PageResultSize"] = p.PageResultSize
	}
	if p.StartDateTime != nil {
		m["StartDateTime"] = p.StartDateTime
	}
	if p.State != nil {
		m["State"] = p.State
	}
	return m
}

// GetAdaptiveScopeMembers runs the Get-AdaptiveScopeMembers cmdlet.
func (s *Service) GetAdaptiveScopeMembers(ctx context.Context, p GetAdaptiveScopeMembersParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-AdaptiveScopeMembers", p.params())
}

// GetAdminAuditLogConfigParams are the parameters of Get-AdminAuditLogConfig.
type GetAdminAuditLogConfigParams struct {
}

func (p GetAdminAuditLogConfigParams) params() map[string]any {
	m := map[string]any{}
	return m
}

// GetAdminAuditLogConfig runs the Get-AdminAuditLogConfig cmdlet.
func (s *Service) GetAdminAuditLogConfig(ctx context.Context, p GetAdminAuditLogConfigParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-AdminAuditLogConfig", p.params())
}

// GetAdministrativeUnitExtensionParams are the parameters of Get-AdministrativeUnitExtension.
// DefaultParameterSetName: Identity
type GetAdministrativeUnitExtensionParams struct {
	AdministrativeUnits any `ps:"AdministrativeUnits"`
	Identity            any `ps:"Identity"`
}

func (p GetAdministrativeUnitExtensionParams) params() map[string]any {
	m := map[string]any{}
	if p.AdministrativeUnits != nil {
		m["AdministrativeUnits"] = p.AdministrativeUnits
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetAdministrativeUnitExtension runs the Get-AdministrativeUnitExtension cmdlet.
func (s *Service) GetAdministrativeUnitExtension(ctx context.Context, p GetAdministrativeUnitExtensionParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-AdministrativeUnitExtension", p.params())
}

// GetAppRetentionCompliancePolicyParams are the parameters of Get-AppRetentionCompliancePolicy.
// DefaultParameterSetName: Identity
type GetAppRetentionCompliancePolicyParams struct {
	DistributionDetail bool `ps:"DistributionDetail"`
	ErrorPolicyOnly    bool `ps:"ErrorPolicyOnly"`
	Identity           any  `ps:"Identity"`
	RetentionRuleTypes bool `ps:"RetentionRuleTypes"`
}

func (p GetAppRetentionCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.DistributionDetail {
		m["DistributionDetail"] = true
	}
	if p.ErrorPolicyOnly {
		m["ErrorPolicyOnly"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.RetentionRuleTypes {
		m["RetentionRuleTypes"] = true
	}
	return m
}

// GetAppRetentionCompliancePolicy runs the Get-AppRetentionCompliancePolicy cmdlet.
func (s *Service) GetAppRetentionCompliancePolicy(ctx context.Context, p GetAppRetentionCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-AppRetentionCompliancePolicy", p.params())
}

// GetAppRetentionComplianceRuleParams are the parameters of Get-AppRetentionComplianceRule.
// DefaultParameterSetName: Identity
type GetAppRetentionComplianceRuleParams struct {
	Identity any `ps:"Identity"`
	Policy   any `ps:"Policy"`
}

func (p GetAppRetentionComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	return m
}

// GetAppRetentionComplianceRule runs the Get-AppRetentionComplianceRule cmdlet.
func (s *Service) GetAppRetentionComplianceRule(ctx context.Context, p GetAppRetentionComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-AppRetentionComplianceRule", p.params())
}

// GetAuditConfigParams are the parameters of Get-AuditConfig.
type GetAuditConfigParams struct {
	DomainController any `ps:"DomainController"`
}

func (p GetAuditConfigParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	return m
}

// GetAuditConfig runs the Get-AuditConfig cmdlet.
func (s *Service) GetAuditConfig(ctx context.Context, p GetAuditConfigParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-AuditConfig", p.params())
}

// GetAuditConfigurationPolicyParams are the parameters of Get-AuditConfigurationPolicy.
// DefaultParameterSetName: Identity
type GetAuditConfigurationPolicyParams struct {
	DomainController any `ps:"DomainController"`
	Identity         any `ps:"Identity"`
}

func (p GetAuditConfigurationPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetAuditConfigurationPolicy runs the Get-AuditConfigurationPolicy cmdlet.
func (s *Service) GetAuditConfigurationPolicy(ctx context.Context, p GetAuditConfigurationPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-AuditConfigurationPolicy", p.params())
}

// GetAuditConfigurationRuleParams are the parameters of Get-AuditConfigurationRule.
// DefaultParameterSetName: Identity
type GetAuditConfigurationRuleParams struct {
	DomainController any `ps:"DomainController"`
	Identity         any `ps:"Identity"`
}

func (p GetAuditConfigurationRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetAuditConfigurationRule runs the Get-AuditConfigurationRule cmdlet.
func (s *Service) GetAuditConfigurationRule(ctx context.Context, p GetAuditConfigurationRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-AuditConfigurationRule", p.params())
}

// GetAutoSensitivityLabelPolicyParams are the parameters of Get-AutoSensitivityLabelPolicy.
// DefaultParameterSetName: Identity
type GetAutoSensitivityLabelPolicyParams struct {
	DistributionDetail          bool `ps:"DistributionDetail"`
	ForceValidate               bool `ps:"ForceValidate"`
	Identity                    any  `ps:"Identity"`
	IncludePolicyGradingDetails bool `ps:"IncludePolicyGradingDetails"`
	IncludePolicyGradingResults bool `ps:"IncludePolicyGradingResults"`
	IncludeProgressFeedback     bool `ps:"IncludeProgressFeedback"`
	IncludeTestModeResults      bool `ps:"IncludeTestModeResults"`
}

func (p GetAutoSensitivityLabelPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.DistributionDetail {
		m["DistributionDetail"] = true
	}
	if p.ForceValidate {
		m["ForceValidate"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IncludePolicyGradingDetails {
		m["IncludePolicyGradingDetails"] = true
	}
	if p.IncludePolicyGradingResults {
		m["IncludePolicyGradingResults"] = true
	}
	if p.IncludeProgressFeedback {
		m["IncludeProgressFeedback"] = true
	}
	if p.IncludeTestModeResults {
		m["IncludeTestModeResults"] = true
	}
	return m
}

// GetAutoSensitivityLabelPolicy runs the Get-AutoSensitivityLabelPolicy cmdlet.
func (s *Service) GetAutoSensitivityLabelPolicy(ctx context.Context, p GetAutoSensitivityLabelPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-AutoSensitivityLabelPolicy", p.params())
}

// GetAutoSensitivityLabelRuleParams are the parameters of Get-AutoSensitivityLabelRule.
type GetAutoSensitivityLabelRuleParams struct {
	ForceValidate                   bool `ps:"ForceValidate"`
	Identity                        any  `ps:"Identity"`
	IncludeExecutionRuleGuids       bool `ps:"IncludeExecutionRuleGuids"`
	IncludeExecutionRuleInformation bool `ps:"IncludeExecutionRuleInformation"`
	Policy                          any  `ps:"Policy"`
}

func (p GetAutoSensitivityLabelRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceValidate {
		m["ForceValidate"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IncludeExecutionRuleGuids {
		m["IncludeExecutionRuleGuids"] = true
	}
	if p.IncludeExecutionRuleInformation {
		m["IncludeExecutionRuleInformation"] = true
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	return m
}

// GetAutoSensitivityLabelRule runs the Get-AutoSensitivityLabelRule cmdlet.
func (s *Service) GetAutoSensitivityLabelRule(ctx context.Context, p GetAutoSensitivityLabelRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-AutoSensitivityLabelRule", p.params())
}

// GetCaseHoldPolicyParams are the parameters of Get-CaseHoldPolicy.
// DefaultParameterSetName: Identity
type GetCaseHoldPolicyParams struct {
	Case                string `ps:"Case"`
	DistributionDetail  bool   `ps:"DistributionDetail"`
	Identity            any    `ps:"Identity"`
	IncludeBindings     bool   `ps:"IncludeBindings"`
	IncludeBindingsOnly bool   `ps:"IncludeBindingsOnly"`
}

func (p GetCaseHoldPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Case != "" {
		m["Case"] = p.Case
	}
	if p.DistributionDetail {
		m["DistributionDetail"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IncludeBindings {
		m["IncludeBindings"] = true
	}
	if p.IncludeBindingsOnly {
		m["IncludeBindingsOnly"] = true
	}
	return m
}

// GetCaseHoldPolicy runs the Get-CaseHoldPolicy cmdlet.
func (s *Service) GetCaseHoldPolicy(ctx context.Context, p GetCaseHoldPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-CaseHoldPolicy", p.params())
}

// GetCaseHoldRuleParams are the parameters of Get-CaseHoldRule.
// DefaultParameterSetName: Identity
type GetCaseHoldRuleParams struct {
	Identity any `ps:"Identity"`
	Policy   any `ps:"Policy"`
}

func (p GetCaseHoldRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	return m
}

// GetCaseHoldRule runs the Get-CaseHoldRule cmdlet.
func (s *Service) GetCaseHoldRule(ctx context.Context, p GetCaseHoldRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-CaseHoldRule", p.params())
}

// GetClassificationGradingPolicyParams are the parameters of Get-ClassificationGradingPolicy.
type GetClassificationGradingPolicyParams struct {
	Identity any `ps:"Identity"`
}

func (p GetClassificationGradingPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetClassificationGradingPolicy runs the Get-ClassificationGradingPolicy cmdlet.
func (s *Service) GetClassificationGradingPolicy(ctx context.Context, p GetClassificationGradingPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ClassificationGradingPolicy", p.params())
}

// GetComplianceBoundaryParams are the parameters of Get-ComplianceBoundary.
// DefaultParameterSetName: Identity
type GetComplianceBoundaryParams struct {
	Identity any `ps:"Identity"`
}

func (p GetComplianceBoundaryParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetComplianceBoundary runs the Get-ComplianceBoundary cmdlet.
func (s *Service) GetComplianceBoundary(ctx context.Context, p GetComplianceBoundaryParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ComplianceBoundary", p.params())
}

// GetComplianceCaseParams are the parameters of Get-ComplianceCase.
// DefaultParameterSetName: Identity
type GetComplianceCaseParams struct {
	CaseType         any    `ps:"CaseType"`
	DomainController any    `ps:"DomainController"`
	Identity         any    `ps:"Identity"`
	RecentOnly       bool   `ps:"RecentOnly"`
	RoleGroup        string `ps:"RoleGroup"`
}

func (p GetComplianceCaseParams) params() map[string]any {
	m := map[string]any{}
	if p.CaseType != nil {
		m["CaseType"] = p.CaseType
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.RecentOnly {
		m["RecentOnly"] = true
	}
	if p.RoleGroup != "" {
		m["RoleGroup"] = p.RoleGroup
	}
	return m
}

// GetComplianceCase runs the Get-ComplianceCase cmdlet.
func (s *Service) GetComplianceCase(ctx context.Context, p GetComplianceCaseParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ComplianceCase", p.params())
}

// GetComplianceCaseMemberParams are the parameters of Get-ComplianceCaseMember.
// DefaultParameterSetName: Identity
type GetComplianceCaseMemberParams struct {
	Case             string `ps:"Case"`
	DomainController any    `ps:"DomainController"`
	ResultSize       any    `ps:"ResultSize"`
	RoleGroupOnly    bool   `ps:"RoleGroupOnly"` // one of: True
	ShowCaseAdmin    bool   `ps:"ShowCaseAdmin"`
}

func (p GetComplianceCaseMemberParams) params() map[string]any {
	m := map[string]any{}
	if p.Case != "" {
		m["Case"] = p.Case
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.ResultSize != nil {
		m["ResultSize"] = p.ResultSize
	}
	if p.RoleGroupOnly {
		m["RoleGroupOnly"] = true
	}
	if p.ShowCaseAdmin {
		m["ShowCaseAdmin"] = true
	}
	return m
}

// GetComplianceCaseMember runs the Get-ComplianceCaseMember cmdlet.
func (s *Service) GetComplianceCaseMember(ctx context.Context, p GetComplianceCaseMemberParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ComplianceCaseMember", p.params())
}

// GetComplianceCaseMemberCandidateParams are the parameters of Get-ComplianceCaseMemberCandidate.
// DefaultParameterSetName: Identity
type GetComplianceCaseMemberCandidateParams struct {
	DomainController any  `ps:"DomainController"`
	ResultSize       any  `ps:"ResultSize"`
	RoleGroupOnly    bool `ps:"RoleGroupOnly"`
}

func (p GetComplianceCaseMemberCandidateParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.ResultSize != nil {
		m["ResultSize"] = p.ResultSize
	}
	if p.RoleGroupOnly {
		m["RoleGroupOnly"] = true
	}
	return m
}

// GetComplianceCaseMemberCandidate runs the Get-ComplianceCaseMemberCandidate cmdlet.
func (s *Service) GetComplianceCaseMemberCandidate(ctx context.Context, p GetComplianceCaseMemberCandidateParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ComplianceCaseMemberCandidate", p.params())
}

// GetComplianceCaseStatisticsParams are the parameters of Get-ComplianceCaseStatistics.
// DefaultParameterSetName: Identity
type GetComplianceCaseStatisticsParams struct {
	CaseType any `ps:"CaseType"`
}

func (p GetComplianceCaseStatisticsParams) params() map[string]any {
	m := map[string]any{}
	if p.CaseType != nil {
		m["CaseType"] = p.CaseType
	}
	return m
}

// GetComplianceCaseStatistics runs the Get-ComplianceCaseStatistics cmdlet.
func (s *Service) GetComplianceCaseStatistics(ctx context.Context, p GetComplianceCaseStatisticsParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ComplianceCaseStatistics", p.params())
}

// GetComplianceCustodianParams are the parameters of Get-ComplianceCustodian.
// DefaultParameterSetName: Identity
type GetComplianceCustodianParams struct {
	Case     string `ps:"Case"`
	Identity any    `ps:"Identity"`
}

func (p GetComplianceCustodianParams) params() map[string]any {
	m := map[string]any{}
	if p.Case != "" {
		m["Case"] = p.Case
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetComplianceCustodian runs the Get-ComplianceCustodian cmdlet.
func (s *Service) GetComplianceCustodian(ctx context.Context, p GetComplianceCustodianParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ComplianceCustodian", p.params())
}

// GetComplianceEmailSettingParams are the parameters of Get-ComplianceEmailSetting.
// DefaultParameterSetName: Identity
type GetComplianceEmailSettingParams struct {
	Identity any    `ps:"Identity"`
	NoticeId string `ps:"NoticeId"`
}

func (p GetComplianceEmailSettingParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.NoticeId != "" {
		m["NoticeId"] = p.NoticeId
	}
	return m
}

// GetComplianceEmailSetting runs the Get-ComplianceEmailSetting cmdlet.
func (s *Service) GetComplianceEmailSetting(ctx context.Context, p GetComplianceEmailSettingParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ComplianceEmailSetting", p.params())
}

// GetComplianceNoticeParams are the parameters of Get-ComplianceNotice.
// DefaultParameterSetName: Identity
type GetComplianceNoticeParams struct {
	Case     string `ps:"Case"`
	Identity any    `ps:"Identity"`
}

func (p GetComplianceNoticeParams) params() map[string]any {
	m := map[string]any{}
	if p.Case != "" {
		m["Case"] = p.Case
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetComplianceNotice runs the Get-ComplianceNotice cmdlet.
func (s *Service) GetComplianceNotice(ctx context.Context, p GetComplianceNoticeParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ComplianceNotice", p.params())
}

// GetCompliancePreservationSettingParams are the parameters of Get-CompliancePreservationSetting.
type GetCompliancePreservationSettingParams struct {
}

func (p GetCompliancePreservationSettingParams) params() map[string]any {
	m := map[string]any{}
	return m
}

// GetCompliancePreservationSetting runs the Get-CompliancePreservationSetting cmdlet.
func (s *Service) GetCompliancePreservationSetting(ctx context.Context, p GetCompliancePreservationSettingParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-CompliancePreservationSetting", p.params())
}

// GetComplianceRetentionEventParams are the parameters of Get-ComplianceRetentionEvent.
// DefaultParameterSetName: Identity
type GetComplianceRetentionEventParams struct {
	BeginDateTime    any  `ps:"BeginDateTime"`
	DomainController any  `ps:"DomainController"`
	EndDateTime      any  `ps:"EndDateTime"`
	Identity         any  `ps:"Identity"`
	PreviewOnly      bool `ps:"PreviewOnly"`
}

func (p GetComplianceRetentionEventParams) params() map[string]any {
	m := map[string]any{}
	if p.BeginDateTime != nil {
		m["BeginDateTime"] = p.BeginDateTime
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.EndDateTime != nil {
		m["EndDateTime"] = p.EndDateTime
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.PreviewOnly {
		m["PreviewOnly"] = true
	}
	return m
}

// GetComplianceRetentionEvent runs the Get-ComplianceRetentionEvent cmdlet.
func (s *Service) GetComplianceRetentionEvent(ctx context.Context, p GetComplianceRetentionEventParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ComplianceRetentionEvent", p.params())
}

// GetComplianceRetentionEventTypeParams are the parameters of Get-ComplianceRetentionEventType.
// DefaultParameterSetName: Identity
type GetComplianceRetentionEventTypeParams struct {
	Identity any  `ps:"Identity"`
	LoadTag  bool `ps:"LoadTag"`
}

func (p GetComplianceRetentionEventTypeParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.LoadTag {
		m["LoadTag"] = true
	}
	return m
}

// GetComplianceRetentionEventType runs the Get-ComplianceRetentionEventType cmdlet.
func (s *Service) GetComplianceRetentionEventType(ctx context.Context, p GetComplianceRetentionEventTypeParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ComplianceRetentionEventType", p.params())
}

// GetComplianceSearchParams are the parameters of Get-ComplianceSearch.
// DefaultParameterSetName: Identity
type GetComplianceSearchParams struct {
	Case       string `ps:"Case"`
	Identity   any    `ps:"Identity"`
	ResultSize any    `ps:"ResultSize"`
}

func (p GetComplianceSearchParams) params() map[string]any {
	m := map[string]any{}
	if p.Case != "" {
		m["Case"] = p.Case
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.ResultSize != nil {
		m["ResultSize"] = p.ResultSize
	}
	return m
}

// GetComplianceSearch runs the Get-ComplianceSearch cmdlet.
func (s *Service) GetComplianceSearch(ctx context.Context, p GetComplianceSearchParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ComplianceSearch", p.params())
}

// GetComplianceSearchActionParams are the parameters of Get-ComplianceSearchAction.
// DefaultParameterSetName: Identity
type GetComplianceSearchActionParams struct {
	Case              string `ps:"Case"`
	Details           bool   `ps:"Details"`
	Export            bool   `ps:"Export"`
	Identity          any    `ps:"Identity"`
	IncludeCredential bool   `ps:"IncludeCredential"`
	Preview           bool   `ps:"Preview"`
	ResultSize        any    `ps:"ResultSize"`
}

func (p GetComplianceSearchActionParams) params() map[string]any {
	m := map[string]any{}
	if p.Case != "" {
		m["Case"] = p.Case
	}
	if p.Details {
		m["Details"] = true
	}
	if p.Export {
		m["Export"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IncludeCredential {
		m["IncludeCredential"] = true
	}
	if p.Preview {
		m["Preview"] = true
	}
	if p.ResultSize != nil {
		m["ResultSize"] = p.ResultSize
	}
	return m
}

// GetComplianceSearchAction runs the Get-ComplianceSearchAction cmdlet.
func (s *Service) GetComplianceSearchAction(ctx context.Context, p GetComplianceSearchActionParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ComplianceSearchAction", p.params())
}

// GetComplianceSecurityFilterParams are the parameters of Get-ComplianceSecurityFilter.
// DefaultParameterSetName: Identity
type GetComplianceSecurityFilterParams struct {
	User string `ps:"User"`
}

func (p GetComplianceSecurityFilterParams) params() map[string]any {
	m := map[string]any{}
	if p.User != "" {
		m["User"] = p.User
	}
	return m
}

// GetComplianceSecurityFilter runs the Get-ComplianceSecurityFilter cmdlet.
func (s *Service) GetComplianceSecurityFilter(ctx context.Context, p GetComplianceSecurityFilterParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ComplianceSecurityFilter", p.params())
}

// GetComplianceTagParams are the parameters of Get-ComplianceTag.
// DefaultParameterSetName: Identity
type GetComplianceTagParams struct {
	Identity            any  `ps:"Identity"`
	IncludingLabelState bool `ps:"IncludingLabelState"`
}

func (p GetComplianceTagParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IncludingLabelState {
		m["IncludingLabelState"] = true
	}
	return m
}

// GetComplianceTag runs the Get-ComplianceTag cmdlet.
func (s *Service) GetComplianceTag(ctx context.Context, p GetComplianceTagParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ComplianceTag", p.params())
}

// GetComplianceTagStorageParams are the parameters of Get-ComplianceTagStorage.
// DefaultParameterSetName: Identity
type GetComplianceTagStorageParams struct {
	Identity any `ps:"Identity"`
}

func (p GetComplianceTagStorageParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetComplianceTagStorage runs the Get-ComplianceTagStorage cmdlet.
func (s *Service) GetComplianceTagStorage(ctx context.Context, p GetComplianceTagStorageParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ComplianceTagStorage", p.params())
}

// GetCustomDlpEmailTemplatesParams are the parameters of Get-CustomDlpEmailTemplates.
// DefaultParameterSetName: Identity
type GetCustomDlpEmailTemplatesParams struct {
}

func (p GetCustomDlpEmailTemplatesParams) params() map[string]any {
	m := map[string]any{}
	return m
}

// GetCustomDlpEmailTemplates runs the Get-CustomDlpEmailTemplates cmdlet.
func (s *Service) GetCustomDlpEmailTemplates(ctx context.Context, p GetCustomDlpEmailTemplatesParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-CustomDlpEmailTemplates", p.params())
}

// GetDataRetentionReportParams are the parameters of Get-DataRetentionReport.
type GetDataRetentionReportParams struct {
	EndDate   any `ps:"EndDate"`
	Page      int `ps:"Page"`
	PageSize  int `ps:"PageSize"`
	StartDate any `ps:"StartDate"`
}

func (p GetDataRetentionReportParams) params() map[string]any {
	m := map[string]any{}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	if p.Page != 0 {
		m["Page"] = p.Page
	}
	if p.PageSize != 0 {
		m["PageSize"] = p.PageSize
	}
	if p.StartDate != nil {
		m["StartDate"] = p.StartDate
	}
	return m
}

// GetDataRetentionReport runs the Get-DataRetentionReport cmdlet.
func (s *Service) GetDataRetentionReport(ctx context.Context, p GetDataRetentionReportParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DataRetentionReport", p.params())
}

// GetDeviceComplianceDetailsReportParams are the parameters of Get-DeviceComplianceDetailsReport.
type GetDeviceComplianceDetailsReportParams struct {
	DeviceCompliancePolicy any    `ps:"DeviceCompliancePolicy"`
	DeviceComplianceStatus string `ps:"DeviceComplianceStatus"`
	DevicePlatform         string `ps:"DevicePlatform"`
	DeviceUserName         string `ps:"DeviceUserName"`
	EndDate                any    `ps:"EndDate"`
	Expression             any    `ps:"Expression"`
}

func (p GetDeviceComplianceDetailsReportParams) params() map[string]any {
	m := map[string]any{}
	if p.DeviceCompliancePolicy != nil {
		m["DeviceCompliancePolicy"] = p.DeviceCompliancePolicy
	}
	if p.DeviceComplianceStatus != "" {
		m["DeviceComplianceStatus"] = p.DeviceComplianceStatus
	}
	if p.DevicePlatform != "" {
		m["DevicePlatform"] = p.DevicePlatform
	}
	if p.DeviceUserName != "" {
		m["DeviceUserName"] = p.DeviceUserName
	}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	if p.Expression != nil {
		m["Expression"] = p.Expression
	}
	return m
}

// GetDeviceComplianceDetailsReport runs the Get-DeviceComplianceDetailsReport cmdlet.
func (s *Service) GetDeviceComplianceDetailsReport(ctx context.Context, p GetDeviceComplianceDetailsReportParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DeviceComplianceDetailsReport", p.params())
}

// GetDeviceComplianceDetailsReportFilterParams are the parameters of Get-DeviceComplianceDetailsReportFilter.
type GetDeviceComplianceDetailsReportFilterParams struct {
	EndDate any `ps:"EndDate"`
}

func (p GetDeviceComplianceDetailsReportFilterParams) params() map[string]any {
	m := map[string]any{}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	return m
}

// GetDeviceComplianceDetailsReportFilter runs the Get-DeviceComplianceDetailsReportFilter cmdlet.
func (s *Service) GetDeviceComplianceDetailsReportFilter(ctx context.Context, p GetDeviceComplianceDetailsReportFilterParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DeviceComplianceDetailsReportFilter", p.params())
}

// GetDeviceCompliancePolicyInventoryParams are the parameters of Get-DeviceCompliancePolicyInventory.
type GetDeviceCompliancePolicyInventoryParams struct {
	EndDate    any `ps:"EndDate"`
	Expression any `ps:"Expression"`
}

func (p GetDeviceCompliancePolicyInventoryParams) params() map[string]any {
	m := map[string]any{}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	if p.Expression != nil {
		m["Expression"] = p.Expression
	}
	return m
}

// GetDeviceCompliancePolicyInventory runs the Get-DeviceCompliancePolicyInventory cmdlet.
func (s *Service) GetDeviceCompliancePolicyInventory(ctx context.Context, p GetDeviceCompliancePolicyInventoryParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DeviceCompliancePolicyInventory", p.params())
}

// GetDeviceComplianceReportDateParams are the parameters of Get-DeviceComplianceReportDate.
type GetDeviceComplianceReportDateParams struct {
	Expression any `ps:"Expression"`
}

func (p GetDeviceComplianceReportDateParams) params() map[string]any {
	m := map[string]any{}
	if p.Expression != nil {
		m["Expression"] = p.Expression
	}
	return m
}

// GetDeviceComplianceReportDate runs the Get-DeviceComplianceReportDate cmdlet.
func (s *Service) GetDeviceComplianceReportDate(ctx context.Context, p GetDeviceComplianceReportDateParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DeviceComplianceReportDate", p.params())
}

// GetDeviceComplianceSummaryReportParams are the parameters of Get-DeviceComplianceSummaryReport.
type GetDeviceComplianceSummaryReportParams struct {
	DeviceCompliancePolicy any    `ps:"DeviceCompliancePolicy"`
	DeviceComplianceStatus string `ps:"DeviceComplianceStatus"`
	DevicePlatform         string `ps:"DevicePlatform"`
	DeviceUserName         string `ps:"DeviceUserName"`
	EndDate                any    `ps:"EndDate"`
	Expression             any    `ps:"Expression"`
}

func (p GetDeviceComplianceSummaryReportParams) params() map[string]any {
	m := map[string]any{}
	if p.DeviceCompliancePolicy != nil {
		m["DeviceCompliancePolicy"] = p.DeviceCompliancePolicy
	}
	if p.DeviceComplianceStatus != "" {
		m["DeviceComplianceStatus"] = p.DeviceComplianceStatus
	}
	if p.DevicePlatform != "" {
		m["DevicePlatform"] = p.DevicePlatform
	}
	if p.DeviceUserName != "" {
		m["DeviceUserName"] = p.DeviceUserName
	}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	if p.Expression != nil {
		m["Expression"] = p.Expression
	}
	return m
}

// GetDeviceComplianceSummaryReport runs the Get-DeviceComplianceSummaryReport cmdlet.
func (s *Service) GetDeviceComplianceSummaryReport(ctx context.Context, p GetDeviceComplianceSummaryReportParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DeviceComplianceSummaryReport", p.params())
}

// GetDeviceComplianceUserInventoryParams are the parameters of Get-DeviceComplianceUserInventory.
type GetDeviceComplianceUserInventoryParams struct {
	EndDate    any `ps:"EndDate"`
	Expression any `ps:"Expression"`
}

func (p GetDeviceComplianceUserInventoryParams) params() map[string]any {
	m := map[string]any{}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	if p.Expression != nil {
		m["Expression"] = p.Expression
	}
	return m
}

// GetDeviceComplianceUserInventory runs the Get-DeviceComplianceUserInventory cmdlet.
func (s *Service) GetDeviceComplianceUserInventory(ctx context.Context, p GetDeviceComplianceUserInventoryParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DeviceComplianceUserInventory", p.params())
}

// GetDeviceComplianceUserReportParams are the parameters of Get-DeviceComplianceUserReport.
type GetDeviceComplianceUserReportParams struct {
	DeviceId   any `ps:"DeviceId"`
	EndDate    any `ps:"EndDate"`
	Expression any `ps:"Expression"`
}

func (p GetDeviceComplianceUserReportParams) params() map[string]any {
	m := map[string]any{}
	if p.DeviceId != nil {
		m["DeviceId"] = p.DeviceId
	}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	if p.Expression != nil {
		m["Expression"] = p.Expression
	}
	return m
}

// GetDeviceComplianceUserReport runs the Get-DeviceComplianceUserReport cmdlet.
func (s *Service) GetDeviceComplianceUserReport(ctx context.Context, p GetDeviceComplianceUserReportParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DeviceComplianceUserReport", p.params())
}

// GetDeviceConditionalAccessPolicyParams are the parameters of Get-DeviceConditionalAccessPolicy.
// DefaultParameterSetName: Identity
type GetDeviceConditionalAccessPolicyParams struct {
	Identity any `ps:"Identity"`
}

func (p GetDeviceConditionalAccessPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetDeviceConditionalAccessPolicy runs the Get-DeviceConditionalAccessPolicy cmdlet.
func (s *Service) GetDeviceConditionalAccessPolicy(ctx context.Context, p GetDeviceConditionalAccessPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DeviceConditionalAccessPolicy", p.params())
}

// GetDeviceConditionalAccessRuleParams are the parameters of Get-DeviceConditionalAccessRule.
// DefaultParameterSetName: Identity
type GetDeviceConditionalAccessRuleParams struct {
	CompareToWorkload bool `ps:"CompareToWorkload"`
	DomainController  any  `ps:"DomainController"`
	Identity          any  `ps:"Identity"`
}

func (p GetDeviceConditionalAccessRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.CompareToWorkload {
		m["CompareToWorkload"] = true
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetDeviceConditionalAccessRule runs the Get-DeviceConditionalAccessRule cmdlet.
func (s *Service) GetDeviceConditionalAccessRule(ctx context.Context, p GetDeviceConditionalAccessRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DeviceConditionalAccessRule", p.params())
}

// GetDeviceConfigurationPolicyParams are the parameters of Get-DeviceConfigurationPolicy.
// DefaultParameterSetName: Identity
type GetDeviceConfigurationPolicyParams struct {
	Identity any `ps:"Identity"`
}

func (p GetDeviceConfigurationPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetDeviceConfigurationPolicy runs the Get-DeviceConfigurationPolicy cmdlet.
func (s *Service) GetDeviceConfigurationPolicy(ctx context.Context, p GetDeviceConfigurationPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DeviceConfigurationPolicy", p.params())
}

// GetDeviceConfigurationRuleParams are the parameters of Get-DeviceConfigurationRule.
// DefaultParameterSetName: Identity
type GetDeviceConfigurationRuleParams struct {
	CompareToWorkload bool `ps:"CompareToWorkload"`
	DomainController  any  `ps:"DomainController"`
	Identity          any  `ps:"Identity"`
}

func (p GetDeviceConfigurationRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.CompareToWorkload {
		m["CompareToWorkload"] = true
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetDeviceConfigurationRule runs the Get-DeviceConfigurationRule cmdlet.
func (s *Service) GetDeviceConfigurationRule(ctx context.Context, p GetDeviceConfigurationRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DeviceConfigurationRule", p.params())
}

// GetDevicePolicyParams are the parameters of Get-DevicePolicy.
// DefaultParameterSetName: Identity
type GetDevicePolicyParams struct {
	Identity any `ps:"Identity"`
}

func (p GetDevicePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetDevicePolicy runs the Get-DevicePolicy cmdlet.
func (s *Service) GetDevicePolicy(ctx context.Context, p GetDevicePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DevicePolicy", p.params())
}

// GetDeviceTenantPolicyParams are the parameters of Get-DeviceTenantPolicy.
// DefaultParameterSetName: Identity
type GetDeviceTenantPolicyParams struct {
	Identity any `ps:"Identity"`
}

func (p GetDeviceTenantPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetDeviceTenantPolicy runs the Get-DeviceTenantPolicy cmdlet.
func (s *Service) GetDeviceTenantPolicy(ctx context.Context, p GetDeviceTenantPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DeviceTenantPolicy", p.params())
}

// GetDeviceTenantRuleParams are the parameters of Get-DeviceTenantRule.
// DefaultParameterSetName: Identity
type GetDeviceTenantRuleParams struct {
	DomainController any `ps:"DomainController"`
	Identity         any `ps:"Identity"`
}

func (p GetDeviceTenantRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetDeviceTenantRule runs the Get-DeviceTenantRule cmdlet.
func (s *Service) GetDeviceTenantRule(ctx context.Context, p GetDeviceTenantRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DeviceTenantRule", p.params())
}

// GetDlpAlertTuningRuleParams are the parameters of Get-DlpAlertTuningRule.
// DefaultParameterSetName: Identity
type GetDlpAlertTuningRuleParams struct {
}

func (p GetDlpAlertTuningRuleParams) params() map[string]any {
	m := map[string]any{}
	return m
}

// GetDlpAlertTuningRule runs the Get-DlpAlertTuningRule cmdlet.
func (s *Service) GetDlpAlertTuningRule(ctx context.Context, p GetDlpAlertTuningRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DlpAlertTuningRule", p.params())
}

// GetDlpCompliancePolicyParams are the parameters of Get-DlpCompliancePolicy.
// DefaultParameterSetName: Identity
type GetDlpCompliancePolicyParams struct {
	DisplayName                  string `ps:"DisplayName"`
	DistributionDetail           bool   `ps:"DistributionDetail"`
	ForceValidate                bool   `ps:"ForceValidate"`
	Identity                     any    `ps:"Identity"`
	IncludeExtendedProperties    bool   `ps:"IncludeExtendedProperties"`
	IncludeRulesMetadata         bool   `ps:"IncludeRulesMetadata"`
	IncludeSimulationResults     bool   `ps:"IncludeSimulationResults"`
	IRMUserRiskConfiguredAnyRule bool   `ps:"IRMUserRiskConfiguredAnyRule"`
	Summary                      bool   `ps:"Summary"`
}

func (p GetDlpCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.DistributionDetail {
		m["DistributionDetail"] = true
	}
	if p.ForceValidate {
		m["ForceValidate"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IncludeExtendedProperties {
		m["IncludeExtendedProperties"] = true
	}
	if p.IncludeRulesMetadata {
		m["IncludeRulesMetadata"] = true
	}
	if p.IncludeSimulationResults {
		m["IncludeSimulationResults"] = true
	}
	if p.IRMUserRiskConfiguredAnyRule {
		m["IRMUserRiskConfiguredAnyRule"] = true
	}
	if p.Summary {
		m["Summary"] = true
	}
	return m
}

// GetDlpCompliancePolicy runs the Get-DlpCompliancePolicy cmdlet.
func (s *Service) GetDlpCompliancePolicy(ctx context.Context, p GetDlpCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DlpCompliancePolicy", p.params())
}

// GetDlpComplianceRuleParams are the parameters of Get-DlpComplianceRule.
// DefaultParameterSetName: Identity
type GetDlpComplianceRuleParams struct {
	DisplayName               string `ps:"DisplayName"`
	ForceValidate             bool   `ps:"ForceValidate"`
	Identity                  any    `ps:"Identity"`
	IncludeExecutionRuleGuids bool   `ps:"IncludeExecutionRuleGuids"`
	Policy                    any    `ps:"Policy"`
	Summary                   bool   `ps:"Summary"`
}

func (p GetDlpComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.ForceValidate {
		m["ForceValidate"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IncludeExecutionRuleGuids {
		m["IncludeExecutionRuleGuids"] = true
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	if p.Summary {
		m["Summary"] = true
	}
	return m
}

// GetDlpComplianceRule runs the Get-DlpComplianceRule cmdlet.
func (s *Service) GetDlpComplianceRule(ctx context.Context, p GetDlpComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DlpComplianceRule", p.params())
}

// GetDlpDetailReportParams are the parameters of Get-DlpDetailReport.
type GetDlpDetailReportParams struct {
	Action              any `ps:"Action"`
	Actor               any `ps:"Actor"`
	DlpCompliancePolicy any `ps:"DlpCompliancePolicy"`
	DlpComplianceRule   any `ps:"DlpComplianceRule"`
	EndDate             any `ps:"EndDate"`
	EventType           any `ps:"EventType"`
	Page                int `ps:"Page"`
	PageSize            int `ps:"PageSize"`
	Source              any `ps:"Source"`
	StartDate           any `ps:"StartDate"`
}

func (p GetDlpDetailReportParams) params() map[string]any {
	m := map[string]any{}
	if p.Action != nil {
		m["Action"] = p.Action
	}
	if p.Actor != nil {
		m["Actor"] = p.Actor
	}
	if p.DlpCompliancePolicy != nil {
		m["DlpCompliancePolicy"] = p.DlpCompliancePolicy
	}
	if p.DlpComplianceRule != nil {
		m["DlpComplianceRule"] = p.DlpComplianceRule
	}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	if p.EventType != nil {
		m["EventType"] = p.EventType
	}
	if p.Page != 0 {
		m["Page"] = p.Page
	}
	if p.PageSize != 0 {
		m["PageSize"] = p.PageSize
	}
	if p.Source != nil {
		m["Source"] = p.Source
	}
	if p.StartDate != nil {
		m["StartDate"] = p.StartDate
	}
	return m
}

// GetDlpDetailReport runs the Get-DlpDetailReport cmdlet.
func (s *Service) GetDlpDetailReport(ctx context.Context, p GetDlpDetailReportParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DlpDetailReport", p.params())
}

// GetDlpDetectionsReportParams are the parameters of Get-DlpDetectionsReport.
type GetDlpDetectionsReportParams struct {
	Action              any    `ps:"Action"`
	AggregateBy         string `ps:"AggregateBy"`
	DlpCompliancePolicy any    `ps:"DlpCompliancePolicy"`
	DlpComplianceRule   any    `ps:"DlpComplianceRule"`
	EndDate             any    `ps:"EndDate"`
	EventType           any    `ps:"EventType"`
	Expression          any    `ps:"Expression"`
	Page                int    `ps:"Page"`
	PageSize            int    `ps:"PageSize"`
	Source              any    `ps:"Source"`
	StartDate           any    `ps:"StartDate"`
	SummarizeBy         any    `ps:"SummarizeBy"`
}

func (p GetDlpDetectionsReportParams) params() map[string]any {
	m := map[string]any{}
	if p.Action != nil {
		m["Action"] = p.Action
	}
	if p.AggregateBy != "" {
		m["AggregateBy"] = p.AggregateBy
	}
	if p.DlpCompliancePolicy != nil {
		m["DlpCompliancePolicy"] = p.DlpCompliancePolicy
	}
	if p.DlpComplianceRule != nil {
		m["DlpComplianceRule"] = p.DlpComplianceRule
	}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	if p.EventType != nil {
		m["EventType"] = p.EventType
	}
	if p.Expression != nil {
		m["Expression"] = p.Expression
	}
	if p.Page != 0 {
		m["Page"] = p.Page
	}
	if p.PageSize != 0 {
		m["PageSize"] = p.PageSize
	}
	if p.Source != nil {
		m["Source"] = p.Source
	}
	if p.StartDate != nil {
		m["StartDate"] = p.StartDate
	}
	if p.SummarizeBy != nil {
		m["SummarizeBy"] = p.SummarizeBy
	}
	return m
}

// GetDlpDetectionsReport runs the Get-DlpDetectionsReport cmdlet.
func (s *Service) GetDlpDetectionsReport(ctx context.Context, p GetDlpDetectionsReportParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DlpDetectionsReport", p.params())
}

// GetDlpEdmSchemaParams are the parameters of Get-DlpEdmSchema.
// DefaultParameterSetName: Identity
type GetDlpEdmSchemaParams struct {
	Identity any `ps:"Identity"`
}

func (p GetDlpEdmSchemaParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetDlpEdmSchema runs the Get-DlpEdmSchema cmdlet.
func (s *Service) GetDlpEdmSchema(ctx context.Context, p GetDlpEdmSchemaParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DlpEdmSchema", p.params())
}

// GetDlpEdmSessionParams are the parameters of Get-DlpEdmSession.
type GetDlpEdmSessionParams struct {
	DataStoreName string `ps:"DataStoreName"`
	SessionId     string `ps:"SessionId"`
}

func (p GetDlpEdmSessionParams) params() map[string]any {
	m := map[string]any{}
	if p.DataStoreName != "" {
		m["DataStoreName"] = p.DataStoreName
	}
	if p.SessionId != "" {
		m["SessionId"] = p.SessionId
	}
	return m
}

// GetDlpEdmSession runs the Get-DlpEdmSession cmdlet.
func (s *Service) GetDlpEdmSession(ctx context.Context, p GetDlpEdmSessionParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DlpEdmSession", p.params())
}

// GetDlpIncidentDetailReportParams are the parameters of Get-DlpIncidentDetailReport.
type GetDlpIncidentDetailReportParams struct {
	Action              any `ps:"Action"`
	Actor               any `ps:"Actor"`
	DlpCompliancePolicy any `ps:"DlpCompliancePolicy"`
	DlpComplianceRule   any `ps:"DlpComplianceRule"`
	EndDate             any `ps:"EndDate"`
	EventType           any `ps:"EventType"`
	Page                int `ps:"Page"`
	PageSize            int `ps:"PageSize"`
	Source              any `ps:"Source"`
	StartDate           any `ps:"StartDate"`
}

func (p GetDlpIncidentDetailReportParams) params() map[string]any {
	m := map[string]any{}
	if p.Action != nil {
		m["Action"] = p.Action
	}
	if p.Actor != nil {
		m["Actor"] = p.Actor
	}
	if p.DlpCompliancePolicy != nil {
		m["DlpCompliancePolicy"] = p.DlpCompliancePolicy
	}
	if p.DlpComplianceRule != nil {
		m["DlpComplianceRule"] = p.DlpComplianceRule
	}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	if p.EventType != nil {
		m["EventType"] = p.EventType
	}
	if p.Page != 0 {
		m["Page"] = p.Page
	}
	if p.PageSize != 0 {
		m["PageSize"] = p.PageSize
	}
	if p.Source != nil {
		m["Source"] = p.Source
	}
	if p.StartDate != nil {
		m["StartDate"] = p.StartDate
	}
	return m
}

// GetDlpIncidentDetailReport runs the Get-DlpIncidentDetailReport cmdlet.
func (s *Service) GetDlpIncidentDetailReport(ctx context.Context, p GetDlpIncidentDetailReportParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DlpIncidentDetailReport", p.params())
}

// GetDlpKeywordDictionaryParams are the parameters of Get-DlpKeywordDictionary.
type GetDlpKeywordDictionaryParams struct {
	Name string `ps:"Name"`
}

func (p GetDlpKeywordDictionaryParams) params() map[string]any {
	m := map[string]any{}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	return m
}

// GetDlpKeywordDictionary runs the Get-DlpKeywordDictionary cmdlet.
func (s *Service) GetDlpKeywordDictionary(ctx context.Context, p GetDlpKeywordDictionaryParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DlpKeywordDictionary", p.params())
}

// GetDlpSensitiveInformationTypeParams are the parameters of Get-DlpSensitiveInformationType.
type GetDlpSensitiveInformationTypeParams struct {
	Capability      any  `ps:"Capability"`
	Identity        any  `ps:"Identity"`
	IncludeDetails  bool `ps:"IncludeDetails"`
	IncludeElements any  `ps:"IncludeElements"`
	Organization    any  `ps:"Organization"`
}

func (p GetDlpSensitiveInformationTypeParams) params() map[string]any {
	m := map[string]any{}
	if p.Capability != nil {
		m["Capability"] = p.Capability
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IncludeDetails {
		m["IncludeDetails"] = true
	}
	if p.IncludeElements != nil {
		m["IncludeElements"] = p.IncludeElements
	}
	if p.Organization != nil {
		m["Organization"] = p.Organization
	}
	return m
}

// GetDlpSensitiveInformationType runs the Get-DlpSensitiveInformationType cmdlet.
func (s *Service) GetDlpSensitiveInformationType(ctx context.Context, p GetDlpSensitiveInformationTypeParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DlpSensitiveInformationType", p.params())
}

// GetDlpSensitiveInformationTypeConfigParams are the parameters of Get-DlpSensitiveInformationTypeConfig.
type GetDlpSensitiveInformationTypeConfigParams struct {
}

func (p GetDlpSensitiveInformationTypeConfigParams) params() map[string]any {
	m := map[string]any{}
	return m
}

// GetDlpSensitiveInformationTypeConfig runs the Get-DlpSensitiveInformationTypeConfig cmdlet.
func (s *Service) GetDlpSensitiveInformationTypeConfig(ctx context.Context, p GetDlpSensitiveInformationTypeConfigParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DlpSensitiveInformationTypeConfig", p.params())
}

// GetDlpSensitiveInformationTypeRulePackageParams are the parameters of Get-DlpSensitiveInformationTypeRulePackage.
// DefaultParameterSetName: Identity
type GetDlpSensitiveInformationTypeRulePackageParams struct {
	Capability any `ps:"Capability"`
	Identity   any `ps:"Identity"`
}

func (p GetDlpSensitiveInformationTypeRulePackageParams) params() map[string]any {
	m := map[string]any{}
	if p.Capability != nil {
		m["Capability"] = p.Capability
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetDlpSensitiveInformationTypeRulePackage runs the Get-DlpSensitiveInformationTypeRulePackage cmdlet.
func (s *Service) GetDlpSensitiveInformationTypeRulePackage(ctx context.Context, p GetDlpSensitiveInformationTypeRulePackageParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DlpSensitiveInformationTypeRulePackage", p.params())
}

// GetDlpSiDetectionsReportParams are the parameters of Get-DlpSiDetectionsReport.
type GetDlpSiDetectionsReportParams struct {
	EndDate       any `ps:"EndDate"`
	SensitiveType any `ps:"SensitiveType"`
}

func (p GetDlpSiDetectionsReportParams) params() map[string]any {
	m := map[string]any{}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	if p.SensitiveType != nil {
		m["SensitiveType"] = p.SensitiveType
	}
	return m
}

// GetDlpSiDetectionsReport runs the Get-DlpSiDetectionsReport cmdlet.
func (s *Service) GetDlpSiDetectionsReport(ctx context.Context, p GetDlpSiDetectionsReportParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DlpSiDetectionsReport", p.params())
}

// GetDspmPolicyParams are the parameters of Get-DspmPolicy.
// DefaultParameterSetName: Identity
type GetDspmPolicyParams struct {
	Identity any `ps:"Identity"`
}

func (p GetDspmPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetDspmPolicy runs the Get-DspmPolicy cmdlet.
func (s *Service) GetDspmPolicy(ctx context.Context, p GetDspmPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-DspmPolicy", p.params())
}

// GetFeatureConfigurationParams are the parameters of Get-FeatureConfiguration.
type GetFeatureConfigurationParams struct {
	FeatureScenario any `ps:"FeatureScenario"` // one of: KnowYourData, TrustContainer, ProtectExcludedFolders, UnsavedDocument, JustInTimeAudit
	Identity        any `ps:"Identity"`
}

func (p GetFeatureConfigurationParams) params() map[string]any {
	m := map[string]any{}
	if p.FeatureScenario != nil {
		m["FeatureScenario"] = p.FeatureScenario
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetFeatureConfiguration runs the Get-FeatureConfiguration cmdlet.
func (s *Service) GetFeatureConfiguration(ctx context.Context, p GetFeatureConfigurationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-FeatureConfiguration", p.params())
}

// GetFilePlanPropertyAuthorityParams are the parameters of Get-FilePlanPropertyAuthority.
// DefaultParameterSetName: Identity
type GetFilePlanPropertyAuthorityParams struct {
	Identity any `ps:"Identity"`
}

func (p GetFilePlanPropertyAuthorityParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetFilePlanPropertyAuthority runs the Get-FilePlanPropertyAuthority cmdlet.
func (s *Service) GetFilePlanPropertyAuthority(ctx context.Context, p GetFilePlanPropertyAuthorityParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-FilePlanPropertyAuthority", p.params())
}

// GetFilePlanPropertyCategoryParams are the parameters of Get-FilePlanPropertyCategory.
// DefaultParameterSetName: Identity
type GetFilePlanPropertyCategoryParams struct {
	Identity any `ps:"Identity"`
}

func (p GetFilePlanPropertyCategoryParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetFilePlanPropertyCategory runs the Get-FilePlanPropertyCategory cmdlet.
func (s *Service) GetFilePlanPropertyCategory(ctx context.Context, p GetFilePlanPropertyCategoryParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-FilePlanPropertyCategory", p.params())
}

// GetFilePlanPropertyCitationParams are the parameters of Get-FilePlanPropertyCitation.
// DefaultParameterSetName: Identity
type GetFilePlanPropertyCitationParams struct {
	Identity any `ps:"Identity"`
}

func (p GetFilePlanPropertyCitationParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetFilePlanPropertyCitation runs the Get-FilePlanPropertyCitation cmdlet.
func (s *Service) GetFilePlanPropertyCitation(ctx context.Context, p GetFilePlanPropertyCitationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-FilePlanPropertyCitation", p.params())
}

// GetFilePlanPropertyDepartmentParams are the parameters of Get-FilePlanPropertyDepartment.
// DefaultParameterSetName: Identity
type GetFilePlanPropertyDepartmentParams struct {
	Identity any `ps:"Identity"`
}

func (p GetFilePlanPropertyDepartmentParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetFilePlanPropertyDepartment runs the Get-FilePlanPropertyDepartment cmdlet.
func (s *Service) GetFilePlanPropertyDepartment(ctx context.Context, p GetFilePlanPropertyDepartmentParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-FilePlanPropertyDepartment", p.params())
}

// GetFilePlanPropertyReferenceIdParams are the parameters of Get-FilePlanPropertyReferenceId.
// DefaultParameterSetName: Identity
type GetFilePlanPropertyReferenceIdParams struct {
	Identity any `ps:"Identity"`
}

func (p GetFilePlanPropertyReferenceIdParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetFilePlanPropertyReferenceId runs the Get-FilePlanPropertyReferenceId cmdlet.
func (s *Service) GetFilePlanPropertyReferenceId(ctx context.Context, p GetFilePlanPropertyReferenceIdParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-FilePlanPropertyReferenceId", p.params())
}

// GetFilePlanPropertyStructureParams are the parameters of Get-FilePlanPropertyStructure.
type GetFilePlanPropertyStructureParams struct {
	IncludeAdditionalInfo bool `ps:"IncludeAdditionalInfo"`
}

func (p GetFilePlanPropertyStructureParams) params() map[string]any {
	m := map[string]any{}
	if p.IncludeAdditionalInfo {
		m["IncludeAdditionalInfo"] = true
	}
	return m
}

// GetFilePlanPropertyStructure runs the Get-FilePlanPropertyStructure cmdlet.
func (s *Service) GetFilePlanPropertyStructure(ctx context.Context, p GetFilePlanPropertyStructureParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-FilePlanPropertyStructure", p.params())
}

// GetFilePlanPropertySubCategoryParams are the parameters of Get-FilePlanPropertySubCategory.
// DefaultParameterSetName: Identity
type GetFilePlanPropertySubCategoryParams struct {
	Identity any `ps:"Identity"`
}

func (p GetFilePlanPropertySubCategoryParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetFilePlanPropertySubCategory runs the Get-FilePlanPropertySubCategory cmdlet.
func (s *Service) GetFilePlanPropertySubCategory(ctx context.Context, p GetFilePlanPropertySubCategoryParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-FilePlanPropertySubCategory", p.params())
}

// GetGlobalListParams are the parameters of Get-GlobalList.
// DefaultParameterSetName: Identity
type GetGlobalListParams struct {
	GlobalListType any `ps:"GlobalListType"`
	Identity       any `ps:"Identity"`
	Organization   any `ps:"Organization"`
}

func (p GetGlobalListParams) params() map[string]any {
	m := map[string]any{}
	if p.GlobalListType != nil {
		m["GlobalListType"] = p.GlobalListType
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Organization != nil {
		m["Organization"] = p.Organization
	}
	return m
}

// GetGlobalList runs the Get-GlobalList cmdlet.
func (s *Service) GetGlobalList(ctx context.Context, p GetGlobalListParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-GlobalList", p.params())
}

// GetGroupParams are the parameters of Get-Group.
// DefaultParameterSetName: Identity
type GetGroupParams struct {
	Identity             any      `ps:"Identity"`
	RecipientTypeDetails []string `ps:"RecipientTypeDetails"`
	ResultSize           any      `ps:"ResultSize"`
	SortBy               string   `ps:"SortBy"`
}

func (p GetGroupParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if len(p.RecipientTypeDetails) > 0 {
		m["RecipientTypeDetails"] = p.RecipientTypeDetails
	}
	if p.ResultSize != nil {
		m["ResultSize"] = p.ResultSize
	}
	if p.SortBy != "" {
		m["SortBy"] = p.SortBy
	}
	return m
}

// GetGroup runs the Get-Group cmdlet.
func (s *Service) GetGroup(ctx context.Context, p GetGroupParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-Group", p.params())
}

// GetHoldCompliancePolicyParams are the parameters of Get-HoldCompliancePolicy.
// DefaultParameterSetName: Identity
type GetHoldCompliancePolicyParams struct {
	DistributionDetail bool `ps:"DistributionDetail"`
	Identity           any  `ps:"Identity"`
}

func (p GetHoldCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.DistributionDetail {
		m["DistributionDetail"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetHoldCompliancePolicy runs the Get-HoldCompliancePolicy cmdlet.
func (s *Service) GetHoldCompliancePolicy(ctx context.Context, p GetHoldCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-HoldCompliancePolicy", p.params())
}

// GetHoldComplianceRuleParams are the parameters of Get-HoldComplianceRule.
// DefaultParameterSetName: Identity
type GetHoldComplianceRuleParams struct {
	Identity any `ps:"Identity"`
	Policy   any `ps:"Policy"`
}

func (p GetHoldComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	return m
}

// GetHoldComplianceRule runs the Get-HoldComplianceRule cmdlet.
func (s *Service) GetHoldComplianceRule(ctx context.Context, p GetHoldComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-HoldComplianceRule", p.params())
}

// GetInformationBarrierPoliciesApplicationStatusParams are the parameters of Get-InformationBarrierPoliciesApplicationStatus.
// DefaultParameterSetName: Default
type GetInformationBarrierPoliciesApplicationStatusParams struct {
	All      bool `ps:"All"`
	Identity any  `ps:"Identity"`
}

func (p GetInformationBarrierPoliciesApplicationStatusParams) params() map[string]any {
	m := map[string]any{}
	if p.All {
		m["All"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetInformationBarrierPoliciesApplicationStatus runs the Get-InformationBarrierPoliciesApplicationStatus cmdlet.
func (s *Service) GetInformationBarrierPoliciesApplicationStatus(ctx context.Context, p GetInformationBarrierPoliciesApplicationStatusParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-InformationBarrierPoliciesApplicationStatus", p.params())
}

// GetInformationBarrierPolicyParams are the parameters of Get-InformationBarrierPolicy.
// DefaultParameterSetName: InformationBarrierDefault
type GetInformationBarrierPolicyParams struct {
	ExoPolicyId any `ps:"ExoPolicyId"`
	Identity    any `ps:"Identity"`
}

func (p GetInformationBarrierPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.ExoPolicyId != nil {
		m["ExoPolicyId"] = p.ExoPolicyId
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetInformationBarrierPolicy runs the Get-InformationBarrierPolicy cmdlet.
func (s *Service) GetInformationBarrierPolicy(ctx context.Context, p GetInformationBarrierPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-InformationBarrierPolicy", p.params())
}

// GetInformationBarrierRecipientStatusParams are the parameters of Get-InformationBarrierRecipientStatus.
// DefaultParameterSetName: Identity
type GetInformationBarrierRecipientStatusParams struct {
	Identity  any `ps:"Identity"`
	Identity2 any `ps:"Identity2"`
}

func (p GetInformationBarrierRecipientStatusParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Identity2 != nil {
		m["Identity2"] = p.Identity2
	}
	return m
}

// GetInformationBarrierRecipientStatus runs the Get-InformationBarrierRecipientStatus cmdlet.
func (s *Service) GetInformationBarrierRecipientStatus(ctx context.Context, p GetInformationBarrierRecipientStatusParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-InformationBarrierRecipientStatus", p.params())
}

// GetInsiderRiskEntityListParams are the parameters of Get-InsiderRiskEntityList.
// DefaultParameterSetName: Identity
type GetInsiderRiskEntityListParams struct {
	Identity        any  `ps:"Identity"`
	IncludeDeleted  bool `ps:"IncludeDeleted"`
	IncludeEntities bool `ps:"IncludeEntities"`
	Type            any  `ps:"Type"`
}

func (p GetInsiderRiskEntityListParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IncludeDeleted {
		m["IncludeDeleted"] = true
	}
	if p.IncludeEntities {
		m["IncludeEntities"] = true
	}
	if p.Type != nil {
		m["Type"] = p.Type
	}
	return m
}

// GetInsiderRiskEntityList runs the Get-InsiderRiskEntityList cmdlet.
func (s *Service) GetInsiderRiskEntityList(ctx context.Context, p GetInsiderRiskEntityListParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-InsiderRiskEntityList", p.params())
}

// GetInsiderRiskPolicyParams are the parameters of Get-InsiderRiskPolicy.
// DefaultParameterSetName: Identity
type GetInsiderRiskPolicyParams struct {
	DistributionDetail bool `ps:"DistributionDetail"`
	Identity           any  `ps:"Identity"`
	MetaDataOnly       bool `ps:"MetaDataOnly"`
	SkipPolicyHealth   bool `ps:"SkipPolicyHealth"`
}

func (p GetInsiderRiskPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.DistributionDetail {
		m["DistributionDetail"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.MetaDataOnly {
		m["MetaDataOnly"] = true
	}
	if p.SkipPolicyHealth {
		m["SkipPolicyHealth"] = true
	}
	return m
}

// GetInsiderRiskPolicy runs the Get-InsiderRiskPolicy cmdlet.
func (s *Service) GetInsiderRiskPolicy(ctx context.Context, p GetInsiderRiskPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-InsiderRiskPolicy", p.params())
}

// GetJitConfigurationParams are the parameters of Get-JitConfiguration.
// DefaultParameterSetName: Identity
type GetJitConfigurationParams struct {
	Identity any `ps:"Identity"`
}

func (p GetJitConfigurationParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetJitConfiguration runs the Get-JitConfiguration cmdlet.
func (s *Service) GetJitConfiguration(ctx context.Context, p GetJitConfigurationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-JitConfiguration", p.params())
}

// GetLabelParams are the parameters of Get-Label.
// DefaultParameterSetName: Identity
type GetLabelParams struct {
	Identity                    any  `ps:"Identity"`
	IncludeDetailedLabelActions bool `ps:"IncludeDetailedLabelActions"`
	ReturnModernLabelScheme     bool `ps:"ReturnModernLabelScheme"`
	SkipValidations             bool `ps:"SkipValidations"`
	ValidateContentTypeRemoval  bool `ps:"ValidateContentTypeRemoval"`
}

func (p GetLabelParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IncludeDetailedLabelActions {
		m["IncludeDetailedLabelActions"] = true
	}
	if p.ReturnModernLabelScheme {
		m["ReturnModernLabelScheme"] = true
	}
	if p.SkipValidations {
		m["SkipValidations"] = true
	}
	if p.ValidateContentTypeRemoval {
		m["ValidateContentTypeRemoval"] = true
	}
	return m
}

// GetLabel runs the Get-Label cmdlet.
func (s *Service) GetLabel(ctx context.Context, p GetLabelParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-Label", p.params())
}

// GetLabelExplorerConfigParams are the parameters of Get-LabelExplorerConfig.
type GetLabelExplorerConfigParams struct {
	DisplayName string `ps:"DisplayName"`
	Identity    any    `ps:"Identity"`
	User        string `ps:"User"`
}

func (p GetLabelExplorerConfigParams) params() map[string]any {
	m := map[string]any{}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.User != "" {
		m["User"] = p.User
	}
	return m
}

// GetLabelExplorerConfig runs the Get-LabelExplorerConfig cmdlet.
func (s *Service) GetLabelExplorerConfig(ctx context.Context, p GetLabelExplorerConfigParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-LabelExplorerConfig", p.params())
}

// GetLabelPolicyParams are the parameters of Get-LabelPolicy.
// DefaultParameterSetName: Identity
type GetLabelPolicyParams struct {
	ForceValidate bool `ps:"ForceValidate"`
	Identity      any  `ps:"Identity"`
}

func (p GetLabelPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceValidate {
		m["ForceValidate"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetLabelPolicy runs the Get-LabelPolicy cmdlet.
func (s *Service) GetLabelPolicy(ctx context.Context, p GetLabelPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-LabelPolicy", p.params())
}

// GetLabelPolicyRuleParams are the parameters of Get-LabelPolicyRule.
// DefaultParameterSetName: Identity
type GetLabelPolicyRuleParams struct {
	Identity any `ps:"Identity"`
	Policy   any `ps:"Policy"`
}

func (p GetLabelPolicyRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	return m
}

// GetLabelPolicyRule runs the Get-LabelPolicyRule cmdlet.
func (s *Service) GetLabelPolicyRule(ctx context.Context, p GetLabelPolicyRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-LabelPolicyRule", p.params())
}

// GetLongTermAuditItemsParams are the parameters of Get-LongTermAuditItems.
// DefaultParameterSetName: Identity
type GetLongTermAuditItemsParams struct {
	EndDate    any    `ps:"EndDate"`
	LabelId    string `ps:"LabelId"`
	LTAAddress string `ps:"LTAAddress"`
	PageOffset int    `ps:"PageOffset"`
	StartDate  any    `ps:"StartDate"`
	Workload   string `ps:"Workload"`
}

func (p GetLongTermAuditItemsParams) params() map[string]any {
	m := map[string]any{}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	if p.LabelId != "" {
		m["LabelId"] = p.LabelId
	}
	if p.LTAAddress != "" {
		m["LTAAddress"] = p.LTAAddress
	}
	if p.PageOffset != 0 {
		m["PageOffset"] = p.PageOffset
	}
	if p.StartDate != nil {
		m["StartDate"] = p.StartDate
	}
	if p.Workload != "" {
		m["Workload"] = p.Workload
	}
	return m
}

// GetLongTermAuditItems runs the Get-LongTermAuditItems cmdlet.
func (s *Service) GetLongTermAuditItems(ctx context.Context, p GetLongTermAuditItemsParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-LongTermAuditItems", p.params())
}

// GetLongTermAuditStatsParams are the parameters of Get-LongTermAuditStats.
// DefaultParameterSetName: Identity
type GetLongTermAuditStatsParams struct {
	LtaAddress string `ps:"LtaAddress"`
}

func (p GetLongTermAuditStatsParams) params() map[string]any {
	m := map[string]any{}
	if p.LtaAddress != "" {
		m["LtaAddress"] = p.LtaAddress
	}
	return m
}

// GetLongTermAuditStats runs the Get-LongTermAuditStats cmdlet.
func (s *Service) GetLongTermAuditStats(ctx context.Context, p GetLongTermAuditStatsParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-LongTermAuditStats", p.params())
}

// GetMachineAssistedTagResourceParams are the parameters of Get-MachineAssistedTagResource.
type GetMachineAssistedTagResourceParams struct {
	ClassificationResultType any    `ps:"ClassificationResultType"`
	Identity                 any    `ps:"Identity"`
	Limit                    int    `ps:"Limit"`
	ModelId                  any    `ps:"ModelId"`
	ModuleId                 any    `ps:"ModuleId"`
	Offset                   string `ps:"Offset"`
	Resource                 any    `ps:"Resource"`
	TaskId                   any    `ps:"TaskId"`
	TrainingId               any    `ps:"TrainingId"`
	TrainingType             any    `ps:"TrainingType"`
}

func (p GetMachineAssistedTagResourceParams) params() map[string]any {
	m := map[string]any{}
	if p.ClassificationResultType != nil {
		m["ClassificationResultType"] = p.ClassificationResultType
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Limit != 0 {
		m["Limit"] = p.Limit
	}
	if p.ModelId != nil {
		m["ModelId"] = p.ModelId
	}
	if p.ModuleId != nil {
		m["ModuleId"] = p.ModuleId
	}
	if p.Offset != "" {
		m["Offset"] = p.Offset
	}
	if p.Resource != nil {
		m["Resource"] = p.Resource
	}
	if p.TaskId != nil {
		m["TaskId"] = p.TaskId
	}
	if p.TrainingId != nil {
		m["TrainingId"] = p.TrainingId
	}
	if p.TrainingType != nil {
		m["TrainingType"] = p.TrainingType
	}
	return m
}

// GetMachineAssistedTagResource runs the Get-MachineAssistedTagResource cmdlet.
func (s *Service) GetMachineAssistedTagResource(ctx context.Context, p GetMachineAssistedTagResourceParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-MachineAssistedTagResource", p.params())
}

// GetMailDetailEncryptionReportParams are the parameters of Get-MailDetailEncryptionReport.
type GetMailDetailEncryptionReportParams struct {
	AggregateBy    string `ps:"AggregateBy"`
	Direction      any    `ps:"Direction"`
	Domain         any    `ps:"Domain"`
	EndDate        any    `ps:"EndDate"`
	EventType      any    `ps:"EventType"`
	MessageId      any    `ps:"MessageId"`
	MessageTraceId any    `ps:"MessageTraceId"`
	Organization   any    `ps:"Organization"`
	Page           int    `ps:"Page"`
	PageSize       int    `ps:"PageSize"`
	ProbeTag       string `ps:"ProbeTag"`
	StartDate      any    `ps:"StartDate"`
}

func (p GetMailDetailEncryptionReportParams) params() map[string]any {
	m := map[string]any{}
	if p.AggregateBy != "" {
		m["AggregateBy"] = p.AggregateBy
	}
	if p.Direction != nil {
		m["Direction"] = p.Direction
	}
	if p.Domain != nil {
		m["Domain"] = p.Domain
	}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	if p.EventType != nil {
		m["EventType"] = p.EventType
	}
	if p.MessageId != nil {
		m["MessageId"] = p.MessageId
	}
	if p.MessageTraceId != nil {
		m["MessageTraceId"] = p.MessageTraceId
	}
	if p.Organization != nil {
		m["Organization"] = p.Organization
	}
	if p.Page != 0 {
		m["Page"] = p.Page
	}
	if p.PageSize != 0 {
		m["PageSize"] = p.PageSize
	}
	if p.ProbeTag != "" {
		m["ProbeTag"] = p.ProbeTag
	}
	if p.StartDate != nil {
		m["StartDate"] = p.StartDate
	}
	return m
}

// GetMailDetailEncryptionReport runs the Get-MailDetailEncryptionReport cmdlet.
func (s *Service) GetMailDetailEncryptionReport(ctx context.Context, p GetMailDetailEncryptionReportParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-MailDetailEncryptionReport", p.params())
}

// GetMailFilterListReportParams are the parameters of Get-MailFilterListReport.
type GetMailFilterListReportParams struct {
	Domain          any    `ps:"Domain"`
	Expression      any    `ps:"Expression"`
	ProbeTag        string `ps:"ProbeTag"`
	SelectionTarget any    `ps:"SelectionTarget"`
}

func (p GetMailFilterListReportParams) params() map[string]any {
	m := map[string]any{}
	if p.Domain != nil {
		m["Domain"] = p.Domain
	}
	if p.Expression != nil {
		m["Expression"] = p.Expression
	}
	if p.ProbeTag != "" {
		m["ProbeTag"] = p.ProbeTag
	}
	if p.SelectionTarget != nil {
		m["SelectionTarget"] = p.SelectionTarget
	}
	return m
}

// GetMailFilterListReport runs the Get-MailFilterListReport cmdlet.
func (s *Service) GetMailFilterListReport(ctx context.Context, p GetMailFilterListReportParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-MailFilterListReport", p.params())
}

// GetMailTrafficEncryptionReportParams are the parameters of Get-MailTrafficEncryptionReport.
type GetMailTrafficEncryptionReportParams struct {
	Action      any    `ps:"Action"`
	AggregateBy string `ps:"AggregateBy"`
	Direction   any    `ps:"Direction"`
	Domain      any    `ps:"Domain"`
	EndDate     any    `ps:"EndDate"`
	EventType   any    `ps:"EventType"`
	Page        int    `ps:"Page"`
	PageSize    int    `ps:"PageSize"`
	ProbeTag    string `ps:"ProbeTag"`
	StartDate   any    `ps:"StartDate"`
	SummarizeBy any    `ps:"SummarizeBy"`
}

func (p GetMailTrafficEncryptionReportParams) params() map[string]any {
	m := map[string]any{}
	if p.Action != nil {
		m["Action"] = p.Action
	}
	if p.AggregateBy != "" {
		m["AggregateBy"] = p.AggregateBy
	}
	if p.Direction != nil {
		m["Direction"] = p.Direction
	}
	if p.Domain != nil {
		m["Domain"] = p.Domain
	}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	if p.EventType != nil {
		m["EventType"] = p.EventType
	}
	if p.Page != 0 {
		m["Page"] = p.Page
	}
	if p.PageSize != 0 {
		m["PageSize"] = p.PageSize
	}
	if p.ProbeTag != "" {
		m["ProbeTag"] = p.ProbeTag
	}
	if p.StartDate != nil {
		m["StartDate"] = p.StartDate
	}
	if p.SummarizeBy != nil {
		m["SummarizeBy"] = p.SummarizeBy
	}
	return m
}

// GetMailTrafficEncryptionReport runs the Get-MailTrafficEncryptionReport cmdlet.
func (s *Service) GetMailTrafficEncryptionReport(ctx context.Context, p GetMailTrafficEncryptionReportParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-MailTrafficEncryptionReport", p.params())
}

// GetMultiTenantOrganizationConfigParams are the parameters of Get-MultiTenantOrganizationConfig.
type GetMultiTenantOrganizationConfigParams struct {
	Identity any `ps:"Identity"`
}

func (p GetMultiTenantOrganizationConfigParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetMultiTenantOrganizationConfig runs the Get-MultiTenantOrganizationConfig cmdlet.
func (s *Service) GetMultiTenantOrganizationConfig(ctx context.Context, p GetMultiTenantOrganizationConfigParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-MultiTenantOrganizationConfig", p.params())
}

// GetOcrConfigurationParams are the parameters of Get-OcrConfiguration.
// DefaultParameterSetName: Identity
type GetOcrConfigurationParams struct {
	Identity any `ps:"Identity"`
}

func (p GetOcrConfigurationParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetOcrConfiguration runs the Get-OcrConfiguration cmdlet.
func (s *Service) GetOcrConfiguration(ctx context.Context, p GetOcrConfigurationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-OcrConfiguration", p.params())
}

// GetOrganizationSegmentParams are the parameters of Get-OrganizationSegment.
// DefaultParameterSetName: OrganizationSegmentsDefault
type GetOrganizationSegmentParams struct {
	Identity any `ps:"Identity"`
}

func (p GetOrganizationSegmentParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetOrganizationSegment runs the Get-OrganizationSegment cmdlet.
func (s *Service) GetOrganizationSegment(ctx context.Context, p GetOrganizationSegmentParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-OrganizationSegment", p.params())
}

// GetPolicyConfigParams are the parameters of Get-PolicyConfig.
type GetPolicyConfigParams struct {
	Identity any `ps:"Identity"`
}

func (p GetPolicyConfigParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetPolicyConfig runs the Get-PolicyConfig cmdlet.
func (s *Service) GetPolicyConfig(ctx context.Context, p GetPolicyConfigParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-PolicyConfig", p.params())
}

// GetPrivacyManagementCaseAdminParams are the parameters of Get-PrivacyManagementCaseAdmin.
// DefaultParameterSetName: Identity
type GetPrivacyManagementCaseAdminParams struct {
	DomainController any `ps:"DomainController"`
	ResultSize       any `ps:"ResultSize"`
}

func (p GetPrivacyManagementCaseAdminParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.ResultSize != nil {
		m["ResultSize"] = p.ResultSize
	}
	return m
}

// GetPrivacyManagementCaseAdmin runs the Get-PrivacyManagementCaseAdmin cmdlet.
func (s *Service) GetPrivacyManagementCaseAdmin(ctx context.Context, p GetPrivacyManagementCaseAdminParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-PrivacyManagementCaseAdmin", p.params())
}

// GetPrivacyManagementComplianceCaseMemberParams are the parameters of Get-PrivacyManagementComplianceCaseMember.
// DefaultParameterSetName: Identity
type GetPrivacyManagementComplianceCaseMemberParams struct {
	Case             string `ps:"Case"`
	DomainController any    `ps:"DomainController"`
	ResultSize       any    `ps:"ResultSize"`
	RoleGroupOnly    bool   `ps:"RoleGroupOnly"` // one of: True
	ShowCaseAdmin    bool   `ps:"ShowCaseAdmin"`
}

func (p GetPrivacyManagementComplianceCaseMemberParams) params() map[string]any {
	m := map[string]any{}
	if p.Case != "" {
		m["Case"] = p.Case
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.ResultSize != nil {
		m["ResultSize"] = p.ResultSize
	}
	if p.RoleGroupOnly {
		m["RoleGroupOnly"] = true
	}
	if p.ShowCaseAdmin {
		m["ShowCaseAdmin"] = true
	}
	return m
}

// GetPrivacyManagementComplianceCaseMember runs the Get-PrivacyManagementComplianceCaseMember cmdlet.
func (s *Service) GetPrivacyManagementComplianceCaseMember(ctx context.Context, p GetPrivacyManagementComplianceCaseMemberParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-PrivacyManagementComplianceCaseMember", p.params())
}

// GetPrivacyManagementComplianceTagParams are the parameters of Get-PrivacyManagementComplianceTag.
// DefaultParameterSetName: Identity
type GetPrivacyManagementComplianceTagParams struct {
	Identity     any `ps:"Identity"`
	Organization any `ps:"Organization"`
}

func (p GetPrivacyManagementComplianceTagParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Organization != nil {
		m["Organization"] = p.Organization
	}
	return m
}

// GetPrivacyManagementComplianceTag runs the Get-PrivacyManagementComplianceTag cmdlet.
func (s *Service) GetPrivacyManagementComplianceTag(ctx context.Context, p GetPrivacyManagementComplianceTagParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-PrivacyManagementComplianceTag", p.params())
}

// GetPrivacyManagementPolicyParams are the parameters of Get-PrivacyManagementPolicy.
// DefaultParameterSetName: Identity
type GetPrivacyManagementPolicyParams struct {
	DistributionDetail        bool `ps:"DistributionDetail"`
	Identity                  any  `ps:"Identity"`
	PrivacyManagementScenario any  `ps:"PrivacyManagementScenario"`
}

func (p GetPrivacyManagementPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.DistributionDetail {
		m["DistributionDetail"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.PrivacyManagementScenario != nil {
		m["PrivacyManagementScenario"] = p.PrivacyManagementScenario
	}
	return m
}

// GetPrivacyManagementPolicy runs the Get-PrivacyManagementPolicy cmdlet.
func (s *Service) GetPrivacyManagementPolicy(ctx context.Context, p GetPrivacyManagementPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-PrivacyManagementPolicy", p.params())
}

// GetPrivacyManagementRuleParams are the parameters of Get-PrivacyManagementRule.
// DefaultParameterSetName: Identity
type GetPrivacyManagementRuleParams struct {
	Identity any `ps:"Identity"`
}

func (p GetPrivacyManagementRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetPrivacyManagementRule runs the Get-PrivacyManagementRule cmdlet.
func (s *Service) GetPrivacyManagementRule(ctx context.Context, p GetPrivacyManagementRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-PrivacyManagementRule", p.params())
}

// GetProtectionAlertParams are the parameters of Get-ProtectionAlert.
// DefaultParameterSetName: Identity
type GetProtectionAlertParams struct {
	Identity       any  `ps:"Identity"`
	IncludeRuleXml bool `ps:"IncludeRuleXml"`
}

func (p GetProtectionAlertParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IncludeRuleXml {
		m["IncludeRuleXml"] = true
	}
	return m
}

// GetProtectionAlert runs the Get-ProtectionAlert cmdlet.
func (s *Service) GetProtectionAlert(ctx context.Context, p GetProtectionAlertParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ProtectionAlert", p.params())
}

// GetProtectionCompliancePolicyParams are the parameters of Get-ProtectionCompliancePolicy.
// DefaultParameterSetName: Identity
type GetProtectionCompliancePolicyParams struct {
	DistributionDetail bool `ps:"DistributionDetail"`
	Identity           any  `ps:"Identity"`
}

func (p GetProtectionCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.DistributionDetail {
		m["DistributionDetail"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetProtectionCompliancePolicy runs the Get-ProtectionCompliancePolicy cmdlet.
func (s *Service) GetProtectionCompliancePolicy(ctx context.Context, p GetProtectionCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ProtectionCompliancePolicy", p.params())
}

// GetProtectionComplianceRuleParams are the parameters of Get-ProtectionComplianceRule.
type GetProtectionComplianceRuleParams struct {
	Identity any `ps:"Identity"`
	Policy   any `ps:"Policy"`
}

func (p GetProtectionComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	return m
}

// GetProtectionComplianceRule runs the Get-ProtectionComplianceRule cmdlet.
func (s *Service) GetProtectionComplianceRule(ctx context.Context, p GetProtectionComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ProtectionComplianceRule", p.params())
}

// GetQuarantineMessageParams are the parameters of Get-QuarantineMessage.
// DefaultParameterSetName: Summary
type GetQuarantineMessageParams struct {
	Direction                               any      `ps:"Direction"`
	EndExpiresDate                          any      `ps:"EndExpiresDate"`
	EndReceivedDate                         any      `ps:"EndReceivedDate"`
	EntityType                              any      `ps:"EntityType"`
	Identity                                any      `ps:"Identity"`
	IncludeMessagesFromBlockedSenderAddress bool     `ps:"IncludeMessagesFromBlockedSenderAddress"`
	MessageId                               string   `ps:"MessageId"`
	MoveToQuarantineAdminActionTakenBy      string   `ps:"MoveToQuarantineAdminActionTakenBy"`
	MoveToQuarantineApprovalId              any      `ps:"MoveToQuarantineApprovalId"`
	MyItems                                 bool     `ps:"MyItems"`
	NetworkMessageId                        string   `ps:"NetworkMessageId"`
	Page                                    any      `ps:"Page"`
	PageSize                                any      `ps:"PageSize"`
	PolicyName                              string   `ps:"PolicyName"`
	PolicyTypes                             []string `ps:"PolicyTypes"`
	QuarantineTypes                         []string `ps:"QuarantineTypes"`
	RecipientAddress                        []string `ps:"RecipientAddress"`
	RecipientTag                            []string `ps:"RecipientTag"`
	ReleaseStatus                           []string `ps:"ReleaseStatus"`
	Reported                                any      `ps:"Reported"`
	SearchTemporaryDeletedMessage           bool     `ps:"SearchTemporaryDeletedMessage"`
	SenderAddress                           []string `ps:"SenderAddress"`
	StartExpiresDate                        any      `ps:"StartExpiresDate"`
	StartReceivedDate                       any      `ps:"StartReceivedDate"`
	Subject                                 string   `ps:"Subject"`
	TeamsConversationTypes                  []string `ps:"TeamsConversationTypes"`
	Type                                    any      `ps:"Type"`
}

func (p GetQuarantineMessageParams) params() map[string]any {
	m := map[string]any{}
	if p.Direction != nil {
		m["Direction"] = p.Direction
	}
	if p.EndExpiresDate != nil {
		m["EndExpiresDate"] = p.EndExpiresDate
	}
	if p.EndReceivedDate != nil {
		m["EndReceivedDate"] = p.EndReceivedDate
	}
	if p.EntityType != nil {
		m["EntityType"] = p.EntityType
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IncludeMessagesFromBlockedSenderAddress {
		m["IncludeMessagesFromBlockedSenderAddress"] = true
	}
	if p.MessageId != "" {
		m["MessageId"] = p.MessageId
	}
	if p.MoveToQuarantineAdminActionTakenBy != "" {
		m["MoveToQuarantineAdminActionTakenBy"] = p.MoveToQuarantineAdminActionTakenBy
	}
	if p.MoveToQuarantineApprovalId != nil {
		m["MoveToQuarantineApprovalId"] = p.MoveToQuarantineApprovalId
	}
	if p.MyItems {
		m["MyItems"] = true
	}
	if p.NetworkMessageId != "" {
		m["NetworkMessageId"] = p.NetworkMessageId
	}
	if p.Page != nil {
		m["Page"] = p.Page
	}
	if p.PageSize != nil {
		m["PageSize"] = p.PageSize
	}
	if p.PolicyName != "" {
		m["PolicyName"] = p.PolicyName
	}
	if len(p.PolicyTypes) > 0 {
		m["PolicyTypes"] = p.PolicyTypes
	}
	if len(p.QuarantineTypes) > 0 {
		m["QuarantineTypes"] = p.QuarantineTypes
	}
	if len(p.RecipientAddress) > 0 {
		m["RecipientAddress"] = p.RecipientAddress
	}
	if len(p.RecipientTag) > 0 {
		m["RecipientTag"] = p.RecipientTag
	}
	if len(p.ReleaseStatus) > 0 {
		m["ReleaseStatus"] = p.ReleaseStatus
	}
	if p.Reported != nil {
		m["Reported"] = p.Reported
	}
	if p.SearchTemporaryDeletedMessage {
		m["SearchTemporaryDeletedMessage"] = true
	}
	if len(p.SenderAddress) > 0 {
		m["SenderAddress"] = p.SenderAddress
	}
	if p.StartExpiresDate != nil {
		m["StartExpiresDate"] = p.StartExpiresDate
	}
	if p.StartReceivedDate != nil {
		m["StartReceivedDate"] = p.StartReceivedDate
	}
	if p.Subject != "" {
		m["Subject"] = p.Subject
	}
	if len(p.TeamsConversationTypes) > 0 {
		m["TeamsConversationTypes"] = p.TeamsConversationTypes
	}
	if p.Type != nil {
		m["Type"] = p.Type
	}
	return m
}

// GetQuarantineMessage runs the Get-QuarantineMessage cmdlet.
func (s *Service) GetQuarantineMessage(ctx context.Context, p GetQuarantineMessageParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-QuarantineMessage", p.params())
}

// GetQuarantineMessageHeaderParams are the parameters of Get-QuarantineMessageHeader.
type GetQuarantineMessageHeaderParams struct {
	EntityType       any    `ps:"EntityType"`
	Identity         any    `ps:"Identity"`
	RecipientAddress string `ps:"RecipientAddress"`
}

func (p GetQuarantineMessageHeaderParams) params() map[string]any {
	m := map[string]any{}
	if p.EntityType != nil {
		m["EntityType"] = p.EntityType
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.RecipientAddress != "" {
		m["RecipientAddress"] = p.RecipientAddress
	}
	return m
}

// GetQuarantineMessageHeader runs the Get-QuarantineMessageHeader cmdlet.
func (s *Service) GetQuarantineMessageHeader(ctx context.Context, p GetQuarantineMessageHeaderParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-QuarantineMessageHeader", p.params())
}

// GetRecipientParams are the parameters of Get-Recipient.
// DefaultParameterSetName: Identity
type GetRecipientParams struct {
	Anr                    string   `ps:"Anr"`
	AuthenticationType     any      `ps:"AuthenticationType"`
	BookmarkDisplayName    string   `ps:"BookmarkDisplayName"`
	Capabilities           any      `ps:"Capabilities"`
	Filter                 string   `ps:"Filter"`
	Identity               any      `ps:"Identity"`
	IncludeBookmarkObject  bool     `ps:"IncludeBookmarkObject"`
	Properties             []string `ps:"Properties"`
	PropertySet            any      `ps:"PropertySet"`
	RecipientPreviewFilter string   `ps:"RecipientPreviewFilter"`
	RecipientType          []string `ps:"RecipientType"`
	RecipientTypeDetails   []string `ps:"RecipientTypeDetails"`
	ResultSize             any      `ps:"ResultSize"`
	SortBy                 string   `ps:"SortBy"`
}

func (p GetRecipientParams) params() map[string]any {
	m := map[string]any{}
	if p.Anr != "" {
		m["Anr"] = p.Anr
	}
	if p.AuthenticationType != nil {
		m["AuthenticationType"] = p.AuthenticationType
	}
	if p.BookmarkDisplayName != "" {
		m["BookmarkDisplayName"] = p.BookmarkDisplayName
	}
	if p.Capabilities != nil {
		m["Capabilities"] = p.Capabilities
	}
	if p.Filter != "" {
		m["Filter"] = p.Filter
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IncludeBookmarkObject {
		m["IncludeBookmarkObject"] = true
	}
	if len(p.Properties) > 0 {
		m["Properties"] = p.Properties
	}
	if p.PropertySet != nil {
		m["PropertySet"] = p.PropertySet
	}
	if p.RecipientPreviewFilter != "" {
		m["RecipientPreviewFilter"] = p.RecipientPreviewFilter
	}
	if len(p.RecipientType) > 0 {
		m["RecipientType"] = p.RecipientType
	}
	if len(p.RecipientTypeDetails) > 0 {
		m["RecipientTypeDetails"] = p.RecipientTypeDetails
	}
	if p.ResultSize != nil {
		m["ResultSize"] = p.ResultSize
	}
	if p.SortBy != "" {
		m["SortBy"] = p.SortBy
	}
	return m
}

// GetRecipient runs the Get-Recipient cmdlet.
func (s *Service) GetRecipient(ctx context.Context, p GetRecipientParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-Recipient", p.params())
}

// GetRecordReviewNotificationTemplateConfigParams are the parameters of Get-RecordReviewNotificationTemplateConfig.
type GetRecordReviewNotificationTemplateConfigParams struct {
}

func (p GetRecordReviewNotificationTemplateConfigParams) params() map[string]any {
	m := map[string]any{}
	return m
}

// GetRecordReviewNotificationTemplateConfig runs the Get-RecordReviewNotificationTemplateConfig cmdlet.
func (s *Service) GetRecordReviewNotificationTemplateConfig(ctx context.Context, p GetRecordReviewNotificationTemplateConfigParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-RecordReviewNotificationTemplateConfig", p.params())
}

// GetRegulatoryComplianceUIParams are the parameters of Get-RegulatoryComplianceUI.
// DefaultParameterSetName: Identity
type GetRegulatoryComplianceUIParams struct {
}

func (p GetRegulatoryComplianceUIParams) params() map[string]any {
	m := map[string]any{}
	return m
}

// GetRegulatoryComplianceUI runs the Get-RegulatoryComplianceUI cmdlet.
func (s *Service) GetRegulatoryComplianceUI(ctx context.Context, p GetRegulatoryComplianceUIParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-RegulatoryComplianceUI", p.params())
}

// GetRetentionCompliancePolicyParams are the parameters of Get-RetentionCompliancePolicy.
// DefaultParameterSetName: Identity
type GetRetentionCompliancePolicyParams struct {
	DistributionDetail     bool `ps:"DistributionDetail"`
	ErrorPolicyOnly        bool `ps:"ErrorPolicyOnly"`
	ExcludeTeamsPolicy     bool `ps:"ExcludeTeamsPolicy"`
	Identity               any  `ps:"Identity"`
	IncludeTestModeResults bool `ps:"IncludeTestModeResults"`
	RetentionRuleTypes     bool `ps:"RetentionRuleTypes"`
	TeamsPolicyOnly        bool `ps:"TeamsPolicyOnly"`
}

func (p GetRetentionCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.DistributionDetail {
		m["DistributionDetail"] = true
	}
	if p.ErrorPolicyOnly {
		m["ErrorPolicyOnly"] = true
	}
	if p.ExcludeTeamsPolicy {
		m["ExcludeTeamsPolicy"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IncludeTestModeResults {
		m["IncludeTestModeResults"] = true
	}
	if p.RetentionRuleTypes {
		m["RetentionRuleTypes"] = true
	}
	if p.TeamsPolicyOnly {
		m["TeamsPolicyOnly"] = true
	}
	return m
}

// GetRetentionCompliancePolicy runs the Get-RetentionCompliancePolicy cmdlet.
func (s *Service) GetRetentionCompliancePolicy(ctx context.Context, p GetRetentionCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-RetentionCompliancePolicy", p.params())
}

// GetRetentionComplianceRuleParams are the parameters of Get-RetentionComplianceRule.
// DefaultParameterSetName: Identity
type GetRetentionComplianceRuleParams struct {
	Identity any `ps:"Identity"`
	Policy   any `ps:"Policy"`
}

func (p GetRetentionComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	return m
}

// GetRetentionComplianceRule runs the Get-RetentionComplianceRule cmdlet.
func (s *Service) GetRetentionComplianceRule(ctx context.Context, p GetRetentionComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-RetentionComplianceRule", p.params())
}

// GetRetentionEventParams are the parameters of Get-RetentionEvent.
// DefaultParameterSetName: Identity
type GetRetentionEventParams struct {
	DomainController any `ps:"DomainController"`
	Identity         any `ps:"Identity"`
}

func (p GetRetentionEventParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetRetentionEvent runs the Get-RetentionEvent cmdlet.
func (s *Service) GetRetentionEvent(ctx context.Context, p GetRetentionEventParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-RetentionEvent", p.params())
}

// GetRoleGroupParams are the parameters of Get-RoleGroup.
// DefaultParameterSetName: Identity
type GetRoleGroupParams struct {
	Filter            string `ps:"Filter"`
	Identity          any    `ps:"Identity"`
	ResultSize        any    `ps:"ResultSize"`
	ShowPartnerLinked bool   `ps:"ShowPartnerLinked"`
	SortBy            string `ps:"SortBy"`
}

func (p GetRoleGroupParams) params() map[string]any {
	m := map[string]any{}
	if p.Filter != "" {
		m["Filter"] = p.Filter
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.ResultSize != nil {
		m["ResultSize"] = p.ResultSize
	}
	if p.ShowPartnerLinked {
		m["ShowPartnerLinked"] = true
	}
	if p.SortBy != "" {
		m["SortBy"] = p.SortBy
	}
	return m
}

// GetRoleGroup runs the Get-RoleGroup cmdlet.
func (s *Service) GetRoleGroup(ctx context.Context, p GetRoleGroupParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-RoleGroup", p.params())
}

// GetRoleGroupMemberParams are the parameters of Get-RoleGroupMember.
type GetRoleGroupMemberParams struct {
	Identity   any `ps:"Identity"`
	ResultSize any `ps:"ResultSize"`
}

func (p GetRoleGroupMemberParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.ResultSize != nil {
		m["ResultSize"] = p.ResultSize
	}
	return m
}

// GetRoleGroupMember runs the Get-RoleGroupMember cmdlet.
func (s *Service) GetRoleGroupMember(ctx context.Context, p GetRoleGroupMemberParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-RoleGroupMember", p.params())
}

// GetSCInsightsParams are the parameters of Get-SCInsights.
type GetSCInsightsParams struct {
	EndDate any `ps:"EndDate"`
}

func (p GetSCInsightsParams) params() map[string]any {
	m := map[string]any{}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	return m
}

// GetSCInsights runs the Get-SCInsights cmdlet.
func (s *Service) GetSCInsights(ctx context.Context, p GetSCInsightsParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-SCInsights", p.params())
}

// GetScopeAdminsParams are the parameters of Get-ScopeAdmins.
// DefaultParameterSetName: Identity
type GetScopeAdminsParams struct {
	Identity any `ps:"Identity"`
}

func (p GetScopeAdminsParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetScopeAdmins runs the Get-ScopeAdmins cmdlet.
func (s *Service) GetScopeAdmins(ctx context.Context, p GetScopeAdminsParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ScopeAdmins", p.params())
}

// GetScopeEntitiesParams are the parameters of Get-ScopeEntities.
// DefaultParameterSetName: Identity
type GetScopeEntitiesParams struct {
	Filter             string `ps:"Filter"`
	Identity           any    `ps:"Identity"`
	ResultSize         any    `ps:"ResultSize"`
	ScopeRecipientType string `ps:"ScopeRecipientType"`
	SortBy             string `ps:"SortBy"`
}

func (p GetScopeEntitiesParams) params() map[string]any {
	m := map[string]any{}
	if p.Filter != "" {
		m["Filter"] = p.Filter
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.ResultSize != nil {
		m["ResultSize"] = p.ResultSize
	}
	if p.ScopeRecipientType != "" {
		m["ScopeRecipientType"] = p.ScopeRecipientType
	}
	if p.SortBy != "" {
		m["SortBy"] = p.SortBy
	}
	return m
}

// GetScopeEntities runs the Get-ScopeEntities cmdlet.
func (s *Service) GetScopeEntities(ctx context.Context, p GetScopeEntitiesParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ScopeEntities", p.params())
}

// GetSensitiveInformationScanParams are the parameters of Get-SensitiveInformationScan.
type GetSensitiveInformationScanParams struct {
	Identity                         any  `ps:"Identity"`
	IncludeImpactAssessment          bool `ps:"IncludeImpactAssessment"`
	IncludeProgressForAllActiveScans bool `ps:"IncludeProgressForAllActiveScans"`
	IncludeScanProgress              bool `ps:"IncludeScanProgress"`
}

func (p GetSensitiveInformationScanParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IncludeImpactAssessment {
		m["IncludeImpactAssessment"] = true
	}
	if p.IncludeProgressForAllActiveScans {
		m["IncludeProgressForAllActiveScans"] = true
	}
	if p.IncludeScanProgress {
		m["IncludeScanProgress"] = true
	}
	return m
}

// GetSensitiveInformationScan runs the Get-SensitiveInformationScan cmdlet.
func (s *Service) GetSensitiveInformationScan(ctx context.Context, p GetSensitiveInformationScanParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-SensitiveInformationScan", p.params())
}

// GetSensitiveInformationScanRuleParams are the parameters of Get-SensitiveInformationScanRule.
type GetSensitiveInformationScanRuleParams struct {
	Identity any `ps:"Identity"`
	Policy   any `ps:"Policy"`
}

func (p GetSensitiveInformationScanRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	return m
}

// GetSensitiveInformationScanRule runs the Get-SensitiveInformationScanRule cmdlet.
func (s *Service) GetSensitiveInformationScanRule(ctx context.Context, p GetSensitiveInformationScanRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-SensitiveInformationScanRule", p.params())
}

// GetServiceDomainGroupParams are the parameters of Get-ServiceDomainGroup.
// DefaultParameterSetName: Identity
type GetServiceDomainGroupParams struct {
	Identity     any `ps:"Identity"`
	Organization any `ps:"Organization"`
}

func (p GetServiceDomainGroupParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Organization != nil {
		m["Organization"] = p.Organization
	}
	return m
}

// GetServiceDomainGroup runs the Get-ServiceDomainGroup cmdlet.
func (s *Service) GetServiceDomainGroup(ctx context.Context, p GetServiceDomainGroupParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ServiceDomainGroup", p.params())
}

// GetServiceDomainGroupStorageParams are the parameters of Get-ServiceDomainGroupStorage.
// DefaultParameterSetName: Identity
type GetServiceDomainGroupStorageParams struct {
}

func (p GetServiceDomainGroupStorageParams) params() map[string]any {
	m := map[string]any{}
	return m
}

// GetServiceDomainGroupStorage runs the Get-ServiceDomainGroupStorage cmdlet.
func (s *Service) GetServiceDomainGroupStorage(ctx context.Context, p GetServiceDomainGroupStorageParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ServiceDomainGroupStorage", p.params())
}

// GetServicePrincipalParams are the parameters of Get-ServicePrincipal.
// DefaultParameterSetName: Identity
type GetServicePrincipalParams struct {
	Identity     any `ps:"Identity"`
	Organization any `ps:"Organization"`
}

func (p GetServicePrincipalParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Organization != nil {
		m["Organization"] = p.Organization
	}
	return m
}

// GetServicePrincipal runs the Get-ServicePrincipal cmdlet.
func (s *Service) GetServicePrincipal(ctx context.Context, p GetServicePrincipalParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ServicePrincipal", p.params())
}

// GetSupervisoryReviewActivityParams are the parameters of Get-SupervisoryReviewActivity.
// DefaultParameterSetName: Identity
type GetSupervisoryReviewActivityParams struct {
	EndDate   any    `ps:"EndDate"`
	PolicyId  string `ps:"PolicyId"`
	StartDate any    `ps:"StartDate"`
}

func (p GetSupervisoryReviewActivityParams) params() map[string]any {
	m := map[string]any{}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	if p.PolicyId != "" {
		m["PolicyId"] = p.PolicyId
	}
	if p.StartDate != nil {
		m["StartDate"] = p.StartDate
	}
	return m
}

// GetSupervisoryReviewActivity runs the Get-SupervisoryReviewActivity cmdlet.
func (s *Service) GetSupervisoryReviewActivity(ctx context.Context, p GetSupervisoryReviewActivityParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-SupervisoryReviewActivity", p.params())
}

// GetSupervisoryReviewOverallProgressReportParams are the parameters of Get-SupervisoryReviewOverallProgressReport.
type GetSupervisoryReviewOverallProgressReportParams struct {
	EndDate   any `ps:"EndDate"`
	Page      int `ps:"Page"`
	PageSize  int `ps:"PageSize"`
	StartDate any `ps:"StartDate"`
}

func (p GetSupervisoryReviewOverallProgressReportParams) params() map[string]any {
	m := map[string]any{}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	if p.Page != 0 {
		m["Page"] = p.Page
	}
	if p.PageSize != 0 {
		m["PageSize"] = p.PageSize
	}
	if p.StartDate != nil {
		m["StartDate"] = p.StartDate
	}
	return m
}

// GetSupervisoryReviewOverallProgressReport runs the Get-SupervisoryReviewOverallProgressReport cmdlet.
func (s *Service) GetSupervisoryReviewOverallProgressReport(ctx context.Context, p GetSupervisoryReviewOverallProgressReportParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-SupervisoryReviewOverallProgressReport", p.params())
}

// GetSupervisoryReviewPolicyReportParams are the parameters of Get-SupervisoryReviewPolicyReport.
type GetSupervisoryReviewPolicyReportParams struct {
	EndDate   any `ps:"EndDate"`
	Page      int `ps:"Page"`
	PageSize  int `ps:"PageSize"`
	Policies  any `ps:"Policies"`
	StartDate any `ps:"StartDate"`
}

func (p GetSupervisoryReviewPolicyReportParams) params() map[string]any {
	m := map[string]any{}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	if p.Page != 0 {
		m["Page"] = p.Page
	}
	if p.PageSize != 0 {
		m["PageSize"] = p.PageSize
	}
	if p.Policies != nil {
		m["Policies"] = p.Policies
	}
	if p.StartDate != nil {
		m["StartDate"] = p.StartDate
	}
	return m
}

// GetSupervisoryReviewPolicyReport runs the Get-SupervisoryReviewPolicyReport cmdlet.
func (s *Service) GetSupervisoryReviewPolicyReport(ctx context.Context, p GetSupervisoryReviewPolicyReportParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-SupervisoryReviewPolicyReport", p.params())
}

// GetSupervisoryReviewPolicyV2Params are the parameters of Get-SupervisoryReviewPolicyV2.
// DefaultParameterSetName: Identity
type GetSupervisoryReviewPolicyV2Params struct {
	Identity any `ps:"Identity"`
}

func (p GetSupervisoryReviewPolicyV2Params) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetSupervisoryReviewPolicyV2 runs the Get-SupervisoryReviewPolicyV2 cmdlet.
func (s *Service) GetSupervisoryReviewPolicyV2(ctx context.Context, p GetSupervisoryReviewPolicyV2Params) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-SupervisoryReviewPolicyV2", p.params())
}

// GetSupervisoryReviewReportParams are the parameters of Get-SupervisoryReviewReport.
type GetSupervisoryReviewReportParams struct {
	EndDate   any `ps:"EndDate"`
	Page      int `ps:"Page"`
	PageSize  int `ps:"PageSize"`
	Policies  any `ps:"Policies"`
	Reviewers any `ps:"Reviewers"`
	StartDate any `ps:"StartDate"`
}

func (p GetSupervisoryReviewReportParams) params() map[string]any {
	m := map[string]any{}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	if p.Page != 0 {
		m["Page"] = p.Page
	}
	if p.PageSize != 0 {
		m["PageSize"] = p.PageSize
	}
	if p.Policies != nil {
		m["Policies"] = p.Policies
	}
	if p.Reviewers != nil {
		m["Reviewers"] = p.Reviewers
	}
	if p.StartDate != nil {
		m["StartDate"] = p.StartDate
	}
	return m
}

// GetSupervisoryReviewReport runs the Get-SupervisoryReviewReport cmdlet.
func (s *Service) GetSupervisoryReviewReport(ctx context.Context, p GetSupervisoryReviewReportParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-SupervisoryReviewReport", p.params())
}

// GetSupervisoryReviewRuleParams are the parameters of Get-SupervisoryReviewRule.
// DefaultParameterSetName: Identity
type GetSupervisoryReviewRuleParams struct {
	Identity       any  `ps:"Identity"`
	IncludeRuleXml bool `ps:"IncludeRuleXml"`
	Policy         any  `ps:"Policy"`
}

func (p GetSupervisoryReviewRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IncludeRuleXml {
		m["IncludeRuleXml"] = true
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	return m
}

// GetSupervisoryReviewRule runs the Get-SupervisoryReviewRule cmdlet.
func (s *Service) GetSupervisoryReviewRule(ctx context.Context, p GetSupervisoryReviewRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-SupervisoryReviewRule", p.params())
}

// GetSupervisoryReviewTopCasesReportParams are the parameters of Get-SupervisoryReviewTopCasesReport.
type GetSupervisoryReviewTopCasesReportParams struct {
	EndDate   any `ps:"EndDate"`
	Page      int `ps:"Page"`
	PageSize  int `ps:"PageSize"`
	StartDate any `ps:"StartDate"`
}

func (p GetSupervisoryReviewTopCasesReportParams) params() map[string]any {
	m := map[string]any{}
	if p.EndDate != nil {
		m["EndDate"] = p.EndDate
	}
	if p.Page != 0 {
		m["Page"] = p.Page
	}
	if p.PageSize != 0 {
		m["PageSize"] = p.PageSize
	}
	if p.StartDate != nil {
		m["StartDate"] = p.StartDate
	}
	return m
}

// GetSupervisoryReviewTopCasesReport runs the Get-SupervisoryReviewTopCasesReport cmdlet.
func (s *Service) GetSupervisoryReviewTopCasesReport(ctx context.Context, p GetSupervisoryReviewTopCasesReportParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-SupervisoryReviewTopCasesReport", p.params())
}

// GetTeamsRetentionCompliancePolicyParams are the parameters of Get-TeamsRetentionCompliancePolicy.
type GetTeamsRetentionCompliancePolicyParams struct {
	DistributionDetail bool `ps:"DistributionDetail"`
	Identity           any  `ps:"Identity"`
}

func (p GetTeamsRetentionCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.DistributionDetail {
		m["DistributionDetail"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetTeamsRetentionCompliancePolicy runs the Get-TeamsRetentionCompliancePolicy cmdlet.
func (s *Service) GetTeamsRetentionCompliancePolicy(ctx context.Context, p GetTeamsRetentionCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-TeamsRetentionCompliancePolicy", p.params())
}

// GetTeamsRetentionComplianceRuleParams are the parameters of Get-TeamsRetentionComplianceRule.
// DefaultParameterSetName: Identity
type GetTeamsRetentionComplianceRuleParams struct {
	Identity any `ps:"Identity"`
	Policy   any `ps:"Policy"`
}

func (p GetTeamsRetentionComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	return m
}

// GetTeamsRetentionComplianceRule runs the Get-TeamsRetentionComplianceRule cmdlet.
func (s *Service) GetTeamsRetentionComplianceRule(ctx context.Context, p GetTeamsRetentionComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-TeamsRetentionComplianceRule", p.params())
}

// GetThreatResponsePolicyParams are the parameters of Get-ThreatResponsePolicy.
// DefaultParameterSetName: Identity
type GetThreatResponsePolicyParams struct {
	DistributionDetail bool `ps:"DistributionDetail"`
	Identity           any  `ps:"Identity"`
}

func (p GetThreatResponsePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.DistributionDetail {
		m["DistributionDetail"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// GetThreatResponsePolicy runs the Get-ThreatResponsePolicy cmdlet.
func (s *Service) GetThreatResponsePolicy(ctx context.Context, p GetThreatResponsePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ThreatResponsePolicy", p.params())
}

// GetThreatResponseRuleParams are the parameters of Get-ThreatResponseRule.
// DefaultParameterSetName: Identity
type GetThreatResponseRuleParams struct {
	Identity any `ps:"Identity"`
	Policy   any `ps:"Policy"`
}

func (p GetThreatResponseRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	return m
}

// GetThreatResponseRule runs the Get-ThreatResponseRule cmdlet.
func (s *Service) GetThreatResponseRule(ctx context.Context, p GetThreatResponseRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-ThreatResponseRule", p.params())
}

// GetUnifiedAuditLogRetentionPolicyParams are the parameters of Get-UnifiedAuditLogRetentionPolicy.
// DefaultParameterSetName: Identity
type GetUnifiedAuditLogRetentionPolicyParams struct {
	Operation         string `ps:"Operation"`
	RecordType        any    `ps:"RecordType"`
	RetentionDuration any    `ps:"RetentionDuration"`
	UserId            string `ps:"UserId"`
}

func (p GetUnifiedAuditLogRetentionPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Operation != "" {
		m["Operation"] = p.Operation
	}
	if p.RecordType != nil {
		m["RecordType"] = p.RecordType
	}
	if p.RetentionDuration != nil {
		m["RetentionDuration"] = p.RetentionDuration
	}
	if p.UserId != "" {
		m["UserId"] = p.UserId
	}
	return m
}

// GetUnifiedAuditLogRetentionPolicy runs the Get-UnifiedAuditLogRetentionPolicy cmdlet.
func (s *Service) GetUnifiedAuditLogRetentionPolicy(ctx context.Context, p GetUnifiedAuditLogRetentionPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-UnifiedAuditLogRetentionPolicy", p.params())
}

// GetUserParams are the parameters of Get-User.
// DefaultParameterSetName: Identity
type GetUserParams struct {
	Filter               string   `ps:"Filter"`
	Identity             any      `ps:"Identity"`
	PublicFolder         bool     `ps:"PublicFolder"`
	RecipientTypeDetails []string `ps:"RecipientTypeDetails"`
	ResultSize           any      `ps:"ResultSize"`
	SortBy               string   `ps:"SortBy"`
}

func (p GetUserParams) params() map[string]any {
	m := map[string]any{}
	if p.Filter != "" {
		m["Filter"] = p.Filter
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.PublicFolder {
		m["PublicFolder"] = true
	}
	if len(p.RecipientTypeDetails) > 0 {
		m["RecipientTypeDetails"] = p.RecipientTypeDetails
	}
	if p.ResultSize != nil {
		m["ResultSize"] = p.ResultSize
	}
	if p.SortBy != "" {
		m["SortBy"] = p.SortBy
	}
	return m
}

// GetUser runs the Get-User cmdlet.
func (s *Service) GetUser(ctx context.Context, p GetUserParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Get-User", p.params())
}

// ImportDlpComplianceRuleCollectionParams are the parameters of Import-DlpComplianceRuleCollection.
type ImportDlpComplianceRuleCollectionParams struct {
	ExtendedWorkloadPolicies []string `ps:"ExtendedWorkloadPolicies"`
	FileData                 []string `ps:"FileData"`
	Force                    bool     `ps:"Force"`
	ImportErrorAction        any      `ps:"ImportErrorAction"`
	ImportMode               any      `ps:"ImportMode"`
	PolicyNames              any      `ps:"PolicyNames"`
}

func (p ImportDlpComplianceRuleCollectionParams) params() map[string]any {
	m := map[string]any{}
	if len(p.ExtendedWorkloadPolicies) > 0 {
		m["ExtendedWorkloadPolicies"] = p.ExtendedWorkloadPolicies
	}
	if len(p.FileData) > 0 {
		m["FileData"] = p.FileData
	}
	if p.Force {
		m["Force"] = true
	}
	if p.ImportErrorAction != nil {
		m["ImportErrorAction"] = p.ImportErrorAction
	}
	if p.ImportMode != nil {
		m["ImportMode"] = p.ImportMode
	}
	if p.PolicyNames != nil {
		m["PolicyNames"] = p.PolicyNames
	}
	return m
}

// ImportDlpComplianceRuleCollection runs the Import-DlpComplianceRuleCollection cmdlet.
func (s *Service) ImportDlpComplianceRuleCollection(ctx context.Context, p ImportDlpComplianceRuleCollectionParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Import-DlpComplianceRuleCollection", p.params())
}

// ImportExchangeDlpPolicyParams are the parameters of Import-ExchangeDlpPolicy.
type ImportExchangeDlpPolicyParams struct {
	Check bool `ps:"Check"`
}

func (p ImportExchangeDlpPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Check {
		m["Check"] = true
	}
	return m
}

// ImportExchangeDlpPolicy runs the Import-ExchangeDlpPolicy cmdlet.
func (s *Service) ImportExchangeDlpPolicy(ctx context.Context, p ImportExchangeDlpPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Import-ExchangeDlpPolicy", p.params())
}

// ImportFilePlanPropertyParams are the parameters of Import-FilePlanProperty.
// DefaultParameterSetName: Identity
type ImportFilePlanPropertyParams struct {
	DomainController any    `ps:"DomainController"`
	Force            bool   `ps:"Force"`
	RawCsv           string `ps:"RawCsv"`
	ValidateOnly     bool   `ps:"ValidateOnly"`
}

func (p ImportFilePlanPropertyParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.Force {
		m["Force"] = true
	}
	if p.RawCsv != "" {
		m["RawCsv"] = p.RawCsv
	}
	if p.ValidateOnly {
		m["ValidateOnly"] = true
	}
	return m
}

// ImportFilePlanProperty runs the Import-FilePlanProperty cmdlet.
func (s *Service) ImportFilePlanProperty(ctx context.Context, p ImportFilePlanPropertyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Import-FilePlanProperty", p.params())
}

// InstallUnifiedCompliancePrerequisiteParams are the parameters of Install-UnifiedCompliancePrerequisite.
// DefaultParameterSetName: Initialize
type InstallUnifiedCompliancePrerequisiteParams struct {
	ForceInitialize       bool `ps:"ForceInitialize"`
	LoadOnly              bool `ps:"LoadOnly"`
	PolicyCenterSiteOwner any  `ps:"PolicyCenterSiteOwner"`
}

func (p InstallUnifiedCompliancePrerequisiteParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceInitialize {
		m["ForceInitialize"] = true
	}
	if p.LoadOnly {
		m["LoadOnly"] = true
	}
	if p.PolicyCenterSiteOwner != nil {
		m["PolicyCenterSiteOwner"] = p.PolicyCenterSiteOwner
	}
	return m
}

// InstallUnifiedCompliancePrerequisite runs the Install-UnifiedCompliancePrerequisite cmdlet.
func (s *Service) InstallUnifiedCompliancePrerequisite(ctx context.Context, p InstallUnifiedCompliancePrerequisiteParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Install-UnifiedCompliancePrerequisite", p.params())
}

// InvokeComplianceSecurityFilterActionParams are the parameters of Invoke-ComplianceSecurityFilterAction.
type InvokeComplianceSecurityFilterActionParams struct {
	Action        string `ps:"Action"` // one of: GetStatus, Set
	EmailAddress  string `ps:"EmailAddress"`
	PropertyName  string `ps:"PropertyName"`
	PropertyValue string `ps:"PropertyValue"`
	SiteUrl       string `ps:"SiteUrl"`
}

func (p InvokeComplianceSecurityFilterActionParams) params() map[string]any {
	m := map[string]any{}
	if p.Action != "" {
		m["Action"] = p.Action
	}
	if p.EmailAddress != "" {
		m["EmailAddress"] = p.EmailAddress
	}
	if p.PropertyName != "" {
		m["PropertyName"] = p.PropertyName
	}
	if p.PropertyValue != "" {
		m["PropertyValue"] = p.PropertyValue
	}
	if p.SiteUrl != "" {
		m["SiteUrl"] = p.SiteUrl
	}
	return m
}

// InvokeComplianceSecurityFilterAction runs the Invoke-ComplianceSecurityFilterAction cmdlet.
func (s *Service) InvokeComplianceSecurityFilterAction(ctx context.Context, p InvokeComplianceSecurityFilterActionParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Invoke-ComplianceSecurityFilterAction", p.params())
}

// InvokeHoldRemovalActionParams are the parameters of Invoke-HoldRemovalAction.
type InvokeHoldRemovalActionParams struct {
	Action             any    `ps:"Action"` // one of: RemoveHold, GetHolds, GetHoldRemovals
	ExchangeLocation   string `ps:"ExchangeLocation"`
	Force              bool   `ps:"Force"`
	HoldId             string `ps:"HoldId"`
	SharePointLocation string `ps:"SharePointLocation"`
}

func (p InvokeHoldRemovalActionParams) params() map[string]any {
	m := map[string]any{}
	if p.Action != nil {
		m["Action"] = p.Action
	}
	if p.ExchangeLocation != "" {
		m["ExchangeLocation"] = p.ExchangeLocation
	}
	if p.Force {
		m["Force"] = true
	}
	if p.HoldId != "" {
		m["HoldId"] = p.HoldId
	}
	if p.SharePointLocation != "" {
		m["SharePointLocation"] = p.SharePointLocation
	}
	return m
}

// InvokeHoldRemovalAction runs the Invoke-HoldRemovalAction cmdlet.
func (s *Service) InvokeHoldRemovalAction(ctx context.Context, p InvokeHoldRemovalActionParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Invoke-HoldRemovalAction", p.params())
}

// InvokeMachineAssistedTagActionParams are the parameters of Invoke-MachineAssistedTagAction.
type InvokeMachineAssistedTagActionParams struct {
	Identity                        any    `ps:"Identity"`
	ModelIds                        any    `ps:"ModelIds"`
	ModuleId                        any    `ps:"ModuleId"`
	Resource                        any    `ps:"Resource"`
	TaskId                          any    `ps:"TaskId"`
	TrainingId                      any    `ps:"TrainingId"`
	TrueNegativeSharePointLocations any    `ps:"TrueNegativeSharePointLocations"`
	TruePositiveSharePointLocations any    `ps:"TruePositiveSharePointLocations"`
	Unsupervised                    string `ps:"Unsupervised"`
}

func (p InvokeMachineAssistedTagActionParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.ModelIds != nil {
		m["ModelIds"] = p.ModelIds
	}
	if p.ModuleId != nil {
		m["ModuleId"] = p.ModuleId
	}
	if p.Resource != nil {
		m["Resource"] = p.Resource
	}
	if p.TaskId != nil {
		m["TaskId"] = p.TaskId
	}
	if p.TrainingId != nil {
		m["TrainingId"] = p.TrainingId
	}
	if p.TrueNegativeSharePointLocations != nil {
		m["TrueNegativeSharePointLocations"] = p.TrueNegativeSharePointLocations
	}
	if p.TruePositiveSharePointLocations != nil {
		m["TruePositiveSharePointLocations"] = p.TruePositiveSharePointLocations
	}
	if p.Unsupervised != "" {
		m["Unsupervised"] = p.Unsupervised
	}
	return m
}

// InvokeMachineAssistedTagAction runs the Invoke-MachineAssistedTagAction cmdlet.
func (s *Service) InvokeMachineAssistedTagAction(ctx context.Context, p InvokeMachineAssistedTagActionParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Invoke-MachineAssistedTagAction", p.params())
}

// MigrateDlpFingerprintParams are the parameters of Migrate-DlpFingerprint.
type MigrateDlpFingerprintParams struct {
}

func (p MigrateDlpFingerprintParams) params() map[string]any {
	m := map[string]any{}
	return m
}

// MigrateDlpFingerprint runs the Migrate-DlpFingerprint cmdlet.
func (s *Service) MigrateDlpFingerprint(ctx context.Context, p MigrateDlpFingerprintParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Migrate-DlpFingerprint", p.params())
}

// NewActivityAlertParams are the parameters of New-ActivityAlert.
type NewActivityAlertParams struct {
	Category     any    `ps:"Category"`
	Condition    string `ps:"Condition"`
	Description  string `ps:"Description"`
	Disabled     bool   `ps:"Disabled"`
	EmailCulture any    `ps:"EmailCulture"`
	Multiplier   any    `ps:"Multiplier"`
	Name         string `ps:"Name"`
	NotifyUser   any    `ps:"NotifyUser"`
	Operation    any    `ps:"Operation"`
	RecordType   any    `ps:"RecordType"`
	ScopeLevel   any    `ps:"ScopeLevel"`
	Severity     any    `ps:"Severity"`
	Threshold    int    `ps:"Threshold"`
	TimeWindow   int    `ps:"TimeWindow"`
	Type         any    `ps:"Type"`
	UserId       any    `ps:"UserId"`
}

func (p NewActivityAlertParams) params() map[string]any {
	m := map[string]any{}
	if p.Category != nil {
		m["Category"] = p.Category
	}
	if p.Condition != "" {
		m["Condition"] = p.Condition
	}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if p.Disabled {
		m["Disabled"] = true
	}
	if p.EmailCulture != nil {
		m["EmailCulture"] = p.EmailCulture
	}
	if p.Multiplier != nil {
		m["Multiplier"] = p.Multiplier
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.NotifyUser != nil {
		m["NotifyUser"] = p.NotifyUser
	}
	if p.Operation != nil {
		m["Operation"] = p.Operation
	}
	if p.RecordType != nil {
		m["RecordType"] = p.RecordType
	}
	if p.ScopeLevel != nil {
		m["ScopeLevel"] = p.ScopeLevel
	}
	if p.Severity != nil {
		m["Severity"] = p.Severity
	}
	if p.Threshold != 0 {
		m["Threshold"] = p.Threshold
	}
	if p.TimeWindow != 0 {
		m["TimeWindow"] = p.TimeWindow
	}
	if p.Type != nil {
		m["Type"] = p.Type
	}
	if p.UserId != nil {
		m["UserId"] = p.UserId
	}
	return m
}

// NewActivityAlert runs the New-ActivityAlert cmdlet.
func (s *Service) NewActivityAlert(ctx context.Context, p NewActivityAlertParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-ActivityAlert", p.params())
}

// NewAdaptiveScopeParams are the parameters of New-AdaptiveScope.
type NewAdaptiveScopeParams struct {
	AdministrativeUnit           any      `ps:"AdministrativeUnit"`
	Comment                      string   `ps:"Comment"`
	EnabledStates                []string `ps:"EnabledStates"`
	FilterConditions             any      `ps:"FilterConditions"`
	LinkedRecipientEnabledStates []string `ps:"LinkedRecipientEnabledStates"`
	LocationType                 any      `ps:"LocationType"`
	Name                         string   `ps:"Name"`
	RawQuery                     string   `ps:"RawQuery"`
	UseKql                       any      `ps:"UseKql"`
}

func (p NewAdaptiveScopeParams) params() map[string]any {
	m := map[string]any{}
	if p.AdministrativeUnit != nil {
		m["AdministrativeUnit"] = p.AdministrativeUnit
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if len(p.EnabledStates) > 0 {
		m["EnabledStates"] = p.EnabledStates
	}
	if p.FilterConditions != nil {
		m["FilterConditions"] = p.FilterConditions
	}
	if len(p.LinkedRecipientEnabledStates) > 0 {
		m["LinkedRecipientEnabledStates"] = p.LinkedRecipientEnabledStates
	}
	if p.LocationType != nil {
		m["LocationType"] = p.LocationType
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.RawQuery != "" {
		m["RawQuery"] = p.RawQuery
	}
	if p.UseKql != nil {
		m["UseKql"] = p.UseKql
	}
	return m
}

// NewAdaptiveScope runs the New-AdaptiveScope cmdlet.
func (s *Service) NewAdaptiveScope(ctx context.Context, p NewAdaptiveScopeParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-AdaptiveScope", p.params())
}

// NewAdministrativeUnitExtensionParams are the parameters of New-AdministrativeUnitExtension.
type NewAdministrativeUnitExtensionParams struct {
	AdministrativeUnit           any      `ps:"AdministrativeUnit"`
	EnabledStates                []string `ps:"EnabledStates"`
	FilterConditions             any      `ps:"FilterConditions"`
	LinkedRecipientEnabledStates []string `ps:"LinkedRecipientEnabledStates"`
	LocationType                 any      `ps:"LocationType"`
	RawQuery                     string   `ps:"RawQuery"`
	UseKql                       any      `ps:"UseKql"`
}

func (p NewAdministrativeUnitExtensionParams) params() map[string]any {
	m := map[string]any{}
	if p.AdministrativeUnit != nil {
		m["AdministrativeUnit"] = p.AdministrativeUnit
	}
	if len(p.EnabledStates) > 0 {
		m["EnabledStates"] = p.EnabledStates
	}
	if p.FilterConditions != nil {
		m["FilterConditions"] = p.FilterConditions
	}
	if len(p.LinkedRecipientEnabledStates) > 0 {
		m["LinkedRecipientEnabledStates"] = p.LinkedRecipientEnabledStates
	}
	if p.LocationType != nil {
		m["LocationType"] = p.LocationType
	}
	if p.RawQuery != "" {
		m["RawQuery"] = p.RawQuery
	}
	if p.UseKql != nil {
		m["UseKql"] = p.UseKql
	}
	return m
}

// NewAdministrativeUnitExtension runs the New-AdministrativeUnitExtension cmdlet.
func (s *Service) NewAdministrativeUnitExtension(ctx context.Context, p NewAdministrativeUnitExtensionParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-AdministrativeUnitExtension", p.params())
}

// NewAppRetentionCompliancePolicyParams are the parameters of New-AppRetentionCompliancePolicy.
type NewAppRetentionCompliancePolicyParams struct {
	AdaptiveScopeLocation        any      `ps:"AdaptiveScopeLocation"`
	Applications                 []string `ps:"Applications"`
	Comment                      string   `ps:"Comment"`
	Enabled                      bool     `ps:"Enabled"`
	ExchangeLocation             any      `ps:"ExchangeLocation"`
	ExchangeLocationException    any      `ps:"ExchangeLocationException"`
	Force                        bool     `ps:"Force"`
	Locations                    string   `ps:"Locations"`
	ModernGroupLocation          any      `ps:"ModernGroupLocation"`
	ModernGroupLocationException any      `ps:"ModernGroupLocationException"`
	Name                         string   `ps:"Name"`
	PolicyRBACScopes             any      `ps:"PolicyRBACScopes"`
	RestrictiveRetention         bool     `ps:"RestrictiveRetention"`
}

func (p NewAppRetentionCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.AdaptiveScopeLocation != nil {
		m["AdaptiveScopeLocation"] = p.AdaptiveScopeLocation
	}
	if len(p.Applications) > 0 {
		m["Applications"] = p.Applications
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.ExchangeLocation != nil {
		m["ExchangeLocation"] = p.ExchangeLocation
	}
	if p.ExchangeLocationException != nil {
		m["ExchangeLocationException"] = p.ExchangeLocationException
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Locations != "" {
		m["Locations"] = p.Locations
	}
	if p.ModernGroupLocation != nil {
		m["ModernGroupLocation"] = p.ModernGroupLocation
	}
	if p.ModernGroupLocationException != nil {
		m["ModernGroupLocationException"] = p.ModernGroupLocationException
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.RestrictiveRetention {
		m["RestrictiveRetention"] = true
	}
	return m
}

// NewAppRetentionCompliancePolicy runs the New-AppRetentionCompliancePolicy cmdlet.
func (s *Service) NewAppRetentionCompliancePolicy(ctx context.Context, p NewAppRetentionCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-AppRetentionCompliancePolicy", p.params())
}

// NewAppRetentionComplianceRuleParams are the parameters of New-AppRetentionComplianceRule.
type NewAppRetentionComplianceRuleParams struct {
	ApplyComplianceTag                  string   `ps:"ApplyComplianceTag"`
	Comment                             string   `ps:"Comment"`
	ContentContainsSensitiveInformation []string `ps:"ContentContainsSensitiveInformation"`
	ContentMatchQuery                   string   `ps:"ContentMatchQuery"`
	ExcludedItemClasses                 any      `ps:"ExcludedItemClasses"`
	ExpirationDateOption                string   `ps:"ExpirationDateOption"`
	MachineLearningModelIDs             any      `ps:"MachineLearningModelIDs"`
	Name                                string   `ps:"Name"`
	Policy                              any      `ps:"Policy"`
	PublishComplianceTag                string   `ps:"PublishComplianceTag"`
	RetentionComplianceAction           string   `ps:"RetentionComplianceAction"`
	RetentionDuration                   any      `ps:"RetentionDuration"`
	RetentionDurationDisplayHint        any      `ps:"RetentionDurationDisplayHint"`
}

func (p NewAppRetentionComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.ApplyComplianceTag != "" {
		m["ApplyComplianceTag"] = p.ApplyComplianceTag
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if len(p.ContentContainsSensitiveInformation) > 0 {
		m["ContentContainsSensitiveInformation"] = p.ContentContainsSensitiveInformation
	}
	if p.ContentMatchQuery != "" {
		m["ContentMatchQuery"] = p.ContentMatchQuery
	}
	if p.ExcludedItemClasses != nil {
		m["ExcludedItemClasses"] = p.ExcludedItemClasses
	}
	if p.ExpirationDateOption != "" {
		m["ExpirationDateOption"] = p.ExpirationDateOption
	}
	if p.MachineLearningModelIDs != nil {
		m["MachineLearningModelIDs"] = p.MachineLearningModelIDs
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	if p.PublishComplianceTag != "" {
		m["PublishComplianceTag"] = p.PublishComplianceTag
	}
	if p.RetentionComplianceAction != "" {
		m["RetentionComplianceAction"] = p.RetentionComplianceAction
	}
	if p.RetentionDuration != nil {
		m["RetentionDuration"] = p.RetentionDuration
	}
	if p.RetentionDurationDisplayHint != nil {
		m["RetentionDurationDisplayHint"] = p.RetentionDurationDisplayHint
	}
	return m
}

// NewAppRetentionComplianceRule runs the New-AppRetentionComplianceRule cmdlet.
func (s *Service) NewAppRetentionComplianceRule(ctx context.Context, p NewAppRetentionComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-AppRetentionComplianceRule", p.params())
}

// NewAuditConfigurationPolicyParams are the parameters of New-AuditConfigurationPolicy.
type NewAuditConfigurationPolicyParams struct {
	DomainController any `ps:"DomainController"`
	Workload         any `ps:"Workload"`
}

func (p NewAuditConfigurationPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.Workload != nil {
		m["Workload"] = p.Workload
	}
	return m
}

// NewAuditConfigurationPolicy runs the New-AuditConfigurationPolicy cmdlet.
func (s *Service) NewAuditConfigurationPolicy(ctx context.Context, p NewAuditConfigurationPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-AuditConfigurationPolicy", p.params())
}

// NewAuditConfigurationRuleParams are the parameters of New-AuditConfigurationRule.
type NewAuditConfigurationRuleParams struct {
	AuditOperation   any `ps:"AuditOperation"`
	DomainController any `ps:"DomainController"`
	Workload         any `ps:"Workload"`
}

func (p NewAuditConfigurationRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.AuditOperation != nil {
		m["AuditOperation"] = p.AuditOperation
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.Workload != nil {
		m["Workload"] = p.Workload
	}
	return m
}

// NewAuditConfigurationRule runs the New-AuditConfigurationRule cmdlet.
func (s *Service) NewAuditConfigurationRule(ctx context.Context, p NewAuditConfigurationRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-AuditConfigurationRule", p.params())
}

// NewAutoSensitivityLabelPolicyParams are the parameters of New-AutoSensitivityLabelPolicy.
type NewAutoSensitivityLabelPolicyParams struct {
	ApplySensitivityLabel                   string   `ps:"ApplySensitivityLabel"`
	ApplySensitivityLabelOverwriteWorkloads any      `ps:"ApplySensitivityLabelOverwriteWorkloads"`
	Comment                                 string   `ps:"Comment"`
	EndpointDlpLocation                     any      `ps:"EndpointDlpLocation"`
	EndpointDlpLocationException            any      `ps:"EndpointDlpLocationException"`
	EnforcementPlanes                       any      `ps:"EnforcementPlanes"`
	ExceptIfOneDriveSharedBy                []string `ps:"ExceptIfOneDriveSharedBy"`
	ExceptIfOneDriveSharedByMemberOf        []string `ps:"ExceptIfOneDriveSharedByMemberOf"`
	ExchangeAdaptiveScopes                  any      `ps:"ExchangeAdaptiveScopes"`
	ExchangeAdaptiveScopesException         any      `ps:"ExchangeAdaptiveScopesException"`
	ExchangeLocation                        any      `ps:"ExchangeLocation"`
	ExchangeSender                          []string `ps:"ExchangeSender"`
	ExchangeSenderException                 []string `ps:"ExchangeSenderException"`
	ExchangeSenderMemberOf                  []string `ps:"ExchangeSenderMemberOf"`
	ExchangeSenderMemberOfException         []string `ps:"ExchangeSenderMemberOfException"`
	ExternalMailRightsManagementOwner       any      `ps:"ExternalMailRightsManagementOwner"`
	Force                                   bool     `ps:"Force"`
	Locations                               string   `ps:"Locations"`
	Mode                                    any      `ps:"Mode"`
	Name                                    string   `ps:"Name"`
	OneDriveAdaptiveScopes                  any      `ps:"OneDriveAdaptiveScopes"`
	OneDriveAdaptiveScopesException         any      `ps:"OneDriveAdaptiveScopesException"`
	OneDriveLocation                        any      `ps:"OneDriveLocation"`
	OneDriveLocationException               any      `ps:"OneDriveLocationException"`
	OneDriveSharedBy                        []string `ps:"OneDriveSharedBy"`
	OneDriveSharedByMemberOf                []string `ps:"OneDriveSharedByMemberOf"`
	OverwriteLabel                          bool     `ps:"OverwriteLabel"`
	PolicyRBACScopes                        any      `ps:"PolicyRBACScopes"`
	PolicyTemplateInfo                      any      `ps:"PolicyTemplateInfo"`
	PowerBILocation                         any      `ps:"PowerBILocation"`
	PowerBILocationException                any      `ps:"PowerBILocationException"`
	Priority                                any      `ps:"Priority"`
	RemoveLabel                             bool     `ps:"RemoveLabel"`
	SharePointAdaptiveScopes                any      `ps:"SharePointAdaptiveScopes"`
	SharePointAdaptiveScopesException       any      `ps:"SharePointAdaptiveScopesException"`
	SharePointLocation                      any      `ps:"SharePointLocation"`
	SharePointLocationException             any      `ps:"SharePointLocationException"`
	UnifiedAuditLogEnabled                  bool     `ps:"UnifiedAuditLogEnabled"`
}

func (p NewAutoSensitivityLabelPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.ApplySensitivityLabel != "" {
		m["ApplySensitivityLabel"] = p.ApplySensitivityLabel
	}
	if p.ApplySensitivityLabelOverwriteWorkloads != nil {
		m["ApplySensitivityLabelOverwriteWorkloads"] = p.ApplySensitivityLabelOverwriteWorkloads
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.EndpointDlpLocation != nil {
		m["EndpointDlpLocation"] = p.EndpointDlpLocation
	}
	if p.EndpointDlpLocationException != nil {
		m["EndpointDlpLocationException"] = p.EndpointDlpLocationException
	}
	if p.EnforcementPlanes != nil {
		m["EnforcementPlanes"] = p.EnforcementPlanes
	}
	if len(p.ExceptIfOneDriveSharedBy) > 0 {
		m["ExceptIfOneDriveSharedBy"] = p.ExceptIfOneDriveSharedBy
	}
	if len(p.ExceptIfOneDriveSharedByMemberOf) > 0 {
		m["ExceptIfOneDriveSharedByMemberOf"] = p.ExceptIfOneDriveSharedByMemberOf
	}
	if p.ExchangeAdaptiveScopes != nil {
		m["ExchangeAdaptiveScopes"] = p.ExchangeAdaptiveScopes
	}
	if p.ExchangeAdaptiveScopesException != nil {
		m["ExchangeAdaptiveScopesException"] = p.ExchangeAdaptiveScopesException
	}
	if p.ExchangeLocation != nil {
		m["ExchangeLocation"] = p.ExchangeLocation
	}
	if len(p.ExchangeSender) > 0 {
		m["ExchangeSender"] = p.ExchangeSender
	}
	if len(p.ExchangeSenderException) > 0 {
		m["ExchangeSenderException"] = p.ExchangeSenderException
	}
	if len(p.ExchangeSenderMemberOf) > 0 {
		m["ExchangeSenderMemberOf"] = p.ExchangeSenderMemberOf
	}
	if len(p.ExchangeSenderMemberOfException) > 0 {
		m["ExchangeSenderMemberOfException"] = p.ExchangeSenderMemberOfException
	}
	if p.ExternalMailRightsManagementOwner != nil {
		m["ExternalMailRightsManagementOwner"] = p.ExternalMailRightsManagementOwner
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Locations != "" {
		m["Locations"] = p.Locations
	}
	if p.Mode != nil {
		m["Mode"] = p.Mode
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.OneDriveAdaptiveScopes != nil {
		m["OneDriveAdaptiveScopes"] = p.OneDriveAdaptiveScopes
	}
	if p.OneDriveAdaptiveScopesException != nil {
		m["OneDriveAdaptiveScopesException"] = p.OneDriveAdaptiveScopesException
	}
	if p.OneDriveLocation != nil {
		m["OneDriveLocation"] = p.OneDriveLocation
	}
	if p.OneDriveLocationException != nil {
		m["OneDriveLocationException"] = p.OneDriveLocationException
	}
	if len(p.OneDriveSharedBy) > 0 {
		m["OneDriveSharedBy"] = p.OneDriveSharedBy
	}
	if len(p.OneDriveSharedByMemberOf) > 0 {
		m["OneDriveSharedByMemberOf"] = p.OneDriveSharedByMemberOf
	}
	if p.OverwriteLabel {
		m["OverwriteLabel"] = true
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.PolicyTemplateInfo != nil {
		m["PolicyTemplateInfo"] = p.PolicyTemplateInfo
	}
	if p.PowerBILocation != nil {
		m["PowerBILocation"] = p.PowerBILocation
	}
	if p.PowerBILocationException != nil {
		m["PowerBILocationException"] = p.PowerBILocationException
	}
	if p.Priority != nil {
		m["Priority"] = p.Priority
	}
	if p.RemoveLabel {
		m["RemoveLabel"] = true
	}
	if p.SharePointAdaptiveScopes != nil {
		m["SharePointAdaptiveScopes"] = p.SharePointAdaptiveScopes
	}
	if p.SharePointAdaptiveScopesException != nil {
		m["SharePointAdaptiveScopesException"] = p.SharePointAdaptiveScopesException
	}
	if p.SharePointLocation != nil {
		m["SharePointLocation"] = p.SharePointLocation
	}
	if p.SharePointLocationException != nil {
		m["SharePointLocationException"] = p.SharePointLocationException
	}
	if p.UnifiedAuditLogEnabled {
		m["UnifiedAuditLogEnabled"] = true
	}
	return m
}

// NewAutoSensitivityLabelPolicy runs the New-AutoSensitivityLabelPolicy cmdlet.
func (s *Service) NewAutoSensitivityLabelPolicy(ctx context.Context, p NewAutoSensitivityLabelPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-AutoSensitivityLabelPolicy", p.params())
}

// NewAutoSensitivityLabelRuleParams are the parameters of New-AutoSensitivityLabelRule.
type NewAutoSensitivityLabelRuleParams struct {
	AccessScope                                  any      `ps:"AccessScope"`
	ActivationDate                               any      `ps:"ActivationDate"`
	AdvancedRule                                 string   `ps:"AdvancedRule"`
	AnyOfRecipientAddressContainsWords           any      `ps:"AnyOfRecipientAddressContainsWords"`
	AnyOfRecipientAddressMatchesPatterns         any      `ps:"AnyOfRecipientAddressMatchesPatterns"`
	Comment                                      string   `ps:"Comment"`
	ContentContainsSensitiveInformation          []string `ps:"ContentContainsSensitiveInformation"`
	ContentExtensionMatchesWords                 any      `ps:"ContentExtensionMatchesWords"`
	ContentIsNotLabeled                          bool     `ps:"ContentIsNotLabeled"`
	ContentPropertyContainsWords                 any      `ps:"ContentPropertyContainsWords"`
	DefaultSpoDocLibraryHasLabel                 bool     `ps:"DefaultSpoDocLibraryHasLabel"`
	Disabled                                     bool     `ps:"Disabled"`
	DocumentCreatedBy                            any      `ps:"DocumentCreatedBy"`
	DocumentIsPasswordProtected                  bool     `ps:"DocumentIsPasswordProtected"`
	DocumentIsUnsupported                        bool     `ps:"DocumentIsUnsupported"`
	DocumentNameMatchesWords                     any      `ps:"DocumentNameMatchesWords"`
	DocumentSizeOver                             any      `ps:"DocumentSizeOver"`
	ExceptIfAccessScope                          any      `ps:"ExceptIfAccessScope"`
	ExceptIfAnyOfRecipientAddressContainsWords   any      `ps:"ExceptIfAnyOfRecipientAddressContainsWords"`
	ExceptIfAnyOfRecipientAddressMatchesPatterns any      `ps:"ExceptIfAnyOfRecipientAddressMatchesPatterns"`
	ExceptIfContentContainsSensitiveInformation  []string `ps:"ExceptIfContentContainsSensitiveInformation"`
	ExceptIfContentExtensionMatchesWords         any      `ps:"ExceptIfContentExtensionMatchesWords"`
	ExceptIfContentPropertyContainsWords         any      `ps:"ExceptIfContentPropertyContainsWords"`
	ExceptIfDocumentCreatedBy                    any      `ps:"ExceptIfDocumentCreatedBy"`
	ExceptIfDocumentIsPasswordProtected          bool     `ps:"ExceptIfDocumentIsPasswordProtected"`
	ExceptIfDocumentIsUnsupported                bool     `ps:"ExceptIfDocumentIsUnsupported"`
	ExceptIfDocumentNameMatchesWords             any      `ps:"ExceptIfDocumentNameMatchesWords"`
	ExceptIfDocumentSizeOver                     any      `ps:"ExceptIfDocumentSizeOver"`
	ExceptIfFrom                                 []string `ps:"ExceptIfFrom"`
	ExceptIfFromAddressContainsWords             any      `ps:"ExceptIfFromAddressContainsWords"`
	ExceptIfFromAddressMatchesPatterns           any      `ps:"ExceptIfFromAddressMatchesPatterns"`
	ExceptIfFromMemberOf                         []string `ps:"ExceptIfFromMemberOf"`
	ExceptIfHeaderMatchesPatterns                any      `ps:"ExceptIfHeaderMatchesPatterns"`
	ExceptIfProcessingLimitExceeded              bool     `ps:"ExceptIfProcessingLimitExceeded"`
	ExceptIfRecipientDomainIs                    any      `ps:"ExceptIfRecipientDomainIs"`
	ExceptIfSenderDomainIs                       any      `ps:"ExceptIfSenderDomainIs"`
	ExceptIfSenderIPRanges                       any      `ps:"ExceptIfSenderIPRanges"`
	ExceptIfSentTo                               any      `ps:"ExceptIfSentTo"`
	ExceptIfSentToMemberOf                       []string `ps:"ExceptIfSentToMemberOf"`
	ExceptIfSharedWithDomain                     any      `ps:"ExceptIfSharedWithDomain"`
	ExceptIfSubjectMatchesPatterns               any      `ps:"ExceptIfSubjectMatchesPatterns"`
	ExpiryDate                                   any      `ps:"ExpiryDate"`
	From                                         []string `ps:"From"`
	FromAddressContainsWords                     any      `ps:"FromAddressContainsWords"`
	FromAddressMatchesPatterns                   any      `ps:"FromAddressMatchesPatterns"`
	FromMemberOf                                 []string `ps:"FromMemberOf"`
	HeaderMatchesPatterns                        any      `ps:"HeaderMatchesPatterns"`
	ImmutableId                                  any      `ps:"ImmutableId"`
	Name                                         string   `ps:"Name"`
	Policy                                       any      `ps:"Policy"`
	Priority                                     any      `ps:"Priority"`
	ProcessingLimitExceeded                      bool     `ps:"ProcessingLimitExceeded"`
	RecipientDomainIs                            any      `ps:"RecipientDomainIs"`
	ReportSeverityLevel                          any      `ps:"ReportSeverityLevel"`
	RuleErrorAction                              any      `ps:"RuleErrorAction"`
	SenderDomainIs                               any      `ps:"SenderDomainIs"`
	SenderIPRanges                               any      `ps:"SenderIPRanges"`
	SentTo                                       any      `ps:"SentTo"`
	SentToMemberOf                               []string `ps:"SentToMemberOf"`
	SharedWithDomain                             any      `ps:"SharedWithDomain"`
	SourceType                                   string   `ps:"SourceType"`
	SubjectMatchesPatterns                       any      `ps:"SubjectMatchesPatterns"`
	Workload                                     any      `ps:"Workload"`
}

func (p NewAutoSensitivityLabelRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.AccessScope != nil {
		m["AccessScope"] = p.AccessScope
	}
	if p.ActivationDate != nil {
		m["ActivationDate"] = p.ActivationDate
	}
	if p.AdvancedRule != "" {
		m["AdvancedRule"] = p.AdvancedRule
	}
	if p.AnyOfRecipientAddressContainsWords != nil {
		m["AnyOfRecipientAddressContainsWords"] = p.AnyOfRecipientAddressContainsWords
	}
	if p.AnyOfRecipientAddressMatchesPatterns != nil {
		m["AnyOfRecipientAddressMatchesPatterns"] = p.AnyOfRecipientAddressMatchesPatterns
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if len(p.ContentContainsSensitiveInformation) > 0 {
		m["ContentContainsSensitiveInformation"] = p.ContentContainsSensitiveInformation
	}
	if p.ContentExtensionMatchesWords != nil {
		m["ContentExtensionMatchesWords"] = p.ContentExtensionMatchesWords
	}
	if p.ContentIsNotLabeled {
		m["ContentIsNotLabeled"] = true
	}
	if p.ContentPropertyContainsWords != nil {
		m["ContentPropertyContainsWords"] = p.ContentPropertyContainsWords
	}
	if p.DefaultSpoDocLibraryHasLabel {
		m["DefaultSpoDocLibraryHasLabel"] = true
	}
	if p.Disabled {
		m["Disabled"] = true
	}
	if p.DocumentCreatedBy != nil {
		m["DocumentCreatedBy"] = p.DocumentCreatedBy
	}
	if p.DocumentIsPasswordProtected {
		m["DocumentIsPasswordProtected"] = true
	}
	if p.DocumentIsUnsupported {
		m["DocumentIsUnsupported"] = true
	}
	if p.DocumentNameMatchesWords != nil {
		m["DocumentNameMatchesWords"] = p.DocumentNameMatchesWords
	}
	if p.DocumentSizeOver != nil {
		m["DocumentSizeOver"] = p.DocumentSizeOver
	}
	if p.ExceptIfAccessScope != nil {
		m["ExceptIfAccessScope"] = p.ExceptIfAccessScope
	}
	if p.ExceptIfAnyOfRecipientAddressContainsWords != nil {
		m["ExceptIfAnyOfRecipientAddressContainsWords"] = p.ExceptIfAnyOfRecipientAddressContainsWords
	}
	if p.ExceptIfAnyOfRecipientAddressMatchesPatterns != nil {
		m["ExceptIfAnyOfRecipientAddressMatchesPatterns"] = p.ExceptIfAnyOfRecipientAddressMatchesPatterns
	}
	if len(p.ExceptIfContentContainsSensitiveInformation) > 0 {
		m["ExceptIfContentContainsSensitiveInformation"] = p.ExceptIfContentContainsSensitiveInformation
	}
	if p.ExceptIfContentExtensionMatchesWords != nil {
		m["ExceptIfContentExtensionMatchesWords"] = p.ExceptIfContentExtensionMatchesWords
	}
	if p.ExceptIfContentPropertyContainsWords != nil {
		m["ExceptIfContentPropertyContainsWords"] = p.ExceptIfContentPropertyContainsWords
	}
	if p.ExceptIfDocumentCreatedBy != nil {
		m["ExceptIfDocumentCreatedBy"] = p.ExceptIfDocumentCreatedBy
	}
	if p.ExceptIfDocumentIsPasswordProtected {
		m["ExceptIfDocumentIsPasswordProtected"] = true
	}
	if p.ExceptIfDocumentIsUnsupported {
		m["ExceptIfDocumentIsUnsupported"] = true
	}
	if p.ExceptIfDocumentNameMatchesWords != nil {
		m["ExceptIfDocumentNameMatchesWords"] = p.ExceptIfDocumentNameMatchesWords
	}
	if p.ExceptIfDocumentSizeOver != nil {
		m["ExceptIfDocumentSizeOver"] = p.ExceptIfDocumentSizeOver
	}
	if len(p.ExceptIfFrom) > 0 {
		m["ExceptIfFrom"] = p.ExceptIfFrom
	}
	if p.ExceptIfFromAddressContainsWords != nil {
		m["ExceptIfFromAddressContainsWords"] = p.ExceptIfFromAddressContainsWords
	}
	if p.ExceptIfFromAddressMatchesPatterns != nil {
		m["ExceptIfFromAddressMatchesPatterns"] = p.ExceptIfFromAddressMatchesPatterns
	}
	if len(p.ExceptIfFromMemberOf) > 0 {
		m["ExceptIfFromMemberOf"] = p.ExceptIfFromMemberOf
	}
	if p.ExceptIfHeaderMatchesPatterns != nil {
		m["ExceptIfHeaderMatchesPatterns"] = p.ExceptIfHeaderMatchesPatterns
	}
	if p.ExceptIfProcessingLimitExceeded {
		m["ExceptIfProcessingLimitExceeded"] = true
	}
	if p.ExceptIfRecipientDomainIs != nil {
		m["ExceptIfRecipientDomainIs"] = p.ExceptIfRecipientDomainIs
	}
	if p.ExceptIfSenderDomainIs != nil {
		m["ExceptIfSenderDomainIs"] = p.ExceptIfSenderDomainIs
	}
	if p.ExceptIfSenderIPRanges != nil {
		m["ExceptIfSenderIPRanges"] = p.ExceptIfSenderIPRanges
	}
	if p.ExceptIfSentTo != nil {
		m["ExceptIfSentTo"] = p.ExceptIfSentTo
	}
	if len(p.ExceptIfSentToMemberOf) > 0 {
		m["ExceptIfSentToMemberOf"] = p.ExceptIfSentToMemberOf
	}
	if p.ExceptIfSharedWithDomain != nil {
		m["ExceptIfSharedWithDomain"] = p.ExceptIfSharedWithDomain
	}
	if p.ExceptIfSubjectMatchesPatterns != nil {
		m["ExceptIfSubjectMatchesPatterns"] = p.ExceptIfSubjectMatchesPatterns
	}
	if p.ExpiryDate != nil {
		m["ExpiryDate"] = p.ExpiryDate
	}
	if len(p.From) > 0 {
		m["From"] = p.From
	}
	if p.FromAddressContainsWords != nil {
		m["FromAddressContainsWords"] = p.FromAddressContainsWords
	}
	if p.FromAddressMatchesPatterns != nil {
		m["FromAddressMatchesPatterns"] = p.FromAddressMatchesPatterns
	}
	if len(p.FromMemberOf) > 0 {
		m["FromMemberOf"] = p.FromMemberOf
	}
	if p.HeaderMatchesPatterns != nil {
		m["HeaderMatchesPatterns"] = p.HeaderMatchesPatterns
	}
	if p.ImmutableId != nil {
		m["ImmutableId"] = p.ImmutableId
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	if p.Priority != nil {
		m["Priority"] = p.Priority
	}
	if p.ProcessingLimitExceeded {
		m["ProcessingLimitExceeded"] = true
	}
	if p.RecipientDomainIs != nil {
		m["RecipientDomainIs"] = p.RecipientDomainIs
	}
	if p.ReportSeverityLevel != nil {
		m["ReportSeverityLevel"] = p.ReportSeverityLevel
	}
	if p.RuleErrorAction != nil {
		m["RuleErrorAction"] = p.RuleErrorAction
	}
	if p.SenderDomainIs != nil {
		m["SenderDomainIs"] = p.SenderDomainIs
	}
	if p.SenderIPRanges != nil {
		m["SenderIPRanges"] = p.SenderIPRanges
	}
	if p.SentTo != nil {
		m["SentTo"] = p.SentTo
	}
	if len(p.SentToMemberOf) > 0 {
		m["SentToMemberOf"] = p.SentToMemberOf
	}
	if p.SharedWithDomain != nil {
		m["SharedWithDomain"] = p.SharedWithDomain
	}
	if p.SourceType != "" {
		m["SourceType"] = p.SourceType
	}
	if p.SubjectMatchesPatterns != nil {
		m["SubjectMatchesPatterns"] = p.SubjectMatchesPatterns
	}
	if p.Workload != nil {
		m["Workload"] = p.Workload
	}
	return m
}

// NewAutoSensitivityLabelRule runs the New-AutoSensitivityLabelRule cmdlet.
func (s *Service) NewAutoSensitivityLabelRule(ctx context.Context, p NewAutoSensitivityLabelRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-AutoSensitivityLabelRule", p.params())
}

// NewCaseHoldPolicyParams are the parameters of New-CaseHoldPolicy.
type NewCaseHoldPolicyParams struct {
	Case                 string `ps:"Case"`
	Comment              string `ps:"Comment"`
	Enabled              bool   `ps:"Enabled"`
	ExchangeLocation     any    `ps:"ExchangeLocation"`
	Force                bool   `ps:"Force"`
	Name                 string `ps:"Name"`
	PublicFolderLocation any    `ps:"PublicFolderLocation"`
	SharePointLocation   any    `ps:"SharePointLocation"`
}

func (p NewCaseHoldPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Case != "" {
		m["Case"] = p.Case
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.ExchangeLocation != nil {
		m["ExchangeLocation"] = p.ExchangeLocation
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.PublicFolderLocation != nil {
		m["PublicFolderLocation"] = p.PublicFolderLocation
	}
	if p.SharePointLocation != nil {
		m["SharePointLocation"] = p.SharePointLocation
	}
	return m
}

// NewCaseHoldPolicy runs the New-CaseHoldPolicy cmdlet.
func (s *Service) NewCaseHoldPolicy(ctx context.Context, p NewCaseHoldPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-CaseHoldPolicy", p.params())
}

// NewCaseHoldRuleParams are the parameters of New-CaseHoldRule.
type NewCaseHoldRuleParams struct {
	Comment           string `ps:"Comment"`
	ContentDateFrom   any    `ps:"ContentDateFrom"`
	ContentDateTo     any    `ps:"ContentDateTo"`
	ContentMatchQuery string `ps:"ContentMatchQuery"`
	Disabled          bool   `ps:"Disabled"`
	Name              string `ps:"Name"`
	Policy            any    `ps:"Policy"`
}

func (p NewCaseHoldRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.ContentDateFrom != nil {
		m["ContentDateFrom"] = p.ContentDateFrom
	}
	if p.ContentDateTo != nil {
		m["ContentDateTo"] = p.ContentDateTo
	}
	if p.ContentMatchQuery != "" {
		m["ContentMatchQuery"] = p.ContentMatchQuery
	}
	if p.Disabled {
		m["Disabled"] = true
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	return m
}

// NewCaseHoldRule runs the New-CaseHoldRule cmdlet.
func (s *Service) NewCaseHoldRule(ctx context.Context, p NewCaseHoldRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-CaseHoldRule", p.params())
}

// NewClassificationGradingPolicyParams are the parameters of New-ClassificationGradingPolicy.
type NewClassificationGradingPolicyParams struct {
	EndDateTime                 any      `ps:"EndDateTime"`
	GraderId                    any      `ps:"GraderId"`
	Name                        string   `ps:"Name"`
	SensitiveInformationTypeIds []string `ps:"SensitiveInformationTypeIds"`
	StartDateTime               any      `ps:"StartDateTime"`
}

func (p NewClassificationGradingPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.EndDateTime != nil {
		m["EndDateTime"] = p.EndDateTime
	}
	if p.GraderId != nil {
		m["GraderId"] = p.GraderId
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if len(p.SensitiveInformationTypeIds) > 0 {
		m["SensitiveInformationTypeIds"] = p.SensitiveInformationTypeIds
	}
	if p.StartDateTime != nil {
		m["StartDateTime"] = p.StartDateTime
	}
	return m
}

// NewClassificationGradingPolicy runs the New-ClassificationGradingPolicy cmdlet.
func (s *Service) NewClassificationGradingPolicy(ctx context.Context, p NewClassificationGradingPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-ClassificationGradingPolicy", p.params())
}

// NewComplianceBoundaryParams are the parameters of New-ComplianceBoundary.
type NewComplianceBoundaryParams struct {
	Comment               string `ps:"Comment"`
	CustomRecipientFilter string `ps:"CustomRecipientFilter"`
	Name                  string `ps:"Name"`
}

func (p NewComplianceBoundaryParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.CustomRecipientFilter != "" {
		m["CustomRecipientFilter"] = p.CustomRecipientFilter
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	return m
}

// NewComplianceBoundary runs the New-ComplianceBoundary cmdlet.
func (s *Service) NewComplianceBoundary(ctx context.Context, p NewComplianceBoundaryParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-ComplianceBoundary", p.params())
}

// NewComplianceCaseParams are the parameters of New-ComplianceCase.
// DefaultParameterSetName: Identity
type NewComplianceCaseParams struct {
	CaseType          any      `ps:"CaseType"`
	Description       string   `ps:"Description"`
	DomainController  any      `ps:"DomainController"`
	ExternalId        string   `ps:"ExternalId"`
	Name              string   `ps:"Name"`
	SecondaryCaseType string   `ps:"SecondaryCaseType"`
	SourceCaseType    string   `ps:"SourceCaseType"`
	Sources           []string `ps:"Sources"`
}

func (p NewComplianceCaseParams) params() map[string]any {
	m := map[string]any{}
	if p.CaseType != nil {
		m["CaseType"] = p.CaseType
	}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.ExternalId != "" {
		m["ExternalId"] = p.ExternalId
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.SecondaryCaseType != "" {
		m["SecondaryCaseType"] = p.SecondaryCaseType
	}
	if p.SourceCaseType != "" {
		m["SourceCaseType"] = p.SourceCaseType
	}
	if len(p.Sources) > 0 {
		m["Sources"] = p.Sources
	}
	return m
}

// NewComplianceCase runs the New-ComplianceCase cmdlet.
func (s *Service) NewComplianceCase(ctx context.Context, p NewComplianceCaseParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-ComplianceCase", p.params())
}

// NewComplianceCustodianParams are the parameters of New-ComplianceCustodian.
// DefaultParameterSetName: Identity
type NewComplianceCustodianParams struct {
	Case                 string   `ps:"Case"`
	ContactEmail         string   `ps:"ContactEmail"`
	ExchangeLocation     []string `ps:"ExchangeLocation"`
	Name                 string   `ps:"Name"`
	PublicFolderLocation []string `ps:"PublicFolderLocation"`
	SharePointLocation   []string `ps:"SharePointLocation"`
}

func (p NewComplianceCustodianParams) params() map[string]any {
	m := map[string]any{}
	if p.Case != "" {
		m["Case"] = p.Case
	}
	if p.ContactEmail != "" {
		m["ContactEmail"] = p.ContactEmail
	}
	if len(p.ExchangeLocation) > 0 {
		m["ExchangeLocation"] = p.ExchangeLocation
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if len(p.PublicFolderLocation) > 0 {
		m["PublicFolderLocation"] = p.PublicFolderLocation
	}
	if len(p.SharePointLocation) > 0 {
		m["SharePointLocation"] = p.SharePointLocation
	}
	return m
}

// NewComplianceCustodian runs the New-ComplianceCustodian cmdlet.
func (s *Service) NewComplianceCustodian(ctx context.Context, p NewComplianceCustodianParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-ComplianceCustodian", p.params())
}

// NewComplianceNoticeParams are the parameters of New-ComplianceNotice.
// DefaultParameterSetName: Identity
type NewComplianceNoticeParams struct {
	Case           string   `ps:"Case"`
	ContentBody    string   `ps:"ContentBody"`
	CustodianIds   []string `ps:"CustodianIds"`
	From           string   `ps:"From"`
	IsHighPriority bool     `ps:"IsHighPriority"`
	Name           string   `ps:"Name"`
	Subject        string   `ps:"Subject"`
	Type           string   `ps:"Type"`
	UrlPattern     string   `ps:"UrlPattern"`
}

func (p NewComplianceNoticeParams) params() map[string]any {
	m := map[string]any{}
	if p.Case != "" {
		m["Case"] = p.Case
	}
	if p.ContentBody != "" {
		m["ContentBody"] = p.ContentBody
	}
	if len(p.CustodianIds) > 0 {
		m["CustodianIds"] = p.CustodianIds
	}
	if p.From != "" {
		m["From"] = p.From
	}
	if p.IsHighPriority {
		m["IsHighPriority"] = true
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Subject != "" {
		m["Subject"] = p.Subject
	}
	if p.Type != "" {
		m["Type"] = p.Type
	}
	if p.UrlPattern != "" {
		m["UrlPattern"] = p.UrlPattern
	}
	return m
}

// NewComplianceNotice runs the New-ComplianceNotice cmdlet.
func (s *Service) NewComplianceNotice(ctx context.Context, p NewComplianceNoticeParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-ComplianceNotice", p.params())
}

// NewComplianceRetentionEventParams are the parameters of New-ComplianceRetentionEvent.
type NewComplianceRetentionEventParams struct {
	AssetId                string `ps:"AssetId"`
	Comment                string `ps:"Comment"`
	DomainController       any    `ps:"DomainController"`
	EventDateTime          any    `ps:"EventDateTime"`
	EventTags              any    `ps:"EventTags"`
	EventType              any    `ps:"EventType"`
	ExchangeAssetIdQuery   string `ps:"ExchangeAssetIdQuery"`
	Name                   string `ps:"Name"`
	PreviewOnly            bool   `ps:"PreviewOnly"`
	SharePointAssetIdQuery string `ps:"SharePointAssetIdQuery"`
}

func (p NewComplianceRetentionEventParams) params() map[string]any {
	m := map[string]any{}
	if p.AssetId != "" {
		m["AssetId"] = p.AssetId
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.EventDateTime != nil {
		m["EventDateTime"] = p.EventDateTime
	}
	if p.EventTags != nil {
		m["EventTags"] = p.EventTags
	}
	if p.EventType != nil {
		m["EventType"] = p.EventType
	}
	if p.ExchangeAssetIdQuery != "" {
		m["ExchangeAssetIdQuery"] = p.ExchangeAssetIdQuery
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.PreviewOnly {
		m["PreviewOnly"] = true
	}
	if p.SharePointAssetIdQuery != "" {
		m["SharePointAssetIdQuery"] = p.SharePointAssetIdQuery
	}
	return m
}

// NewComplianceRetentionEvent runs the New-ComplianceRetentionEvent cmdlet.
func (s *Service) NewComplianceRetentionEvent(ctx context.Context, p NewComplianceRetentionEventParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-ComplianceRetentionEvent", p.params())
}

// NewComplianceRetentionEventTypeParams are the parameters of New-ComplianceRetentionEventType.
type NewComplianceRetentionEventTypeParams struct {
	Comment string `ps:"Comment"`
	Name    string `ps:"Name"`
}

func (p NewComplianceRetentionEventTypeParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	return m
}

// NewComplianceRetentionEventType runs the New-ComplianceRetentionEventType cmdlet.
func (s *Service) NewComplianceRetentionEventType(ctx context.Context, p NewComplianceRetentionEventTypeParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-ComplianceRetentionEventType", p.params())
}

// NewComplianceSearchParams are the parameters of New-ComplianceSearch.
// DefaultParameterSetName: Identity
type NewComplianceSearchParams struct {
	AllowNotFoundExchangeLocationsEnabled bool     `ps:"AllowNotFoundExchangeLocationsEnabled"`
	Case                                  string   `ps:"Case"`
	ContentMatchQuery                     string   `ps:"ContentMatchQuery"`
	Description                           string   `ps:"Description"`
	ExchangeLocation                      []string `ps:"ExchangeLocation"`
	ExchangeLocationExclusion             []string `ps:"ExchangeLocationExclusion"`
	Force                                 bool     `ps:"Force"`
	HoldNames                             []string `ps:"HoldNames"`
	IncludeOrgContent                     bool     `ps:"IncludeOrgContent"`
	IncludeUserAppContent                 bool     `ps:"IncludeUserAppContent"`
	Language                              any      `ps:"Language"`
	LogLevel                              any      `ps:"LogLevel"`
	Name                                  string   `ps:"Name"`
	OneDriveLocation                      []string `ps:"OneDriveLocation"`
	OneDriveLocationExclusion             []string `ps:"OneDriveLocationExclusion"`
	PublicFolderLocation                  []string `ps:"PublicFolderLocation"`
	PublicFolderLocationExclusion         []string `ps:"PublicFolderLocationExclusion"`
	RefinerNames                          []string `ps:"RefinerNames"`
	SearchNames                           []string `ps:"SearchNames"`
	SharePointLocation                    []string `ps:"SharePointLocation"`
	SharePointLocationExclusion           []string `ps:"SharePointLocationExclusion"`
	StatusMailRecipients                  []string `ps:"StatusMailRecipients"`
}

func (p NewComplianceSearchParams) params() map[string]any {
	m := map[string]any{}
	if p.AllowNotFoundExchangeLocationsEnabled {
		m["AllowNotFoundExchangeLocationsEnabled"] = true
	}
	if p.Case != "" {
		m["Case"] = p.Case
	}
	if p.ContentMatchQuery != "" {
		m["ContentMatchQuery"] = p.ContentMatchQuery
	}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if len(p.ExchangeLocation) > 0 {
		m["ExchangeLocation"] = p.ExchangeLocation
	}
	if len(p.ExchangeLocationExclusion) > 0 {
		m["ExchangeLocationExclusion"] = p.ExchangeLocationExclusion
	}
	if p.Force {
		m["Force"] = true
	}
	if len(p.HoldNames) > 0 {
		m["HoldNames"] = p.HoldNames
	}
	if p.IncludeOrgContent {
		m["IncludeOrgContent"] = true
	}
	if p.IncludeUserAppContent {
		m["IncludeUserAppContent"] = true
	}
	if p.Language != nil {
		m["Language"] = p.Language
	}
	if p.LogLevel != nil {
		m["LogLevel"] = p.LogLevel
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if len(p.OneDriveLocation) > 0 {
		m["OneDriveLocation"] = p.OneDriveLocation
	}
	if len(p.OneDriveLocationExclusion) > 0 {
		m["OneDriveLocationExclusion"] = p.OneDriveLocationExclusion
	}
	if len(p.PublicFolderLocation) > 0 {
		m["PublicFolderLocation"] = p.PublicFolderLocation
	}
	if len(p.PublicFolderLocationExclusion) > 0 {
		m["PublicFolderLocationExclusion"] = p.PublicFolderLocationExclusion
	}
	if len(p.RefinerNames) > 0 {
		m["RefinerNames"] = p.RefinerNames
	}
	if len(p.SearchNames) > 0 {
		m["SearchNames"] = p.SearchNames
	}
	if len(p.SharePointLocation) > 0 {
		m["SharePointLocation"] = p.SharePointLocation
	}
	if len(p.SharePointLocationExclusion) > 0 {
		m["SharePointLocationExclusion"] = p.SharePointLocationExclusion
	}
	if len(p.StatusMailRecipients) > 0 {
		m["StatusMailRecipients"] = p.StatusMailRecipients
	}
	return m
}

// NewComplianceSearch runs the New-ComplianceSearch cmdlet.
func (s *Service) NewComplianceSearch(ctx context.Context, p NewComplianceSearchParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-ComplianceSearch", p.params())
}

// NewComplianceSearchActionParams are the parameters of New-ComplianceSearchAction.
// DefaultParameterSetName: Identity
type NewComplianceSearchActionParams struct {
	ActionName                          string   `ps:"ActionName"`
	EnableDedupe                        bool     `ps:"EnableDedupe"`
	FileTypeExclusionsForUnindexedItems []string `ps:"FileTypeExclusionsForUnindexedItems"`
	Force                               bool     `ps:"Force"`
	IncludeCredential                   bool     `ps:"IncludeCredential"`
	IncludeSharePointDocumentVersions   bool     `ps:"IncludeSharePointDocumentVersions"`
	JobOptions                          int      `ps:"JobOptions"`
	MaxUnindexedSize                    int      `ps:"MaxUnindexedSize"`
	ReferenceActionName                 string   `ps:"ReferenceActionName"`
	Region                              string   `ps:"Region"`
	Report                              bool     `ps:"Report"`
	RetentionReport                     bool     `ps:"RetentionReport"`
	RetryOnError                        bool     `ps:"RetryOnError"`
	Scope                               any      `ps:"Scope"` // one of: IndexedItemsOnly, UnindexedItemsOnly, BothIndexedAndUnindexedItems
	ScopeDetails                        []string `ps:"ScopeDetails"`
	SearchName                          []string `ps:"SearchName"`
	SearchNames                         []string `ps:"SearchNames"`
	Version                             string   `ps:"Version"`
}

func (p NewComplianceSearchActionParams) params() map[string]any {
	m := map[string]any{}
	if p.ActionName != "" {
		m["ActionName"] = p.ActionName
	}
	if p.EnableDedupe {
		m["EnableDedupe"] = true
	}
	if len(p.FileTypeExclusionsForUnindexedItems) > 0 {
		m["FileTypeExclusionsForUnindexedItems"] = p.FileTypeExclusionsForUnindexedItems
	}
	if p.Force {
		m["Force"] = true
	}
	if p.IncludeCredential {
		m["IncludeCredential"] = true
	}
	if p.IncludeSharePointDocumentVersions {
		m["IncludeSharePointDocumentVersions"] = true
	}
	if p.JobOptions != 0 {
		m["JobOptions"] = p.JobOptions
	}
	if p.MaxUnindexedSize != 0 {
		m["MaxUnindexedSize"] = p.MaxUnindexedSize
	}
	if p.ReferenceActionName != "" {
		m["ReferenceActionName"] = p.ReferenceActionName
	}
	if p.Region != "" {
		m["Region"] = p.Region
	}
	if p.Report {
		m["Report"] = true
	}
	if p.RetentionReport {
		m["RetentionReport"] = true
	}
	if p.RetryOnError {
		m["RetryOnError"] = true
	}
	if p.Scope != nil {
		m["Scope"] = p.Scope
	}
	if len(p.ScopeDetails) > 0 {
		m["ScopeDetails"] = p.ScopeDetails
	}
	if len(p.SearchName) > 0 {
		m["SearchName"] = p.SearchName
	}
	if len(p.SearchNames) > 0 {
		m["SearchNames"] = p.SearchNames
	}
	if p.Version != "" {
		m["Version"] = p.Version
	}
	return m
}

// NewComplianceSearchAction runs the New-ComplianceSearchAction cmdlet.
func (s *Service) NewComplianceSearchAction(ctx context.Context, p NewComplianceSearchActionParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-ComplianceSearchAction", p.params())
}

// NewComplianceTagParams are the parameters of New-ComplianceTag.
// DefaultParameterSetName: Default
type NewComplianceTagParams struct {
	ArchiveEnabled            bool     `ps:"ArchiveEnabled"`
	ArchiveTriggerBasedOn     string   `ps:"ArchiveTriggerBasedOn"`
	ArchiveTriggerInDays      any      `ps:"ArchiveTriggerInDays"`
	AutoApprovalPeriod        any      `ps:"AutoApprovalPeriod"`
	Comment                   string   `ps:"Comment"`
	ComplianceTagForNextStage string   `ps:"ComplianceTagForNextStage"`
	EventType                 any      `ps:"EventType"`
	FilePlanProperty          string   `ps:"FilePlanProperty"`
	FlowId                    any      `ps:"FlowId"`
	Force                     bool     `ps:"Force"`
	IsRecordLabel             any      `ps:"IsRecordLabel"`
	IsRecordUnlockedAsDefault bool     `ps:"IsRecordUnlockedAsDefault"`
	MultiStageReviewProperty  string   `ps:"MultiStageReviewProperty"`
	Name                      string   `ps:"Name"`
	Notes                     string   `ps:"Notes"`
	Regulatory                bool     `ps:"Regulatory"`
	RetentionAction           string   `ps:"RetentionAction"`
	RetentionDuration         any      `ps:"RetentionDuration"`
	RetentionType             string   `ps:"RetentionType"`
	ReviewerEmail             []string `ps:"ReviewerEmail"`
	WebHookUrl                string   `ps:"WebHookUrl"`
}

func (p NewComplianceTagParams) params() map[string]any {
	m := map[string]any{}
	if p.ArchiveEnabled {
		m["ArchiveEnabled"] = true
	}
	if p.ArchiveTriggerBasedOn != "" {
		m["ArchiveTriggerBasedOn"] = p.ArchiveTriggerBasedOn
	}
	if p.ArchiveTriggerInDays != nil {
		m["ArchiveTriggerInDays"] = p.ArchiveTriggerInDays
	}
	if p.AutoApprovalPeriod != nil {
		m["AutoApprovalPeriod"] = p.AutoApprovalPeriod
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.ComplianceTagForNextStage != "" {
		m["ComplianceTagForNextStage"] = p.ComplianceTagForNextStage
	}
	if p.EventType != nil {
		m["EventType"] = p.EventType
	}
	if p.FilePlanProperty != "" {
		m["FilePlanProperty"] = p.FilePlanProperty
	}
	if p.FlowId != nil {
		m["FlowId"] = p.FlowId
	}
	if p.Force {
		m["Force"] = true
	}
	if p.IsRecordLabel != nil {
		m["IsRecordLabel"] = p.IsRecordLabel
	}
	if p.IsRecordUnlockedAsDefault {
		m["IsRecordUnlockedAsDefault"] = true
	}
	if p.MultiStageReviewProperty != "" {
		m["MultiStageReviewProperty"] = p.MultiStageReviewProperty
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Notes != "" {
		m["Notes"] = p.Notes
	}
	if p.Regulatory {
		m["Regulatory"] = true
	}
	if p.RetentionAction != "" {
		m["RetentionAction"] = p.RetentionAction
	}
	if p.RetentionDuration != nil {
		m["RetentionDuration"] = p.RetentionDuration
	}
	if p.RetentionType != "" {
		m["RetentionType"] = p.RetentionType
	}
	if len(p.ReviewerEmail) > 0 {
		m["ReviewerEmail"] = p.ReviewerEmail
	}
	if p.WebHookUrl != "" {
		m["WebHookUrl"] = p.WebHookUrl
	}
	return m
}

// NewComplianceTag runs the New-ComplianceTag cmdlet.
func (s *Service) NewComplianceTag(ctx context.Context, p NewComplianceTagParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-ComplianceTag", p.params())
}

// NewCustomDlpEmailTemplateParams are the parameters of New-CustomDlpEmailTemplate.
// DefaultParameterSetName: Identity
type NewCustomDlpEmailTemplateParams struct {
	Bcc               []string `ps:"Bcc"`
	Body              string   `ps:"Body"`
	Cc                []string `ps:"Cc"`
	Description       string   `ps:"Description"`
	From              any      `ps:"From"`
	Importance        any      `ps:"Importance"`
	Name              string   `ps:"Name"`
	NotificationTitle string   `ps:"NotificationTitle"`
	Subject           string   `ps:"Subject"`
	To                []string `ps:"To"`
}

func (p NewCustomDlpEmailTemplateParams) params() map[string]any {
	m := map[string]any{}
	if len(p.Bcc) > 0 {
		m["Bcc"] = p.Bcc
	}
	if p.Body != "" {
		m["Body"] = p.Body
	}
	if len(p.Cc) > 0 {
		m["Cc"] = p.Cc
	}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if p.From != nil {
		m["From"] = p.From
	}
	if p.Importance != nil {
		m["Importance"] = p.Importance
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.NotificationTitle != "" {
		m["NotificationTitle"] = p.NotificationTitle
	}
	if p.Subject != "" {
		m["Subject"] = p.Subject
	}
	if len(p.To) > 0 {
		m["To"] = p.To
	}
	return m
}

// NewCustomDlpEmailTemplate runs the New-CustomDlpEmailTemplate cmdlet.
func (s *Service) NewCustomDlpEmailTemplate(ctx context.Context, p NewCustomDlpEmailTemplateParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-CustomDlpEmailTemplate", p.params())
}

// NewDeviceConditionalAccessPolicyParams are the parameters of New-DeviceConditionalAccessPolicy.
type NewDeviceConditionalAccessPolicyParams struct {
	Comment string `ps:"Comment"`
	Enabled bool   `ps:"Enabled"`
	Force   bool   `ps:"Force"`
	Name    string `ps:"Name"`
}

func (p NewDeviceConditionalAccessPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	return m
}

// NewDeviceConditionalAccessPolicy runs the New-DeviceConditionalAccessPolicy cmdlet.
func (s *Service) NewDeviceConditionalAccessPolicy(ctx context.Context, p NewDeviceConditionalAccessPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-DeviceConditionalAccessPolicy", p.params())
}

// NewDeviceConditionalAccessRuleParams are the parameters of New-DeviceConditionalAccessRule.
type NewDeviceConditionalAccessRuleParams struct {
	AccountName                   string `ps:"AccountName"`
	AccountUserName               string `ps:"AccountUserName"`
	AllowAppStore                 any    `ps:"AllowAppStore"`
	AllowAssistantWhileLocked     any    `ps:"AllowAssistantWhileLocked"`
	AllowConvenienceLogon         any    `ps:"AllowConvenienceLogon"`
	AllowDiagnosticSubmission     any    `ps:"AllowDiagnosticSubmission"`
	AllowiCloudBackup             any    `ps:"AllowiCloudBackup"`
	AllowiCloudDocSync            any    `ps:"AllowiCloudDocSync"`
	AllowiCloudPhotoSync          any    `ps:"AllowiCloudPhotoSync"`
	AllowJailbroken               any    `ps:"AllowJailbroken"`
	AllowPassbookWhileLocked      any    `ps:"AllowPassbookWhileLocked"`
	AllowScreenshot               any    `ps:"AllowScreenshot"`
	AllowSimplePassword           any    `ps:"AllowSimplePassword"`
	AllowVideoConferencing        any    `ps:"AllowVideoConferencing"`
	AllowVoiceAssistant           any    `ps:"AllowVoiceAssistant"`
	AllowVoiceDialing             any    `ps:"AllowVoiceDialing"`
	AntiVirusSignatureStatus      any    `ps:"AntiVirusSignatureStatus"`
	AntiVirusStatus               any    `ps:"AntiVirusStatus"`
	AppsRating                    any    `ps:"AppsRating"`
	AutoUpdateStatus              any    `ps:"AutoUpdateStatus"`
	BluetoothEnabled              any    `ps:"BluetoothEnabled"`
	CameraEnabled                 any    `ps:"CameraEnabled"`
	DomainController              any    `ps:"DomainController"`
	EmailAddress                  string `ps:"EmailAddress"`
	EnableRemovableStorage        any    `ps:"EnableRemovableStorage"`
	ExchangeActiveSyncHost        string `ps:"ExchangeActiveSyncHost"`
	FirewallStatus                any    `ps:"FirewallStatus"`
	ForceAppStorePassword         any    `ps:"ForceAppStorePassword"`
	ForceEncryptedBackup          any    `ps:"ForceEncryptedBackup"`
	MaxPasswordAttemptsBeforeWipe any    `ps:"MaxPasswordAttemptsBeforeWipe"`
	MaxPasswordGracePeriod        any    `ps:"MaxPasswordGracePeriod"`
	MoviesRating                  any    `ps:"MoviesRating"`
	PasswordComplexity            any    `ps:"PasswordComplexity"`
	PasswordExpirationDays        any    `ps:"PasswordExpirationDays"`
	PasswordHistoryCount          any    `ps:"PasswordHistoryCount"`
	PasswordMinComplexChars       any    `ps:"PasswordMinComplexChars"`
	PasswordMinimumLength         any    `ps:"PasswordMinimumLength"`
	PasswordQuality               any    `ps:"PasswordQuality"`
	PasswordRequired              any    `ps:"PasswordRequired"`
	PasswordTimeout               any    `ps:"PasswordTimeout"`
	PhoneMemoryEncrypted          any    `ps:"PhoneMemoryEncrypted"`
	Policy                        any    `ps:"Policy"`
	RegionRatings                 any    `ps:"RegionRatings"`
	RequireEmailProfile           any    `ps:"RequireEmailProfile"`
	SmartScreenEnabled            any    `ps:"SmartScreenEnabled"`
	SystemSecurityTLS             any    `ps:"SystemSecurityTLS"`
	TargetGroups                  any    `ps:"TargetGroups"`
	TVShowsRating                 any    `ps:"TVShowsRating"`
	UserAccountControlStatus      any    `ps:"UserAccountControlStatus"`
	WLANEnabled                   any    `ps:"WLANEnabled"`
	WorkFoldersSyncUrl            string `ps:"WorkFoldersSyncUrl"`
}

func (p NewDeviceConditionalAccessRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.AccountName != "" {
		m["AccountName"] = p.AccountName
	}
	if p.AccountUserName != "" {
		m["AccountUserName"] = p.AccountUserName
	}
	if p.AllowAppStore != nil {
		m["AllowAppStore"] = p.AllowAppStore
	}
	if p.AllowAssistantWhileLocked != nil {
		m["AllowAssistantWhileLocked"] = p.AllowAssistantWhileLocked
	}
	if p.AllowConvenienceLogon != nil {
		m["AllowConvenienceLogon"] = p.AllowConvenienceLogon
	}
	if p.AllowDiagnosticSubmission != nil {
		m["AllowDiagnosticSubmission"] = p.AllowDiagnosticSubmission
	}
	if p.AllowiCloudBackup != nil {
		m["AllowiCloudBackup"] = p.AllowiCloudBackup
	}
	if p.AllowiCloudDocSync != nil {
		m["AllowiCloudDocSync"] = p.AllowiCloudDocSync
	}
	if p.AllowiCloudPhotoSync != nil {
		m["AllowiCloudPhotoSync"] = p.AllowiCloudPhotoSync
	}
	if p.AllowJailbroken != nil {
		m["AllowJailbroken"] = p.AllowJailbroken
	}
	if p.AllowPassbookWhileLocked != nil {
		m["AllowPassbookWhileLocked"] = p.AllowPassbookWhileLocked
	}
	if p.AllowScreenshot != nil {
		m["AllowScreenshot"] = p.AllowScreenshot
	}
	if p.AllowSimplePassword != nil {
		m["AllowSimplePassword"] = p.AllowSimplePassword
	}
	if p.AllowVideoConferencing != nil {
		m["AllowVideoConferencing"] = p.AllowVideoConferencing
	}
	if p.AllowVoiceAssistant != nil {
		m["AllowVoiceAssistant"] = p.AllowVoiceAssistant
	}
	if p.AllowVoiceDialing != nil {
		m["AllowVoiceDialing"] = p.AllowVoiceDialing
	}
	if p.AntiVirusSignatureStatus != nil {
		m["AntiVirusSignatureStatus"] = p.AntiVirusSignatureStatus
	}
	if p.AntiVirusStatus != nil {
		m["AntiVirusStatus"] = p.AntiVirusStatus
	}
	if p.AppsRating != nil {
		m["AppsRating"] = p.AppsRating
	}
	if p.AutoUpdateStatus != nil {
		m["AutoUpdateStatus"] = p.AutoUpdateStatus
	}
	if p.BluetoothEnabled != nil {
		m["BluetoothEnabled"] = p.BluetoothEnabled
	}
	if p.CameraEnabled != nil {
		m["CameraEnabled"] = p.CameraEnabled
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.EmailAddress != "" {
		m["EmailAddress"] = p.EmailAddress
	}
	if p.EnableRemovableStorage != nil {
		m["EnableRemovableStorage"] = p.EnableRemovableStorage
	}
	if p.ExchangeActiveSyncHost != "" {
		m["ExchangeActiveSyncHost"] = p.ExchangeActiveSyncHost
	}
	if p.FirewallStatus != nil {
		m["FirewallStatus"] = p.FirewallStatus
	}
	if p.ForceAppStorePassword != nil {
		m["ForceAppStorePassword"] = p.ForceAppStorePassword
	}
	if p.ForceEncryptedBackup != nil {
		m["ForceEncryptedBackup"] = p.ForceEncryptedBackup
	}
	if p.MaxPasswordAttemptsBeforeWipe != nil {
		m["MaxPasswordAttemptsBeforeWipe"] = p.MaxPasswordAttemptsBeforeWipe
	}
	if p.MaxPasswordGracePeriod != nil {
		m["MaxPasswordGracePeriod"] = p.MaxPasswordGracePeriod
	}
	if p.MoviesRating != nil {
		m["MoviesRating"] = p.MoviesRating
	}
	if p.PasswordComplexity != nil {
		m["PasswordComplexity"] = p.PasswordComplexity
	}
	if p.PasswordExpirationDays != nil {
		m["PasswordExpirationDays"] = p.PasswordExpirationDays
	}
	if p.PasswordHistoryCount != nil {
		m["PasswordHistoryCount"] = p.PasswordHistoryCount
	}
	if p.PasswordMinComplexChars != nil {
		m["PasswordMinComplexChars"] = p.PasswordMinComplexChars
	}
	if p.PasswordMinimumLength != nil {
		m["PasswordMinimumLength"] = p.PasswordMinimumLength
	}
	if p.PasswordQuality != nil {
		m["PasswordQuality"] = p.PasswordQuality
	}
	if p.PasswordRequired != nil {
		m["PasswordRequired"] = p.PasswordRequired
	}
	if p.PasswordTimeout != nil {
		m["PasswordTimeout"] = p.PasswordTimeout
	}
	if p.PhoneMemoryEncrypted != nil {
		m["PhoneMemoryEncrypted"] = p.PhoneMemoryEncrypted
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	if p.RegionRatings != nil {
		m["RegionRatings"] = p.RegionRatings
	}
	if p.RequireEmailProfile != nil {
		m["RequireEmailProfile"] = p.RequireEmailProfile
	}
	if p.SmartScreenEnabled != nil {
		m["SmartScreenEnabled"] = p.SmartScreenEnabled
	}
	if p.SystemSecurityTLS != nil {
		m["SystemSecurityTLS"] = p.SystemSecurityTLS
	}
	if p.TargetGroups != nil {
		m["TargetGroups"] = p.TargetGroups
	}
	if p.TVShowsRating != nil {
		m["TVShowsRating"] = p.TVShowsRating
	}
	if p.UserAccountControlStatus != nil {
		m["UserAccountControlStatus"] = p.UserAccountControlStatus
	}
	if p.WLANEnabled != nil {
		m["WLANEnabled"] = p.WLANEnabled
	}
	if p.WorkFoldersSyncUrl != "" {
		m["WorkFoldersSyncUrl"] = p.WorkFoldersSyncUrl
	}
	return m
}

// NewDeviceConditionalAccessRule runs the New-DeviceConditionalAccessRule cmdlet.
func (s *Service) NewDeviceConditionalAccessRule(ctx context.Context, p NewDeviceConditionalAccessRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-DeviceConditionalAccessRule", p.params())
}

// NewDeviceConfigurationPolicyParams are the parameters of New-DeviceConfigurationPolicy.
type NewDeviceConfigurationPolicyParams struct {
	Comment string `ps:"Comment"`
	Enabled bool   `ps:"Enabled"`
	Force   bool   `ps:"Force"`
	Name    string `ps:"Name"`
}

func (p NewDeviceConfigurationPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	return m
}

// NewDeviceConfigurationPolicy runs the New-DeviceConfigurationPolicy cmdlet.
func (s *Service) NewDeviceConfigurationPolicy(ctx context.Context, p NewDeviceConfigurationPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-DeviceConfigurationPolicy", p.params())
}

// NewDeviceConfigurationRuleParams are the parameters of New-DeviceConfigurationRule.
type NewDeviceConfigurationRuleParams struct {
	AccountName                   string `ps:"AccountName"`
	AccountUserName               string `ps:"AccountUserName"`
	AllowAppStore                 any    `ps:"AllowAppStore"`
	AllowAssistantWhileLocked     any    `ps:"AllowAssistantWhileLocked"`
	AllowConvenienceLogon         any    `ps:"AllowConvenienceLogon"`
	AllowDiagnosticSubmission     any    `ps:"AllowDiagnosticSubmission"`
	AllowiCloudBackup             any    `ps:"AllowiCloudBackup"`
	AllowiCloudDocSync            any    `ps:"AllowiCloudDocSync"`
	AllowiCloudPhotoSync          any    `ps:"AllowiCloudPhotoSync"`
	AllowPassbookWhileLocked      any    `ps:"AllowPassbookWhileLocked"`
	AllowScreenshot               any    `ps:"AllowScreenshot"`
	AllowSimplePassword           any    `ps:"AllowSimplePassword"`
	AllowVideoConferencing        any    `ps:"AllowVideoConferencing"`
	AllowVoiceAssistant           any    `ps:"AllowVoiceAssistant"`
	AllowVoiceDialing             any    `ps:"AllowVoiceDialing"`
	AntiVirusSignatureStatus      any    `ps:"AntiVirusSignatureStatus"`
	AntiVirusStatus               any    `ps:"AntiVirusStatus"`
	AppsRating                    any    `ps:"AppsRating"`
	AutoUpdateStatus              any    `ps:"AutoUpdateStatus"`
	BluetoothEnabled              any    `ps:"BluetoothEnabled"`
	CameraEnabled                 any    `ps:"CameraEnabled"`
	DomainController              any    `ps:"DomainController"`
	EmailAddress                  string `ps:"EmailAddress"`
	EnableRemovableStorage        any    `ps:"EnableRemovableStorage"`
	ExchangeActiveSyncHost        string `ps:"ExchangeActiveSyncHost"`
	FirewallStatus                any    `ps:"FirewallStatus"`
	ForceAppStorePassword         any    `ps:"ForceAppStorePassword"`
	ForceEncryptedBackup          any    `ps:"ForceEncryptedBackup"`
	MaxPasswordAttemptsBeforeWipe any    `ps:"MaxPasswordAttemptsBeforeWipe"`
	MaxPasswordGracePeriod        any    `ps:"MaxPasswordGracePeriod"`
	MoviesRating                  any    `ps:"MoviesRating"`
	PasswordComplexity            any    `ps:"PasswordComplexity"`
	PasswordExpirationDays        any    `ps:"PasswordExpirationDays"`
	PasswordHistoryCount          any    `ps:"PasswordHistoryCount"`
	PasswordMinComplexChars       any    `ps:"PasswordMinComplexChars"`
	PasswordMinimumLength         any    `ps:"PasswordMinimumLength"`
	PasswordQuality               any    `ps:"PasswordQuality"`
	PasswordRequired              any    `ps:"PasswordRequired"`
	PasswordTimeout               any    `ps:"PasswordTimeout"`
	PhoneMemoryEncrypted          any    `ps:"PhoneMemoryEncrypted"`
	Policy                        any    `ps:"Policy"`
	RegionRatings                 any    `ps:"RegionRatings"`
	RequireEmailProfile           any    `ps:"RequireEmailProfile"`
	SmartScreenEnabled            any    `ps:"SmartScreenEnabled"`
	SystemSecurityTLS             any    `ps:"SystemSecurityTLS"`
	TargetGroups                  any    `ps:"TargetGroups"`
	TVShowsRating                 any    `ps:"TVShowsRating"`
	UserAccountControlStatus      any    `ps:"UserAccountControlStatus"`
	WLANEnabled                   any    `ps:"WLANEnabled"`
	WorkFoldersSyncUrl            string `ps:"WorkFoldersSyncUrl"`
}

func (p NewDeviceConfigurationRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.AccountName != "" {
		m["AccountName"] = p.AccountName
	}
	if p.AccountUserName != "" {
		m["AccountUserName"] = p.AccountUserName
	}
	if p.AllowAppStore != nil {
		m["AllowAppStore"] = p.AllowAppStore
	}
	if p.AllowAssistantWhileLocked != nil {
		m["AllowAssistantWhileLocked"] = p.AllowAssistantWhileLocked
	}
	if p.AllowConvenienceLogon != nil {
		m["AllowConvenienceLogon"] = p.AllowConvenienceLogon
	}
	if p.AllowDiagnosticSubmission != nil {
		m["AllowDiagnosticSubmission"] = p.AllowDiagnosticSubmission
	}
	if p.AllowiCloudBackup != nil {
		m["AllowiCloudBackup"] = p.AllowiCloudBackup
	}
	if p.AllowiCloudDocSync != nil {
		m["AllowiCloudDocSync"] = p.AllowiCloudDocSync
	}
	if p.AllowiCloudPhotoSync != nil {
		m["AllowiCloudPhotoSync"] = p.AllowiCloudPhotoSync
	}
	if p.AllowPassbookWhileLocked != nil {
		m["AllowPassbookWhileLocked"] = p.AllowPassbookWhileLocked
	}
	if p.AllowScreenshot != nil {
		m["AllowScreenshot"] = p.AllowScreenshot
	}
	if p.AllowSimplePassword != nil {
		m["AllowSimplePassword"] = p.AllowSimplePassword
	}
	if p.AllowVideoConferencing != nil {
		m["AllowVideoConferencing"] = p.AllowVideoConferencing
	}
	if p.AllowVoiceAssistant != nil {
		m["AllowVoiceAssistant"] = p.AllowVoiceAssistant
	}
	if p.AllowVoiceDialing != nil {
		m["AllowVoiceDialing"] = p.AllowVoiceDialing
	}
	if p.AntiVirusSignatureStatus != nil {
		m["AntiVirusSignatureStatus"] = p.AntiVirusSignatureStatus
	}
	if p.AntiVirusStatus != nil {
		m["AntiVirusStatus"] = p.AntiVirusStatus
	}
	if p.AppsRating != nil {
		m["AppsRating"] = p.AppsRating
	}
	if p.AutoUpdateStatus != nil {
		m["AutoUpdateStatus"] = p.AutoUpdateStatus
	}
	if p.BluetoothEnabled != nil {
		m["BluetoothEnabled"] = p.BluetoothEnabled
	}
	if p.CameraEnabled != nil {
		m["CameraEnabled"] = p.CameraEnabled
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.EmailAddress != "" {
		m["EmailAddress"] = p.EmailAddress
	}
	if p.EnableRemovableStorage != nil {
		m["EnableRemovableStorage"] = p.EnableRemovableStorage
	}
	if p.ExchangeActiveSyncHost != "" {
		m["ExchangeActiveSyncHost"] = p.ExchangeActiveSyncHost
	}
	if p.FirewallStatus != nil {
		m["FirewallStatus"] = p.FirewallStatus
	}
	if p.ForceAppStorePassword != nil {
		m["ForceAppStorePassword"] = p.ForceAppStorePassword
	}
	if p.ForceEncryptedBackup != nil {
		m["ForceEncryptedBackup"] = p.ForceEncryptedBackup
	}
	if p.MaxPasswordAttemptsBeforeWipe != nil {
		m["MaxPasswordAttemptsBeforeWipe"] = p.MaxPasswordAttemptsBeforeWipe
	}
	if p.MaxPasswordGracePeriod != nil {
		m["MaxPasswordGracePeriod"] = p.MaxPasswordGracePeriod
	}
	if p.MoviesRating != nil {
		m["MoviesRating"] = p.MoviesRating
	}
	if p.PasswordComplexity != nil {
		m["PasswordComplexity"] = p.PasswordComplexity
	}
	if p.PasswordExpirationDays != nil {
		m["PasswordExpirationDays"] = p.PasswordExpirationDays
	}
	if p.PasswordHistoryCount != nil {
		m["PasswordHistoryCount"] = p.PasswordHistoryCount
	}
	if p.PasswordMinComplexChars != nil {
		m["PasswordMinComplexChars"] = p.PasswordMinComplexChars
	}
	if p.PasswordMinimumLength != nil {
		m["PasswordMinimumLength"] = p.PasswordMinimumLength
	}
	if p.PasswordQuality != nil {
		m["PasswordQuality"] = p.PasswordQuality
	}
	if p.PasswordRequired != nil {
		m["PasswordRequired"] = p.PasswordRequired
	}
	if p.PasswordTimeout != nil {
		m["PasswordTimeout"] = p.PasswordTimeout
	}
	if p.PhoneMemoryEncrypted != nil {
		m["PhoneMemoryEncrypted"] = p.PhoneMemoryEncrypted
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	if p.RegionRatings != nil {
		m["RegionRatings"] = p.RegionRatings
	}
	if p.RequireEmailProfile != nil {
		m["RequireEmailProfile"] = p.RequireEmailProfile
	}
	if p.SmartScreenEnabled != nil {
		m["SmartScreenEnabled"] = p.SmartScreenEnabled
	}
	if p.SystemSecurityTLS != nil {
		m["SystemSecurityTLS"] = p.SystemSecurityTLS
	}
	if p.TargetGroups != nil {
		m["TargetGroups"] = p.TargetGroups
	}
	if p.TVShowsRating != nil {
		m["TVShowsRating"] = p.TVShowsRating
	}
	if p.UserAccountControlStatus != nil {
		m["UserAccountControlStatus"] = p.UserAccountControlStatus
	}
	if p.WLANEnabled != nil {
		m["WLANEnabled"] = p.WLANEnabled
	}
	if p.WorkFoldersSyncUrl != "" {
		m["WorkFoldersSyncUrl"] = p.WorkFoldersSyncUrl
	}
	return m
}

// NewDeviceConfigurationRule runs the New-DeviceConfigurationRule cmdlet.
func (s *Service) NewDeviceConfigurationRule(ctx context.Context, p NewDeviceConfigurationRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-DeviceConfigurationRule", p.params())
}

// NewDeviceTenantPolicyParams are the parameters of New-DeviceTenantPolicy.
type NewDeviceTenantPolicyParams struct {
	Comment string `ps:"Comment"`
	Enabled bool   `ps:"Enabled"`
	Force   bool   `ps:"Force"`
}

func (p NewDeviceTenantPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.Force {
		m["Force"] = true
	}
	return m
}

// NewDeviceTenantPolicy runs the New-DeviceTenantPolicy cmdlet.
func (s *Service) NewDeviceTenantPolicy(ctx context.Context, p NewDeviceTenantPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-DeviceTenantPolicy", p.params())
}

// NewDeviceTenantRuleParams are the parameters of New-DeviceTenantRule.
type NewDeviceTenantRuleParams struct {
	ApplyPolicyTo           any `ps:"ApplyPolicyTo"`
	BlockUnsupportedDevices any `ps:"BlockUnsupportedDevices"`
	DomainController        any `ps:"DomainController"`
	ExclusionList           any `ps:"ExclusionList"`
}

func (p NewDeviceTenantRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.ApplyPolicyTo != nil {
		m["ApplyPolicyTo"] = p.ApplyPolicyTo
	}
	if p.BlockUnsupportedDevices != nil {
		m["BlockUnsupportedDevices"] = p.BlockUnsupportedDevices
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.ExclusionList != nil {
		m["ExclusionList"] = p.ExclusionList
	}
	return m
}

// NewDeviceTenantRule runs the New-DeviceTenantRule cmdlet.
func (s *Service) NewDeviceTenantRule(ctx context.Context, p NewDeviceTenantRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-DeviceTenantRule", p.params())
}

// NewDlpAIPoweredSensitiveInformationTypeParams are the parameters of New-DlpAIPoweredSensitiveInformationType.
type NewDlpAIPoweredSensitiveInformationTypeParams struct {
	BaseId any `ps:"BaseId"`
}

func (p NewDlpAIPoweredSensitiveInformationTypeParams) params() map[string]any {
	m := map[string]any{}
	if p.BaseId != nil {
		m["BaseId"] = p.BaseId
	}
	return m
}

// NewDlpAIPoweredSensitiveInformationType runs the New-DlpAIPoweredSensitiveInformationType cmdlet.
func (s *Service) NewDlpAIPoweredSensitiveInformationType(ctx context.Context, p NewDlpAIPoweredSensitiveInformationTypeParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-DlpAIPoweredSensitiveInformationType", p.params())
}

// NewDlpAlertTuningRuleParams are the parameters of New-DlpAlertTuningRule.
// DefaultParameterSetName: Identity
type NewDlpAlertTuningRuleParams struct {
	Actions        string `ps:"Actions"`
	ApplicableTags string `ps:"ApplicableTags"`
	Conditions     string `ps:"Conditions"`
	Description    string `ps:"Description"`
	Enabled        string `ps:"Enabled"`
	Name           string `ps:"Name"`
}

func (p NewDlpAlertTuningRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Actions != "" {
		m["Actions"] = p.Actions
	}
	if p.ApplicableTags != "" {
		m["ApplicableTags"] = p.ApplicableTags
	}
	if p.Conditions != "" {
		m["Conditions"] = p.Conditions
	}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if p.Enabled != "" {
		m["Enabled"] = p.Enabled
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	return m
}

// NewDlpAlertTuningRule runs the New-DlpAlertTuningRule cmdlet.
func (s *Service) NewDlpAlertTuningRule(ctx context.Context, p NewDlpAlertTuningRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-DlpAlertTuningRule", p.params())
}

// NewDlpCompliancePolicyParams are the parameters of New-DlpCompliancePolicy.
type NewDlpCompliancePolicyParams struct {
	Comment                               string   `ps:"Comment"`
	DisplayName                           string   `ps:"DisplayName"`
	EndpointDlpAdaptiveScopes             any      `ps:"EndpointDlpAdaptiveScopes"`
	EndpointDlpAdaptiveScopesException    any      `ps:"EndpointDlpAdaptiveScopesException"`
	EndpointDlpExtendedLocations          string   `ps:"EndpointDlpExtendedLocations"`
	EndpointDlpLocation                   any      `ps:"EndpointDlpLocation"`
	EndpointDlpLocationException          any      `ps:"EndpointDlpLocationException"`
	EnforcementPlanes                     any      `ps:"EnforcementPlanes"`
	ExceptIfOneDriveSharedBy              []string `ps:"ExceptIfOneDriveSharedBy"`
	ExceptIfOneDriveSharedByMemberOf      []string `ps:"ExceptIfOneDriveSharedByMemberOf"`
	ExchangeAdaptiveScopes                any      `ps:"ExchangeAdaptiveScopes"`
	ExchangeAdaptiveScopesException       any      `ps:"ExchangeAdaptiveScopesException"`
	ExchangeLocation                      any      `ps:"ExchangeLocation"`
	ExchangeOnPremisesLocation            any      `ps:"ExchangeOnPremisesLocation"`
	ExchangeSender                        []string `ps:"ExchangeSender"`
	ExchangeSenderException               []string `ps:"ExchangeSenderException"`
	ExchangeSenderMemberOf                []string `ps:"ExchangeSenderMemberOf"`
	ExchangeSenderMemberOfException       []string `ps:"ExchangeSenderMemberOfException"`
	Force                                 bool     `ps:"Force"`
	IsFromSmartInsights                   any      `ps:"IsFromSmartInsights"`
	Locations                             string   `ps:"Locations"`
	Mode                                  any      `ps:"Mode"`
	Name                                  string   `ps:"Name"`
	OneDriveAdaptiveScopes                any      `ps:"OneDriveAdaptiveScopes"`
	OneDriveAdaptiveScopesException       any      `ps:"OneDriveAdaptiveScopesException"`
	OneDriveLocation                      any      `ps:"OneDriveLocation"`
	OneDriveLocationException             any      `ps:"OneDriveLocationException"`
	OneDriveSharedBy                      []string `ps:"OneDriveSharedBy"`
	OneDriveSharedByMemberOf              []string `ps:"OneDriveSharedByMemberOf"`
	OnPremisesScannerDlpLocation          any      `ps:"OnPremisesScannerDlpLocation"`
	OnPremisesScannerDlpLocationException any      `ps:"OnPremisesScannerDlpLocationException"`
	PolicyRBACScopes                      any      `ps:"PolicyRBACScopes"`
	PolicyTemplateInfo                    any      `ps:"PolicyTemplateInfo"`
	PowerBIDlpLocation                    any      `ps:"PowerBIDlpLocation"`
	PowerBIDlpLocationException           any      `ps:"PowerBIDlpLocationException"`
	Priority                              any      `ps:"Priority"`
	SharePointAdaptiveScopes              any      `ps:"SharePointAdaptiveScopes"`
	SharePointAdaptiveScopesException     any      `ps:"SharePointAdaptiveScopesException"`
	SharePointLocation                    any      `ps:"SharePointLocation"`
	SharePointLocationException           any      `ps:"SharePointLocationException"`
	SharePointOnPremisesLocationException any      `ps:"SharePointOnPremisesLocationException"`
	SharePointServerLocation              any      `ps:"SharePointServerLocation"`
	TeamsAdaptiveScopes                   any      `ps:"TeamsAdaptiveScopes"`
	TeamsAdaptiveScopesException          any      `ps:"TeamsAdaptiveScopesException"`
	TeamsLocation                         any      `ps:"TeamsLocation"`
	TeamsLocationException                any      `ps:"TeamsLocationException"`
	ThirdPartyAppDlpLocation              any      `ps:"ThirdPartyAppDlpLocation"`
	ThirdPartyAppDlpLocationException     any      `ps:"ThirdPartyAppDlpLocationException"`
	ValidatePolicy                        bool     `ps:"ValidatePolicy"`
}

func (p NewDlpCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.EndpointDlpAdaptiveScopes != nil {
		m["EndpointDlpAdaptiveScopes"] = p.EndpointDlpAdaptiveScopes
	}
	if p.EndpointDlpAdaptiveScopesException != nil {
		m["EndpointDlpAdaptiveScopesException"] = p.EndpointDlpAdaptiveScopesException
	}
	if p.EndpointDlpExtendedLocations != "" {
		m["EndpointDlpExtendedLocations"] = p.EndpointDlpExtendedLocations
	}
	if p.EndpointDlpLocation != nil {
		m["EndpointDlpLocation"] = p.EndpointDlpLocation
	}
	if p.EndpointDlpLocationException != nil {
		m["EndpointDlpLocationException"] = p.EndpointDlpLocationException
	}
	if p.EnforcementPlanes != nil {
		m["EnforcementPlanes"] = p.EnforcementPlanes
	}
	if len(p.ExceptIfOneDriveSharedBy) > 0 {
		m["ExceptIfOneDriveSharedBy"] = p.ExceptIfOneDriveSharedBy
	}
	if len(p.ExceptIfOneDriveSharedByMemberOf) > 0 {
		m["ExceptIfOneDriveSharedByMemberOf"] = p.ExceptIfOneDriveSharedByMemberOf
	}
	if p.ExchangeAdaptiveScopes != nil {
		m["ExchangeAdaptiveScopes"] = p.ExchangeAdaptiveScopes
	}
	if p.ExchangeAdaptiveScopesException != nil {
		m["ExchangeAdaptiveScopesException"] = p.ExchangeAdaptiveScopesException
	}
	if p.ExchangeLocation != nil {
		m["ExchangeLocation"] = p.ExchangeLocation
	}
	if p.ExchangeOnPremisesLocation != nil {
		m["ExchangeOnPremisesLocation"] = p.ExchangeOnPremisesLocation
	}
	if len(p.ExchangeSender) > 0 {
		m["ExchangeSender"] = p.ExchangeSender
	}
	if len(p.ExchangeSenderException) > 0 {
		m["ExchangeSenderException"] = p.ExchangeSenderException
	}
	if len(p.ExchangeSenderMemberOf) > 0 {
		m["ExchangeSenderMemberOf"] = p.ExchangeSenderMemberOf
	}
	if len(p.ExchangeSenderMemberOfException) > 0 {
		m["ExchangeSenderMemberOfException"] = p.ExchangeSenderMemberOfException
	}
	if p.Force {
		m["Force"] = true
	}
	if p.IsFromSmartInsights != nil {
		m["IsFromSmartInsights"] = p.IsFromSmartInsights
	}
	if p.Locations != "" {
		m["Locations"] = p.Locations
	}
	if p.Mode != nil {
		m["Mode"] = p.Mode
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.OneDriveAdaptiveScopes != nil {
		m["OneDriveAdaptiveScopes"] = p.OneDriveAdaptiveScopes
	}
	if p.OneDriveAdaptiveScopesException != nil {
		m["OneDriveAdaptiveScopesException"] = p.OneDriveAdaptiveScopesException
	}
	if p.OneDriveLocation != nil {
		m["OneDriveLocation"] = p.OneDriveLocation
	}
	if p.OneDriveLocationException != nil {
		m["OneDriveLocationException"] = p.OneDriveLocationException
	}
	if len(p.OneDriveSharedBy) > 0 {
		m["OneDriveSharedBy"] = p.OneDriveSharedBy
	}
	if len(p.OneDriveSharedByMemberOf) > 0 {
		m["OneDriveSharedByMemberOf"] = p.OneDriveSharedByMemberOf
	}
	if p.OnPremisesScannerDlpLocation != nil {
		m["OnPremisesScannerDlpLocation"] = p.OnPremisesScannerDlpLocation
	}
	if p.OnPremisesScannerDlpLocationException != nil {
		m["OnPremisesScannerDlpLocationException"] = p.OnPremisesScannerDlpLocationException
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.PolicyTemplateInfo != nil {
		m["PolicyTemplateInfo"] = p.PolicyTemplateInfo
	}
	if p.PowerBIDlpLocation != nil {
		m["PowerBIDlpLocation"] = p.PowerBIDlpLocation
	}
	if p.PowerBIDlpLocationException != nil {
		m["PowerBIDlpLocationException"] = p.PowerBIDlpLocationException
	}
	if p.Priority != nil {
		m["Priority"] = p.Priority
	}
	if p.SharePointAdaptiveScopes != nil {
		m["SharePointAdaptiveScopes"] = p.SharePointAdaptiveScopes
	}
	if p.SharePointAdaptiveScopesException != nil {
		m["SharePointAdaptiveScopesException"] = p.SharePointAdaptiveScopesException
	}
	if p.SharePointLocation != nil {
		m["SharePointLocation"] = p.SharePointLocation
	}
	if p.SharePointLocationException != nil {
		m["SharePointLocationException"] = p.SharePointLocationException
	}
	if p.SharePointOnPremisesLocationException != nil {
		m["SharePointOnPremisesLocationException"] = p.SharePointOnPremisesLocationException
	}
	if p.SharePointServerLocation != nil {
		m["SharePointServerLocation"] = p.SharePointServerLocation
	}
	if p.TeamsAdaptiveScopes != nil {
		m["TeamsAdaptiveScopes"] = p.TeamsAdaptiveScopes
	}
	if p.TeamsAdaptiveScopesException != nil {
		m["TeamsAdaptiveScopesException"] = p.TeamsAdaptiveScopesException
	}
	if p.TeamsLocation != nil {
		m["TeamsLocation"] = p.TeamsLocation
	}
	if p.TeamsLocationException != nil {
		m["TeamsLocationException"] = p.TeamsLocationException
	}
	if p.ThirdPartyAppDlpLocation != nil {
		m["ThirdPartyAppDlpLocation"] = p.ThirdPartyAppDlpLocation
	}
	if p.ThirdPartyAppDlpLocationException != nil {
		m["ThirdPartyAppDlpLocationException"] = p.ThirdPartyAppDlpLocationException
	}
	if p.ValidatePolicy {
		m["ValidatePolicy"] = true
	}
	return m
}

// NewDlpCompliancePolicy runs the New-DlpCompliancePolicy cmdlet.
func (s *Service) NewDlpCompliancePolicy(ctx context.Context, p NewDlpCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-DlpCompliancePolicy", p.params())
}

// NewDlpComplianceRuleParams are the parameters of New-DlpComplianceRule.
type NewDlpComplianceRuleParams struct {
	AccessedBy                                   []string `ps:"AccessedBy"`
	AccessedByMemberOf                           []string `ps:"AccessedByMemberOf"`
	AccessScope                                  any      `ps:"AccessScope"`
	AccessTimeControl                            any      `ps:"AccessTimeControl"`
	ActivationDate                               any      `ps:"ActivationDate"`
	AddRecipients                                any      `ps:"AddRecipients"`
	AdvancedRule                                 string   `ps:"AdvancedRule"`
	AlertProperties                              any      `ps:"AlertProperties"`
	AnyOfRecipientAddressContainsWords           any      `ps:"AnyOfRecipientAddressContainsWords"`
	AnyOfRecipientAddressMatchesPatterns         any      `ps:"AnyOfRecipientAddressMatchesPatterns"`
	ApplyBrandingTemplate                        string   `ps:"ApplyBrandingTemplate"`
	ApplyHtmlDisclaimer                          any      `ps:"ApplyHtmlDisclaimer"`
	AttachmentCountOver                          any      `ps:"AttachmentCountOver"`
	AttachmentIsNotLabeled                       bool     `ps:"AttachmentIsNotLabeled"`
	BlockAccess                                  bool     `ps:"BlockAccess"`
	BlockAccessScope                             any      `ps:"BlockAccessScope"`
	Comment                                      string   `ps:"Comment"`
	ContentCharacterSetContainsWords             any      `ps:"ContentCharacterSetContainsWords"`
	ContentContainsSensitiveInformation          []string `ps:"ContentContainsSensitiveInformation"`
	ContentExtensionMatchesWords                 any      `ps:"ContentExtensionMatchesWords"`
	ContentFileTypeMatches                       any      `ps:"ContentFileTypeMatches"`
	ContentIsNotLabeled                          bool     `ps:"ContentIsNotLabeled"`
	ContentIsShared                              bool     `ps:"ContentIsShared"`
	ContentMissingSensitivityLabel               any      `ps:"ContentMissingSensitivityLabel"`
	ContentPropertyContainsWords                 any      `ps:"ContentPropertyContainsWords"`
	ContextPropertiesContainWords                any      `ps:"ContextPropertiesContainWords"`
	DeviceManagementType                         any      `ps:"DeviceManagementType"` // one of: Managed, Unmanaged
	Disabled                                     bool     `ps:"Disabled"`
	DisplayName                                  string   `ps:"DisplayName"`
	DocumentContainsWords                        any      `ps:"DocumentContainsWords"`
	DocumentCreatedBy                            any      `ps:"DocumentCreatedBy"`
	DocumentCreatedByMemberOf                    []string `ps:"DocumentCreatedByMemberOf"`
	DocumentIsPasswordProtected                  bool     `ps:"DocumentIsPasswordProtected"`
	DocumentIsUnsupported                        bool     `ps:"DocumentIsUnsupported"`
	DocumentMatchesPatterns                      any      `ps:"DocumentMatchesPatterns"`
	DocumentNameMatchesPatterns                  any      `ps:"DocumentNameMatchesPatterns"`
	DocumentNameMatchesWords                     any      `ps:"DocumentNameMatchesWords"`
	DocumentSizeOver                             any      `ps:"DocumentSizeOver"`
	EncryptRMSTemplate                           any      `ps:"EncryptRMSTemplate"`
	EndpointDlpBrowserRestrictions               []string `ps:"EndpointDlpBrowserRestrictions"`
	EndpointDlpRestrictions                      []string `ps:"EndpointDlpRestrictions"`
	EnforcePortalAccess                          bool     `ps:"EnforcePortalAccess"`
	EvaluateRulePerComponent                     bool     `ps:"EvaluateRulePerComponent"`
	ExceptIfAccessedBy                           []string `ps:"ExceptIfAccessedBy"`
	ExceptIfAccessedByMemberOf                   []string `ps:"ExceptIfAccessedByMemberOf"`
	ExceptIfAccessScope                          any      `ps:"ExceptIfAccessScope"`
	ExceptIfAnyOfRecipientAddressContainsWords   any      `ps:"ExceptIfAnyOfRecipientAddressContainsWords"`
	ExceptIfAnyOfRecipientAddressMatchesPatterns any      `ps:"ExceptIfAnyOfRecipientAddressMatchesPatterns"`
	ExceptIfContentCharacterSetContainsWords     any      `ps:"ExceptIfContentCharacterSetContainsWords"`
	ExceptIfContentContainsSensitiveInformation  []string `ps:"ExceptIfContentContainsSensitiveInformation"`
	ExceptIfContentExtensionMatchesWords         any      `ps:"ExceptIfContentExtensionMatchesWords"`
	ExceptIfContentFileTypeMatches               any      `ps:"ExceptIfContentFileTypeMatches"`
	ExceptIfContentIsShared                      bool     `ps:"ExceptIfContentIsShared"`
	ExceptIfContentPropertyContainsWords         any      `ps:"ExceptIfContentPropertyContainsWords"`
	ExceptIfContextPropertiesContainWords        any      `ps:"ExceptIfContextPropertiesContainWords"`
	ExceptIfDeviceManagementType                 any      `ps:"ExceptIfDeviceManagementType"` // one of: Managed, Unmanaged
	ExceptIfDocumentContainsWords                any      `ps:"ExceptIfDocumentContainsWords"`
	ExceptIfDocumentCreatedBy                    any      `ps:"ExceptIfDocumentCreatedBy"`
	ExceptIfDocumentCreatedByMemberOf            []string `ps:"ExceptIfDocumentCreatedByMemberOf"`
	ExceptIfDocumentIsPasswordProtected          bool     `ps:"ExceptIfDocumentIsPasswordProtected"`
	ExceptIfDocumentIsUnsupported                bool     `ps:"ExceptIfDocumentIsUnsupported"`
	ExceptIfDocumentMatchesPatterns              any      `ps:"ExceptIfDocumentMatchesPatterns"`
	ExceptIfDocumentNameMatchesPatterns          any      `ps:"ExceptIfDocumentNameMatchesPatterns"`
	ExceptIfDocumentNameMatchesWords             any      `ps:"ExceptIfDocumentNameMatchesWords"`
	ExceptIfDocumentSizeOver                     any      `ps:"ExceptIfDocumentSizeOver"`
	ExceptIfFrom                                 []string `ps:"ExceptIfFrom"`
	ExceptIfFromAddressContainsWords             any      `ps:"ExceptIfFromAddressContainsWords"`
	ExceptIfFromAddressMatchesPatterns           any      `ps:"ExceptIfFromAddressMatchesPatterns"`
	ExceptIfFromMemberOf                         []string `ps:"ExceptIfFromMemberOf"`
	ExceptIfFromScope                            any      `ps:"ExceptIfFromScope"`
	ExceptIfHasSenderOverride                    bool     `ps:"ExceptIfHasSenderOverride"`
	ExceptIfHeaderContainsTokens                 any      `ps:"ExceptIfHeaderContainsTokens"`
	ExceptIfHeaderContainsWords                  any      `ps:"ExceptIfHeaderContainsWords"`
	ExceptIfHeaderMatchesPatterns                any      `ps:"ExceptIfHeaderMatchesPatterns"`
	ExceptIfMessageLabelChangeDetected           any      `ps:"ExceptIfMessageLabelChangeDetected"`
	ExceptIfMessageSizeOver                      any      `ps:"ExceptIfMessageSizeOver"`
	ExceptIfMessageTypeMatches                   any      `ps:"ExceptIfMessageTypeMatches"`
	ExceptIfProcessingLimitExceeded              bool     `ps:"ExceptIfProcessingLimitExceeded"`
	ExceptIfRecipientADAttributeContainsWords    any      `ps:"ExceptIfRecipientADAttributeContainsWords"`
	ExceptIfRecipientADAttributeMatchesPatterns  any      `ps:"ExceptIfRecipientADAttributeMatchesPatterns"`
	ExceptIfRecipientDomainIs                    any      `ps:"ExceptIfRecipientDomainIs"`
	ExceptIfSenderADAttributeContainsWords       any      `ps:"ExceptIfSenderADAttributeContainsWords"`
	ExceptIfSenderADAttributeMatchesPatterns     any      `ps:"ExceptIfSenderADAttributeMatchesPatterns"`
	ExceptIfSenderDomainIs                       any      `ps:"ExceptIfSenderDomainIs"`
	ExceptIfSenderIPRanges                       any      `ps:"ExceptIfSenderIPRanges"`
	ExceptIfSentTo                               any      `ps:"ExceptIfSentTo"`
	ExceptIfSentToMemberOf                       []string `ps:"ExceptIfSentToMemberOf"`
	ExceptIfSharedWithDomain                     any      `ps:"ExceptIfSharedWithDomain"`
	ExceptIfSubjectContainsWords                 any      `ps:"ExceptIfSubjectContainsWords"`
	ExceptIfSubjectMatchesPatterns               any      `ps:"ExceptIfSubjectMatchesPatterns"`
	ExceptIfSubjectOrBodyContainsWords           any      `ps:"ExceptIfSubjectOrBodyContainsWords"`
	ExceptIfSubjectOrBodyMatchesPatterns         any      `ps:"ExceptIfSubjectOrBodyMatchesPatterns"`
	ExceptIfTeamsSharedWithParticipantCategory   any      `ps:"ExceptIfTeamsSharedWithParticipantCategory"`
	ExceptIfUnscannableDocumentExtensionIs       any      `ps:"ExceptIfUnscannableDocumentExtensionIs"`
	ExceptIfUrlContainsText                      any      `ps:"ExceptIfUrlContainsText"`
	ExceptIfWithImportance                       any      `ps:"ExceptIfWithImportance"`
	ExpiryDate                                   any      `ps:"ExpiryDate"`
	From                                         []string `ps:"From"`
	FromAddressContainsWords                     any      `ps:"FromAddressContainsWords"`
	FromAddressMatchesPatterns                   any      `ps:"FromAddressMatchesPatterns"`
	FromMemberOf                                 []string `ps:"FromMemberOf"`
	FromScope                                    any      `ps:"FromScope"`
	GenerateAlert                                any      `ps:"GenerateAlert"`
	GenerateIncidentReport                       any      `ps:"GenerateIncidentReport"`
	HasActivity                                  any      `ps:"HasActivity"`
	HasLabelDowngradedFrom                       any      `ps:"HasLabelDowngradedFrom"`
	HasSenderOverride                            bool     `ps:"HasSenderOverride"`
	HeaderContainsTokens                         any      `ps:"HeaderContainsTokens"`
	HeaderContainsWords                          any      `ps:"HeaderContainsWords"`
	HeaderMatchesPatterns                        any      `ps:"HeaderMatchesPatterns"`
	ImmutableId                                  any      `ps:"ImmutableId"`
	IncidentReportContent                        []string `ps:"IncidentReportContent"`
	MapRecipients                                any      `ps:"MapRecipients"`
	MessageIsNotLabeled                          bool     `ps:"MessageIsNotLabeled"`
	MessageLabelChangeDetected                   any      `ps:"MessageLabelChangeDetected"`
	MessageSizeOver                              any      `ps:"MessageSizeOver"`
	MessageTypeMatches                           any      `ps:"MessageTypeMatches"`
	MipRestrictAccess                            []string `ps:"MipRestrictAccess"`
	Moderate                                     any      `ps:"Moderate"`
	ModifySubject                                any      `ps:"ModifySubject"`
	MoveToQuarantineLocation                     bool     `ps:"MoveToQuarantineLocation"`
	Name                                         string   `ps:"Name"`
	NonBifurcatingAccessScope                    any      `ps:"NonBifurcatingAccessScope"`
	NotifyAllowOverride                          []string `ps:"NotifyAllowOverride"`
	NotifyEmailCustomSenderDisplayName           string   `ps:"NotifyEmailCustomSenderDisplayName"`
	NotifyEmailCustomSubject                     string   `ps:"NotifyEmailCustomSubject"`
	NotifyEmailCustomText                        string   `ps:"NotifyEmailCustomText"`
	NotifyEmailExchangeIncludeAttachment         bool     `ps:"NotifyEmailExchangeIncludeAttachment"`
	NotifyEmailOnedriveRemediationActions        any      `ps:"NotifyEmailOnedriveRemediationActions"`
	NotifyEndpointUser                           any      `ps:"NotifyEndpointUser"`
	NotifyJustificationCustomText                string   `ps:"NotifyJustificationCustomText"`
	NotifyJustificationCustomTextTranslations    any      `ps:"NotifyJustificationCustomTextTranslations"`
	NotifyOverrideRequirements                   any      `ps:"NotifyOverrideRequirements"`
	NotifyPolicyTipCustomDialog                  string   `ps:"NotifyPolicyTipCustomDialog"`
	NotifyPolicyTipCustomText                    string   `ps:"NotifyPolicyTipCustomText"`
	NotifyPolicyTipCustomTextTranslations        any      `ps:"NotifyPolicyTipCustomTextTranslations"`
	NotifyPolicyTipDisplayOption                 any      `ps:"NotifyPolicyTipDisplayOption"`
	NotifyPolicyTipUrl                           string   `ps:"NotifyPolicyTipUrl"`
	NotifyUser                                   any      `ps:"NotifyUser"`
	NotifyUserType                               any      `ps:"NotifyUserType"`
	OnPremisesScannerDlpRestrictions             []string `ps:"OnPremisesScannerDlpRestrictions"`
	Policy                                       any      `ps:"Policy"`
	PowerBIDlpRestrictions                       []string `ps:"PowerBIDlpRestrictions"`
	PrependSubject                               string   `ps:"PrependSubject"`
	Priority                                     any      `ps:"Priority"`
	ProcessingLimitExceeded                      bool     `ps:"ProcessingLimitExceeded"`
	Quarantine                                   bool     `ps:"Quarantine"`
	RecipientADAttributeContainsWords            any      `ps:"RecipientADAttributeContainsWords"`
	RecipientADAttributeMatchesPatterns          any      `ps:"RecipientADAttributeMatchesPatterns"`
	RecipientDomainIs                            any      `ps:"RecipientDomainIs"`
	RedactSensitiveInfo                          any      `ps:"RedactSensitiveInfo"`
	RedirectMessageTo                            []string `ps:"RedirectMessageTo"`
	RemoveHeader                                 any      `ps:"RemoveHeader"`
	RemoveRMSTemplate                            bool     `ps:"RemoveRMSTemplate"`
	ReportSeverityLevel                          any      `ps:"ReportSeverityLevel"`
	RestrictAccess                               []string `ps:"RestrictAccess"`
	RestrictBrowserAccess                        bool     `ps:"RestrictBrowserAccess"`
	RestrictWebGrounding                         bool     `ps:"RestrictWebGrounding"`
	RuleErrorAction                              any      `ps:"RuleErrorAction"`
	SenderADAttributeContainsWords               any      `ps:"SenderADAttributeContainsWords"`
	SenderADAttributeMatchesPatterns             any      `ps:"SenderADAttributeMatchesPatterns"`
	SenderAddressLocation                        any      `ps:"SenderAddressLocation"`
	SenderDomainIs                               any      `ps:"SenderDomainIs"`
	SenderIPRanges                               any      `ps:"SenderIPRanges"`
	SenderType                                   any      `ps:"SenderType"`
	SentTo                                       any      `ps:"SentTo"`
	SentToMemberOf                               []string `ps:"SentToMemberOf"`
	SetHeader                                    any      `ps:"SetHeader"`
	SharedByIRMAgentRisk                         any      `ps:"SharedByIRMAgentRisk"`
	SharedByIRMUserRisk                          any      `ps:"SharedByIRMUserRisk"`
	SharedWithDomain                             any      `ps:"SharedWithDomain"`
	SharepointBlockDomains                       any      `ps:"SharepointBlockDomains"`
	SharepointBlockDomainsExcept                 any      `ps:"SharepointBlockDomainsExcept"`
	SharepointBlockDomainsOrUsers                bool     `ps:"SharepointBlockDomainsOrUsers"`
	SharepointBlockUsers                         any      `ps:"SharepointBlockUsers"`
	SharepointBlockUsersExcept                   any      `ps:"SharepointBlockUsersExcept"`
	SharepointMoveToQuarantineLocation           bool     `ps:"SharepointMoveToQuarantineLocation"`
	StopPolicyProcessing                         bool     `ps:"StopPolicyProcessing"`
	SubjectContainsWords                         any      `ps:"SubjectContainsWords"`
	SubjectMatchesPatterns                       any      `ps:"SubjectMatchesPatterns"`
	SubjectOrBodyContainsWords                   any      `ps:"SubjectOrBodyContainsWords"`
	SubjectOrBodyMatchesPatterns                 any      `ps:"SubjectOrBodyMatchesPatterns"`
	TeamsSharedWithParticipantCategory           any      `ps:"TeamsSharedWithParticipantCategory"`
	ThirdPartyAppDlpRestrictions                 []string `ps:"ThirdPartyAppDlpRestrictions"`
	TriggerPowerAutomateFlow                     string   `ps:"TriggerPowerAutomateFlow"`
	UnscannableDocumentExtensionIs               any      `ps:"UnscannableDocumentExtensionIs"`
	UrlContainsText                              any      `ps:"UrlContainsText"`
	ValidateRule                                 bool     `ps:"ValidateRule"`
	WithImportance                               any      `ps:"WithImportance"`
}

func (p NewDlpComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if len(p.AccessedBy) > 0 {
		m["AccessedBy"] = p.AccessedBy
	}
	if len(p.AccessedByMemberOf) > 0 {
		m["AccessedByMemberOf"] = p.AccessedByMemberOf
	}
	if p.AccessScope != nil {
		m["AccessScope"] = p.AccessScope
	}
	if p.AccessTimeControl != nil {
		m["AccessTimeControl"] = p.AccessTimeControl
	}
	if p.ActivationDate != nil {
		m["ActivationDate"] = p.ActivationDate
	}
	if p.AddRecipients != nil {
		m["AddRecipients"] = p.AddRecipients
	}
	if p.AdvancedRule != "" {
		m["AdvancedRule"] = p.AdvancedRule
	}
	if p.AlertProperties != nil {
		m["AlertProperties"] = p.AlertProperties
	}
	if p.AnyOfRecipientAddressContainsWords != nil {
		m["AnyOfRecipientAddressContainsWords"] = p.AnyOfRecipientAddressContainsWords
	}
	if p.AnyOfRecipientAddressMatchesPatterns != nil {
		m["AnyOfRecipientAddressMatchesPatterns"] = p.AnyOfRecipientAddressMatchesPatterns
	}
	if p.ApplyBrandingTemplate != "" {
		m["ApplyBrandingTemplate"] = p.ApplyBrandingTemplate
	}
	if p.ApplyHtmlDisclaimer != nil {
		m["ApplyHtmlDisclaimer"] = p.ApplyHtmlDisclaimer
	}
	if p.AttachmentCountOver != nil {
		m["AttachmentCountOver"] = p.AttachmentCountOver
	}
	if p.AttachmentIsNotLabeled {
		m["AttachmentIsNotLabeled"] = true
	}
	if p.BlockAccess {
		m["BlockAccess"] = true
	}
	if p.BlockAccessScope != nil {
		m["BlockAccessScope"] = p.BlockAccessScope
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.ContentCharacterSetContainsWords != nil {
		m["ContentCharacterSetContainsWords"] = p.ContentCharacterSetContainsWords
	}
	if len(p.ContentContainsSensitiveInformation) > 0 {
		m["ContentContainsSensitiveInformation"] = p.ContentContainsSensitiveInformation
	}
	if p.ContentExtensionMatchesWords != nil {
		m["ContentExtensionMatchesWords"] = p.ContentExtensionMatchesWords
	}
	if p.ContentFileTypeMatches != nil {
		m["ContentFileTypeMatches"] = p.ContentFileTypeMatches
	}
	if p.ContentIsNotLabeled {
		m["ContentIsNotLabeled"] = true
	}
	if p.ContentIsShared {
		m["ContentIsShared"] = true
	}
	if p.ContentMissingSensitivityLabel != nil {
		m["ContentMissingSensitivityLabel"] = p.ContentMissingSensitivityLabel
	}
	if p.ContentPropertyContainsWords != nil {
		m["ContentPropertyContainsWords"] = p.ContentPropertyContainsWords
	}
	if p.ContextPropertiesContainWords != nil {
		m["ContextPropertiesContainWords"] = p.ContextPropertiesContainWords
	}
	if p.DeviceManagementType != nil {
		m["DeviceManagementType"] = p.DeviceManagementType
	}
	if p.Disabled {
		m["Disabled"] = true
	}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.DocumentContainsWords != nil {
		m["DocumentContainsWords"] = p.DocumentContainsWords
	}
	if p.DocumentCreatedBy != nil {
		m["DocumentCreatedBy"] = p.DocumentCreatedBy
	}
	if len(p.DocumentCreatedByMemberOf) > 0 {
		m["DocumentCreatedByMemberOf"] = p.DocumentCreatedByMemberOf
	}
	if p.DocumentIsPasswordProtected {
		m["DocumentIsPasswordProtected"] = true
	}
	if p.DocumentIsUnsupported {
		m["DocumentIsUnsupported"] = true
	}
	if p.DocumentMatchesPatterns != nil {
		m["DocumentMatchesPatterns"] = p.DocumentMatchesPatterns
	}
	if p.DocumentNameMatchesPatterns != nil {
		m["DocumentNameMatchesPatterns"] = p.DocumentNameMatchesPatterns
	}
	if p.DocumentNameMatchesWords != nil {
		m["DocumentNameMatchesWords"] = p.DocumentNameMatchesWords
	}
	if p.DocumentSizeOver != nil {
		m["DocumentSizeOver"] = p.DocumentSizeOver
	}
	if p.EncryptRMSTemplate != nil {
		m["EncryptRMSTemplate"] = p.EncryptRMSTemplate
	}
	if len(p.EndpointDlpBrowserRestrictions) > 0 {
		m["EndpointDlpBrowserRestrictions"] = p.EndpointDlpBrowserRestrictions
	}
	if len(p.EndpointDlpRestrictions) > 0 {
		m["EndpointDlpRestrictions"] = p.EndpointDlpRestrictions
	}
	if p.EnforcePortalAccess {
		m["EnforcePortalAccess"] = true
	}
	if p.EvaluateRulePerComponent {
		m["EvaluateRulePerComponent"] = true
	}
	if len(p.ExceptIfAccessedBy) > 0 {
		m["ExceptIfAccessedBy"] = p.ExceptIfAccessedBy
	}
	if len(p.ExceptIfAccessedByMemberOf) > 0 {
		m["ExceptIfAccessedByMemberOf"] = p.ExceptIfAccessedByMemberOf
	}
	if p.ExceptIfAccessScope != nil {
		m["ExceptIfAccessScope"] = p.ExceptIfAccessScope
	}
	if p.ExceptIfAnyOfRecipientAddressContainsWords != nil {
		m["ExceptIfAnyOfRecipientAddressContainsWords"] = p.ExceptIfAnyOfRecipientAddressContainsWords
	}
	if p.ExceptIfAnyOfRecipientAddressMatchesPatterns != nil {
		m["ExceptIfAnyOfRecipientAddressMatchesPatterns"] = p.ExceptIfAnyOfRecipientAddressMatchesPatterns
	}
	if p.ExceptIfContentCharacterSetContainsWords != nil {
		m["ExceptIfContentCharacterSetContainsWords"] = p.ExceptIfContentCharacterSetContainsWords
	}
	if len(p.ExceptIfContentContainsSensitiveInformation) > 0 {
		m["ExceptIfContentContainsSensitiveInformation"] = p.ExceptIfContentContainsSensitiveInformation
	}
	if p.ExceptIfContentExtensionMatchesWords != nil {
		m["ExceptIfContentExtensionMatchesWords"] = p.ExceptIfContentExtensionMatchesWords
	}
	if p.ExceptIfContentFileTypeMatches != nil {
		m["ExceptIfContentFileTypeMatches"] = p.ExceptIfContentFileTypeMatches
	}
	if p.ExceptIfContentIsShared {
		m["ExceptIfContentIsShared"] = true
	}
	if p.ExceptIfContentPropertyContainsWords != nil {
		m["ExceptIfContentPropertyContainsWords"] = p.ExceptIfContentPropertyContainsWords
	}
	if p.ExceptIfContextPropertiesContainWords != nil {
		m["ExceptIfContextPropertiesContainWords"] = p.ExceptIfContextPropertiesContainWords
	}
	if p.ExceptIfDeviceManagementType != nil {
		m["ExceptIfDeviceManagementType"] = p.ExceptIfDeviceManagementType
	}
	if p.ExceptIfDocumentContainsWords != nil {
		m["ExceptIfDocumentContainsWords"] = p.ExceptIfDocumentContainsWords
	}
	if p.ExceptIfDocumentCreatedBy != nil {
		m["ExceptIfDocumentCreatedBy"] = p.ExceptIfDocumentCreatedBy
	}
	if len(p.ExceptIfDocumentCreatedByMemberOf) > 0 {
		m["ExceptIfDocumentCreatedByMemberOf"] = p.ExceptIfDocumentCreatedByMemberOf
	}
	if p.ExceptIfDocumentIsPasswordProtected {
		m["ExceptIfDocumentIsPasswordProtected"] = true
	}
	if p.ExceptIfDocumentIsUnsupported {
		m["ExceptIfDocumentIsUnsupported"] = true
	}
	if p.ExceptIfDocumentMatchesPatterns != nil {
		m["ExceptIfDocumentMatchesPatterns"] = p.ExceptIfDocumentMatchesPatterns
	}
	if p.ExceptIfDocumentNameMatchesPatterns != nil {
		m["ExceptIfDocumentNameMatchesPatterns"] = p.ExceptIfDocumentNameMatchesPatterns
	}
	if p.ExceptIfDocumentNameMatchesWords != nil {
		m["ExceptIfDocumentNameMatchesWords"] = p.ExceptIfDocumentNameMatchesWords
	}
	if p.ExceptIfDocumentSizeOver != nil {
		m["ExceptIfDocumentSizeOver"] = p.ExceptIfDocumentSizeOver
	}
	if len(p.ExceptIfFrom) > 0 {
		m["ExceptIfFrom"] = p.ExceptIfFrom
	}
	if p.ExceptIfFromAddressContainsWords != nil {
		m["ExceptIfFromAddressContainsWords"] = p.ExceptIfFromAddressContainsWords
	}
	if p.ExceptIfFromAddressMatchesPatterns != nil {
		m["ExceptIfFromAddressMatchesPatterns"] = p.ExceptIfFromAddressMatchesPatterns
	}
	if len(p.ExceptIfFromMemberOf) > 0 {
		m["ExceptIfFromMemberOf"] = p.ExceptIfFromMemberOf
	}
	if p.ExceptIfFromScope != nil {
		m["ExceptIfFromScope"] = p.ExceptIfFromScope
	}
	if p.ExceptIfHasSenderOverride {
		m["ExceptIfHasSenderOverride"] = true
	}
	if p.ExceptIfHeaderContainsTokens != nil {
		m["ExceptIfHeaderContainsTokens"] = p.ExceptIfHeaderContainsTokens
	}
	if p.ExceptIfHeaderContainsWords != nil {
		m["ExceptIfHeaderContainsWords"] = p.ExceptIfHeaderContainsWords
	}
	if p.ExceptIfHeaderMatchesPatterns != nil {
		m["ExceptIfHeaderMatchesPatterns"] = p.ExceptIfHeaderMatchesPatterns
	}
	if p.ExceptIfMessageLabelChangeDetected != nil {
		m["ExceptIfMessageLabelChangeDetected"] = p.ExceptIfMessageLabelChangeDetected
	}
	if p.ExceptIfMessageSizeOver != nil {
		m["ExceptIfMessageSizeOver"] = p.ExceptIfMessageSizeOver
	}
	if p.ExceptIfMessageTypeMatches != nil {
		m["ExceptIfMessageTypeMatches"] = p.ExceptIfMessageTypeMatches
	}
	if p.ExceptIfProcessingLimitExceeded {
		m["ExceptIfProcessingLimitExceeded"] = true
	}
	if p.ExceptIfRecipientADAttributeContainsWords != nil {
		m["ExceptIfRecipientADAttributeContainsWords"] = p.ExceptIfRecipientADAttributeContainsWords
	}
	if p.ExceptIfRecipientADAttributeMatchesPatterns != nil {
		m["ExceptIfRecipientADAttributeMatchesPatterns"] = p.ExceptIfRecipientADAttributeMatchesPatterns
	}
	if p.ExceptIfRecipientDomainIs != nil {
		m["ExceptIfRecipientDomainIs"] = p.ExceptIfRecipientDomainIs
	}
	if p.ExceptIfSenderADAttributeContainsWords != nil {
		m["ExceptIfSenderADAttributeContainsWords"] = p.ExceptIfSenderADAttributeContainsWords
	}
	if p.ExceptIfSenderADAttributeMatchesPatterns != nil {
		m["ExceptIfSenderADAttributeMatchesPatterns"] = p.ExceptIfSenderADAttributeMatchesPatterns
	}
	if p.ExceptIfSenderDomainIs != nil {
		m["ExceptIfSenderDomainIs"] = p.ExceptIfSenderDomainIs
	}
	if p.ExceptIfSenderIPRanges != nil {
		m["ExceptIfSenderIPRanges"] = p.ExceptIfSenderIPRanges
	}
	if p.ExceptIfSentTo != nil {
		m["ExceptIfSentTo"] = p.ExceptIfSentTo
	}
	if len(p.ExceptIfSentToMemberOf) > 0 {
		m["ExceptIfSentToMemberOf"] = p.ExceptIfSentToMemberOf
	}
	if p.ExceptIfSharedWithDomain != nil {
		m["ExceptIfSharedWithDomain"] = p.ExceptIfSharedWithDomain
	}
	if p.ExceptIfSubjectContainsWords != nil {
		m["ExceptIfSubjectContainsWords"] = p.ExceptIfSubjectContainsWords
	}
	if p.ExceptIfSubjectMatchesPatterns != nil {
		m["ExceptIfSubjectMatchesPatterns"] = p.ExceptIfSubjectMatchesPatterns
	}
	if p.ExceptIfSubjectOrBodyContainsWords != nil {
		m["ExceptIfSubjectOrBodyContainsWords"] = p.ExceptIfSubjectOrBodyContainsWords
	}
	if p.ExceptIfSubjectOrBodyMatchesPatterns != nil {
		m["ExceptIfSubjectOrBodyMatchesPatterns"] = p.ExceptIfSubjectOrBodyMatchesPatterns
	}
	if p.ExceptIfTeamsSharedWithParticipantCategory != nil {
		m["ExceptIfTeamsSharedWithParticipantCategory"] = p.ExceptIfTeamsSharedWithParticipantCategory
	}
	if p.ExceptIfUnscannableDocumentExtensionIs != nil {
		m["ExceptIfUnscannableDocumentExtensionIs"] = p.ExceptIfUnscannableDocumentExtensionIs
	}
	if p.ExceptIfUrlContainsText != nil {
		m["ExceptIfUrlContainsText"] = p.ExceptIfUrlContainsText
	}
	if p.ExceptIfWithImportance != nil {
		m["ExceptIfWithImportance"] = p.ExceptIfWithImportance
	}
	if p.ExpiryDate != nil {
		m["ExpiryDate"] = p.ExpiryDate
	}
	if len(p.From) > 0 {
		m["From"] = p.From
	}
	if p.FromAddressContainsWords != nil {
		m["FromAddressContainsWords"] = p.FromAddressContainsWords
	}
	if p.FromAddressMatchesPatterns != nil {
		m["FromAddressMatchesPatterns"] = p.FromAddressMatchesPatterns
	}
	if len(p.FromMemberOf) > 0 {
		m["FromMemberOf"] = p.FromMemberOf
	}
	if p.FromScope != nil {
		m["FromScope"] = p.FromScope
	}
	if p.GenerateAlert != nil {
		m["GenerateAlert"] = p.GenerateAlert
	}
	if p.GenerateIncidentReport != nil {
		m["GenerateIncidentReport"] = p.GenerateIncidentReport
	}
	if p.HasActivity != nil {
		m["HasActivity"] = p.HasActivity
	}
	if p.HasLabelDowngradedFrom != nil {
		m["HasLabelDowngradedFrom"] = p.HasLabelDowngradedFrom
	}
	if p.HasSenderOverride {
		m["HasSenderOverride"] = true
	}
	if p.HeaderContainsTokens != nil {
		m["HeaderContainsTokens"] = p.HeaderContainsTokens
	}
	if p.HeaderContainsWords != nil {
		m["HeaderContainsWords"] = p.HeaderContainsWords
	}
	if p.HeaderMatchesPatterns != nil {
		m["HeaderMatchesPatterns"] = p.HeaderMatchesPatterns
	}
	if p.ImmutableId != nil {
		m["ImmutableId"] = p.ImmutableId
	}
	if len(p.IncidentReportContent) > 0 {
		m["IncidentReportContent"] = p.IncidentReportContent
	}
	if p.MapRecipients != nil {
		m["MapRecipients"] = p.MapRecipients
	}
	if p.MessageIsNotLabeled {
		m["MessageIsNotLabeled"] = true
	}
	if p.MessageLabelChangeDetected != nil {
		m["MessageLabelChangeDetected"] = p.MessageLabelChangeDetected
	}
	if p.MessageSizeOver != nil {
		m["MessageSizeOver"] = p.MessageSizeOver
	}
	if p.MessageTypeMatches != nil {
		m["MessageTypeMatches"] = p.MessageTypeMatches
	}
	if len(p.MipRestrictAccess) > 0 {
		m["MipRestrictAccess"] = p.MipRestrictAccess
	}
	if p.Moderate != nil {
		m["Moderate"] = p.Moderate
	}
	if p.ModifySubject != nil {
		m["ModifySubject"] = p.ModifySubject
	}
	if p.MoveToQuarantineLocation {
		m["MoveToQuarantineLocation"] = true
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.NonBifurcatingAccessScope != nil {
		m["NonBifurcatingAccessScope"] = p.NonBifurcatingAccessScope
	}
	if len(p.NotifyAllowOverride) > 0 {
		m["NotifyAllowOverride"] = p.NotifyAllowOverride
	}
	if p.NotifyEmailCustomSenderDisplayName != "" {
		m["NotifyEmailCustomSenderDisplayName"] = p.NotifyEmailCustomSenderDisplayName
	}
	if p.NotifyEmailCustomSubject != "" {
		m["NotifyEmailCustomSubject"] = p.NotifyEmailCustomSubject
	}
	if p.NotifyEmailCustomText != "" {
		m["NotifyEmailCustomText"] = p.NotifyEmailCustomText
	}
	if p.NotifyEmailExchangeIncludeAttachment {
		m["NotifyEmailExchangeIncludeAttachment"] = true
	}
	if p.NotifyEmailOnedriveRemediationActions != nil {
		m["NotifyEmailOnedriveRemediationActions"] = p.NotifyEmailOnedriveRemediationActions
	}
	if p.NotifyEndpointUser != nil {
		m["NotifyEndpointUser"] = p.NotifyEndpointUser
	}
	if p.NotifyJustificationCustomText != "" {
		m["NotifyJustificationCustomText"] = p.NotifyJustificationCustomText
	}
	if p.NotifyJustificationCustomTextTranslations != nil {
		m["NotifyJustificationCustomTextTranslations"] = p.NotifyJustificationCustomTextTranslations
	}
	if p.NotifyOverrideRequirements != nil {
		m["NotifyOverrideRequirements"] = p.NotifyOverrideRequirements
	}
	if p.NotifyPolicyTipCustomDialog != "" {
		m["NotifyPolicyTipCustomDialog"] = p.NotifyPolicyTipCustomDialog
	}
	if p.NotifyPolicyTipCustomText != "" {
		m["NotifyPolicyTipCustomText"] = p.NotifyPolicyTipCustomText
	}
	if p.NotifyPolicyTipCustomTextTranslations != nil {
		m["NotifyPolicyTipCustomTextTranslations"] = p.NotifyPolicyTipCustomTextTranslations
	}
	if p.NotifyPolicyTipDisplayOption != nil {
		m["NotifyPolicyTipDisplayOption"] = p.NotifyPolicyTipDisplayOption
	}
	if p.NotifyPolicyTipUrl != "" {
		m["NotifyPolicyTipUrl"] = p.NotifyPolicyTipUrl
	}
	if p.NotifyUser != nil {
		m["NotifyUser"] = p.NotifyUser
	}
	if p.NotifyUserType != nil {
		m["NotifyUserType"] = p.NotifyUserType
	}
	if len(p.OnPremisesScannerDlpRestrictions) > 0 {
		m["OnPremisesScannerDlpRestrictions"] = p.OnPremisesScannerDlpRestrictions
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	if len(p.PowerBIDlpRestrictions) > 0 {
		m["PowerBIDlpRestrictions"] = p.PowerBIDlpRestrictions
	}
	if p.PrependSubject != "" {
		m["PrependSubject"] = p.PrependSubject
	}
	if p.Priority != nil {
		m["Priority"] = p.Priority
	}
	if p.ProcessingLimitExceeded {
		m["ProcessingLimitExceeded"] = true
	}
	if p.Quarantine {
		m["Quarantine"] = true
	}
	if p.RecipientADAttributeContainsWords != nil {
		m["RecipientADAttributeContainsWords"] = p.RecipientADAttributeContainsWords
	}
	if p.RecipientADAttributeMatchesPatterns != nil {
		m["RecipientADAttributeMatchesPatterns"] = p.RecipientADAttributeMatchesPatterns
	}
	if p.RecipientDomainIs != nil {
		m["RecipientDomainIs"] = p.RecipientDomainIs
	}
	if p.RedactSensitiveInfo != nil {
		m["RedactSensitiveInfo"] = p.RedactSensitiveInfo
	}
	if len(p.RedirectMessageTo) > 0 {
		m["RedirectMessageTo"] = p.RedirectMessageTo
	}
	if p.RemoveHeader != nil {
		m["RemoveHeader"] = p.RemoveHeader
	}
	if p.RemoveRMSTemplate {
		m["RemoveRMSTemplate"] = true
	}
	if p.ReportSeverityLevel != nil {
		m["ReportSeverityLevel"] = p.ReportSeverityLevel
	}
	if len(p.RestrictAccess) > 0 {
		m["RestrictAccess"] = p.RestrictAccess
	}
	if p.RestrictBrowserAccess {
		m["RestrictBrowserAccess"] = true
	}
	if p.RestrictWebGrounding {
		m["RestrictWebGrounding"] = true
	}
	if p.RuleErrorAction != nil {
		m["RuleErrorAction"] = p.RuleErrorAction
	}
	if p.SenderADAttributeContainsWords != nil {
		m["SenderADAttributeContainsWords"] = p.SenderADAttributeContainsWords
	}
	if p.SenderADAttributeMatchesPatterns != nil {
		m["SenderADAttributeMatchesPatterns"] = p.SenderADAttributeMatchesPatterns
	}
	if p.SenderAddressLocation != nil {
		m["SenderAddressLocation"] = p.SenderAddressLocation
	}
	if p.SenderDomainIs != nil {
		m["SenderDomainIs"] = p.SenderDomainIs
	}
	if p.SenderIPRanges != nil {
		m["SenderIPRanges"] = p.SenderIPRanges
	}
	if p.SenderType != nil {
		m["SenderType"] = p.SenderType
	}
	if p.SentTo != nil {
		m["SentTo"] = p.SentTo
	}
	if len(p.SentToMemberOf) > 0 {
		m["SentToMemberOf"] = p.SentToMemberOf
	}
	if p.SetHeader != nil {
		m["SetHeader"] = p.SetHeader
	}
	if p.SharedByIRMAgentRisk != nil {
		m["SharedByIRMAgentRisk"] = p.SharedByIRMAgentRisk
	}
	if p.SharedByIRMUserRisk != nil {
		m["SharedByIRMUserRisk"] = p.SharedByIRMUserRisk
	}
	if p.SharedWithDomain != nil {
		m["SharedWithDomain"] = p.SharedWithDomain
	}
	if p.SharepointBlockDomains != nil {
		m["SharepointBlockDomains"] = p.SharepointBlockDomains
	}
	if p.SharepointBlockDomainsExcept != nil {
		m["SharepointBlockDomainsExcept"] = p.SharepointBlockDomainsExcept
	}
	if p.SharepointBlockDomainsOrUsers {
		m["SharepointBlockDomainsOrUsers"] = true
	}
	if p.SharepointBlockUsers != nil {
		m["SharepointBlockUsers"] = p.SharepointBlockUsers
	}
	if p.SharepointBlockUsersExcept != nil {
		m["SharepointBlockUsersExcept"] = p.SharepointBlockUsersExcept
	}
	if p.SharepointMoveToQuarantineLocation {
		m["SharepointMoveToQuarantineLocation"] = true
	}
	if p.StopPolicyProcessing {
		m["StopPolicyProcessing"] = true
	}
	if p.SubjectContainsWords != nil {
		m["SubjectContainsWords"] = p.SubjectContainsWords
	}
	if p.SubjectMatchesPatterns != nil {
		m["SubjectMatchesPatterns"] = p.SubjectMatchesPatterns
	}
	if p.SubjectOrBodyContainsWords != nil {
		m["SubjectOrBodyContainsWords"] = p.SubjectOrBodyContainsWords
	}
	if p.SubjectOrBodyMatchesPatterns != nil {
		m["SubjectOrBodyMatchesPatterns"] = p.SubjectOrBodyMatchesPatterns
	}
	if p.TeamsSharedWithParticipantCategory != nil {
		m["TeamsSharedWithParticipantCategory"] = p.TeamsSharedWithParticipantCategory
	}
	if len(p.ThirdPartyAppDlpRestrictions) > 0 {
		m["ThirdPartyAppDlpRestrictions"] = p.ThirdPartyAppDlpRestrictions
	}
	if p.TriggerPowerAutomateFlow != "" {
		m["TriggerPowerAutomateFlow"] = p.TriggerPowerAutomateFlow
	}
	if p.UnscannableDocumentExtensionIs != nil {
		m["UnscannableDocumentExtensionIs"] = p.UnscannableDocumentExtensionIs
	}
	if p.UrlContainsText != nil {
		m["UrlContainsText"] = p.UrlContainsText
	}
	if p.ValidateRule {
		m["ValidateRule"] = true
	}
	if p.WithImportance != nil {
		m["WithImportance"] = p.WithImportance
	}
	return m
}

// NewDlpComplianceRule runs the New-DlpComplianceRule cmdlet.
func (s *Service) NewDlpComplianceRule(ctx context.Context, p NewDlpComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-DlpComplianceRule", p.params())
}

// NewDlpEdmSchemaParams are the parameters of New-DlpEdmSchema.
type NewDlpEdmSchemaParams struct {
	FileData []string `ps:"FileData"`
}

func (p NewDlpEdmSchemaParams) params() map[string]any {
	m := map[string]any{}
	if len(p.FileData) > 0 {
		m["FileData"] = p.FileData
	}
	return m
}

// NewDlpEdmSchema runs the New-DlpEdmSchema cmdlet.
func (s *Service) NewDlpEdmSchema(ctx context.Context, p NewDlpEdmSchemaParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-DlpEdmSchema", p.params())
}

// NewDlpFingerprintParams are the parameters of New-DlpFingerprint.
type NewDlpFingerprintParams struct {
	Description     string   `ps:"Description"`
	FileData        []string `ps:"FileData"`
	IsExact         bool     `ps:"IsExact"`
	Threshold       any      `ps:"Threshold"`
	ThresholdConfig any      `ps:"ThresholdConfig"`
}

func (p NewDlpFingerprintParams) params() map[string]any {
	m := map[string]any{}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if len(p.FileData) > 0 {
		m["FileData"] = p.FileData
	}
	if p.IsExact {
		m["IsExact"] = true
	}
	if p.Threshold != nil {
		m["Threshold"] = p.Threshold
	}
	if p.ThresholdConfig != nil {
		m["ThresholdConfig"] = p.ThresholdConfig
	}
	return m
}

// NewDlpFingerprint runs the New-DlpFingerprint cmdlet.
func (s *Service) NewDlpFingerprint(ctx context.Context, p NewDlpFingerprintParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-DlpFingerprint", p.params())
}

// NewDlpKeywordDictionaryParams are the parameters of New-DlpKeywordDictionary.
type NewDlpKeywordDictionaryParams struct {
	Description            string   `ps:"Description"`
	DoNotPersistDictionary bool     `ps:"DoNotPersistDictionary"`
	DoNotPersistKeywords   bool     `ps:"DoNotPersistKeywords"`
	FileData               []string `ps:"FileData"`
	MatchStyle             string   `ps:"MatchStyle"` // one of: word, string
	Name                   string   `ps:"Name"`
	Organization           any      `ps:"Organization"`
}

func (p NewDlpKeywordDictionaryParams) params() map[string]any {
	m := map[string]any{}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if p.DoNotPersistDictionary {
		m["DoNotPersistDictionary"] = true
	}
	if p.DoNotPersistKeywords {
		m["DoNotPersistKeywords"] = true
	}
	if len(p.FileData) > 0 {
		m["FileData"] = p.FileData
	}
	if p.MatchStyle != "" {
		m["MatchStyle"] = p.MatchStyle
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Organization != nil {
		m["Organization"] = p.Organization
	}
	return m
}

// NewDlpKeywordDictionary runs the New-DlpKeywordDictionary cmdlet.
func (s *Service) NewDlpKeywordDictionary(ctx context.Context, p NewDlpKeywordDictionaryParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-DlpKeywordDictionary", p.params())
}

// NewDlpSensitiveInformationTypeParams are the parameters of New-DlpSensitiveInformationType.
type NewDlpSensitiveInformationTypeParams struct {
	Description     string   `ps:"Description"`
	FileData        []string `ps:"FileData"`
	Fingerprints    any      `ps:"Fingerprints"`
	IsExact         bool     `ps:"IsExact"`
	Locale          any      `ps:"Locale"`
	Name            string   `ps:"Name"`
	Threshold       any      `ps:"Threshold"`
	ThresholdConfig any      `ps:"ThresholdConfig"`
}

func (p NewDlpSensitiveInformationTypeParams) params() map[string]any {
	m := map[string]any{}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if len(p.FileData) > 0 {
		m["FileData"] = p.FileData
	}
	if p.Fingerprints != nil {
		m["Fingerprints"] = p.Fingerprints
	}
	if p.IsExact {
		m["IsExact"] = true
	}
	if p.Locale != nil {
		m["Locale"] = p.Locale
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Threshold != nil {
		m["Threshold"] = p.Threshold
	}
	if p.ThresholdConfig != nil {
		m["ThresholdConfig"] = p.ThresholdConfig
	}
	return m
}

// NewDlpSensitiveInformationType runs the New-DlpSensitiveInformationType cmdlet.
func (s *Service) NewDlpSensitiveInformationType(ctx context.Context, p NewDlpSensitiveInformationTypeParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-DlpSensitiveInformationType", p.params())
}

// NewDlpSensitiveInformationTypeRulePackageParams are the parameters of New-DlpSensitiveInformationTypeRulePackage.
type NewDlpSensitiveInformationTypeRulePackageParams struct {
	FileData    []string `ps:"FileData"`
	MigrationId string   `ps:"MigrationId"`
}

func (p NewDlpSensitiveInformationTypeRulePackageParams) params() map[string]any {
	m := map[string]any{}
	if len(p.FileData) > 0 {
		m["FileData"] = p.FileData
	}
	if p.MigrationId != "" {
		m["MigrationId"] = p.MigrationId
	}
	return m
}

// NewDlpSensitiveInformationTypeRulePackage runs the New-DlpSensitiveInformationTypeRulePackage cmdlet.
func (s *Service) NewDlpSensitiveInformationTypeRulePackage(ctx context.Context, p NewDlpSensitiveInformationTypeRulePackageParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-DlpSensitiveInformationTypeRulePackage", p.params())
}

// NewDspmPolicyParams are the parameters of New-DspmPolicy.
// DefaultParameterSetName: Identity
type NewDspmPolicyParams struct {
	Enabled       bool   `ps:"Enabled"`
	Name          string `ps:"Name"`
	TenantSetting string `ps:"TenantSetting"`
}

func (p NewDspmPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.TenantSetting != "" {
		m["TenantSetting"] = p.TenantSetting
	}
	return m
}

// NewDspmPolicy runs the New-DspmPolicy cmdlet.
func (s *Service) NewDspmPolicy(ctx context.Context, p NewDspmPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-DspmPolicy", p.params())
}

// NewFeatureConfigurationParams are the parameters of New-FeatureConfiguration.
type NewFeatureConfigurationParams struct {
	Comment          string `ps:"Comment"`
	FeatureScenario  any    `ps:"FeatureScenario"`
	Locations        string `ps:"Locations"`
	Mode             any    `ps:"Mode"`
	Name             string `ps:"Name"`
	PolicyRBACScopes any    `ps:"PolicyRBACScopes"`
	ScenarioConfig   string `ps:"ScenarioConfig"`
}

func (p NewFeatureConfigurationParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.FeatureScenario != nil {
		m["FeatureScenario"] = p.FeatureScenario
	}
	if p.Locations != "" {
		m["Locations"] = p.Locations
	}
	if p.Mode != nil {
		m["Mode"] = p.Mode
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.ScenarioConfig != "" {
		m["ScenarioConfig"] = p.ScenarioConfig
	}
	return m
}

// NewFeatureConfiguration runs the New-FeatureConfiguration cmdlet.
func (s *Service) NewFeatureConfiguration(ctx context.Context, p NewFeatureConfigurationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-FeatureConfiguration", p.params())
}

// NewFilePlanPropertyAuthorityParams are the parameters of New-FilePlanPropertyAuthority.
type NewFilePlanPropertyAuthorityParams struct {
	Name string `ps:"Name"`
}

func (p NewFilePlanPropertyAuthorityParams) params() map[string]any {
	m := map[string]any{}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	return m
}

// NewFilePlanPropertyAuthority runs the New-FilePlanPropertyAuthority cmdlet.
func (s *Service) NewFilePlanPropertyAuthority(ctx context.Context, p NewFilePlanPropertyAuthorityParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-FilePlanPropertyAuthority", p.params())
}

// NewFilePlanPropertyCategoryParams are the parameters of New-FilePlanPropertyCategory.
type NewFilePlanPropertyCategoryParams struct {
	Name string `ps:"Name"`
}

func (p NewFilePlanPropertyCategoryParams) params() map[string]any {
	m := map[string]any{}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	return m
}

// NewFilePlanPropertyCategory runs the New-FilePlanPropertyCategory cmdlet.
func (s *Service) NewFilePlanPropertyCategory(ctx context.Context, p NewFilePlanPropertyCategoryParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-FilePlanPropertyCategory", p.params())
}

// NewFilePlanPropertyCitationParams are the parameters of New-FilePlanPropertyCitation.
type NewFilePlanPropertyCitationParams struct {
	CitationJurisdiction string `ps:"CitationJurisdiction"`
	CitationUrl          string `ps:"CitationUrl"`
	Name                 string `ps:"Name"`
}

func (p NewFilePlanPropertyCitationParams) params() map[string]any {
	m := map[string]any{}
	if p.CitationJurisdiction != "" {
		m["CitationJurisdiction"] = p.CitationJurisdiction
	}
	if p.CitationUrl != "" {
		m["CitationUrl"] = p.CitationUrl
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	return m
}

// NewFilePlanPropertyCitation runs the New-FilePlanPropertyCitation cmdlet.
func (s *Service) NewFilePlanPropertyCitation(ctx context.Context, p NewFilePlanPropertyCitationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-FilePlanPropertyCitation", p.params())
}

// NewFilePlanPropertyDepartmentParams are the parameters of New-FilePlanPropertyDepartment.
type NewFilePlanPropertyDepartmentParams struct {
	Name string `ps:"Name"`
}

func (p NewFilePlanPropertyDepartmentParams) params() map[string]any {
	m := map[string]any{}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	return m
}

// NewFilePlanPropertyDepartment runs the New-FilePlanPropertyDepartment cmdlet.
func (s *Service) NewFilePlanPropertyDepartment(ctx context.Context, p NewFilePlanPropertyDepartmentParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-FilePlanPropertyDepartment", p.params())
}

// NewFilePlanPropertyReferenceIdParams are the parameters of New-FilePlanPropertyReferenceId.
type NewFilePlanPropertyReferenceIdParams struct {
	Name string `ps:"Name"`
}

func (p NewFilePlanPropertyReferenceIdParams) params() map[string]any {
	m := map[string]any{}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	return m
}

// NewFilePlanPropertyReferenceId runs the New-FilePlanPropertyReferenceId cmdlet.
func (s *Service) NewFilePlanPropertyReferenceId(ctx context.Context, p NewFilePlanPropertyReferenceIdParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-FilePlanPropertyReferenceId", p.params())
}

// NewFilePlanPropertySubCategoryParams are the parameters of New-FilePlanPropertySubCategory.
type NewFilePlanPropertySubCategoryParams struct {
	Name     string `ps:"Name"`
	ParentId any    `ps:"ParentId"`
}

func (p NewFilePlanPropertySubCategoryParams) params() map[string]any {
	m := map[string]any{}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.ParentId != nil {
		m["ParentId"] = p.ParentId
	}
	return m
}

// NewFilePlanPropertySubCategory runs the New-FilePlanPropertySubCategory cmdlet.
func (s *Service) NewFilePlanPropertySubCategory(ctx context.Context, p NewFilePlanPropertySubCategoryParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-FilePlanPropertySubCategory", p.params())
}

// NewGlobalListParams are the parameters of New-GlobalList.
// DefaultParameterSetName: Default
type NewGlobalListParams struct {
	Comment         string `ps:"Comment"`
	GlobalListType  any    `ps:"GlobalListType"`
	GlobalListValue any    `ps:"GlobalListValue"`
	Name            string `ps:"Name"`
	Organization    any    `ps:"Organization"`
}

func (p NewGlobalListParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.GlobalListType != nil {
		m["GlobalListType"] = p.GlobalListType
	}
	if p.GlobalListValue != nil {
		m["GlobalListValue"] = p.GlobalListValue
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Organization != nil {
		m["Organization"] = p.Organization
	}
	return m
}

// NewGlobalList runs the New-GlobalList cmdlet.
func (s *Service) NewGlobalList(ctx context.Context, p NewGlobalListParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-GlobalList", p.params())
}

// NewHoldCompliancePolicyParams are the parameters of New-HoldCompliancePolicy.
type NewHoldCompliancePolicyParams struct {
	Comment              string `ps:"Comment"`
	Enabled              bool   `ps:"Enabled"`
	ExchangeLocation     any    `ps:"ExchangeLocation"`
	Force                bool   `ps:"Force"`
	Name                 string `ps:"Name"`
	PublicFolderLocation any    `ps:"PublicFolderLocation"`
	SharePointLocation   any    `ps:"SharePointLocation"`
}

func (p NewHoldCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.ExchangeLocation != nil {
		m["ExchangeLocation"] = p.ExchangeLocation
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.PublicFolderLocation != nil {
		m["PublicFolderLocation"] = p.PublicFolderLocation
	}
	if p.SharePointLocation != nil {
		m["SharePointLocation"] = p.SharePointLocation
	}
	return m
}

// NewHoldCompliancePolicy runs the New-HoldCompliancePolicy cmdlet.
func (s *Service) NewHoldCompliancePolicy(ctx context.Context, p NewHoldCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-HoldCompliancePolicy", p.params())
}

// NewHoldComplianceRuleParams are the parameters of New-HoldComplianceRule.
type NewHoldComplianceRuleParams struct {
	Comment                 string `ps:"Comment"`
	ContentDateFrom         any    `ps:"ContentDateFrom"`
	ContentDateTo           any    `ps:"ContentDateTo"`
	ContentMatchQuery       string `ps:"ContentMatchQuery"`
	Disabled                bool   `ps:"Disabled"`
	HoldContent             any    `ps:"HoldContent"`
	HoldDurationDisplayHint any    `ps:"HoldDurationDisplayHint"`
	Name                    string `ps:"Name"`
	Policy                  any    `ps:"Policy"`
}

func (p NewHoldComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.ContentDateFrom != nil {
		m["ContentDateFrom"] = p.ContentDateFrom
	}
	if p.ContentDateTo != nil {
		m["ContentDateTo"] = p.ContentDateTo
	}
	if p.ContentMatchQuery != "" {
		m["ContentMatchQuery"] = p.ContentMatchQuery
	}
	if p.Disabled {
		m["Disabled"] = true
	}
	if p.HoldContent != nil {
		m["HoldContent"] = p.HoldContent
	}
	if p.HoldDurationDisplayHint != nil {
		m["HoldDurationDisplayHint"] = p.HoldDurationDisplayHint
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	return m
}

// NewHoldComplianceRule runs the New-HoldComplianceRule cmdlet.
func (s *Service) NewHoldComplianceRule(ctx context.Context, p NewHoldComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-HoldComplianceRule", p.params())
}

// NewInformationBarrierPolicyParams are the parameters of New-InformationBarrierPolicy.
// DefaultParameterSetName: InformationBarrierDefault
type NewInformationBarrierPolicyParams struct {
	AssignedSegment      string `ps:"AssignedSegment"`
	Comment              string `ps:"Comment"`
	Force                bool   `ps:"Force"`
	ModerationAllowed    bool   `ps:"ModerationAllowed"`
	Name                 string `ps:"Name"`
	SegmentAllowedFilter string `ps:"SegmentAllowedFilter"`
	SegmentsAllowed      any    `ps:"SegmentsAllowed"`
	SegmentsBlocked      any    `ps:"SegmentsBlocked"`
	State                any    `ps:"State"`
}

func (p NewInformationBarrierPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.AssignedSegment != "" {
		m["AssignedSegment"] = p.AssignedSegment
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Force {
		m["Force"] = true
	}
	if p.ModerationAllowed {
		m["ModerationAllowed"] = true
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.SegmentAllowedFilter != "" {
		m["SegmentAllowedFilter"] = p.SegmentAllowedFilter
	}
	if p.SegmentsAllowed != nil {
		m["SegmentsAllowed"] = p.SegmentsAllowed
	}
	if p.SegmentsBlocked != nil {
		m["SegmentsBlocked"] = p.SegmentsBlocked
	}
	if p.State != nil {
		m["State"] = p.State
	}
	return m
}

// NewInformationBarrierPolicy runs the New-InformationBarrierPolicy cmdlet.
func (s *Service) NewInformationBarrierPolicy(ctx context.Context, p NewInformationBarrierPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-InformationBarrierPolicy", p.params())
}

// NewInsiderRiskEntityListParams are the parameters of New-InsiderRiskEntityList.
type NewInsiderRiskEntityListParams struct {
	Description string `ps:"Description"`
	DisplayName string `ps:"DisplayName"`
	Entities    any    `ps:"Entities"`
	Name        string `ps:"Name"`
	Type        any    `ps:"Type"`
}

func (p NewInsiderRiskEntityListParams) params() map[string]any {
	m := map[string]any{}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.Entities != nil {
		m["Entities"] = p.Entities
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Type != nil {
		m["Type"] = p.Type
	}
	return m
}

// NewInsiderRiskEntityList runs the New-InsiderRiskEntityList cmdlet.
func (s *Service) NewInsiderRiskEntityList(ctx context.Context, p NewInsiderRiskEntityListParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-InsiderRiskEntityList", p.params())
}

// NewInsiderRiskPolicyParams are the parameters of New-InsiderRiskPolicy.
type NewInsiderRiskPolicyParams struct {
	CCPolicyName                   string `ps:"CCPolicyName"`
	CCPolicySdsId                  string `ps:"CCPolicySdsId"`
	Comment                        string `ps:"Comment"`
	CustomTags                     any    `ps:"CustomTags"`
	DlpPoliciesAsTrigger           any    `ps:"DlpPoliciesAsTrigger"`
	DlpPolicy                      any    `ps:"DlpPolicy"`
	DlpSensitiveTypes              any    `ps:"DlpSensitiveTypes"`
	Enabled                        bool   `ps:"Enabled"`
	ExchangeLocation               any    `ps:"ExchangeLocation"`
	ExchangeLocationException      any    `ps:"ExchangeLocationException"`
	ExtensibleIndicators           any    `ps:"ExtensibleIndicators"`
	ExtensibleTriggerInsightGroups any    `ps:"ExtensibleTriggerInsightGroups"`
	FileExtensions                 any    `ps:"FileExtensions"`
	FutureTerminationWindow        int    `ps:"FutureTerminationWindow"`
	HistoricTimeSpan               int    `ps:"HistoricTimeSpan"`
	Indicators                     any    `ps:"Indicators"`
	InScopeTimeSpan                int    `ps:"InScopeTimeSpan"`
	InsiderRiskScenario            any    `ps:"InsiderRiskScenario"`
	IrmAdaptiveScopeLocation       any    `ps:"IrmAdaptiveScopeLocation"`
	IsCustom                       bool   `ps:"IsCustom"`
	IsPriorityContentOnlyScoring   bool   `ps:"IsPriorityContentOnlyScoring"`
	MLClassifierTypes              any    `ps:"MLClassifierTypes"`
	ModernGroupLocation            any    `ps:"ModernGroupLocation"`
	ModernGroupLocationException   any    `ps:"ModernGroupLocationException"`
	Name                           string `ps:"Name"`
	OptInDrpForDlp                 bool   `ps:"OptInDrpForDlp"`
	PastTerminationWindow          int    `ps:"PastTerminationWindow"`
	PolicyRBACScopes               any    `ps:"PolicyRBACScopes"`
	PostTerminationActivity        bool   `ps:"PostTerminationActivity"`
	SchemaVersion                  int    `ps:"SchemaVersion"`
	SensitivityLabels              any    `ps:"SensitivityLabels"`
	SessionRecordingSettings       string `ps:"SessionRecordingSettings"`
	SharepointSites                any    `ps:"SharepointSites"`
	TeamsSites                     any    `ps:"TeamsSites"`
	TenantSetting                  string `ps:"TenantSetting"`
	TriggerInsightGroups           any    `ps:"TriggerInsightGroups"`
	Triggers                       any    `ps:"Triggers"`
	UseDefaultTemplate             int    `ps:"UseDefaultTemplate"`
	UserMailList                   any    `ps:"UserMailList"`
}

func (p NewInsiderRiskPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.CCPolicyName != "" {
		m["CCPolicyName"] = p.CCPolicyName
	}
	if p.CCPolicySdsId != "" {
		m["CCPolicySdsId"] = p.CCPolicySdsId
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.CustomTags != nil {
		m["CustomTags"] = p.CustomTags
	}
	if p.DlpPoliciesAsTrigger != nil {
		m["DlpPoliciesAsTrigger"] = p.DlpPoliciesAsTrigger
	}
	if p.DlpPolicy != nil {
		m["DlpPolicy"] = p.DlpPolicy
	}
	if p.DlpSensitiveTypes != nil {
		m["DlpSensitiveTypes"] = p.DlpSensitiveTypes
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.ExchangeLocation != nil {
		m["ExchangeLocation"] = p.ExchangeLocation
	}
	if p.ExchangeLocationException != nil {
		m["ExchangeLocationException"] = p.ExchangeLocationException
	}
	if p.ExtensibleIndicators != nil {
		m["ExtensibleIndicators"] = p.ExtensibleIndicators
	}
	if p.ExtensibleTriggerInsightGroups != nil {
		m["ExtensibleTriggerInsightGroups"] = p.ExtensibleTriggerInsightGroups
	}
	if p.FileExtensions != nil {
		m["FileExtensions"] = p.FileExtensions
	}
	if p.FutureTerminationWindow != 0 {
		m["FutureTerminationWindow"] = p.FutureTerminationWindow
	}
	if p.HistoricTimeSpan != 0 {
		m["HistoricTimeSpan"] = p.HistoricTimeSpan
	}
	if p.Indicators != nil {
		m["Indicators"] = p.Indicators
	}
	if p.InScopeTimeSpan != 0 {
		m["InScopeTimeSpan"] = p.InScopeTimeSpan
	}
	if p.InsiderRiskScenario != nil {
		m["InsiderRiskScenario"] = p.InsiderRiskScenario
	}
	if p.IrmAdaptiveScopeLocation != nil {
		m["IrmAdaptiveScopeLocation"] = p.IrmAdaptiveScopeLocation
	}
	if p.IsCustom {
		m["IsCustom"] = true
	}
	if p.IsPriorityContentOnlyScoring {
		m["IsPriorityContentOnlyScoring"] = true
	}
	if p.MLClassifierTypes != nil {
		m["MLClassifierTypes"] = p.MLClassifierTypes
	}
	if p.ModernGroupLocation != nil {
		m["ModernGroupLocation"] = p.ModernGroupLocation
	}
	if p.ModernGroupLocationException != nil {
		m["ModernGroupLocationException"] = p.ModernGroupLocationException
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.OptInDrpForDlp {
		m["OptInDrpForDlp"] = true
	}
	if p.PastTerminationWindow != 0 {
		m["PastTerminationWindow"] = p.PastTerminationWindow
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.PostTerminationActivity {
		m["PostTerminationActivity"] = true
	}
	if p.SchemaVersion != 0 {
		m["SchemaVersion"] = p.SchemaVersion
	}
	if p.SensitivityLabels != nil {
		m["SensitivityLabels"] = p.SensitivityLabels
	}
	if p.SessionRecordingSettings != "" {
		m["SessionRecordingSettings"] = p.SessionRecordingSettings
	}
	if p.SharepointSites != nil {
		m["SharepointSites"] = p.SharepointSites
	}
	if p.TeamsSites != nil {
		m["TeamsSites"] = p.TeamsSites
	}
	if p.TenantSetting != "" {
		m["TenantSetting"] = p.TenantSetting
	}
	if p.TriggerInsightGroups != nil {
		m["TriggerInsightGroups"] = p.TriggerInsightGroups
	}
	if p.Triggers != nil {
		m["Triggers"] = p.Triggers
	}
	if p.UseDefaultTemplate != 0 {
		m["UseDefaultTemplate"] = p.UseDefaultTemplate
	}
	if p.UserMailList != nil {
		m["UserMailList"] = p.UserMailList
	}
	return m
}

// NewInsiderRiskPolicy runs the New-InsiderRiskPolicy cmdlet.
func (s *Service) NewInsiderRiskPolicy(ctx context.Context, p NewInsiderRiskPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-InsiderRiskPolicy", p.params())
}

// NewJitConfigurationParams are the parameters of New-JitConfiguration.
type NewJitConfigurationParams struct {
	Comment                          string   `ps:"Comment"`
	EndpointConfig                   string   `ps:"EndpointConfig"`
	EndpointDlpLocation              any      `ps:"EndpointDlpLocation"`
	EndpointDlpLocationException     any      `ps:"EndpointDlpLocationException"`
	ExceptIfOneDriveSharedBy         []string `ps:"ExceptIfOneDriveSharedBy"`
	ExceptIfOneDriveSharedByMemberOf []string `ps:"ExceptIfOneDriveSharedByMemberOf"`
	Mode                             any      `ps:"Mode"`
	Name                             string   `ps:"Name"`
	OneDriveLocation                 any      `ps:"OneDriveLocation"`
	OneDriveLocationException        any      `ps:"OneDriveLocationException"`
	OneDriveSharedBy                 []string `ps:"OneDriveSharedBy"`
	OneDriveSharedByMemberOf         []string `ps:"OneDriveSharedByMemberOf"`
	Organization                     any      `ps:"Organization"`
	SharePointLocation               any      `ps:"SharePointLocation"`
	SharePointLocationException      any      `ps:"SharePointLocationException"`
}

func (p NewJitConfigurationParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.EndpointConfig != "" {
		m["EndpointConfig"] = p.EndpointConfig
	}
	if p.EndpointDlpLocation != nil {
		m["EndpointDlpLocation"] = p.EndpointDlpLocation
	}
	if p.EndpointDlpLocationException != nil {
		m["EndpointDlpLocationException"] = p.EndpointDlpLocationException
	}
	if len(p.ExceptIfOneDriveSharedBy) > 0 {
		m["ExceptIfOneDriveSharedBy"] = p.ExceptIfOneDriveSharedBy
	}
	if len(p.ExceptIfOneDriveSharedByMemberOf) > 0 {
		m["ExceptIfOneDriveSharedByMemberOf"] = p.ExceptIfOneDriveSharedByMemberOf
	}
	if p.Mode != nil {
		m["Mode"] = p.Mode
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.OneDriveLocation != nil {
		m["OneDriveLocation"] = p.OneDriveLocation
	}
	if p.OneDriveLocationException != nil {
		m["OneDriveLocationException"] = p.OneDriveLocationException
	}
	if len(p.OneDriveSharedBy) > 0 {
		m["OneDriveSharedBy"] = p.OneDriveSharedBy
	}
	if len(p.OneDriveSharedByMemberOf) > 0 {
		m["OneDriveSharedByMemberOf"] = p.OneDriveSharedByMemberOf
	}
	if p.Organization != nil {
		m["Organization"] = p.Organization
	}
	if p.SharePointLocation != nil {
		m["SharePointLocation"] = p.SharePointLocation
	}
	if p.SharePointLocationException != nil {
		m["SharePointLocationException"] = p.SharePointLocationException
	}
	return m
}

// NewJitConfiguration runs the New-JitConfiguration cmdlet.
func (s *Service) NewJitConfiguration(ctx context.Context, p NewJitConfigurationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-JitConfiguration", p.params())
}

// NewLabelParams are the parameters of New-Label.
type NewLabelParams struct {
	AbacEnabled                                           any    `ps:"AbacEnabled"`
	AdvancedSettings                                      any    `ps:"AdvancedSettings"`
	ApplyContentMarkingFooterAlignment                    any    `ps:"ApplyContentMarkingFooterAlignment"`
	ApplyContentMarkingFooterEnabled                      any    `ps:"ApplyContentMarkingFooterEnabled"`
	ApplyContentMarkingFooterFontColor                    string `ps:"ApplyContentMarkingFooterFontColor"`
	ApplyContentMarkingFooterFontName                     string `ps:"ApplyContentMarkingFooterFontName"`
	ApplyContentMarkingFooterFontSize                     any    `ps:"ApplyContentMarkingFooterFontSize"`
	ApplyContentMarkingFooterMargin                       any    `ps:"ApplyContentMarkingFooterMargin"`
	ApplyContentMarkingFooterText                         string `ps:"ApplyContentMarkingFooterText"`
	ApplyContentMarkingHeaderAlignment                    any    `ps:"ApplyContentMarkingHeaderAlignment"`
	ApplyContentMarkingHeaderEnabled                      any    `ps:"ApplyContentMarkingHeaderEnabled"`
	ApplyContentMarkingHeaderFontColor                    string `ps:"ApplyContentMarkingHeaderFontColor"`
	ApplyContentMarkingHeaderFontName                     string `ps:"ApplyContentMarkingHeaderFontName"`
	ApplyContentMarkingHeaderFontSize                     any    `ps:"ApplyContentMarkingHeaderFontSize"`
	ApplyContentMarkingHeaderMargin                       any    `ps:"ApplyContentMarkingHeaderMargin"`
	ApplyContentMarkingHeaderText                         string `ps:"ApplyContentMarkingHeaderText"`
	ApplyDynamicWatermarkingEnabled                       any    `ps:"ApplyDynamicWatermarkingEnabled"`
	ApplyMessageMetaData                                  any    `ps:"ApplyMessageMetaData"`
	ApplyMessageMetaDataEnabled                           any    `ps:"ApplyMessageMetaDataEnabled"`
	ApplyWaterMarkingEnabled                              any    `ps:"ApplyWaterMarkingEnabled"`
	ApplyWaterMarkingFontColor                            string `ps:"ApplyWaterMarkingFontColor"`
	ApplyWaterMarkingFontName                             string `ps:"ApplyWaterMarkingFontName"`
	ApplyWaterMarkingFontSize                             any    `ps:"ApplyWaterMarkingFontSize"`
	ApplyWaterMarkingLayout                               any    `ps:"ApplyWaterMarkingLayout"`
	ApplyWaterMarkingText                                 string `ps:"ApplyWaterMarkingText"`
	ColumnAssetCondition                                  string `ps:"ColumnAssetCondition"`
	Comment                                               string `ps:"Comment"`
	ConditionalAccessAuthContext                          string `ps:"ConditionalAccessAuthContext"`
	Conditions                                            any    `ps:"Conditions"`
	ContentType                                           any    `ps:"ContentType"`
	DefaultContentLabel                                   string `ps:"DefaultContentLabel"`
	DenyAccessEnabled                                     any    `ps:"DenyAccessEnabled"`
	DenyAccessExcludedPrincipals                          string `ps:"DenyAccessExcludedPrincipals"`
	DenyAccessIncludedPrincipals                          string `ps:"DenyAccessIncludedPrincipals"`
	DenyAccessRights                                      string `ps:"DenyAccessRights"`
	DenyAccessWorkload                                    string `ps:"DenyAccessWorkload"`
	DisplayName                                           string `ps:"DisplayName"`
	DynamicWatermarkDisplay                               string `ps:"DynamicWatermarkDisplay"`
	EncryptionAipTemplateScopes                           string `ps:"EncryptionAipTemplateScopes"`
	EncryptionContentExpiredOnDateInDaysOrNever           string `ps:"EncryptionContentExpiredOnDateInDaysOrNever"`
	EncryptionDisableAutomaticOwnerRights                 any    `ps:"EncryptionDisableAutomaticOwnerRights"`
	EncryptionDoNotForward                                any    `ps:"EncryptionDoNotForward"`
	EncryptionDoubleKeyEncryptionUrl                      string `ps:"EncryptionDoubleKeyEncryptionUrl"`
	EncryptionEnabled                                     any    `ps:"EncryptionEnabled"`
	EncryptionEncryptOnly                                 any    `ps:"EncryptionEncryptOnly"`
	EncryptionLinkedTemplateId                            string `ps:"EncryptionLinkedTemplateId"`
	EncryptionOfflineAccessDays                           any    `ps:"EncryptionOfflineAccessDays"`
	EncryptionPromptUser                                  any    `ps:"EncryptionPromptUser"`
	EncryptionProtectionType                              any    `ps:"EncryptionProtectionType"`
	EncryptionRightsDefinitions                           any    `ps:"EncryptionRightsDefinitions"`
	EncryptionRightsUrl                                   string `ps:"EncryptionRightsUrl"`
	EncryptionTemplateId                                  string `ps:"EncryptionTemplateId"`
	Identity                                              any    `ps:"Identity"`
	InformationProtectionAttributeRequired                any    `ps:"InformationProtectionAttributeRequired"`
	InheritToChildItems                                   any    `ps:"InheritToChildItems"`
	IsLabelGroup                                          bool   `ps:"IsLabelGroup"`
	LabelActions                                          any    `ps:"LabelActions"`
	LocaleSettings                                        any    `ps:"LocaleSettings"`
	MigrationId                                           string `ps:"MigrationId"`
	Name                                                  string `ps:"Name"`
	ParentId                                              any    `ps:"ParentId"`
	RuleBlob                                              string `ps:"RuleBlob"`
	SchematizedDataCondition                              string `ps:"SchematizedDataCondition"`
	Setting                                               any    `ps:"Setting"`
	Settings                                              any    `ps:"Settings"`
	SiteAndGroupProtectionAllowAccessToGuestUsers         any    `ps:"SiteAndGroupProtectionAllowAccessToGuestUsers"`
	SiteAndGroupProtectionAllowEmailFromGuestUsers        any    `ps:"SiteAndGroupProtectionAllowEmailFromGuestUsers"`
	SiteAndGroupProtectionAllowFullAccess                 any    `ps:"SiteAndGroupProtectionAllowFullAccess"`
	SiteAndGroupProtectionAllowLimitedAccess              any    `ps:"SiteAndGroupProtectionAllowLimitedAccess"`
	SiteAndGroupProtectionBlockAccess                     any    `ps:"SiteAndGroupProtectionBlockAccess"`
	SiteAndGroupProtectionEnabled                         any    `ps:"SiteAndGroupProtectionEnabled"`
	SiteAndGroupProtectionLevel                           any    `ps:"SiteAndGroupProtectionLevel"`
	SiteAndGroupProtectionPrivacy                         any    `ps:"SiteAndGroupProtectionPrivacy"`
	SiteExternalSharingControlType                        any    `ps:"SiteExternalSharingControlType"`
	TeamsAllowedPresenters                                any    `ps:"TeamsAllowedPresenters"`
	TeamsAllowMeetingChat                                 any    `ps:"TeamsAllowMeetingChat"`
	TeamsAllowPrivateTeamsToBeDiscoverableUsingSearch     any    `ps:"TeamsAllowPrivateTeamsToBeDiscoverableUsingSearch"`
	TeamsBypassLobbyForDialInUsers                        any    `ps:"TeamsBypassLobbyForDialInUsers"`
	TeamsChannelProtectionEnabled                         any    `ps:"TeamsChannelProtectionEnabled"`
	TeamsChannelSharedWithExternalTenants                 any    `ps:"TeamsChannelSharedWithExternalTenants"`
	TeamsChannelSharedWithPrivateTeamsOnly                any    `ps:"TeamsChannelSharedWithPrivateTeamsOnly"`
	TeamsChannelSharedWithSameLabelOnly                   any    `ps:"TeamsChannelSharedWithSameLabelOnly"`
	TeamsCopyRestrictionEnforced                          any    `ps:"TeamsCopyRestrictionEnforced"`
	TeamsDetectSensitiveContentDuringScreenSharingEnabled any    `ps:"TeamsDetectSensitiveContentDuringScreenSharingEnabled"`
	TeamsDisableLobby                                     any    `ps:"TeamsDisableLobby"`
	TeamsEndToEndEncryptionEnabled                        any    `ps:"TeamsEndToEndEncryptionEnabled"`
	TeamsLobbyBypassScope                                 any    `ps:"TeamsLobbyBypassScope"`
	TeamsLobbyRestrictionEnforced                         any    `ps:"TeamsLobbyRestrictionEnforced"`
	TeamsPresentersRestrictionEnforced                    any    `ps:"TeamsPresentersRestrictionEnforced"`
	TeamsProtectionEnabled                                any    `ps:"TeamsProtectionEnabled"`
	TeamsRecordAutomatically                              any    `ps:"TeamsRecordAutomatically"`
	TeamsVideoWatermark                                   any    `ps:"TeamsVideoWatermark"`
	TeamsWhoCanRecord                                     any    `ps:"TeamsWhoCanRecord"`
	Tooltip                                               string `ps:"Tooltip"`
}

func (p NewLabelParams) params() map[string]any {
	m := map[string]any{}
	if p.AbacEnabled != nil {
		m["AbacEnabled"] = p.AbacEnabled
	}
	if p.AdvancedSettings != nil {
		m["AdvancedSettings"] = p.AdvancedSettings
	}
	if p.ApplyContentMarkingFooterAlignment != nil {
		m["ApplyContentMarkingFooterAlignment"] = p.ApplyContentMarkingFooterAlignment
	}
	if p.ApplyContentMarkingFooterEnabled != nil {
		m["ApplyContentMarkingFooterEnabled"] = p.ApplyContentMarkingFooterEnabled
	}
	if p.ApplyContentMarkingFooterFontColor != "" {
		m["ApplyContentMarkingFooterFontColor"] = p.ApplyContentMarkingFooterFontColor
	}
	if p.ApplyContentMarkingFooterFontName != "" {
		m["ApplyContentMarkingFooterFontName"] = p.ApplyContentMarkingFooterFontName
	}
	if p.ApplyContentMarkingFooterFontSize != nil {
		m["ApplyContentMarkingFooterFontSize"] = p.ApplyContentMarkingFooterFontSize
	}
	if p.ApplyContentMarkingFooterMargin != nil {
		m["ApplyContentMarkingFooterMargin"] = p.ApplyContentMarkingFooterMargin
	}
	if p.ApplyContentMarkingFooterText != "" {
		m["ApplyContentMarkingFooterText"] = p.ApplyContentMarkingFooterText
	}
	if p.ApplyContentMarkingHeaderAlignment != nil {
		m["ApplyContentMarkingHeaderAlignment"] = p.ApplyContentMarkingHeaderAlignment
	}
	if p.ApplyContentMarkingHeaderEnabled != nil {
		m["ApplyContentMarkingHeaderEnabled"] = p.ApplyContentMarkingHeaderEnabled
	}
	if p.ApplyContentMarkingHeaderFontColor != "" {
		m["ApplyContentMarkingHeaderFontColor"] = p.ApplyContentMarkingHeaderFontColor
	}
	if p.ApplyContentMarkingHeaderFontName != "" {
		m["ApplyContentMarkingHeaderFontName"] = p.ApplyContentMarkingHeaderFontName
	}
	if p.ApplyContentMarkingHeaderFontSize != nil {
		m["ApplyContentMarkingHeaderFontSize"] = p.ApplyContentMarkingHeaderFontSize
	}
	if p.ApplyContentMarkingHeaderMargin != nil {
		m["ApplyContentMarkingHeaderMargin"] = p.ApplyContentMarkingHeaderMargin
	}
	if p.ApplyContentMarkingHeaderText != "" {
		m["ApplyContentMarkingHeaderText"] = p.ApplyContentMarkingHeaderText
	}
	if p.ApplyDynamicWatermarkingEnabled != nil {
		m["ApplyDynamicWatermarkingEnabled"] = p.ApplyDynamicWatermarkingEnabled
	}
	if p.ApplyMessageMetaData != nil {
		m["ApplyMessageMetaData"] = p.ApplyMessageMetaData
	}
	if p.ApplyMessageMetaDataEnabled != nil {
		m["ApplyMessageMetaDataEnabled"] = p.ApplyMessageMetaDataEnabled
	}
	if p.ApplyWaterMarkingEnabled != nil {
		m["ApplyWaterMarkingEnabled"] = p.ApplyWaterMarkingEnabled
	}
	if p.ApplyWaterMarkingFontColor != "" {
		m["ApplyWaterMarkingFontColor"] = p.ApplyWaterMarkingFontColor
	}
	if p.ApplyWaterMarkingFontName != "" {
		m["ApplyWaterMarkingFontName"] = p.ApplyWaterMarkingFontName
	}
	if p.ApplyWaterMarkingFontSize != nil {
		m["ApplyWaterMarkingFontSize"] = p.ApplyWaterMarkingFontSize
	}
	if p.ApplyWaterMarkingLayout != nil {
		m["ApplyWaterMarkingLayout"] = p.ApplyWaterMarkingLayout
	}
	if p.ApplyWaterMarkingText != "" {
		m["ApplyWaterMarkingText"] = p.ApplyWaterMarkingText
	}
	if p.ColumnAssetCondition != "" {
		m["ColumnAssetCondition"] = p.ColumnAssetCondition
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.ConditionalAccessAuthContext != "" {
		m["ConditionalAccessAuthContext"] = p.ConditionalAccessAuthContext
	}
	if p.Conditions != nil {
		m["Conditions"] = p.Conditions
	}
	if p.ContentType != nil {
		m["ContentType"] = p.ContentType
	}
	if p.DefaultContentLabel != "" {
		m["DefaultContentLabel"] = p.DefaultContentLabel
	}
	if p.DenyAccessEnabled != nil {
		m["DenyAccessEnabled"] = p.DenyAccessEnabled
	}
	if p.DenyAccessExcludedPrincipals != "" {
		m["DenyAccessExcludedPrincipals"] = p.DenyAccessExcludedPrincipals
	}
	if p.DenyAccessIncludedPrincipals != "" {
		m["DenyAccessIncludedPrincipals"] = p.DenyAccessIncludedPrincipals
	}
	if p.DenyAccessRights != "" {
		m["DenyAccessRights"] = p.DenyAccessRights
	}
	if p.DenyAccessWorkload != "" {
		m["DenyAccessWorkload"] = p.DenyAccessWorkload
	}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.DynamicWatermarkDisplay != "" {
		m["DynamicWatermarkDisplay"] = p.DynamicWatermarkDisplay
	}
	if p.EncryptionAipTemplateScopes != "" {
		m["EncryptionAipTemplateScopes"] = p.EncryptionAipTemplateScopes
	}
	if p.EncryptionContentExpiredOnDateInDaysOrNever != "" {
		m["EncryptionContentExpiredOnDateInDaysOrNever"] = p.EncryptionContentExpiredOnDateInDaysOrNever
	}
	if p.EncryptionDisableAutomaticOwnerRights != nil {
		m["EncryptionDisableAutomaticOwnerRights"] = p.EncryptionDisableAutomaticOwnerRights
	}
	if p.EncryptionDoNotForward != nil {
		m["EncryptionDoNotForward"] = p.EncryptionDoNotForward
	}
	if p.EncryptionDoubleKeyEncryptionUrl != "" {
		m["EncryptionDoubleKeyEncryptionUrl"] = p.EncryptionDoubleKeyEncryptionUrl
	}
	if p.EncryptionEnabled != nil {
		m["EncryptionEnabled"] = p.EncryptionEnabled
	}
	if p.EncryptionEncryptOnly != nil {
		m["EncryptionEncryptOnly"] = p.EncryptionEncryptOnly
	}
	if p.EncryptionLinkedTemplateId != "" {
		m["EncryptionLinkedTemplateId"] = p.EncryptionLinkedTemplateId
	}
	if p.EncryptionOfflineAccessDays != nil {
		m["EncryptionOfflineAccessDays"] = p.EncryptionOfflineAccessDays
	}
	if p.EncryptionPromptUser != nil {
		m["EncryptionPromptUser"] = p.EncryptionPromptUser
	}
	if p.EncryptionProtectionType != nil {
		m["EncryptionProtectionType"] = p.EncryptionProtectionType
	}
	if p.EncryptionRightsDefinitions != nil {
		m["EncryptionRightsDefinitions"] = p.EncryptionRightsDefinitions
	}
	if p.EncryptionRightsUrl != "" {
		m["EncryptionRightsUrl"] = p.EncryptionRightsUrl
	}
	if p.EncryptionTemplateId != "" {
		m["EncryptionTemplateId"] = p.EncryptionTemplateId
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.InformationProtectionAttributeRequired != nil {
		m["InformationProtectionAttributeRequired"] = p.InformationProtectionAttributeRequired
	}
	if p.InheritToChildItems != nil {
		m["InheritToChildItems"] = p.InheritToChildItems
	}
	if p.IsLabelGroup {
		m["IsLabelGroup"] = true
	}
	if p.LabelActions != nil {
		m["LabelActions"] = p.LabelActions
	}
	if p.LocaleSettings != nil {
		m["LocaleSettings"] = p.LocaleSettings
	}
	if p.MigrationId != "" {
		m["MigrationId"] = p.MigrationId
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.ParentId != nil {
		m["ParentId"] = p.ParentId
	}
	if p.RuleBlob != "" {
		m["RuleBlob"] = p.RuleBlob
	}
	if p.SchematizedDataCondition != "" {
		m["SchematizedDataCondition"] = p.SchematizedDataCondition
	}
	if p.Setting != nil {
		m["Setting"] = p.Setting
	}
	if p.Settings != nil {
		m["Settings"] = p.Settings
	}
	if p.SiteAndGroupProtectionAllowAccessToGuestUsers != nil {
		m["SiteAndGroupProtectionAllowAccessToGuestUsers"] = p.SiteAndGroupProtectionAllowAccessToGuestUsers
	}
	if p.SiteAndGroupProtectionAllowEmailFromGuestUsers != nil {
		m["SiteAndGroupProtectionAllowEmailFromGuestUsers"] = p.SiteAndGroupProtectionAllowEmailFromGuestUsers
	}
	if p.SiteAndGroupProtectionAllowFullAccess != nil {
		m["SiteAndGroupProtectionAllowFullAccess"] = p.SiteAndGroupProtectionAllowFullAccess
	}
	if p.SiteAndGroupProtectionAllowLimitedAccess != nil {
		m["SiteAndGroupProtectionAllowLimitedAccess"] = p.SiteAndGroupProtectionAllowLimitedAccess
	}
	if p.SiteAndGroupProtectionBlockAccess != nil {
		m["SiteAndGroupProtectionBlockAccess"] = p.SiteAndGroupProtectionBlockAccess
	}
	if p.SiteAndGroupProtectionEnabled != nil {
		m["SiteAndGroupProtectionEnabled"] = p.SiteAndGroupProtectionEnabled
	}
	if p.SiteAndGroupProtectionLevel != nil {
		m["SiteAndGroupProtectionLevel"] = p.SiteAndGroupProtectionLevel
	}
	if p.SiteAndGroupProtectionPrivacy != nil {
		m["SiteAndGroupProtectionPrivacy"] = p.SiteAndGroupProtectionPrivacy
	}
	if p.SiteExternalSharingControlType != nil {
		m["SiteExternalSharingControlType"] = p.SiteExternalSharingControlType
	}
	if p.TeamsAllowedPresenters != nil {
		m["TeamsAllowedPresenters"] = p.TeamsAllowedPresenters
	}
	if p.TeamsAllowMeetingChat != nil {
		m["TeamsAllowMeetingChat"] = p.TeamsAllowMeetingChat
	}
	if p.TeamsAllowPrivateTeamsToBeDiscoverableUsingSearch != nil {
		m["TeamsAllowPrivateTeamsToBeDiscoverableUsingSearch"] = p.TeamsAllowPrivateTeamsToBeDiscoverableUsingSearch
	}
	if p.TeamsBypassLobbyForDialInUsers != nil {
		m["TeamsBypassLobbyForDialInUsers"] = p.TeamsBypassLobbyForDialInUsers
	}
	if p.TeamsChannelProtectionEnabled != nil {
		m["TeamsChannelProtectionEnabled"] = p.TeamsChannelProtectionEnabled
	}
	if p.TeamsChannelSharedWithExternalTenants != nil {
		m["TeamsChannelSharedWithExternalTenants"] = p.TeamsChannelSharedWithExternalTenants
	}
	if p.TeamsChannelSharedWithPrivateTeamsOnly != nil {
		m["TeamsChannelSharedWithPrivateTeamsOnly"] = p.TeamsChannelSharedWithPrivateTeamsOnly
	}
	if p.TeamsChannelSharedWithSameLabelOnly != nil {
		m["TeamsChannelSharedWithSameLabelOnly"] = p.TeamsChannelSharedWithSameLabelOnly
	}
	if p.TeamsCopyRestrictionEnforced != nil {
		m["TeamsCopyRestrictionEnforced"] = p.TeamsCopyRestrictionEnforced
	}
	if p.TeamsDetectSensitiveContentDuringScreenSharingEnabled != nil {
		m["TeamsDetectSensitiveContentDuringScreenSharingEnabled"] = p.TeamsDetectSensitiveContentDuringScreenSharingEnabled
	}
	if p.TeamsDisableLobby != nil {
		m["TeamsDisableLobby"] = p.TeamsDisableLobby
	}
	if p.TeamsEndToEndEncryptionEnabled != nil {
		m["TeamsEndToEndEncryptionEnabled"] = p.TeamsEndToEndEncryptionEnabled
	}
	if p.TeamsLobbyBypassScope != nil {
		m["TeamsLobbyBypassScope"] = p.TeamsLobbyBypassScope
	}
	if p.TeamsLobbyRestrictionEnforced != nil {
		m["TeamsLobbyRestrictionEnforced"] = p.TeamsLobbyRestrictionEnforced
	}
	if p.TeamsPresentersRestrictionEnforced != nil {
		m["TeamsPresentersRestrictionEnforced"] = p.TeamsPresentersRestrictionEnforced
	}
	if p.TeamsProtectionEnabled != nil {
		m["TeamsProtectionEnabled"] = p.TeamsProtectionEnabled
	}
	if p.TeamsRecordAutomatically != nil {
		m["TeamsRecordAutomatically"] = p.TeamsRecordAutomatically
	}
	if p.TeamsVideoWatermark != nil {
		m["TeamsVideoWatermark"] = p.TeamsVideoWatermark
	}
	if p.TeamsWhoCanRecord != nil {
		m["TeamsWhoCanRecord"] = p.TeamsWhoCanRecord
	}
	if p.Tooltip != "" {
		m["Tooltip"] = p.Tooltip
	}
	return m
}

// NewLabel runs the New-Label cmdlet.
func (s *Service) NewLabel(ctx context.Context, p NewLabelParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-Label", p.params())
}

// NewLabelExplorerConfigParams are the parameters of New-LabelExplorerConfig.
type NewLabelExplorerConfigParams struct {
	Actions     any    `ps:"Actions"`
	DisplayName string `ps:"DisplayName"`
	Filter      string `ps:"Filter"`
	Labels      any    `ps:"Labels"`
	Name        string `ps:"Name"`
	Users       any    `ps:"Users"`
}

func (p NewLabelExplorerConfigParams) params() map[string]any {
	m := map[string]any{}
	if p.Actions != nil {
		m["Actions"] = p.Actions
	}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.Filter != "" {
		m["Filter"] = p.Filter
	}
	if p.Labels != nil {
		m["Labels"] = p.Labels
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Users != nil {
		m["Users"] = p.Users
	}
	return m
}

// NewLabelExplorerConfig runs the New-LabelExplorerConfig cmdlet.
func (s *Service) NewLabelExplorerConfig(ctx context.Context, p NewLabelExplorerConfigParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-LabelExplorerConfig", p.params())
}

// NewLabelPolicyParams are the parameters of New-LabelPolicy.
type NewLabelPolicyParams struct {
	AdvancedSettings                any    `ps:"AdvancedSettings"`
	Comment                         string `ps:"Comment"`
	ExchangeAdaptiveScopes          any    `ps:"ExchangeAdaptiveScopes"`
	ExchangeAdaptiveScopesException any    `ps:"ExchangeAdaptiveScopesException"`
	ExchangeLocation                any    `ps:"ExchangeLocation"`
	ExchangeLocationException       any    `ps:"ExchangeLocationException"`
	Force                           bool   `ps:"Force"`
	Labels                          any    `ps:"Labels"`
	Locations                       string `ps:"Locations"`
	MigrationId                     string `ps:"MigrationId"`
	ModernGroupLocation             any    `ps:"ModernGroupLocation"`
	ModernGroupLocationException    any    `ps:"ModernGroupLocationException"`
	Name                            string `ps:"Name"`
	OneDriveLocation                any    `ps:"OneDriveLocation"`
	OneDriveLocationException       any    `ps:"OneDriveLocationException"`
	PolicyRBACScopes                any    `ps:"PolicyRBACScopes"`
	PublicFolderLocation            any    `ps:"PublicFolderLocation"`
	Setting                         any    `ps:"Setting"`
	Settings                        any    `ps:"Settings"`
	SharePointLocation              any    `ps:"SharePointLocation"`
	SharePointLocationException     any    `ps:"SharePointLocationException"`
	SkypeLocation                   any    `ps:"SkypeLocation"`
	SkypeLocationException          any    `ps:"SkypeLocationException"`
}

func (p NewLabelPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.AdvancedSettings != nil {
		m["AdvancedSettings"] = p.AdvancedSettings
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.ExchangeAdaptiveScopes != nil {
		m["ExchangeAdaptiveScopes"] = p.ExchangeAdaptiveScopes
	}
	if p.ExchangeAdaptiveScopesException != nil {
		m["ExchangeAdaptiveScopesException"] = p.ExchangeAdaptiveScopesException
	}
	if p.ExchangeLocation != nil {
		m["ExchangeLocation"] = p.ExchangeLocation
	}
	if p.ExchangeLocationException != nil {
		m["ExchangeLocationException"] = p.ExchangeLocationException
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Labels != nil {
		m["Labels"] = p.Labels
	}
	if p.Locations != "" {
		m["Locations"] = p.Locations
	}
	if p.MigrationId != "" {
		m["MigrationId"] = p.MigrationId
	}
	if p.ModernGroupLocation != nil {
		m["ModernGroupLocation"] = p.ModernGroupLocation
	}
	if p.ModernGroupLocationException != nil {
		m["ModernGroupLocationException"] = p.ModernGroupLocationException
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.OneDriveLocation != nil {
		m["OneDriveLocation"] = p.OneDriveLocation
	}
	if p.OneDriveLocationException != nil {
		m["OneDriveLocationException"] = p.OneDriveLocationException
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.PublicFolderLocation != nil {
		m["PublicFolderLocation"] = p.PublicFolderLocation
	}
	if p.Setting != nil {
		m["Setting"] = p.Setting
	}
	if p.Settings != nil {
		m["Settings"] = p.Settings
	}
	if p.SharePointLocation != nil {
		m["SharePointLocation"] = p.SharePointLocation
	}
	if p.SharePointLocationException != nil {
		m["SharePointLocationException"] = p.SharePointLocationException
	}
	if p.SkypeLocation != nil {
		m["SkypeLocation"] = p.SkypeLocation
	}
	if p.SkypeLocationException != nil {
		m["SkypeLocationException"] = p.SkypeLocationException
	}
	return m
}

// NewLabelPolicy runs the New-LabelPolicy cmdlet.
func (s *Service) NewLabelPolicy(ctx context.Context, p NewLabelPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-LabelPolicy", p.params())
}

// NewMachineAssistedTagResourceParams are the parameters of New-MachineAssistedTagResource.
type NewMachineAssistedTagResourceParams struct {
	Description  string `ps:"Description"`
	ModelIds     any    `ps:"ModelIds"`
	Name         string `ps:"Name"`
	Resource     any    `ps:"Resource"`
	TaskId       any    `ps:"TaskId"`
	TaskType     any    `ps:"TaskType"`
	TrainingId   any    `ps:"TrainingId"`
	TrainingSite string `ps:"TrainingSite"`
}

func (p NewMachineAssistedTagResourceParams) params() map[string]any {
	m := map[string]any{}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if p.ModelIds != nil {
		m["ModelIds"] = p.ModelIds
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Resource != nil {
		m["Resource"] = p.Resource
	}
	if p.TaskId != nil {
		m["TaskId"] = p.TaskId
	}
	if p.TaskType != nil {
		m["TaskType"] = p.TaskType
	}
	if p.TrainingId != nil {
		m["TrainingId"] = p.TrainingId
	}
	if p.TrainingSite != "" {
		m["TrainingSite"] = p.TrainingSite
	}
	return m
}

// NewMachineAssistedTagResource runs the New-MachineAssistedTagResource cmdlet.
func (s *Service) NewMachineAssistedTagResource(ctx context.Context, p NewMachineAssistedTagResourceParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-MachineAssistedTagResource", p.params())
}

// NewOcrConfigurationParams are the parameters of New-OcrConfiguration.
type NewOcrConfigurationParams struct {
	Comment                              string   `ps:"Comment"`
	EndpointDlpLocation                  any      `ps:"EndpointDlpLocation"`
	EndpointDlpLocationException         any      `ps:"EndpointDlpLocationException"`
	EnforcementPlanes                    any      `ps:"EnforcementPlanes"`
	ExceptIfOneDriveSharedBy             []string `ps:"ExceptIfOneDriveSharedBy"`
	ExceptIfOneDriveSharedByMemberOf     []string `ps:"ExceptIfOneDriveSharedByMemberOf"`
	ExchangeLocation                     any      `ps:"ExchangeLocation"`
	ExchangeLocationException            any      `ps:"ExchangeLocationException"`
	ExchangeScopeOcrAnyRecipientExternal bool     `ps:"ExchangeScopeOcrAnyRecipientExternal"`
	ExchangeSender                       []string `ps:"ExchangeSender"`
	ExchangeSenderException              []string `ps:"ExchangeSenderException"`
	ExchangeSenderMemberOf               []string `ps:"ExchangeSenderMemberOf"`
	ExchangeSenderMemberOfException      []string `ps:"ExchangeSenderMemberOfException"`
	Locations                            string   `ps:"Locations"`
	Mode                                 any      `ps:"Mode"`
	Name                                 string   `ps:"Name"`
	OcrMode                              any      `ps:"OcrMode"`
	OneDriveLocation                     any      `ps:"OneDriveLocation"`
	OneDriveLocationException            any      `ps:"OneDriveLocationException"`
	OneDriveSharedBy                     []string `ps:"OneDriveSharedBy"`
	OneDriveSharedByMemberOf             []string `ps:"OneDriveSharedByMemberOf"`
	SharePointLocation                   any      `ps:"SharePointLocation"`
	SharePointLocationException          any      `ps:"SharePointLocationException"`
	TeamsLocation                        any      `ps:"TeamsLocation"`
	TeamsLocationException               any      `ps:"TeamsLocationException"`
}

func (p NewOcrConfigurationParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.EndpointDlpLocation != nil {
		m["EndpointDlpLocation"] = p.EndpointDlpLocation
	}
	if p.EndpointDlpLocationException != nil {
		m["EndpointDlpLocationException"] = p.EndpointDlpLocationException
	}
	if p.EnforcementPlanes != nil {
		m["EnforcementPlanes"] = p.EnforcementPlanes
	}
	if len(p.ExceptIfOneDriveSharedBy) > 0 {
		m["ExceptIfOneDriveSharedBy"] = p.ExceptIfOneDriveSharedBy
	}
	if len(p.ExceptIfOneDriveSharedByMemberOf) > 0 {
		m["ExceptIfOneDriveSharedByMemberOf"] = p.ExceptIfOneDriveSharedByMemberOf
	}
	if p.ExchangeLocation != nil {
		m["ExchangeLocation"] = p.ExchangeLocation
	}
	if p.ExchangeLocationException != nil {
		m["ExchangeLocationException"] = p.ExchangeLocationException
	}
	if p.ExchangeScopeOcrAnyRecipientExternal {
		m["ExchangeScopeOcrAnyRecipientExternal"] = true
	}
	if len(p.ExchangeSender) > 0 {
		m["ExchangeSender"] = p.ExchangeSender
	}
	if len(p.ExchangeSenderException) > 0 {
		m["ExchangeSenderException"] = p.ExchangeSenderException
	}
	if len(p.ExchangeSenderMemberOf) > 0 {
		m["ExchangeSenderMemberOf"] = p.ExchangeSenderMemberOf
	}
	if len(p.ExchangeSenderMemberOfException) > 0 {
		m["ExchangeSenderMemberOfException"] = p.ExchangeSenderMemberOfException
	}
	if p.Locations != "" {
		m["Locations"] = p.Locations
	}
	if p.Mode != nil {
		m["Mode"] = p.Mode
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.OcrMode != nil {
		m["OcrMode"] = p.OcrMode
	}
	if p.OneDriveLocation != nil {
		m["OneDriveLocation"] = p.OneDriveLocation
	}
	if p.OneDriveLocationException != nil {
		m["OneDriveLocationException"] = p.OneDriveLocationException
	}
	if len(p.OneDriveSharedBy) > 0 {
		m["OneDriveSharedBy"] = p.OneDriveSharedBy
	}
	if len(p.OneDriveSharedByMemberOf) > 0 {
		m["OneDriveSharedByMemberOf"] = p.OneDriveSharedByMemberOf
	}
	if p.SharePointLocation != nil {
		m["SharePointLocation"] = p.SharePointLocation
	}
	if p.SharePointLocationException != nil {
		m["SharePointLocationException"] = p.SharePointLocationException
	}
	if p.TeamsLocation != nil {
		m["TeamsLocation"] = p.TeamsLocation
	}
	if p.TeamsLocationException != nil {
		m["TeamsLocationException"] = p.TeamsLocationException
	}
	return m
}

// NewOcrConfiguration runs the New-OcrConfiguration cmdlet.
func (s *Service) NewOcrConfiguration(ctx context.Context, p NewOcrConfigurationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-OcrConfiguration", p.params())
}

// NewOrganizationSegmentParams are the parameters of New-OrganizationSegment.
// DefaultParameterSetName: OrganizationSegmentsDefault
type NewOrganizationSegmentParams struct {
	Name            string `ps:"Name"`
	UserGroupFilter string `ps:"UserGroupFilter"`
}

func (p NewOrganizationSegmentParams) params() map[string]any {
	m := map[string]any{}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.UserGroupFilter != "" {
		m["UserGroupFilter"] = p.UserGroupFilter
	}
	return m
}

// NewOrganizationSegment runs the New-OrganizationSegment cmdlet.
func (s *Service) NewOrganizationSegment(ctx context.Context, p NewOrganizationSegmentParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-OrganizationSegment", p.params())
}

// NewPrivacyManagementComplianceTagParams are the parameters of New-PrivacyManagementComplianceTag.
type NewPrivacyManagementComplianceTagParams struct {
	Comment           string `ps:"Comment"`
	Force             bool   `ps:"Force"`
	Name              string `ps:"Name"`
	RetentionAction   string `ps:"RetentionAction"`
	RetentionDuration any    `ps:"RetentionDuration"`
}

func (p NewPrivacyManagementComplianceTagParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.RetentionAction != "" {
		m["RetentionAction"] = p.RetentionAction
	}
	if p.RetentionDuration != nil {
		m["RetentionDuration"] = p.RetentionDuration
	}
	return m
}

// NewPrivacyManagementComplianceTag runs the New-PrivacyManagementComplianceTag cmdlet.
func (s *Service) NewPrivacyManagementComplianceTag(ctx context.Context, p NewPrivacyManagementComplianceTagParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-PrivacyManagementComplianceTag", p.params())
}

// NewPrivacyManagementPolicyParams are the parameters of New-PrivacyManagementPolicy.
// DefaultParameterSetName: Identity
type NewPrivacyManagementPolicyParams struct {
	Comment                     string `ps:"Comment"`
	Enabled                     bool   `ps:"Enabled"`
	ExchangeLocation            any    `ps:"ExchangeLocation"`
	ExchangeLocationException   any    `ps:"ExchangeLocationException"`
	Mode                        any    `ps:"Mode"`
	Name                        string `ps:"Name"`
	OneDriveLocation            any    `ps:"OneDriveLocation"`
	OneDriveLocationException   any    `ps:"OneDriveLocationException"`
	PrivacyManagementScenario   any    `ps:"PrivacyManagementScenario"`
	SharePointLocation          any    `ps:"SharePointLocation"`
	SharePointLocationException any    `ps:"SharePointLocationException"`
	TeamsLocation               any    `ps:"TeamsLocation"`
	TeamsLocationException      any    `ps:"TeamsLocationException"`
}

func (p NewPrivacyManagementPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.ExchangeLocation != nil {
		m["ExchangeLocation"] = p.ExchangeLocation
	}
	if p.ExchangeLocationException != nil {
		m["ExchangeLocationException"] = p.ExchangeLocationException
	}
	if p.Mode != nil {
		m["Mode"] = p.Mode
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.OneDriveLocation != nil {
		m["OneDriveLocation"] = p.OneDriveLocation
	}
	if p.OneDriveLocationException != nil {
		m["OneDriveLocationException"] = p.OneDriveLocationException
	}
	if p.PrivacyManagementScenario != nil {
		m["PrivacyManagementScenario"] = p.PrivacyManagementScenario
	}
	if p.SharePointLocation != nil {
		m["SharePointLocation"] = p.SharePointLocation
	}
	if p.SharePointLocationException != nil {
		m["SharePointLocationException"] = p.SharePointLocationException
	}
	if p.TeamsLocation != nil {
		m["TeamsLocation"] = p.TeamsLocation
	}
	if p.TeamsLocationException != nil {
		m["TeamsLocationException"] = p.TeamsLocationException
	}
	return m
}

// NewPrivacyManagementPolicy runs the New-PrivacyManagementPolicy cmdlet.
func (s *Service) NewPrivacyManagementPolicy(ctx context.Context, p NewPrivacyManagementPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-PrivacyManagementPolicy", p.params())
}

// NewPrivacyManagementRuleParams are the parameters of New-PrivacyManagementRule.
// DefaultParameterSetName: Identity
type NewPrivacyManagementRuleParams struct {
	Comment                             string   `ps:"Comment"`
	ContentContainsSensitiveInformation []string `ps:"ContentContainsSensitiveInformation"`
	CrossBoundaryTransfers              []string `ps:"CrossBoundaryTransfers"`
	EndpointOperations                  []string `ps:"EndpointOperations"`
	ExchangeSites                       any      `ps:"ExchangeSites"`
	Groups                              any      `ps:"Groups"`
	LastModifiedThresholdInDays         any      `ps:"LastModifiedThresholdInDays"`
	Locations                           string   `ps:"Locations"`
	Name                                string   `ps:"Name"`
	OnedriveSites                       any      `ps:"OnedriveSites"`
	Policy                              any      `ps:"Policy"`
	PolicySettings                      string   `ps:"PolicySettings"`
	PrivacyAccessLevel                  []string `ps:"PrivacyAccessLevel"`
	PurviewSites                        any      `ps:"PurviewSites"`
	SettingsVersion                     string   `ps:"SettingsVersion"`
	SharepointSiteOversharingEnabled    any      `ps:"SharepointSiteOversharingEnabled"`
	SharepointSites                     any      `ps:"SharepointSites"`
	TeamsSites                          any      `ps:"TeamsSites"`
	TenantSettings                      string   `ps:"TenantSettings"`
	Users                               any      `ps:"Users"`
}

func (p NewPrivacyManagementRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if len(p.ContentContainsSensitiveInformation) > 0 {
		m["ContentContainsSensitiveInformation"] = p.ContentContainsSensitiveInformation
	}
	if len(p.CrossBoundaryTransfers) > 0 {
		m["CrossBoundaryTransfers"] = p.CrossBoundaryTransfers
	}
	if len(p.EndpointOperations) > 0 {
		m["EndpointOperations"] = p.EndpointOperations
	}
	if p.ExchangeSites != nil {
		m["ExchangeSites"] = p.ExchangeSites
	}
	if p.Groups != nil {
		m["Groups"] = p.Groups
	}
	if p.LastModifiedThresholdInDays != nil {
		m["LastModifiedThresholdInDays"] = p.LastModifiedThresholdInDays
	}
	if p.Locations != "" {
		m["Locations"] = p.Locations
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.OnedriveSites != nil {
		m["OnedriveSites"] = p.OnedriveSites
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	if p.PolicySettings != "" {
		m["PolicySettings"] = p.PolicySettings
	}
	if len(p.PrivacyAccessLevel) > 0 {
		m["PrivacyAccessLevel"] = p.PrivacyAccessLevel
	}
	if p.PurviewSites != nil {
		m["PurviewSites"] = p.PurviewSites
	}
	if p.SettingsVersion != "" {
		m["SettingsVersion"] = p.SettingsVersion
	}
	if p.SharepointSiteOversharingEnabled != nil {
		m["SharepointSiteOversharingEnabled"] = p.SharepointSiteOversharingEnabled
	}
	if p.SharepointSites != nil {
		m["SharepointSites"] = p.SharepointSites
	}
	if p.TeamsSites != nil {
		m["TeamsSites"] = p.TeamsSites
	}
	if p.TenantSettings != "" {
		m["TenantSettings"] = p.TenantSettings
	}
	if p.Users != nil {
		m["Users"] = p.Users
	}
	return m
}

// NewPrivacyManagementRule runs the New-PrivacyManagementRule cmdlet.
func (s *Service) NewPrivacyManagementRule(ctx context.Context, p NewPrivacyManagementRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-PrivacyManagementRule", p.params())
}

// NewProcessComplianceReviewRecordsParams are the parameters of New-ProcessComplianceReviewRecords.
// DefaultParameterSetName: Identity
type NewProcessComplianceReviewRecordsParams struct {
	RecordReviewDecisionPayload string `ps:"RecordReviewDecisionPayload"`
}

func (p NewProcessComplianceReviewRecordsParams) params() map[string]any {
	m := map[string]any{}
	if p.RecordReviewDecisionPayload != "" {
		m["RecordReviewDecisionPayload"] = p.RecordReviewDecisionPayload
	}
	return m
}

// NewProcessComplianceReviewRecords runs the New-ProcessComplianceReviewRecords cmdlet.
func (s *Service) NewProcessComplianceReviewRecords(ctx context.Context, p NewProcessComplianceReviewRecordsParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-ProcessComplianceReviewRecords", p.params())
}

// NewProtectionAlertParams are the parameters of New-ProtectionAlert.
// DefaultParameterSetName: Identity
type NewProtectionAlertParams struct {
	AggregationType                                             any    `ps:"AggregationType"`
	AlertBy                                                     any    `ps:"AlertBy"`
	AlertFor                                                    any    `ps:"AlertFor"`
	Category                                                    any    `ps:"Category"`
	Comment                                                     string `ps:"Comment"`
	CorrelationPolicyId                                         any    `ps:"CorrelationPolicyId"`
	CustomProperties                                            any    `ps:"CustomProperties"`
	Description                                                 string `ps:"Description"`
	Disabled                                                    bool   `ps:"Disabled"`
	ExternalScenarioData                                        any    `ps:"ExternalScenarioData"`
	Filter                                                      string `ps:"Filter"`
	LogicalOperationName                                        string `ps:"LogicalOperationName"`
	Name                                                        string `ps:"Name"`
	NotificationCulture                                         any    `ps:"NotificationCulture"`
	NotificationEnabled                                         bool   `ps:"NotificationEnabled"`
	NotifyUser                                                  any    `ps:"NotifyUser"`
	NotifyUserOnFilterMatch                                     bool   `ps:"NotifyUserOnFilterMatch"`
	NotifyUserSuppressionExpiryDate                             any    `ps:"NotifyUserSuppressionExpiryDate"`
	NotifyUserThrottleThreshold                                 any    `ps:"NotifyUserThrottleThreshold"`
	NotifyUserThrottleWindow                                    any    `ps:"NotifyUserThrottleWindow"`
	Operation                                                   any    `ps:"Operation"`
	PrivacyManagementScopedSensitiveInformationTypes            any    `ps:"PrivacyManagementScopedSensitiveInformationTypes"`
	PrivacyManagementScopedSensitiveInformationTypesForCounting any    `ps:"PrivacyManagementScopedSensitiveInformationTypesForCounting"`
	PrivacyManagementScopedSensitiveInformationTypesThreshold   any    `ps:"PrivacyManagementScopedSensitiveInformationTypesThreshold"`
	Severity                                                    any    `ps:"Severity"`
	StreamType                                                  any    `ps:"StreamType"`
	ThreatType                                                  any    `ps:"ThreatType"`
	Threshold                                                   any    `ps:"Threshold"`
	TimeWindow                                                  any    `ps:"TimeWindow"`
	UseCreatedDateTime                                          any    `ps:"UseCreatedDateTime"`
	VolumeThreshold                                             any    `ps:"VolumeThreshold"`
}

func (p NewProtectionAlertParams) params() map[string]any {
	m := map[string]any{}
	if p.AggregationType != nil {
		m["AggregationType"] = p.AggregationType
	}
	if p.AlertBy != nil {
		m["AlertBy"] = p.AlertBy
	}
	if p.AlertFor != nil {
		m["AlertFor"] = p.AlertFor
	}
	if p.Category != nil {
		m["Category"] = p.Category
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.CorrelationPolicyId != nil {
		m["CorrelationPolicyId"] = p.CorrelationPolicyId
	}
	if p.CustomProperties != nil {
		m["CustomProperties"] = p.CustomProperties
	}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if p.Disabled {
		m["Disabled"] = true
	}
	if p.ExternalScenarioData != nil {
		m["ExternalScenarioData"] = p.ExternalScenarioData
	}
	if p.Filter != "" {
		m["Filter"] = p.Filter
	}
	if p.LogicalOperationName != "" {
		m["LogicalOperationName"] = p.LogicalOperationName
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.NotificationCulture != nil {
		m["NotificationCulture"] = p.NotificationCulture
	}
	if p.NotificationEnabled {
		m["NotificationEnabled"] = true
	}
	if p.NotifyUser != nil {
		m["NotifyUser"] = p.NotifyUser
	}
	if p.NotifyUserOnFilterMatch {
		m["NotifyUserOnFilterMatch"] = true
	}
	if p.NotifyUserSuppressionExpiryDate != nil {
		m["NotifyUserSuppressionExpiryDate"] = p.NotifyUserSuppressionExpiryDate
	}
	if p.NotifyUserThrottleThreshold != nil {
		m["NotifyUserThrottleThreshold"] = p.NotifyUserThrottleThreshold
	}
	if p.NotifyUserThrottleWindow != nil {
		m["NotifyUserThrottleWindow"] = p.NotifyUserThrottleWindow
	}
	if p.Operation != nil {
		m["Operation"] = p.Operation
	}
	if p.PrivacyManagementScopedSensitiveInformationTypes != nil {
		m["PrivacyManagementScopedSensitiveInformationTypes"] = p.PrivacyManagementScopedSensitiveInformationTypes
	}
	if p.PrivacyManagementScopedSensitiveInformationTypesForCounting != nil {
		m["PrivacyManagementScopedSensitiveInformationTypesForCounting"] = p.PrivacyManagementScopedSensitiveInformationTypesForCounting
	}
	if p.PrivacyManagementScopedSensitiveInformationTypesThreshold != nil {
		m["PrivacyManagementScopedSensitiveInformationTypesThreshold"] = p.PrivacyManagementScopedSensitiveInformationTypesThreshold
	}
	if p.Severity != nil {
		m["Severity"] = p.Severity
	}
	if p.StreamType != nil {
		m["StreamType"] = p.StreamType
	}
	if p.ThreatType != nil {
		m["ThreatType"] = p.ThreatType
	}
	if p.Threshold != nil {
		m["Threshold"] = p.Threshold
	}
	if p.TimeWindow != nil {
		m["TimeWindow"] = p.TimeWindow
	}
	if p.UseCreatedDateTime != nil {
		m["UseCreatedDateTime"] = p.UseCreatedDateTime
	}
	if p.VolumeThreshold != nil {
		m["VolumeThreshold"] = p.VolumeThreshold
	}
	return m
}

// NewProtectionAlert runs the New-ProtectionAlert cmdlet.
func (s *Service) NewProtectionAlert(ctx context.Context, p NewProtectionAlertParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-ProtectionAlert", p.params())
}

// NewProtectionCompliancePolicyParams are the parameters of New-ProtectionCompliancePolicy.
type NewProtectionCompliancePolicyParams struct {
	Enabled   bool   `ps:"Enabled"`
	Locations string `ps:"Locations"`
	Name      string `ps:"Name"`
}

func (p NewProtectionCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.Locations != "" {
		m["Locations"] = p.Locations
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	return m
}

// NewProtectionCompliancePolicy runs the New-ProtectionCompliancePolicy cmdlet.
func (s *Service) NewProtectionCompliancePolicy(ctx context.Context, p NewProtectionCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-ProtectionCompliancePolicy", p.params())
}

// NewProtectionComplianceRuleParams are the parameters of New-ProtectionComplianceRule.
type NewProtectionComplianceRuleParams struct {
	AdvancedRule   string   `ps:"AdvancedRule"`
	ContainsLabels any      `ps:"ContainsLabels"`
	LabelActions   []string `ps:"LabelActions"`
	Name           string   `ps:"Name"`
	Policy         any      `ps:"Policy"`
}

func (p NewProtectionComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.AdvancedRule != "" {
		m["AdvancedRule"] = p.AdvancedRule
	}
	if p.ContainsLabels != nil {
		m["ContainsLabels"] = p.ContainsLabels
	}
	if len(p.LabelActions) > 0 {
		m["LabelActions"] = p.LabelActions
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	return m
}

// NewProtectionComplianceRule runs the New-ProtectionComplianceRule cmdlet.
func (s *Service) NewProtectionComplianceRule(ctx context.Context, p NewProtectionComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-ProtectionComplianceRule", p.params())
}

// NewRetentionCompliancePolicyParams are the parameters of New-RetentionCompliancePolicy.
type NewRetentionCompliancePolicyParams struct {
	AdaptiveScopeLocation         any    `ps:"AdaptiveScopeLocation"`
	Applications                  any    `ps:"Applications"`
	Comment                       string `ps:"Comment"`
	Enabled                       bool   `ps:"Enabled"`
	ExchangeLocation              any    `ps:"ExchangeLocation"`
	ExchangeLocationException     any    `ps:"ExchangeLocationException"`
	ExcludeSPE                    bool   `ps:"ExcludeSPE"`
	Force                         bool   `ps:"Force"`
	IncludeUserOwnedContainers    bool   `ps:"IncludeUserOwnedContainers"`
	IsSimulation                  bool   `ps:"IsSimulation"`
	ModernGroupLocation           any    `ps:"ModernGroupLocation"`
	ModernGroupLocationException  any    `ps:"ModernGroupLocationException"`
	Name                          string `ps:"Name"`
	OneDriveLocation              any    `ps:"OneDriveLocation"`
	OneDriveLocationException     any    `ps:"OneDriveLocationException"`
	PolicyRBACScopes              any    `ps:"PolicyRBACScopes"`
	PolicyTemplateInfo            any    `ps:"PolicyTemplateInfo"`
	PublicFolderLocation          any    `ps:"PublicFolderLocation"`
	RestrictiveRetention          bool   `ps:"RestrictiveRetention"`
	RetainCloudAttachment         bool   `ps:"RetainCloudAttachment"`
	SharePointLocation            any    `ps:"SharePointLocation"`
	SharePointLocationException   any    `ps:"SharePointLocationException"`
	SkypeLocation                 any    `ps:"SkypeLocation"`
	SkypeLocationException        any    `ps:"SkypeLocationException"`
	TeamsChannelLocation          any    `ps:"TeamsChannelLocation"`
	TeamsChannelLocationException any    `ps:"TeamsChannelLocationException"`
	TeamsChatLocation             any    `ps:"TeamsChatLocation"`
	TeamsChatLocationException    any    `ps:"TeamsChatLocationException"`
}

func (p NewRetentionCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.AdaptiveScopeLocation != nil {
		m["AdaptiveScopeLocation"] = p.AdaptiveScopeLocation
	}
	if p.Applications != nil {
		m["Applications"] = p.Applications
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.ExchangeLocation != nil {
		m["ExchangeLocation"] = p.ExchangeLocation
	}
	if p.ExchangeLocationException != nil {
		m["ExchangeLocationException"] = p.ExchangeLocationException
	}
	if p.ExcludeSPE {
		m["ExcludeSPE"] = true
	}
	if p.Force {
		m["Force"] = true
	}
	if p.IncludeUserOwnedContainers {
		m["IncludeUserOwnedContainers"] = true
	}
	if p.IsSimulation {
		m["IsSimulation"] = true
	}
	if p.ModernGroupLocation != nil {
		m["ModernGroupLocation"] = p.ModernGroupLocation
	}
	if p.ModernGroupLocationException != nil {
		m["ModernGroupLocationException"] = p.ModernGroupLocationException
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.OneDriveLocation != nil {
		m["OneDriveLocation"] = p.OneDriveLocation
	}
	if p.OneDriveLocationException != nil {
		m["OneDriveLocationException"] = p.OneDriveLocationException
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.PolicyTemplateInfo != nil {
		m["PolicyTemplateInfo"] = p.PolicyTemplateInfo
	}
	if p.PublicFolderLocation != nil {
		m["PublicFolderLocation"] = p.PublicFolderLocation
	}
	if p.RestrictiveRetention {
		m["RestrictiveRetention"] = true
	}
	if p.RetainCloudAttachment {
		m["RetainCloudAttachment"] = true
	}
	if p.SharePointLocation != nil {
		m["SharePointLocation"] = p.SharePointLocation
	}
	if p.SharePointLocationException != nil {
		m["SharePointLocationException"] = p.SharePointLocationException
	}
	if p.SkypeLocation != nil {
		m["SkypeLocation"] = p.SkypeLocation
	}
	if p.SkypeLocationException != nil {
		m["SkypeLocationException"] = p.SkypeLocationException
	}
	if p.TeamsChannelLocation != nil {
		m["TeamsChannelLocation"] = p.TeamsChannelLocation
	}
	if p.TeamsChannelLocationException != nil {
		m["TeamsChannelLocationException"] = p.TeamsChannelLocationException
	}
	if p.TeamsChatLocation != nil {
		m["TeamsChatLocation"] = p.TeamsChatLocation
	}
	if p.TeamsChatLocationException != nil {
		m["TeamsChatLocationException"] = p.TeamsChatLocationException
	}
	return m
}

// NewRetentionCompliancePolicy runs the New-RetentionCompliancePolicy cmdlet.
func (s *Service) NewRetentionCompliancePolicy(ctx context.Context, p NewRetentionCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-RetentionCompliancePolicy", p.params())
}

// NewRetentionComplianceRuleParams are the parameters of New-RetentionComplianceRule.
type NewRetentionComplianceRuleParams struct {
	ApplyComplianceTag                  string   `ps:"ApplyComplianceTag"`
	ArchiveEnabled                      bool     `ps:"ArchiveEnabled"`
	ArchiveTriggerBasedOn               string   `ps:"ArchiveTriggerBasedOn"`
	ArchiveTriggerInDays                any      `ps:"ArchiveTriggerInDays"`
	Comment                             string   `ps:"Comment"`
	ContentContainsSensitiveInformation []string `ps:"ContentContainsSensitiveInformation"`
	ContentMatchQuery                   string   `ps:"ContentMatchQuery"`
	ExcludedItemClasses                 any      `ps:"ExcludedItemClasses"`
	ExpirationDateOption                string   `ps:"ExpirationDateOption"`
	IRMRiskyUserProfiles                any      `ps:"IRMRiskyUserProfiles"`
	MachineLearningModelIDs             any      `ps:"MachineLearningModelIDs"`
	Name                                string   `ps:"Name"`
	Policy                              any      `ps:"Policy"`
	PublishComplianceTag                string   `ps:"PublishComplianceTag"`
	RetentionComplianceAction           string   `ps:"RetentionComplianceAction"`
	RetentionDuration                   any      `ps:"RetentionDuration"`
	RetentionDurationDisplayHint        any      `ps:"RetentionDurationDisplayHint"`
}

func (p NewRetentionComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.ApplyComplianceTag != "" {
		m["ApplyComplianceTag"] = p.ApplyComplianceTag
	}
	if p.ArchiveEnabled {
		m["ArchiveEnabled"] = true
	}
	if p.ArchiveTriggerBasedOn != "" {
		m["ArchiveTriggerBasedOn"] = p.ArchiveTriggerBasedOn
	}
	if p.ArchiveTriggerInDays != nil {
		m["ArchiveTriggerInDays"] = p.ArchiveTriggerInDays
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if len(p.ContentContainsSensitiveInformation) > 0 {
		m["ContentContainsSensitiveInformation"] = p.ContentContainsSensitiveInformation
	}
	if p.ContentMatchQuery != "" {
		m["ContentMatchQuery"] = p.ContentMatchQuery
	}
	if p.ExcludedItemClasses != nil {
		m["ExcludedItemClasses"] = p.ExcludedItemClasses
	}
	if p.ExpirationDateOption != "" {
		m["ExpirationDateOption"] = p.ExpirationDateOption
	}
	if p.IRMRiskyUserProfiles != nil {
		m["IRMRiskyUserProfiles"] = p.IRMRiskyUserProfiles
	}
	if p.MachineLearningModelIDs != nil {
		m["MachineLearningModelIDs"] = p.MachineLearningModelIDs
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	if p.PublishComplianceTag != "" {
		m["PublishComplianceTag"] = p.PublishComplianceTag
	}
	if p.RetentionComplianceAction != "" {
		m["RetentionComplianceAction"] = p.RetentionComplianceAction
	}
	if p.RetentionDuration != nil {
		m["RetentionDuration"] = p.RetentionDuration
	}
	if p.RetentionDurationDisplayHint != nil {
		m["RetentionDurationDisplayHint"] = p.RetentionDurationDisplayHint
	}
	return m
}

// NewRetentionComplianceRule runs the New-RetentionComplianceRule cmdlet.
func (s *Service) NewRetentionComplianceRule(ctx context.Context, p NewRetentionComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-RetentionComplianceRule", p.params())
}

// NewSensitiveInformationScanParams are the parameters of New-SensitiveInformationScan.
type NewSensitiveInformationScanParams struct {
	Comment                           string   `ps:"Comment"`
	EndpointDlpLocation               any      `ps:"EndpointDlpLocation"`
	EndpointDlpLocationException      any      `ps:"EndpointDlpLocationException"`
	ExceptIfOneDriveSharedBy          []string `ps:"ExceptIfOneDriveSharedBy"`
	ExceptIfOneDriveSharedByMemberOf  []string `ps:"ExceptIfOneDriveSharedByMemberOf"`
	ExchangeAdaptiveScopes            any      `ps:"ExchangeAdaptiveScopes"`
	ExchangeAdaptiveScopesException   any      `ps:"ExchangeAdaptiveScopesException"`
	ExchangeLocation                  any      `ps:"ExchangeLocation"`
	ExchangeSender                    []string `ps:"ExchangeSender"`
	ExchangeSenderException           []string `ps:"ExchangeSenderException"`
	ExchangeSenderMemberOf            []string `ps:"ExchangeSenderMemberOf"`
	ExchangeSenderMemberOfException   []string `ps:"ExchangeSenderMemberOfException"`
	Name                              string   `ps:"Name"`
	OneDriveAdaptiveScopes            any      `ps:"OneDriveAdaptiveScopes"`
	OneDriveAdaptiveScopesException   any      `ps:"OneDriveAdaptiveScopesException"`
	OneDriveLocation                  any      `ps:"OneDriveLocation"`
	OneDriveLocationException         any      `ps:"OneDriveLocationException"`
	OneDriveSharedBy                  []string `ps:"OneDriveSharedBy"`
	OneDriveSharedByMemberOf          []string `ps:"OneDriveSharedByMemberOf"`
	PolicyRBACScopes                  any      `ps:"PolicyRBACScopes"`
	ScanBudget                        any      `ps:"ScanBudget"`
	ScanType                          string   `ps:"ScanType"`
	SharePointAdaptiveScopes          any      `ps:"SharePointAdaptiveScopes"`
	SharePointAdaptiveScopesException any      `ps:"SharePointAdaptiveScopesException"`
	SharePointLocation                any      `ps:"SharePointLocation"`
	SharePointLocationException       any      `ps:"SharePointLocationException"`
	TeamsAdaptiveScopes               any      `ps:"TeamsAdaptiveScopes"`
	TeamsAdaptiveScopesException      any      `ps:"TeamsAdaptiveScopesException"`
	TeamsLocation                     any      `ps:"TeamsLocation"`
	TeamsLocationException            any      `ps:"TeamsLocationException"`
}

func (p NewSensitiveInformationScanParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.EndpointDlpLocation != nil {
		m["EndpointDlpLocation"] = p.EndpointDlpLocation
	}
	if p.EndpointDlpLocationException != nil {
		m["EndpointDlpLocationException"] = p.EndpointDlpLocationException
	}
	if len(p.ExceptIfOneDriveSharedBy) > 0 {
		m["ExceptIfOneDriveSharedBy"] = p.ExceptIfOneDriveSharedBy
	}
	if len(p.ExceptIfOneDriveSharedByMemberOf) > 0 {
		m["ExceptIfOneDriveSharedByMemberOf"] = p.ExceptIfOneDriveSharedByMemberOf
	}
	if p.ExchangeAdaptiveScopes != nil {
		m["ExchangeAdaptiveScopes"] = p.ExchangeAdaptiveScopes
	}
	if p.ExchangeAdaptiveScopesException != nil {
		m["ExchangeAdaptiveScopesException"] = p.ExchangeAdaptiveScopesException
	}
	if p.ExchangeLocation != nil {
		m["ExchangeLocation"] = p.ExchangeLocation
	}
	if len(p.ExchangeSender) > 0 {
		m["ExchangeSender"] = p.ExchangeSender
	}
	if len(p.ExchangeSenderException) > 0 {
		m["ExchangeSenderException"] = p.ExchangeSenderException
	}
	if len(p.ExchangeSenderMemberOf) > 0 {
		m["ExchangeSenderMemberOf"] = p.ExchangeSenderMemberOf
	}
	if len(p.ExchangeSenderMemberOfException) > 0 {
		m["ExchangeSenderMemberOfException"] = p.ExchangeSenderMemberOfException
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.OneDriveAdaptiveScopes != nil {
		m["OneDriveAdaptiveScopes"] = p.OneDriveAdaptiveScopes
	}
	if p.OneDriveAdaptiveScopesException != nil {
		m["OneDriveAdaptiveScopesException"] = p.OneDriveAdaptiveScopesException
	}
	if p.OneDriveLocation != nil {
		m["OneDriveLocation"] = p.OneDriveLocation
	}
	if p.OneDriveLocationException != nil {
		m["OneDriveLocationException"] = p.OneDriveLocationException
	}
	if len(p.OneDriveSharedBy) > 0 {
		m["OneDriveSharedBy"] = p.OneDriveSharedBy
	}
	if len(p.OneDriveSharedByMemberOf) > 0 {
		m["OneDriveSharedByMemberOf"] = p.OneDriveSharedByMemberOf
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.ScanBudget != nil {
		m["ScanBudget"] = p.ScanBudget
	}
	if p.ScanType != "" {
		m["ScanType"] = p.ScanType
	}
	if p.SharePointAdaptiveScopes != nil {
		m["SharePointAdaptiveScopes"] = p.SharePointAdaptiveScopes
	}
	if p.SharePointAdaptiveScopesException != nil {
		m["SharePointAdaptiveScopesException"] = p.SharePointAdaptiveScopesException
	}
	if p.SharePointLocation != nil {
		m["SharePointLocation"] = p.SharePointLocation
	}
	if p.SharePointLocationException != nil {
		m["SharePointLocationException"] = p.SharePointLocationException
	}
	if p.TeamsAdaptiveScopes != nil {
		m["TeamsAdaptiveScopes"] = p.TeamsAdaptiveScopes
	}
	if p.TeamsAdaptiveScopesException != nil {
		m["TeamsAdaptiveScopesException"] = p.TeamsAdaptiveScopesException
	}
	if p.TeamsLocation != nil {
		m["TeamsLocation"] = p.TeamsLocation
	}
	if p.TeamsLocationException != nil {
		m["TeamsLocationException"] = p.TeamsLocationException
	}
	return m
}

// NewSensitiveInformationScan runs the New-SensitiveInformationScan cmdlet.
func (s *Service) NewSensitiveInformationScan(ctx context.Context, p NewSensitiveInformationScanParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-SensitiveInformationScan", p.params())
}

// NewSensitiveInformationScanRuleParams are the parameters of New-SensitiveInformationScanRule.
type NewSensitiveInformationScanRuleParams struct {
	Comment                                        string   `ps:"Comment"`
	ContentContainsSensitiveInformation            []string `ps:"ContentContainsSensitiveInformation"`
	ContentCreatedOrUpdatedDateFrom                any      `ps:"ContentCreatedOrUpdatedDateFrom"`
	ContentCreatedOrUpdatedDateTo                  any      `ps:"ContentCreatedOrUpdatedDateTo"`
	ContentExtensionMatchesWords                   any      `ps:"ContentExtensionMatchesWords"`
	ContentIsNotLabeled                            bool     `ps:"ContentIsNotLabeled"`
	ContentPropertyContainsWords                   any      `ps:"ContentPropertyContainsWords"`
	DocumentSizeOver                               any      `ps:"DocumentSizeOver"`
	ExceptIfContentContainsSensitiveInformation    []string `ps:"ExceptIfContentContainsSensitiveInformation"`
	ExceptIfContentExtensionMatchesWords           any      `ps:"ExceptIfContentExtensionMatchesWords"`
	ExceptIfContentPropertyContainsWords           any      `ps:"ExceptIfContentPropertyContainsWords"`
	ExceptIfDocumentSizeOver                       any      `ps:"ExceptIfDocumentSizeOver"`
	ExceptIfOdcContentContainsSensitiveInformation []string `ps:"ExceptIfOdcContentContainsSensitiveInformation"`
	Name                                           string   `ps:"Name"`
	OdcContentContainsSensitiveInformation         []string `ps:"OdcContentContainsSensitiveInformation"`
	Policy                                         any      `ps:"Policy"`
	StartImpactAssessment                          bool     `ps:"StartImpactAssessment"`
	Workload                                       any      `ps:"Workload"`
}

func (p NewSensitiveInformationScanRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if len(p.ContentContainsSensitiveInformation) > 0 {
		m["ContentContainsSensitiveInformation"] = p.ContentContainsSensitiveInformation
	}
	if p.ContentCreatedOrUpdatedDateFrom != nil {
		m["ContentCreatedOrUpdatedDateFrom"] = p.ContentCreatedOrUpdatedDateFrom
	}
	if p.ContentCreatedOrUpdatedDateTo != nil {
		m["ContentCreatedOrUpdatedDateTo"] = p.ContentCreatedOrUpdatedDateTo
	}
	if p.ContentExtensionMatchesWords != nil {
		m["ContentExtensionMatchesWords"] = p.ContentExtensionMatchesWords
	}
	if p.ContentIsNotLabeled {
		m["ContentIsNotLabeled"] = true
	}
	if p.ContentPropertyContainsWords != nil {
		m["ContentPropertyContainsWords"] = p.ContentPropertyContainsWords
	}
	if p.DocumentSizeOver != nil {
		m["DocumentSizeOver"] = p.DocumentSizeOver
	}
	if len(p.ExceptIfContentContainsSensitiveInformation) > 0 {
		m["ExceptIfContentContainsSensitiveInformation"] = p.ExceptIfContentContainsSensitiveInformation
	}
	if p.ExceptIfContentExtensionMatchesWords != nil {
		m["ExceptIfContentExtensionMatchesWords"] = p.ExceptIfContentExtensionMatchesWords
	}
	if p.ExceptIfContentPropertyContainsWords != nil {
		m["ExceptIfContentPropertyContainsWords"] = p.ExceptIfContentPropertyContainsWords
	}
	if p.ExceptIfDocumentSizeOver != nil {
		m["ExceptIfDocumentSizeOver"] = p.ExceptIfDocumentSizeOver
	}
	if len(p.ExceptIfOdcContentContainsSensitiveInformation) > 0 {
		m["ExceptIfOdcContentContainsSensitiveInformation"] = p.ExceptIfOdcContentContainsSensitiveInformation
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if len(p.OdcContentContainsSensitiveInformation) > 0 {
		m["OdcContentContainsSensitiveInformation"] = p.OdcContentContainsSensitiveInformation
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	if p.StartImpactAssessment {
		m["StartImpactAssessment"] = true
	}
	if p.Workload != nil {
		m["Workload"] = p.Workload
	}
	return m
}

// NewSensitiveInformationScanRule runs the New-SensitiveInformationScanRule cmdlet.
func (s *Service) NewSensitiveInformationScanRule(ctx context.Context, p NewSensitiveInformationScanRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-SensitiveInformationScanRule", p.params())
}

// NewServiceDomainGroupParams are the parameters of New-ServiceDomainGroup.
// DefaultParameterSetName: Default
type NewServiceDomainGroupParams struct {
	Name          string `ps:"Name"`
	Organization  any    `ps:"Organization"`
	SiteAddresses string `ps:"SiteAddresses"`
	SiteGroupId   string `ps:"SiteGroupId"`
}

func (p NewServiceDomainGroupParams) params() map[string]any {
	m := map[string]any{}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Organization != nil {
		m["Organization"] = p.Organization
	}
	if p.SiteAddresses != "" {
		m["SiteAddresses"] = p.SiteAddresses
	}
	if p.SiteGroupId != "" {
		m["SiteGroupId"] = p.SiteGroupId
	}
	return m
}

// NewServiceDomainGroup runs the New-ServiceDomainGroup cmdlet.
func (s *Service) NewServiceDomainGroup(ctx context.Context, p NewServiceDomainGroupParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-ServiceDomainGroup", p.params())
}

// NewServicePrincipalParams are the parameters of New-ServicePrincipal.
type NewServicePrincipalParams struct {
	AppId        string `ps:"AppId"`
	DisplayName  string `ps:"DisplayName"`
	ObjectId     string `ps:"ObjectId"`
	Organization any    `ps:"Organization"`
	ServiceId    string `ps:"ServiceId"`
}

func (p NewServicePrincipalParams) params() map[string]any {
	m := map[string]any{}
	if p.AppId != "" {
		m["AppId"] = p.AppId
	}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.ObjectId != "" {
		m["ObjectId"] = p.ObjectId
	}
	if p.Organization != nil {
		m["Organization"] = p.Organization
	}
	if p.ServiceId != "" {
		m["ServiceId"] = p.ServiceId
	}
	return m
}

// NewServicePrincipal runs the New-ServicePrincipal cmdlet.
func (s *Service) NewServicePrincipal(ctx context.Context, p NewServicePrincipalParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-ServicePrincipal", p.params())
}

// NewSupervisoryReviewPolicyMailboxFoldersParams are the parameters of New-SupervisoryReviewPolicyMailboxFolders.
type NewSupervisoryReviewPolicyMailboxFoldersParams struct {
	PolicyName string   `ps:"PolicyName"`
	Reviewers  []string `ps:"Reviewers"`
}

func (p NewSupervisoryReviewPolicyMailboxFoldersParams) params() map[string]any {
	m := map[string]any{}
	if p.PolicyName != "" {
		m["PolicyName"] = p.PolicyName
	}
	if len(p.Reviewers) > 0 {
		m["Reviewers"] = p.Reviewers
	}
	return m
}

// NewSupervisoryReviewPolicyMailboxFolders runs the New-SupervisoryReviewPolicyMailboxFolders cmdlet.
func (s *Service) NewSupervisoryReviewPolicyMailboxFolders(ctx context.Context, p NewSupervisoryReviewPolicyMailboxFoldersParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-SupervisoryReviewPolicyMailboxFolders", p.params())
}

// NewSupervisoryReviewPolicyV2Params are the parameters of New-SupervisoryReviewPolicyV2.
type NewSupervisoryReviewPolicyV2Params struct {
	Comment                  string   `ps:"Comment"`
	Enabled                  bool     `ps:"Enabled"`
	Force                    bool     `ps:"Force"`
	Name                     string   `ps:"Name"`
	PolicyRBACScopes         any      `ps:"PolicyRBACScopes"`
	PreservationPeriodInDays int      `ps:"PreservationPeriodInDays"`
	Reviewers                []string `ps:"Reviewers"`
	UserReportingWorkloads   []string `ps:"UserReportingWorkloads"`
}

func (p NewSupervisoryReviewPolicyV2Params) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.PreservationPeriodInDays != 0 {
		m["PreservationPeriodInDays"] = p.PreservationPeriodInDays
	}
	if len(p.Reviewers) > 0 {
		m["Reviewers"] = p.Reviewers
	}
	if len(p.UserReportingWorkloads) > 0 {
		m["UserReportingWorkloads"] = p.UserReportingWorkloads
	}
	return m
}

// NewSupervisoryReviewPolicyV2 runs the New-SupervisoryReviewPolicyV2 cmdlet.
func (s *Service) NewSupervisoryReviewPolicyV2(ctx context.Context, p NewSupervisoryReviewPolicyV2Params) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-SupervisoryReviewPolicyV2", p.params())
}

// NewSupervisoryReviewRuleParams are the parameters of New-SupervisoryReviewRule.
type NewSupervisoryReviewRuleParams struct {
	AdvancedRule                        string   `ps:"AdvancedRule"`
	CcsiDataModelOperator               string   `ps:"CcsiDataModelOperator"`
	Condition                           string   `ps:"Condition"`
	ContentContainsSensitiveInformation []string `ps:"ContentContainsSensitiveInformation"`
	ContentMatchesDataModel             string   `ps:"ContentMatchesDataModel"`
	ContentSources                      []string `ps:"ContentSources"`
	DayXInsights                        bool     `ps:"DayXInsights"`
	ExceptIfFrom                        any      `ps:"ExceptIfFrom"`
	ExceptIfRecipientDomainIs           any      `ps:"ExceptIfRecipientDomainIs"`
	ExceptIfRevieweeIs                  any      `ps:"ExceptIfRevieweeIs"`
	ExceptIfSenderDomainIs              any      `ps:"ExceptIfSenderDomainIs"`
	ExceptIfSentTo                      any      `ps:"ExceptIfSentTo"`
	ExceptIfSubjectOrBodyContainsWords  any      `ps:"ExceptIfSubjectOrBodyContainsWords"`
	From                                any      `ps:"From"`
	IncludeAdaptiveScopes               []string `ps:"IncludeAdaptiveScopes"`
	InPurviewFilter                     string   `ps:"InPurviewFilter"`
	Name                                string   `ps:"Name"`
	Ocr                                 bool     `ps:"Ocr"`
	Policy                              any      `ps:"Policy"`
	PolicyRBACScopes                    any      `ps:"PolicyRBACScopes"`
	RateOfSampling                      string   `ps:"RateOfSampling"`
	SamplingRate                        int      `ps:"SamplingRate"`
	SentTo                              any      `ps:"SentTo"`
}

func (p NewSupervisoryReviewRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.AdvancedRule != "" {
		m["AdvancedRule"] = p.AdvancedRule
	}
	if p.CcsiDataModelOperator != "" {
		m["CcsiDataModelOperator"] = p.CcsiDataModelOperator
	}
	if p.Condition != "" {
		m["Condition"] = p.Condition
	}
	if len(p.ContentContainsSensitiveInformation) > 0 {
		m["ContentContainsSensitiveInformation"] = p.ContentContainsSensitiveInformation
	}
	if p.ContentMatchesDataModel != "" {
		m["ContentMatchesDataModel"] = p.ContentMatchesDataModel
	}
	if len(p.ContentSources) > 0 {
		m["ContentSources"] = p.ContentSources
	}
	if p.DayXInsights {
		m["DayXInsights"] = true
	}
	if p.ExceptIfFrom != nil {
		m["ExceptIfFrom"] = p.ExceptIfFrom
	}
	if p.ExceptIfRecipientDomainIs != nil {
		m["ExceptIfRecipientDomainIs"] = p.ExceptIfRecipientDomainIs
	}
	if p.ExceptIfRevieweeIs != nil {
		m["ExceptIfRevieweeIs"] = p.ExceptIfRevieweeIs
	}
	if p.ExceptIfSenderDomainIs != nil {
		m["ExceptIfSenderDomainIs"] = p.ExceptIfSenderDomainIs
	}
	if p.ExceptIfSentTo != nil {
		m["ExceptIfSentTo"] = p.ExceptIfSentTo
	}
	if p.ExceptIfSubjectOrBodyContainsWords != nil {
		m["ExceptIfSubjectOrBodyContainsWords"] = p.ExceptIfSubjectOrBodyContainsWords
	}
	if p.From != nil {
		m["From"] = p.From
	}
	if len(p.IncludeAdaptiveScopes) > 0 {
		m["IncludeAdaptiveScopes"] = p.IncludeAdaptiveScopes
	}
	if p.InPurviewFilter != "" {
		m["InPurviewFilter"] = p.InPurviewFilter
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Ocr {
		m["Ocr"] = true
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.RateOfSampling != "" {
		m["RateOfSampling"] = p.RateOfSampling
	}
	if p.SamplingRate != 0 {
		m["SamplingRate"] = p.SamplingRate
	}
	if p.SentTo != nil {
		m["SentTo"] = p.SentTo
	}
	return m
}

// NewSupervisoryReviewRule runs the New-SupervisoryReviewRule cmdlet.
func (s *Service) NewSupervisoryReviewRule(ctx context.Context, p NewSupervisoryReviewRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-SupervisoryReviewRule", p.params())
}

// NewTeamsRetentionCompliancePolicyParams are the parameters of New-TeamsRetentionCompliancePolicy.
type NewTeamsRetentionCompliancePolicyParams struct {
	Comment                       string `ps:"Comment"`
	Enabled                       bool   `ps:"Enabled"`
	Force                         bool   `ps:"Force"`
	Name                          string `ps:"Name"`
	RestrictiveRetention          bool   `ps:"RestrictiveRetention"`
	TeamsChannelLocation          any    `ps:"TeamsChannelLocation"`
	TeamsChannelLocationException any    `ps:"TeamsChannelLocationException"`
	TeamsChatLocation             any    `ps:"TeamsChatLocation"`
	TeamsChatLocationException    any    `ps:"TeamsChatLocationException"`
}

func (p NewTeamsRetentionCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.RestrictiveRetention {
		m["RestrictiveRetention"] = true
	}
	if p.TeamsChannelLocation != nil {
		m["TeamsChannelLocation"] = p.TeamsChannelLocation
	}
	if p.TeamsChannelLocationException != nil {
		m["TeamsChannelLocationException"] = p.TeamsChannelLocationException
	}
	if p.TeamsChatLocation != nil {
		m["TeamsChatLocation"] = p.TeamsChatLocation
	}
	if p.TeamsChatLocationException != nil {
		m["TeamsChatLocationException"] = p.TeamsChatLocationException
	}
	return m
}

// NewTeamsRetentionCompliancePolicy runs the New-TeamsRetentionCompliancePolicy cmdlet.
func (s *Service) NewTeamsRetentionCompliancePolicy(ctx context.Context, p NewTeamsRetentionCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-TeamsRetentionCompliancePolicy", p.params())
}

// NewTeamsRetentionComplianceRuleParams are the parameters of New-TeamsRetentionComplianceRule.
type NewTeamsRetentionComplianceRuleParams struct {
	Comment                   string `ps:"Comment"`
	Name                      string `ps:"Name"`
	Policy                    any    `ps:"Policy"`
	RetentionComplianceAction string `ps:"RetentionComplianceAction"`
	RetentionDuration         any    `ps:"RetentionDuration"`
}

func (p NewTeamsRetentionComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	if p.RetentionComplianceAction != "" {
		m["RetentionComplianceAction"] = p.RetentionComplianceAction
	}
	if p.RetentionDuration != nil {
		m["RetentionDuration"] = p.RetentionDuration
	}
	return m
}

// NewTeamsRetentionComplianceRule runs the New-TeamsRetentionComplianceRule cmdlet.
func (s *Service) NewTeamsRetentionComplianceRule(ctx context.Context, p NewTeamsRetentionComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-TeamsRetentionComplianceRule", p.params())
}

// NewTenantWideRetentionOperationParams are the parameters of New-TenantWideRetentionOperation.
type NewTenantWideRetentionOperationParams struct {
	AdaptiveScopeLocation        any    `ps:"AdaptiveScopeLocation"`
	Applications                 any    `ps:"Applications"`
	Comment                      string `ps:"Comment"`
	ContentMatchQuery            string `ps:"ContentMatchQuery"`
	DestinationLabel             string `ps:"DestinationLabel"`
	Enabled                      bool   `ps:"Enabled"`
	ExchangeLocation             any    `ps:"ExchangeLocation"`
	ExchangeLocationException    any    `ps:"ExchangeLocationException"`
	Force                        bool   `ps:"Force"`
	ModernGroupLocation          any    `ps:"ModernGroupLocation"`
	ModernGroupLocationException any    `ps:"ModernGroupLocationException"`
	Name                         string `ps:"Name"`
	OneDriveLocation             any    `ps:"OneDriveLocation"`
	OneDriveLocationException    any    `ps:"OneDriveLocationException"`
	Operation                    string `ps:"Operation"`
	SharePointLocation           any    `ps:"SharePointLocation"`
	SharePointLocationException  any    `ps:"SharePointLocationException"`
	SourceLabel                  any    `ps:"SourceLabel"`
}

func (p NewTenantWideRetentionOperationParams) params() map[string]any {
	m := map[string]any{}
	if p.AdaptiveScopeLocation != nil {
		m["AdaptiveScopeLocation"] = p.AdaptiveScopeLocation
	}
	if p.Applications != nil {
		m["Applications"] = p.Applications
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.ContentMatchQuery != "" {
		m["ContentMatchQuery"] = p.ContentMatchQuery
	}
	if p.DestinationLabel != "" {
		m["DestinationLabel"] = p.DestinationLabel
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.ExchangeLocation != nil {
		m["ExchangeLocation"] = p.ExchangeLocation
	}
	if p.ExchangeLocationException != nil {
		m["ExchangeLocationException"] = p.ExchangeLocationException
	}
	if p.Force {
		m["Force"] = true
	}
	if p.ModernGroupLocation != nil {
		m["ModernGroupLocation"] = p.ModernGroupLocation
	}
	if p.ModernGroupLocationException != nil {
		m["ModernGroupLocationException"] = p.ModernGroupLocationException
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.OneDriveLocation != nil {
		m["OneDriveLocation"] = p.OneDriveLocation
	}
	if p.OneDriveLocationException != nil {
		m["OneDriveLocationException"] = p.OneDriveLocationException
	}
	if p.Operation != "" {
		m["Operation"] = p.Operation
	}
	if p.SharePointLocation != nil {
		m["SharePointLocation"] = p.SharePointLocation
	}
	if p.SharePointLocationException != nil {
		m["SharePointLocationException"] = p.SharePointLocationException
	}
	if p.SourceLabel != nil {
		m["SourceLabel"] = p.SourceLabel
	}
	return m
}

// NewTenantWideRetentionOperation runs the New-TenantWideRetentionOperation cmdlet.
func (s *Service) NewTenantWideRetentionOperation(ctx context.Context, p NewTenantWideRetentionOperationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-TenantWideRetentionOperation", p.params())
}

// NewThreatResponsePolicyParams are the parameters of New-ThreatResponsePolicy.
type NewThreatResponsePolicyParams struct {
	Comment          string `ps:"Comment"`
	ExchangeLocation any    `ps:"ExchangeLocation"`
	Mode             any    `ps:"Mode"`
	Name             string `ps:"Name"`
}

func (p NewThreatResponsePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.ExchangeLocation != nil {
		m["ExchangeLocation"] = p.ExchangeLocation
	}
	if p.Mode != nil {
		m["Mode"] = p.Mode
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	return m
}

// NewThreatResponsePolicy runs the New-ThreatResponsePolicy cmdlet.
func (s *Service) NewThreatResponsePolicy(ctx context.Context, p NewThreatResponsePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-ThreatResponsePolicy", p.params())
}

// NewThreatResponseRuleParams are the parameters of New-ThreatResponseRule.
type NewThreatResponseRuleParams struct {
	DisableActiveSync                     bool   `ps:"DisableActiveSync"`
	DisableAutoForwardDefaultRemoteDomain bool   `ps:"DisableAutoForwardDefaultRemoteDomain"`
	DisableBasicAuth                      bool   `ps:"DisableBasicAuth"`
	DisableEWS                            bool   `ps:"DisableEWS"`
	DisableForwardingSMTPAddress          bool   `ps:"DisableForwardingSMTPAddress"`
	DisableIMAP                           bool   `ps:"DisableIMAP"`
	DisableMAPI                           bool   `ps:"DisableMAPI"`
	DisableMAPIOverHTTP                   bool   `ps:"DisableMAPIOverHTTP"`
	DisableOWA                            bool   `ps:"DisableOWA"`
	DisablePOP                            bool   `ps:"DisablePOP"`
	DisableTransportRuleForwarding        bool   `ps:"DisableTransportRuleForwarding"`
	ExceptIfHighRiskUser                  bool   `ps:"ExceptIfHighRiskUser"`
	ExceptIfMemberOfDL                    any    `ps:"ExceptIfMemberOfDL"`
	HighRiskUser                          bool   `ps:"HighRiskUser"`
	MemberOfDL                            any    `ps:"MemberOfDL"`
	Name                                  string `ps:"Name"`
	Policy                                any    `ps:"Policy"`
}

func (p NewThreatResponseRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.DisableActiveSync {
		m["DisableActiveSync"] = true
	}
	if p.DisableAutoForwardDefaultRemoteDomain {
		m["DisableAutoForwardDefaultRemoteDomain"] = true
	}
	if p.DisableBasicAuth {
		m["DisableBasicAuth"] = true
	}
	if p.DisableEWS {
		m["DisableEWS"] = true
	}
	if p.DisableForwardingSMTPAddress {
		m["DisableForwardingSMTPAddress"] = true
	}
	if p.DisableIMAP {
		m["DisableIMAP"] = true
	}
	if p.DisableMAPI {
		m["DisableMAPI"] = true
	}
	if p.DisableMAPIOverHTTP {
		m["DisableMAPIOverHTTP"] = true
	}
	if p.DisableOWA {
		m["DisableOWA"] = true
	}
	if p.DisablePOP {
		m["DisablePOP"] = true
	}
	if p.DisableTransportRuleForwarding {
		m["DisableTransportRuleForwarding"] = true
	}
	if p.ExceptIfHighRiskUser {
		m["ExceptIfHighRiskUser"] = true
	}
	if p.ExceptIfMemberOfDL != nil {
		m["ExceptIfMemberOfDL"] = p.ExceptIfMemberOfDL
	}
	if p.HighRiskUser {
		m["HighRiskUser"] = true
	}
	if p.MemberOfDL != nil {
		m["MemberOfDL"] = p.MemberOfDL
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Policy != nil {
		m["Policy"] = p.Policy
	}
	return m
}

// NewThreatResponseRule runs the New-ThreatResponseRule cmdlet.
func (s *Service) NewThreatResponseRule(ctx context.Context, p NewThreatResponseRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-ThreatResponseRule", p.params())
}

// NewUnifiedAuditLogRetentionPolicyParams are the parameters of New-UnifiedAuditLogRetentionPolicy.
type NewUnifiedAuditLogRetentionPolicyParams struct {
	Description       string `ps:"Description"`
	Name              string `ps:"Name"`
	Operations        any    `ps:"Operations"`
	Priority          int    `ps:"Priority"`
	RecordTypes       any    `ps:"RecordTypes"`
	RetentionDuration any    `ps:"RetentionDuration"`
	UserIds           any    `ps:"UserIds"`
}

func (p NewUnifiedAuditLogRetentionPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Operations != nil {
		m["Operations"] = p.Operations
	}
	if p.Priority != 0 {
		m["Priority"] = p.Priority
	}
	if p.RecordTypes != nil {
		m["RecordTypes"] = p.RecordTypes
	}
	if p.RetentionDuration != nil {
		m["RetentionDuration"] = p.RetentionDuration
	}
	if p.UserIds != nil {
		m["UserIds"] = p.UserIds
	}
	return m
}

// NewUnifiedAuditLogRetentionPolicy runs the New-UnifiedAuditLogRetentionPolicy cmdlet.
func (s *Service) NewUnifiedAuditLogRetentionPolicy(ctx context.Context, p NewUnifiedAuditLogRetentionPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "New-UnifiedAuditLogRetentionPolicy", p.params())
}

// PreviewQuarantineMessageParams are the parameters of Preview-QuarantineMessage.
type PreviewQuarantineMessageParams struct {
	EntityType       any    `ps:"EntityType"`
	Identity         any    `ps:"Identity"`
	RecipientAddress string `ps:"RecipientAddress"`
}

func (p PreviewQuarantineMessageParams) params() map[string]any {
	m := map[string]any{}
	if p.EntityType != nil {
		m["EntityType"] = p.EntityType
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.RecipientAddress != "" {
		m["RecipientAddress"] = p.RecipientAddress
	}
	return m
}

// PreviewQuarantineMessage runs the Preview-QuarantineMessage cmdlet.
func (s *Service) PreviewQuarantineMessage(ctx context.Context, p PreviewQuarantineMessageParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Preview-QuarantineMessage", p.params())
}

// PreviewQuarantineMessageV1Params are the parameters of Preview-QuarantineMessageV1.
type PreviewQuarantineMessageV1Params struct {
	EntityType       any    `ps:"EntityType"`
	Identity         any    `ps:"Identity"`
	RecipientAddress string `ps:"RecipientAddress"`
}

func (p PreviewQuarantineMessageV1Params) params() map[string]any {
	m := map[string]any{}
	if p.EntityType != nil {
		m["EntityType"] = p.EntityType
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.RecipientAddress != "" {
		m["RecipientAddress"] = p.RecipientAddress
	}
	return m
}

// PreviewQuarantineMessageV1 runs the Preview-QuarantineMessageV1 cmdlet.
func (s *Service) PreviewQuarantineMessageV1(ctx context.Context, p PreviewQuarantineMessageV1Params) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Preview-QuarantineMessageV1", p.params())
}

// ReleaseQuarantineMessageParams are the parameters of Release-QuarantineMessage.
// DefaultParameterSetName: IdentityOnly
type ReleaseQuarantineMessageParams struct {
	ActionType          any      `ps:"ActionType"`
	AllowSender         bool     `ps:"AllowSender"`
	DecryptionHint      string   `ps:"DecryptionHint"`
	Delete              bool     `ps:"Delete"`
	EntityType          any      `ps:"EntityType"`
	Force               bool     `ps:"Force"`
	Identities          []string `ps:"Identities"`
	Identity            any      `ps:"Identity"`
	ReleaseToAll        bool     `ps:"ReleaseToAll"`
	ReportFalsePositive bool     `ps:"ReportFalsePositive"`
	User                []string `ps:"User"`
}

func (p ReleaseQuarantineMessageParams) params() map[string]any {
	m := map[string]any{}
	if p.ActionType != nil {
		m["ActionType"] = p.ActionType
	}
	if p.AllowSender {
		m["AllowSender"] = true
	}
	if p.DecryptionHint != "" {
		m["DecryptionHint"] = p.DecryptionHint
	}
	if p.Delete {
		m["Delete"] = true
	}
	if p.EntityType != nil {
		m["EntityType"] = p.EntityType
	}
	if p.Force {
		m["Force"] = true
	}
	if len(p.Identities) > 0 {
		m["Identities"] = p.Identities
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.ReleaseToAll {
		m["ReleaseToAll"] = true
	}
	if p.ReportFalsePositive {
		m["ReportFalsePositive"] = true
	}
	if len(p.User) > 0 {
		m["User"] = p.User
	}
	return m
}

// ReleaseQuarantineMessage runs the Release-QuarantineMessage cmdlet.
func (s *Service) ReleaseQuarantineMessage(ctx context.Context, p ReleaseQuarantineMessageParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Release-QuarantineMessage", p.params())
}

// RemoveActivityAlertParams are the parameters of Remove-ActivityAlert.
type RemoveActivityAlertParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveActivityAlertParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveActivityAlert runs the Remove-ActivityAlert cmdlet.
func (s *Service) RemoveActivityAlert(ctx context.Context, p RemoveActivityAlertParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ActivityAlert", p.params())
}

// RemoveAdaptiveScopeParams are the parameters of Remove-AdaptiveScope.
type RemoveAdaptiveScopeParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveAdaptiveScopeParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveAdaptiveScope runs the Remove-AdaptiveScope cmdlet.
func (s *Service) RemoveAdaptiveScope(ctx context.Context, p RemoveAdaptiveScopeParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-AdaptiveScope", p.params())
}

// RemoveAppRetentionCompliancePolicyParams are the parameters of Remove-AppRetentionCompliancePolicy.
type RemoveAppRetentionCompliancePolicyParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveAppRetentionCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveAppRetentionCompliancePolicy runs the Remove-AppRetentionCompliancePolicy cmdlet.
func (s *Service) RemoveAppRetentionCompliancePolicy(ctx context.Context, p RemoveAppRetentionCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-AppRetentionCompliancePolicy", p.params())
}

// RemoveAppRetentionComplianceRuleParams are the parameters of Remove-AppRetentionComplianceRule.
type RemoveAppRetentionComplianceRuleParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveAppRetentionComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveAppRetentionComplianceRule runs the Remove-AppRetentionComplianceRule cmdlet.
func (s *Service) RemoveAppRetentionComplianceRule(ctx context.Context, p RemoveAppRetentionComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-AppRetentionComplianceRule", p.params())
}

// RemoveAuditConfigurationPolicyParams are the parameters of Remove-AuditConfigurationPolicy.
type RemoveAuditConfigurationPolicyParams struct {
	DomainController any `ps:"DomainController"`
	Identity         any `ps:"Identity"`
}

func (p RemoveAuditConfigurationPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveAuditConfigurationPolicy runs the Remove-AuditConfigurationPolicy cmdlet.
func (s *Service) RemoveAuditConfigurationPolicy(ctx context.Context, p RemoveAuditConfigurationPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-AuditConfigurationPolicy", p.params())
}

// RemoveAuditConfigurationRuleParams are the parameters of Remove-AuditConfigurationRule.
type RemoveAuditConfigurationRuleParams struct {
	DomainController any `ps:"DomainController"`
	Identity         any `ps:"Identity"`
}

func (p RemoveAuditConfigurationRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveAuditConfigurationRule runs the Remove-AuditConfigurationRule cmdlet.
func (s *Service) RemoveAuditConfigurationRule(ctx context.Context, p RemoveAuditConfigurationRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-AuditConfigurationRule", p.params())
}

// RemoveAutoSensitivityLabelPolicyParams are the parameters of Remove-AutoSensitivityLabelPolicy.
type RemoveAutoSensitivityLabelPolicyParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveAutoSensitivityLabelPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveAutoSensitivityLabelPolicy runs the Remove-AutoSensitivityLabelPolicy cmdlet.
func (s *Service) RemoveAutoSensitivityLabelPolicy(ctx context.Context, p RemoveAutoSensitivityLabelPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-AutoSensitivityLabelPolicy", p.params())
}

// RemoveAutoSensitivityLabelRuleParams are the parameters of Remove-AutoSensitivityLabelRule.
type RemoveAutoSensitivityLabelRuleParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveAutoSensitivityLabelRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveAutoSensitivityLabelRule runs the Remove-AutoSensitivityLabelRule cmdlet.
func (s *Service) RemoveAutoSensitivityLabelRule(ctx context.Context, p RemoveAutoSensitivityLabelRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-AutoSensitivityLabelRule", p.params())
}

// RemoveCaseHoldPolicyParams are the parameters of Remove-CaseHoldPolicy.
type RemoveCaseHoldPolicyParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveCaseHoldPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveCaseHoldPolicy runs the Remove-CaseHoldPolicy cmdlet.
func (s *Service) RemoveCaseHoldPolicy(ctx context.Context, p RemoveCaseHoldPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-CaseHoldPolicy", p.params())
}

// RemoveCaseHoldRuleParams are the parameters of Remove-CaseHoldRule.
type RemoveCaseHoldRuleParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveCaseHoldRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveCaseHoldRule runs the Remove-CaseHoldRule cmdlet.
func (s *Service) RemoveCaseHoldRule(ctx context.Context, p RemoveCaseHoldRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-CaseHoldRule", p.params())
}

// RemoveClassificationGradingPolicyParams are the parameters of Remove-ClassificationGradingPolicy.
type RemoveClassificationGradingPolicyParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveClassificationGradingPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveClassificationGradingPolicy runs the Remove-ClassificationGradingPolicy cmdlet.
func (s *Service) RemoveClassificationGradingPolicy(ctx context.Context, p RemoveClassificationGradingPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ClassificationGradingPolicy", p.params())
}

// RemoveComplianceBoundaryParams are the parameters of Remove-ComplianceBoundary.
type RemoveComplianceBoundaryParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveComplianceBoundaryParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveComplianceBoundary runs the Remove-ComplianceBoundary cmdlet.
func (s *Service) RemoveComplianceBoundary(ctx context.Context, p RemoveComplianceBoundaryParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ComplianceBoundary", p.params())
}

// RemoveComplianceCaseParams are the parameters of Remove-ComplianceCase.
// DefaultParameterSetName: Identity
type RemoveComplianceCaseParams struct {
	DomainController any `ps:"DomainController"`
	Identity         any `ps:"Identity"`
}

func (p RemoveComplianceCaseParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveComplianceCase runs the Remove-ComplianceCase cmdlet.
func (s *Service) RemoveComplianceCase(ctx context.Context, p RemoveComplianceCaseParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ComplianceCase", p.params())
}

// RemoveComplianceCaseMemberParams are the parameters of Remove-ComplianceCaseMember.
// DefaultParameterSetName: Identity
type RemoveComplianceCaseMemberParams struct {
	Case   string `ps:"Case"`
	Member string `ps:"Member"`
}

func (p RemoveComplianceCaseMemberParams) params() map[string]any {
	m := map[string]any{}
	if p.Case != "" {
		m["Case"] = p.Case
	}
	if p.Member != "" {
		m["Member"] = p.Member
	}
	return m
}

// RemoveComplianceCaseMember runs the Remove-ComplianceCaseMember cmdlet.
func (s *Service) RemoveComplianceCaseMember(ctx context.Context, p RemoveComplianceCaseMemberParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ComplianceCaseMember", p.params())
}

// RemoveComplianceCustodianParams are the parameters of Remove-ComplianceCustodian.
// DefaultParameterSetName: Identity
type RemoveComplianceCustodianParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveComplianceCustodianParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveComplianceCustodian runs the Remove-ComplianceCustodian cmdlet.
func (s *Service) RemoveComplianceCustodian(ctx context.Context, p RemoveComplianceCustodianParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ComplianceCustodian", p.params())
}

// RemoveComplianceEmailSettingParams are the parameters of Remove-ComplianceEmailSetting.
// DefaultParameterSetName: Identity
type RemoveComplianceEmailSettingParams struct {
	DomainController any `ps:"DomainController"`
	Identity         any `ps:"Identity"`
}

func (p RemoveComplianceEmailSettingParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveComplianceEmailSetting runs the Remove-ComplianceEmailSetting cmdlet.
func (s *Service) RemoveComplianceEmailSetting(ctx context.Context, p RemoveComplianceEmailSettingParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ComplianceEmailSetting", p.params())
}

// RemoveComplianceNoticeParams are the parameters of Remove-ComplianceNotice.
// DefaultParameterSetName: Identity
type RemoveComplianceNoticeParams struct {
	DomainController any `ps:"DomainController"`
	Identity         any `ps:"Identity"`
}

func (p RemoveComplianceNoticeParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveComplianceNotice runs the Remove-ComplianceNotice cmdlet.
func (s *Service) RemoveComplianceNotice(ctx context.Context, p RemoveComplianceNoticeParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ComplianceNotice", p.params())
}

// RemoveComplianceRetentionEventParams are the parameters of Remove-ComplianceRetentionEvent.
type RemoveComplianceRetentionEventParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
	PreviewOnly   bool `ps:"PreviewOnly"`
}

func (p RemoveComplianceRetentionEventParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.PreviewOnly {
		m["PreviewOnly"] = true
	}
	return m
}

// RemoveComplianceRetentionEvent runs the Remove-ComplianceRetentionEvent cmdlet.
func (s *Service) RemoveComplianceRetentionEvent(ctx context.Context, p RemoveComplianceRetentionEventParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ComplianceRetentionEvent", p.params())
}

// RemoveComplianceRetentionEventTypeParams are the parameters of Remove-ComplianceRetentionEventType.
type RemoveComplianceRetentionEventTypeParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveComplianceRetentionEventTypeParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveComplianceRetentionEventType runs the Remove-ComplianceRetentionEventType cmdlet.
func (s *Service) RemoveComplianceRetentionEventType(ctx context.Context, p RemoveComplianceRetentionEventTypeParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ComplianceRetentionEventType", p.params())
}

// RemoveComplianceSearchParams are the parameters of Remove-ComplianceSearch.
type RemoveComplianceSearchParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveComplianceSearchParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveComplianceSearch runs the Remove-ComplianceSearch cmdlet.
func (s *Service) RemoveComplianceSearch(ctx context.Context, p RemoveComplianceSearchParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ComplianceSearch", p.params())
}

// RemoveComplianceSearchActionParams are the parameters of Remove-ComplianceSearchAction.
type RemoveComplianceSearchActionParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveComplianceSearchActionParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveComplianceSearchAction runs the Remove-ComplianceSearchAction cmdlet.
func (s *Service) RemoveComplianceSearchAction(ctx context.Context, p RemoveComplianceSearchActionParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ComplianceSearchAction", p.params())
}

// RemoveComplianceTagParams are the parameters of Remove-ComplianceTag.
type RemoveComplianceTagParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveComplianceTagParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveComplianceTag runs the Remove-ComplianceTag cmdlet.
func (s *Service) RemoveComplianceTag(ctx context.Context, p RemoveComplianceTagParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ComplianceTag", p.params())
}

// RemoveCustomDlpEmailTemplateParams are the parameters of Remove-CustomDlpEmailTemplate.
type RemoveCustomDlpEmailTemplateParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveCustomDlpEmailTemplateParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveCustomDlpEmailTemplate runs the Remove-CustomDlpEmailTemplate cmdlet.
func (s *Service) RemoveCustomDlpEmailTemplate(ctx context.Context, p RemoveCustomDlpEmailTemplateParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-CustomDlpEmailTemplate", p.params())
}

// RemoveDeviceConditionalAccessPolicyParams are the parameters of Remove-DeviceConditionalAccessPolicy.
type RemoveDeviceConditionalAccessPolicyParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveDeviceConditionalAccessPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveDeviceConditionalAccessPolicy runs the Remove-DeviceConditionalAccessPolicy cmdlet.
func (s *Service) RemoveDeviceConditionalAccessPolicy(ctx context.Context, p RemoveDeviceConditionalAccessPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-DeviceConditionalAccessPolicy", p.params())
}

// RemoveDeviceConditionalAccessRuleParams are the parameters of Remove-DeviceConditionalAccessRule.
type RemoveDeviceConditionalAccessRuleParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveDeviceConditionalAccessRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveDeviceConditionalAccessRule runs the Remove-DeviceConditionalAccessRule cmdlet.
func (s *Service) RemoveDeviceConditionalAccessRule(ctx context.Context, p RemoveDeviceConditionalAccessRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-DeviceConditionalAccessRule", p.params())
}

// RemoveDeviceConfigurationPolicyParams are the parameters of Remove-DeviceConfigurationPolicy.
type RemoveDeviceConfigurationPolicyParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveDeviceConfigurationPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveDeviceConfigurationPolicy runs the Remove-DeviceConfigurationPolicy cmdlet.
func (s *Service) RemoveDeviceConfigurationPolicy(ctx context.Context, p RemoveDeviceConfigurationPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-DeviceConfigurationPolicy", p.params())
}

// RemoveDeviceConfigurationRuleParams are the parameters of Remove-DeviceConfigurationRule.
type RemoveDeviceConfigurationRuleParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveDeviceConfigurationRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveDeviceConfigurationRule runs the Remove-DeviceConfigurationRule cmdlet.
func (s *Service) RemoveDeviceConfigurationRule(ctx context.Context, p RemoveDeviceConfigurationRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-DeviceConfigurationRule", p.params())
}

// RemoveDeviceTenantPolicyParams are the parameters of Remove-DeviceTenantPolicy.
type RemoveDeviceTenantPolicyParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveDeviceTenantPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveDeviceTenantPolicy runs the Remove-DeviceTenantPolicy cmdlet.
func (s *Service) RemoveDeviceTenantPolicy(ctx context.Context, p RemoveDeviceTenantPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-DeviceTenantPolicy", p.params())
}

// RemoveDeviceTenantRuleParams are the parameters of Remove-DeviceTenantRule.
type RemoveDeviceTenantRuleParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveDeviceTenantRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveDeviceTenantRule runs the Remove-DeviceTenantRule cmdlet.
func (s *Service) RemoveDeviceTenantRule(ctx context.Context, p RemoveDeviceTenantRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-DeviceTenantRule", p.params())
}

// RemoveDlpAlertTuningRuleParams are the parameters of Remove-DlpAlertTuningRule.
type RemoveDlpAlertTuningRuleParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveDlpAlertTuningRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveDlpAlertTuningRule runs the Remove-DlpAlertTuningRule cmdlet.
func (s *Service) RemoveDlpAlertTuningRule(ctx context.Context, p RemoveDlpAlertTuningRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-DlpAlertTuningRule", p.params())
}

// RemoveDlpCompliancePolicyParams are the parameters of Remove-DlpCompliancePolicy.
type RemoveDlpCompliancePolicyParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveDlpCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveDlpCompliancePolicy runs the Remove-DlpCompliancePolicy cmdlet.
func (s *Service) RemoveDlpCompliancePolicy(ctx context.Context, p RemoveDlpCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-DlpCompliancePolicy", p.params())
}

// RemoveDlpComplianceRuleParams are the parameters of Remove-DlpComplianceRule.
type RemoveDlpComplianceRuleParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveDlpComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveDlpComplianceRule runs the Remove-DlpComplianceRule cmdlet.
func (s *Service) RemoveDlpComplianceRule(ctx context.Context, p RemoveDlpComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-DlpComplianceRule", p.params())
}

// RemoveDlpEdmSchemaParams are the parameters of Remove-DlpEdmSchema.
type RemoveDlpEdmSchemaParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveDlpEdmSchemaParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveDlpEdmSchema runs the Remove-DlpEdmSchema cmdlet.
func (s *Service) RemoveDlpEdmSchema(ctx context.Context, p RemoveDlpEdmSchemaParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-DlpEdmSchema", p.params())
}

// RemoveDlpKeywordDictionaryParams are the parameters of Remove-DlpKeywordDictionary.
type RemoveDlpKeywordDictionaryParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveDlpKeywordDictionaryParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveDlpKeywordDictionary runs the Remove-DlpKeywordDictionary cmdlet.
func (s *Service) RemoveDlpKeywordDictionary(ctx context.Context, p RemoveDlpKeywordDictionaryParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-DlpKeywordDictionary", p.params())
}

// RemoveDlpSensitiveInformationTypeParams are the parameters of Remove-DlpSensitiveInformationType.
type RemoveDlpSensitiveInformationTypeParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveDlpSensitiveInformationTypeParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveDlpSensitiveInformationType runs the Remove-DlpSensitiveInformationType cmdlet.
func (s *Service) RemoveDlpSensitiveInformationType(ctx context.Context, p RemoveDlpSensitiveInformationTypeParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-DlpSensitiveInformationType", p.params())
}

// RemoveDlpSensitiveInformationTypeRulePackageParams are the parameters of Remove-DlpSensitiveInformationTypeRulePackage.
type RemoveDlpSensitiveInformationTypeRulePackageParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveDlpSensitiveInformationTypeRulePackageParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveDlpSensitiveInformationTypeRulePackage runs the Remove-DlpSensitiveInformationTypeRulePackage cmdlet.
func (s *Service) RemoveDlpSensitiveInformationTypeRulePackage(ctx context.Context, p RemoveDlpSensitiveInformationTypeRulePackageParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-DlpSensitiveInformationTypeRulePackage", p.params())
}

// RemoveDspmPolicyParams are the parameters of Remove-DspmPolicy.
type RemoveDspmPolicyParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveDspmPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveDspmPolicy runs the Remove-DspmPolicy cmdlet.
func (s *Service) RemoveDspmPolicy(ctx context.Context, p RemoveDspmPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-DspmPolicy", p.params())
}

// RemoveFeatureConfigurationParams are the parameters of Remove-FeatureConfiguration.
type RemoveFeatureConfigurationParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveFeatureConfigurationParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveFeatureConfiguration runs the Remove-FeatureConfiguration cmdlet.
func (s *Service) RemoveFeatureConfiguration(ctx context.Context, p RemoveFeatureConfigurationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-FeatureConfiguration", p.params())
}

// RemoveFilePlanPropertyAuthorityParams are the parameters of Remove-FilePlanPropertyAuthority.
type RemoveFilePlanPropertyAuthorityParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveFilePlanPropertyAuthorityParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveFilePlanPropertyAuthority runs the Remove-FilePlanPropertyAuthority cmdlet.
func (s *Service) RemoveFilePlanPropertyAuthority(ctx context.Context, p RemoveFilePlanPropertyAuthorityParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-FilePlanPropertyAuthority", p.params())
}

// RemoveFilePlanPropertyCategoryParams are the parameters of Remove-FilePlanPropertyCategory.
type RemoveFilePlanPropertyCategoryParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveFilePlanPropertyCategoryParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveFilePlanPropertyCategory runs the Remove-FilePlanPropertyCategory cmdlet.
func (s *Service) RemoveFilePlanPropertyCategory(ctx context.Context, p RemoveFilePlanPropertyCategoryParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-FilePlanPropertyCategory", p.params())
}

// RemoveFilePlanPropertyCitationParams are the parameters of Remove-FilePlanPropertyCitation.
type RemoveFilePlanPropertyCitationParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveFilePlanPropertyCitationParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveFilePlanPropertyCitation runs the Remove-FilePlanPropertyCitation cmdlet.
func (s *Service) RemoveFilePlanPropertyCitation(ctx context.Context, p RemoveFilePlanPropertyCitationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-FilePlanPropertyCitation", p.params())
}

// RemoveFilePlanPropertyDepartmentParams are the parameters of Remove-FilePlanPropertyDepartment.
type RemoveFilePlanPropertyDepartmentParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveFilePlanPropertyDepartmentParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveFilePlanPropertyDepartment runs the Remove-FilePlanPropertyDepartment cmdlet.
func (s *Service) RemoveFilePlanPropertyDepartment(ctx context.Context, p RemoveFilePlanPropertyDepartmentParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-FilePlanPropertyDepartment", p.params())
}

// RemoveFilePlanPropertyReferenceIdParams are the parameters of Remove-FilePlanPropertyReferenceId.
type RemoveFilePlanPropertyReferenceIdParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveFilePlanPropertyReferenceIdParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveFilePlanPropertyReferenceId runs the Remove-FilePlanPropertyReferenceId cmdlet.
func (s *Service) RemoveFilePlanPropertyReferenceId(ctx context.Context, p RemoveFilePlanPropertyReferenceIdParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-FilePlanPropertyReferenceId", p.params())
}

// RemoveFilePlanPropertySubCategoryParams are the parameters of Remove-FilePlanPropertySubCategory.
type RemoveFilePlanPropertySubCategoryParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveFilePlanPropertySubCategoryParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveFilePlanPropertySubCategory runs the Remove-FilePlanPropertySubCategory cmdlet.
func (s *Service) RemoveFilePlanPropertySubCategory(ctx context.Context, p RemoveFilePlanPropertySubCategoryParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-FilePlanPropertySubCategory", p.params())
}

// RemoveGlobalListParams are the parameters of Remove-GlobalList.
type RemoveGlobalListParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveGlobalListParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveGlobalList runs the Remove-GlobalList cmdlet.
func (s *Service) RemoveGlobalList(ctx context.Context, p RemoveGlobalListParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-GlobalList", p.params())
}

// RemoveHoldCompliancePolicyParams are the parameters of Remove-HoldCompliancePolicy.
type RemoveHoldCompliancePolicyParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveHoldCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveHoldCompliancePolicy runs the Remove-HoldCompliancePolicy cmdlet.
func (s *Service) RemoveHoldCompliancePolicy(ctx context.Context, p RemoveHoldCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-HoldCompliancePolicy", p.params())
}

// RemoveHoldComplianceRuleParams are the parameters of Remove-HoldComplianceRule.
type RemoveHoldComplianceRuleParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveHoldComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveHoldComplianceRule runs the Remove-HoldComplianceRule cmdlet.
func (s *Service) RemoveHoldComplianceRule(ctx context.Context, p RemoveHoldComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-HoldComplianceRule", p.params())
}

// RemoveInformationBarrierPolicyParams are the parameters of Remove-InformationBarrierPolicy.
type RemoveInformationBarrierPolicyParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveInformationBarrierPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveInformationBarrierPolicy runs the Remove-InformationBarrierPolicy cmdlet.
func (s *Service) RemoveInformationBarrierPolicy(ctx context.Context, p RemoveInformationBarrierPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-InformationBarrierPolicy", p.params())
}

// RemoveInsiderRiskEntityListParams are the parameters of Remove-InsiderRiskEntityList.
type RemoveInsiderRiskEntityListParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveInsiderRiskEntityListParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveInsiderRiskEntityList runs the Remove-InsiderRiskEntityList cmdlet.
func (s *Service) RemoveInsiderRiskEntityList(ctx context.Context, p RemoveInsiderRiskEntityListParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-InsiderRiskEntityList", p.params())
}

// RemoveInsiderRiskPolicyParams are the parameters of Remove-InsiderRiskPolicy.
type RemoveInsiderRiskPolicyParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveInsiderRiskPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveInsiderRiskPolicy runs the Remove-InsiderRiskPolicy cmdlet.
func (s *Service) RemoveInsiderRiskPolicy(ctx context.Context, p RemoveInsiderRiskPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-InsiderRiskPolicy", p.params())
}

// RemoveJitConfigurationParams are the parameters of Remove-JitConfiguration.
type RemoveJitConfigurationParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveJitConfigurationParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveJitConfiguration runs the Remove-JitConfiguration cmdlet.
func (s *Service) RemoveJitConfiguration(ctx context.Context, p RemoveJitConfigurationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-JitConfiguration", p.params())
}

// RemoveLabelParams are the parameters of Remove-Label.
type RemoveLabelParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveLabelParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveLabel runs the Remove-Label cmdlet.
func (s *Service) RemoveLabel(ctx context.Context, p RemoveLabelParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-Label", p.params())
}

// RemoveLabelExplorerConfigParams are the parameters of Remove-LabelExplorerConfig.
type RemoveLabelExplorerConfigParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveLabelExplorerConfigParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveLabelExplorerConfig runs the Remove-LabelExplorerConfig cmdlet.
func (s *Service) RemoveLabelExplorerConfig(ctx context.Context, p RemoveLabelExplorerConfigParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-LabelExplorerConfig", p.params())
}

// RemoveLabelPolicyParams are the parameters of Remove-LabelPolicy.
type RemoveLabelPolicyParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveLabelPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveLabelPolicy runs the Remove-LabelPolicy cmdlet.
func (s *Service) RemoveLabelPolicy(ctx context.Context, p RemoveLabelPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-LabelPolicy", p.params())
}

// RemoveMachineAssistedTagResourceParams are the parameters of Remove-MachineAssistedTagResource.
type RemoveMachineAssistedTagResourceParams struct {
	Identity any `ps:"Identity"`
	ModuleId any `ps:"ModuleId"`
	Resource any `ps:"Resource"`
	TaskId   any `ps:"TaskId"`
}

func (p RemoveMachineAssistedTagResourceParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.ModuleId != nil {
		m["ModuleId"] = p.ModuleId
	}
	if p.Resource != nil {
		m["Resource"] = p.Resource
	}
	if p.TaskId != nil {
		m["TaskId"] = p.TaskId
	}
	return m
}

// RemoveMachineAssistedTagResource runs the Remove-MachineAssistedTagResource cmdlet.
func (s *Service) RemoveMachineAssistedTagResource(ctx context.Context, p RemoveMachineAssistedTagResourceParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-MachineAssistedTagResource", p.params())
}

// RemoveOcrConfigurationParams are the parameters of Remove-OcrConfiguration.
type RemoveOcrConfigurationParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveOcrConfigurationParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveOcrConfiguration runs the Remove-OcrConfiguration cmdlet.
func (s *Service) RemoveOcrConfiguration(ctx context.Context, p RemoveOcrConfigurationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-OcrConfiguration", p.params())
}

// RemoveOrganizationSegmentParams are the parameters of Remove-OrganizationSegment.
type RemoveOrganizationSegmentParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveOrganizationSegmentParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveOrganizationSegment runs the Remove-OrganizationSegment cmdlet.
func (s *Service) RemoveOrganizationSegment(ctx context.Context, p RemoveOrganizationSegmentParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-OrganizationSegment", p.params())
}

// RemovePrivacyManagementCaseAdminParams are the parameters of Remove-PrivacyManagementCaseAdmin.
// DefaultParameterSetName: Identity
type RemovePrivacyManagementCaseAdminParams struct {
	User string `ps:"User"`
}

func (p RemovePrivacyManagementCaseAdminParams) params() map[string]any {
	m := map[string]any{}
	if p.User != "" {
		m["User"] = p.User
	}
	return m
}

// RemovePrivacyManagementCaseAdmin runs the Remove-PrivacyManagementCaseAdmin cmdlet.
func (s *Service) RemovePrivacyManagementCaseAdmin(ctx context.Context, p RemovePrivacyManagementCaseAdminParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-PrivacyManagementCaseAdmin", p.params())
}

// RemovePrivacyManagementComplianceCaseMemberParams are the parameters of Remove-PrivacyManagementComplianceCaseMember.
// DefaultParameterSetName: Identity
type RemovePrivacyManagementComplianceCaseMemberParams struct {
	Case   string `ps:"Case"`
	Member string `ps:"Member"`
}

func (p RemovePrivacyManagementComplianceCaseMemberParams) params() map[string]any {
	m := map[string]any{}
	if p.Case != "" {
		m["Case"] = p.Case
	}
	if p.Member != "" {
		m["Member"] = p.Member
	}
	return m
}

// RemovePrivacyManagementComplianceCaseMember runs the Remove-PrivacyManagementComplianceCaseMember cmdlet.
func (s *Service) RemovePrivacyManagementComplianceCaseMember(ctx context.Context, p RemovePrivacyManagementComplianceCaseMemberParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-PrivacyManagementComplianceCaseMember", p.params())
}

// RemovePrivacyManagementComplianceTagParams are the parameters of Remove-PrivacyManagementComplianceTag.
type RemovePrivacyManagementComplianceTagParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemovePrivacyManagementComplianceTagParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemovePrivacyManagementComplianceTag runs the Remove-PrivacyManagementComplianceTag cmdlet.
func (s *Service) RemovePrivacyManagementComplianceTag(ctx context.Context, p RemovePrivacyManagementComplianceTagParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-PrivacyManagementComplianceTag", p.params())
}

// RemovePrivacyManagementPolicyParams are the parameters of Remove-PrivacyManagementPolicy.
type RemovePrivacyManagementPolicyParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemovePrivacyManagementPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemovePrivacyManagementPolicy runs the Remove-PrivacyManagementPolicy cmdlet.
func (s *Service) RemovePrivacyManagementPolicy(ctx context.Context, p RemovePrivacyManagementPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-PrivacyManagementPolicy", p.params())
}

// RemovePrivacyManagementRuleParams are the parameters of Remove-PrivacyManagementRule.
// DefaultParameterSetName: Identity
type RemovePrivacyManagementRuleParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemovePrivacyManagementRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemovePrivacyManagementRule runs the Remove-PrivacyManagementRule cmdlet.
func (s *Service) RemovePrivacyManagementRule(ctx context.Context, p RemovePrivacyManagementRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-PrivacyManagementRule", p.params())
}

// RemoveProtectionAlertParams are the parameters of Remove-ProtectionAlert.
// DefaultParameterSetName: Identity
type RemoveProtectionAlertParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveProtectionAlertParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveProtectionAlert runs the Remove-ProtectionAlert cmdlet.
func (s *Service) RemoveProtectionAlert(ctx context.Context, p RemoveProtectionAlertParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ProtectionAlert", p.params())
}

// RemoveProtectionCompliancePolicyParams are the parameters of Remove-ProtectionCompliancePolicy.
type RemoveProtectionCompliancePolicyParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveProtectionCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveProtectionCompliancePolicy runs the Remove-ProtectionCompliancePolicy cmdlet.
func (s *Service) RemoveProtectionCompliancePolicy(ctx context.Context, p RemoveProtectionCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ProtectionCompliancePolicy", p.params())
}

// RemoveProtectionComplianceRuleParams are the parameters of Remove-ProtectionComplianceRule.
type RemoveProtectionComplianceRuleParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveProtectionComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveProtectionComplianceRule runs the Remove-ProtectionComplianceRule cmdlet.
func (s *Service) RemoveProtectionComplianceRule(ctx context.Context, p RemoveProtectionComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ProtectionComplianceRule", p.params())
}

// RemoveRecordLabelParams are the parameters of Remove-RecordLabel.
type RemoveRecordLabelParams struct {
	ItemUrl   string `ps:"ItemUrl"`
	LabelName string `ps:"LabelName"`
}

func (p RemoveRecordLabelParams) params() map[string]any {
	m := map[string]any{}
	if p.ItemUrl != "" {
		m["ItemUrl"] = p.ItemUrl
	}
	if p.LabelName != "" {
		m["LabelName"] = p.LabelName
	}
	return m
}

// RemoveRecordLabel runs the Remove-RecordLabel cmdlet.
func (s *Service) RemoveRecordLabel(ctx context.Context, p RemoveRecordLabelParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-RecordLabel", p.params())
}

// RemoveRetentionCompliancePolicyParams are the parameters of Remove-RetentionCompliancePolicy.
type RemoveRetentionCompliancePolicyParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveRetentionCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveRetentionCompliancePolicy runs the Remove-RetentionCompliancePolicy cmdlet.
func (s *Service) RemoveRetentionCompliancePolicy(ctx context.Context, p RemoveRetentionCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-RetentionCompliancePolicy", p.params())
}

// RemoveRetentionComplianceRuleParams are the parameters of Remove-RetentionComplianceRule.
type RemoveRetentionComplianceRuleParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveRetentionComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveRetentionComplianceRule runs the Remove-RetentionComplianceRule cmdlet.
func (s *Service) RemoveRetentionComplianceRule(ctx context.Context, p RemoveRetentionComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-RetentionComplianceRule", p.params())
}

// RemoveRoleGroupMemberParams are the parameters of Remove-RoleGroupMember.
type RemoveRoleGroupMemberParams struct {
}

func (p RemoveRoleGroupMemberParams) params() map[string]any {
	m := map[string]any{}
	return m
}

// RemoveRoleGroupMember runs the Remove-RoleGroupMember cmdlet.
func (s *Service) RemoveRoleGroupMember(ctx context.Context, p RemoveRoleGroupMemberParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-RoleGroupMember", p.params())
}

// RemoveSensitiveInformationScanParams are the parameters of Remove-SensitiveInformationScan.
type RemoveSensitiveInformationScanParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveSensitiveInformationScanParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveSensitiveInformationScan runs the Remove-SensitiveInformationScan cmdlet.
func (s *Service) RemoveSensitiveInformationScan(ctx context.Context, p RemoveSensitiveInformationScanParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-SensitiveInformationScan", p.params())
}

// RemoveSensitiveInformationScanRuleParams are the parameters of Remove-SensitiveInformationScanRule.
type RemoveSensitiveInformationScanRuleParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveSensitiveInformationScanRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveSensitiveInformationScanRule runs the Remove-SensitiveInformationScanRule cmdlet.
func (s *Service) RemoveSensitiveInformationScanRule(ctx context.Context, p RemoveSensitiveInformationScanRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-SensitiveInformationScanRule", p.params())
}

// RemoveServiceDomainGroupParams are the parameters of Remove-ServiceDomainGroup.
type RemoveServiceDomainGroupParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveServiceDomainGroupParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveServiceDomainGroup runs the Remove-ServiceDomainGroup cmdlet.
func (s *Service) RemoveServiceDomainGroup(ctx context.Context, p RemoveServiceDomainGroupParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ServiceDomainGroup", p.params())
}

// RemoveServicePrincipalParams are the parameters of Remove-ServicePrincipal.
type RemoveServicePrincipalParams struct {
	Identity any `ps:"Identity"`
}

func (p RemoveServicePrincipalParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveServicePrincipal runs the Remove-ServicePrincipal cmdlet.
func (s *Service) RemoveServicePrincipal(ctx context.Context, p RemoveServicePrincipalParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ServicePrincipal", p.params())
}

// RemoveSupervisoryReviewPolicyV2Params are the parameters of Remove-SupervisoryReviewPolicyV2.
type RemoveSupervisoryReviewPolicyV2Params struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveSupervisoryReviewPolicyV2Params) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveSupervisoryReviewPolicyV2 runs the Remove-SupervisoryReviewPolicyV2 cmdlet.
func (s *Service) RemoveSupervisoryReviewPolicyV2(ctx context.Context, p RemoveSupervisoryReviewPolicyV2Params) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-SupervisoryReviewPolicyV2", p.params())
}

// RemoveTeamsRetentionCompliancePolicyParams are the parameters of Remove-TeamsRetentionCompliancePolicy.
type RemoveTeamsRetentionCompliancePolicyParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveTeamsRetentionCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveTeamsRetentionCompliancePolicy runs the Remove-TeamsRetentionCompliancePolicy cmdlet.
func (s *Service) RemoveTeamsRetentionCompliancePolicy(ctx context.Context, p RemoveTeamsRetentionCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-TeamsRetentionCompliancePolicy", p.params())
}

// RemoveTeamsRetentionComplianceRuleParams are the parameters of Remove-TeamsRetentionComplianceRule.
type RemoveTeamsRetentionComplianceRuleParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveTeamsRetentionComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveTeamsRetentionComplianceRule runs the Remove-TeamsRetentionComplianceRule cmdlet.
func (s *Service) RemoveTeamsRetentionComplianceRule(ctx context.Context, p RemoveTeamsRetentionComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-TeamsRetentionComplianceRule", p.params())
}

// RemoveThreatResponsePolicyParams are the parameters of Remove-ThreatResponsePolicy.
type RemoveThreatResponsePolicyParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveThreatResponsePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveThreatResponsePolicy runs the Remove-ThreatResponsePolicy cmdlet.
func (s *Service) RemoveThreatResponsePolicy(ctx context.Context, p RemoveThreatResponsePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ThreatResponsePolicy", p.params())
}

// RemoveThreatResponseRuleParams are the parameters of Remove-ThreatResponseRule.
type RemoveThreatResponseRuleParams struct {
	ForceDeletion bool `ps:"ForceDeletion"`
	Identity      any  `ps:"Identity"`
}

func (p RemoveThreatResponseRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveThreatResponseRule runs the Remove-ThreatResponseRule cmdlet.
func (s *Service) RemoveThreatResponseRule(ctx context.Context, p RemoveThreatResponseRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-ThreatResponseRule", p.params())
}

// RemoveUnifiedAuditLogRetentionPolicyParams are the parameters of Remove-UnifiedAuditLogRetentionPolicy.
type RemoveUnifiedAuditLogRetentionPolicyParams struct {
	DomainController any  `ps:"DomainController"`
	ForceDeletion    bool `ps:"ForceDeletion"`
	Identity         any  `ps:"Identity"`
}

func (p RemoveUnifiedAuditLogRetentionPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.ForceDeletion {
		m["ForceDeletion"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// RemoveUnifiedAuditLogRetentionPolicy runs the Remove-UnifiedAuditLogRetentionPolicy cmdlet.
func (s *Service) RemoveUnifiedAuditLogRetentionPolicy(ctx context.Context, p RemoveUnifiedAuditLogRetentionPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Remove-UnifiedAuditLogRetentionPolicy", p.params())
}

// RestoreItemsDeletedByRetentionSystemParams are the parameters of Restore-ItemsDeletedByRetentionSystem.
type RestoreItemsDeletedByRetentionSystemParams struct {
	PolicyIdOrRuleIdOrTagTd any    `ps:"PolicyIdOrRuleIdOrTagTd"`
	RestoreFrom             any    `ps:"RestoreFrom"`
	RestoreTo               any    `ps:"RestoreTo"`
	SiteIds                 string `ps:"SiteIds"`
}

func (p RestoreItemsDeletedByRetentionSystemParams) params() map[string]any {
	m := map[string]any{}
	if p.PolicyIdOrRuleIdOrTagTd != nil {
		m["PolicyIdOrRuleIdOrTagTd"] = p.PolicyIdOrRuleIdOrTagTd
	}
	if p.RestoreFrom != nil {
		m["RestoreFrom"] = p.RestoreFrom
	}
	if p.RestoreTo != nil {
		m["RestoreTo"] = p.RestoreTo
	}
	if p.SiteIds != "" {
		m["SiteIds"] = p.SiteIds
	}
	return m
}

// RestoreItemsDeletedByRetentionSystem runs the Restore-ItemsDeletedByRetentionSystem cmdlet.
func (s *Service) RestoreItemsDeletedByRetentionSystem(ctx context.Context, p RestoreItemsDeletedByRetentionSystemParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Restore-ItemsDeletedByRetentionSystem", p.params())
}

// RestoreTemporaryDeletedQuarantineMessageParams are the parameters of Restore-TemporaryDeletedQuarantineMessage.
type RestoreTemporaryDeletedQuarantineMessageParams struct {
	Identities       []string `ps:"Identities"`
	Identity         any      `ps:"Identity"`
	RecipientAddress []string `ps:"RecipientAddress"`
}

func (p RestoreTemporaryDeletedQuarantineMessageParams) params() map[string]any {
	m := map[string]any{}
	if len(p.Identities) > 0 {
		m["Identities"] = p.Identities
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if len(p.RecipientAddress) > 0 {
		m["RecipientAddress"] = p.RecipientAddress
	}
	return m
}

// RestoreTemporaryDeletedQuarantineMessage runs the Restore-TemporaryDeletedQuarantineMessage cmdlet.
func (s *Service) RestoreTemporaryDeletedQuarantineMessage(ctx context.Context, p RestoreTemporaryDeletedQuarantineMessageParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Restore-TemporaryDeletedQuarantineMessage", p.params())
}

// SetActivityAlertParams are the parameters of Set-ActivityAlert.
// DefaultParameterSetName: Identity
type SetActivityAlertParams struct {
	Category     any    `ps:"Category"`
	Condition    string `ps:"Condition"`
	Description  string `ps:"Description"`
	Disabled     bool   `ps:"Disabled"`
	EmailCulture any    `ps:"EmailCulture"`
	Identity     any    `ps:"Identity"`
	Multiplier   any    `ps:"Multiplier"`
	NotifyUser   any    `ps:"NotifyUser"`
	Operation    any    `ps:"Operation"`
	RecordType   any    `ps:"RecordType"`
	ScopeLevel   any    `ps:"ScopeLevel"`
	Severity     any    `ps:"Severity"`
	Threshold    int    `ps:"Threshold"`
	TimeWindow   int    `ps:"TimeWindow"`
	UserId       any    `ps:"UserId"`
}

func (p SetActivityAlertParams) params() map[string]any {
	m := map[string]any{}
	if p.Category != nil {
		m["Category"] = p.Category
	}
	if p.Condition != "" {
		m["Condition"] = p.Condition
	}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if p.Disabled {
		m["Disabled"] = true
	}
	if p.EmailCulture != nil {
		m["EmailCulture"] = p.EmailCulture
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Multiplier != nil {
		m["Multiplier"] = p.Multiplier
	}
	if p.NotifyUser != nil {
		m["NotifyUser"] = p.NotifyUser
	}
	if p.Operation != nil {
		m["Operation"] = p.Operation
	}
	if p.RecordType != nil {
		m["RecordType"] = p.RecordType
	}
	if p.ScopeLevel != nil {
		m["ScopeLevel"] = p.ScopeLevel
	}
	if p.Severity != nil {
		m["Severity"] = p.Severity
	}
	if p.Threshold != 0 {
		m["Threshold"] = p.Threshold
	}
	if p.TimeWindow != 0 {
		m["TimeWindow"] = p.TimeWindow
	}
	if p.UserId != nil {
		m["UserId"] = p.UserId
	}
	return m
}

// SetActivityAlert runs the Set-ActivityAlert cmdlet.
func (s *Service) SetActivityAlert(ctx context.Context, p SetActivityAlertParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-ActivityAlert", p.params())
}

// SetAdaptiveScopeParams are the parameters of Set-AdaptiveScope.
type SetAdaptiveScopeParams struct {
	AdministrativeUnit           any      `ps:"AdministrativeUnit"`
	Comment                      string   `ps:"Comment"`
	EnabledStates                []string `ps:"EnabledStates"`
	FilterConditions             any      `ps:"FilterConditions"`
	Identity                     any      `ps:"Identity"`
	LinkedRecipientEnabledStates []string `ps:"LinkedRecipientEnabledStates"`
	RawQuery                     string   `ps:"RawQuery"`
}

func (p SetAdaptiveScopeParams) params() map[string]any {
	m := map[string]any{}
	if p.AdministrativeUnit != nil {
		m["AdministrativeUnit"] = p.AdministrativeUnit
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if len(p.EnabledStates) > 0 {
		m["EnabledStates"] = p.EnabledStates
	}
	if p.FilterConditions != nil {
		m["FilterConditions"] = p.FilterConditions
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if len(p.LinkedRecipientEnabledStates) > 0 {
		m["LinkedRecipientEnabledStates"] = p.LinkedRecipientEnabledStates
	}
	if p.RawQuery != "" {
		m["RawQuery"] = p.RawQuery
	}
	return m
}

// SetAdaptiveScope runs the Set-AdaptiveScope cmdlet.
func (s *Service) SetAdaptiveScope(ctx context.Context, p SetAdaptiveScopeParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-AdaptiveScope", p.params())
}

// SetAdministrativeUnitExtensionParams are the parameters of Set-AdministrativeUnitExtension.
type SetAdministrativeUnitExtensionParams struct {
	EnabledStates                []string `ps:"EnabledStates"`
	FilterConditions             any      `ps:"FilterConditions"`
	Identity                     any      `ps:"Identity"`
	LinkedRecipientEnabledStates []string `ps:"LinkedRecipientEnabledStates"`
	RawQuery                     string   `ps:"RawQuery"`
}

func (p SetAdministrativeUnitExtensionParams) params() map[string]any {
	m := map[string]any{}
	if len(p.EnabledStates) > 0 {
		m["EnabledStates"] = p.EnabledStates
	}
	if p.FilterConditions != nil {
		m["FilterConditions"] = p.FilterConditions
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if len(p.LinkedRecipientEnabledStates) > 0 {
		m["LinkedRecipientEnabledStates"] = p.LinkedRecipientEnabledStates
	}
	if p.RawQuery != "" {
		m["RawQuery"] = p.RawQuery
	}
	return m
}

// SetAdministrativeUnitExtension runs the Set-AdministrativeUnitExtension cmdlet.
func (s *Service) SetAdministrativeUnitExtension(ctx context.Context, p SetAdministrativeUnitExtensionParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-AdministrativeUnitExtension", p.params())
}

// SetAppRetentionCompliancePolicyParams are the parameters of Set-AppRetentionCompliancePolicy.
// DefaultParameterSetName: Identity
type SetAppRetentionCompliancePolicyParams struct {
	AddAdaptiveScopeLocation           any      `ps:"AddAdaptiveScopeLocation"`
	AddExchangeLocation                any      `ps:"AddExchangeLocation"`
	AddExchangeLocationException       any      `ps:"AddExchangeLocationException"`
	AddModernGroupLocation             any      `ps:"AddModernGroupLocation"`
	AddModernGroupLocationException    any      `ps:"AddModernGroupLocationException"`
	Applications                       []string `ps:"Applications"`
	Comment                            string   `ps:"Comment"`
	DeletedResources                   string   `ps:"DeletedResources"`
	Enabled                            bool     `ps:"Enabled"`
	Force                              bool     `ps:"Force"`
	Identity                           any      `ps:"Identity"`
	PolicyRBACScopes                   any      `ps:"PolicyRBACScopes"`
	RemoveAdaptiveScopeLocation        any      `ps:"RemoveAdaptiveScopeLocation"`
	RemoveExchangeLocation             any      `ps:"RemoveExchangeLocation"`
	RemoveExchangeLocationException    any      `ps:"RemoveExchangeLocationException"`
	RemoveModernGroupLocation          any      `ps:"RemoveModernGroupLocation"`
	RemoveModernGroupLocationException any      `ps:"RemoveModernGroupLocationException"`
	RestrictiveRetention               bool     `ps:"RestrictiveRetention"`
	RetryDistribution                  bool     `ps:"RetryDistribution"`
}

func (p SetAppRetentionCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.AddAdaptiveScopeLocation != nil {
		m["AddAdaptiveScopeLocation"] = p.AddAdaptiveScopeLocation
	}
	if p.AddExchangeLocation != nil {
		m["AddExchangeLocation"] = p.AddExchangeLocation
	}
	if p.AddExchangeLocationException != nil {
		m["AddExchangeLocationException"] = p.AddExchangeLocationException
	}
	if p.AddModernGroupLocation != nil {
		m["AddModernGroupLocation"] = p.AddModernGroupLocation
	}
	if p.AddModernGroupLocationException != nil {
		m["AddModernGroupLocationException"] = p.AddModernGroupLocationException
	}
	if len(p.Applications) > 0 {
		m["Applications"] = p.Applications
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.DeletedResources != "" {
		m["DeletedResources"] = p.DeletedResources
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.RemoveAdaptiveScopeLocation != nil {
		m["RemoveAdaptiveScopeLocation"] = p.RemoveAdaptiveScopeLocation
	}
	if p.RemoveExchangeLocation != nil {
		m["RemoveExchangeLocation"] = p.RemoveExchangeLocation
	}
	if p.RemoveExchangeLocationException != nil {
		m["RemoveExchangeLocationException"] = p.RemoveExchangeLocationException
	}
	if p.RemoveModernGroupLocation != nil {
		m["RemoveModernGroupLocation"] = p.RemoveModernGroupLocation
	}
	if p.RemoveModernGroupLocationException != nil {
		m["RemoveModernGroupLocationException"] = p.RemoveModernGroupLocationException
	}
	if p.RestrictiveRetention {
		m["RestrictiveRetention"] = true
	}
	if p.RetryDistribution {
		m["RetryDistribution"] = true
	}
	return m
}

// SetAppRetentionCompliancePolicy runs the Set-AppRetentionCompliancePolicy cmdlet.
func (s *Service) SetAppRetentionCompliancePolicy(ctx context.Context, p SetAppRetentionCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-AppRetentionCompliancePolicy", p.params())
}

// SetAppRetentionComplianceRuleParams are the parameters of Set-AppRetentionComplianceRule.
// DefaultParameterSetName: Identity
type SetAppRetentionComplianceRuleParams struct {
	ApplyComplianceTag                  string   `ps:"ApplyComplianceTag"`
	Comment                             string   `ps:"Comment"`
	ContentContainsSensitiveInformation []string `ps:"ContentContainsSensitiveInformation"`
	ContentDateFrom                     any      `ps:"ContentDateFrom"`
	ContentDateTo                       any      `ps:"ContentDateTo"`
	ContentMatchQuery                   string   `ps:"ContentMatchQuery"`
	ExcludedItemClasses                 any      `ps:"ExcludedItemClasses"`
	ExpirationDateOption                string   `ps:"ExpirationDateOption"`
	Identity                            any      `ps:"Identity"`
	RetentionComplianceAction           string   `ps:"RetentionComplianceAction"`
	RetentionDuration                   any      `ps:"RetentionDuration"`
	RetentionDurationDisplayHint        any      `ps:"RetentionDurationDisplayHint"`
}

func (p SetAppRetentionComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.ApplyComplianceTag != "" {
		m["ApplyComplianceTag"] = p.ApplyComplianceTag
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if len(p.ContentContainsSensitiveInformation) > 0 {
		m["ContentContainsSensitiveInformation"] = p.ContentContainsSensitiveInformation
	}
	if p.ContentDateFrom != nil {
		m["ContentDateFrom"] = p.ContentDateFrom
	}
	if p.ContentDateTo != nil {
		m["ContentDateTo"] = p.ContentDateTo
	}
	if p.ContentMatchQuery != "" {
		m["ContentMatchQuery"] = p.ContentMatchQuery
	}
	if p.ExcludedItemClasses != nil {
		m["ExcludedItemClasses"] = p.ExcludedItemClasses
	}
	if p.ExpirationDateOption != "" {
		m["ExpirationDateOption"] = p.ExpirationDateOption
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.RetentionComplianceAction != "" {
		m["RetentionComplianceAction"] = p.RetentionComplianceAction
	}
	if p.RetentionDuration != nil {
		m["RetentionDuration"] = p.RetentionDuration
	}
	if p.RetentionDurationDisplayHint != nil {
		m["RetentionDurationDisplayHint"] = p.RetentionDurationDisplayHint
	}
	return m
}

// SetAppRetentionComplianceRule runs the Set-AppRetentionComplianceRule cmdlet.
func (s *Service) SetAppRetentionComplianceRule(ctx context.Context, p SetAppRetentionComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-AppRetentionComplianceRule", p.params())
}

// SetAuditConfigParams are the parameters of Set-AuditConfig.
// DefaultParameterSetName: Identity
type SetAuditConfigParams struct {
	DomainController any `ps:"DomainController"`
	Workload         any `ps:"Workload"`
}

func (p SetAuditConfigParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.Workload != nil {
		m["Workload"] = p.Workload
	}
	return m
}

// SetAuditConfig runs the Set-AuditConfig cmdlet.
func (s *Service) SetAuditConfig(ctx context.Context, p SetAuditConfigParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-AuditConfig", p.params())
}

// SetAuditConfigurationRuleParams are the parameters of Set-AuditConfigurationRule.
// DefaultParameterSetName: Identity
type SetAuditConfigurationRuleParams struct {
	AuditOperation   any `ps:"AuditOperation"`
	DomainController any `ps:"DomainController"`
	Identity         any `ps:"Identity"`
}

func (p SetAuditConfigurationRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.AuditOperation != nil {
		m["AuditOperation"] = p.AuditOperation
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// SetAuditConfigurationRule runs the Set-AuditConfigurationRule cmdlet.
func (s *Service) SetAuditConfigurationRule(ctx context.Context, p SetAuditConfigurationRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-AuditConfigurationRule", p.params())
}

// SetAutoSensitivityLabelPolicyParams are the parameters of Set-AutoSensitivityLabelPolicy.
// DefaultParameterSetName: Identity
type SetAutoSensitivityLabelPolicyParams struct {
	AddExchangeLocation                     any      `ps:"AddExchangeLocation"`
	AddOneDriveLocation                     any      `ps:"AddOneDriveLocation"`
	AddOneDriveLocationException            any      `ps:"AddOneDriveLocationException"`
	AddPowerBILocation                      any      `ps:"AddPowerBILocation"`
	AddPowerBILocationException             any      `ps:"AddPowerBILocationException"`
	AddSharePointLocation                   any      `ps:"AddSharePointLocation"`
	AddSharePointLocationException          any      `ps:"AddSharePointLocationException"`
	ApplySensitivityLabel                   string   `ps:"ApplySensitivityLabel"`
	ApplySensitivityLabelOverwriteWorkloads any      `ps:"ApplySensitivityLabelOverwriteWorkloads"`
	AutoEnableAfter                         any      `ps:"AutoEnableAfter"`
	Comment                                 string   `ps:"Comment"`
	Enabled                                 bool     `ps:"Enabled"`
	EnforcementPlanes                       any      `ps:"EnforcementPlanes"`
	ExceptIfOneDriveSharedBy                []string `ps:"ExceptIfOneDriveSharedBy"`
	ExceptIfOneDriveSharedByMemberOf        []string `ps:"ExceptIfOneDriveSharedByMemberOf"`
	ExchangeAdaptiveScopes                  any      `ps:"ExchangeAdaptiveScopes"`
	ExchangeAdaptiveScopesException         any      `ps:"ExchangeAdaptiveScopesException"`
	ExchangeSender                          []string `ps:"ExchangeSender"`
	ExchangeSenderException                 []string `ps:"ExchangeSenderException"`
	ExchangeSenderMemberOf                  []string `ps:"ExchangeSenderMemberOf"`
	ExchangeSenderMemberOfException         []string `ps:"ExchangeSenderMemberOfException"`
	ExternalMailRightsManagementOwner       any      `ps:"ExternalMailRightsManagementOwner"`
	Force                                   bool     `ps:"Force"`
	Identity                                any      `ps:"Identity"`
	IncludeSimulationPolicyGrading          bool     `ps:"IncludeSimulationPolicyGrading"`
	Locations                               string   `ps:"Locations"`
	Mode                                    any      `ps:"Mode"`
	OneDriveAdaptiveScopes                  any      `ps:"OneDriveAdaptiveScopes"`
	OneDriveAdaptiveScopesException         any      `ps:"OneDriveAdaptiveScopesException"`
	OneDriveSharedBy                        []string `ps:"OneDriveSharedBy"`
	OneDriveSharedByMemberOf                []string `ps:"OneDriveSharedByMemberOf"`
	OverwriteLabel                          bool     `ps:"OverwriteLabel"`
	PolicyRBACScopes                        any      `ps:"PolicyRBACScopes"`
	PolicyTemplateInfo                      any      `ps:"PolicyTemplateInfo"`
	Priority                                any      `ps:"Priority"`
	RemoveExchangeLocation                  any      `ps:"RemoveExchangeLocation"`
	RemoveOneDriveLocation                  any      `ps:"RemoveOneDriveLocation"`
	RemoveOneDriveLocationException         any      `ps:"RemoveOneDriveLocationException"`
	RemovePowerBILocation                   any      `ps:"RemovePowerBILocation"`
	RemovePowerBILocationException          any      `ps:"RemovePowerBILocationException"`
	RemoveSharePointLocation                any      `ps:"RemoveSharePointLocation"`
	RemoveSharePointLocationException       any      `ps:"RemoveSharePointLocationException"`
	RetryDistribution                       bool     `ps:"RetryDistribution"`
	SharePointAdaptiveScopes                any      `ps:"SharePointAdaptiveScopes"`
	SharePointAdaptiveScopesException       any      `ps:"SharePointAdaptiveScopesException"`
	SpoAipIntegrationEnabled                bool     `ps:"SpoAipIntegrationEnabled"`
	StartSimulation                         bool     `ps:"StartSimulation"`
}

func (p SetAutoSensitivityLabelPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.AddExchangeLocation != nil {
		m["AddExchangeLocation"] = p.AddExchangeLocation
	}
	if p.AddOneDriveLocation != nil {
		m["AddOneDriveLocation"] = p.AddOneDriveLocation
	}
	if p.AddOneDriveLocationException != nil {
		m["AddOneDriveLocationException"] = p.AddOneDriveLocationException
	}
	if p.AddPowerBILocation != nil {
		m["AddPowerBILocation"] = p.AddPowerBILocation
	}
	if p.AddPowerBILocationException != nil {
		m["AddPowerBILocationException"] = p.AddPowerBILocationException
	}
	if p.AddSharePointLocation != nil {
		m["AddSharePointLocation"] = p.AddSharePointLocation
	}
	if p.AddSharePointLocationException != nil {
		m["AddSharePointLocationException"] = p.AddSharePointLocationException
	}
	if p.ApplySensitivityLabel != "" {
		m["ApplySensitivityLabel"] = p.ApplySensitivityLabel
	}
	if p.ApplySensitivityLabelOverwriteWorkloads != nil {
		m["ApplySensitivityLabelOverwriteWorkloads"] = p.ApplySensitivityLabelOverwriteWorkloads
	}
	if p.AutoEnableAfter != nil {
		m["AutoEnableAfter"] = p.AutoEnableAfter
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.EnforcementPlanes != nil {
		m["EnforcementPlanes"] = p.EnforcementPlanes
	}
	if len(p.ExceptIfOneDriveSharedBy) > 0 {
		m["ExceptIfOneDriveSharedBy"] = p.ExceptIfOneDriveSharedBy
	}
	if len(p.ExceptIfOneDriveSharedByMemberOf) > 0 {
		m["ExceptIfOneDriveSharedByMemberOf"] = p.ExceptIfOneDriveSharedByMemberOf
	}
	if p.ExchangeAdaptiveScopes != nil {
		m["ExchangeAdaptiveScopes"] = p.ExchangeAdaptiveScopes
	}
	if p.ExchangeAdaptiveScopesException != nil {
		m["ExchangeAdaptiveScopesException"] = p.ExchangeAdaptiveScopesException
	}
	if len(p.ExchangeSender) > 0 {
		m["ExchangeSender"] = p.ExchangeSender
	}
	if len(p.ExchangeSenderException) > 0 {
		m["ExchangeSenderException"] = p.ExchangeSenderException
	}
	if len(p.ExchangeSenderMemberOf) > 0 {
		m["ExchangeSenderMemberOf"] = p.ExchangeSenderMemberOf
	}
	if len(p.ExchangeSenderMemberOfException) > 0 {
		m["ExchangeSenderMemberOfException"] = p.ExchangeSenderMemberOfException
	}
	if p.ExternalMailRightsManagementOwner != nil {
		m["ExternalMailRightsManagementOwner"] = p.ExternalMailRightsManagementOwner
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IncludeSimulationPolicyGrading {
		m["IncludeSimulationPolicyGrading"] = true
	}
	if p.Locations != "" {
		m["Locations"] = p.Locations
	}
	if p.Mode != nil {
		m["Mode"] = p.Mode
	}
	if p.OneDriveAdaptiveScopes != nil {
		m["OneDriveAdaptiveScopes"] = p.OneDriveAdaptiveScopes
	}
	if p.OneDriveAdaptiveScopesException != nil {
		m["OneDriveAdaptiveScopesException"] = p.OneDriveAdaptiveScopesException
	}
	if len(p.OneDriveSharedBy) > 0 {
		m["OneDriveSharedBy"] = p.OneDriveSharedBy
	}
	if len(p.OneDriveSharedByMemberOf) > 0 {
		m["OneDriveSharedByMemberOf"] = p.OneDriveSharedByMemberOf
	}
	if p.OverwriteLabel {
		m["OverwriteLabel"] = true
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.PolicyTemplateInfo != nil {
		m["PolicyTemplateInfo"] = p.PolicyTemplateInfo
	}
	if p.Priority != nil {
		m["Priority"] = p.Priority
	}
	if p.RemoveExchangeLocation != nil {
		m["RemoveExchangeLocation"] = p.RemoveExchangeLocation
	}
	if p.RemoveOneDriveLocation != nil {
		m["RemoveOneDriveLocation"] = p.RemoveOneDriveLocation
	}
	if p.RemoveOneDriveLocationException != nil {
		m["RemoveOneDriveLocationException"] = p.RemoveOneDriveLocationException
	}
	if p.RemovePowerBILocation != nil {
		m["RemovePowerBILocation"] = p.RemovePowerBILocation
	}
	if p.RemovePowerBILocationException != nil {
		m["RemovePowerBILocationException"] = p.RemovePowerBILocationException
	}
	if p.RemoveSharePointLocation != nil {
		m["RemoveSharePointLocation"] = p.RemoveSharePointLocation
	}
	if p.RemoveSharePointLocationException != nil {
		m["RemoveSharePointLocationException"] = p.RemoveSharePointLocationException
	}
	if p.RetryDistribution {
		m["RetryDistribution"] = true
	}
	if p.SharePointAdaptiveScopes != nil {
		m["SharePointAdaptiveScopes"] = p.SharePointAdaptiveScopes
	}
	if p.SharePointAdaptiveScopesException != nil {
		m["SharePointAdaptiveScopesException"] = p.SharePointAdaptiveScopesException
	}
	if p.SpoAipIntegrationEnabled {
		m["SpoAipIntegrationEnabled"] = true
	}
	if p.StartSimulation {
		m["StartSimulation"] = true
	}
	return m
}

// SetAutoSensitivityLabelPolicy runs the Set-AutoSensitivityLabelPolicy cmdlet.
func (s *Service) SetAutoSensitivityLabelPolicy(ctx context.Context, p SetAutoSensitivityLabelPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-AutoSensitivityLabelPolicy", p.params())
}

// SetAutoSensitivityLabelRuleParams are the parameters of Set-AutoSensitivityLabelRule.
type SetAutoSensitivityLabelRuleParams struct {
	AccessScope                                  any      `ps:"AccessScope"`
	ActivationDate                               any      `ps:"ActivationDate"`
	AdvancedRule                                 string   `ps:"AdvancedRule"`
	AnyOfRecipientAddressContainsWords           any      `ps:"AnyOfRecipientAddressContainsWords"`
	AnyOfRecipientAddressMatchesPatterns         any      `ps:"AnyOfRecipientAddressMatchesPatterns"`
	Comment                                      string   `ps:"Comment"`
	ContentContainsSensitiveInformation          []string `ps:"ContentContainsSensitiveInformation"`
	ContentExtensionMatchesWords                 any      `ps:"ContentExtensionMatchesWords"`
	ContentIsNotLabeled                          bool     `ps:"ContentIsNotLabeled"`
	ContentPropertyContainsWords                 any      `ps:"ContentPropertyContainsWords"`
	DefaultSpoDocLibraryHasLabel                 bool     `ps:"DefaultSpoDocLibraryHasLabel"`
	Disabled                                     bool     `ps:"Disabled"`
	DocumentCreatedBy                            any      `ps:"DocumentCreatedBy"`
	DocumentIsPasswordProtected                  bool     `ps:"DocumentIsPasswordProtected"`
	DocumentIsUnsupported                        bool     `ps:"DocumentIsUnsupported"`
	DocumentNameMatchesWords                     any      `ps:"DocumentNameMatchesWords"`
	DocumentSizeOver                             any      `ps:"DocumentSizeOver"`
	ExceptIfAccessScope                          any      `ps:"ExceptIfAccessScope"`
	ExceptIfAnyOfRecipientAddressContainsWords   any      `ps:"ExceptIfAnyOfRecipientAddressContainsWords"`
	ExceptIfAnyOfRecipientAddressMatchesPatterns any      `ps:"ExceptIfAnyOfRecipientAddressMatchesPatterns"`
	ExceptIfContentContainsSensitiveInformation  []string `ps:"ExceptIfContentContainsSensitiveInformation"`
	ExceptIfContentExtensionMatchesWords         any      `ps:"ExceptIfContentExtensionMatchesWords"`
	ExceptIfContentPropertyContainsWords         any      `ps:"ExceptIfContentPropertyContainsWords"`
	ExceptIfDocumentCreatedBy                    any      `ps:"ExceptIfDocumentCreatedBy"`
	ExceptIfDocumentIsPasswordProtected          bool     `ps:"ExceptIfDocumentIsPasswordProtected"`
	ExceptIfDocumentIsUnsupported                bool     `ps:"ExceptIfDocumentIsUnsupported"`
	ExceptIfDocumentNameMatchesWords             any      `ps:"ExceptIfDocumentNameMatchesWords"`
	ExceptIfDocumentSizeOver                     any      `ps:"ExceptIfDocumentSizeOver"`
	ExceptIfFrom                                 []string `ps:"ExceptIfFrom"`
	ExceptIfFromAddressContainsWords             any      `ps:"ExceptIfFromAddressContainsWords"`
	ExceptIfFromAddressMatchesPatterns           any      `ps:"ExceptIfFromAddressMatchesPatterns"`
	ExceptIfFromMemberOf                         []string `ps:"ExceptIfFromMemberOf"`
	ExceptIfHeaderMatchesPatterns                any      `ps:"ExceptIfHeaderMatchesPatterns"`
	ExceptIfProcessingLimitExceeded              bool     `ps:"ExceptIfProcessingLimitExceeded"`
	ExceptIfRecipientDomainIs                    any      `ps:"ExceptIfRecipientDomainIs"`
	ExceptIfSenderDomainIs                       any      `ps:"ExceptIfSenderDomainIs"`
	ExceptIfSenderIPRanges                       any      `ps:"ExceptIfSenderIPRanges"`
	ExceptIfSentTo                               any      `ps:"ExceptIfSentTo"`
	ExceptIfSentToMemberOf                       []string `ps:"ExceptIfSentToMemberOf"`
	ExceptIfSharedWithDomain                     any      `ps:"ExceptIfSharedWithDomain"`
	ExceptIfSubjectMatchesPatterns               any      `ps:"ExceptIfSubjectMatchesPatterns"`
	ExpiryDate                                   any      `ps:"ExpiryDate"`
	From                                         []string `ps:"From"`
	FromAddressContainsWords                     any      `ps:"FromAddressContainsWords"`
	FromAddressMatchesPatterns                   any      `ps:"FromAddressMatchesPatterns"`
	FromMemberOf                                 []string `ps:"FromMemberOf"`
	HeaderMatchesPatterns                        any      `ps:"HeaderMatchesPatterns"`
	Identity                                     any      `ps:"Identity"`
	Priority                                     any      `ps:"Priority"`
	ProcessingLimitExceeded                      bool     `ps:"ProcessingLimitExceeded"`
	RecipientDomainIs                            any      `ps:"RecipientDomainIs"`
	ReportSeverityLevel                          any      `ps:"ReportSeverityLevel"`
	RuleErrorAction                              any      `ps:"RuleErrorAction"`
	SenderDomainIs                               any      `ps:"SenderDomainIs"`
	SenderIPRanges                               any      `ps:"SenderIPRanges"`
	SentTo                                       any      `ps:"SentTo"`
	SentToMemberOf                               []string `ps:"SentToMemberOf"`
	SharedWithDomain                             any      `ps:"SharedWithDomain"`
	SourceType                                   string   `ps:"SourceType"`
	SubjectMatchesPatterns                       any      `ps:"SubjectMatchesPatterns"`
	Workload                                     any      `ps:"Workload"`
}

func (p SetAutoSensitivityLabelRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.AccessScope != nil {
		m["AccessScope"] = p.AccessScope
	}
	if p.ActivationDate != nil {
		m["ActivationDate"] = p.ActivationDate
	}
	if p.AdvancedRule != "" {
		m["AdvancedRule"] = p.AdvancedRule
	}
	if p.AnyOfRecipientAddressContainsWords != nil {
		m["AnyOfRecipientAddressContainsWords"] = p.AnyOfRecipientAddressContainsWords
	}
	if p.AnyOfRecipientAddressMatchesPatterns != nil {
		m["AnyOfRecipientAddressMatchesPatterns"] = p.AnyOfRecipientAddressMatchesPatterns
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if len(p.ContentContainsSensitiveInformation) > 0 {
		m["ContentContainsSensitiveInformation"] = p.ContentContainsSensitiveInformation
	}
	if p.ContentExtensionMatchesWords != nil {
		m["ContentExtensionMatchesWords"] = p.ContentExtensionMatchesWords
	}
	if p.ContentIsNotLabeled {
		m["ContentIsNotLabeled"] = true
	}
	if p.ContentPropertyContainsWords != nil {
		m["ContentPropertyContainsWords"] = p.ContentPropertyContainsWords
	}
	if p.DefaultSpoDocLibraryHasLabel {
		m["DefaultSpoDocLibraryHasLabel"] = true
	}
	if p.Disabled {
		m["Disabled"] = true
	}
	if p.DocumentCreatedBy != nil {
		m["DocumentCreatedBy"] = p.DocumentCreatedBy
	}
	if p.DocumentIsPasswordProtected {
		m["DocumentIsPasswordProtected"] = true
	}
	if p.DocumentIsUnsupported {
		m["DocumentIsUnsupported"] = true
	}
	if p.DocumentNameMatchesWords != nil {
		m["DocumentNameMatchesWords"] = p.DocumentNameMatchesWords
	}
	if p.DocumentSizeOver != nil {
		m["DocumentSizeOver"] = p.DocumentSizeOver
	}
	if p.ExceptIfAccessScope != nil {
		m["ExceptIfAccessScope"] = p.ExceptIfAccessScope
	}
	if p.ExceptIfAnyOfRecipientAddressContainsWords != nil {
		m["ExceptIfAnyOfRecipientAddressContainsWords"] = p.ExceptIfAnyOfRecipientAddressContainsWords
	}
	if p.ExceptIfAnyOfRecipientAddressMatchesPatterns != nil {
		m["ExceptIfAnyOfRecipientAddressMatchesPatterns"] = p.ExceptIfAnyOfRecipientAddressMatchesPatterns
	}
	if len(p.ExceptIfContentContainsSensitiveInformation) > 0 {
		m["ExceptIfContentContainsSensitiveInformation"] = p.ExceptIfContentContainsSensitiveInformation
	}
	if p.ExceptIfContentExtensionMatchesWords != nil {
		m["ExceptIfContentExtensionMatchesWords"] = p.ExceptIfContentExtensionMatchesWords
	}
	if p.ExceptIfContentPropertyContainsWords != nil {
		m["ExceptIfContentPropertyContainsWords"] = p.ExceptIfContentPropertyContainsWords
	}
	if p.ExceptIfDocumentCreatedBy != nil {
		m["ExceptIfDocumentCreatedBy"] = p.ExceptIfDocumentCreatedBy
	}
	if p.ExceptIfDocumentIsPasswordProtected {
		m["ExceptIfDocumentIsPasswordProtected"] = true
	}
	if p.ExceptIfDocumentIsUnsupported {
		m["ExceptIfDocumentIsUnsupported"] = true
	}
	if p.ExceptIfDocumentNameMatchesWords != nil {
		m["ExceptIfDocumentNameMatchesWords"] = p.ExceptIfDocumentNameMatchesWords
	}
	if p.ExceptIfDocumentSizeOver != nil {
		m["ExceptIfDocumentSizeOver"] = p.ExceptIfDocumentSizeOver
	}
	if len(p.ExceptIfFrom) > 0 {
		m["ExceptIfFrom"] = p.ExceptIfFrom
	}
	if p.ExceptIfFromAddressContainsWords != nil {
		m["ExceptIfFromAddressContainsWords"] = p.ExceptIfFromAddressContainsWords
	}
	if p.ExceptIfFromAddressMatchesPatterns != nil {
		m["ExceptIfFromAddressMatchesPatterns"] = p.ExceptIfFromAddressMatchesPatterns
	}
	if len(p.ExceptIfFromMemberOf) > 0 {
		m["ExceptIfFromMemberOf"] = p.ExceptIfFromMemberOf
	}
	if p.ExceptIfHeaderMatchesPatterns != nil {
		m["ExceptIfHeaderMatchesPatterns"] = p.ExceptIfHeaderMatchesPatterns
	}
	if p.ExceptIfProcessingLimitExceeded {
		m["ExceptIfProcessingLimitExceeded"] = true
	}
	if p.ExceptIfRecipientDomainIs != nil {
		m["ExceptIfRecipientDomainIs"] = p.ExceptIfRecipientDomainIs
	}
	if p.ExceptIfSenderDomainIs != nil {
		m["ExceptIfSenderDomainIs"] = p.ExceptIfSenderDomainIs
	}
	if p.ExceptIfSenderIPRanges != nil {
		m["ExceptIfSenderIPRanges"] = p.ExceptIfSenderIPRanges
	}
	if p.ExceptIfSentTo != nil {
		m["ExceptIfSentTo"] = p.ExceptIfSentTo
	}
	if len(p.ExceptIfSentToMemberOf) > 0 {
		m["ExceptIfSentToMemberOf"] = p.ExceptIfSentToMemberOf
	}
	if p.ExceptIfSharedWithDomain != nil {
		m["ExceptIfSharedWithDomain"] = p.ExceptIfSharedWithDomain
	}
	if p.ExceptIfSubjectMatchesPatterns != nil {
		m["ExceptIfSubjectMatchesPatterns"] = p.ExceptIfSubjectMatchesPatterns
	}
	if p.ExpiryDate != nil {
		m["ExpiryDate"] = p.ExpiryDate
	}
	if len(p.From) > 0 {
		m["From"] = p.From
	}
	if p.FromAddressContainsWords != nil {
		m["FromAddressContainsWords"] = p.FromAddressContainsWords
	}
	if p.FromAddressMatchesPatterns != nil {
		m["FromAddressMatchesPatterns"] = p.FromAddressMatchesPatterns
	}
	if len(p.FromMemberOf) > 0 {
		m["FromMemberOf"] = p.FromMemberOf
	}
	if p.HeaderMatchesPatterns != nil {
		m["HeaderMatchesPatterns"] = p.HeaderMatchesPatterns
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Priority != nil {
		m["Priority"] = p.Priority
	}
	if p.ProcessingLimitExceeded {
		m["ProcessingLimitExceeded"] = true
	}
	if p.RecipientDomainIs != nil {
		m["RecipientDomainIs"] = p.RecipientDomainIs
	}
	if p.ReportSeverityLevel != nil {
		m["ReportSeverityLevel"] = p.ReportSeverityLevel
	}
	if p.RuleErrorAction != nil {
		m["RuleErrorAction"] = p.RuleErrorAction
	}
	if p.SenderDomainIs != nil {
		m["SenderDomainIs"] = p.SenderDomainIs
	}
	if p.SenderIPRanges != nil {
		m["SenderIPRanges"] = p.SenderIPRanges
	}
	if p.SentTo != nil {
		m["SentTo"] = p.SentTo
	}
	if len(p.SentToMemberOf) > 0 {
		m["SentToMemberOf"] = p.SentToMemberOf
	}
	if p.SharedWithDomain != nil {
		m["SharedWithDomain"] = p.SharedWithDomain
	}
	if p.SourceType != "" {
		m["SourceType"] = p.SourceType
	}
	if p.SubjectMatchesPatterns != nil {
		m["SubjectMatchesPatterns"] = p.SubjectMatchesPatterns
	}
	if p.Workload != nil {
		m["Workload"] = p.Workload
	}
	return m
}

// SetAutoSensitivityLabelRule runs the Set-AutoSensitivityLabelRule cmdlet.
func (s *Service) SetAutoSensitivityLabelRule(ctx context.Context, p SetAutoSensitivityLabelRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-AutoSensitivityLabelRule", p.params())
}

// SetCaseHoldPolicyParams are the parameters of Set-CaseHoldPolicy.
// DefaultParameterSetName: Identity
type SetCaseHoldPolicyParams struct {
	AddExchangeLocation        any    `ps:"AddExchangeLocation"`
	AddPublicFolderLocation    any    `ps:"AddPublicFolderLocation"`
	AddSharePointLocation      any    `ps:"AddSharePointLocation"`
	Comment                    string `ps:"Comment"`
	Enabled                    bool   `ps:"Enabled"`
	Force                      bool   `ps:"Force"`
	Identity                   any    `ps:"Identity"`
	RemoveExchangeLocation     any    `ps:"RemoveExchangeLocation"`
	RemovePublicFolderLocation any    `ps:"RemovePublicFolderLocation"`
	RemoveSharePointLocation   any    `ps:"RemoveSharePointLocation"`
	RetryDistribution          bool   `ps:"RetryDistribution"`
	UpdateStatistics           bool   `ps:"UpdateStatistics"`
}

func (p SetCaseHoldPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.AddExchangeLocation != nil {
		m["AddExchangeLocation"] = p.AddExchangeLocation
	}
	if p.AddPublicFolderLocation != nil {
		m["AddPublicFolderLocation"] = p.AddPublicFolderLocation
	}
	if p.AddSharePointLocation != nil {
		m["AddSharePointLocation"] = p.AddSharePointLocation
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.RemoveExchangeLocation != nil {
		m["RemoveExchangeLocation"] = p.RemoveExchangeLocation
	}
	if p.RemovePublicFolderLocation != nil {
		m["RemovePublicFolderLocation"] = p.RemovePublicFolderLocation
	}
	if p.RemoveSharePointLocation != nil {
		m["RemoveSharePointLocation"] = p.RemoveSharePointLocation
	}
	if p.RetryDistribution {
		m["RetryDistribution"] = true
	}
	if p.UpdateStatistics {
		m["UpdateStatistics"] = true
	}
	return m
}

// SetCaseHoldPolicy runs the Set-CaseHoldPolicy cmdlet.
func (s *Service) SetCaseHoldPolicy(ctx context.Context, p SetCaseHoldPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-CaseHoldPolicy", p.params())
}

// SetCaseHoldRuleParams are the parameters of Set-CaseHoldRule.
// DefaultParameterSetName: Identity
type SetCaseHoldRuleParams struct {
	Comment           string `ps:"Comment"`
	ContentDateFrom   any    `ps:"ContentDateFrom"`
	ContentDateTo     any    `ps:"ContentDateTo"`
	ContentMatchQuery string `ps:"ContentMatchQuery"`
	Disabled          bool   `ps:"Disabled"`
	Identity          any    `ps:"Identity"`
}

func (p SetCaseHoldRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.ContentDateFrom != nil {
		m["ContentDateFrom"] = p.ContentDateFrom
	}
	if p.ContentDateTo != nil {
		m["ContentDateTo"] = p.ContentDateTo
	}
	if p.ContentMatchQuery != "" {
		m["ContentMatchQuery"] = p.ContentMatchQuery
	}
	if p.Disabled {
		m["Disabled"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// SetCaseHoldRule runs the Set-CaseHoldRule cmdlet.
func (s *Service) SetCaseHoldRule(ctx context.Context, p SetCaseHoldRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-CaseHoldRule", p.params())
}

// SetClassificationGradingPolicyParams are the parameters of Set-ClassificationGradingPolicy.
// DefaultParameterSetName: Identity
type SetClassificationGradingPolicyParams struct {
	EndDateTime                 any      `ps:"EndDateTime"`
	Identity                    any      `ps:"Identity"`
	SensitiveInformationTypeIds []string `ps:"SensitiveInformationTypeIds"`
}

func (p SetClassificationGradingPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.EndDateTime != nil {
		m["EndDateTime"] = p.EndDateTime
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if len(p.SensitiveInformationTypeIds) > 0 {
		m["SensitiveInformationTypeIds"] = p.SensitiveInformationTypeIds
	}
	return m
}

// SetClassificationGradingPolicy runs the Set-ClassificationGradingPolicy cmdlet.
func (s *Service) SetClassificationGradingPolicy(ctx context.Context, p SetClassificationGradingPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-ClassificationGradingPolicy", p.params())
}

// SetComplianceBoundaryParams are the parameters of Set-ComplianceBoundary.
type SetComplianceBoundaryParams struct {
	Comment  string `ps:"Comment"`
	Identity any    `ps:"Identity"`
}

func (p SetComplianceBoundaryParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// SetComplianceBoundary runs the Set-ComplianceBoundary cmdlet.
func (s *Service) SetComplianceBoundary(ctx context.Context, p SetComplianceBoundaryParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-ComplianceBoundary", p.params())
}

// SetComplianceCaseParams are the parameters of Set-ComplianceCase.
// DefaultParameterSetName: Identity
type SetComplianceCaseParams struct {
	AddOrUpdateSources bool     `ps:"AddOrUpdateSources"`
	CaseType           any      `ps:"CaseType"`
	Close              bool     `ps:"Close"`
	Description        string   `ps:"Description"`
	DomainController   any      `ps:"DomainController"`
	ExternalId         string   `ps:"ExternalId"`
	Identity           any      `ps:"Identity"`
	Name               string   `ps:"Name"`
	RemoveSources      bool     `ps:"RemoveSources"`
	Reopen             bool     `ps:"Reopen"`
	Sources            []string `ps:"Sources"`
}

func (p SetComplianceCaseParams) params() map[string]any {
	m := map[string]any{}
	if p.AddOrUpdateSources {
		m["AddOrUpdateSources"] = true
	}
	if p.CaseType != nil {
		m["CaseType"] = p.CaseType
	}
	if p.Close {
		m["Close"] = true
	}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.ExternalId != "" {
		m["ExternalId"] = p.ExternalId
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.RemoveSources {
		m["RemoveSources"] = true
	}
	if p.Reopen {
		m["Reopen"] = true
	}
	if len(p.Sources) > 0 {
		m["Sources"] = p.Sources
	}
	return m
}

// SetComplianceCase runs the Set-ComplianceCase cmdlet.
func (s *Service) SetComplianceCase(ctx context.Context, p SetComplianceCaseParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-ComplianceCase", p.params())
}

// SetComplianceCustodianParams are the parameters of Set-ComplianceCustodian.
// DefaultParameterSetName: Identity
type SetComplianceCustodianParams struct {
	AddExchangeLocation        []string `ps:"AddExchangeLocation"`
	AddPublicFolderLocation    []string `ps:"AddPublicFolderLocation"`
	AddSharePointLocation      []string `ps:"AddSharePointLocation"`
	ContactEmail               string   `ps:"ContactEmail"`
	ExchangeLocation           []string `ps:"ExchangeLocation"`
	Identity                   any      `ps:"Identity"`
	PublicFolderLocation       []string `ps:"PublicFolderLocation"`
	RemoveExchangeLocation     []string `ps:"RemoveExchangeLocation"`
	RemovePublicFolderLocation []string `ps:"RemovePublicFolderLocation"`
	RemoveSharePointLocation   []string `ps:"RemoveSharePointLocation"`
	SharePointLocation         []string `ps:"SharePointLocation"`
}

func (p SetComplianceCustodianParams) params() map[string]any {
	m := map[string]any{}
	if len(p.AddExchangeLocation) > 0 {
		m["AddExchangeLocation"] = p.AddExchangeLocation
	}
	if len(p.AddPublicFolderLocation) > 0 {
		m["AddPublicFolderLocation"] = p.AddPublicFolderLocation
	}
	if len(p.AddSharePointLocation) > 0 {
		m["AddSharePointLocation"] = p.AddSharePointLocation
	}
	if p.ContactEmail != "" {
		m["ContactEmail"] = p.ContactEmail
	}
	if len(p.ExchangeLocation) > 0 {
		m["ExchangeLocation"] = p.ExchangeLocation
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if len(p.PublicFolderLocation) > 0 {
		m["PublicFolderLocation"] = p.PublicFolderLocation
	}
	if len(p.RemoveExchangeLocation) > 0 {
		m["RemoveExchangeLocation"] = p.RemoveExchangeLocation
	}
	if len(p.RemovePublicFolderLocation) > 0 {
		m["RemovePublicFolderLocation"] = p.RemovePublicFolderLocation
	}
	if len(p.RemoveSharePointLocation) > 0 {
		m["RemoveSharePointLocation"] = p.RemoveSharePointLocation
	}
	if len(p.SharePointLocation) > 0 {
		m["SharePointLocation"] = p.SharePointLocation
	}
	return m
}

// SetComplianceCustodian runs the Set-ComplianceCustodian cmdlet.
func (s *Service) SetComplianceCustodian(ctx context.Context, p SetComplianceCustodianParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-ComplianceCustodian", p.params())
}

// SetComplianceNoticeParams are the parameters of Set-ComplianceNotice.
// DefaultParameterSetName: Identity
type SetComplianceNoticeParams struct {
	Acknowledged           bool     `ps:"Acknowledged"`
	AcknowledgedCustodians []string `ps:"AcknowledgedCustodians"`
	AddedCustodians        []string `ps:"AddedCustodians"`
	ContentBody            string   `ps:"ContentBody"`
	Force                  bool     `ps:"Force"`
	From                   string   `ps:"From"`
	Identity               any      `ps:"Identity"`
	IsHighPriority         bool     `ps:"IsHighPriority"`
	Name                   string   `ps:"Name"`
	Subject                string   `ps:"Subject"`
	Type                   string   `ps:"Type"`
}

func (p SetComplianceNoticeParams) params() map[string]any {
	m := map[string]any{}
	if p.Acknowledged {
		m["Acknowledged"] = true
	}
	if len(p.AcknowledgedCustodians) > 0 {
		m["AcknowledgedCustodians"] = p.AcknowledgedCustodians
	}
	if len(p.AddedCustodians) > 0 {
		m["AddedCustodians"] = p.AddedCustodians
	}
	if p.ContentBody != "" {
		m["ContentBody"] = p.ContentBody
	}
	if p.Force {
		m["Force"] = true
	}
	if p.From != "" {
		m["From"] = p.From
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IsHighPriority {
		m["IsHighPriority"] = true
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Subject != "" {
		m["Subject"] = p.Subject
	}
	if p.Type != "" {
		m["Type"] = p.Type
	}
	return m
}

// SetComplianceNotice runs the Set-ComplianceNotice cmdlet.
func (s *Service) SetComplianceNotice(ctx context.Context, p SetComplianceNoticeParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-ComplianceNotice", p.params())
}

// SetCompliancePreservationSettingParams are the parameters of Set-CompliancePreservationSetting.
type SetCompliancePreservationSettingParams struct {
	CaseHoldCloudAttachmentEnabled bool `ps:"CaseHoldCloudAttachmentEnabled"`
}

func (p SetCompliancePreservationSettingParams) params() map[string]any {
	m := map[string]any{}
	if p.CaseHoldCloudAttachmentEnabled {
		m["CaseHoldCloudAttachmentEnabled"] = true
	}
	return m
}

// SetCompliancePreservationSetting runs the Set-CompliancePreservationSetting cmdlet.
func (s *Service) SetCompliancePreservationSetting(ctx context.Context, p SetCompliancePreservationSettingParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-CompliancePreservationSetting", p.params())
}

// SetComplianceRetentionEventParams are the parameters of Set-ComplianceRetentionEvent.
// DefaultParameterSetName: Identity
type SetComplianceRetentionEventParams struct {
	Action                 any    `ps:"Action"`
	AssetId                string `ps:"AssetId"`
	Comment                string `ps:"Comment"`
	DomainController       any    `ps:"DomainController"`
	EventTags              any    `ps:"EventTags"`
	EventType              any    `ps:"EventType"`
	ExchangeAssetIdQuery   string `ps:"ExchangeAssetIdQuery"`
	Identity               any    `ps:"Identity"`
	SharePointAssetIdQuery string `ps:"SharePointAssetIdQuery"`
}

func (p SetComplianceRetentionEventParams) params() map[string]any {
	m := map[string]any{}
	if p.Action != nil {
		m["Action"] = p.Action
	}
	if p.AssetId != "" {
		m["AssetId"] = p.AssetId
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.EventTags != nil {
		m["EventTags"] = p.EventTags
	}
	if p.EventType != nil {
		m["EventType"] = p.EventType
	}
	if p.ExchangeAssetIdQuery != "" {
		m["ExchangeAssetIdQuery"] = p.ExchangeAssetIdQuery
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.SharePointAssetIdQuery != "" {
		m["SharePointAssetIdQuery"] = p.SharePointAssetIdQuery
	}
	return m
}

// SetComplianceRetentionEvent runs the Set-ComplianceRetentionEvent cmdlet.
func (s *Service) SetComplianceRetentionEvent(ctx context.Context, p SetComplianceRetentionEventParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-ComplianceRetentionEvent", p.params())
}

// SetComplianceRetentionEventTypeParams are the parameters of Set-ComplianceRetentionEventType.
// DefaultParameterSetName: Identity
type SetComplianceRetentionEventTypeParams struct {
	Comment  string `ps:"Comment"`
	Identity any    `ps:"Identity"`
}

func (p SetComplianceRetentionEventTypeParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// SetComplianceRetentionEventType runs the Set-ComplianceRetentionEventType cmdlet.
func (s *Service) SetComplianceRetentionEventType(ctx context.Context, p SetComplianceRetentionEventTypeParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-ComplianceRetentionEventType", p.params())
}

// SetComplianceSearchParams are the parameters of Set-ComplianceSearch.
// DefaultParameterSetName: Identity
type SetComplianceSearchParams struct {
	AddExchangeLocation                   []string `ps:"AddExchangeLocation"`
	AddExchangeLocationExclusion          []string `ps:"AddExchangeLocationExclusion"`
	AddOneDriveLocation                   []string `ps:"AddOneDriveLocation"`
	AddOneDriveLocationExclusion          []string `ps:"AddOneDriveLocationExclusion"`
	AddPublicFolderLocation               []string `ps:"AddPublicFolderLocation"`
	AddPublicFolderLocationExclusion      []string `ps:"AddPublicFolderLocationExclusion"`
	AddSharePointLocation                 []string `ps:"AddSharePointLocation"`
	AddSharePointLocationExclusion        []string `ps:"AddSharePointLocationExclusion"`
	AllowNotFoundExchangeLocationsEnabled bool     `ps:"AllowNotFoundExchangeLocationsEnabled"`
	ContentMatchQuery                     string   `ps:"ContentMatchQuery"`
	Description                           string   `ps:"Description"`
	ExchangeLocation                      []string `ps:"ExchangeLocation"`
	ExchangeLocationExclusion             []string `ps:"ExchangeLocationExclusion"`
	Force                                 bool     `ps:"Force"`
	HoldNames                             []string `ps:"HoldNames"`
	Identity                              any      `ps:"Identity"`
	IncludeOrgContent                     bool     `ps:"IncludeOrgContent"`
	IncludeUserAppContent                 bool     `ps:"IncludeUserAppContent"`
	Language                              any      `ps:"Language"`
	Name                                  string   `ps:"Name"`
	OneDriveLocation                      []string `ps:"OneDriveLocation"`
	OneDriveLocationExclusion             []string `ps:"OneDriveLocationExclusion"`
	PublicFolderLocation                  []string `ps:"PublicFolderLocation"`
	PublicFolderLocationExclusion         []string `ps:"PublicFolderLocationExclusion"`
	RefinerNames                          []string `ps:"RefinerNames"`
	RemoveExchangeLocation                []string `ps:"RemoveExchangeLocation"`
	RemoveExchangeLocationExclusion       []string `ps:"RemoveExchangeLocationExclusion"`
	RemoveOneDriveLocation                []string `ps:"RemoveOneDriveLocation"`
	RemoveOneDriveLocationExclusion       []string `ps:"RemoveOneDriveLocationExclusion"`
	RemovePublicFolderLocation            []string `ps:"RemovePublicFolderLocation"`
	RemovePublicFolderLocationExclusion   []string `ps:"RemovePublicFolderLocationExclusion"`
	RemoveSharePointLocation              []string `ps:"RemoveSharePointLocation"`
	RemoveSharePointLocationExclusion     []string `ps:"RemoveSharePointLocationExclusion"`
	SearchNames                           []string `ps:"SearchNames"`
	SharePointLocation                    []string `ps:"SharePointLocation"`
	SharePointLocationExclusion           []string `ps:"SharePointLocationExclusion"`
}

func (p SetComplianceSearchParams) params() map[string]any {
	m := map[string]any{}
	if len(p.AddExchangeLocation) > 0 {
		m["AddExchangeLocation"] = p.AddExchangeLocation
	}
	if len(p.AddExchangeLocationExclusion) > 0 {
		m["AddExchangeLocationExclusion"] = p.AddExchangeLocationExclusion
	}
	if len(p.AddOneDriveLocation) > 0 {
		m["AddOneDriveLocation"] = p.AddOneDriveLocation
	}
	if len(p.AddOneDriveLocationExclusion) > 0 {
		m["AddOneDriveLocationExclusion"] = p.AddOneDriveLocationExclusion
	}
	if len(p.AddPublicFolderLocation) > 0 {
		m["AddPublicFolderLocation"] = p.AddPublicFolderLocation
	}
	if len(p.AddPublicFolderLocationExclusion) > 0 {
		m["AddPublicFolderLocationExclusion"] = p.AddPublicFolderLocationExclusion
	}
	if len(p.AddSharePointLocation) > 0 {
		m["AddSharePointLocation"] = p.AddSharePointLocation
	}
	if len(p.AddSharePointLocationExclusion) > 0 {
		m["AddSharePointLocationExclusion"] = p.AddSharePointLocationExclusion
	}
	if p.AllowNotFoundExchangeLocationsEnabled {
		m["AllowNotFoundExchangeLocationsEnabled"] = true
	}
	if p.ContentMatchQuery != "" {
		m["ContentMatchQuery"] = p.ContentMatchQuery
	}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if len(p.ExchangeLocation) > 0 {
		m["ExchangeLocation"] = p.ExchangeLocation
	}
	if len(p.ExchangeLocationExclusion) > 0 {
		m["ExchangeLocationExclusion"] = p.ExchangeLocationExclusion
	}
	if p.Force {
		m["Force"] = true
	}
	if len(p.HoldNames) > 0 {
		m["HoldNames"] = p.HoldNames
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IncludeOrgContent {
		m["IncludeOrgContent"] = true
	}
	if p.IncludeUserAppContent {
		m["IncludeUserAppContent"] = true
	}
	if p.Language != nil {
		m["Language"] = p.Language
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if len(p.OneDriveLocation) > 0 {
		m["OneDriveLocation"] = p.OneDriveLocation
	}
	if len(p.OneDriveLocationExclusion) > 0 {
		m["OneDriveLocationExclusion"] = p.OneDriveLocationExclusion
	}
	if len(p.PublicFolderLocation) > 0 {
		m["PublicFolderLocation"] = p.PublicFolderLocation
	}
	if len(p.PublicFolderLocationExclusion) > 0 {
		m["PublicFolderLocationExclusion"] = p.PublicFolderLocationExclusion
	}
	if len(p.RefinerNames) > 0 {
		m["RefinerNames"] = p.RefinerNames
	}
	if len(p.RemoveExchangeLocation) > 0 {
		m["RemoveExchangeLocation"] = p.RemoveExchangeLocation
	}
	if len(p.RemoveExchangeLocationExclusion) > 0 {
		m["RemoveExchangeLocationExclusion"] = p.RemoveExchangeLocationExclusion
	}
	if len(p.RemoveOneDriveLocation) > 0 {
		m["RemoveOneDriveLocation"] = p.RemoveOneDriveLocation
	}
	if len(p.RemoveOneDriveLocationExclusion) > 0 {
		m["RemoveOneDriveLocationExclusion"] = p.RemoveOneDriveLocationExclusion
	}
	if len(p.RemovePublicFolderLocation) > 0 {
		m["RemovePublicFolderLocation"] = p.RemovePublicFolderLocation
	}
	if len(p.RemovePublicFolderLocationExclusion) > 0 {
		m["RemovePublicFolderLocationExclusion"] = p.RemovePublicFolderLocationExclusion
	}
	if len(p.RemoveSharePointLocation) > 0 {
		m["RemoveSharePointLocation"] = p.RemoveSharePointLocation
	}
	if len(p.RemoveSharePointLocationExclusion) > 0 {
		m["RemoveSharePointLocationExclusion"] = p.RemoveSharePointLocationExclusion
	}
	if len(p.SearchNames) > 0 {
		m["SearchNames"] = p.SearchNames
	}
	if len(p.SharePointLocation) > 0 {
		m["SharePointLocation"] = p.SharePointLocation
	}
	if len(p.SharePointLocationExclusion) > 0 {
		m["SharePointLocationExclusion"] = p.SharePointLocationExclusion
	}
	return m
}

// SetComplianceSearch runs the Set-ComplianceSearch cmdlet.
func (s *Service) SetComplianceSearch(ctx context.Context, p SetComplianceSearchParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-ComplianceSearch", p.params())
}

// SetComplianceSearchActionParams are the parameters of Set-ComplianceSearchAction.
// DefaultParameterSetName: Identity
type SetComplianceSearchActionParams struct {
	ChangeExportKey bool `ps:"ChangeExportKey"`
	Identity        any  `ps:"Identity"`
}

func (p SetComplianceSearchActionParams) params() map[string]any {
	m := map[string]any{}
	if p.ChangeExportKey {
		m["ChangeExportKey"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// SetComplianceSearchAction runs the Set-ComplianceSearchAction cmdlet.
func (s *Service) SetComplianceSearchAction(ctx context.Context, p SetComplianceSearchActionParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-ComplianceSearchAction", p.params())
}

// SetComplianceTagParams are the parameters of Set-ComplianceTag.
// DefaultParameterSetName: Identity
type SetComplianceTagParams struct {
	ArchiveEnabled            bool     `ps:"ArchiveEnabled"`
	ArchiveTriggerBasedOn     string   `ps:"ArchiveTriggerBasedOn"`
	ArchiveTriggerInDays      any      `ps:"ArchiveTriggerInDays"`
	AutoApprovalPeriod        any      `ps:"AutoApprovalPeriod"`
	Comment                   string   `ps:"Comment"`
	ComplianceTagForNextStage string   `ps:"ComplianceTagForNextStage"`
	EventType                 any      `ps:"EventType"`
	FilePlanProperty          string   `ps:"FilePlanProperty"`
	FlowId                    any      `ps:"FlowId"`
	Force                     bool     `ps:"Force"`
	Identity                  any      `ps:"Identity"`
	MultiStageReviewProperty  string   `ps:"MultiStageReviewProperty"`
	Notes                     string   `ps:"Notes"`
	RetentionDuration         any      `ps:"RetentionDuration"`
	ReviewerEmail             []string `ps:"ReviewerEmail"`
	WebHookUrl                string   `ps:"WebHookUrl"`
}

func (p SetComplianceTagParams) params() map[string]any {
	m := map[string]any{}
	if p.ArchiveEnabled {
		m["ArchiveEnabled"] = true
	}
	if p.ArchiveTriggerBasedOn != "" {
		m["ArchiveTriggerBasedOn"] = p.ArchiveTriggerBasedOn
	}
	if p.ArchiveTriggerInDays != nil {
		m["ArchiveTriggerInDays"] = p.ArchiveTriggerInDays
	}
	if p.AutoApprovalPeriod != nil {
		m["AutoApprovalPeriod"] = p.AutoApprovalPeriod
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.ComplianceTagForNextStage != "" {
		m["ComplianceTagForNextStage"] = p.ComplianceTagForNextStage
	}
	if p.EventType != nil {
		m["EventType"] = p.EventType
	}
	if p.FilePlanProperty != "" {
		m["FilePlanProperty"] = p.FilePlanProperty
	}
	if p.FlowId != nil {
		m["FlowId"] = p.FlowId
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.MultiStageReviewProperty != "" {
		m["MultiStageReviewProperty"] = p.MultiStageReviewProperty
	}
	if p.Notes != "" {
		m["Notes"] = p.Notes
	}
	if p.RetentionDuration != nil {
		m["RetentionDuration"] = p.RetentionDuration
	}
	if len(p.ReviewerEmail) > 0 {
		m["ReviewerEmail"] = p.ReviewerEmail
	}
	if p.WebHookUrl != "" {
		m["WebHookUrl"] = p.WebHookUrl
	}
	return m
}

// SetComplianceTag runs the Set-ComplianceTag cmdlet.
func (s *Service) SetComplianceTag(ctx context.Context, p SetComplianceTagParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-ComplianceTag", p.params())
}

// SetCustomDlpEmailTemplateParams are the parameters of Set-CustomDlpEmailTemplate.
// DefaultParameterSetName: Identity
type SetCustomDlpEmailTemplateParams struct {
	Bcc               []string `ps:"Bcc"`
	Body              string   `ps:"Body"`
	Cc                []string `ps:"Cc"`
	Description       string   `ps:"Description"`
	From              any      `ps:"From"`
	Identity          any      `ps:"Identity"`
	Importance        any      `ps:"Importance"`
	NotificationTitle string   `ps:"NotificationTitle"`
	Subject           string   `ps:"Subject"`
	To                []string `ps:"To"`
}

func (p SetCustomDlpEmailTemplateParams) params() map[string]any {
	m := map[string]any{}
	if len(p.Bcc) > 0 {
		m["Bcc"] = p.Bcc
	}
	if p.Body != "" {
		m["Body"] = p.Body
	}
	if len(p.Cc) > 0 {
		m["Cc"] = p.Cc
	}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if p.From != nil {
		m["From"] = p.From
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Importance != nil {
		m["Importance"] = p.Importance
	}
	if p.NotificationTitle != "" {
		m["NotificationTitle"] = p.NotificationTitle
	}
	if p.Subject != "" {
		m["Subject"] = p.Subject
	}
	if len(p.To) > 0 {
		m["To"] = p.To
	}
	return m
}

// SetCustomDlpEmailTemplate runs the Set-CustomDlpEmailTemplate cmdlet.
func (s *Service) SetCustomDlpEmailTemplate(ctx context.Context, p SetCustomDlpEmailTemplateParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-CustomDlpEmailTemplate", p.params())
}

// SetDeviceConditionalAccessPolicyParams are the parameters of Set-DeviceConditionalAccessPolicy.
// DefaultParameterSetName: Identity
type SetDeviceConditionalAccessPolicyParams struct {
	Comment           string `ps:"Comment"`
	Enabled           bool   `ps:"Enabled"`
	Force             bool   `ps:"Force"`
	Identity          any    `ps:"Identity"`
	RetryDistribution bool   `ps:"RetryDistribution"`
}

func (p SetDeviceConditionalAccessPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.RetryDistribution {
		m["RetryDistribution"] = true
	}
	return m
}

// SetDeviceConditionalAccessPolicy runs the Set-DeviceConditionalAccessPolicy cmdlet.
func (s *Service) SetDeviceConditionalAccessPolicy(ctx context.Context, p SetDeviceConditionalAccessPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-DeviceConditionalAccessPolicy", p.params())
}

// SetDeviceConditionalAccessRuleParams are the parameters of Set-DeviceConditionalAccessRule.
// DefaultParameterSetName: Identity
type SetDeviceConditionalAccessRuleParams struct {
	AccountName                   string `ps:"AccountName"`
	AccountUserName               string `ps:"AccountUserName"`
	AllowAppStore                 any    `ps:"AllowAppStore"`
	AllowAssistantWhileLocked     any    `ps:"AllowAssistantWhileLocked"`
	AllowConvenienceLogon         any    `ps:"AllowConvenienceLogon"`
	AllowDiagnosticSubmission     any    `ps:"AllowDiagnosticSubmission"`
	AllowiCloudBackup             any    `ps:"AllowiCloudBackup"`
	AllowiCloudDocSync            any    `ps:"AllowiCloudDocSync"`
	AllowiCloudPhotoSync          any    `ps:"AllowiCloudPhotoSync"`
	AllowJailbroken               any    `ps:"AllowJailbroken"`
	AllowPassbookWhileLocked      any    `ps:"AllowPassbookWhileLocked"`
	AllowScreenshot               any    `ps:"AllowScreenshot"`
	AllowSimplePassword           any    `ps:"AllowSimplePassword"`
	AllowVideoConferencing        any    `ps:"AllowVideoConferencing"`
	AllowVoiceAssistant           any    `ps:"AllowVoiceAssistant"`
	AllowVoiceDialing             any    `ps:"AllowVoiceDialing"`
	AntiVirusSignatureStatus      any    `ps:"AntiVirusSignatureStatus"`
	AntiVirusStatus               any    `ps:"AntiVirusStatus"`
	AppsRating                    any    `ps:"AppsRating"`
	AutoUpdateStatus              any    `ps:"AutoUpdateStatus"`
	BluetoothEnabled              any    `ps:"BluetoothEnabled"`
	CameraEnabled                 any    `ps:"CameraEnabled"`
	DomainController              any    `ps:"DomainController"`
	EmailAddress                  string `ps:"EmailAddress"`
	EnableRemovableStorage        any    `ps:"EnableRemovableStorage"`
	ExchangeActiveSyncHost        string `ps:"ExchangeActiveSyncHost"`
	FirewallStatus                any    `ps:"FirewallStatus"`
	ForceAppStorePassword         any    `ps:"ForceAppStorePassword"`
	ForceEncryptedBackup          any    `ps:"ForceEncryptedBackup"`
	Identity                      any    `ps:"Identity"`
	MaxPasswordAttemptsBeforeWipe any    `ps:"MaxPasswordAttemptsBeforeWipe"`
	MaxPasswordGracePeriod        any    `ps:"MaxPasswordGracePeriod"`
	MoviesRating                  any    `ps:"MoviesRating"`
	PasswordComplexity            any    `ps:"PasswordComplexity"`
	PasswordExpirationDays        any    `ps:"PasswordExpirationDays"`
	PasswordHistoryCount          any    `ps:"PasswordHistoryCount"`
	PasswordMinComplexChars       any    `ps:"PasswordMinComplexChars"`
	PasswordMinimumLength         any    `ps:"PasswordMinimumLength"`
	PasswordQuality               any    `ps:"PasswordQuality"`
	PasswordRequired              any    `ps:"PasswordRequired"`
	PasswordTimeout               any    `ps:"PasswordTimeout"`
	PhoneMemoryEncrypted          any    `ps:"PhoneMemoryEncrypted"`
	RegionRatings                 any    `ps:"RegionRatings"`
	RequireEmailProfile           any    `ps:"RequireEmailProfile"`
	SmartScreenEnabled            any    `ps:"SmartScreenEnabled"`
	SystemSecurityTLS             any    `ps:"SystemSecurityTLS"`
	TargetGroups                  any    `ps:"TargetGroups"`
	TVShowsRating                 any    `ps:"TVShowsRating"`
	UserAccountControlStatus      any    `ps:"UserAccountControlStatus"`
	WLANEnabled                   any    `ps:"WLANEnabled"`
	WorkFoldersSyncUrl            string `ps:"WorkFoldersSyncUrl"`
}

func (p SetDeviceConditionalAccessRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.AccountName != "" {
		m["AccountName"] = p.AccountName
	}
	if p.AccountUserName != "" {
		m["AccountUserName"] = p.AccountUserName
	}
	if p.AllowAppStore != nil {
		m["AllowAppStore"] = p.AllowAppStore
	}
	if p.AllowAssistantWhileLocked != nil {
		m["AllowAssistantWhileLocked"] = p.AllowAssistantWhileLocked
	}
	if p.AllowConvenienceLogon != nil {
		m["AllowConvenienceLogon"] = p.AllowConvenienceLogon
	}
	if p.AllowDiagnosticSubmission != nil {
		m["AllowDiagnosticSubmission"] = p.AllowDiagnosticSubmission
	}
	if p.AllowiCloudBackup != nil {
		m["AllowiCloudBackup"] = p.AllowiCloudBackup
	}
	if p.AllowiCloudDocSync != nil {
		m["AllowiCloudDocSync"] = p.AllowiCloudDocSync
	}
	if p.AllowiCloudPhotoSync != nil {
		m["AllowiCloudPhotoSync"] = p.AllowiCloudPhotoSync
	}
	if p.AllowJailbroken != nil {
		m["AllowJailbroken"] = p.AllowJailbroken
	}
	if p.AllowPassbookWhileLocked != nil {
		m["AllowPassbookWhileLocked"] = p.AllowPassbookWhileLocked
	}
	if p.AllowScreenshot != nil {
		m["AllowScreenshot"] = p.AllowScreenshot
	}
	if p.AllowSimplePassword != nil {
		m["AllowSimplePassword"] = p.AllowSimplePassword
	}
	if p.AllowVideoConferencing != nil {
		m["AllowVideoConferencing"] = p.AllowVideoConferencing
	}
	if p.AllowVoiceAssistant != nil {
		m["AllowVoiceAssistant"] = p.AllowVoiceAssistant
	}
	if p.AllowVoiceDialing != nil {
		m["AllowVoiceDialing"] = p.AllowVoiceDialing
	}
	if p.AntiVirusSignatureStatus != nil {
		m["AntiVirusSignatureStatus"] = p.AntiVirusSignatureStatus
	}
	if p.AntiVirusStatus != nil {
		m["AntiVirusStatus"] = p.AntiVirusStatus
	}
	if p.AppsRating != nil {
		m["AppsRating"] = p.AppsRating
	}
	if p.AutoUpdateStatus != nil {
		m["AutoUpdateStatus"] = p.AutoUpdateStatus
	}
	if p.BluetoothEnabled != nil {
		m["BluetoothEnabled"] = p.BluetoothEnabled
	}
	if p.CameraEnabled != nil {
		m["CameraEnabled"] = p.CameraEnabled
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.EmailAddress != "" {
		m["EmailAddress"] = p.EmailAddress
	}
	if p.EnableRemovableStorage != nil {
		m["EnableRemovableStorage"] = p.EnableRemovableStorage
	}
	if p.ExchangeActiveSyncHost != "" {
		m["ExchangeActiveSyncHost"] = p.ExchangeActiveSyncHost
	}
	if p.FirewallStatus != nil {
		m["FirewallStatus"] = p.FirewallStatus
	}
	if p.ForceAppStorePassword != nil {
		m["ForceAppStorePassword"] = p.ForceAppStorePassword
	}
	if p.ForceEncryptedBackup != nil {
		m["ForceEncryptedBackup"] = p.ForceEncryptedBackup
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.MaxPasswordAttemptsBeforeWipe != nil {
		m["MaxPasswordAttemptsBeforeWipe"] = p.MaxPasswordAttemptsBeforeWipe
	}
	if p.MaxPasswordGracePeriod != nil {
		m["MaxPasswordGracePeriod"] = p.MaxPasswordGracePeriod
	}
	if p.MoviesRating != nil {
		m["MoviesRating"] = p.MoviesRating
	}
	if p.PasswordComplexity != nil {
		m["PasswordComplexity"] = p.PasswordComplexity
	}
	if p.PasswordExpirationDays != nil {
		m["PasswordExpirationDays"] = p.PasswordExpirationDays
	}
	if p.PasswordHistoryCount != nil {
		m["PasswordHistoryCount"] = p.PasswordHistoryCount
	}
	if p.PasswordMinComplexChars != nil {
		m["PasswordMinComplexChars"] = p.PasswordMinComplexChars
	}
	if p.PasswordMinimumLength != nil {
		m["PasswordMinimumLength"] = p.PasswordMinimumLength
	}
	if p.PasswordQuality != nil {
		m["PasswordQuality"] = p.PasswordQuality
	}
	if p.PasswordRequired != nil {
		m["PasswordRequired"] = p.PasswordRequired
	}
	if p.PasswordTimeout != nil {
		m["PasswordTimeout"] = p.PasswordTimeout
	}
	if p.PhoneMemoryEncrypted != nil {
		m["PhoneMemoryEncrypted"] = p.PhoneMemoryEncrypted
	}
	if p.RegionRatings != nil {
		m["RegionRatings"] = p.RegionRatings
	}
	if p.RequireEmailProfile != nil {
		m["RequireEmailProfile"] = p.RequireEmailProfile
	}
	if p.SmartScreenEnabled != nil {
		m["SmartScreenEnabled"] = p.SmartScreenEnabled
	}
	if p.SystemSecurityTLS != nil {
		m["SystemSecurityTLS"] = p.SystemSecurityTLS
	}
	if p.TargetGroups != nil {
		m["TargetGroups"] = p.TargetGroups
	}
	if p.TVShowsRating != nil {
		m["TVShowsRating"] = p.TVShowsRating
	}
	if p.UserAccountControlStatus != nil {
		m["UserAccountControlStatus"] = p.UserAccountControlStatus
	}
	if p.WLANEnabled != nil {
		m["WLANEnabled"] = p.WLANEnabled
	}
	if p.WorkFoldersSyncUrl != "" {
		m["WorkFoldersSyncUrl"] = p.WorkFoldersSyncUrl
	}
	return m
}

// SetDeviceConditionalAccessRule runs the Set-DeviceConditionalAccessRule cmdlet.
func (s *Service) SetDeviceConditionalAccessRule(ctx context.Context, p SetDeviceConditionalAccessRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-DeviceConditionalAccessRule", p.params())
}

// SetDeviceConfigurationPolicyParams are the parameters of Set-DeviceConfigurationPolicy.
// DefaultParameterSetName: Identity
type SetDeviceConfigurationPolicyParams struct {
	Comment           string `ps:"Comment"`
	Enabled           bool   `ps:"Enabled"`
	Force             bool   `ps:"Force"`
	Identity          any    `ps:"Identity"`
	RetryDistribution bool   `ps:"RetryDistribution"`
}

func (p SetDeviceConfigurationPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.RetryDistribution {
		m["RetryDistribution"] = true
	}
	return m
}

// SetDeviceConfigurationPolicy runs the Set-DeviceConfigurationPolicy cmdlet.
func (s *Service) SetDeviceConfigurationPolicy(ctx context.Context, p SetDeviceConfigurationPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-DeviceConfigurationPolicy", p.params())
}

// SetDeviceConfigurationRuleParams are the parameters of Set-DeviceConfigurationRule.
// DefaultParameterSetName: Identity
type SetDeviceConfigurationRuleParams struct {
	AccountName                   string `ps:"AccountName"`
	AccountUserName               string `ps:"AccountUserName"`
	AllowAppStore                 any    `ps:"AllowAppStore"`
	AllowAssistantWhileLocked     any    `ps:"AllowAssistantWhileLocked"`
	AllowConvenienceLogon         any    `ps:"AllowConvenienceLogon"`
	AllowDiagnosticSubmission     any    `ps:"AllowDiagnosticSubmission"`
	AllowiCloudBackup             any    `ps:"AllowiCloudBackup"`
	AllowiCloudDocSync            any    `ps:"AllowiCloudDocSync"`
	AllowiCloudPhotoSync          any    `ps:"AllowiCloudPhotoSync"`
	AllowPassbookWhileLocked      any    `ps:"AllowPassbookWhileLocked"`
	AllowScreenshot               any    `ps:"AllowScreenshot"`
	AllowSimplePassword           any    `ps:"AllowSimplePassword"`
	AllowVideoConferencing        any    `ps:"AllowVideoConferencing"`
	AllowVoiceAssistant           any    `ps:"AllowVoiceAssistant"`
	AllowVoiceDialing             any    `ps:"AllowVoiceDialing"`
	AntiVirusSignatureStatus      any    `ps:"AntiVirusSignatureStatus"`
	AntiVirusStatus               any    `ps:"AntiVirusStatus"`
	AppsRating                    any    `ps:"AppsRating"`
	AutoUpdateStatus              any    `ps:"AutoUpdateStatus"`
	BluetoothEnabled              any    `ps:"BluetoothEnabled"`
	CameraEnabled                 any    `ps:"CameraEnabled"`
	DomainController              any    `ps:"DomainController"`
	EmailAddress                  string `ps:"EmailAddress"`
	EnableRemovableStorage        any    `ps:"EnableRemovableStorage"`
	ExchangeActiveSyncHost        string `ps:"ExchangeActiveSyncHost"`
	FirewallStatus                any    `ps:"FirewallStatus"`
	ForceAppStorePassword         any    `ps:"ForceAppStorePassword"`
	ForceEncryptedBackup          any    `ps:"ForceEncryptedBackup"`
	Identity                      any    `ps:"Identity"`
	MaxPasswordAttemptsBeforeWipe any    `ps:"MaxPasswordAttemptsBeforeWipe"`
	MaxPasswordGracePeriod        any    `ps:"MaxPasswordGracePeriod"`
	MoviesRating                  any    `ps:"MoviesRating"`
	PasswordComplexity            any    `ps:"PasswordComplexity"`
	PasswordExpirationDays        any    `ps:"PasswordExpirationDays"`
	PasswordHistoryCount          any    `ps:"PasswordHistoryCount"`
	PasswordMinComplexChars       any    `ps:"PasswordMinComplexChars"`
	PasswordMinimumLength         any    `ps:"PasswordMinimumLength"`
	PasswordQuality               any    `ps:"PasswordQuality"`
	PasswordRequired              any    `ps:"PasswordRequired"`
	PasswordTimeout               any    `ps:"PasswordTimeout"`
	PhoneMemoryEncrypted          any    `ps:"PhoneMemoryEncrypted"`
	RegionRatings                 any    `ps:"RegionRatings"`
	RequireEmailProfile           any    `ps:"RequireEmailProfile"`
	SmartScreenEnabled            any    `ps:"SmartScreenEnabled"`
	SystemSecurityTLS             any    `ps:"SystemSecurityTLS"`
	TargetGroups                  any    `ps:"TargetGroups"`
	TVShowsRating                 any    `ps:"TVShowsRating"`
	UserAccountControlStatus      any    `ps:"UserAccountControlStatus"`
	WLANEnabled                   any    `ps:"WLANEnabled"`
	WorkFoldersSyncUrl            string `ps:"WorkFoldersSyncUrl"`
}

func (p SetDeviceConfigurationRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.AccountName != "" {
		m["AccountName"] = p.AccountName
	}
	if p.AccountUserName != "" {
		m["AccountUserName"] = p.AccountUserName
	}
	if p.AllowAppStore != nil {
		m["AllowAppStore"] = p.AllowAppStore
	}
	if p.AllowAssistantWhileLocked != nil {
		m["AllowAssistantWhileLocked"] = p.AllowAssistantWhileLocked
	}
	if p.AllowConvenienceLogon != nil {
		m["AllowConvenienceLogon"] = p.AllowConvenienceLogon
	}
	if p.AllowDiagnosticSubmission != nil {
		m["AllowDiagnosticSubmission"] = p.AllowDiagnosticSubmission
	}
	if p.AllowiCloudBackup != nil {
		m["AllowiCloudBackup"] = p.AllowiCloudBackup
	}
	if p.AllowiCloudDocSync != nil {
		m["AllowiCloudDocSync"] = p.AllowiCloudDocSync
	}
	if p.AllowiCloudPhotoSync != nil {
		m["AllowiCloudPhotoSync"] = p.AllowiCloudPhotoSync
	}
	if p.AllowPassbookWhileLocked != nil {
		m["AllowPassbookWhileLocked"] = p.AllowPassbookWhileLocked
	}
	if p.AllowScreenshot != nil {
		m["AllowScreenshot"] = p.AllowScreenshot
	}
	if p.AllowSimplePassword != nil {
		m["AllowSimplePassword"] = p.AllowSimplePassword
	}
	if p.AllowVideoConferencing != nil {
		m["AllowVideoConferencing"] = p.AllowVideoConferencing
	}
	if p.AllowVoiceAssistant != nil {
		m["AllowVoiceAssistant"] = p.AllowVoiceAssistant
	}
	if p.AllowVoiceDialing != nil {
		m["AllowVoiceDialing"] = p.AllowVoiceDialing
	}
	if p.AntiVirusSignatureStatus != nil {
		m["AntiVirusSignatureStatus"] = p.AntiVirusSignatureStatus
	}
	if p.AntiVirusStatus != nil {
		m["AntiVirusStatus"] = p.AntiVirusStatus
	}
	if p.AppsRating != nil {
		m["AppsRating"] = p.AppsRating
	}
	if p.AutoUpdateStatus != nil {
		m["AutoUpdateStatus"] = p.AutoUpdateStatus
	}
	if p.BluetoothEnabled != nil {
		m["BluetoothEnabled"] = p.BluetoothEnabled
	}
	if p.CameraEnabled != nil {
		m["CameraEnabled"] = p.CameraEnabled
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.EmailAddress != "" {
		m["EmailAddress"] = p.EmailAddress
	}
	if p.EnableRemovableStorage != nil {
		m["EnableRemovableStorage"] = p.EnableRemovableStorage
	}
	if p.ExchangeActiveSyncHost != "" {
		m["ExchangeActiveSyncHost"] = p.ExchangeActiveSyncHost
	}
	if p.FirewallStatus != nil {
		m["FirewallStatus"] = p.FirewallStatus
	}
	if p.ForceAppStorePassword != nil {
		m["ForceAppStorePassword"] = p.ForceAppStorePassword
	}
	if p.ForceEncryptedBackup != nil {
		m["ForceEncryptedBackup"] = p.ForceEncryptedBackup
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.MaxPasswordAttemptsBeforeWipe != nil {
		m["MaxPasswordAttemptsBeforeWipe"] = p.MaxPasswordAttemptsBeforeWipe
	}
	if p.MaxPasswordGracePeriod != nil {
		m["MaxPasswordGracePeriod"] = p.MaxPasswordGracePeriod
	}
	if p.MoviesRating != nil {
		m["MoviesRating"] = p.MoviesRating
	}
	if p.PasswordComplexity != nil {
		m["PasswordComplexity"] = p.PasswordComplexity
	}
	if p.PasswordExpirationDays != nil {
		m["PasswordExpirationDays"] = p.PasswordExpirationDays
	}
	if p.PasswordHistoryCount != nil {
		m["PasswordHistoryCount"] = p.PasswordHistoryCount
	}
	if p.PasswordMinComplexChars != nil {
		m["PasswordMinComplexChars"] = p.PasswordMinComplexChars
	}
	if p.PasswordMinimumLength != nil {
		m["PasswordMinimumLength"] = p.PasswordMinimumLength
	}
	if p.PasswordQuality != nil {
		m["PasswordQuality"] = p.PasswordQuality
	}
	if p.PasswordRequired != nil {
		m["PasswordRequired"] = p.PasswordRequired
	}
	if p.PasswordTimeout != nil {
		m["PasswordTimeout"] = p.PasswordTimeout
	}
	if p.PhoneMemoryEncrypted != nil {
		m["PhoneMemoryEncrypted"] = p.PhoneMemoryEncrypted
	}
	if p.RegionRatings != nil {
		m["RegionRatings"] = p.RegionRatings
	}
	if p.RequireEmailProfile != nil {
		m["RequireEmailProfile"] = p.RequireEmailProfile
	}
	if p.SmartScreenEnabled != nil {
		m["SmartScreenEnabled"] = p.SmartScreenEnabled
	}
	if p.SystemSecurityTLS != nil {
		m["SystemSecurityTLS"] = p.SystemSecurityTLS
	}
	if p.TargetGroups != nil {
		m["TargetGroups"] = p.TargetGroups
	}
	if p.TVShowsRating != nil {
		m["TVShowsRating"] = p.TVShowsRating
	}
	if p.UserAccountControlStatus != nil {
		m["UserAccountControlStatus"] = p.UserAccountControlStatus
	}
	if p.WLANEnabled != nil {
		m["WLANEnabled"] = p.WLANEnabled
	}
	if p.WorkFoldersSyncUrl != "" {
		m["WorkFoldersSyncUrl"] = p.WorkFoldersSyncUrl
	}
	return m
}

// SetDeviceConfigurationRule runs the Set-DeviceConfigurationRule cmdlet.
func (s *Service) SetDeviceConfigurationRule(ctx context.Context, p SetDeviceConfigurationRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-DeviceConfigurationRule", p.params())
}

// SetDeviceTenantPolicyParams are the parameters of Set-DeviceTenantPolicy.
// DefaultParameterSetName: Identity
type SetDeviceTenantPolicyParams struct {
	Comment           string `ps:"Comment"`
	Enabled           bool   `ps:"Enabled"`
	Force             bool   `ps:"Force"`
	Identity          any    `ps:"Identity"`
	RetryDistribution bool   `ps:"RetryDistribution"`
}

func (p SetDeviceTenantPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.RetryDistribution {
		m["RetryDistribution"] = true
	}
	return m
}

// SetDeviceTenantPolicy runs the Set-DeviceTenantPolicy cmdlet.
func (s *Service) SetDeviceTenantPolicy(ctx context.Context, p SetDeviceTenantPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-DeviceTenantPolicy", p.params())
}

// SetDeviceTenantRuleParams are the parameters of Set-DeviceTenantRule.
// DefaultParameterSetName: Identity
type SetDeviceTenantRuleParams struct {
	ApplyPolicyTo           any `ps:"ApplyPolicyTo"`
	BlockUnsupportedDevices any `ps:"BlockUnsupportedDevices"`
	DomainController        any `ps:"DomainController"`
	ExclusionList           any `ps:"ExclusionList"`
	Identity                any `ps:"Identity"`
}

func (p SetDeviceTenantRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.ApplyPolicyTo != nil {
		m["ApplyPolicyTo"] = p.ApplyPolicyTo
	}
	if p.BlockUnsupportedDevices != nil {
		m["BlockUnsupportedDevices"] = p.BlockUnsupportedDevices
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.ExclusionList != nil {
		m["ExclusionList"] = p.ExclusionList
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// SetDeviceTenantRule runs the Set-DeviceTenantRule cmdlet.
func (s *Service) SetDeviceTenantRule(ctx context.Context, p SetDeviceTenantRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-DeviceTenantRule", p.params())
}

// SetDlpAlertTuningRuleParams are the parameters of Set-DlpAlertTuningRule.
// DefaultParameterSetName: Identity
type SetDlpAlertTuningRuleParams struct {
	Actions        string `ps:"Actions"`
	ApplicableTags string `ps:"ApplicableTags"`
	Conditions     string `ps:"Conditions"`
	Description    string `ps:"Description"`
	Identity       any    `ps:"Identity"`
}

func (p SetDlpAlertTuningRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Actions != "" {
		m["Actions"] = p.Actions
	}
	if p.ApplicableTags != "" {
		m["ApplicableTags"] = p.ApplicableTags
	}
	if p.Conditions != "" {
		m["Conditions"] = p.Conditions
	}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// SetDlpAlertTuningRule runs the Set-DlpAlertTuningRule cmdlet.
func (s *Service) SetDlpAlertTuningRule(ctx context.Context, p SetDlpAlertTuningRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-DlpAlertTuningRule", p.params())
}

// SetDlpCompliancePolicyParams are the parameters of Set-DlpCompliancePolicy.
// DefaultParameterSetName: Identity
type SetDlpCompliancePolicyParams struct {
	AddEndpointDlpLocation                      any      `ps:"AddEndpointDlpLocation"`
	AddEndpointDlpLocationException             any      `ps:"AddEndpointDlpLocationException"`
	AddExchangeLocation                         any      `ps:"AddExchangeLocation"`
	AddOneDriveLocation                         any      `ps:"AddOneDriveLocation"`
	AddOneDriveLocationException                any      `ps:"AddOneDriveLocationException"`
	AddOnPremisesScannerDlpLocation             any      `ps:"AddOnPremisesScannerDlpLocation"`
	AddOnPremisesScannerDlpLocationException    any      `ps:"AddOnPremisesScannerDlpLocationException"`
	AddPowerBIDlpLocation                       any      `ps:"AddPowerBIDlpLocation"`
	AddPowerBIDlpLocationException              any      `ps:"AddPowerBIDlpLocationException"`
	AddSharePointLocation                       any      `ps:"AddSharePointLocation"`
	AddSharePointLocationException              any      `ps:"AddSharePointLocationException"`
	AddTeamsLocation                            any      `ps:"AddTeamsLocation"`
	AddTeamsLocationException                   any      `ps:"AddTeamsLocationException"`
	AddThirdPartyAppDlpLocation                 any      `ps:"AddThirdPartyAppDlpLocation"`
	AddThirdPartyAppDlpLocationException        any      `ps:"AddThirdPartyAppDlpLocationException"`
	AutoEnableAfter                             any      `ps:"AutoEnableAfter"`
	Comment                                     string   `ps:"Comment"`
	DisplayName                                 string   `ps:"DisplayName"`
	EndpointDlpAdaptiveScopes                   any      `ps:"EndpointDlpAdaptiveScopes"`
	EndpointDlpAdaptiveScopesException          any      `ps:"EndpointDlpAdaptiveScopesException"`
	EndpointDlpExtendedLocations                string   `ps:"EndpointDlpExtendedLocations"`
	EnforcementPlanes                           any      `ps:"EnforcementPlanes"`
	ExceptIfOneDriveSharedBy                    []string `ps:"ExceptIfOneDriveSharedBy"`
	ExceptIfOneDriveSharedByMemberOf            []string `ps:"ExceptIfOneDriveSharedByMemberOf"`
	ExchangeAdaptiveScopes                      any      `ps:"ExchangeAdaptiveScopes"`
	ExchangeAdaptiveScopesException             any      `ps:"ExchangeAdaptiveScopesException"`
	ExchangeSender                              []string `ps:"ExchangeSender"`
	ExchangeSenderException                     []string `ps:"ExchangeSenderException"`
	ExchangeSenderMemberOf                      []string `ps:"ExchangeSenderMemberOf"`
	ExchangeSenderMemberOfException             []string `ps:"ExchangeSenderMemberOfException"`
	Force                                       bool     `ps:"Force"`
	Identity                                    any      `ps:"Identity"`
	IsFromSmartInsights                         any      `ps:"IsFromSmartInsights"`
	Locations                                   string   `ps:"Locations"`
	Mode                                        any      `ps:"Mode"`
	NewName                                     string   `ps:"NewName"`
	OneDriveAdaptiveScopes                      any      `ps:"OneDriveAdaptiveScopes"`
	OneDriveAdaptiveScopesException             any      `ps:"OneDriveAdaptiveScopesException"`
	OneDriveSharedBy                            []string `ps:"OneDriveSharedBy"`
	OneDriveSharedByMemberOf                    []string `ps:"OneDriveSharedByMemberOf"`
	PolicyRBACScopes                            any      `ps:"PolicyRBACScopes"`
	PolicyTemplateInfo                          any      `ps:"PolicyTemplateInfo"`
	Priority                                    any      `ps:"Priority"`
	RemoveEndpointDlpLocation                   any      `ps:"RemoveEndpointDlpLocation"`
	RemoveEndpointDlpLocationException          any      `ps:"RemoveEndpointDlpLocationException"`
	RemoveExchangeLocation                      any      `ps:"RemoveExchangeLocation"`
	RemoveOneDriveLocation                      any      `ps:"RemoveOneDriveLocation"`
	RemoveOneDriveLocationException             any      `ps:"RemoveOneDriveLocationException"`
	RemoveOnPremisesScannerDlpLocation          any      `ps:"RemoveOnPremisesScannerDlpLocation"`
	RemoveOnPremisesScannerDlpLocationException any      `ps:"RemoveOnPremisesScannerDlpLocationException"`
	RemovePowerBIDlpLocation                    any      `ps:"RemovePowerBIDlpLocation"`
	RemovePowerBIDlpLocationException           any      `ps:"RemovePowerBIDlpLocationException"`
	RemoveSharePointLocation                    any      `ps:"RemoveSharePointLocation"`
	RemoveSharePointLocationException           any      `ps:"RemoveSharePointLocationException"`
	RemoveTeamsLocation                         any      `ps:"RemoveTeamsLocation"`
	RemoveTeamsLocationException                any      `ps:"RemoveTeamsLocationException"`
	RemoveThirdPartyAppDlpLocation              any      `ps:"RemoveThirdPartyAppDlpLocation"`
	RemoveThirdPartyAppDlpLocationException     any      `ps:"RemoveThirdPartyAppDlpLocationException"`
	RetryDistribution                           bool     `ps:"RetryDistribution"`
	SharePointAdaptiveScopes                    any      `ps:"SharePointAdaptiveScopes"`
	SharePointAdaptiveScopesException           any      `ps:"SharePointAdaptiveScopesException"`
	StartSimulation                             bool     `ps:"StartSimulation"`
	TeamsAdaptiveScopes                         any      `ps:"TeamsAdaptiveScopes"`
	TeamsAdaptiveScopesException                any      `ps:"TeamsAdaptiveScopesException"`
}

func (p SetDlpCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.AddEndpointDlpLocation != nil {
		m["AddEndpointDlpLocation"] = p.AddEndpointDlpLocation
	}
	if p.AddEndpointDlpLocationException != nil {
		m["AddEndpointDlpLocationException"] = p.AddEndpointDlpLocationException
	}
	if p.AddExchangeLocation != nil {
		m["AddExchangeLocation"] = p.AddExchangeLocation
	}
	if p.AddOneDriveLocation != nil {
		m["AddOneDriveLocation"] = p.AddOneDriveLocation
	}
	if p.AddOneDriveLocationException != nil {
		m["AddOneDriveLocationException"] = p.AddOneDriveLocationException
	}
	if p.AddOnPremisesScannerDlpLocation != nil {
		m["AddOnPremisesScannerDlpLocation"] = p.AddOnPremisesScannerDlpLocation
	}
	if p.AddOnPremisesScannerDlpLocationException != nil {
		m["AddOnPremisesScannerDlpLocationException"] = p.AddOnPremisesScannerDlpLocationException
	}
	if p.AddPowerBIDlpLocation != nil {
		m["AddPowerBIDlpLocation"] = p.AddPowerBIDlpLocation
	}
	if p.AddPowerBIDlpLocationException != nil {
		m["AddPowerBIDlpLocationException"] = p.AddPowerBIDlpLocationException
	}
	if p.AddSharePointLocation != nil {
		m["AddSharePointLocation"] = p.AddSharePointLocation
	}
	if p.AddSharePointLocationException != nil {
		m["AddSharePointLocationException"] = p.AddSharePointLocationException
	}
	if p.AddTeamsLocation != nil {
		m["AddTeamsLocation"] = p.AddTeamsLocation
	}
	if p.AddTeamsLocationException != nil {
		m["AddTeamsLocationException"] = p.AddTeamsLocationException
	}
	if p.AddThirdPartyAppDlpLocation != nil {
		m["AddThirdPartyAppDlpLocation"] = p.AddThirdPartyAppDlpLocation
	}
	if p.AddThirdPartyAppDlpLocationException != nil {
		m["AddThirdPartyAppDlpLocationException"] = p.AddThirdPartyAppDlpLocationException
	}
	if p.AutoEnableAfter != nil {
		m["AutoEnableAfter"] = p.AutoEnableAfter
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.EndpointDlpAdaptiveScopes != nil {
		m["EndpointDlpAdaptiveScopes"] = p.EndpointDlpAdaptiveScopes
	}
	if p.EndpointDlpAdaptiveScopesException != nil {
		m["EndpointDlpAdaptiveScopesException"] = p.EndpointDlpAdaptiveScopesException
	}
	if p.EndpointDlpExtendedLocations != "" {
		m["EndpointDlpExtendedLocations"] = p.EndpointDlpExtendedLocations
	}
	if p.EnforcementPlanes != nil {
		m["EnforcementPlanes"] = p.EnforcementPlanes
	}
	if len(p.ExceptIfOneDriveSharedBy) > 0 {
		m["ExceptIfOneDriveSharedBy"] = p.ExceptIfOneDriveSharedBy
	}
	if len(p.ExceptIfOneDriveSharedByMemberOf) > 0 {
		m["ExceptIfOneDriveSharedByMemberOf"] = p.ExceptIfOneDriveSharedByMemberOf
	}
	if p.ExchangeAdaptiveScopes != nil {
		m["ExchangeAdaptiveScopes"] = p.ExchangeAdaptiveScopes
	}
	if p.ExchangeAdaptiveScopesException != nil {
		m["ExchangeAdaptiveScopesException"] = p.ExchangeAdaptiveScopesException
	}
	if len(p.ExchangeSender) > 0 {
		m["ExchangeSender"] = p.ExchangeSender
	}
	if len(p.ExchangeSenderException) > 0 {
		m["ExchangeSenderException"] = p.ExchangeSenderException
	}
	if len(p.ExchangeSenderMemberOf) > 0 {
		m["ExchangeSenderMemberOf"] = p.ExchangeSenderMemberOf
	}
	if len(p.ExchangeSenderMemberOfException) > 0 {
		m["ExchangeSenderMemberOfException"] = p.ExchangeSenderMemberOfException
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IsFromSmartInsights != nil {
		m["IsFromSmartInsights"] = p.IsFromSmartInsights
	}
	if p.Locations != "" {
		m["Locations"] = p.Locations
	}
	if p.Mode != nil {
		m["Mode"] = p.Mode
	}
	if p.NewName != "" {
		m["NewName"] = p.NewName
	}
	if p.OneDriveAdaptiveScopes != nil {
		m["OneDriveAdaptiveScopes"] = p.OneDriveAdaptiveScopes
	}
	if p.OneDriveAdaptiveScopesException != nil {
		m["OneDriveAdaptiveScopesException"] = p.OneDriveAdaptiveScopesException
	}
	if len(p.OneDriveSharedBy) > 0 {
		m["OneDriveSharedBy"] = p.OneDriveSharedBy
	}
	if len(p.OneDriveSharedByMemberOf) > 0 {
		m["OneDriveSharedByMemberOf"] = p.OneDriveSharedByMemberOf
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.PolicyTemplateInfo != nil {
		m["PolicyTemplateInfo"] = p.PolicyTemplateInfo
	}
	if p.Priority != nil {
		m["Priority"] = p.Priority
	}
	if p.RemoveEndpointDlpLocation != nil {
		m["RemoveEndpointDlpLocation"] = p.RemoveEndpointDlpLocation
	}
	if p.RemoveEndpointDlpLocationException != nil {
		m["RemoveEndpointDlpLocationException"] = p.RemoveEndpointDlpLocationException
	}
	if p.RemoveExchangeLocation != nil {
		m["RemoveExchangeLocation"] = p.RemoveExchangeLocation
	}
	if p.RemoveOneDriveLocation != nil {
		m["RemoveOneDriveLocation"] = p.RemoveOneDriveLocation
	}
	if p.RemoveOneDriveLocationException != nil {
		m["RemoveOneDriveLocationException"] = p.RemoveOneDriveLocationException
	}
	if p.RemoveOnPremisesScannerDlpLocation != nil {
		m["RemoveOnPremisesScannerDlpLocation"] = p.RemoveOnPremisesScannerDlpLocation
	}
	if p.RemoveOnPremisesScannerDlpLocationException != nil {
		m["RemoveOnPremisesScannerDlpLocationException"] = p.RemoveOnPremisesScannerDlpLocationException
	}
	if p.RemovePowerBIDlpLocation != nil {
		m["RemovePowerBIDlpLocation"] = p.RemovePowerBIDlpLocation
	}
	if p.RemovePowerBIDlpLocationException != nil {
		m["RemovePowerBIDlpLocationException"] = p.RemovePowerBIDlpLocationException
	}
	if p.RemoveSharePointLocation != nil {
		m["RemoveSharePointLocation"] = p.RemoveSharePointLocation
	}
	if p.RemoveSharePointLocationException != nil {
		m["RemoveSharePointLocationException"] = p.RemoveSharePointLocationException
	}
	if p.RemoveTeamsLocation != nil {
		m["RemoveTeamsLocation"] = p.RemoveTeamsLocation
	}
	if p.RemoveTeamsLocationException != nil {
		m["RemoveTeamsLocationException"] = p.RemoveTeamsLocationException
	}
	if p.RemoveThirdPartyAppDlpLocation != nil {
		m["RemoveThirdPartyAppDlpLocation"] = p.RemoveThirdPartyAppDlpLocation
	}
	if p.RemoveThirdPartyAppDlpLocationException != nil {
		m["RemoveThirdPartyAppDlpLocationException"] = p.RemoveThirdPartyAppDlpLocationException
	}
	if p.RetryDistribution {
		m["RetryDistribution"] = true
	}
	if p.SharePointAdaptiveScopes != nil {
		m["SharePointAdaptiveScopes"] = p.SharePointAdaptiveScopes
	}
	if p.SharePointAdaptiveScopesException != nil {
		m["SharePointAdaptiveScopesException"] = p.SharePointAdaptiveScopesException
	}
	if p.StartSimulation {
		m["StartSimulation"] = true
	}
	if p.TeamsAdaptiveScopes != nil {
		m["TeamsAdaptiveScopes"] = p.TeamsAdaptiveScopes
	}
	if p.TeamsAdaptiveScopesException != nil {
		m["TeamsAdaptiveScopesException"] = p.TeamsAdaptiveScopesException
	}
	return m
}

// SetDlpCompliancePolicy runs the Set-DlpCompliancePolicy cmdlet.
func (s *Service) SetDlpCompliancePolicy(ctx context.Context, p SetDlpCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-DlpCompliancePolicy", p.params())
}

// SetDlpComplianceRuleParams are the parameters of Set-DlpComplianceRule.
// DefaultParameterSetName: Identity
type SetDlpComplianceRuleParams struct {
	AccessedBy                                   []string `ps:"AccessedBy"`
	AccessedByMemberOf                           []string `ps:"AccessedByMemberOf"`
	AccessScope                                  any      `ps:"AccessScope"`
	AccessTimeControl                            any      `ps:"AccessTimeControl"`
	ActivationDate                               any      `ps:"ActivationDate"`
	AddRecipients                                any      `ps:"AddRecipients"`
	AdvancedRule                                 string   `ps:"AdvancedRule"`
	AlertProperties                              any      `ps:"AlertProperties"`
	AnyOfRecipientAddressContainsWords           any      `ps:"AnyOfRecipientAddressContainsWords"`
	AnyOfRecipientAddressMatchesPatterns         any      `ps:"AnyOfRecipientAddressMatchesPatterns"`
	ApplyBrandingTemplate                        string   `ps:"ApplyBrandingTemplate"`
	ApplyHtmlDisclaimer                          any      `ps:"ApplyHtmlDisclaimer"`
	AttachmentCountOver                          any      `ps:"AttachmentCountOver"`
	AttachmentIsNotLabeled                       bool     `ps:"AttachmentIsNotLabeled"`
	BlockAccess                                  bool     `ps:"BlockAccess"`
	BlockAccessScope                             any      `ps:"BlockAccessScope"`
	Comment                                      string   `ps:"Comment"`
	ContentCharacterSetContainsWords             any      `ps:"ContentCharacterSetContainsWords"`
	ContentContainsSensitiveInformation          []string `ps:"ContentContainsSensitiveInformation"`
	ContentExtensionMatchesWords                 any      `ps:"ContentExtensionMatchesWords"`
	ContentFileTypeMatches                       any      `ps:"ContentFileTypeMatches"`
	ContentIsNotLabeled                          bool     `ps:"ContentIsNotLabeled"`
	ContentIsShared                              bool     `ps:"ContentIsShared"`
	ContentMissingSensitivityLabel               any      `ps:"ContentMissingSensitivityLabel"`
	ContentPropertyContainsWords                 any      `ps:"ContentPropertyContainsWords"`
	ContextPropertiesContainWords                any      `ps:"ContextPropertiesContainWords"`
	DeviceManagementType                         any      `ps:"DeviceManagementType"`
	Disabled                                     bool     `ps:"Disabled"`
	DisplayName                                  string   `ps:"DisplayName"`
	DocumentContainsWords                        any      `ps:"DocumentContainsWords"`
	DocumentCreatedBy                            any      `ps:"DocumentCreatedBy"`
	DocumentCreatedByMemberOf                    []string `ps:"DocumentCreatedByMemberOf"`
	DocumentIsPasswordProtected                  bool     `ps:"DocumentIsPasswordProtected"`
	DocumentIsUnsupported                        bool     `ps:"DocumentIsUnsupported"`
	DocumentMatchesPatterns                      any      `ps:"DocumentMatchesPatterns"`
	DocumentNameMatchesPatterns                  any      `ps:"DocumentNameMatchesPatterns"`
	DocumentNameMatchesWords                     any      `ps:"DocumentNameMatchesWords"`
	DocumentSizeOver                             any      `ps:"DocumentSizeOver"`
	EncryptRMSTemplate                           any      `ps:"EncryptRMSTemplate"`
	EndpointDlpBrowserRestrictions               []string `ps:"EndpointDlpBrowserRestrictions"`
	EndpointDlpRestrictions                      []string `ps:"EndpointDlpRestrictions"`
	EnforcePortalAccess                          bool     `ps:"EnforcePortalAccess"`
	EvaluateRulePerComponent                     bool     `ps:"EvaluateRulePerComponent"`
	ExceptIfAccessedBy                           []string `ps:"ExceptIfAccessedBy"`
	ExceptIfAccessedByMemberOf                   []string `ps:"ExceptIfAccessedByMemberOf"`
	ExceptIfAccessScope                          any      `ps:"ExceptIfAccessScope"`
	ExceptIfAnyOfRecipientAddressContainsWords   any      `ps:"ExceptIfAnyOfRecipientAddressContainsWords"`
	ExceptIfAnyOfRecipientAddressMatchesPatterns any      `ps:"ExceptIfAnyOfRecipientAddressMatchesPatterns"`
	ExceptIfContentCharacterSetContainsWords     any      `ps:"ExceptIfContentCharacterSetContainsWords"`
	ExceptIfContentContainsSensitiveInformation  []string `ps:"ExceptIfContentContainsSensitiveInformation"`
	ExceptIfContentExtensionMatchesWords         any      `ps:"ExceptIfContentExtensionMatchesWords"`
	ExceptIfContentFileTypeMatches               any      `ps:"ExceptIfContentFileTypeMatches"`
	ExceptIfContentIsShared                      bool     `ps:"ExceptIfContentIsShared"`
	ExceptIfContentPropertyContainsWords         any      `ps:"ExceptIfContentPropertyContainsWords"`
	ExceptIfContextPropertiesContainWords        any      `ps:"ExceptIfContextPropertiesContainWords"`
	ExceptIfDeviceManagementType                 any      `ps:"ExceptIfDeviceManagementType"` // one of: Managed, Unmanaged
	ExceptIfDocumentContainsWords                any      `ps:"ExceptIfDocumentContainsWords"`
	ExceptIfDocumentCreatedBy                    any      `ps:"ExceptIfDocumentCreatedBy"`
	ExceptIfDocumentCreatedByMemberOf            []string `ps:"ExceptIfDocumentCreatedByMemberOf"`
	ExceptIfDocumentIsPasswordProtected          bool     `ps:"ExceptIfDocumentIsPasswordProtected"`
	ExceptIfDocumentIsUnsupported                bool     `ps:"ExceptIfDocumentIsUnsupported"`
	ExceptIfDocumentMatchesPatterns              any      `ps:"ExceptIfDocumentMatchesPatterns"`
	ExceptIfDocumentNameMatchesPatterns          any      `ps:"ExceptIfDocumentNameMatchesPatterns"`
	ExceptIfDocumentNameMatchesWords             any      `ps:"ExceptIfDocumentNameMatchesWords"`
	ExceptIfDocumentSizeOver                     any      `ps:"ExceptIfDocumentSizeOver"`
	ExceptIfFrom                                 []string `ps:"ExceptIfFrom"`
	ExceptIfFromAddressContainsWords             any      `ps:"ExceptIfFromAddressContainsWords"`
	ExceptIfFromAddressMatchesPatterns           any      `ps:"ExceptIfFromAddressMatchesPatterns"`
	ExceptIfFromMemberOf                         []string `ps:"ExceptIfFromMemberOf"`
	ExceptIfFromScope                            any      `ps:"ExceptIfFromScope"`
	ExceptIfHasSenderOverride                    bool     `ps:"ExceptIfHasSenderOverride"`
	ExceptIfHeaderContainsTokens                 any      `ps:"ExceptIfHeaderContainsTokens"`
	ExceptIfHeaderContainsWords                  any      `ps:"ExceptIfHeaderContainsWords"`
	ExceptIfHeaderMatchesPatterns                any      `ps:"ExceptIfHeaderMatchesPatterns"`
	ExceptIfMessageLabelChangeDetected           any      `ps:"ExceptIfMessageLabelChangeDetected"`
	ExceptIfMessageSizeOver                      any      `ps:"ExceptIfMessageSizeOver"`
	ExceptIfMessageTypeMatches                   any      `ps:"ExceptIfMessageTypeMatches"`
	ExceptIfProcessingLimitExceeded              bool     `ps:"ExceptIfProcessingLimitExceeded"`
	ExceptIfRecipientADAttributeContainsWords    any      `ps:"ExceptIfRecipientADAttributeContainsWords"`
	ExceptIfRecipientADAttributeMatchesPatterns  any      `ps:"ExceptIfRecipientADAttributeMatchesPatterns"`
	ExceptIfRecipientDomainIs                    any      `ps:"ExceptIfRecipientDomainIs"`
	ExceptIfSenderADAttributeContainsWords       any      `ps:"ExceptIfSenderADAttributeContainsWords"`
	ExceptIfSenderADAttributeMatchesPatterns     any      `ps:"ExceptIfSenderADAttributeMatchesPatterns"`
	ExceptIfSenderDomainIs                       any      `ps:"ExceptIfSenderDomainIs"`
	ExceptIfSenderIPRanges                       any      `ps:"ExceptIfSenderIPRanges"`
	ExceptIfSentTo                               any      `ps:"ExceptIfSentTo"`
	ExceptIfSentToMemberOf                       []string `ps:"ExceptIfSentToMemberOf"`
	ExceptIfSharedWithDomain                     any      `ps:"ExceptIfSharedWithDomain"`
	ExceptIfSubjectContainsWords                 any      `ps:"ExceptIfSubjectContainsWords"`
	ExceptIfSubjectMatchesPatterns               any      `ps:"ExceptIfSubjectMatchesPatterns"`
	ExceptIfSubjectOrBodyContainsWords           any      `ps:"ExceptIfSubjectOrBodyContainsWords"`
	ExceptIfSubjectOrBodyMatchesPatterns         any      `ps:"ExceptIfSubjectOrBodyMatchesPatterns"`
	ExceptIfTeamsSharedWithParticipantCategory   any      `ps:"ExceptIfTeamsSharedWithParticipantCategory"`
	ExceptIfUnscannableDocumentExtensionIs       any      `ps:"ExceptIfUnscannableDocumentExtensionIs"`
	ExceptIfWithImportance                       any      `ps:"ExceptIfWithImportance"`
	ExpiryDate                                   any      `ps:"ExpiryDate"`
	From                                         []string `ps:"From"`
	FromAddressContainsWords                     any      `ps:"FromAddressContainsWords"`
	FromAddressMatchesPatterns                   any      `ps:"FromAddressMatchesPatterns"`
	FromMemberOf                                 []string `ps:"FromMemberOf"`
	FromScope                                    any      `ps:"FromScope"`
	GenerateAlert                                any      `ps:"GenerateAlert"`
	GenerateIncidentReport                       any      `ps:"GenerateIncidentReport"`
	HasActivity                                  any      `ps:"HasActivity"`
	HasLabelDowngradedFrom                       any      `ps:"HasLabelDowngradedFrom"`
	HasSenderOverride                            bool     `ps:"HasSenderOverride"`
	HeaderContainsTokens                         any      `ps:"HeaderContainsTokens"`
	HeaderContainsWords                          any      `ps:"HeaderContainsWords"`
	HeaderMatchesPatterns                        any      `ps:"HeaderMatchesPatterns"`
	Identity                                     any      `ps:"Identity"`
	IncidentReportContent                        []string `ps:"IncidentReportContent"`
	MapRecipients                                any      `ps:"MapRecipients"`
	MessageIsNotLabeled                          bool     `ps:"MessageIsNotLabeled"`
	MessageLabelChangeDetected                   any      `ps:"MessageLabelChangeDetected"`
	MessageSizeOver                              any      `ps:"MessageSizeOver"`
	MessageTypeMatches                           any      `ps:"MessageTypeMatches"`
	MipRestrictAccess                            []string `ps:"MipRestrictAccess"`
	Moderate                                     any      `ps:"Moderate"`
	ModifySubject                                any      `ps:"ModifySubject"`
	MoveToQuarantineLocation                     bool     `ps:"MoveToQuarantineLocation"`
	NonBifurcatingAccessScope                    any      `ps:"NonBifurcatingAccessScope"`
	NotifyAllowOverride                          []string `ps:"NotifyAllowOverride"`
	NotifyEmailCustomSenderDisplayName           string   `ps:"NotifyEmailCustomSenderDisplayName"`
	NotifyEmailCustomSubject                     string   `ps:"NotifyEmailCustomSubject"`
	NotifyEmailCustomText                        string   `ps:"NotifyEmailCustomText"`
	NotifyEmailExchangeIncludeAttachment         bool     `ps:"NotifyEmailExchangeIncludeAttachment"`
	NotifyEmailOnedriveRemediationActions        any      `ps:"NotifyEmailOnedriveRemediationActions"`
	NotifyEndpointUser                           any      `ps:"NotifyEndpointUser"`
	NotifyJustificationCustomText                string   `ps:"NotifyJustificationCustomText"`
	NotifyJustificationCustomTextTranslations    any      `ps:"NotifyJustificationCustomTextTranslations"`
	NotifyOverrideRequirements                   any      `ps:"NotifyOverrideRequirements"`
	NotifyPolicyTipCustomDialog                  string   `ps:"NotifyPolicyTipCustomDialog"`
	NotifyPolicyTipCustomText                    string   `ps:"NotifyPolicyTipCustomText"`
	NotifyPolicyTipCustomTextTranslations        any      `ps:"NotifyPolicyTipCustomTextTranslations"`
	NotifyPolicyTipDisplayOption                 any      `ps:"NotifyPolicyTipDisplayOption"`
	NotifyPolicyTipUrl                           string   `ps:"NotifyPolicyTipUrl"`
	NotifyUser                                   any      `ps:"NotifyUser"`
	NotifyUserType                               any      `ps:"NotifyUserType"`
	OnPremisesScannerDlpRestrictions             []string `ps:"OnPremisesScannerDlpRestrictions"`
	PowerBIDlpRestrictions                       []string `ps:"PowerBIDlpRestrictions"`
	PrependSubject                               string   `ps:"PrependSubject"`
	Priority                                     any      `ps:"Priority"`
	ProcessingLimitExceeded                      bool     `ps:"ProcessingLimitExceeded"`
	Quarantine                                   bool     `ps:"Quarantine"`
	RecipientADAttributeContainsWords            any      `ps:"RecipientADAttributeContainsWords"`
	RecipientADAttributeMatchesPatterns          any      `ps:"RecipientADAttributeMatchesPatterns"`
	RecipientDomainIs                            any      `ps:"RecipientDomainIs"`
	RedactSensitiveInfo                          any      `ps:"RedactSensitiveInfo"`
	RedirectMessageTo                            []string `ps:"RedirectMessageTo"`
	RemoveHeader                                 any      `ps:"RemoveHeader"`
	RemoveRMSTemplate                            bool     `ps:"RemoveRMSTemplate"`
	ReportSeverityLevel                          any      `ps:"ReportSeverityLevel"`
	RestrictAccess                               []string `ps:"RestrictAccess"`
	RestrictBrowserAccess                        bool     `ps:"RestrictBrowserAccess"`
	RestrictWebGrounding                         bool     `ps:"RestrictWebGrounding"`
	RuleErrorAction                              any      `ps:"RuleErrorAction"`
	SenderADAttributeContainsWords               any      `ps:"SenderADAttributeContainsWords"`
	SenderADAttributeMatchesPatterns             any      `ps:"SenderADAttributeMatchesPatterns"`
	SenderAddressLocation                        any      `ps:"SenderAddressLocation"`
	SenderDomainIs                               any      `ps:"SenderDomainIs"`
	SenderIPRanges                               any      `ps:"SenderIPRanges"`
	SenderType                                   any      `ps:"SenderType"`
	SentTo                                       any      `ps:"SentTo"`
	SentToMemberOf                               []string `ps:"SentToMemberOf"`
	SetHeader                                    any      `ps:"SetHeader"`
	SharedByIRMAgentRisk                         any      `ps:"SharedByIRMAgentRisk"`
	SharedByIRMUserRisk                          any      `ps:"SharedByIRMUserRisk"`
	SharedWithDomain                             any      `ps:"SharedWithDomain"`
	SharepointBlockDomains                       any      `ps:"SharepointBlockDomains"`
	SharepointBlockDomainsExcept                 any      `ps:"SharepointBlockDomainsExcept"`
	SharepointBlockDomainsOrUsers                bool     `ps:"SharepointBlockDomainsOrUsers"`
	SharepointBlockUsers                         any      `ps:"SharepointBlockUsers"`
	SharepointBlockUsersExcept                   any      `ps:"SharepointBlockUsersExcept"`
	SharepointMoveToQuarantineLocation           bool     `ps:"SharepointMoveToQuarantineLocation"`
	StopPolicyProcessing                         bool     `ps:"StopPolicyProcessing"`
	SubjectContainsWords                         any      `ps:"SubjectContainsWords"`
	SubjectMatchesPatterns                       any      `ps:"SubjectMatchesPatterns"`
	SubjectOrBodyContainsWords                   any      `ps:"SubjectOrBodyContainsWords"`
	SubjectOrBodyMatchesPatterns                 any      `ps:"SubjectOrBodyMatchesPatterns"`
	TeamsSharedWithParticipantCategory           any      `ps:"TeamsSharedWithParticipantCategory"`
	ThirdPartyAppDlpRestrictions                 []string `ps:"ThirdPartyAppDlpRestrictions"`
	TriggerPowerAutomateFlow                     string   `ps:"TriggerPowerAutomateFlow"`
	UnscannableDocumentExtensionIs               any      `ps:"UnscannableDocumentExtensionIs"`
	WithImportance                               any      `ps:"WithImportance"`
}

func (p SetDlpComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if len(p.AccessedBy) > 0 {
		m["AccessedBy"] = p.AccessedBy
	}
	if len(p.AccessedByMemberOf) > 0 {
		m["AccessedByMemberOf"] = p.AccessedByMemberOf
	}
	if p.AccessScope != nil {
		m["AccessScope"] = p.AccessScope
	}
	if p.AccessTimeControl != nil {
		m["AccessTimeControl"] = p.AccessTimeControl
	}
	if p.ActivationDate != nil {
		m["ActivationDate"] = p.ActivationDate
	}
	if p.AddRecipients != nil {
		m["AddRecipients"] = p.AddRecipients
	}
	if p.AdvancedRule != "" {
		m["AdvancedRule"] = p.AdvancedRule
	}
	if p.AlertProperties != nil {
		m["AlertProperties"] = p.AlertProperties
	}
	if p.AnyOfRecipientAddressContainsWords != nil {
		m["AnyOfRecipientAddressContainsWords"] = p.AnyOfRecipientAddressContainsWords
	}
	if p.AnyOfRecipientAddressMatchesPatterns != nil {
		m["AnyOfRecipientAddressMatchesPatterns"] = p.AnyOfRecipientAddressMatchesPatterns
	}
	if p.ApplyBrandingTemplate != "" {
		m["ApplyBrandingTemplate"] = p.ApplyBrandingTemplate
	}
	if p.ApplyHtmlDisclaimer != nil {
		m["ApplyHtmlDisclaimer"] = p.ApplyHtmlDisclaimer
	}
	if p.AttachmentCountOver != nil {
		m["AttachmentCountOver"] = p.AttachmentCountOver
	}
	if p.AttachmentIsNotLabeled {
		m["AttachmentIsNotLabeled"] = true
	}
	if p.BlockAccess {
		m["BlockAccess"] = true
	}
	if p.BlockAccessScope != nil {
		m["BlockAccessScope"] = p.BlockAccessScope
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.ContentCharacterSetContainsWords != nil {
		m["ContentCharacterSetContainsWords"] = p.ContentCharacterSetContainsWords
	}
	if len(p.ContentContainsSensitiveInformation) > 0 {
		m["ContentContainsSensitiveInformation"] = p.ContentContainsSensitiveInformation
	}
	if p.ContentExtensionMatchesWords != nil {
		m["ContentExtensionMatchesWords"] = p.ContentExtensionMatchesWords
	}
	if p.ContentFileTypeMatches != nil {
		m["ContentFileTypeMatches"] = p.ContentFileTypeMatches
	}
	if p.ContentIsNotLabeled {
		m["ContentIsNotLabeled"] = true
	}
	if p.ContentIsShared {
		m["ContentIsShared"] = true
	}
	if p.ContentMissingSensitivityLabel != nil {
		m["ContentMissingSensitivityLabel"] = p.ContentMissingSensitivityLabel
	}
	if p.ContentPropertyContainsWords != nil {
		m["ContentPropertyContainsWords"] = p.ContentPropertyContainsWords
	}
	if p.ContextPropertiesContainWords != nil {
		m["ContextPropertiesContainWords"] = p.ContextPropertiesContainWords
	}
	if p.DeviceManagementType != nil {
		m["DeviceManagementType"] = p.DeviceManagementType
	}
	if p.Disabled {
		m["Disabled"] = true
	}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.DocumentContainsWords != nil {
		m["DocumentContainsWords"] = p.DocumentContainsWords
	}
	if p.DocumentCreatedBy != nil {
		m["DocumentCreatedBy"] = p.DocumentCreatedBy
	}
	if len(p.DocumentCreatedByMemberOf) > 0 {
		m["DocumentCreatedByMemberOf"] = p.DocumentCreatedByMemberOf
	}
	if p.DocumentIsPasswordProtected {
		m["DocumentIsPasswordProtected"] = true
	}
	if p.DocumentIsUnsupported {
		m["DocumentIsUnsupported"] = true
	}
	if p.DocumentMatchesPatterns != nil {
		m["DocumentMatchesPatterns"] = p.DocumentMatchesPatterns
	}
	if p.DocumentNameMatchesPatterns != nil {
		m["DocumentNameMatchesPatterns"] = p.DocumentNameMatchesPatterns
	}
	if p.DocumentNameMatchesWords != nil {
		m["DocumentNameMatchesWords"] = p.DocumentNameMatchesWords
	}
	if p.DocumentSizeOver != nil {
		m["DocumentSizeOver"] = p.DocumentSizeOver
	}
	if p.EncryptRMSTemplate != nil {
		m["EncryptRMSTemplate"] = p.EncryptRMSTemplate
	}
	if len(p.EndpointDlpBrowserRestrictions) > 0 {
		m["EndpointDlpBrowserRestrictions"] = p.EndpointDlpBrowserRestrictions
	}
	if len(p.EndpointDlpRestrictions) > 0 {
		m["EndpointDlpRestrictions"] = p.EndpointDlpRestrictions
	}
	if p.EnforcePortalAccess {
		m["EnforcePortalAccess"] = true
	}
	if p.EvaluateRulePerComponent {
		m["EvaluateRulePerComponent"] = true
	}
	if len(p.ExceptIfAccessedBy) > 0 {
		m["ExceptIfAccessedBy"] = p.ExceptIfAccessedBy
	}
	if len(p.ExceptIfAccessedByMemberOf) > 0 {
		m["ExceptIfAccessedByMemberOf"] = p.ExceptIfAccessedByMemberOf
	}
	if p.ExceptIfAccessScope != nil {
		m["ExceptIfAccessScope"] = p.ExceptIfAccessScope
	}
	if p.ExceptIfAnyOfRecipientAddressContainsWords != nil {
		m["ExceptIfAnyOfRecipientAddressContainsWords"] = p.ExceptIfAnyOfRecipientAddressContainsWords
	}
	if p.ExceptIfAnyOfRecipientAddressMatchesPatterns != nil {
		m["ExceptIfAnyOfRecipientAddressMatchesPatterns"] = p.ExceptIfAnyOfRecipientAddressMatchesPatterns
	}
	if p.ExceptIfContentCharacterSetContainsWords != nil {
		m["ExceptIfContentCharacterSetContainsWords"] = p.ExceptIfContentCharacterSetContainsWords
	}
	if len(p.ExceptIfContentContainsSensitiveInformation) > 0 {
		m["ExceptIfContentContainsSensitiveInformation"] = p.ExceptIfContentContainsSensitiveInformation
	}
	if p.ExceptIfContentExtensionMatchesWords != nil {
		m["ExceptIfContentExtensionMatchesWords"] = p.ExceptIfContentExtensionMatchesWords
	}
	if p.ExceptIfContentFileTypeMatches != nil {
		m["ExceptIfContentFileTypeMatches"] = p.ExceptIfContentFileTypeMatches
	}
	if p.ExceptIfContentIsShared {
		m["ExceptIfContentIsShared"] = true
	}
	if p.ExceptIfContentPropertyContainsWords != nil {
		m["ExceptIfContentPropertyContainsWords"] = p.ExceptIfContentPropertyContainsWords
	}
	if p.ExceptIfContextPropertiesContainWords != nil {
		m["ExceptIfContextPropertiesContainWords"] = p.ExceptIfContextPropertiesContainWords
	}
	if p.ExceptIfDeviceManagementType != nil {
		m["ExceptIfDeviceManagementType"] = p.ExceptIfDeviceManagementType
	}
	if p.ExceptIfDocumentContainsWords != nil {
		m["ExceptIfDocumentContainsWords"] = p.ExceptIfDocumentContainsWords
	}
	if p.ExceptIfDocumentCreatedBy != nil {
		m["ExceptIfDocumentCreatedBy"] = p.ExceptIfDocumentCreatedBy
	}
	if len(p.ExceptIfDocumentCreatedByMemberOf) > 0 {
		m["ExceptIfDocumentCreatedByMemberOf"] = p.ExceptIfDocumentCreatedByMemberOf
	}
	if p.ExceptIfDocumentIsPasswordProtected {
		m["ExceptIfDocumentIsPasswordProtected"] = true
	}
	if p.ExceptIfDocumentIsUnsupported {
		m["ExceptIfDocumentIsUnsupported"] = true
	}
	if p.ExceptIfDocumentMatchesPatterns != nil {
		m["ExceptIfDocumentMatchesPatterns"] = p.ExceptIfDocumentMatchesPatterns
	}
	if p.ExceptIfDocumentNameMatchesPatterns != nil {
		m["ExceptIfDocumentNameMatchesPatterns"] = p.ExceptIfDocumentNameMatchesPatterns
	}
	if p.ExceptIfDocumentNameMatchesWords != nil {
		m["ExceptIfDocumentNameMatchesWords"] = p.ExceptIfDocumentNameMatchesWords
	}
	if p.ExceptIfDocumentSizeOver != nil {
		m["ExceptIfDocumentSizeOver"] = p.ExceptIfDocumentSizeOver
	}
	if len(p.ExceptIfFrom) > 0 {
		m["ExceptIfFrom"] = p.ExceptIfFrom
	}
	if p.ExceptIfFromAddressContainsWords != nil {
		m["ExceptIfFromAddressContainsWords"] = p.ExceptIfFromAddressContainsWords
	}
	if p.ExceptIfFromAddressMatchesPatterns != nil {
		m["ExceptIfFromAddressMatchesPatterns"] = p.ExceptIfFromAddressMatchesPatterns
	}
	if len(p.ExceptIfFromMemberOf) > 0 {
		m["ExceptIfFromMemberOf"] = p.ExceptIfFromMemberOf
	}
	if p.ExceptIfFromScope != nil {
		m["ExceptIfFromScope"] = p.ExceptIfFromScope
	}
	if p.ExceptIfHasSenderOverride {
		m["ExceptIfHasSenderOverride"] = true
	}
	if p.ExceptIfHeaderContainsTokens != nil {
		m["ExceptIfHeaderContainsTokens"] = p.ExceptIfHeaderContainsTokens
	}
	if p.ExceptIfHeaderContainsWords != nil {
		m["ExceptIfHeaderContainsWords"] = p.ExceptIfHeaderContainsWords
	}
	if p.ExceptIfHeaderMatchesPatterns != nil {
		m["ExceptIfHeaderMatchesPatterns"] = p.ExceptIfHeaderMatchesPatterns
	}
	if p.ExceptIfMessageLabelChangeDetected != nil {
		m["ExceptIfMessageLabelChangeDetected"] = p.ExceptIfMessageLabelChangeDetected
	}
	if p.ExceptIfMessageSizeOver != nil {
		m["ExceptIfMessageSizeOver"] = p.ExceptIfMessageSizeOver
	}
	if p.ExceptIfMessageTypeMatches != nil {
		m["ExceptIfMessageTypeMatches"] = p.ExceptIfMessageTypeMatches
	}
	if p.ExceptIfProcessingLimitExceeded {
		m["ExceptIfProcessingLimitExceeded"] = true
	}
	if p.ExceptIfRecipientADAttributeContainsWords != nil {
		m["ExceptIfRecipientADAttributeContainsWords"] = p.ExceptIfRecipientADAttributeContainsWords
	}
	if p.ExceptIfRecipientADAttributeMatchesPatterns != nil {
		m["ExceptIfRecipientADAttributeMatchesPatterns"] = p.ExceptIfRecipientADAttributeMatchesPatterns
	}
	if p.ExceptIfRecipientDomainIs != nil {
		m["ExceptIfRecipientDomainIs"] = p.ExceptIfRecipientDomainIs
	}
	if p.ExceptIfSenderADAttributeContainsWords != nil {
		m["ExceptIfSenderADAttributeContainsWords"] = p.ExceptIfSenderADAttributeContainsWords
	}
	if p.ExceptIfSenderADAttributeMatchesPatterns != nil {
		m["ExceptIfSenderADAttributeMatchesPatterns"] = p.ExceptIfSenderADAttributeMatchesPatterns
	}
	if p.ExceptIfSenderDomainIs != nil {
		m["ExceptIfSenderDomainIs"] = p.ExceptIfSenderDomainIs
	}
	if p.ExceptIfSenderIPRanges != nil {
		m["ExceptIfSenderIPRanges"] = p.ExceptIfSenderIPRanges
	}
	if p.ExceptIfSentTo != nil {
		m["ExceptIfSentTo"] = p.ExceptIfSentTo
	}
	if len(p.ExceptIfSentToMemberOf) > 0 {
		m["ExceptIfSentToMemberOf"] = p.ExceptIfSentToMemberOf
	}
	if p.ExceptIfSharedWithDomain != nil {
		m["ExceptIfSharedWithDomain"] = p.ExceptIfSharedWithDomain
	}
	if p.ExceptIfSubjectContainsWords != nil {
		m["ExceptIfSubjectContainsWords"] = p.ExceptIfSubjectContainsWords
	}
	if p.ExceptIfSubjectMatchesPatterns != nil {
		m["ExceptIfSubjectMatchesPatterns"] = p.ExceptIfSubjectMatchesPatterns
	}
	if p.ExceptIfSubjectOrBodyContainsWords != nil {
		m["ExceptIfSubjectOrBodyContainsWords"] = p.ExceptIfSubjectOrBodyContainsWords
	}
	if p.ExceptIfSubjectOrBodyMatchesPatterns != nil {
		m["ExceptIfSubjectOrBodyMatchesPatterns"] = p.ExceptIfSubjectOrBodyMatchesPatterns
	}
	if p.ExceptIfTeamsSharedWithParticipantCategory != nil {
		m["ExceptIfTeamsSharedWithParticipantCategory"] = p.ExceptIfTeamsSharedWithParticipantCategory
	}
	if p.ExceptIfUnscannableDocumentExtensionIs != nil {
		m["ExceptIfUnscannableDocumentExtensionIs"] = p.ExceptIfUnscannableDocumentExtensionIs
	}
	if p.ExceptIfWithImportance != nil {
		m["ExceptIfWithImportance"] = p.ExceptIfWithImportance
	}
	if p.ExpiryDate != nil {
		m["ExpiryDate"] = p.ExpiryDate
	}
	if len(p.From) > 0 {
		m["From"] = p.From
	}
	if p.FromAddressContainsWords != nil {
		m["FromAddressContainsWords"] = p.FromAddressContainsWords
	}
	if p.FromAddressMatchesPatterns != nil {
		m["FromAddressMatchesPatterns"] = p.FromAddressMatchesPatterns
	}
	if len(p.FromMemberOf) > 0 {
		m["FromMemberOf"] = p.FromMemberOf
	}
	if p.FromScope != nil {
		m["FromScope"] = p.FromScope
	}
	if p.GenerateAlert != nil {
		m["GenerateAlert"] = p.GenerateAlert
	}
	if p.GenerateIncidentReport != nil {
		m["GenerateIncidentReport"] = p.GenerateIncidentReport
	}
	if p.HasActivity != nil {
		m["HasActivity"] = p.HasActivity
	}
	if p.HasLabelDowngradedFrom != nil {
		m["HasLabelDowngradedFrom"] = p.HasLabelDowngradedFrom
	}
	if p.HasSenderOverride {
		m["HasSenderOverride"] = true
	}
	if p.HeaderContainsTokens != nil {
		m["HeaderContainsTokens"] = p.HeaderContainsTokens
	}
	if p.HeaderContainsWords != nil {
		m["HeaderContainsWords"] = p.HeaderContainsWords
	}
	if p.HeaderMatchesPatterns != nil {
		m["HeaderMatchesPatterns"] = p.HeaderMatchesPatterns
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if len(p.IncidentReportContent) > 0 {
		m["IncidentReportContent"] = p.IncidentReportContent
	}
	if p.MapRecipients != nil {
		m["MapRecipients"] = p.MapRecipients
	}
	if p.MessageIsNotLabeled {
		m["MessageIsNotLabeled"] = true
	}
	if p.MessageLabelChangeDetected != nil {
		m["MessageLabelChangeDetected"] = p.MessageLabelChangeDetected
	}
	if p.MessageSizeOver != nil {
		m["MessageSizeOver"] = p.MessageSizeOver
	}
	if p.MessageTypeMatches != nil {
		m["MessageTypeMatches"] = p.MessageTypeMatches
	}
	if len(p.MipRestrictAccess) > 0 {
		m["MipRestrictAccess"] = p.MipRestrictAccess
	}
	if p.Moderate != nil {
		m["Moderate"] = p.Moderate
	}
	if p.ModifySubject != nil {
		m["ModifySubject"] = p.ModifySubject
	}
	if p.MoveToQuarantineLocation {
		m["MoveToQuarantineLocation"] = true
	}
	if p.NonBifurcatingAccessScope != nil {
		m["NonBifurcatingAccessScope"] = p.NonBifurcatingAccessScope
	}
	if len(p.NotifyAllowOverride) > 0 {
		m["NotifyAllowOverride"] = p.NotifyAllowOverride
	}
	if p.NotifyEmailCustomSenderDisplayName != "" {
		m["NotifyEmailCustomSenderDisplayName"] = p.NotifyEmailCustomSenderDisplayName
	}
	if p.NotifyEmailCustomSubject != "" {
		m["NotifyEmailCustomSubject"] = p.NotifyEmailCustomSubject
	}
	if p.NotifyEmailCustomText != "" {
		m["NotifyEmailCustomText"] = p.NotifyEmailCustomText
	}
	if p.NotifyEmailExchangeIncludeAttachment {
		m["NotifyEmailExchangeIncludeAttachment"] = true
	}
	if p.NotifyEmailOnedriveRemediationActions != nil {
		m["NotifyEmailOnedriveRemediationActions"] = p.NotifyEmailOnedriveRemediationActions
	}
	if p.NotifyEndpointUser != nil {
		m["NotifyEndpointUser"] = p.NotifyEndpointUser
	}
	if p.NotifyJustificationCustomText != "" {
		m["NotifyJustificationCustomText"] = p.NotifyJustificationCustomText
	}
	if p.NotifyJustificationCustomTextTranslations != nil {
		m["NotifyJustificationCustomTextTranslations"] = p.NotifyJustificationCustomTextTranslations
	}
	if p.NotifyOverrideRequirements != nil {
		m["NotifyOverrideRequirements"] = p.NotifyOverrideRequirements
	}
	if p.NotifyPolicyTipCustomDialog != "" {
		m["NotifyPolicyTipCustomDialog"] = p.NotifyPolicyTipCustomDialog
	}
	if p.NotifyPolicyTipCustomText != "" {
		m["NotifyPolicyTipCustomText"] = p.NotifyPolicyTipCustomText
	}
	if p.NotifyPolicyTipCustomTextTranslations != nil {
		m["NotifyPolicyTipCustomTextTranslations"] = p.NotifyPolicyTipCustomTextTranslations
	}
	if p.NotifyPolicyTipDisplayOption != nil {
		m["NotifyPolicyTipDisplayOption"] = p.NotifyPolicyTipDisplayOption
	}
	if p.NotifyPolicyTipUrl != "" {
		m["NotifyPolicyTipUrl"] = p.NotifyPolicyTipUrl
	}
	if p.NotifyUser != nil {
		m["NotifyUser"] = p.NotifyUser
	}
	if p.NotifyUserType != nil {
		m["NotifyUserType"] = p.NotifyUserType
	}
	if len(p.OnPremisesScannerDlpRestrictions) > 0 {
		m["OnPremisesScannerDlpRestrictions"] = p.OnPremisesScannerDlpRestrictions
	}
	if len(p.PowerBIDlpRestrictions) > 0 {
		m["PowerBIDlpRestrictions"] = p.PowerBIDlpRestrictions
	}
	if p.PrependSubject != "" {
		m["PrependSubject"] = p.PrependSubject
	}
	if p.Priority != nil {
		m["Priority"] = p.Priority
	}
	if p.ProcessingLimitExceeded {
		m["ProcessingLimitExceeded"] = true
	}
	if p.Quarantine {
		m["Quarantine"] = true
	}
	if p.RecipientADAttributeContainsWords != nil {
		m["RecipientADAttributeContainsWords"] = p.RecipientADAttributeContainsWords
	}
	if p.RecipientADAttributeMatchesPatterns != nil {
		m["RecipientADAttributeMatchesPatterns"] = p.RecipientADAttributeMatchesPatterns
	}
	if p.RecipientDomainIs != nil {
		m["RecipientDomainIs"] = p.RecipientDomainIs
	}
	if p.RedactSensitiveInfo != nil {
		m["RedactSensitiveInfo"] = p.RedactSensitiveInfo
	}
	if len(p.RedirectMessageTo) > 0 {
		m["RedirectMessageTo"] = p.RedirectMessageTo
	}
	if p.RemoveHeader != nil {
		m["RemoveHeader"] = p.RemoveHeader
	}
	if p.RemoveRMSTemplate {
		m["RemoveRMSTemplate"] = true
	}
	if p.ReportSeverityLevel != nil {
		m["ReportSeverityLevel"] = p.ReportSeverityLevel
	}
	if len(p.RestrictAccess) > 0 {
		m["RestrictAccess"] = p.RestrictAccess
	}
	if p.RestrictBrowserAccess {
		m["RestrictBrowserAccess"] = true
	}
	if p.RestrictWebGrounding {
		m["RestrictWebGrounding"] = true
	}
	if p.RuleErrorAction != nil {
		m["RuleErrorAction"] = p.RuleErrorAction
	}
	if p.SenderADAttributeContainsWords != nil {
		m["SenderADAttributeContainsWords"] = p.SenderADAttributeContainsWords
	}
	if p.SenderADAttributeMatchesPatterns != nil {
		m["SenderADAttributeMatchesPatterns"] = p.SenderADAttributeMatchesPatterns
	}
	if p.SenderAddressLocation != nil {
		m["SenderAddressLocation"] = p.SenderAddressLocation
	}
	if p.SenderDomainIs != nil {
		m["SenderDomainIs"] = p.SenderDomainIs
	}
	if p.SenderIPRanges != nil {
		m["SenderIPRanges"] = p.SenderIPRanges
	}
	if p.SenderType != nil {
		m["SenderType"] = p.SenderType
	}
	if p.SentTo != nil {
		m["SentTo"] = p.SentTo
	}
	if len(p.SentToMemberOf) > 0 {
		m["SentToMemberOf"] = p.SentToMemberOf
	}
	if p.SetHeader != nil {
		m["SetHeader"] = p.SetHeader
	}
	if p.SharedByIRMAgentRisk != nil {
		m["SharedByIRMAgentRisk"] = p.SharedByIRMAgentRisk
	}
	if p.SharedByIRMUserRisk != nil {
		m["SharedByIRMUserRisk"] = p.SharedByIRMUserRisk
	}
	if p.SharedWithDomain != nil {
		m["SharedWithDomain"] = p.SharedWithDomain
	}
	if p.SharepointBlockDomains != nil {
		m["SharepointBlockDomains"] = p.SharepointBlockDomains
	}
	if p.SharepointBlockDomainsExcept != nil {
		m["SharepointBlockDomainsExcept"] = p.SharepointBlockDomainsExcept
	}
	if p.SharepointBlockDomainsOrUsers {
		m["SharepointBlockDomainsOrUsers"] = true
	}
	if p.SharepointBlockUsers != nil {
		m["SharepointBlockUsers"] = p.SharepointBlockUsers
	}
	if p.SharepointBlockUsersExcept != nil {
		m["SharepointBlockUsersExcept"] = p.SharepointBlockUsersExcept
	}
	if p.SharepointMoveToQuarantineLocation {
		m["SharepointMoveToQuarantineLocation"] = true
	}
	if p.StopPolicyProcessing {
		m["StopPolicyProcessing"] = true
	}
	if p.SubjectContainsWords != nil {
		m["SubjectContainsWords"] = p.SubjectContainsWords
	}
	if p.SubjectMatchesPatterns != nil {
		m["SubjectMatchesPatterns"] = p.SubjectMatchesPatterns
	}
	if p.SubjectOrBodyContainsWords != nil {
		m["SubjectOrBodyContainsWords"] = p.SubjectOrBodyContainsWords
	}
	if p.SubjectOrBodyMatchesPatterns != nil {
		m["SubjectOrBodyMatchesPatterns"] = p.SubjectOrBodyMatchesPatterns
	}
	if p.TeamsSharedWithParticipantCategory != nil {
		m["TeamsSharedWithParticipantCategory"] = p.TeamsSharedWithParticipantCategory
	}
	if len(p.ThirdPartyAppDlpRestrictions) > 0 {
		m["ThirdPartyAppDlpRestrictions"] = p.ThirdPartyAppDlpRestrictions
	}
	if p.TriggerPowerAutomateFlow != "" {
		m["TriggerPowerAutomateFlow"] = p.TriggerPowerAutomateFlow
	}
	if p.UnscannableDocumentExtensionIs != nil {
		m["UnscannableDocumentExtensionIs"] = p.UnscannableDocumentExtensionIs
	}
	if p.WithImportance != nil {
		m["WithImportance"] = p.WithImportance
	}
	return m
}

// SetDlpComplianceRule runs the Set-DlpComplianceRule cmdlet.
func (s *Service) SetDlpComplianceRule(ctx context.Context, p SetDlpComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-DlpComplianceRule", p.params())
}

// SetDlpEdmSchemaParams are the parameters of Set-DlpEdmSchema.
type SetDlpEdmSchemaParams struct {
	FileData []string `ps:"FileData"`
}

func (p SetDlpEdmSchemaParams) params() map[string]any {
	m := map[string]any{}
	if len(p.FileData) > 0 {
		m["FileData"] = p.FileData
	}
	return m
}

// SetDlpEdmSchema runs the Set-DlpEdmSchema cmdlet.
func (s *Service) SetDlpEdmSchema(ctx context.Context, p SetDlpEdmSchemaParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-DlpEdmSchema", p.params())
}

// SetDlpKeywordDictionaryParams are the parameters of Set-DlpKeywordDictionary.
type SetDlpKeywordDictionaryParams struct {
	Description          string   `ps:"Description"`
	DoNotPersistKeywords bool     `ps:"DoNotPersistKeywords"`
	FileData             []string `ps:"FileData"`
	Identity             any      `ps:"Identity"`
	MatchStyle           string   `ps:"MatchStyle"` // one of: word, string
	Name                 string   `ps:"Name"`
}

func (p SetDlpKeywordDictionaryParams) params() map[string]any {
	m := map[string]any{}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if p.DoNotPersistKeywords {
		m["DoNotPersistKeywords"] = true
	}
	if len(p.FileData) > 0 {
		m["FileData"] = p.FileData
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.MatchStyle != "" {
		m["MatchStyle"] = p.MatchStyle
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	return m
}

// SetDlpKeywordDictionary runs the Set-DlpKeywordDictionary cmdlet.
func (s *Service) SetDlpKeywordDictionary(ctx context.Context, p SetDlpKeywordDictionaryParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-DlpKeywordDictionary", p.params())
}

// SetDlpSensitiveInformationTypeParams are the parameters of Set-DlpSensitiveInformationType.
type SetDlpSensitiveInformationTypeParams struct {
	Description     string   `ps:"Description"`
	FileData        []string `ps:"FileData"`
	Fingerprints    any      `ps:"Fingerprints"`
	Identity        any      `ps:"Identity"`
	IsExact         bool     `ps:"IsExact"`
	Locale          any      `ps:"Locale"`
	Name            string   `ps:"Name"`
	Threshold       any      `ps:"Threshold"`
	ThresholdConfig any      `ps:"ThresholdConfig"`
}

func (p SetDlpSensitiveInformationTypeParams) params() map[string]any {
	m := map[string]any{}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if len(p.FileData) > 0 {
		m["FileData"] = p.FileData
	}
	if p.Fingerprints != nil {
		m["Fingerprints"] = p.Fingerprints
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IsExact {
		m["IsExact"] = true
	}
	if p.Locale != nil {
		m["Locale"] = p.Locale
	}
	if p.Name != "" {
		m["Name"] = p.Name
	}
	if p.Threshold != nil {
		m["Threshold"] = p.Threshold
	}
	if p.ThresholdConfig != nil {
		m["ThresholdConfig"] = p.ThresholdConfig
	}
	return m
}

// SetDlpSensitiveInformationType runs the Set-DlpSensitiveInformationType cmdlet.
func (s *Service) SetDlpSensitiveInformationType(ctx context.Context, p SetDlpSensitiveInformationTypeParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-DlpSensitiveInformationType", p.params())
}

// SetDlpSensitiveInformationTypeConfigParams are the parameters of Set-DlpSensitiveInformationTypeConfig.
// DefaultParameterSetName: Identity
type SetDlpSensitiveInformationTypeConfigParams struct {
	FingerprintThreshold int  `ps:"FingerprintThreshold"`
	RegExGrammarLimit    bool `ps:"RegExGrammarLimit"`
}

func (p SetDlpSensitiveInformationTypeConfigParams) params() map[string]any {
	m := map[string]any{}
	if p.FingerprintThreshold != 0 {
		m["FingerprintThreshold"] = p.FingerprintThreshold
	}
	if p.RegExGrammarLimit {
		m["RegExGrammarLimit"] = true
	}
	return m
}

// SetDlpSensitiveInformationTypeConfig runs the Set-DlpSensitiveInformationTypeConfig cmdlet.
func (s *Service) SetDlpSensitiveInformationTypeConfig(ctx context.Context, p SetDlpSensitiveInformationTypeConfigParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-DlpSensitiveInformationTypeConfig", p.params())
}

// SetDlpSensitiveInformationTypeRulePackageParams are the parameters of Set-DlpSensitiveInformationTypeRulePackage.
type SetDlpSensitiveInformationTypeRulePackageParams struct {
	FileData []string `ps:"FileData"`
}

func (p SetDlpSensitiveInformationTypeRulePackageParams) params() map[string]any {
	m := map[string]any{}
	if len(p.FileData) > 0 {
		m["FileData"] = p.FileData
	}
	return m
}

// SetDlpSensitiveInformationTypeRulePackage runs the Set-DlpSensitiveInformationTypeRulePackage cmdlet.
func (s *Service) SetDlpSensitiveInformationTypeRulePackage(ctx context.Context, p SetDlpSensitiveInformationTypeRulePackageParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-DlpSensitiveInformationTypeRulePackage", p.params())
}

// SetDspmPolicyParams are the parameters of Set-DspmPolicy.
// DefaultParameterSetName: Identity
type SetDspmPolicyParams struct {
	Enabled       bool   `ps:"Enabled"`
	Identity      any    `ps:"Identity"`
	TenantSetting string `ps:"TenantSetting"`
}

func (p SetDspmPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.TenantSetting != "" {
		m["TenantSetting"] = p.TenantSetting
	}
	return m
}

// SetDspmPolicy runs the Set-DspmPolicy cmdlet.
func (s *Service) SetDspmPolicy(ctx context.Context, p SetDspmPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-DspmPolicy", p.params())
}

// SetFeatureConfigurationParams are the parameters of Set-FeatureConfiguration.
// DefaultParameterSetName: Identity
type SetFeatureConfigurationParams struct {
	Comment          string `ps:"Comment"`
	Identity         any    `ps:"Identity"`
	Locations        string `ps:"Locations"`
	Mode             any    `ps:"Mode"`
	PolicyRBACScopes any    `ps:"PolicyRBACScopes"`
	ScenarioConfig   string `ps:"ScenarioConfig"`
}

func (p SetFeatureConfigurationParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Locations != "" {
		m["Locations"] = p.Locations
	}
	if p.Mode != nil {
		m["Mode"] = p.Mode
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.ScenarioConfig != "" {
		m["ScenarioConfig"] = p.ScenarioConfig
	}
	return m
}

// SetFeatureConfiguration runs the Set-FeatureConfiguration cmdlet.
func (s *Service) SetFeatureConfiguration(ctx context.Context, p SetFeatureConfigurationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-FeatureConfiguration", p.params())
}

// SetFilePlanPropertyAuthorityParams are the parameters of Set-FilePlanPropertyAuthority.
// DefaultParameterSetName: Identity
type SetFilePlanPropertyAuthorityParams struct {
	DisplayName string `ps:"DisplayName"`
	Identity    any    `ps:"Identity"`
}

func (p SetFilePlanPropertyAuthorityParams) params() map[string]any {
	m := map[string]any{}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// SetFilePlanPropertyAuthority runs the Set-FilePlanPropertyAuthority cmdlet.
func (s *Service) SetFilePlanPropertyAuthority(ctx context.Context, p SetFilePlanPropertyAuthorityParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-FilePlanPropertyAuthority", p.params())
}

// SetFilePlanPropertyCategoryParams are the parameters of Set-FilePlanPropertyCategory.
// DefaultParameterSetName: Identity
type SetFilePlanPropertyCategoryParams struct {
	DisplayName string `ps:"DisplayName"`
	Identity    any    `ps:"Identity"`
}

func (p SetFilePlanPropertyCategoryParams) params() map[string]any {
	m := map[string]any{}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// SetFilePlanPropertyCategory runs the Set-FilePlanPropertyCategory cmdlet.
func (s *Service) SetFilePlanPropertyCategory(ctx context.Context, p SetFilePlanPropertyCategoryParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-FilePlanPropertyCategory", p.params())
}

// SetFilePlanPropertyCitationParams are the parameters of Set-FilePlanPropertyCitation.
// DefaultParameterSetName: Identity
type SetFilePlanPropertyCitationParams struct {
	CitationJurisdiction string `ps:"CitationJurisdiction"`
	CitationUrl          string `ps:"CitationUrl"`
	DisplayName          string `ps:"DisplayName"`
	Identity             any    `ps:"Identity"`
}

func (p SetFilePlanPropertyCitationParams) params() map[string]any {
	m := map[string]any{}
	if p.CitationJurisdiction != "" {
		m["CitationJurisdiction"] = p.CitationJurisdiction
	}
	if p.CitationUrl != "" {
		m["CitationUrl"] = p.CitationUrl
	}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// SetFilePlanPropertyCitation runs the Set-FilePlanPropertyCitation cmdlet.
func (s *Service) SetFilePlanPropertyCitation(ctx context.Context, p SetFilePlanPropertyCitationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-FilePlanPropertyCitation", p.params())
}

// SetFilePlanPropertyDepartmentParams are the parameters of Set-FilePlanPropertyDepartment.
// DefaultParameterSetName: Identity
type SetFilePlanPropertyDepartmentParams struct {
	DisplayName string `ps:"DisplayName"`
	Identity    any    `ps:"Identity"`
}

func (p SetFilePlanPropertyDepartmentParams) params() map[string]any {
	m := map[string]any{}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// SetFilePlanPropertyDepartment runs the Set-FilePlanPropertyDepartment cmdlet.
func (s *Service) SetFilePlanPropertyDepartment(ctx context.Context, p SetFilePlanPropertyDepartmentParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-FilePlanPropertyDepartment", p.params())
}

// SetFilePlanPropertyReferenceIdParams are the parameters of Set-FilePlanPropertyReferenceId.
// DefaultParameterSetName: Identity
type SetFilePlanPropertyReferenceIdParams struct {
	DisplayName string `ps:"DisplayName"`
	Identity    any    `ps:"Identity"`
}

func (p SetFilePlanPropertyReferenceIdParams) params() map[string]any {
	m := map[string]any{}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// SetFilePlanPropertyReferenceId runs the Set-FilePlanPropertyReferenceId cmdlet.
func (s *Service) SetFilePlanPropertyReferenceId(ctx context.Context, p SetFilePlanPropertyReferenceIdParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-FilePlanPropertyReferenceId", p.params())
}

// SetFilePlanPropertySubCategoryParams are the parameters of Set-FilePlanPropertySubCategory.
// DefaultParameterSetName: Identity
type SetFilePlanPropertySubCategoryParams struct {
	DisplayName string `ps:"DisplayName"`
	Identity    any    `ps:"Identity"`
}

func (p SetFilePlanPropertySubCategoryParams) params() map[string]any {
	m := map[string]any{}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// SetFilePlanPropertySubCategory runs the Set-FilePlanPropertySubCategory cmdlet.
func (s *Service) SetFilePlanPropertySubCategory(ctx context.Context, p SetFilePlanPropertySubCategoryParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-FilePlanPropertySubCategory", p.params())
}

// SetGlobalListParams are the parameters of Set-GlobalList.
type SetGlobalListParams struct {
	AddItems        any    `ps:"AddItems"`
	Comment         string `ps:"Comment"`
	GlobalListValue any    `ps:"GlobalListValue"`
	Identity        any    `ps:"Identity"`
	RemoveItems     any    `ps:"RemoveItems"`
}

func (p SetGlobalListParams) params() map[string]any {
	m := map[string]any{}
	if p.AddItems != nil {
		m["AddItems"] = p.AddItems
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.GlobalListValue != nil {
		m["GlobalListValue"] = p.GlobalListValue
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.RemoveItems != nil {
		m["RemoveItems"] = p.RemoveItems
	}
	return m
}

// SetGlobalList runs the Set-GlobalList cmdlet.
func (s *Service) SetGlobalList(ctx context.Context, p SetGlobalListParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-GlobalList", p.params())
}

// SetHoldCompliancePolicyParams are the parameters of Set-HoldCompliancePolicy.
// DefaultParameterSetName: Identity
type SetHoldCompliancePolicyParams struct {
	AddExchangeLocation        any    `ps:"AddExchangeLocation"`
	AddPublicFolderLocation    any    `ps:"AddPublicFolderLocation"`
	AddSharePointLocation      any    `ps:"AddSharePointLocation"`
	Comment                    string `ps:"Comment"`
	Enabled                    bool   `ps:"Enabled"`
	Force                      bool   `ps:"Force"`
	Identity                   any    `ps:"Identity"`
	RemoveExchangeLocation     any    `ps:"RemoveExchangeLocation"`
	RemovePublicFolderLocation any    `ps:"RemovePublicFolderLocation"`
	RemoveSharePointLocation   any    `ps:"RemoveSharePointLocation"`
	RetryDistribution          bool   `ps:"RetryDistribution"`
}

func (p SetHoldCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.AddExchangeLocation != nil {
		m["AddExchangeLocation"] = p.AddExchangeLocation
	}
	if p.AddPublicFolderLocation != nil {
		m["AddPublicFolderLocation"] = p.AddPublicFolderLocation
	}
	if p.AddSharePointLocation != nil {
		m["AddSharePointLocation"] = p.AddSharePointLocation
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.RemoveExchangeLocation != nil {
		m["RemoveExchangeLocation"] = p.RemoveExchangeLocation
	}
	if p.RemovePublicFolderLocation != nil {
		m["RemovePublicFolderLocation"] = p.RemovePublicFolderLocation
	}
	if p.RemoveSharePointLocation != nil {
		m["RemoveSharePointLocation"] = p.RemoveSharePointLocation
	}
	if p.RetryDistribution {
		m["RetryDistribution"] = true
	}
	return m
}

// SetHoldCompliancePolicy runs the Set-HoldCompliancePolicy cmdlet.
func (s *Service) SetHoldCompliancePolicy(ctx context.Context, p SetHoldCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-HoldCompliancePolicy", p.params())
}

// SetHoldComplianceRuleParams are the parameters of Set-HoldComplianceRule.
// DefaultParameterSetName: Identity
type SetHoldComplianceRuleParams struct {
	Comment                 string `ps:"Comment"`
	ContentDateFrom         any    `ps:"ContentDateFrom"`
	ContentDateTo           any    `ps:"ContentDateTo"`
	ContentMatchQuery       string `ps:"ContentMatchQuery"`
	Disabled                bool   `ps:"Disabled"`
	HoldContent             any    `ps:"HoldContent"`
	HoldDurationDisplayHint any    `ps:"HoldDurationDisplayHint"`
	Identity                any    `ps:"Identity"`
}

func (p SetHoldComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.ContentDateFrom != nil {
		m["ContentDateFrom"] = p.ContentDateFrom
	}
	if p.ContentDateTo != nil {
		m["ContentDateTo"] = p.ContentDateTo
	}
	if p.ContentMatchQuery != "" {
		m["ContentMatchQuery"] = p.ContentMatchQuery
	}
	if p.Disabled {
		m["Disabled"] = true
	}
	if p.HoldContent != nil {
		m["HoldContent"] = p.HoldContent
	}
	if p.HoldDurationDisplayHint != nil {
		m["HoldDurationDisplayHint"] = p.HoldDurationDisplayHint
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// SetHoldComplianceRule runs the Set-HoldComplianceRule cmdlet.
func (s *Service) SetHoldComplianceRule(ctx context.Context, p SetHoldComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-HoldComplianceRule", p.params())
}

// SetInformationBarrierPolicyParams are the parameters of Set-InformationBarrierPolicy.
// DefaultParameterSetName: InformationBarrierDefault
type SetInformationBarrierPolicyParams struct {
	Comment              string `ps:"Comment"`
	Force                bool   `ps:"Force"`
	Identity             any    `ps:"Identity"`
	ModerationAllowed    bool   `ps:"ModerationAllowed"`
	SegmentAllowedFilter string `ps:"SegmentAllowedFilter"`
	SegmentsAllowed      any    `ps:"SegmentsAllowed"`
	SegmentsBlocked      any    `ps:"SegmentsBlocked"`
	State                any    `ps:"State"`
}

func (p SetInformationBarrierPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.ModerationAllowed {
		m["ModerationAllowed"] = true
	}
	if p.SegmentAllowedFilter != "" {
		m["SegmentAllowedFilter"] = p.SegmentAllowedFilter
	}
	if p.SegmentsAllowed != nil {
		m["SegmentsAllowed"] = p.SegmentsAllowed
	}
	if p.SegmentsBlocked != nil {
		m["SegmentsBlocked"] = p.SegmentsBlocked
	}
	if p.State != nil {
		m["State"] = p.State
	}
	return m
}

// SetInformationBarrierPolicy runs the Set-InformationBarrierPolicy cmdlet.
func (s *Service) SetInformationBarrierPolicy(ctx context.Context, p SetInformationBarrierPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-InformationBarrierPolicy", p.params())
}

// SetInsiderRiskEntityListParams are the parameters of Set-InsiderRiskEntityList.
// DefaultParameterSetName: Identity
type SetInsiderRiskEntityListParams struct {
	AddEntities    any    `ps:"AddEntities"`
	Description    string `ps:"Description"`
	DisplayName    string `ps:"DisplayName"`
	Identity       any    `ps:"Identity"`
	RemoveEntities any    `ps:"RemoveEntities"`
}

func (p SetInsiderRiskEntityListParams) params() map[string]any {
	m := map[string]any{}
	if p.AddEntities != nil {
		m["AddEntities"] = p.AddEntities
	}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.RemoveEntities != nil {
		m["RemoveEntities"] = p.RemoveEntities
	}
	return m
}

// SetInsiderRiskEntityList runs the Set-InsiderRiskEntityList cmdlet.
func (s *Service) SetInsiderRiskEntityList(ctx context.Context, p SetInsiderRiskEntityListParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-InsiderRiskEntityList", p.params())
}

// SetInsiderRiskPolicyParams are the parameters of Set-InsiderRiskPolicy.
// DefaultParameterSetName: Identity
type SetInsiderRiskPolicyParams struct {
	AddExchangeLocation                any    `ps:"AddExchangeLocation"`
	AddExchangeLocationException       any    `ps:"AddExchangeLocationException"`
	AddIrmAdaptiveScopeLocation        any    `ps:"AddIrmAdaptiveScopeLocation"`
	AddModernGroupLocation             any    `ps:"AddModernGroupLocation"`
	AddModernGroupLocationException    any    `ps:"AddModernGroupLocationException"`
	AddUserMailList                    any    `ps:"AddUserMailList"`
	CCPolicyName                       string `ps:"CCPolicyName"`
	CCPolicySdsId                      string `ps:"CCPolicySdsId"`
	Comment                            string `ps:"Comment"`
	CustomTags                         any    `ps:"CustomTags"`
	DlpPoliciesAsTrigger               any    `ps:"DlpPoliciesAsTrigger"`
	DlpPolicy                          any    `ps:"DlpPolicy"`
	DlpSensitiveTypes                  any    `ps:"DlpSensitiveTypes"`
	Enabled                            bool   `ps:"Enabled"`
	ExtensibleIndicators               any    `ps:"ExtensibleIndicators"`
	ExtensibleTriggerInsightGroups     any    `ps:"ExtensibleTriggerInsightGroups"`
	FileExtensions                     any    `ps:"FileExtensions"`
	FutureTerminationWindow            int    `ps:"FutureTerminationWindow"`
	HistoricTimeSpan                   int    `ps:"HistoricTimeSpan"`
	Identity                           any    `ps:"Identity"`
	Indicators                         any    `ps:"Indicators"`
	InScopeTimeSpan                    int    `ps:"InScopeTimeSpan"`
	IsCustom                           bool   `ps:"IsCustom"`
	IsPriorityContentOnlyScoring       bool   `ps:"IsPriorityContentOnlyScoring"`
	MLClassifierTypes                  any    `ps:"MLClassifierTypes"`
	OptInDrpForDlp                     bool   `ps:"OptInDrpForDlp"`
	PastTerminationWindow              int    `ps:"PastTerminationWindow"`
	PolicyRBACScopes                   any    `ps:"PolicyRBACScopes"`
	PostTerminationActivity            bool   `ps:"PostTerminationActivity"`
	RemoveExchangeLocation             any    `ps:"RemoveExchangeLocation"`
	RemoveExchangeLocationException    any    `ps:"RemoveExchangeLocationException"`
	RemoveIrmAdaptiveScopeLocation     any    `ps:"RemoveIrmAdaptiveScopeLocation"`
	RemoveModernGroupLocation          any    `ps:"RemoveModernGroupLocation"`
	RemoveModernGroupLocationException any    `ps:"RemoveModernGroupLocationException"`
	RetryDistribution                  bool   `ps:"RetryDistribution"`
	SchemaVersion                      int    `ps:"SchemaVersion"`
	SensitivityLabels                  any    `ps:"SensitivityLabels"`
	SessionRecordingSettings           string `ps:"SessionRecordingSettings"`
	SharepointSites                    any    `ps:"SharepointSites"`
	TeamsSites                         any    `ps:"TeamsSites"`
	TenantSetting                      string `ps:"TenantSetting"`
	TriggerInsightGroups               any    `ps:"TriggerInsightGroups"`
	Triggers                           any    `ps:"Triggers"`
	TurnOnAnalytics                    string `ps:"TurnOnAnalytics"`
	TurnOnDLPUserRiskSync              string `ps:"TurnOnDLPUserRiskSync"`
}

func (p SetInsiderRiskPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.AddExchangeLocation != nil {
		m["AddExchangeLocation"] = p.AddExchangeLocation
	}
	if p.AddExchangeLocationException != nil {
		m["AddExchangeLocationException"] = p.AddExchangeLocationException
	}
	if p.AddIrmAdaptiveScopeLocation != nil {
		m["AddIrmAdaptiveScopeLocation"] = p.AddIrmAdaptiveScopeLocation
	}
	if p.AddModernGroupLocation != nil {
		m["AddModernGroupLocation"] = p.AddModernGroupLocation
	}
	if p.AddModernGroupLocationException != nil {
		m["AddModernGroupLocationException"] = p.AddModernGroupLocationException
	}
	if p.AddUserMailList != nil {
		m["AddUserMailList"] = p.AddUserMailList
	}
	if p.CCPolicyName != "" {
		m["CCPolicyName"] = p.CCPolicyName
	}
	if p.CCPolicySdsId != "" {
		m["CCPolicySdsId"] = p.CCPolicySdsId
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.CustomTags != nil {
		m["CustomTags"] = p.CustomTags
	}
	if p.DlpPoliciesAsTrigger != nil {
		m["DlpPoliciesAsTrigger"] = p.DlpPoliciesAsTrigger
	}
	if p.DlpPolicy != nil {
		m["DlpPolicy"] = p.DlpPolicy
	}
	if p.DlpSensitiveTypes != nil {
		m["DlpSensitiveTypes"] = p.DlpSensitiveTypes
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.ExtensibleIndicators != nil {
		m["ExtensibleIndicators"] = p.ExtensibleIndicators
	}
	if p.ExtensibleTriggerInsightGroups != nil {
		m["ExtensibleTriggerInsightGroups"] = p.ExtensibleTriggerInsightGroups
	}
	if p.FileExtensions != nil {
		m["FileExtensions"] = p.FileExtensions
	}
	if p.FutureTerminationWindow != 0 {
		m["FutureTerminationWindow"] = p.FutureTerminationWindow
	}
	if p.HistoricTimeSpan != 0 {
		m["HistoricTimeSpan"] = p.HistoricTimeSpan
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Indicators != nil {
		m["Indicators"] = p.Indicators
	}
	if p.InScopeTimeSpan != 0 {
		m["InScopeTimeSpan"] = p.InScopeTimeSpan
	}
	if p.IsCustom {
		m["IsCustom"] = true
	}
	if p.IsPriorityContentOnlyScoring {
		m["IsPriorityContentOnlyScoring"] = true
	}
	if p.MLClassifierTypes != nil {
		m["MLClassifierTypes"] = p.MLClassifierTypes
	}
	if p.OptInDrpForDlp {
		m["OptInDrpForDlp"] = true
	}
	if p.PastTerminationWindow != 0 {
		m["PastTerminationWindow"] = p.PastTerminationWindow
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.PostTerminationActivity {
		m["PostTerminationActivity"] = true
	}
	if p.RemoveExchangeLocation != nil {
		m["RemoveExchangeLocation"] = p.RemoveExchangeLocation
	}
	if p.RemoveExchangeLocationException != nil {
		m["RemoveExchangeLocationException"] = p.RemoveExchangeLocationException
	}
	if p.RemoveIrmAdaptiveScopeLocation != nil {
		m["RemoveIrmAdaptiveScopeLocation"] = p.RemoveIrmAdaptiveScopeLocation
	}
	if p.RemoveModernGroupLocation != nil {
		m["RemoveModernGroupLocation"] = p.RemoveModernGroupLocation
	}
	if p.RemoveModernGroupLocationException != nil {
		m["RemoveModernGroupLocationException"] = p.RemoveModernGroupLocationException
	}
	if p.RetryDistribution {
		m["RetryDistribution"] = true
	}
	if p.SchemaVersion != 0 {
		m["SchemaVersion"] = p.SchemaVersion
	}
	if p.SensitivityLabels != nil {
		m["SensitivityLabels"] = p.SensitivityLabels
	}
	if p.SessionRecordingSettings != "" {
		m["SessionRecordingSettings"] = p.SessionRecordingSettings
	}
	if p.SharepointSites != nil {
		m["SharepointSites"] = p.SharepointSites
	}
	if p.TeamsSites != nil {
		m["TeamsSites"] = p.TeamsSites
	}
	if p.TenantSetting != "" {
		m["TenantSetting"] = p.TenantSetting
	}
	if p.TriggerInsightGroups != nil {
		m["TriggerInsightGroups"] = p.TriggerInsightGroups
	}
	if p.Triggers != nil {
		m["Triggers"] = p.Triggers
	}
	if p.TurnOnAnalytics != "" {
		m["TurnOnAnalytics"] = p.TurnOnAnalytics
	}
	if p.TurnOnDLPUserRiskSync != "" {
		m["TurnOnDLPUserRiskSync"] = p.TurnOnDLPUserRiskSync
	}
	return m
}

// SetInsiderRiskPolicy runs the Set-InsiderRiskPolicy cmdlet.
func (s *Service) SetInsiderRiskPolicy(ctx context.Context, p SetInsiderRiskPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-InsiderRiskPolicy", p.params())
}

// SetInsiderRiskPolicyLiteParams are the parameters of Set-InsiderRiskPolicyLite.
// DefaultParameterSetName: Identity
type SetInsiderRiskPolicyLiteParams struct {
	ExtensibleIndicators any  `ps:"ExtensibleIndicators"`
	Identity             any  `ps:"Identity"`
	Indicators           any  `ps:"Indicators"`
	IsCustom             bool `ps:"IsCustom"`
}

func (p SetInsiderRiskPolicyLiteParams) params() map[string]any {
	m := map[string]any{}
	if p.ExtensibleIndicators != nil {
		m["ExtensibleIndicators"] = p.ExtensibleIndicators
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Indicators != nil {
		m["Indicators"] = p.Indicators
	}
	if p.IsCustom {
		m["IsCustom"] = true
	}
	return m
}

// SetInsiderRiskPolicyLite runs the Set-InsiderRiskPolicyLite cmdlet.
func (s *Service) SetInsiderRiskPolicyLite(ctx context.Context, p SetInsiderRiskPolicyLiteParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-InsiderRiskPolicyLite", p.params())
}

// SetJitConfigurationParams are the parameters of Set-JitConfiguration.
// DefaultParameterSetName: Identity
type SetJitConfigurationParams struct {
	AddEndpointDlpLocation             any      `ps:"AddEndpointDlpLocation"`
	AddEndpointDlpLocationException    any      `ps:"AddEndpointDlpLocationException"`
	AddOneDriveLocation                any      `ps:"AddOneDriveLocation"`
	AddOneDriveLocationException       any      `ps:"AddOneDriveLocationException"`
	AddSharePointLocation              any      `ps:"AddSharePointLocation"`
	AddSharePointLocationException     any      `ps:"AddSharePointLocationException"`
	Comment                            string   `ps:"Comment"`
	EndpointConfig                     string   `ps:"EndpointConfig"`
	ExceptIfOneDriveSharedBy           []string `ps:"ExceptIfOneDriveSharedBy"`
	ExceptIfOneDriveSharedByMemberOf   []string `ps:"ExceptIfOneDriveSharedByMemberOf"`
	Identity                           any      `ps:"Identity"`
	Mode                               any      `ps:"Mode"`
	OneDriveSharedBy                   []string `ps:"OneDriveSharedBy"`
	OneDriveSharedByMemberOf           []string `ps:"OneDriveSharedByMemberOf"`
	RemoveEndpointDlpLocation          any      `ps:"RemoveEndpointDlpLocation"`
	RemoveEndpointDlpLocationException any      `ps:"RemoveEndpointDlpLocationException"`
	RemoveOneDriveLocation             any      `ps:"RemoveOneDriveLocation"`
	RemoveOneDriveLocationException    any      `ps:"RemoveOneDriveLocationException"`
	RemoveSharePointLocation           any      `ps:"RemoveSharePointLocation"`
	RemoveSharePointLocationException  any      `ps:"RemoveSharePointLocationException"`
}

func (p SetJitConfigurationParams) params() map[string]any {
	m := map[string]any{}
	if p.AddEndpointDlpLocation != nil {
		m["AddEndpointDlpLocation"] = p.AddEndpointDlpLocation
	}
	if p.AddEndpointDlpLocationException != nil {
		m["AddEndpointDlpLocationException"] = p.AddEndpointDlpLocationException
	}
	if p.AddOneDriveLocation != nil {
		m["AddOneDriveLocation"] = p.AddOneDriveLocation
	}
	if p.AddOneDriveLocationException != nil {
		m["AddOneDriveLocationException"] = p.AddOneDriveLocationException
	}
	if p.AddSharePointLocation != nil {
		m["AddSharePointLocation"] = p.AddSharePointLocation
	}
	if p.AddSharePointLocationException != nil {
		m["AddSharePointLocationException"] = p.AddSharePointLocationException
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.EndpointConfig != "" {
		m["EndpointConfig"] = p.EndpointConfig
	}
	if len(p.ExceptIfOneDriveSharedBy) > 0 {
		m["ExceptIfOneDriveSharedBy"] = p.ExceptIfOneDriveSharedBy
	}
	if len(p.ExceptIfOneDriveSharedByMemberOf) > 0 {
		m["ExceptIfOneDriveSharedByMemberOf"] = p.ExceptIfOneDriveSharedByMemberOf
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Mode != nil {
		m["Mode"] = p.Mode
	}
	if len(p.OneDriveSharedBy) > 0 {
		m["OneDriveSharedBy"] = p.OneDriveSharedBy
	}
	if len(p.OneDriveSharedByMemberOf) > 0 {
		m["OneDriveSharedByMemberOf"] = p.OneDriveSharedByMemberOf
	}
	if p.RemoveEndpointDlpLocation != nil {
		m["RemoveEndpointDlpLocation"] = p.RemoveEndpointDlpLocation
	}
	if p.RemoveEndpointDlpLocationException != nil {
		m["RemoveEndpointDlpLocationException"] = p.RemoveEndpointDlpLocationException
	}
	if p.RemoveOneDriveLocation != nil {
		m["RemoveOneDriveLocation"] = p.RemoveOneDriveLocation
	}
	if p.RemoveOneDriveLocationException != nil {
		m["RemoveOneDriveLocationException"] = p.RemoveOneDriveLocationException
	}
	if p.RemoveSharePointLocation != nil {
		m["RemoveSharePointLocation"] = p.RemoveSharePointLocation
	}
	if p.RemoveSharePointLocationException != nil {
		m["RemoveSharePointLocationException"] = p.RemoveSharePointLocationException
	}
	return m
}

// SetJitConfiguration runs the Set-JitConfiguration cmdlet.
func (s *Service) SetJitConfiguration(ctx context.Context, p SetJitConfigurationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-JitConfiguration", p.params())
}

// SetLabelParams are the parameters of Set-Label.
// DefaultParameterSetName: Identity
type SetLabelParams struct {
	AbacEnabled                                           any    `ps:"AbacEnabled"`
	AdvancedSettings                                      any    `ps:"AdvancedSettings"`
	ApplyContentMarkingFooterAlignment                    any    `ps:"ApplyContentMarkingFooterAlignment"`
	ApplyContentMarkingFooterEnabled                      any    `ps:"ApplyContentMarkingFooterEnabled"`
	ApplyContentMarkingFooterFontColor                    string `ps:"ApplyContentMarkingFooterFontColor"`
	ApplyContentMarkingFooterFontName                     string `ps:"ApplyContentMarkingFooterFontName"`
	ApplyContentMarkingFooterFontSize                     any    `ps:"ApplyContentMarkingFooterFontSize"`
	ApplyContentMarkingFooterMargin                       any    `ps:"ApplyContentMarkingFooterMargin"`
	ApplyContentMarkingFooterText                         string `ps:"ApplyContentMarkingFooterText"`
	ApplyContentMarkingHeaderAlignment                    any    `ps:"ApplyContentMarkingHeaderAlignment"`
	ApplyContentMarkingHeaderEnabled                      any    `ps:"ApplyContentMarkingHeaderEnabled"`
	ApplyContentMarkingHeaderFontColor                    string `ps:"ApplyContentMarkingHeaderFontColor"`
	ApplyContentMarkingHeaderFontName                     string `ps:"ApplyContentMarkingHeaderFontName"`
	ApplyContentMarkingHeaderFontSize                     any    `ps:"ApplyContentMarkingHeaderFontSize"`
	ApplyContentMarkingHeaderMargin                       any    `ps:"ApplyContentMarkingHeaderMargin"`
	ApplyContentMarkingHeaderText                         string `ps:"ApplyContentMarkingHeaderText"`
	ApplyDynamicWatermarkingEnabled                       any    `ps:"ApplyDynamicWatermarkingEnabled"`
	ApplyMessageMetaData                                  any    `ps:"ApplyMessageMetaData"`
	ApplyMessageMetaDataEnabled                           any    `ps:"ApplyMessageMetaDataEnabled"`
	ApplyWaterMarkingEnabled                              any    `ps:"ApplyWaterMarkingEnabled"`
	ApplyWaterMarkingFontColor                            string `ps:"ApplyWaterMarkingFontColor"`
	ApplyWaterMarkingFontName                             string `ps:"ApplyWaterMarkingFontName"`
	ApplyWaterMarkingFontSize                             any    `ps:"ApplyWaterMarkingFontSize"`
	ApplyWaterMarkingLayout                               any    `ps:"ApplyWaterMarkingLayout"`
	ApplyWaterMarkingText                                 string `ps:"ApplyWaterMarkingText"`
	ColumnAssetCondition                                  string `ps:"ColumnAssetCondition"`
	Comment                                               string `ps:"Comment"`
	ConditionalAccessAuthContext                          string `ps:"ConditionalAccessAuthContext"`
	Conditions                                            any    `ps:"Conditions"`
	ContentType                                           any    `ps:"ContentType"`
	DefaultContentLabel                                   string `ps:"DefaultContentLabel"`
	DenyAccessEnabled                                     any    `ps:"DenyAccessEnabled"`
	DenyAccessExcludedPrincipals                          string `ps:"DenyAccessExcludedPrincipals"`
	DenyAccessIncludedPrincipals                          string `ps:"DenyAccessIncludedPrincipals"`
	DenyAccessRights                                      string `ps:"DenyAccessRights"`
	DenyAccessWorkload                                    string `ps:"DenyAccessWorkload"`
	DisplayName                                           string `ps:"DisplayName"`
	DynamicWatermarkDisplay                               string `ps:"DynamicWatermarkDisplay"`
	EncryptionContentExpiredOnDateInDaysOrNever           string `ps:"EncryptionContentExpiredOnDateInDaysOrNever"`
	EncryptionDisableAutomaticOwnerRights                 any    `ps:"EncryptionDisableAutomaticOwnerRights"`
	EncryptionDoNotForward                                any    `ps:"EncryptionDoNotForward"`
	EncryptionDoubleKeyEncryptionUrl                      string `ps:"EncryptionDoubleKeyEncryptionUrl"`
	EncryptionEnabled                                     any    `ps:"EncryptionEnabled"`
	EncryptionEncryptOnly                                 any    `ps:"EncryptionEncryptOnly"`
	EncryptionOfflineAccessDays                           any    `ps:"EncryptionOfflineAccessDays"`
	EncryptionPromptUser                                  any    `ps:"EncryptionPromptUser"`
	EncryptionProtectionType                              any    `ps:"EncryptionProtectionType"`
	EncryptionRightsDefinitions                           any    `ps:"EncryptionRightsDefinitions"`
	EncryptionRightsUrl                                   string `ps:"EncryptionRightsUrl"`
	Identity                                              any    `ps:"Identity"`
	InformationProtectionAttributeRequired                any    `ps:"InformationProtectionAttributeRequired"`
	InheritToChildItems                                   any    `ps:"InheritToChildItems"`
	LabelActions                                          any    `ps:"LabelActions"`
	LocaleSettings                                        any    `ps:"LocaleSettings"`
	MigrationId                                           string `ps:"MigrationId"`
	NextLabel                                             any    `ps:"NextLabel"`
	ParentId                                              any    `ps:"ParentId"`
	PreviousLabel                                         any    `ps:"PreviousLabel"`
	Priority                                              any    `ps:"Priority"`
	RemoveParentLink                                      bool   `ps:"RemoveParentLink"`
	RuleBlob                                              string `ps:"RuleBlob"`
	SchematizedDataCondition                              string `ps:"SchematizedDataCondition"`
	Setting                                               any    `ps:"Setting"`
	Settings                                              any    `ps:"Settings"`
	SiteAndGroupProtectionAllowAccessToGuestUsers         any    `ps:"SiteAndGroupProtectionAllowAccessToGuestUsers"`
	SiteAndGroupProtectionAllowEmailFromGuestUsers        any    `ps:"SiteAndGroupProtectionAllowEmailFromGuestUsers"`
	SiteAndGroupProtectionAllowFullAccess                 any    `ps:"SiteAndGroupProtectionAllowFullAccess"`
	SiteAndGroupProtectionAllowLimitedAccess              any    `ps:"SiteAndGroupProtectionAllowLimitedAccess"`
	SiteAndGroupProtectionBlockAccess                     any    `ps:"SiteAndGroupProtectionBlockAccess"`
	SiteAndGroupProtectionEnabled                         any    `ps:"SiteAndGroupProtectionEnabled"`
	SiteAndGroupProtectionLevel                           any    `ps:"SiteAndGroupProtectionLevel"`
	SiteAndGroupProtectionPrivacy                         any    `ps:"SiteAndGroupProtectionPrivacy"`
	SiteExternalSharingControlType                        any    `ps:"SiteExternalSharingControlType"`
	TeamsAllowedPresenters                                any    `ps:"TeamsAllowedPresenters"`
	TeamsAllowMeetingChat                                 any    `ps:"TeamsAllowMeetingChat"`
	TeamsAllowPrivateTeamsToBeDiscoverableUsingSearch     any    `ps:"TeamsAllowPrivateTeamsToBeDiscoverableUsingSearch"`
	TeamsBypassLobbyForDialInUsers                        any    `ps:"TeamsBypassLobbyForDialInUsers"`
	TeamsChannelProtectionEnabled                         any    `ps:"TeamsChannelProtectionEnabled"`
	TeamsChannelSharedWithExternalTenants                 any    `ps:"TeamsChannelSharedWithExternalTenants"`
	TeamsChannelSharedWithPrivateTeamsOnly                any    `ps:"TeamsChannelSharedWithPrivateTeamsOnly"`
	TeamsChannelSharedWithSameLabelOnly                   any    `ps:"TeamsChannelSharedWithSameLabelOnly"`
	TeamsCopyRestrictionEnforced                          any    `ps:"TeamsCopyRestrictionEnforced"`
	TeamsDetectSensitiveContentDuringScreenSharingEnabled any    `ps:"TeamsDetectSensitiveContentDuringScreenSharingEnabled"`
	TeamsDisableLobby                                     any    `ps:"TeamsDisableLobby"`
	TeamsEndToEndEncryptionEnabled                        any    `ps:"TeamsEndToEndEncryptionEnabled"`
	TeamsLobbyBypassScope                                 any    `ps:"TeamsLobbyBypassScope"`
	TeamsLobbyRestrictionEnforced                         any    `ps:"TeamsLobbyRestrictionEnforced"`
	TeamsPresentersRestrictionEnforced                    any    `ps:"TeamsPresentersRestrictionEnforced"`
	TeamsProtectionEnabled                                any    `ps:"TeamsProtectionEnabled"`
	TeamsRecordAutomatically                              any    `ps:"TeamsRecordAutomatically"`
	TeamsVideoWatermark                                   any    `ps:"TeamsVideoWatermark"`
	TeamsWhoCanRecord                                     any    `ps:"TeamsWhoCanRecord"`
	Tooltip                                               string `ps:"Tooltip"`
}

func (p SetLabelParams) params() map[string]any {
	m := map[string]any{}
	if p.AbacEnabled != nil {
		m["AbacEnabled"] = p.AbacEnabled
	}
	if p.AdvancedSettings != nil {
		m["AdvancedSettings"] = p.AdvancedSettings
	}
	if p.ApplyContentMarkingFooterAlignment != nil {
		m["ApplyContentMarkingFooterAlignment"] = p.ApplyContentMarkingFooterAlignment
	}
	if p.ApplyContentMarkingFooterEnabled != nil {
		m["ApplyContentMarkingFooterEnabled"] = p.ApplyContentMarkingFooterEnabled
	}
	if p.ApplyContentMarkingFooterFontColor != "" {
		m["ApplyContentMarkingFooterFontColor"] = p.ApplyContentMarkingFooterFontColor
	}
	if p.ApplyContentMarkingFooterFontName != "" {
		m["ApplyContentMarkingFooterFontName"] = p.ApplyContentMarkingFooterFontName
	}
	if p.ApplyContentMarkingFooterFontSize != nil {
		m["ApplyContentMarkingFooterFontSize"] = p.ApplyContentMarkingFooterFontSize
	}
	if p.ApplyContentMarkingFooterMargin != nil {
		m["ApplyContentMarkingFooterMargin"] = p.ApplyContentMarkingFooterMargin
	}
	if p.ApplyContentMarkingFooterText != "" {
		m["ApplyContentMarkingFooterText"] = p.ApplyContentMarkingFooterText
	}
	if p.ApplyContentMarkingHeaderAlignment != nil {
		m["ApplyContentMarkingHeaderAlignment"] = p.ApplyContentMarkingHeaderAlignment
	}
	if p.ApplyContentMarkingHeaderEnabled != nil {
		m["ApplyContentMarkingHeaderEnabled"] = p.ApplyContentMarkingHeaderEnabled
	}
	if p.ApplyContentMarkingHeaderFontColor != "" {
		m["ApplyContentMarkingHeaderFontColor"] = p.ApplyContentMarkingHeaderFontColor
	}
	if p.ApplyContentMarkingHeaderFontName != "" {
		m["ApplyContentMarkingHeaderFontName"] = p.ApplyContentMarkingHeaderFontName
	}
	if p.ApplyContentMarkingHeaderFontSize != nil {
		m["ApplyContentMarkingHeaderFontSize"] = p.ApplyContentMarkingHeaderFontSize
	}
	if p.ApplyContentMarkingHeaderMargin != nil {
		m["ApplyContentMarkingHeaderMargin"] = p.ApplyContentMarkingHeaderMargin
	}
	if p.ApplyContentMarkingHeaderText != "" {
		m["ApplyContentMarkingHeaderText"] = p.ApplyContentMarkingHeaderText
	}
	if p.ApplyDynamicWatermarkingEnabled != nil {
		m["ApplyDynamicWatermarkingEnabled"] = p.ApplyDynamicWatermarkingEnabled
	}
	if p.ApplyMessageMetaData != nil {
		m["ApplyMessageMetaData"] = p.ApplyMessageMetaData
	}
	if p.ApplyMessageMetaDataEnabled != nil {
		m["ApplyMessageMetaDataEnabled"] = p.ApplyMessageMetaDataEnabled
	}
	if p.ApplyWaterMarkingEnabled != nil {
		m["ApplyWaterMarkingEnabled"] = p.ApplyWaterMarkingEnabled
	}
	if p.ApplyWaterMarkingFontColor != "" {
		m["ApplyWaterMarkingFontColor"] = p.ApplyWaterMarkingFontColor
	}
	if p.ApplyWaterMarkingFontName != "" {
		m["ApplyWaterMarkingFontName"] = p.ApplyWaterMarkingFontName
	}
	if p.ApplyWaterMarkingFontSize != nil {
		m["ApplyWaterMarkingFontSize"] = p.ApplyWaterMarkingFontSize
	}
	if p.ApplyWaterMarkingLayout != nil {
		m["ApplyWaterMarkingLayout"] = p.ApplyWaterMarkingLayout
	}
	if p.ApplyWaterMarkingText != "" {
		m["ApplyWaterMarkingText"] = p.ApplyWaterMarkingText
	}
	if p.ColumnAssetCondition != "" {
		m["ColumnAssetCondition"] = p.ColumnAssetCondition
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.ConditionalAccessAuthContext != "" {
		m["ConditionalAccessAuthContext"] = p.ConditionalAccessAuthContext
	}
	if p.Conditions != nil {
		m["Conditions"] = p.Conditions
	}
	if p.ContentType != nil {
		m["ContentType"] = p.ContentType
	}
	if p.DefaultContentLabel != "" {
		m["DefaultContentLabel"] = p.DefaultContentLabel
	}
	if p.DenyAccessEnabled != nil {
		m["DenyAccessEnabled"] = p.DenyAccessEnabled
	}
	if p.DenyAccessExcludedPrincipals != "" {
		m["DenyAccessExcludedPrincipals"] = p.DenyAccessExcludedPrincipals
	}
	if p.DenyAccessIncludedPrincipals != "" {
		m["DenyAccessIncludedPrincipals"] = p.DenyAccessIncludedPrincipals
	}
	if p.DenyAccessRights != "" {
		m["DenyAccessRights"] = p.DenyAccessRights
	}
	if p.DenyAccessWorkload != "" {
		m["DenyAccessWorkload"] = p.DenyAccessWorkload
	}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.DynamicWatermarkDisplay != "" {
		m["DynamicWatermarkDisplay"] = p.DynamicWatermarkDisplay
	}
	if p.EncryptionContentExpiredOnDateInDaysOrNever != "" {
		m["EncryptionContentExpiredOnDateInDaysOrNever"] = p.EncryptionContentExpiredOnDateInDaysOrNever
	}
	if p.EncryptionDisableAutomaticOwnerRights != nil {
		m["EncryptionDisableAutomaticOwnerRights"] = p.EncryptionDisableAutomaticOwnerRights
	}
	if p.EncryptionDoNotForward != nil {
		m["EncryptionDoNotForward"] = p.EncryptionDoNotForward
	}
	if p.EncryptionDoubleKeyEncryptionUrl != "" {
		m["EncryptionDoubleKeyEncryptionUrl"] = p.EncryptionDoubleKeyEncryptionUrl
	}
	if p.EncryptionEnabled != nil {
		m["EncryptionEnabled"] = p.EncryptionEnabled
	}
	if p.EncryptionEncryptOnly != nil {
		m["EncryptionEncryptOnly"] = p.EncryptionEncryptOnly
	}
	if p.EncryptionOfflineAccessDays != nil {
		m["EncryptionOfflineAccessDays"] = p.EncryptionOfflineAccessDays
	}
	if p.EncryptionPromptUser != nil {
		m["EncryptionPromptUser"] = p.EncryptionPromptUser
	}
	if p.EncryptionProtectionType != nil {
		m["EncryptionProtectionType"] = p.EncryptionProtectionType
	}
	if p.EncryptionRightsDefinitions != nil {
		m["EncryptionRightsDefinitions"] = p.EncryptionRightsDefinitions
	}
	if p.EncryptionRightsUrl != "" {
		m["EncryptionRightsUrl"] = p.EncryptionRightsUrl
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.InformationProtectionAttributeRequired != nil {
		m["InformationProtectionAttributeRequired"] = p.InformationProtectionAttributeRequired
	}
	if p.InheritToChildItems != nil {
		m["InheritToChildItems"] = p.InheritToChildItems
	}
	if p.LabelActions != nil {
		m["LabelActions"] = p.LabelActions
	}
	if p.LocaleSettings != nil {
		m["LocaleSettings"] = p.LocaleSettings
	}
	if p.MigrationId != "" {
		m["MigrationId"] = p.MigrationId
	}
	if p.NextLabel != nil {
		m["NextLabel"] = p.NextLabel
	}
	if p.ParentId != nil {
		m["ParentId"] = p.ParentId
	}
	if p.PreviousLabel != nil {
		m["PreviousLabel"] = p.PreviousLabel
	}
	if p.Priority != nil {
		m["Priority"] = p.Priority
	}
	if p.RemoveParentLink {
		m["RemoveParentLink"] = true
	}
	if p.RuleBlob != "" {
		m["RuleBlob"] = p.RuleBlob
	}
	if p.SchematizedDataCondition != "" {
		m["SchematizedDataCondition"] = p.SchematizedDataCondition
	}
	if p.Setting != nil {
		m["Setting"] = p.Setting
	}
	if p.Settings != nil {
		m["Settings"] = p.Settings
	}
	if p.SiteAndGroupProtectionAllowAccessToGuestUsers != nil {
		m["SiteAndGroupProtectionAllowAccessToGuestUsers"] = p.SiteAndGroupProtectionAllowAccessToGuestUsers
	}
	if p.SiteAndGroupProtectionAllowEmailFromGuestUsers != nil {
		m["SiteAndGroupProtectionAllowEmailFromGuestUsers"] = p.SiteAndGroupProtectionAllowEmailFromGuestUsers
	}
	if p.SiteAndGroupProtectionAllowFullAccess != nil {
		m["SiteAndGroupProtectionAllowFullAccess"] = p.SiteAndGroupProtectionAllowFullAccess
	}
	if p.SiteAndGroupProtectionAllowLimitedAccess != nil {
		m["SiteAndGroupProtectionAllowLimitedAccess"] = p.SiteAndGroupProtectionAllowLimitedAccess
	}
	if p.SiteAndGroupProtectionBlockAccess != nil {
		m["SiteAndGroupProtectionBlockAccess"] = p.SiteAndGroupProtectionBlockAccess
	}
	if p.SiteAndGroupProtectionEnabled != nil {
		m["SiteAndGroupProtectionEnabled"] = p.SiteAndGroupProtectionEnabled
	}
	if p.SiteAndGroupProtectionLevel != nil {
		m["SiteAndGroupProtectionLevel"] = p.SiteAndGroupProtectionLevel
	}
	if p.SiteAndGroupProtectionPrivacy != nil {
		m["SiteAndGroupProtectionPrivacy"] = p.SiteAndGroupProtectionPrivacy
	}
	if p.SiteExternalSharingControlType != nil {
		m["SiteExternalSharingControlType"] = p.SiteExternalSharingControlType
	}
	if p.TeamsAllowedPresenters != nil {
		m["TeamsAllowedPresenters"] = p.TeamsAllowedPresenters
	}
	if p.TeamsAllowMeetingChat != nil {
		m["TeamsAllowMeetingChat"] = p.TeamsAllowMeetingChat
	}
	if p.TeamsAllowPrivateTeamsToBeDiscoverableUsingSearch != nil {
		m["TeamsAllowPrivateTeamsToBeDiscoverableUsingSearch"] = p.TeamsAllowPrivateTeamsToBeDiscoverableUsingSearch
	}
	if p.TeamsBypassLobbyForDialInUsers != nil {
		m["TeamsBypassLobbyForDialInUsers"] = p.TeamsBypassLobbyForDialInUsers
	}
	if p.TeamsChannelProtectionEnabled != nil {
		m["TeamsChannelProtectionEnabled"] = p.TeamsChannelProtectionEnabled
	}
	if p.TeamsChannelSharedWithExternalTenants != nil {
		m["TeamsChannelSharedWithExternalTenants"] = p.TeamsChannelSharedWithExternalTenants
	}
	if p.TeamsChannelSharedWithPrivateTeamsOnly != nil {
		m["TeamsChannelSharedWithPrivateTeamsOnly"] = p.TeamsChannelSharedWithPrivateTeamsOnly
	}
	if p.TeamsChannelSharedWithSameLabelOnly != nil {
		m["TeamsChannelSharedWithSameLabelOnly"] = p.TeamsChannelSharedWithSameLabelOnly
	}
	if p.TeamsCopyRestrictionEnforced != nil {
		m["TeamsCopyRestrictionEnforced"] = p.TeamsCopyRestrictionEnforced
	}
	if p.TeamsDetectSensitiveContentDuringScreenSharingEnabled != nil {
		m["TeamsDetectSensitiveContentDuringScreenSharingEnabled"] = p.TeamsDetectSensitiveContentDuringScreenSharingEnabled
	}
	if p.TeamsDisableLobby != nil {
		m["TeamsDisableLobby"] = p.TeamsDisableLobby
	}
	if p.TeamsEndToEndEncryptionEnabled != nil {
		m["TeamsEndToEndEncryptionEnabled"] = p.TeamsEndToEndEncryptionEnabled
	}
	if p.TeamsLobbyBypassScope != nil {
		m["TeamsLobbyBypassScope"] = p.TeamsLobbyBypassScope
	}
	if p.TeamsLobbyRestrictionEnforced != nil {
		m["TeamsLobbyRestrictionEnforced"] = p.TeamsLobbyRestrictionEnforced
	}
	if p.TeamsPresentersRestrictionEnforced != nil {
		m["TeamsPresentersRestrictionEnforced"] = p.TeamsPresentersRestrictionEnforced
	}
	if p.TeamsProtectionEnabled != nil {
		m["TeamsProtectionEnabled"] = p.TeamsProtectionEnabled
	}
	if p.TeamsRecordAutomatically != nil {
		m["TeamsRecordAutomatically"] = p.TeamsRecordAutomatically
	}
	if p.TeamsVideoWatermark != nil {
		m["TeamsVideoWatermark"] = p.TeamsVideoWatermark
	}
	if p.TeamsWhoCanRecord != nil {
		m["TeamsWhoCanRecord"] = p.TeamsWhoCanRecord
	}
	if p.Tooltip != "" {
		m["Tooltip"] = p.Tooltip
	}
	return m
}

// SetLabel runs the Set-Label cmdlet.
func (s *Service) SetLabel(ctx context.Context, p SetLabelParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-Label", p.params())
}

// SetLabelExplorerConfigParams are the parameters of Set-LabelExplorerConfig.
type SetLabelExplorerConfigParams struct {
	Actions      any    `ps:"Actions"`
	AddLabels    any    `ps:"AddLabels"`
	AddUsers     any    `ps:"AddUsers"`
	DisplayName  string `ps:"DisplayName"`
	Filter       string `ps:"Filter"`
	Identity     any    `ps:"Identity"`
	Labels       any    `ps:"Labels"`
	RemoveLabels any    `ps:"RemoveLabels"`
	RemoveUsers  any    `ps:"RemoveUsers"`
	Users        any    `ps:"Users"`
}

func (p SetLabelExplorerConfigParams) params() map[string]any {
	m := map[string]any{}
	if p.Actions != nil {
		m["Actions"] = p.Actions
	}
	if p.AddLabels != nil {
		m["AddLabels"] = p.AddLabels
	}
	if p.AddUsers != nil {
		m["AddUsers"] = p.AddUsers
	}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.Filter != "" {
		m["Filter"] = p.Filter
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Labels != nil {
		m["Labels"] = p.Labels
	}
	if p.RemoveLabels != nil {
		m["RemoveLabels"] = p.RemoveLabels
	}
	if p.RemoveUsers != nil {
		m["RemoveUsers"] = p.RemoveUsers
	}
	if p.Users != nil {
		m["Users"] = p.Users
	}
	return m
}

// SetLabelExplorerConfig runs the Set-LabelExplorerConfig cmdlet.
func (s *Service) SetLabelExplorerConfig(ctx context.Context, p SetLabelExplorerConfigParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-LabelExplorerConfig", p.params())
}

// SetLabelPolicyParams are the parameters of Set-LabelPolicy.
// DefaultParameterSetName: Identity
type SetLabelPolicyParams struct {
	AddExchangeLocation                any    `ps:"AddExchangeLocation"`
	AddExchangeLocationException       any    `ps:"AddExchangeLocationException"`
	AddLabels                          any    `ps:"AddLabels"`
	AddModernGroupLocation             any    `ps:"AddModernGroupLocation"`
	AddModernGroupLocationException    any    `ps:"AddModernGroupLocationException"`
	AddOneDriveLocation                any    `ps:"AddOneDriveLocation"`
	AddOneDriveLocationException       any    `ps:"AddOneDriveLocationException"`
	AddPublicFolderLocation            any    `ps:"AddPublicFolderLocation"`
	AddSharePointLocation              any    `ps:"AddSharePointLocation"`
	AddSharePointLocationException     any    `ps:"AddSharePointLocationException"`
	AddSkypeLocation                   any    `ps:"AddSkypeLocation"`
	AddSkypeLocationException          any    `ps:"AddSkypeLocationException"`
	AdvancedSettings                   any    `ps:"AdvancedSettings"`
	Comment                            string `ps:"Comment"`
	ExchangeAdaptiveScopes             any    `ps:"ExchangeAdaptiveScopes"`
	ExchangeAdaptiveScopesException    any    `ps:"ExchangeAdaptiveScopesException"`
	Force                              bool   `ps:"Force"`
	Identity                           any    `ps:"Identity"`
	Locations                          string `ps:"Locations"`
	MigrationId                        string `ps:"MigrationId"`
	NextLabelPolicy                    any    `ps:"NextLabelPolicy"`
	PolicyRBACScopes                   any    `ps:"PolicyRBACScopes"`
	PreviousLabelPolicy                any    `ps:"PreviousLabelPolicy"`
	RemoveExchangeLocation             any    `ps:"RemoveExchangeLocation"`
	RemoveExchangeLocationException    any    `ps:"RemoveExchangeLocationException"`
	RemoveLabels                       any    `ps:"RemoveLabels"`
	RemoveModernGroupLocation          any    `ps:"RemoveModernGroupLocation"`
	RemoveModernGroupLocationException any    `ps:"RemoveModernGroupLocationException"`
	RemoveOneDriveLocation             any    `ps:"RemoveOneDriveLocation"`
	RemoveOneDriveLocationException    any    `ps:"RemoveOneDriveLocationException"`
	RemovePublicFolderLocation         any    `ps:"RemovePublicFolderLocation"`
	RemoveSharePointLocation           any    `ps:"RemoveSharePointLocation"`
	RemoveSharePointLocationException  any    `ps:"RemoveSharePointLocationException"`
	RemoveSkypeLocation                any    `ps:"RemoveSkypeLocation"`
	RemoveSkypeLocationException       any    `ps:"RemoveSkypeLocationException"`
	RetryDistribution                  bool   `ps:"RetryDistribution"`
	Setting                            any    `ps:"Setting"`
	Settings                           any    `ps:"Settings"`
}

func (p SetLabelPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.AddExchangeLocation != nil {
		m["AddExchangeLocation"] = p.AddExchangeLocation
	}
	if p.AddExchangeLocationException != nil {
		m["AddExchangeLocationException"] = p.AddExchangeLocationException
	}
	if p.AddLabels != nil {
		m["AddLabels"] = p.AddLabels
	}
	if p.AddModernGroupLocation != nil {
		m["AddModernGroupLocation"] = p.AddModernGroupLocation
	}
	if p.AddModernGroupLocationException != nil {
		m["AddModernGroupLocationException"] = p.AddModernGroupLocationException
	}
	if p.AddOneDriveLocation != nil {
		m["AddOneDriveLocation"] = p.AddOneDriveLocation
	}
	if p.AddOneDriveLocationException != nil {
		m["AddOneDriveLocationException"] = p.AddOneDriveLocationException
	}
	if p.AddPublicFolderLocation != nil {
		m["AddPublicFolderLocation"] = p.AddPublicFolderLocation
	}
	if p.AddSharePointLocation != nil {
		m["AddSharePointLocation"] = p.AddSharePointLocation
	}
	if p.AddSharePointLocationException != nil {
		m["AddSharePointLocationException"] = p.AddSharePointLocationException
	}
	if p.AddSkypeLocation != nil {
		m["AddSkypeLocation"] = p.AddSkypeLocation
	}
	if p.AddSkypeLocationException != nil {
		m["AddSkypeLocationException"] = p.AddSkypeLocationException
	}
	if p.AdvancedSettings != nil {
		m["AdvancedSettings"] = p.AdvancedSettings
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.ExchangeAdaptiveScopes != nil {
		m["ExchangeAdaptiveScopes"] = p.ExchangeAdaptiveScopes
	}
	if p.ExchangeAdaptiveScopesException != nil {
		m["ExchangeAdaptiveScopesException"] = p.ExchangeAdaptiveScopesException
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Locations != "" {
		m["Locations"] = p.Locations
	}
	if p.MigrationId != "" {
		m["MigrationId"] = p.MigrationId
	}
	if p.NextLabelPolicy != nil {
		m["NextLabelPolicy"] = p.NextLabelPolicy
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.PreviousLabelPolicy != nil {
		m["PreviousLabelPolicy"] = p.PreviousLabelPolicy
	}
	if p.RemoveExchangeLocation != nil {
		m["RemoveExchangeLocation"] = p.RemoveExchangeLocation
	}
	if p.RemoveExchangeLocationException != nil {
		m["RemoveExchangeLocationException"] = p.RemoveExchangeLocationException
	}
	if p.RemoveLabels != nil {
		m["RemoveLabels"] = p.RemoveLabels
	}
	if p.RemoveModernGroupLocation != nil {
		m["RemoveModernGroupLocation"] = p.RemoveModernGroupLocation
	}
	if p.RemoveModernGroupLocationException != nil {
		m["RemoveModernGroupLocationException"] = p.RemoveModernGroupLocationException
	}
	if p.RemoveOneDriveLocation != nil {
		m["RemoveOneDriveLocation"] = p.RemoveOneDriveLocation
	}
	if p.RemoveOneDriveLocationException != nil {
		m["RemoveOneDriveLocationException"] = p.RemoveOneDriveLocationException
	}
	if p.RemovePublicFolderLocation != nil {
		m["RemovePublicFolderLocation"] = p.RemovePublicFolderLocation
	}
	if p.RemoveSharePointLocation != nil {
		m["RemoveSharePointLocation"] = p.RemoveSharePointLocation
	}
	if p.RemoveSharePointLocationException != nil {
		m["RemoveSharePointLocationException"] = p.RemoveSharePointLocationException
	}
	if p.RemoveSkypeLocation != nil {
		m["RemoveSkypeLocation"] = p.RemoveSkypeLocation
	}
	if p.RemoveSkypeLocationException != nil {
		m["RemoveSkypeLocationException"] = p.RemoveSkypeLocationException
	}
	if p.RetryDistribution {
		m["RetryDistribution"] = true
	}
	if p.Setting != nil {
		m["Setting"] = p.Setting
	}
	if p.Settings != nil {
		m["Settings"] = p.Settings
	}
	return m
}

// SetLabelPolicy runs the Set-LabelPolicy cmdlet.
func (s *Service) SetLabelPolicy(ctx context.Context, p SetLabelPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-LabelPolicy", p.params())
}

// SetMachineAssistedTagResourceParams are the parameters of Set-MachineAssistedTagResource.
type SetMachineAssistedTagResourceParams struct {
	DiscardModel    string `ps:"DiscardModel"`
	Feedbacks       string `ps:"Feedbacks"`
	Identity        any    `ps:"Identity"`
	ModelId         any    `ps:"ModelId"`
	ModelName       string `ps:"ModelName"`
	Resource        any    `ps:"Resource"`
	ReviewCompleted string `ps:"ReviewCompleted"`
	TaskId          any    `ps:"TaskId"`
	TrainingId      any    `ps:"TrainingId"`
}

func (p SetMachineAssistedTagResourceParams) params() map[string]any {
	m := map[string]any{}
	if p.DiscardModel != "" {
		m["DiscardModel"] = p.DiscardModel
	}
	if p.Feedbacks != "" {
		m["Feedbacks"] = p.Feedbacks
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.ModelId != nil {
		m["ModelId"] = p.ModelId
	}
	if p.ModelName != "" {
		m["ModelName"] = p.ModelName
	}
	if p.Resource != nil {
		m["Resource"] = p.Resource
	}
	if p.ReviewCompleted != "" {
		m["ReviewCompleted"] = p.ReviewCompleted
	}
	if p.TaskId != nil {
		m["TaskId"] = p.TaskId
	}
	if p.TrainingId != nil {
		m["TrainingId"] = p.TrainingId
	}
	return m
}

// SetMachineAssistedTagResource runs the Set-MachineAssistedTagResource cmdlet.
func (s *Service) SetMachineAssistedTagResource(ctx context.Context, p SetMachineAssistedTagResourceParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-MachineAssistedTagResource", p.params())
}

// SetMultiTenantOrganizationConfigParams are the parameters of Set-MultiTenantOrganizationConfig.
// DefaultParameterSetName: Identity
type SetMultiTenantOrganizationConfigParams struct {
	AddAuthorizedMemberTenants           []string `ps:"AddAuthorizedMemberTenants"`
	AuthorizedOwnerTenantForMemberTenant any      `ps:"AuthorizedOwnerTenantForMemberTenant"`
	Identity                             any      `ps:"Identity"`
	IsMultiTenantOrganizationEnabled     any      `ps:"IsMultiTenantOrganizationEnabled"`
	RemoveAuthorizedMemberTenants        []string `ps:"RemoveAuthorizedMemberTenants"`
}

func (p SetMultiTenantOrganizationConfigParams) params() map[string]any {
	m := map[string]any{}
	if len(p.AddAuthorizedMemberTenants) > 0 {
		m["AddAuthorizedMemberTenants"] = p.AddAuthorizedMemberTenants
	}
	if p.AuthorizedOwnerTenantForMemberTenant != nil {
		m["AuthorizedOwnerTenantForMemberTenant"] = p.AuthorizedOwnerTenantForMemberTenant
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IsMultiTenantOrganizationEnabled != nil {
		m["IsMultiTenantOrganizationEnabled"] = p.IsMultiTenantOrganizationEnabled
	}
	if len(p.RemoveAuthorizedMemberTenants) > 0 {
		m["RemoveAuthorizedMemberTenants"] = p.RemoveAuthorizedMemberTenants
	}
	return m
}

// SetMultiTenantOrganizationConfig runs the Set-MultiTenantOrganizationConfig cmdlet.
func (s *Service) SetMultiTenantOrganizationConfig(ctx context.Context, p SetMultiTenantOrganizationConfigParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-MultiTenantOrganizationConfig", p.params())
}

// SetOcrConfigurationParams are the parameters of Set-OcrConfiguration.
// DefaultParameterSetName: Identity
type SetOcrConfigurationParams struct {
	AddEndpointDlpLocation               any      `ps:"AddEndpointDlpLocation"`
	AddEndpointDlpLocationException      any      `ps:"AddEndpointDlpLocationException"`
	AddExchangeLocation                  any      `ps:"AddExchangeLocation"`
	AddExchangeLocationException         any      `ps:"AddExchangeLocationException"`
	AddOneDriveLocation                  any      `ps:"AddOneDriveLocation"`
	AddOneDriveLocationException         any      `ps:"AddOneDriveLocationException"`
	AddSharePointLocation                any      `ps:"AddSharePointLocation"`
	AddSharePointLocationException       any      `ps:"AddSharePointLocationException"`
	AddTeamsLocation                     any      `ps:"AddTeamsLocation"`
	AddTeamsLocationException            any      `ps:"AddTeamsLocationException"`
	Comment                              string   `ps:"Comment"`
	EnforcementPlanes                    any      `ps:"EnforcementPlanes"`
	ExceptIfOneDriveSharedBy             []string `ps:"ExceptIfOneDriveSharedBy"`
	ExceptIfOneDriveSharedByMemberOf     []string `ps:"ExceptIfOneDriveSharedByMemberOf"`
	ExchangeScopeOcrAnyRecipientExternal bool     `ps:"ExchangeScopeOcrAnyRecipientExternal"`
	ExchangeSender                       []string `ps:"ExchangeSender"`
	ExchangeSenderException              []string `ps:"ExchangeSenderException"`
	ExchangeSenderMemberOf               []string `ps:"ExchangeSenderMemberOf"`
	ExchangeSenderMemberOfException      []string `ps:"ExchangeSenderMemberOfException"`
	Identity                             any      `ps:"Identity"`
	Locations                            string   `ps:"Locations"`
	Mode                                 any      `ps:"Mode"`
	OcrMode                              any      `ps:"OcrMode"`
	OneDriveSharedBy                     []string `ps:"OneDriveSharedBy"`
	OneDriveSharedByMemberOf             []string `ps:"OneDriveSharedByMemberOf"`
	RemoveEndpointDlpLocation            any      `ps:"RemoveEndpointDlpLocation"`
	RemoveEndpointDlpLocationException   any      `ps:"RemoveEndpointDlpLocationException"`
	RemoveExchangeLocation               any      `ps:"RemoveExchangeLocation"`
	RemoveExchangeLocationException      any      `ps:"RemoveExchangeLocationException"`
	RemoveOneDriveLocation               any      `ps:"RemoveOneDriveLocation"`
	RemoveOneDriveLocationException      any      `ps:"RemoveOneDriveLocationException"`
	RemoveSharePointLocation             any      `ps:"RemoveSharePointLocation"`
	RemoveSharePointLocationException    any      `ps:"RemoveSharePointLocationException"`
	RemoveTeamsLocation                  any      `ps:"RemoveTeamsLocation"`
	RemoveTeamsLocationException         any      `ps:"RemoveTeamsLocationException"`
}

func (p SetOcrConfigurationParams) params() map[string]any {
	m := map[string]any{}
	if p.AddEndpointDlpLocation != nil {
		m["AddEndpointDlpLocation"] = p.AddEndpointDlpLocation
	}
	if p.AddEndpointDlpLocationException != nil {
		m["AddEndpointDlpLocationException"] = p.AddEndpointDlpLocationException
	}
	if p.AddExchangeLocation != nil {
		m["AddExchangeLocation"] = p.AddExchangeLocation
	}
	if p.AddExchangeLocationException != nil {
		m["AddExchangeLocationException"] = p.AddExchangeLocationException
	}
	if p.AddOneDriveLocation != nil {
		m["AddOneDriveLocation"] = p.AddOneDriveLocation
	}
	if p.AddOneDriveLocationException != nil {
		m["AddOneDriveLocationException"] = p.AddOneDriveLocationException
	}
	if p.AddSharePointLocation != nil {
		m["AddSharePointLocation"] = p.AddSharePointLocation
	}
	if p.AddSharePointLocationException != nil {
		m["AddSharePointLocationException"] = p.AddSharePointLocationException
	}
	if p.AddTeamsLocation != nil {
		m["AddTeamsLocation"] = p.AddTeamsLocation
	}
	if p.AddTeamsLocationException != nil {
		m["AddTeamsLocationException"] = p.AddTeamsLocationException
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.EnforcementPlanes != nil {
		m["EnforcementPlanes"] = p.EnforcementPlanes
	}
	if len(p.ExceptIfOneDriveSharedBy) > 0 {
		m["ExceptIfOneDriveSharedBy"] = p.ExceptIfOneDriveSharedBy
	}
	if len(p.ExceptIfOneDriveSharedByMemberOf) > 0 {
		m["ExceptIfOneDriveSharedByMemberOf"] = p.ExceptIfOneDriveSharedByMemberOf
	}
	if p.ExchangeScopeOcrAnyRecipientExternal {
		m["ExchangeScopeOcrAnyRecipientExternal"] = true
	}
	if len(p.ExchangeSender) > 0 {
		m["ExchangeSender"] = p.ExchangeSender
	}
	if len(p.ExchangeSenderException) > 0 {
		m["ExchangeSenderException"] = p.ExchangeSenderException
	}
	if len(p.ExchangeSenderMemberOf) > 0 {
		m["ExchangeSenderMemberOf"] = p.ExchangeSenderMemberOf
	}
	if len(p.ExchangeSenderMemberOfException) > 0 {
		m["ExchangeSenderMemberOfException"] = p.ExchangeSenderMemberOfException
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Locations != "" {
		m["Locations"] = p.Locations
	}
	if p.Mode != nil {
		m["Mode"] = p.Mode
	}
	if p.OcrMode != nil {
		m["OcrMode"] = p.OcrMode
	}
	if len(p.OneDriveSharedBy) > 0 {
		m["OneDriveSharedBy"] = p.OneDriveSharedBy
	}
	if len(p.OneDriveSharedByMemberOf) > 0 {
		m["OneDriveSharedByMemberOf"] = p.OneDriveSharedByMemberOf
	}
	if p.RemoveEndpointDlpLocation != nil {
		m["RemoveEndpointDlpLocation"] = p.RemoveEndpointDlpLocation
	}
	if p.RemoveEndpointDlpLocationException != nil {
		m["RemoveEndpointDlpLocationException"] = p.RemoveEndpointDlpLocationException
	}
	if p.RemoveExchangeLocation != nil {
		m["RemoveExchangeLocation"] = p.RemoveExchangeLocation
	}
	if p.RemoveExchangeLocationException != nil {
		m["RemoveExchangeLocationException"] = p.RemoveExchangeLocationException
	}
	if p.RemoveOneDriveLocation != nil {
		m["RemoveOneDriveLocation"] = p.RemoveOneDriveLocation
	}
	if p.RemoveOneDriveLocationException != nil {
		m["RemoveOneDriveLocationException"] = p.RemoveOneDriveLocationException
	}
	if p.RemoveSharePointLocation != nil {
		m["RemoveSharePointLocation"] = p.RemoveSharePointLocation
	}
	if p.RemoveSharePointLocationException != nil {
		m["RemoveSharePointLocationException"] = p.RemoveSharePointLocationException
	}
	if p.RemoveTeamsLocation != nil {
		m["RemoveTeamsLocation"] = p.RemoveTeamsLocation
	}
	if p.RemoveTeamsLocationException != nil {
		m["RemoveTeamsLocationException"] = p.RemoveTeamsLocationException
	}
	return m
}

// SetOcrConfiguration runs the Set-OcrConfiguration cmdlet.
func (s *Service) SetOcrConfiguration(ctx context.Context, p SetOcrConfigurationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-OcrConfiguration", p.params())
}

// SetOrganizationSegmentParams are the parameters of Set-OrganizationSegment.
// DefaultParameterSetName: OrganizationSegmentsDefault
type SetOrganizationSegmentParams struct {
	Identity        any    `ps:"Identity"`
	UserGroupFilter string `ps:"UserGroupFilter"`
}

func (p SetOrganizationSegmentParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.UserGroupFilter != "" {
		m["UserGroupFilter"] = p.UserGroupFilter
	}
	return m
}

// SetOrganizationSegment runs the Set-OrganizationSegment cmdlet.
func (s *Service) SetOrganizationSegment(ctx context.Context, p SetOrganizationSegmentParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-OrganizationSegment", p.params())
}

// SetPolicyConfigParams are the parameters of Set-PolicyConfig.
// DefaultParameterSetName: Identity
type SetPolicyConfigParams struct {
	AggregationTimeWindowForDlpAlerts          any      `ps:"AggregationTimeWindowForDlpAlerts"`
	CaseHoldPolicyLimit                        int      `ps:"CaseHoldPolicyLimit"`
	ClassificationScheme                       any      `ps:"ClassificationScheme"`
	ComplianceUrl                              string   `ps:"ComplianceUrl"`
	DlpAlertFoldingConfiguration               any      `ps:"DlpAlertFoldingConfiguration"`
	DlpAlertSlaConfig                          string   `ps:"DlpAlertSlaConfig"`
	DlpAppGroups                               []string `ps:"DlpAppGroups"`
	DlpAppGroupsPsws                           []string `ps:"DlpAppGroupsPsws"`
	DlpErrorHandlingConfig                     string   `ps:"DlpErrorHandlingConfig"`
	DlpExtensionGroups                         any      `ps:"DlpExtensionGroups"`
	DlpNetworkShareGroups                      any      `ps:"DlpNetworkShareGroups"`
	DlpPrinterGroups                           any      `ps:"DlpPrinterGroups"`
	DlpRemovableMediaGroups                    any      `ps:"DlpRemovableMediaGroups"`
	DocumentIsUnsupportedSeverity              any      `ps:"DocumentIsUnsupportedSeverity"`
	EnableAbacFeature                          bool     `ps:"EnableAbacFeature"`
	EnableAdvancedRuleBuilder                  bool     `ps:"EnableAdvancedRuleBuilder"`
	EnableLabelCoauth                          bool     `ps:"EnableLabelCoauth"`
	EnableSensitivityLabelingForPdf            bool     `ps:"EnableSensitivityLabelingForPdf"`
	EnableSpoAipMigration                      bool     `ps:"EnableSpoAipMigration"`
	EndpointDlpGlobalSettings                  []string `ps:"EndpointDlpGlobalSettings"`
	EndpointDlpGlobalSettingsPsws              []string `ps:"EndpointDlpGlobalSettingsPsws"`
	ExtendTeamsDlpPoliciesToSharePointOneDrive bool     `ps:"ExtendTeamsDlpPoliciesToSharePointOneDrive"`
	Identity                                   any      `ps:"Identity"`
	InformationBarrierMode                     any      `ps:"InformationBarrierMode"`
	InformationBarrierPeopleSearchRestriction  any      `ps:"InformationBarrierPeopleSearchRestriction"`
	IsDefaultMlDlpPoliciesOptedOut             bool     `ps:"IsDefaultMlDlpPoliciesOptedOut"`
	IsDlpSimulationOptedIn                     bool     `ps:"IsDlpSimulationOptedIn"`
	IsEventFoldingOnDlpAlertsOptedOut          bool     `ps:"IsEventFoldingOnDlpAlertsOptedOut"`
	IsUserBaseDlpAlertAggregationEnabled       bool     `ps:"IsUserBaseDlpAlertAggregationEnabled"`
	JitEnforcementSettings                     string   `ps:"JitEnforcementSettings"`
	MessageHeadersToRetainInOutlook            []string `ps:"MessageHeadersToRetainInOutlook"`
	MigrateLabelScheme                         bool     `ps:"MigrateLabelScheme"`
	MigrateLabelSchemeDisplayNames             any      `ps:"MigrateLabelSchemeDisplayNames"`
	OnPremisesWorkload                         any      `ps:"OnPremisesWorkload"`
	ProcessingLimitExceededSeverity            any      `ps:"ProcessingLimitExceededSeverity"`
	PurviewLabelConsent                        bool     `ps:"PurviewLabelConsent"`
	ReservedForFutureUse                       bool     `ps:"ReservedForFutureUse"`
	RetentionForwardCrawl                      bool     `ps:"RetentionForwardCrawl"`
	RuleErrorAction                            any      `ps:"RuleErrorAction"`
	SenderAddressLocation                      any      `ps:"SenderAddressLocation"`
	SiteGroups                                 []string `ps:"SiteGroups"`
	SiteGroupsPsws                             []string `ps:"SiteGroupsPsws"`
	SpoQuarantineLocation                      string   `ps:"SpoQuarantineLocation"`
	SpoQuarantineText                          string   `ps:"SpoQuarantineText"`
	TextExtractionConfig                       string   `ps:"TextExtractionConfig"`
}

func (p SetPolicyConfigParams) params() map[string]any {
	m := map[string]any{}
	if p.AggregationTimeWindowForDlpAlerts != nil {
		m["AggregationTimeWindowForDlpAlerts"] = p.AggregationTimeWindowForDlpAlerts
	}
	if p.CaseHoldPolicyLimit != 0 {
		m["CaseHoldPolicyLimit"] = p.CaseHoldPolicyLimit
	}
	if p.ClassificationScheme != nil {
		m["ClassificationScheme"] = p.ClassificationScheme
	}
	if p.ComplianceUrl != "" {
		m["ComplianceUrl"] = p.ComplianceUrl
	}
	if p.DlpAlertFoldingConfiguration != nil {
		m["DlpAlertFoldingConfiguration"] = p.DlpAlertFoldingConfiguration
	}
	if p.DlpAlertSlaConfig != "" {
		m["DlpAlertSlaConfig"] = p.DlpAlertSlaConfig
	}
	if len(p.DlpAppGroups) > 0 {
		m["DlpAppGroups"] = p.DlpAppGroups
	}
	if len(p.DlpAppGroupsPsws) > 0 {
		m["DlpAppGroupsPsws"] = p.DlpAppGroupsPsws
	}
	if p.DlpErrorHandlingConfig != "" {
		m["DlpErrorHandlingConfig"] = p.DlpErrorHandlingConfig
	}
	if p.DlpExtensionGroups != nil {
		m["DlpExtensionGroups"] = p.DlpExtensionGroups
	}
	if p.DlpNetworkShareGroups != nil {
		m["DlpNetworkShareGroups"] = p.DlpNetworkShareGroups
	}
	if p.DlpPrinterGroups != nil {
		m["DlpPrinterGroups"] = p.DlpPrinterGroups
	}
	if p.DlpRemovableMediaGroups != nil {
		m["DlpRemovableMediaGroups"] = p.DlpRemovableMediaGroups
	}
	if p.DocumentIsUnsupportedSeverity != nil {
		m["DocumentIsUnsupportedSeverity"] = p.DocumentIsUnsupportedSeverity
	}
	if p.EnableAbacFeature {
		m["EnableAbacFeature"] = true
	}
	if p.EnableAdvancedRuleBuilder {
		m["EnableAdvancedRuleBuilder"] = true
	}
	if p.EnableLabelCoauth {
		m["EnableLabelCoauth"] = true
	}
	if p.EnableSensitivityLabelingForPdf {
		m["EnableSensitivityLabelingForPdf"] = true
	}
	if p.EnableSpoAipMigration {
		m["EnableSpoAipMigration"] = true
	}
	if len(p.EndpointDlpGlobalSettings) > 0 {
		m["EndpointDlpGlobalSettings"] = p.EndpointDlpGlobalSettings
	}
	if len(p.EndpointDlpGlobalSettingsPsws) > 0 {
		m["EndpointDlpGlobalSettingsPsws"] = p.EndpointDlpGlobalSettingsPsws
	}
	if p.ExtendTeamsDlpPoliciesToSharePointOneDrive {
		m["ExtendTeamsDlpPoliciesToSharePointOneDrive"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.InformationBarrierMode != nil {
		m["InformationBarrierMode"] = p.InformationBarrierMode
	}
	if p.InformationBarrierPeopleSearchRestriction != nil {
		m["InformationBarrierPeopleSearchRestriction"] = p.InformationBarrierPeopleSearchRestriction
	}
	if p.IsDefaultMlDlpPoliciesOptedOut {
		m["IsDefaultMlDlpPoliciesOptedOut"] = true
	}
	if p.IsDlpSimulationOptedIn {
		m["IsDlpSimulationOptedIn"] = true
	}
	if p.IsEventFoldingOnDlpAlertsOptedOut {
		m["IsEventFoldingOnDlpAlertsOptedOut"] = true
	}
	if p.IsUserBaseDlpAlertAggregationEnabled {
		m["IsUserBaseDlpAlertAggregationEnabled"] = true
	}
	if p.JitEnforcementSettings != "" {
		m["JitEnforcementSettings"] = p.JitEnforcementSettings
	}
	if len(p.MessageHeadersToRetainInOutlook) > 0 {
		m["MessageHeadersToRetainInOutlook"] = p.MessageHeadersToRetainInOutlook
	}
	if p.MigrateLabelScheme {
		m["MigrateLabelScheme"] = true
	}
	if p.MigrateLabelSchemeDisplayNames != nil {
		m["MigrateLabelSchemeDisplayNames"] = p.MigrateLabelSchemeDisplayNames
	}
	if p.OnPremisesWorkload != nil {
		m["OnPremisesWorkload"] = p.OnPremisesWorkload
	}
	if p.ProcessingLimitExceededSeverity != nil {
		m["ProcessingLimitExceededSeverity"] = p.ProcessingLimitExceededSeverity
	}
	if p.PurviewLabelConsent {
		m["PurviewLabelConsent"] = true
	}
	if p.ReservedForFutureUse {
		m["ReservedForFutureUse"] = true
	}
	if p.RetentionForwardCrawl {
		m["RetentionForwardCrawl"] = true
	}
	if p.RuleErrorAction != nil {
		m["RuleErrorAction"] = p.RuleErrorAction
	}
	if p.SenderAddressLocation != nil {
		m["SenderAddressLocation"] = p.SenderAddressLocation
	}
	if len(p.SiteGroups) > 0 {
		m["SiteGroups"] = p.SiteGroups
	}
	if len(p.SiteGroupsPsws) > 0 {
		m["SiteGroupsPsws"] = p.SiteGroupsPsws
	}
	if p.SpoQuarantineLocation != "" {
		m["SpoQuarantineLocation"] = p.SpoQuarantineLocation
	}
	if p.SpoQuarantineText != "" {
		m["SpoQuarantineText"] = p.SpoQuarantineText
	}
	if p.TextExtractionConfig != "" {
		m["TextExtractionConfig"] = p.TextExtractionConfig
	}
	return m
}

// SetPolicyConfig runs the Set-PolicyConfig cmdlet.
func (s *Service) SetPolicyConfig(ctx context.Context, p SetPolicyConfigParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-PolicyConfig", p.params())
}

// SetPrivacyManagementComplianceTagParams are the parameters of Set-PrivacyManagementComplianceTag.
// DefaultParameterSetName: Identity
type SetPrivacyManagementComplianceTagParams struct {
	Comment           string `ps:"Comment"`
	Force             bool   `ps:"Force"`
	Identity          any    `ps:"Identity"`
	RetentionAction   string `ps:"RetentionAction"`
	RetentionDuration any    `ps:"RetentionDuration"`
}

func (p SetPrivacyManagementComplianceTagParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.RetentionAction != "" {
		m["RetentionAction"] = p.RetentionAction
	}
	if p.RetentionDuration != nil {
		m["RetentionDuration"] = p.RetentionDuration
	}
	return m
}

// SetPrivacyManagementComplianceTag runs the Set-PrivacyManagementComplianceTag cmdlet.
func (s *Service) SetPrivacyManagementComplianceTag(ctx context.Context, p SetPrivacyManagementComplianceTagParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-PrivacyManagementComplianceTag", p.params())
}

// SetPrivacyManagementPolicyParams are the parameters of Set-PrivacyManagementPolicy.
// DefaultParameterSetName: Identity
type SetPrivacyManagementPolicyParams struct {
	Comment  string `ps:"Comment"`
	Enabled  bool   `ps:"Enabled"`
	Identity any    `ps:"Identity"`
	Mode     any    `ps:"Mode"`
}

func (p SetPrivacyManagementPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Mode != nil {
		m["Mode"] = p.Mode
	}
	return m
}

// SetPrivacyManagementPolicy runs the Set-PrivacyManagementPolicy cmdlet.
func (s *Service) SetPrivacyManagementPolicy(ctx context.Context, p SetPrivacyManagementPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-PrivacyManagementPolicy", p.params())
}

// SetPrivacyManagementRuleParams are the parameters of Set-PrivacyManagementRule.
// DefaultParameterSetName: Identity
type SetPrivacyManagementRuleParams struct {
	Comment                             string   `ps:"Comment"`
	ContentContainsSensitiveInformation []string `ps:"ContentContainsSensitiveInformation"`
	CrossBoundaryTransfers              []string `ps:"CrossBoundaryTransfers"`
	Disabled                            bool     `ps:"Disabled"`
	EndpointOperations                  []string `ps:"EndpointOperations"`
	ExchangeSites                       any      `ps:"ExchangeSites"`
	Groups                              any      `ps:"Groups"`
	Identity                            any      `ps:"Identity"`
	LastModifiedThresholdInDays         any      `ps:"LastModifiedThresholdInDays"`
	Locations                           string   `ps:"Locations"`
	OnedriveSites                       any      `ps:"OnedriveSites"`
	PolicySettings                      string   `ps:"PolicySettings"`
	PrivacyAccessLevel                  []string `ps:"PrivacyAccessLevel"`
	PurviewSites                        any      `ps:"PurviewSites"`
	SharepointSiteOversharingEnabled    any      `ps:"SharepointSiteOversharingEnabled"`
	SharepointSites                     any      `ps:"SharepointSites"`
	TeamsSites                          any      `ps:"TeamsSites"`
	TenantSettings                      string   `ps:"TenantSettings"`
	Users                               any      `ps:"Users"`
}

func (p SetPrivacyManagementRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if len(p.ContentContainsSensitiveInformation) > 0 {
		m["ContentContainsSensitiveInformation"] = p.ContentContainsSensitiveInformation
	}
	if len(p.CrossBoundaryTransfers) > 0 {
		m["CrossBoundaryTransfers"] = p.CrossBoundaryTransfers
	}
	if p.Disabled {
		m["Disabled"] = true
	}
	if len(p.EndpointOperations) > 0 {
		m["EndpointOperations"] = p.EndpointOperations
	}
	if p.ExchangeSites != nil {
		m["ExchangeSites"] = p.ExchangeSites
	}
	if p.Groups != nil {
		m["Groups"] = p.Groups
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.LastModifiedThresholdInDays != nil {
		m["LastModifiedThresholdInDays"] = p.LastModifiedThresholdInDays
	}
	if p.Locations != "" {
		m["Locations"] = p.Locations
	}
	if p.OnedriveSites != nil {
		m["OnedriveSites"] = p.OnedriveSites
	}
	if p.PolicySettings != "" {
		m["PolicySettings"] = p.PolicySettings
	}
	if len(p.PrivacyAccessLevel) > 0 {
		m["PrivacyAccessLevel"] = p.PrivacyAccessLevel
	}
	if p.PurviewSites != nil {
		m["PurviewSites"] = p.PurviewSites
	}
	if p.SharepointSiteOversharingEnabled != nil {
		m["SharepointSiteOversharingEnabled"] = p.SharepointSiteOversharingEnabled
	}
	if p.SharepointSites != nil {
		m["SharepointSites"] = p.SharepointSites
	}
	if p.TeamsSites != nil {
		m["TeamsSites"] = p.TeamsSites
	}
	if p.TenantSettings != "" {
		m["TenantSettings"] = p.TenantSettings
	}
	if p.Users != nil {
		m["Users"] = p.Users
	}
	return m
}

// SetPrivacyManagementRule runs the Set-PrivacyManagementRule cmdlet.
func (s *Service) SetPrivacyManagementRule(ctx context.Context, p SetPrivacyManagementRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-PrivacyManagementRule", p.params())
}

// SetProtectionAlertParams are the parameters of Set-ProtectionAlert.
// DefaultParameterSetName: Identity
type SetProtectionAlertParams struct {
	AggregationType                                             any    `ps:"AggregationType"`
	AlertBy                                                     any    `ps:"AlertBy"`
	AlertFor                                                    any    `ps:"AlertFor"`
	Category                                                    any    `ps:"Category"`
	Comment                                                     string `ps:"Comment"`
	Description                                                 string `ps:"Description"`
	Disabled                                                    bool   `ps:"Disabled"`
	ExternalScenarioData                                        any    `ps:"ExternalScenarioData"`
	Filter                                                      string `ps:"Filter"`
	Identity                                                    any    `ps:"Identity"`
	NotificationCulture                                         any    `ps:"NotificationCulture"`
	NotificationEnabled                                         bool   `ps:"NotificationEnabled"`
	NotifyUser                                                  any    `ps:"NotifyUser"`
	NotifyUserOnFilterMatch                                     bool   `ps:"NotifyUserOnFilterMatch"`
	NotifyUserSuppressionExpiryDate                             any    `ps:"NotifyUserSuppressionExpiryDate"`
	NotifyUserThrottleThreshold                                 any    `ps:"NotifyUserThrottleThreshold"`
	NotifyUserThrottleWindow                                    any    `ps:"NotifyUserThrottleWindow"`
	Operation                                                   any    `ps:"Operation"`
	PrivacyManagementScopedSensitiveInformationTypes            any    `ps:"PrivacyManagementScopedSensitiveInformationTypes"`
	PrivacyManagementScopedSensitiveInformationTypesForCounting any    `ps:"PrivacyManagementScopedSensitiveInformationTypesForCounting"`
	PrivacyManagementScopedSensitiveInformationTypesThreshold   any    `ps:"PrivacyManagementScopedSensitiveInformationTypesThreshold"`
	Severity                                                    any    `ps:"Severity"`
	Threshold                                                   any    `ps:"Threshold"`
	TimeWindow                                                  any    `ps:"TimeWindow"`
	VolumeThreshold                                             any    `ps:"VolumeThreshold"`
}

func (p SetProtectionAlertParams) params() map[string]any {
	m := map[string]any{}
	if p.AggregationType != nil {
		m["AggregationType"] = p.AggregationType
	}
	if p.AlertBy != nil {
		m["AlertBy"] = p.AlertBy
	}
	if p.AlertFor != nil {
		m["AlertFor"] = p.AlertFor
	}
	if p.Category != nil {
		m["Category"] = p.Category
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if p.Disabled {
		m["Disabled"] = true
	}
	if p.ExternalScenarioData != nil {
		m["ExternalScenarioData"] = p.ExternalScenarioData
	}
	if p.Filter != "" {
		m["Filter"] = p.Filter
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.NotificationCulture != nil {
		m["NotificationCulture"] = p.NotificationCulture
	}
	if p.NotificationEnabled {
		m["NotificationEnabled"] = true
	}
	if p.NotifyUser != nil {
		m["NotifyUser"] = p.NotifyUser
	}
	if p.NotifyUserOnFilterMatch {
		m["NotifyUserOnFilterMatch"] = true
	}
	if p.NotifyUserSuppressionExpiryDate != nil {
		m["NotifyUserSuppressionExpiryDate"] = p.NotifyUserSuppressionExpiryDate
	}
	if p.NotifyUserThrottleThreshold != nil {
		m["NotifyUserThrottleThreshold"] = p.NotifyUserThrottleThreshold
	}
	if p.NotifyUserThrottleWindow != nil {
		m["NotifyUserThrottleWindow"] = p.NotifyUserThrottleWindow
	}
	if p.Operation != nil {
		m["Operation"] = p.Operation
	}
	if p.PrivacyManagementScopedSensitiveInformationTypes != nil {
		m["PrivacyManagementScopedSensitiveInformationTypes"] = p.PrivacyManagementScopedSensitiveInformationTypes
	}
	if p.PrivacyManagementScopedSensitiveInformationTypesForCounting != nil {
		m["PrivacyManagementScopedSensitiveInformationTypesForCounting"] = p.PrivacyManagementScopedSensitiveInformationTypesForCounting
	}
	if p.PrivacyManagementScopedSensitiveInformationTypesThreshold != nil {
		m["PrivacyManagementScopedSensitiveInformationTypesThreshold"] = p.PrivacyManagementScopedSensitiveInformationTypesThreshold
	}
	if p.Severity != nil {
		m["Severity"] = p.Severity
	}
	if p.Threshold != nil {
		m["Threshold"] = p.Threshold
	}
	if p.TimeWindow != nil {
		m["TimeWindow"] = p.TimeWindow
	}
	if p.VolumeThreshold != nil {
		m["VolumeThreshold"] = p.VolumeThreshold
	}
	return m
}

// SetProtectionAlert runs the Set-ProtectionAlert cmdlet.
func (s *Service) SetProtectionAlert(ctx context.Context, p SetProtectionAlertParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-ProtectionAlert", p.params())
}

// SetProtectionCompliancePolicyParams are the parameters of Set-ProtectionCompliancePolicy.
// DefaultParameterSetName: Identity
type SetProtectionCompliancePolicyParams struct {
	Enabled   bool   `ps:"Enabled"`
	Identity  any    `ps:"Identity"`
	Locations string `ps:"Locations"`
}

func (p SetProtectionCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Locations != "" {
		m["Locations"] = p.Locations
	}
	return m
}

// SetProtectionCompliancePolicy runs the Set-ProtectionCompliancePolicy cmdlet.
func (s *Service) SetProtectionCompliancePolicy(ctx context.Context, p SetProtectionCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-ProtectionCompliancePolicy", p.params())
}

// SetProtectionComplianceRuleParams are the parameters of Set-ProtectionComplianceRule.
type SetProtectionComplianceRuleParams struct {
	AdvancedRule   string   `ps:"AdvancedRule"`
	ContainsLabels any      `ps:"ContainsLabels"`
	Identity       any      `ps:"Identity"`
	LabelActions   []string `ps:"LabelActions"`
}

func (p SetProtectionComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.AdvancedRule != "" {
		m["AdvancedRule"] = p.AdvancedRule
	}
	if p.ContainsLabels != nil {
		m["ContainsLabels"] = p.ContainsLabels
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if len(p.LabelActions) > 0 {
		m["LabelActions"] = p.LabelActions
	}
	return m
}

// SetProtectionComplianceRule runs the Set-ProtectionComplianceRule cmdlet.
func (s *Service) SetProtectionComplianceRule(ctx context.Context, p SetProtectionComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-ProtectionComplianceRule", p.params())
}

// SetRecordReviewNotificationTemplateConfigParams are the parameters of Set-RecordReviewNotificationTemplateConfig.
type SetRecordReviewNotificationTemplateConfigParams struct {
	CustomizedNotificationDataString string `ps:"CustomizedNotificationDataString"`
	CustomizedReminderDataString     string `ps:"CustomizedReminderDataString"`
	IsCustomizedNotificationTemplate bool   `ps:"IsCustomizedNotificationTemplate"`
	IsCustomizedReminderTemplate     bool   `ps:"IsCustomizedReminderTemplate"`
}

func (p SetRecordReviewNotificationTemplateConfigParams) params() map[string]any {
	m := map[string]any{}
	if p.CustomizedNotificationDataString != "" {
		m["CustomizedNotificationDataString"] = p.CustomizedNotificationDataString
	}
	if p.CustomizedReminderDataString != "" {
		m["CustomizedReminderDataString"] = p.CustomizedReminderDataString
	}
	if p.IsCustomizedNotificationTemplate {
		m["IsCustomizedNotificationTemplate"] = true
	}
	if p.IsCustomizedReminderTemplate {
		m["IsCustomizedReminderTemplate"] = true
	}
	return m
}

// SetRecordReviewNotificationTemplateConfig runs the Set-RecordReviewNotificationTemplateConfig cmdlet.
func (s *Service) SetRecordReviewNotificationTemplateConfig(ctx context.Context, p SetRecordReviewNotificationTemplateConfigParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-RecordReviewNotificationTemplateConfig", p.params())
}

// SetRegulatoryComplianceUIParams are the parameters of Set-RegulatoryComplianceUI.
type SetRegulatoryComplianceUIParams struct {
	Enabled bool `ps:"Enabled"`
}

func (p SetRegulatoryComplianceUIParams) params() map[string]any {
	m := map[string]any{}
	if p.Enabled {
		m["Enabled"] = true
	}
	return m
}

// SetRegulatoryComplianceUI runs the Set-RegulatoryComplianceUI cmdlet.
func (s *Service) SetRegulatoryComplianceUI(ctx context.Context, p SetRegulatoryComplianceUIParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-RegulatoryComplianceUI", p.params())
}

// SetRetentionCompliancePolicyParams are the parameters of Set-RetentionCompliancePolicy.
// DefaultParameterSetName: Identity
type SetRetentionCompliancePolicyParams struct {
	AddAdaptiveScopeLocation            any    `ps:"AddAdaptiveScopeLocation"`
	AddExchangeLocation                 any    `ps:"AddExchangeLocation"`
	AddExchangeLocationException        any    `ps:"AddExchangeLocationException"`
	AddModernGroupLocation              any    `ps:"AddModernGroupLocation"`
	AddModernGroupLocationException     any    `ps:"AddModernGroupLocationException"`
	AddOneDriveLocation                 any    `ps:"AddOneDriveLocation"`
	AddOneDriveLocationException        any    `ps:"AddOneDriveLocationException"`
	AddPublicFolderLocation             any    `ps:"AddPublicFolderLocation"`
	AddSharePointLocation               any    `ps:"AddSharePointLocation"`
	AddSharePointLocationException      any    `ps:"AddSharePointLocationException"`
	AddSkypeLocation                    any    `ps:"AddSkypeLocation"`
	AddSkypeLocationException           any    `ps:"AddSkypeLocationException"`
	AddTeamsChannelLocation             any    `ps:"AddTeamsChannelLocation"`
	AddTeamsChannelLocationException    any    `ps:"AddTeamsChannelLocationException"`
	AddTeamsChatLocation                any    `ps:"AddTeamsChatLocation"`
	AddTeamsChatLocationException       any    `ps:"AddTeamsChatLocationException"`
	Applications                        any    `ps:"Applications"`
	Comment                             string `ps:"Comment"`
	DeletedResources                    string `ps:"DeletedResources"`
	Enabled                             bool   `ps:"Enabled"`
	EnforceSimulationPolicy             bool   `ps:"EnforceSimulationPolicy"`
	ExcludeSPE                          bool   `ps:"ExcludeSPE"`
	Force                               bool   `ps:"Force"`
	ForceBypassScopeCheck               bool   `ps:"ForceBypassScopeCheck"`
	ForceStartSimulation                bool   `ps:"ForceStartSimulation"`
	Identity                            any    `ps:"Identity"`
	IncludeUserOwnedContainers          bool   `ps:"IncludeUserOwnedContainers"`
	PolicyRBACScopes                    any    `ps:"PolicyRBACScopes"`
	PolicyTemplateInfo                  any    `ps:"PolicyTemplateInfo"`
	RemoveAdaptiveScopeLocation         any    `ps:"RemoveAdaptiveScopeLocation"`
	RemoveExchangeLocation              any    `ps:"RemoveExchangeLocation"`
	RemoveExchangeLocationException     any    `ps:"RemoveExchangeLocationException"`
	RemoveModernGroupLocation           any    `ps:"RemoveModernGroupLocation"`
	RemoveModernGroupLocationException  any    `ps:"RemoveModernGroupLocationException"`
	RemoveOneDriveLocation              any    `ps:"RemoveOneDriveLocation"`
	RemoveOneDriveLocationException     any    `ps:"RemoveOneDriveLocationException"`
	RemovePublicFolderLocation          any    `ps:"RemovePublicFolderLocation"`
	RemoveSharePointLocation            any    `ps:"RemoveSharePointLocation"`
	RemoveSharePointLocationException   any    `ps:"RemoveSharePointLocationException"`
	RemoveSkypeLocation                 any    `ps:"RemoveSkypeLocation"`
	RemoveSkypeLocationException        any    `ps:"RemoveSkypeLocationException"`
	RemoveTeamsChannelLocation          any    `ps:"RemoveTeamsChannelLocation"`
	RemoveTeamsChannelLocationException any    `ps:"RemoveTeamsChannelLocationException"`
	RemoveTeamsChatLocation             any    `ps:"RemoveTeamsChatLocation"`
	RemoveTeamsChatLocationException    any    `ps:"RemoveTeamsChatLocationException"`
	RestrictiveRetention                bool   `ps:"RestrictiveRetention"`
	RetryDistribution                   bool   `ps:"RetryDistribution"`
	StartSimulation                     bool   `ps:"StartSimulation"`
}

func (p SetRetentionCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.AddAdaptiveScopeLocation != nil {
		m["AddAdaptiveScopeLocation"] = p.AddAdaptiveScopeLocation
	}
	if p.AddExchangeLocation != nil {
		m["AddExchangeLocation"] = p.AddExchangeLocation
	}
	if p.AddExchangeLocationException != nil {
		m["AddExchangeLocationException"] = p.AddExchangeLocationException
	}
	if p.AddModernGroupLocation != nil {
		m["AddModernGroupLocation"] = p.AddModernGroupLocation
	}
	if p.AddModernGroupLocationException != nil {
		m["AddModernGroupLocationException"] = p.AddModernGroupLocationException
	}
	if p.AddOneDriveLocation != nil {
		m["AddOneDriveLocation"] = p.AddOneDriveLocation
	}
	if p.AddOneDriveLocationException != nil {
		m["AddOneDriveLocationException"] = p.AddOneDriveLocationException
	}
	if p.AddPublicFolderLocation != nil {
		m["AddPublicFolderLocation"] = p.AddPublicFolderLocation
	}
	if p.AddSharePointLocation != nil {
		m["AddSharePointLocation"] = p.AddSharePointLocation
	}
	if p.AddSharePointLocationException != nil {
		m["AddSharePointLocationException"] = p.AddSharePointLocationException
	}
	if p.AddSkypeLocation != nil {
		m["AddSkypeLocation"] = p.AddSkypeLocation
	}
	if p.AddSkypeLocationException != nil {
		m["AddSkypeLocationException"] = p.AddSkypeLocationException
	}
	if p.AddTeamsChannelLocation != nil {
		m["AddTeamsChannelLocation"] = p.AddTeamsChannelLocation
	}
	if p.AddTeamsChannelLocationException != nil {
		m["AddTeamsChannelLocationException"] = p.AddTeamsChannelLocationException
	}
	if p.AddTeamsChatLocation != nil {
		m["AddTeamsChatLocation"] = p.AddTeamsChatLocation
	}
	if p.AddTeamsChatLocationException != nil {
		m["AddTeamsChatLocationException"] = p.AddTeamsChatLocationException
	}
	if p.Applications != nil {
		m["Applications"] = p.Applications
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.DeletedResources != "" {
		m["DeletedResources"] = p.DeletedResources
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.EnforceSimulationPolicy {
		m["EnforceSimulationPolicy"] = true
	}
	if p.ExcludeSPE {
		m["ExcludeSPE"] = true
	}
	if p.Force {
		m["Force"] = true
	}
	if p.ForceBypassScopeCheck {
		m["ForceBypassScopeCheck"] = true
	}
	if p.ForceStartSimulation {
		m["ForceStartSimulation"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IncludeUserOwnedContainers {
		m["IncludeUserOwnedContainers"] = true
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.PolicyTemplateInfo != nil {
		m["PolicyTemplateInfo"] = p.PolicyTemplateInfo
	}
	if p.RemoveAdaptiveScopeLocation != nil {
		m["RemoveAdaptiveScopeLocation"] = p.RemoveAdaptiveScopeLocation
	}
	if p.RemoveExchangeLocation != nil {
		m["RemoveExchangeLocation"] = p.RemoveExchangeLocation
	}
	if p.RemoveExchangeLocationException != nil {
		m["RemoveExchangeLocationException"] = p.RemoveExchangeLocationException
	}
	if p.RemoveModernGroupLocation != nil {
		m["RemoveModernGroupLocation"] = p.RemoveModernGroupLocation
	}
	if p.RemoveModernGroupLocationException != nil {
		m["RemoveModernGroupLocationException"] = p.RemoveModernGroupLocationException
	}
	if p.RemoveOneDriveLocation != nil {
		m["RemoveOneDriveLocation"] = p.RemoveOneDriveLocation
	}
	if p.RemoveOneDriveLocationException != nil {
		m["RemoveOneDriveLocationException"] = p.RemoveOneDriveLocationException
	}
	if p.RemovePublicFolderLocation != nil {
		m["RemovePublicFolderLocation"] = p.RemovePublicFolderLocation
	}
	if p.RemoveSharePointLocation != nil {
		m["RemoveSharePointLocation"] = p.RemoveSharePointLocation
	}
	if p.RemoveSharePointLocationException != nil {
		m["RemoveSharePointLocationException"] = p.RemoveSharePointLocationException
	}
	if p.RemoveSkypeLocation != nil {
		m["RemoveSkypeLocation"] = p.RemoveSkypeLocation
	}
	if p.RemoveSkypeLocationException != nil {
		m["RemoveSkypeLocationException"] = p.RemoveSkypeLocationException
	}
	if p.RemoveTeamsChannelLocation != nil {
		m["RemoveTeamsChannelLocation"] = p.RemoveTeamsChannelLocation
	}
	if p.RemoveTeamsChannelLocationException != nil {
		m["RemoveTeamsChannelLocationException"] = p.RemoveTeamsChannelLocationException
	}
	if p.RemoveTeamsChatLocation != nil {
		m["RemoveTeamsChatLocation"] = p.RemoveTeamsChatLocation
	}
	if p.RemoveTeamsChatLocationException != nil {
		m["RemoveTeamsChatLocationException"] = p.RemoveTeamsChatLocationException
	}
	if p.RestrictiveRetention {
		m["RestrictiveRetention"] = true
	}
	if p.RetryDistribution {
		m["RetryDistribution"] = true
	}
	if p.StartSimulation {
		m["StartSimulation"] = true
	}
	return m
}

// SetRetentionCompliancePolicy runs the Set-RetentionCompliancePolicy cmdlet.
func (s *Service) SetRetentionCompliancePolicy(ctx context.Context, p SetRetentionCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-RetentionCompliancePolicy", p.params())
}

// SetRetentionComplianceRuleParams are the parameters of Set-RetentionComplianceRule.
// DefaultParameterSetName: Identity
type SetRetentionComplianceRuleParams struct {
	ApplyComplianceTag                  string   `ps:"ApplyComplianceTag"`
	ArchiveEnabled                      bool     `ps:"ArchiveEnabled"`
	ArchiveTriggerBasedOn               string   `ps:"ArchiveTriggerBasedOn"`
	ArchiveTriggerInDays                any      `ps:"ArchiveTriggerInDays"`
	Comment                             string   `ps:"Comment"`
	ContentContainsSensitiveInformation []string `ps:"ContentContainsSensitiveInformation"`
	ContentDateFrom                     any      `ps:"ContentDateFrom"`
	ContentDateTo                       any      `ps:"ContentDateTo"`
	ContentMatchQuery                   string   `ps:"ContentMatchQuery"`
	ExcludedItemClasses                 any      `ps:"ExcludedItemClasses"`
	ExpirationDateOption                string   `ps:"ExpirationDateOption"`
	Identity                            any      `ps:"Identity"`
	IRMRiskyUserProfiles                any      `ps:"IRMRiskyUserProfiles"`
	RetentionComplianceAction           string   `ps:"RetentionComplianceAction"`
	RetentionDuration                   any      `ps:"RetentionDuration"`
	RetentionDurationDisplayHint        any      `ps:"RetentionDurationDisplayHint"`
}

func (p SetRetentionComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.ApplyComplianceTag != "" {
		m["ApplyComplianceTag"] = p.ApplyComplianceTag
	}
	if p.ArchiveEnabled {
		m["ArchiveEnabled"] = true
	}
	if p.ArchiveTriggerBasedOn != "" {
		m["ArchiveTriggerBasedOn"] = p.ArchiveTriggerBasedOn
	}
	if p.ArchiveTriggerInDays != nil {
		m["ArchiveTriggerInDays"] = p.ArchiveTriggerInDays
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if len(p.ContentContainsSensitiveInformation) > 0 {
		m["ContentContainsSensitiveInformation"] = p.ContentContainsSensitiveInformation
	}
	if p.ContentDateFrom != nil {
		m["ContentDateFrom"] = p.ContentDateFrom
	}
	if p.ContentDateTo != nil {
		m["ContentDateTo"] = p.ContentDateTo
	}
	if p.ContentMatchQuery != "" {
		m["ContentMatchQuery"] = p.ContentMatchQuery
	}
	if p.ExcludedItemClasses != nil {
		m["ExcludedItemClasses"] = p.ExcludedItemClasses
	}
	if p.ExpirationDateOption != "" {
		m["ExpirationDateOption"] = p.ExpirationDateOption
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.IRMRiskyUserProfiles != nil {
		m["IRMRiskyUserProfiles"] = p.IRMRiskyUserProfiles
	}
	if p.RetentionComplianceAction != "" {
		m["RetentionComplianceAction"] = p.RetentionComplianceAction
	}
	if p.RetentionDuration != nil {
		m["RetentionDuration"] = p.RetentionDuration
	}
	if p.RetentionDurationDisplayHint != nil {
		m["RetentionDurationDisplayHint"] = p.RetentionDurationDisplayHint
	}
	return m
}

// SetRetentionComplianceRule runs the Set-RetentionComplianceRule cmdlet.
func (s *Service) SetRetentionComplianceRule(ctx context.Context, p SetRetentionComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-RetentionComplianceRule", p.params())
}

// SetSensitiveInformationScanParams are the parameters of Set-SensitiveInformationScan.
type SetSensitiveInformationScanParams struct {
	AddEndpointDlpLocation                     any      `ps:"AddEndpointDlpLocation"`
	AddEndpointDlpLocationException            any      `ps:"AddEndpointDlpLocationException"`
	AddExchangeLocation                        any      `ps:"AddExchangeLocation"`
	AddOneDriveLocation                        any      `ps:"AddOneDriveLocation"`
	AddOneDriveLocationException               any      `ps:"AddOneDriveLocationException"`
	AddSharePointLocation                      any      `ps:"AddSharePointLocation"`
	AddSharePointLocationException             any      `ps:"AddSharePointLocationException"`
	AddTeamsLocation                           any      `ps:"AddTeamsLocation"`
	AddTeamsLocationException                  any      `ps:"AddTeamsLocationException"`
	CancelImpactAssessment                     bool     `ps:"CancelImpactAssessment"`
	CancelScan                                 bool     `ps:"CancelScan"`
	Comment                                    string   `ps:"Comment"`
	ExceptIfOneDriveSharedBy                   []string `ps:"ExceptIfOneDriveSharedBy"`
	ExceptIfOneDriveSharedByMemberOf           []string `ps:"ExceptIfOneDriveSharedByMemberOf"`
	ExchangeAdaptiveScopes                     any      `ps:"ExchangeAdaptiveScopes"`
	ExchangeAdaptiveScopesException            any      `ps:"ExchangeAdaptiveScopesException"`
	ExchangeSender                             []string `ps:"ExchangeSender"`
	ExchangeSenderException                    []string `ps:"ExchangeSenderException"`
	ExchangeSenderMemberOf                     []string `ps:"ExchangeSenderMemberOf"`
	ExchangeSenderMemberOfException            []string `ps:"ExchangeSenderMemberOfException"`
	Identity                                   any      `ps:"Identity"`
	Mode                                       any      `ps:"Mode"`
	OneDriveAdaptiveScopes                     any      `ps:"OneDriveAdaptiveScopes"`
	OneDriveAdaptiveScopesException            any      `ps:"OneDriveAdaptiveScopesException"`
	OneDriveSharedBy                           []string `ps:"OneDriveSharedBy"`
	OneDriveSharedByMemberOf                   []string `ps:"OneDriveSharedByMemberOf"`
	PolicyRBACScopes                           any      `ps:"PolicyRBACScopes"`
	RemoveEndpointDlpLocation                  any      `ps:"RemoveEndpointDlpLocation"`
	RemoveEndpointDlpLocationException         any      `ps:"RemoveEndpointDlpLocationException"`
	RemoveExchangeLocation                     any      `ps:"RemoveExchangeLocation"`
	RemoveOneDriveLocation                     any      `ps:"RemoveOneDriveLocation"`
	RemoveOneDriveLocationException            any      `ps:"RemoveOneDriveLocationException"`
	RemoveSharePointLocation                   any      `ps:"RemoveSharePointLocation"`
	RemoveSharePointLocationException          any      `ps:"RemoveSharePointLocationException"`
	RemoveTeamsLocation                        any      `ps:"RemoveTeamsLocation"`
	RemoveTeamsLocationException               any      `ps:"RemoveTeamsLocationException"`
	ScanBudget                                 any      `ps:"ScanBudget"`
	ScanType                                   string   `ps:"ScanType"`
	SharePointAdaptiveScopes                   any      `ps:"SharePointAdaptiveScopes"`
	SharePointAdaptiveScopesException          any      `ps:"SharePointAdaptiveScopesException"`
	StartImpactAssessment                      bool     `ps:"StartImpactAssessment"`
	StopImpactAssessmentAndStartClassification bool     `ps:"StopImpactAssessmentAndStartClassification"`
	TeamsAdaptiveScopes                        any      `ps:"TeamsAdaptiveScopes"`
	TeamsAdaptiveScopesException               any      `ps:"TeamsAdaptiveScopesException"`
}

func (p SetSensitiveInformationScanParams) params() map[string]any {
	m := map[string]any{}
	if p.AddEndpointDlpLocation != nil {
		m["AddEndpointDlpLocation"] = p.AddEndpointDlpLocation
	}
	if p.AddEndpointDlpLocationException != nil {
		m["AddEndpointDlpLocationException"] = p.AddEndpointDlpLocationException
	}
	if p.AddExchangeLocation != nil {
		m["AddExchangeLocation"] = p.AddExchangeLocation
	}
	if p.AddOneDriveLocation != nil {
		m["AddOneDriveLocation"] = p.AddOneDriveLocation
	}
	if p.AddOneDriveLocationException != nil {
		m["AddOneDriveLocationException"] = p.AddOneDriveLocationException
	}
	if p.AddSharePointLocation != nil {
		m["AddSharePointLocation"] = p.AddSharePointLocation
	}
	if p.AddSharePointLocationException != nil {
		m["AddSharePointLocationException"] = p.AddSharePointLocationException
	}
	if p.AddTeamsLocation != nil {
		m["AddTeamsLocation"] = p.AddTeamsLocation
	}
	if p.AddTeamsLocationException != nil {
		m["AddTeamsLocationException"] = p.AddTeamsLocationException
	}
	if p.CancelImpactAssessment {
		m["CancelImpactAssessment"] = true
	}
	if p.CancelScan {
		m["CancelScan"] = true
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if len(p.ExceptIfOneDriveSharedBy) > 0 {
		m["ExceptIfOneDriveSharedBy"] = p.ExceptIfOneDriveSharedBy
	}
	if len(p.ExceptIfOneDriveSharedByMemberOf) > 0 {
		m["ExceptIfOneDriveSharedByMemberOf"] = p.ExceptIfOneDriveSharedByMemberOf
	}
	if p.ExchangeAdaptiveScopes != nil {
		m["ExchangeAdaptiveScopes"] = p.ExchangeAdaptiveScopes
	}
	if p.ExchangeAdaptiveScopesException != nil {
		m["ExchangeAdaptiveScopesException"] = p.ExchangeAdaptiveScopesException
	}
	if len(p.ExchangeSender) > 0 {
		m["ExchangeSender"] = p.ExchangeSender
	}
	if len(p.ExchangeSenderException) > 0 {
		m["ExchangeSenderException"] = p.ExchangeSenderException
	}
	if len(p.ExchangeSenderMemberOf) > 0 {
		m["ExchangeSenderMemberOf"] = p.ExchangeSenderMemberOf
	}
	if len(p.ExchangeSenderMemberOfException) > 0 {
		m["ExchangeSenderMemberOfException"] = p.ExchangeSenderMemberOfException
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Mode != nil {
		m["Mode"] = p.Mode
	}
	if p.OneDriveAdaptiveScopes != nil {
		m["OneDriveAdaptiveScopes"] = p.OneDriveAdaptiveScopes
	}
	if p.OneDriveAdaptiveScopesException != nil {
		m["OneDriveAdaptiveScopesException"] = p.OneDriveAdaptiveScopesException
	}
	if len(p.OneDriveSharedBy) > 0 {
		m["OneDriveSharedBy"] = p.OneDriveSharedBy
	}
	if len(p.OneDriveSharedByMemberOf) > 0 {
		m["OneDriveSharedByMemberOf"] = p.OneDriveSharedByMemberOf
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.RemoveEndpointDlpLocation != nil {
		m["RemoveEndpointDlpLocation"] = p.RemoveEndpointDlpLocation
	}
	if p.RemoveEndpointDlpLocationException != nil {
		m["RemoveEndpointDlpLocationException"] = p.RemoveEndpointDlpLocationException
	}
	if p.RemoveExchangeLocation != nil {
		m["RemoveExchangeLocation"] = p.RemoveExchangeLocation
	}
	if p.RemoveOneDriveLocation != nil {
		m["RemoveOneDriveLocation"] = p.RemoveOneDriveLocation
	}
	if p.RemoveOneDriveLocationException != nil {
		m["RemoveOneDriveLocationException"] = p.RemoveOneDriveLocationException
	}
	if p.RemoveSharePointLocation != nil {
		m["RemoveSharePointLocation"] = p.RemoveSharePointLocation
	}
	if p.RemoveSharePointLocationException != nil {
		m["RemoveSharePointLocationException"] = p.RemoveSharePointLocationException
	}
	if p.RemoveTeamsLocation != nil {
		m["RemoveTeamsLocation"] = p.RemoveTeamsLocation
	}
	if p.RemoveTeamsLocationException != nil {
		m["RemoveTeamsLocationException"] = p.RemoveTeamsLocationException
	}
	if p.ScanBudget != nil {
		m["ScanBudget"] = p.ScanBudget
	}
	if p.ScanType != "" {
		m["ScanType"] = p.ScanType
	}
	if p.SharePointAdaptiveScopes != nil {
		m["SharePointAdaptiveScopes"] = p.SharePointAdaptiveScopes
	}
	if p.SharePointAdaptiveScopesException != nil {
		m["SharePointAdaptiveScopesException"] = p.SharePointAdaptiveScopesException
	}
	if p.StartImpactAssessment {
		m["StartImpactAssessment"] = true
	}
	if p.StopImpactAssessmentAndStartClassification {
		m["StopImpactAssessmentAndStartClassification"] = true
	}
	if p.TeamsAdaptiveScopes != nil {
		m["TeamsAdaptiveScopes"] = p.TeamsAdaptiveScopes
	}
	if p.TeamsAdaptiveScopesException != nil {
		m["TeamsAdaptiveScopesException"] = p.TeamsAdaptiveScopesException
	}
	return m
}

// SetSensitiveInformationScan runs the Set-SensitiveInformationScan cmdlet.
func (s *Service) SetSensitiveInformationScan(ctx context.Context, p SetSensitiveInformationScanParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-SensitiveInformationScan", p.params())
}

// SetSensitiveInformationScanRuleParams are the parameters of Set-SensitiveInformationScanRule.
type SetSensitiveInformationScanRuleParams struct {
	Comment                                        string   `ps:"Comment"`
	ContentContainsSensitiveInformation            []string `ps:"ContentContainsSensitiveInformation"`
	ContentCreatedOrUpdatedDateFrom                any      `ps:"ContentCreatedOrUpdatedDateFrom"`
	ContentCreatedOrUpdatedDateTo                  any      `ps:"ContentCreatedOrUpdatedDateTo"`
	ContentExtensionMatchesWords                   any      `ps:"ContentExtensionMatchesWords"`
	ContentIsNotLabeled                            bool     `ps:"ContentIsNotLabeled"`
	ContentPropertyContainsWords                   any      `ps:"ContentPropertyContainsWords"`
	DocumentSizeOver                               any      `ps:"DocumentSizeOver"`
	ExceptIfContentContainsSensitiveInformation    []string `ps:"ExceptIfContentContainsSensitiveInformation"`
	ExceptIfContentExtensionMatchesWords           any      `ps:"ExceptIfContentExtensionMatchesWords"`
	ExceptIfContentPropertyContainsWords           any      `ps:"ExceptIfContentPropertyContainsWords"`
	ExceptIfDocumentSizeOver                       any      `ps:"ExceptIfDocumentSizeOver"`
	ExceptIfOdcContentContainsSensitiveInformation []string `ps:"ExceptIfOdcContentContainsSensitiveInformation"`
	Identity                                       any      `ps:"Identity"`
	OdcContentContainsSensitiveInformation         []string `ps:"OdcContentContainsSensitiveInformation"`
	StartImpactAssessment                          bool     `ps:"StartImpactAssessment"`
	Workload                                       any      `ps:"Workload"`
}

func (p SetSensitiveInformationScanRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if len(p.ContentContainsSensitiveInformation) > 0 {
		m["ContentContainsSensitiveInformation"] = p.ContentContainsSensitiveInformation
	}
	if p.ContentCreatedOrUpdatedDateFrom != nil {
		m["ContentCreatedOrUpdatedDateFrom"] = p.ContentCreatedOrUpdatedDateFrom
	}
	if p.ContentCreatedOrUpdatedDateTo != nil {
		m["ContentCreatedOrUpdatedDateTo"] = p.ContentCreatedOrUpdatedDateTo
	}
	if p.ContentExtensionMatchesWords != nil {
		m["ContentExtensionMatchesWords"] = p.ContentExtensionMatchesWords
	}
	if p.ContentIsNotLabeled {
		m["ContentIsNotLabeled"] = true
	}
	if p.ContentPropertyContainsWords != nil {
		m["ContentPropertyContainsWords"] = p.ContentPropertyContainsWords
	}
	if p.DocumentSizeOver != nil {
		m["DocumentSizeOver"] = p.DocumentSizeOver
	}
	if len(p.ExceptIfContentContainsSensitiveInformation) > 0 {
		m["ExceptIfContentContainsSensitiveInformation"] = p.ExceptIfContentContainsSensitiveInformation
	}
	if p.ExceptIfContentExtensionMatchesWords != nil {
		m["ExceptIfContentExtensionMatchesWords"] = p.ExceptIfContentExtensionMatchesWords
	}
	if p.ExceptIfContentPropertyContainsWords != nil {
		m["ExceptIfContentPropertyContainsWords"] = p.ExceptIfContentPropertyContainsWords
	}
	if p.ExceptIfDocumentSizeOver != nil {
		m["ExceptIfDocumentSizeOver"] = p.ExceptIfDocumentSizeOver
	}
	if len(p.ExceptIfOdcContentContainsSensitiveInformation) > 0 {
		m["ExceptIfOdcContentContainsSensitiveInformation"] = p.ExceptIfOdcContentContainsSensitiveInformation
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if len(p.OdcContentContainsSensitiveInformation) > 0 {
		m["OdcContentContainsSensitiveInformation"] = p.OdcContentContainsSensitiveInformation
	}
	if p.StartImpactAssessment {
		m["StartImpactAssessment"] = true
	}
	if p.Workload != nil {
		m["Workload"] = p.Workload
	}
	return m
}

// SetSensitiveInformationScanRule runs the Set-SensitiveInformationScanRule cmdlet.
func (s *Service) SetSensitiveInformationScanRule(ctx context.Context, p SetSensitiveInformationScanRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-SensitiveInformationScanRule", p.params())
}

// SetServiceDomainGroupParams are the parameters of Set-ServiceDomainGroup.
type SetServiceDomainGroupParams struct {
	Identity      any    `ps:"Identity"`
	SiteAddresses string `ps:"SiteAddresses"`
}

func (p SetServiceDomainGroupParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.SiteAddresses != "" {
		m["SiteAddresses"] = p.SiteAddresses
	}
	return m
}

// SetServiceDomainGroup runs the Set-ServiceDomainGroup cmdlet.
func (s *Service) SetServiceDomainGroup(ctx context.Context, p SetServiceDomainGroupParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-ServiceDomainGroup", p.params())
}

// SetServicePrincipalParams are the parameters of Set-ServicePrincipal.
// DefaultParameterSetName: Identity
type SetServicePrincipalParams struct {
	DisplayName string `ps:"DisplayName"`
	Identity    any    `ps:"Identity"`
}

func (p SetServicePrincipalParams) params() map[string]any {
	m := map[string]any{}
	if p.DisplayName != "" {
		m["DisplayName"] = p.DisplayName
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// SetServicePrincipal runs the Set-ServicePrincipal cmdlet.
func (s *Service) SetServicePrincipal(ctx context.Context, p SetServicePrincipalParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-ServicePrincipal", p.params())
}

// SetSupervisoryReviewPolicyV2Params are the parameters of Set-SupervisoryReviewPolicyV2.
// DefaultParameterSetName: Identity
type SetSupervisoryReviewPolicyV2Params struct {
	AddReviewers             []string `ps:"AddReviewers"`
	Comment                  string   `ps:"Comment"`
	Enabled                  bool     `ps:"Enabled"`
	Force                    bool     `ps:"Force"`
	Identity                 any      `ps:"Identity"`
	PolicyRBACScopes         any      `ps:"PolicyRBACScopes"`
	PolicyTemplate           string   `ps:"PolicyTemplate"`
	PreservationPeriodInDays int      `ps:"PreservationPeriodInDays"`
	RemoveReviewers          []string `ps:"RemoveReviewers"`
	RetentionPeriodInDays    int      `ps:"RetentionPeriodInDays"`
	Reviewers                []string `ps:"Reviewers"`
	UserReportingWorkloads   []string `ps:"UserReportingWorkloads"`
}

func (p SetSupervisoryReviewPolicyV2Params) params() map[string]any {
	m := map[string]any{}
	if len(p.AddReviewers) > 0 {
		m["AddReviewers"] = p.AddReviewers
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.PolicyTemplate != "" {
		m["PolicyTemplate"] = p.PolicyTemplate
	}
	if p.PreservationPeriodInDays != 0 {
		m["PreservationPeriodInDays"] = p.PreservationPeriodInDays
	}
	if len(p.RemoveReviewers) > 0 {
		m["RemoveReviewers"] = p.RemoveReviewers
	}
	if p.RetentionPeriodInDays != 0 {
		m["RetentionPeriodInDays"] = p.RetentionPeriodInDays
	}
	if len(p.Reviewers) > 0 {
		m["Reviewers"] = p.Reviewers
	}
	if len(p.UserReportingWorkloads) > 0 {
		m["UserReportingWorkloads"] = p.UserReportingWorkloads
	}
	return m
}

// SetSupervisoryReviewPolicyV2 runs the Set-SupervisoryReviewPolicyV2 cmdlet.
func (s *Service) SetSupervisoryReviewPolicyV2(ctx context.Context, p SetSupervisoryReviewPolicyV2Params) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-SupervisoryReviewPolicyV2", p.params())
}

// SetSupervisoryReviewRuleParams are the parameters of Set-SupervisoryReviewRule.
// DefaultParameterSetName: Identity
type SetSupervisoryReviewRuleParams struct {
	AdvancedRule                        string   `ps:"AdvancedRule"`
	CcsiDataModelOperator               string   `ps:"CcsiDataModelOperator"`
	Condition                           string   `ps:"Condition"`
	ContentContainsSensitiveInformation []string `ps:"ContentContainsSensitiveInformation"`
	ContentMatchesDataModel             string   `ps:"ContentMatchesDataModel"`
	ContentSources                      []string `ps:"ContentSources"`
	DayXInsights                        bool     `ps:"DayXInsights"`
	ExceptIfFrom                        any      `ps:"ExceptIfFrom"`
	ExceptIfRecipientDomainIs           any      `ps:"ExceptIfRecipientDomainIs"`
	ExceptIfRevieweeIs                  any      `ps:"ExceptIfRevieweeIs"`
	ExceptIfSenderDomainIs              any      `ps:"ExceptIfSenderDomainIs"`
	ExceptIfSentTo                      any      `ps:"ExceptIfSentTo"`
	ExceptIfSubjectOrBodyContainsWords  any      `ps:"ExceptIfSubjectOrBodyContainsWords"`
	From                                any      `ps:"From"`
	Identity                            any      `ps:"Identity"`
	IncludeAdaptiveScopes               []string `ps:"IncludeAdaptiveScopes"`
	InPurviewFilter                     string   `ps:"InPurviewFilter"`
	Ocr                                 bool     `ps:"Ocr"`
	PolicyRBACScopes                    any      `ps:"PolicyRBACScopes"`
	RateOfSampling                      string   `ps:"RateOfSampling"`
	SamplingRate                        int      `ps:"SamplingRate"`
	SentTo                              any      `ps:"SentTo"`
}

func (p SetSupervisoryReviewRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.AdvancedRule != "" {
		m["AdvancedRule"] = p.AdvancedRule
	}
	if p.CcsiDataModelOperator != "" {
		m["CcsiDataModelOperator"] = p.CcsiDataModelOperator
	}
	if p.Condition != "" {
		m["Condition"] = p.Condition
	}
	if len(p.ContentContainsSensitiveInformation) > 0 {
		m["ContentContainsSensitiveInformation"] = p.ContentContainsSensitiveInformation
	}
	if p.ContentMatchesDataModel != "" {
		m["ContentMatchesDataModel"] = p.ContentMatchesDataModel
	}
	if len(p.ContentSources) > 0 {
		m["ContentSources"] = p.ContentSources
	}
	if p.DayXInsights {
		m["DayXInsights"] = true
	}
	if p.ExceptIfFrom != nil {
		m["ExceptIfFrom"] = p.ExceptIfFrom
	}
	if p.ExceptIfRecipientDomainIs != nil {
		m["ExceptIfRecipientDomainIs"] = p.ExceptIfRecipientDomainIs
	}
	if p.ExceptIfRevieweeIs != nil {
		m["ExceptIfRevieweeIs"] = p.ExceptIfRevieweeIs
	}
	if p.ExceptIfSenderDomainIs != nil {
		m["ExceptIfSenderDomainIs"] = p.ExceptIfSenderDomainIs
	}
	if p.ExceptIfSentTo != nil {
		m["ExceptIfSentTo"] = p.ExceptIfSentTo
	}
	if p.ExceptIfSubjectOrBodyContainsWords != nil {
		m["ExceptIfSubjectOrBodyContainsWords"] = p.ExceptIfSubjectOrBodyContainsWords
	}
	if p.From != nil {
		m["From"] = p.From
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if len(p.IncludeAdaptiveScopes) > 0 {
		m["IncludeAdaptiveScopes"] = p.IncludeAdaptiveScopes
	}
	if p.InPurviewFilter != "" {
		m["InPurviewFilter"] = p.InPurviewFilter
	}
	if p.Ocr {
		m["Ocr"] = true
	}
	if p.PolicyRBACScopes != nil {
		m["PolicyRBACScopes"] = p.PolicyRBACScopes
	}
	if p.RateOfSampling != "" {
		m["RateOfSampling"] = p.RateOfSampling
	}
	if p.SamplingRate != 0 {
		m["SamplingRate"] = p.SamplingRate
	}
	if p.SentTo != nil {
		m["SentTo"] = p.SentTo
	}
	return m
}

// SetSupervisoryReviewRule runs the Set-SupervisoryReviewRule cmdlet.
func (s *Service) SetSupervisoryReviewRule(ctx context.Context, p SetSupervisoryReviewRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-SupervisoryReviewRule", p.params())
}

// SetTeamsRetentionCompliancePolicyParams are the parameters of Set-TeamsRetentionCompliancePolicy.
// DefaultParameterSetName: Identity
type SetTeamsRetentionCompliancePolicyParams struct {
	AddTeamsChannelLocation             any    `ps:"AddTeamsChannelLocation"`
	AddTeamsChannelLocationException    any    `ps:"AddTeamsChannelLocationException"`
	AddTeamsChatLocation                any    `ps:"AddTeamsChatLocation"`
	AddTeamsChatLocationException       any    `ps:"AddTeamsChatLocationException"`
	Comment                             string `ps:"Comment"`
	Enabled                             bool   `ps:"Enabled"`
	Force                               bool   `ps:"Force"`
	Identity                            any    `ps:"Identity"`
	RemoveTeamsChannelLocation          any    `ps:"RemoveTeamsChannelLocation"`
	RemoveTeamsChannelLocationException any    `ps:"RemoveTeamsChannelLocationException"`
	RemoveTeamsChatLocation             any    `ps:"RemoveTeamsChatLocation"`
	RemoveTeamsChatLocationException    any    `ps:"RemoveTeamsChatLocationException"`
	RestrictiveRetention                bool   `ps:"RestrictiveRetention"`
	RetryDistribution                   bool   `ps:"RetryDistribution"`
}

func (p SetTeamsRetentionCompliancePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.AddTeamsChannelLocation != nil {
		m["AddTeamsChannelLocation"] = p.AddTeamsChannelLocation
	}
	if p.AddTeamsChannelLocationException != nil {
		m["AddTeamsChannelLocationException"] = p.AddTeamsChannelLocationException
	}
	if p.AddTeamsChatLocation != nil {
		m["AddTeamsChatLocation"] = p.AddTeamsChatLocation
	}
	if p.AddTeamsChatLocationException != nil {
		m["AddTeamsChatLocationException"] = p.AddTeamsChatLocationException
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Enabled {
		m["Enabled"] = true
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.RemoveTeamsChannelLocation != nil {
		m["RemoveTeamsChannelLocation"] = p.RemoveTeamsChannelLocation
	}
	if p.RemoveTeamsChannelLocationException != nil {
		m["RemoveTeamsChannelLocationException"] = p.RemoveTeamsChannelLocationException
	}
	if p.RemoveTeamsChatLocation != nil {
		m["RemoveTeamsChatLocation"] = p.RemoveTeamsChatLocation
	}
	if p.RemoveTeamsChatLocationException != nil {
		m["RemoveTeamsChatLocationException"] = p.RemoveTeamsChatLocationException
	}
	if p.RestrictiveRetention {
		m["RestrictiveRetention"] = true
	}
	if p.RetryDistribution {
		m["RetryDistribution"] = true
	}
	return m
}

// SetTeamsRetentionCompliancePolicy runs the Set-TeamsRetentionCompliancePolicy cmdlet.
func (s *Service) SetTeamsRetentionCompliancePolicy(ctx context.Context, p SetTeamsRetentionCompliancePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-TeamsRetentionCompliancePolicy", p.params())
}

// SetTeamsRetentionComplianceRuleParams are the parameters of Set-TeamsRetentionComplianceRule.
// DefaultParameterSetName: Identity
type SetTeamsRetentionComplianceRuleParams struct {
	Comment                   string `ps:"Comment"`
	Identity                  any    `ps:"Identity"`
	RetentionComplianceAction string `ps:"RetentionComplianceAction"`
	RetentionDuration         any    `ps:"RetentionDuration"`
}

func (p SetTeamsRetentionComplianceRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.RetentionComplianceAction != "" {
		m["RetentionComplianceAction"] = p.RetentionComplianceAction
	}
	if p.RetentionDuration != nil {
		m["RetentionDuration"] = p.RetentionDuration
	}
	return m
}

// SetTeamsRetentionComplianceRule runs the Set-TeamsRetentionComplianceRule cmdlet.
func (s *Service) SetTeamsRetentionComplianceRule(ctx context.Context, p SetTeamsRetentionComplianceRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-TeamsRetentionComplianceRule", p.params())
}

// SetThreatResponsePolicyParams are the parameters of Set-ThreatResponsePolicy.
// DefaultParameterSetName: Identity
type SetThreatResponsePolicyParams struct {
	Comment           string `ps:"Comment"`
	Force             bool   `ps:"Force"`
	Identity          any    `ps:"Identity"`
	Mode              any    `ps:"Mode"`
	NewName           string `ps:"NewName"`
	RetryDistribution bool   `ps:"RetryDistribution"`
}

func (p SetThreatResponsePolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.Force {
		m["Force"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Mode != nil {
		m["Mode"] = p.Mode
	}
	if p.NewName != "" {
		m["NewName"] = p.NewName
	}
	if p.RetryDistribution {
		m["RetryDistribution"] = true
	}
	return m
}

// SetThreatResponsePolicy runs the Set-ThreatResponsePolicy cmdlet.
func (s *Service) SetThreatResponsePolicy(ctx context.Context, p SetThreatResponsePolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-ThreatResponsePolicy", p.params())
}

// SetThreatResponseRuleParams are the parameters of Set-ThreatResponseRule.
// DefaultParameterSetName: Identity
type SetThreatResponseRuleParams struct {
	ActivationDate                        any    `ps:"ActivationDate"`
	Comment                               string `ps:"Comment"`
	DisableActiveSync                     bool   `ps:"DisableActiveSync"`
	DisableAutoForwardDefaultRemoteDomain bool   `ps:"DisableAutoForwardDefaultRemoteDomain"`
	DisableBasicAuth                      bool   `ps:"DisableBasicAuth"`
	Disabled                              bool   `ps:"Disabled"`
	DisableEWS                            bool   `ps:"DisableEWS"`
	DisableForwardingSMTPAddress          bool   `ps:"DisableForwardingSMTPAddress"`
	DisableIMAP                           bool   `ps:"DisableIMAP"`
	DisableMAPI                           bool   `ps:"DisableMAPI"`
	DisableMAPIOverHTTP                   bool   `ps:"DisableMAPIOverHTTP"`
	DisableOWA                            bool   `ps:"DisableOWA"`
	DisablePOP                            bool   `ps:"DisablePOP"`
	DisableTransportRuleForwarding        bool   `ps:"DisableTransportRuleForwarding"`
	ExceptIfHighRiskUser                  bool   `ps:"ExceptIfHighRiskUser"`
	ExceptIfMemberOfDL                    any    `ps:"ExceptIfMemberOfDL"`
	ExpiryDate                            any    `ps:"ExpiryDate"`
	HighRiskUser                          bool   `ps:"HighRiskUser"`
	Identity                              any    `ps:"Identity"`
	MemberOfDL                            any    `ps:"MemberOfDL"`
}

func (p SetThreatResponseRuleParams) params() map[string]any {
	m := map[string]any{}
	if p.ActivationDate != nil {
		m["ActivationDate"] = p.ActivationDate
	}
	if p.Comment != "" {
		m["Comment"] = p.Comment
	}
	if p.DisableActiveSync {
		m["DisableActiveSync"] = true
	}
	if p.DisableAutoForwardDefaultRemoteDomain {
		m["DisableAutoForwardDefaultRemoteDomain"] = true
	}
	if p.DisableBasicAuth {
		m["DisableBasicAuth"] = true
	}
	if p.Disabled {
		m["Disabled"] = true
	}
	if p.DisableEWS {
		m["DisableEWS"] = true
	}
	if p.DisableForwardingSMTPAddress {
		m["DisableForwardingSMTPAddress"] = true
	}
	if p.DisableIMAP {
		m["DisableIMAP"] = true
	}
	if p.DisableMAPI {
		m["DisableMAPI"] = true
	}
	if p.DisableMAPIOverHTTP {
		m["DisableMAPIOverHTTP"] = true
	}
	if p.DisableOWA {
		m["DisableOWA"] = true
	}
	if p.DisablePOP {
		m["DisablePOP"] = true
	}
	if p.DisableTransportRuleForwarding {
		m["DisableTransportRuleForwarding"] = true
	}
	if p.ExceptIfHighRiskUser {
		m["ExceptIfHighRiskUser"] = true
	}
	if p.ExceptIfMemberOfDL != nil {
		m["ExceptIfMemberOfDL"] = p.ExceptIfMemberOfDL
	}
	if p.ExpiryDate != nil {
		m["ExpiryDate"] = p.ExpiryDate
	}
	if p.HighRiskUser {
		m["HighRiskUser"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.MemberOfDL != nil {
		m["MemberOfDL"] = p.MemberOfDL
	}
	return m
}

// SetThreatResponseRule runs the Set-ThreatResponseRule cmdlet.
func (s *Service) SetThreatResponseRule(ctx context.Context, p SetThreatResponseRuleParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-ThreatResponseRule", p.params())
}

// SetUnifiedAuditLogRetentionPolicyParams are the parameters of Set-UnifiedAuditLogRetentionPolicy.
// DefaultParameterSetName: Identity
type SetUnifiedAuditLogRetentionPolicyParams struct {
	Description       string `ps:"Description"`
	Identity          any    `ps:"Identity"`
	Operations        any    `ps:"Operations"`
	Priority          int    `ps:"Priority"`
	RecordTypes       any    `ps:"RecordTypes"`
	RetentionDuration any    `ps:"RetentionDuration"`
	UserIds           any    `ps:"UserIds"`
}

func (p SetUnifiedAuditLogRetentionPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Description != "" {
		m["Description"] = p.Description
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Operations != nil {
		m["Operations"] = p.Operations
	}
	if p.Priority != 0 {
		m["Priority"] = p.Priority
	}
	if p.RecordTypes != nil {
		m["RecordTypes"] = p.RecordTypes
	}
	if p.RetentionDuration != nil {
		m["RetentionDuration"] = p.RetentionDuration
	}
	if p.UserIds != nil {
		m["UserIds"] = p.UserIds
	}
	return m
}

// SetUnifiedAuditLogRetentionPolicy runs the Set-UnifiedAuditLogRetentionPolicy cmdlet.
func (s *Service) SetUnifiedAuditLogRetentionPolicy(ctx context.Context, p SetUnifiedAuditLogRetentionPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Set-UnifiedAuditLogRetentionPolicy", p.params())
}

// StartComplianceSearchParams are the parameters of Start-ComplianceSearch.
// DefaultParameterSetName: Identity
type StartComplianceSearchParams struct {
	Force        bool `ps:"Force"`
	Identity     any  `ps:"Identity"`
	RetryOnError bool `ps:"RetryOnError"`
	UseBigFunnel bool `ps:"UseBigFunnel"`
}

func (p StartComplianceSearchParams) params() map[string]any {
	m := map[string]any{}
	if p.Force {
		m["Force"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.RetryOnError {
		m["RetryOnError"] = true
	}
	if p.UseBigFunnel {
		m["UseBigFunnel"] = true
	}
	return m
}

// StartComplianceSearch runs the Start-ComplianceSearch cmdlet.
func (s *Service) StartComplianceSearch(ctx context.Context, p StartComplianceSearchParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Start-ComplianceSearch", p.params())
}

// StartDlpSensitiveInformationScanParams are the parameters of Start-DlpSensitiveInformationScan.
// DefaultParameterSetName: Identity
type StartDlpSensitiveInformationScanParams struct {
	SensitiveInformationScanTimeWindowExo any `ps:"SensitiveInformationScanTimeWindowExo"`
}

func (p StartDlpSensitiveInformationScanParams) params() map[string]any {
	m := map[string]any{}
	if p.SensitiveInformationScanTimeWindowExo != nil {
		m["SensitiveInformationScanTimeWindowExo"] = p.SensitiveInformationScanTimeWindowExo
	}
	return m
}

// StartDlpSensitiveInformationScan runs the Start-DlpSensitiveInformationScan cmdlet.
func (s *Service) StartDlpSensitiveInformationScan(ctx context.Context, p StartDlpSensitiveInformationScanParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Start-DlpSensitiveInformationScan", p.params())
}

// StartInformationBarrierPoliciesApplicationParams are the parameters of Start-InformationBarrierPoliciesApplication.
// DefaultParameterSetName: Identity
type StartInformationBarrierPoliciesApplicationParams struct {
	CleanupGroupSegmentLink bool `ps:"CleanupGroupSegmentLink"`
	Identity                any  `ps:"Identity"`
	Rollback                bool `ps:"Rollback"`
	Upgrade                 bool `ps:"Upgrade"`
}

func (p StartInformationBarrierPoliciesApplicationParams) params() map[string]any {
	m := map[string]any{}
	if p.CleanupGroupSegmentLink {
		m["CleanupGroupSegmentLink"] = true
	}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if p.Rollback {
		m["Rollback"] = true
	}
	if p.Upgrade {
		m["Upgrade"] = true
	}
	return m
}

// StartInformationBarrierPoliciesApplication runs the Start-InformationBarrierPoliciesApplication cmdlet.
func (s *Service) StartInformationBarrierPoliciesApplication(ctx context.Context, p StartInformationBarrierPoliciesApplicationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Start-InformationBarrierPoliciesApplication", p.params())
}

// StopComplianceSearchParams are the parameters of Stop-ComplianceSearch.
// DefaultParameterSetName: Identity
type StopComplianceSearchParams struct {
	Identity any `ps:"Identity"`
}

func (p StopComplianceSearchParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// StopComplianceSearch runs the Stop-ComplianceSearch cmdlet.
func (s *Service) StopComplianceSearch(ctx context.Context, p StopComplianceSearchParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Stop-ComplianceSearch", p.params())
}

// StopInformationBarrierPoliciesApplicationParams are the parameters of Stop-InformationBarrierPoliciesApplication.
// DefaultParameterSetName: Identity
type StopInformationBarrierPoliciesApplicationParams struct {
	Identity any `ps:"Identity"`
}

func (p StopInformationBarrierPoliciesApplicationParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// StopInformationBarrierPoliciesApplication runs the Stop-InformationBarrierPoliciesApplication cmdlet.
func (s *Service) StopInformationBarrierPoliciesApplication(ctx context.Context, p StopInformationBarrierPoliciesApplicationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Stop-InformationBarrierPoliciesApplication", p.params())
}

// SyncLabelParams are the parameters of Sync-Label.
// DefaultParameterSetName: Identity
type SyncLabelParams struct {
	Identity      any      `ps:"Identity"`
	MemberTenants []string `ps:"MemberTenants"`
}

func (p SyncLabelParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	if len(p.MemberTenants) > 0 {
		m["MemberTenants"] = p.MemberTenants
	}
	return m
}

// SyncLabel runs the Sync-Label cmdlet.
func (s *Service) SyncLabel(ctx context.Context, p SyncLabelParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Sync-Label", p.params())
}

// TestDataClassificationParams are the parameters of Test-DataClassification.
type TestDataClassificationParams struct {
	ClassificationNames       []string `ps:"ClassificationNames"`
	DiagnosticPlatform        any      `ps:"DiagnosticPlatform"`
	DiagnosticPurpose         any      `ps:"DiagnosticPurpose"`
	DomainController          any      `ps:"DomainController"`
	FileExtension             string   `ps:"FileExtension"`
	TestTextExtractionResults []string `ps:"TestTextExtractionResults"`
	TextToClassify            string   `ps:"TextToClassify"`
}

func (p TestDataClassificationParams) params() map[string]any {
	m := map[string]any{}
	if len(p.ClassificationNames) > 0 {
		m["ClassificationNames"] = p.ClassificationNames
	}
	if p.DiagnosticPlatform != nil {
		m["DiagnosticPlatform"] = p.DiagnosticPlatform
	}
	if p.DiagnosticPurpose != nil {
		m["DiagnosticPurpose"] = p.DiagnosticPurpose
	}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.FileExtension != "" {
		m["FileExtension"] = p.FileExtension
	}
	if len(p.TestTextExtractionResults) > 0 {
		m["TestTextExtractionResults"] = p.TestTextExtractionResults
	}
	if p.TextToClassify != "" {
		m["TextToClassify"] = p.TextToClassify
	}
	return m
}

// TestDataClassification runs the Test-DataClassification cmdlet.
func (s *Service) TestDataClassification(ctx context.Context, p TestDataClassificationParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Test-DataClassification", p.params())
}

// TestInformationBarrierPolicyParams are the parameters of Test-InformationBarrierPolicy.
// DefaultParameterSetName: Identity
type TestInformationBarrierPolicyParams struct {
	Identity any `ps:"Identity"`
}

func (p TestInformationBarrierPolicyParams) params() map[string]any {
	m := map[string]any{}
	if p.Identity != nil {
		m["Identity"] = p.Identity
	}
	return m
}

// TestInformationBarrierPolicy runs the Test-InformationBarrierPolicy cmdlet.
func (s *Service) TestInformationBarrierPolicy(ctx context.Context, p TestInformationBarrierPolicyParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Test-InformationBarrierPolicy", p.params())
}

// TestLabelConfigParams are the parameters of Test-LabelConfig.
type TestLabelConfigParams struct {
	DomainController  any `ps:"DomainController"`
	Scenario          any `ps:"Scenario"`
	UserPrincipalName any `ps:"UserPrincipalName"`
}

func (p TestLabelConfigParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if p.Scenario != nil {
		m["Scenario"] = p.Scenario
	}
	if p.UserPrincipalName != nil {
		m["UserPrincipalName"] = p.UserPrincipalName
	}
	return m
}

// TestLabelConfig runs the Test-LabelConfig cmdlet.
func (s *Service) TestLabelConfig(ctx context.Context, p TestLabelConfigParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Test-LabelConfig", p.params())
}

// TestTextExtractionParams are the parameters of Test-TextExtraction.
type TestTextExtractionParams struct {
	DomainController any      `ps:"DomainController"`
	FileData         []string `ps:"FileData"`
}

func (p TestTextExtractionParams) params() map[string]any {
	m := map[string]any{}
	if p.DomainController != nil {
		m["DomainController"] = p.DomainController
	}
	if len(p.FileData) > 0 {
		m["FileData"] = p.FileData
	}
	return m
}

// TestTextExtraction runs the Test-TextExtraction cmdlet.
func (s *Service) TestTextExtraction(ctx context.Context, p TestTextExtractionParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Test-TextExtraction", p.params())
}

// UpdateComplianceCaseMemberParams are the parameters of Update-ComplianceCaseMember.
type UpdateComplianceCaseMemberParams struct {
	Case    string   `ps:"Case"`
	Members []string `ps:"Members"`
}

func (p UpdateComplianceCaseMemberParams) params() map[string]any {
	m := map[string]any{}
	if p.Case != "" {
		m["Case"] = p.Case
	}
	if len(p.Members) > 0 {
		m["Members"] = p.Members
	}
	return m
}

// UpdateComplianceCaseMember runs the Update-ComplianceCaseMember cmdlet.
func (s *Service) UpdateComplianceCaseMember(ctx context.Context, p UpdateComplianceCaseMemberParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Update-ComplianceCaseMember", p.params())
}

// ValidateRetentionRuleQueryParams are the parameters of Validate-RetentionRuleQuery.
type ValidateRetentionRuleQueryParams struct {
	KqlQueryString string `ps:"KqlQueryString"`
}

func (p ValidateRetentionRuleQueryParams) params() map[string]any {
	m := map[string]any{}
	if p.KqlQueryString != "" {
		m["KqlQueryString"] = p.KqlQueryString
	}
	return m
}

// ValidateRetentionRuleQuery runs the Validate-RetentionRuleQuery cmdlet.
func (s *Service) ValidateRetentionRuleQuery(ctx context.Context, p ValidateRetentionRuleQueryParams) (*adminapi.Result, error) {
	return s.C.Invoke(ctx, "Validate-RetentionRuleQuery", p.params())
}
