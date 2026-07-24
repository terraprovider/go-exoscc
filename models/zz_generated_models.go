// Code generated from EXO-metadata.xml by gen-models. DO NOT EDIT.

package models

import (
	"encoding/json"
	"time"
)

var _ = time.Time{}
var _ json.RawMessage

// BulkMovesEnabled enum.
type BulkMovesEnabled string

const (
	BulkMovesEnabled_NotSet BulkMovesEnabled = "NotSet"
	BulkMovesEnabled_On     BulkMovesEnabled = "On"
	BulkMovesEnabled_Off    BulkMovesEnabled = "Off"
)

// CalendarProcessingFlags enum.
type CalendarProcessingFlags string

const (
	CalendarProcessingFlags_None       CalendarProcessingFlags = "None"
	CalendarProcessingFlags_AutoUpdate CalendarProcessingFlags = "AutoUpdate"
	CalendarProcessingFlags_AutoAccept CalendarProcessingFlags = "AutoAccept"
)

// ElcFolderType enum.
type ElcFolderType string

const (
	ElcFolderType_Calendar              ElcFolderType = "Calendar"
	ElcFolderType_Contacts              ElcFolderType = "Contacts"
	ElcFolderType_DeletedItems          ElcFolderType = "DeletedItems"
	ElcFolderType_Drafts                ElcFolderType = "Drafts"
	ElcFolderType_Inbox                 ElcFolderType = "Inbox"
	ElcFolderType_JunkEmail             ElcFolderType = "JunkEmail"
	ElcFolderType_Journal               ElcFolderType = "Journal"
	ElcFolderType_Notes                 ElcFolderType = "Notes"
	ElcFolderType_Outbox                ElcFolderType = "Outbox"
	ElcFolderType_SentItems             ElcFolderType = "SentItems"
	ElcFolderType_Tasks                 ElcFolderType = "Tasks"
	ElcFolderType_All                   ElcFolderType = "All"
	ElcFolderType_ManagedCustomFolder   ElcFolderType = "ManagedCustomFolder"
	ElcFolderType_RssSubscriptions      ElcFolderType = "RssSubscriptions"
	ElcFolderType_SyncIssues            ElcFolderType = "SyncIssues"
	ElcFolderType_ConversationHistory   ElcFolderType = "ConversationHistory"
	ElcFolderType_Personal              ElcFolderType = "Personal"
	ElcFolderType_RecoverableItems      ElcFolderType = "RecoverableItems"
	ElcFolderType_NonIpmRoot            ElcFolderType = "NonIpmRoot"
	ElcFolderType_LegacyArchiveJournals ElcFolderType = "LegacyArchiveJournals"
	ElcFolderType_Clutter               ElcFolderType = "Clutter"
	ElcFolderType_Archive               ElcFolderType = "Archive"
)

// EncryptedFileTypes enum.
type EncryptedFileTypes string

const (
	EncryptedFileTypes_Acrobat EncryptedFileTypes = "Acrobat"
	EncryptedFileTypes_Archive EncryptedFileTypes = "Archive"
	EncryptedFileTypes_Office  EncryptedFileTypes = "Office"
	EncryptedFileTypes_Others  EncryptedFileTypes = "Others"
)

// ExternalAudience enum.
type ExternalAudience string

const (
	ExternalAudience_None  ExternalAudience = "None"
	ExternalAudience_Known ExternalAudience = "Known"
	ExternalAudience_All   ExternalAudience = "All"
)

// FileTypeFilteringAction enum.
type FileTypeFilteringAction string

const (
	FileTypeFilteringAction_Quarantine FileTypeFilteringAction = "Quarantine"
	FileTypeFilteringAction_Reject     FileTypeFilteringAction = "Reject"
)

// ImpersonationAction enum.
type ImpersonationAction string

const (
	ImpersonationAction_NoAction   ImpersonationAction = "NoAction"
	ImpersonationAction_MoveToJmf  ImpersonationAction = "MoveToJmf"
	ImpersonationAction_Redirect   ImpersonationAction = "Redirect"
	ImpersonationAction_Quarantine ImpersonationAction = "Quarantine"
	ImpersonationAction_Delete     ImpersonationAction = "Delete"
	ImpersonationAction_BccMessage ImpersonationAction = "BccMessage"
)

// ImpersonationProtectionState enum.
type ImpersonationProtectionState string

const (
	ImpersonationProtectionState_Default   ImpersonationProtectionState = "Default"
	ImpersonationProtectionState_Automatic ImpersonationProtectionState = "Automatic"
	ImpersonationProtectionState_Manual    ImpersonationProtectionState = "Manual"
	ImpersonationProtectionState_Off       ImpersonationProtectionState = "Off"
)

// IntraOrgFilterState enum.
type IntraOrgFilterState string

const (
	IntraOrgFilterState_Default             IntraOrgFilterState = "Default"
	IntraOrgFilterState_Disabled            IntraOrgFilterState = "Disabled"
	IntraOrgFilterState_HighConfidencePhish IntraOrgFilterState = "HighConfidencePhish"
	IntraOrgFilterState_Phish               IntraOrgFilterState = "Phish"
	IntraOrgFilterState_HighConfidenceSpam  IntraOrgFilterState = "HighConfidenceSpam"
	IntraOrgFilterState_Spam                IntraOrgFilterState = "Spam"
)

// MalwareFilteringAction enum.
type MalwareFilteringAction string

const (
	MalwareFilteringAction_DeleteMessage                          MalwareFilteringAction = "DeleteMessage"
	MalwareFilteringAction_DeleteAttachmentAndUseDefaultAlertText MalwareFilteringAction = "DeleteAttachmentAndUseDefaultAlertText"
	MalwareFilteringAction_DeleteAttachmentAndUseCustomAlertText  MalwareFilteringAction = "DeleteAttachmentAndUseCustomAlertText"
)

// OofState enum.
type OofState string

const (
	OofState_Disabled  OofState = "Disabled"
	OofState_Enabled   OofState = "Enabled"
	OofState_Scheduled OofState = "Scheduled"
)

// PhishFilteringAction enum.
type PhishFilteringAction string

const (
	PhishFilteringAction_Quarantine PhishFilteringAction = "Quarantine"
	PhishFilteringAction_MoveToJmf  PhishFilteringAction = "MoveToJmf"
	PhishFilteringAction_Redirect   PhishFilteringAction = "Redirect"
)

// RecipientAccessRight enum.
type RecipientAccessRight string

const (
	RecipientAccessRight_SendAs RecipientAccessRight = "SendAs"
)

// RecommendedPolicyType enum.
type RecommendedPolicyType string

const (
	RecommendedPolicyType_Custom     RecommendedPolicyType = "Custom"
	RecommendedPolicyType_Standard   RecommendedPolicyType = "Standard"
	RecommendedPolicyType_Strict     RecommendedPolicyType = "Strict"
	RecommendedPolicyType_Evaluation RecommendedPolicyType = "Evaluation"
)

// RolePrincipalTypes enum.
type RolePrincipalTypes string

const (
	RolePrincipalTypes_User             RolePrincipalTypes = "User"
	RolePrincipalTypes_ServicePrincipal RolePrincipalTypes = "ServicePrincipal"
	RolePrincipalTypes_Group            RolePrincipalTypes = "Group"
)

// SafeAttachmentAction enum.
type SafeAttachmentAction string

const (
	SafeAttachmentAction_Block           SafeAttachmentAction = "Block"
	SafeAttachmentAction_Replace         SafeAttachmentAction = "Replace"
	SafeAttachmentAction_Allow           SafeAttachmentAction = "Allow"
	SafeAttachmentAction_DynamicDelivery SafeAttachmentAction = "DynamicDelivery"
)

// SafeAttachmentOperationMode enum.
type SafeAttachmentOperationMode string

const (
	SafeAttachmentOperationMode_Delay   SafeAttachmentOperationMode = "Delay"
	SafeAttachmentOperationMode_Deliver SafeAttachmentOperationMode = "Deliver"
)

// ScopeRestrictionType enum.
type ScopeRestrictionType string

const (
	ScopeRestrictionType_NotApplicable               ScopeRestrictionType = "NotApplicable"
	ScopeRestrictionType_DomainScopeObsolete         ScopeRestrictionType = "DomainScope_Obsolete"
	ScopeRestrictionType_RecipientScope              ScopeRestrictionType = "RecipientScope"
	ScopeRestrictionType_ServerScope                 ScopeRestrictionType = "ServerScope"
	ScopeRestrictionType_PartnerDelegatedTenantScope ScopeRestrictionType = "PartnerDelegatedTenantScope"
	ScopeRestrictionType_DatabaseScope               ScopeRestrictionType = "DatabaseScope"
)

// SpamFilteringAction enum.
type SpamFilteringAction string

const (
	SpamFilteringAction_MoveToJmf     SpamFilteringAction = "MoveToJmf"
	SpamFilteringAction_AddXHeader    SpamFilteringAction = "AddXHeader"
	SpamFilteringAction_ModifySubject SpamFilteringAction = "ModifySubject"
	SpamFilteringAction_Redirect      SpamFilteringAction = "Redirect"
	SpamFilteringAction_Delete        SpamFilteringAction = "Delete"
	SpamFilteringAction_Quarantine    SpamFilteringAction = "Quarantine"
	SpamFilteringAction_NoAction      SpamFilteringAction = "NoAction"
	SpamFilteringAction_BccMessage    SpamFilteringAction = "BccMessage"
)

// SpamFilteringOption enum.
type SpamFilteringOption string

const (
	SpamFilteringOption_Off  SpamFilteringOption = "Off"
	SpamFilteringOption_On   SpamFilteringOption = "On"
	SpamFilteringOption_Test SpamFilteringOption = "Test"
)

// SpamFilteringTestModeAction enum.
type SpamFilteringTestModeAction string

const (
	SpamFilteringTestModeAction_None       SpamFilteringTestModeAction = "None"
	SpamFilteringTestModeAction_AddXHeader SpamFilteringTestModeAction = "AddXHeader"
	SpamFilteringTestModeAction_BccMessage SpamFilteringTestModeAction = "BccMessage"
)

// SpoofAuthenticationFailAction enum.
type SpoofAuthenticationFailAction string

const (
	SpoofAuthenticationFailAction_MoveToJmf  SpoofAuthenticationFailAction = "MoveToJmf"
	SpoofAuthenticationFailAction_Quarantine SpoofAuthenticationFailAction = "Quarantine"
)

// SpoofDmarcQuarantineAction enum.
type SpoofDmarcQuarantineAction string

const (
	SpoofDmarcQuarantineAction_Quarantine SpoofDmarcQuarantineAction = "Quarantine"
	SpoofDmarcQuarantineAction_MoveToJmf  SpoofDmarcQuarantineAction = "MoveToJmf"
)

// SpoofDmarcRejectAction enum.
type SpoofDmarcRejectAction string

const (
	SpoofDmarcRejectAction_Reject     SpoofDmarcRejectAction = "Reject"
	SpoofDmarcRejectAction_Quarantine SpoofDmarcRejectAction = "Quarantine"
)

// ActiveSyncDeviceAccessRule (OData ComplexType).
type ActiveSyncDeviceAccessRule struct {
	Guid           string `json:"Guid,omitempty"`
	Name           string `json:"Name,omitempty"`
	QueryString    string `json:"QueryString,omitempty"`
	Characteristic string `json:"Characteristic,omitempty"`
	AccessLevel    string `json:"AccessLevel,omitempty"`
}

// ActiveSyncDeviceClass (OData ComplexType).
type ActiveSyncDeviceClass struct {
	Guid        string `json:"Guid,omitempty"`
	Id          string `json:"Id,omitempty"`
	Name        string `json:"Name,omitempty"`
	DeviceModel string `json:"DeviceModel,omitempty"`
	DeviceType  string `json:"DeviceType,omitempty"`
}

// ActiveSyncOrganizationSettings (OData ComplexType).
type ActiveSyncOrganizationSettings struct {
	Identity            string   `json:"Identity,omitempty"`
	Guid                string   `json:"Guid,omitempty"`
	DefaultAccessLevel  string   `json:"DefaultAccessLevel,omitempty"`
	UserMailInsert      string   `json:"UserMailInsert,omitempty"`
	AdminMailRecipients []string `json:"AdminMailRecipients,omitempty"`
}

// AddressBookPolicy (OData ComplexType).
type AddressBookPolicy struct {
	Identity         string    `json:"Identity,omitempty"`
	AdminDisplayName string    `json:"AdminDisplayName,omitempty"`
	IsDefault        bool      `json:"IsDefault,omitempty"`
	IsValid          bool      `json:"IsValid,omitempty"`
	Name             string    `json:"Name,omitempty"`
	WhenChangedUTC   time.Time `json:"WhenChangedUTC,omitempty"`
	WhenCreatedUTC   time.Time `json:"WhenCreatedUTC,omitempty"`
}

// AntiPhishPolicyPresentation (OData ComplexType).
type AntiPhishPolicyPresentation struct {
	Identity                                      string                        `json:"Identity,omitempty"`
	Enabled                                       bool                          `json:"Enabled,omitempty"`
	ImpersonationProtectionState                  ImpersonationProtectionState  `json:"ImpersonationProtectionState,omitempty"`
	EnableTargetedUserProtection                  bool                          `json:"EnableTargetedUserProtection,omitempty"`
	EnableMailboxIntelligenceProtection           bool                          `json:"EnableMailboxIntelligenceProtection,omitempty"`
	EnableTargetedDomainsProtection               bool                          `json:"EnableTargetedDomainsProtection,omitempty"`
	EnableOrganizationDomainsProtection           bool                          `json:"EnableOrganizationDomainsProtection,omitempty"`
	EnableMailboxIntelligence                     bool                          `json:"EnableMailboxIntelligence,omitempty"`
	EnableFirstContactSafetyTips                  bool                          `json:"EnableFirstContactSafetyTips,omitempty"`
	EnableSimilarUsersSafetyTips                  bool                          `json:"EnableSimilarUsersSafetyTips,omitempty"`
	EnableSimilarDomainsSafetyTips                bool                          `json:"EnableSimilarDomainsSafetyTips,omitempty"`
	EnableUnusualCharactersSafetyTips             bool                          `json:"EnableUnusualCharactersSafetyTips,omitempty"`
	TargetedUserProtectionAction                  ImpersonationAction           `json:"TargetedUserProtectionAction,omitempty"`
	MailboxIntelligenceProtectionAction           ImpersonationAction           `json:"MailboxIntelligenceProtectionAction,omitempty"`
	TargetedDomainProtectionAction                ImpersonationAction           `json:"TargetedDomainProtectionAction,omitempty"`
	AuthenticationFailAction                      SpoofAuthenticationFailAction `json:"AuthenticationFailAction,omitempty"`
	EnableSpoofIntelligence                       bool                          `json:"EnableSpoofIntelligence,omitempty"`
	EnableViaTag                                  bool                          `json:"EnableViaTag,omitempty"`
	EnableUnauthenticatedSender                   bool                          `json:"EnableUnauthenticatedSender,omitempty"`
	EnableSuspiciousSafetyTip                     bool                          `json:"EnableSuspiciousSafetyTip,omitempty"`
	PhishThresholdLevel                           int32                         `json:"PhishThresholdLevel,omitempty"`
	TargetedUserActionRecipients                  []string                      `json:"TargetedUserActionRecipients,omitempty"`
	MailboxIntelligenceProtectionActionRecipients []string                      `json:"MailboxIntelligenceProtectionActionRecipients,omitempty"`
	TargetedDomainActionRecipients                []string                      `json:"TargetedDomainActionRecipients,omitempty"`
	ExcludedSenders                               []string                      `json:"ExcludedSenders,omitempty"`
	IsDefault                                     bool                          `json:"IsDefault,omitempty"`
	AdminDisplayName                              string                        `json:"AdminDisplayName,omitempty"`
	PolicyTag                                     string                        `json:"PolicyTag,omitempty"`
	RecommendedPolicyType                         RecommendedPolicyType         `json:"RecommendedPolicyType,omitempty"`
	Name                                          string                        `json:"Name,omitempty"`
	WhenChanged                                   time.Time                     `json:"WhenChanged,omitempty"`
	WhenCreated                                   time.Time                     `json:"WhenCreated,omitempty"`
	ExchangeObjectId                              string                        `json:"ExchangeObjectId,omitempty"`
	OrganizationId                                string                        `json:"OrganizationId,omitempty"`
	Guid                                          string                        `json:"Guid,omitempty"`
	HonorDmarcPolicy                              bool                          `json:"HonorDmarcPolicy,omitempty"`
	DmarcRejectAction                             SpoofDmarcRejectAction        `json:"DmarcRejectAction,omitempty"`
	DmarcQuarantineAction                         SpoofDmarcQuarantineAction    `json:"DmarcQuarantineAction,omitempty"`
}

// AntiPhishRule (OData ComplexType).
type AntiPhishRule struct {
	Identity       string    `json:"Identity,omitempty"`
	State          string    `json:"State,omitempty"`
	Priority       int32     `json:"Priority,omitempty"`
	Comments       string    `json:"Comments,omitempty"`
	Guid           string    `json:"Guid,omitempty"`
	ImmutableId    string    `json:"ImmutableId,omitempty"`
	OrganizationId string    `json:"OrganizationId,omitempty"`
	Name           string    `json:"Name,omitempty"`
	WhenChanged    time.Time `json:"WhenChanged,omitempty"`
}

// AtpPolicyForO365Presentation (OData ComplexType).
type AtpPolicyForO365Presentation struct {
	Identity                string    `json:"Identity,omitempty"`
	EnableATPForSPOTeamsODB bool      `json:"EnableATPForSPOTeamsODB,omitempty"`
	EnableSafeDocs          bool      `json:"EnableSafeDocs,omitempty"`
	AllowSafeDocsOpen       bool      `json:"AllowSafeDocsOpen,omitempty"`
	AdminDisplayName        string    `json:"AdminDisplayName,omitempty"`
	Name                    string    `json:"Name,omitempty"`
	WhenChanged             time.Time `json:"WhenChanged,omitempty"`
	WhenCreated             time.Time `json:"WhenCreated,omitempty"`
	ExchangeObjectId        string    `json:"ExchangeObjectId,omitempty"`
	OrganizationId          string    `json:"OrganizationId,omitempty"`
	Guid                    string    `json:"Guid,omitempty"`
}

// BasicInfo (OData ComplexType).
type BasicInfo struct {
	PropertyName  string `json:"PropertyName,omitempty"`
	PropertyValue string `json:"PropertyValue,omitempty"`
}

// ByteArrayType (OData ComplexType).
type ByteArrayType struct {
	ComplexEntry
	Data []byte `json:"Data,omitempty"`
}

// CalendarProcessing (OData ComplexType).
type CalendarProcessing struct {
	Identity                    string                  `json:"Identity,omitempty"`
	AllBookInPolicy             bool                    `json:"AllBookInPolicy,omitempty"`
	AllRequestInPolicy          bool                    `json:"AllRequestInPolicy,omitempty"`
	AddAdditionalResponse       bool                    `json:"AddAdditionalResponse,omitempty"`
	AdditionalResponse          string                  `json:"AdditionalResponse,omitempty"`
	AllowRecurringMeetings      bool                    `json:"AllowRecurringMeetings,omitempty"`
	AutomateProcessing          CalendarProcessingFlags `json:"AutomateProcessing,omitempty"`
	BookingWindowInDays         int32                   `json:"BookingWindowInDays,omitempty"`
	EnforceSchedulingHorizon    bool                    `json:"EnforceSchedulingHorizon,omitempty"`
	MaximumDurationInMinutes    int32                   `json:"MaximumDurationInMinutes,omitempty"`
	ScheduleOnlyDuringWorkHours bool                    `json:"ScheduleOnlyDuringWorkHours,omitempty"`
	ResourceDelegates           []string                `json:"ResourceDelegates,omitempty"`
}

// CasMailbox (OData ComplexType).
type CasMailbox struct {
	ObjectKey                               string            `json:"ObjectKey,omitempty"`
	ExternalDirectoryObjectId               string            `json:"ExternalDirectoryObjectId,omitempty"`
	Identity                                string            `json:"Identity,omitempty"`
	PrimarySmtpAddress                      string            `json:"PrimarySmtpAddress,omitempty"`
	DisplayName                             string            `json:"DisplayName,omitempty"`
	Name                                    string            `json:"Name,omitempty"`
	Guid                                    string            `json:"Guid,omitempty"`
	ServerLegacyDN                          string            `json:"ServerLegacyDN,omitempty"`
	ExchangeVersion                         string            `json:"ExchangeVersion,omitempty"`
	UniversalOutlookEnabled                 bool              `json:"UniversalOutlookEnabled,omitempty"`
	OutlookMobileEnabled                    bool              `json:"OutlookMobileEnabled,omitempty"`
	MacOutlookEnabled                       bool              `json:"MacOutlookEnabled,omitempty"`
	ECPEnabled                              bool              `json:"ECPEnabled,omitempty"`
	OWAforDevicesEnabled                    bool              `json:"OWAforDevicesEnabled,omitempty"`
	ShowGalAsDefaultView                    bool              `json:"ShowGalAsDefaultView,omitempty"`
	EmailAddresses                          []string          `json:"EmailAddresses,omitempty"`
	LegacyExchangeDN                        string            `json:"LegacyExchangeDN,omitempty"`
	LinkedMasterAccount                     string            `json:"LinkedMasterAccount,omitempty"`
	SamAccountName                          string            `json:"SamAccountName,omitempty"`
	SmtpClientAuthenticationDisabled        bool              `json:"SmtpClientAuthenticationDisabled,omitempty"`
	OneWinNativeOutlookEnabled              bool              `json:"OneWinNativeOutlookEnabled,omitempty"`
	OWAEnabled                              bool              `json:"OWAEnabled,omitempty"`
	DistinguishedName                       string            `json:"DistinguishedName,omitempty"`
	ExchangeObjectId                        string            `json:"ExchangeObjectId,omitempty"`
	ObjectCategory                          string            `json:"ObjectCategory,omitempty"`
	ObjectClass                             []string          `json:"ObjectClass,omitempty"`
	OrganizationId                          string            `json:"OrganizationId,omitempty"`
	PublicFolderClientAccess                bool              `json:"PublicFolderClientAccess,omitempty"`
	WhenChangedUTC                          string            `json:"WhenChangedUTC,omitempty"`
	WhenCreatedUTC                          string            `json:"WhenCreatedUTC,omitempty"`
	WhenChanged                             string            `json:"WhenChanged,omitempty"`
	WhenCreated                             string            `json:"WhenCreated,omitempty"`
	OwaMailboxPolicy                        string            `json:"OwaMailboxPolicy,omitempty"`
	IsOptimizedForAccessibility             bool              `json:"IsOptimizedForAccessibility,omitempty"`
	ImapEnabled                             bool              `json:"ImapEnabled,omitempty"`
	ImapSuppressReadReceipt                 bool              `json:"ImapSuppressReadReceipt,omitempty"`
	ImapEnableExactRFC822Size               bool              `json:"ImapEnableExactRFC822Size,omitempty"`
	ImapMessagesRetrievalMimeFormat         string            `json:"ImapMessagesRetrievalMimeFormat,omitempty"`
	ImapUseProtocolDefaults                 bool              `json:"ImapUseProtocolDefaults,omitempty"`
	ImapForceICalForCalendarRetrievalOption bool              `json:"ImapForceICalForCalendarRetrievalOption,omitempty"`
	PopEnabled                              bool              `json:"PopEnabled,omitempty"`
	PopSuppressReadReceipt                  bool              `json:"PopSuppressReadReceipt,omitempty"`
	PopEnableExactRFC822Size                bool              `json:"PopEnableExactRFC822Size,omitempty"`
	PopMessagesRetrievalMimeFormat          string            `json:"PopMessagesRetrievalMimeFormat,omitempty"`
	PopUseProtocolDefaults                  bool              `json:"PopUseProtocolDefaults,omitempty"`
	PopMessageDeleteEnabled                 bool              `json:"PopMessageDeleteEnabled,omitempty"`
	PopForceICalForCalendarRetrievalOption  bool              `json:"PopForceICalForCalendarRetrievalOption,omitempty"`
	MAPIEnabled                             bool              `json:"MAPIEnabled,omitempty"`
	MAPIBlockOutlookVersions                string            `json:"MAPIBlockOutlookVersions,omitempty"`
	MAPIBlockOutlookRpcHttp                 bool              `json:"MAPIBlockOutlookRpcHttp,omitempty"`
	MapiHttpEnabled                         bool              `json:"MapiHttpEnabled,omitempty"`
	MAPIBlockOutlookNonCachedMode           bool              `json:"MAPIBlockOutlookNonCachedMode,omitempty"`
	MAPIBlockOutlookExternalConnectivity    bool              `json:"MAPIBlockOutlookExternalConnectivity,omitempty"`
	EwsEnabled                              bool              `json:"EwsEnabled,omitempty"`
	EwsAllowOutlook                         bool              `json:"EwsAllowOutlook,omitempty"`
	EwsAllowMacOutlook                      bool              `json:"EwsAllowMacOutlook,omitempty"`
	EwsAllowEntourage                       bool              `json:"EwsAllowEntourage,omitempty"`
	EwsApplicationAccessPolicy              string            `json:"EwsApplicationAccessPolicy,omitempty"`
	EwsAllowList                            []string          `json:"EwsAllowList,omitempty"`
	EwsBlockList                            []string          `json:"EwsBlockList,omitempty"`
	ActiveSyncAllowedDeviceIDs              []string          `json:"ActiveSyncAllowedDeviceIDs,omitempty"`
	ActiveSyncBlockedDeviceIDs              []string          `json:"ActiveSyncBlockedDeviceIDs,omitempty"`
	ActiveSyncEnabled                       bool              `json:"ActiveSyncEnabled,omitempty"`
	ActiveSyncSuppressReadReceipt           bool              `json:"ActiveSyncSuppressReadReceipt,omitempty"`
	ActiveSyncMailboxPolicyIsDefaulted      bool              `json:"ActiveSyncMailboxPolicyIsDefaulted,omitempty"`
	ActiveSyncMailboxPolicy                 string            `json:"ActiveSyncMailboxPolicy,omitempty"`
	HasActiveSyncDevicePartnership          bool              `json:"HasActiveSyncDevicePartnership,omitempty"`
	ActiveSyncDebugLogging                  bool              `json:"ActiveSyncDebugLogging,omitempty"`
	ExternalImapSettings                    string            `json:"ExternalImapSettings,omitempty"`
	InternalImapSettings                    string            `json:"InternalImapSettings,omitempty"`
	ExternalPopSettings                     string            `json:"ExternalPopSettings,omitempty"`
	InternalPopSettings                     string            `json:"InternalPopSettings,omitempty"`
	ExternalSmtpSettings                    string            `json:"ExternalSmtpSettings,omitempty"`
	InternalSmtpSettings                    string            `json:"InternalSmtpSettings,omitempty"`
	DeltaUpdates                            *GenericHashTable `json:"DeltaUpdates,omitempty"`
}

// CmdletInfo (OData ComplexType).
type CmdletInfo struct {
	CmdletName string   `json:"CmdletName,omitempty"`
	Parameters []string `json:"Parameters,omitempty"`
}

// CmdletInvokeInputType (OData ComplexType).
type CmdletInvokeInputType struct {
	ComplexEntry
	CmdletName string            `json:"CmdletName,omitempty"`
	Parameters *GenericHashTable `json:"Parameters,omitempty"`
}

// ComplexEntry (OData ComplexType).
type ComplexEntry struct {
	ChangedProperties []string `json:"ChangedProperties,omitempty"`
}

// ConfigAnalyzerPolicyRecommendation (OData ComplexType).
type ConfigAnalyzerPolicyRecommendation struct {
	Identity               string `json:"Identity,omitempty"`
	PolicyGroup            string `json:"PolicyGroup,omitempty"`
	SettingName            string `json:"SettingName,omitempty"`
	SettingNameDescription string `json:"SettingNameDescription,omitempty"`
	Policy                 string `json:"Policy,omitempty"`
	AppliedTo              string `json:"AppliedTo,omitempty"`
	CurrentConfiguration   string `json:"CurrentConfiguration,omitempty"`
	LastModified           string `json:"LastModified,omitempty"`
	Recommendation         string `json:"Recommendation,omitempty"`
	SettingType            string `json:"SettingType,omitempty"`
	Cmdlet                 string `json:"Cmdlet,omitempty"`
	RuleName               string `json:"RuleName,omitempty"`
}

// DirectMobileDevice (OData ComplexType).
type DirectMobileDevice struct {
	Identity                string    `json:"Identity,omitempty"`
	Guid                    string    `json:"Guid,omitempty"`
	UserDisplayName         string    `json:"UserDisplayName,omitempty"`
	ClientType              string    `json:"ClientType,omitempty"`
	ClientVersion           string    `json:"ClientVersion,omitempty"`
	DeviceAccessControlRule string    `json:"DeviceAccessControlRule,omitempty"`
	DeviceAccessState       string    `json:"DeviceAccessState,omitempty"`
	DeviceAccessStateReason string    `json:"DeviceAccessStateReason,omitempty"`
	DeviceId                string    `json:"DeviceId,omitempty"`
	DeviceImei              string    `json:"DeviceImei,omitempty"`
	DeviceMobileOperator    string    `json:"DeviceMobileOperator,omitempty"`
	DeviceModel             string    `json:"DeviceModel,omitempty"`
	DeviceOS                string    `json:"DeviceOS,omitempty"`
	DeviceOSLanguage        string    `json:"DeviceOSLanguage,omitempty"`
	DeviceTelephoneNumber   string    `json:"DeviceTelephoneNumber,omitempty"`
	DeviceType              string    `json:"DeviceType,omitempty"`
	DeviceUserAgent         string    `json:"DeviceUserAgent,omitempty"`
	FriendlyName            string    `json:"FriendlyName,omitempty"`
	IsCompliant             bool      `json:"IsCompliant,omitempty"`
	IsDisabled              bool      `json:"IsDisabled,omitempty"`
	IsManaged               bool      `json:"IsManaged,omitempty"`
	FirstSyncTime           time.Time `json:"FirstSyncTime,omitempty"`
}

// Divergence (OData ComplexType).
type Divergence struct {
	ObjectId                        string   `json:"ObjectId,omitempty"`
	ErrorMessage                    string   `json:"ErrorMessage,omitempty"`
	IsIncrementalOnly               bool     `json:"IsIncrementalOnly,omitempty"`
	IsLinkRelated                   bool     `json:"IsLinkRelated,omitempty"`
	IsTemporary                     bool     `json:"IsTemporary,omitempty"`
	IsValidationDivergence          bool     `json:"IsValidationDivergence,omitempty"`
	IsIgnoredInHaltCondition        bool     `json:"IsIgnoredInHaltCondition,omitempty"`
	IsHighPriority                  bool     `json:"IsHighPriority,omitempty"`
	WhenNextExecutingUTC            string   `json:"WhenNextExecutingUTC,omitempty"`
	WhenLastExecutedUTC             string   `json:"WhenLastExecutedUTC,omitempty"`
	WhenCreatedUTC                  string   `json:"WhenCreatedUTC,omitempty"`
	ExecutionCount                  int32    `json:"ExecutionCount,omitempty"`
	DivergenceCount                 int32    `json:"DivergenceCount,omitempty"`
	ObjectGuid                      string   `json:"ObjectGuid,omitempty"`
	IsRetriable                     bool     `json:"IsRetriable,omitempty"`
	Comment                         string   `json:"Comment,omitempty"`
	IsTenantWideDivergence          bool     `json:"IsTenantWideDivergence,omitempty"`
	Errors                          []string `json:"Errors,omitempty"`
	ExternalDirectoryObjectClass    string   `json:"ExternalDirectoryObjectClass,omitempty"`
	CmdletName                      string   `json:"CmdletName,omitempty"`
	CmdletParameters                string   `json:"CmdletParameters,omitempty"`
	ExternalDirectoryOrganizationId string   `json:"ExternalDirectoryOrganizationId,omitempty"`
	ExternalDirectoryObjectId       string   `json:"ExternalDirectoryObjectId,omitempty"`
}

// DynamicDistributionGroup (OData ComplexType).
type DynamicDistributionGroup struct {
	Identity                                               string    `json:"Identity,omitempty"`
	Id                                                     string    `json:"Id,omitempty"`
	Guid                                                   string    `json:"Guid,omitempty"`
	Notes                                                  string    `json:"Notes,omitempty"`
	Name                                                   string    `json:"Name,omitempty"`
	DisplayName                                            string    `json:"DisplayName,omitempty"`
	Alias                                                  string    `json:"Alias,omitempty"`
	DistinguishedName                                      string    `json:"DistinguishedName,omitempty"`
	WindowsLiveID                                          string    `json:"WindowsLiveID,omitempty"`
	PrimarySmtpAddress                                     string    `json:"PrimarySmtpAddress,omitempty"`
	IncludedRecipients                                     int32     `json:"IncludedRecipients,omitempty"`
	LdapRecipientFilter                                    string    `json:"LdapRecipientFilter,omitempty"`
	RecipientFilterType                                    string    `json:"RecipientFilterType,omitempty"`
	RecipientFilter                                        string    `json:"RecipientFilter,omitempty"`
	ManagedBy                                              string    `json:"ManagedBy,omitempty"`
	HiddenFromAddressListsEnabled                          bool      `json:"HiddenFromAddressListsEnabled,omitempty"`
	ConditionalCustomAttribute15                           []string  `json:"ConditionalCustomAttribute15,omitempty"`
	ConditionalCustomAttribute14                           []string  `json:"ConditionalCustomAttribute14,omitempty"`
	ConditionalCustomAttribute13                           []string  `json:"ConditionalCustomAttribute13,omitempty"`
	ConditionalCustomAttribute12                           []string  `json:"ConditionalCustomAttribute12,omitempty"`
	ConditionalCustomAttribute11                           []string  `json:"ConditionalCustomAttribute11,omitempty"`
	ConditionalCustomAttribute10                           []string  `json:"ConditionalCustomAttribute10,omitempty"`
	ConditionalCustomAttribute9                            []string  `json:"ConditionalCustomAttribute9,omitempty"`
	ConditionalCustomAttribute8                            []string  `json:"ConditionalCustomAttribute8,omitempty"`
	ConditionalCustomAttribute7                            []string  `json:"ConditionalCustomAttribute7,omitempty"`
	ConditionalCustomAttribute6                            []string  `json:"ConditionalCustomAttribute6,omitempty"`
	ConditionalCustomAttribute5                            []string  `json:"ConditionalCustomAttribute5,omitempty"`
	ConditionalCustomAttribute4                            []string  `json:"ConditionalCustomAttribute4,omitempty"`
	ConditionalCustomAttribute3                            []string  `json:"ConditionalCustomAttribute3,omitempty"`
	ConditionalCustomAttribute2                            []string  `json:"ConditionalCustomAttribute2,omitempty"`
	ConditionalCustomAttribute1                            []string  `json:"ConditionalCustomAttribute1,omitempty"`
	ConditionalStateOrProvince                             []string  `json:"ConditionalStateOrProvince,omitempty"`
	ConditionalCompany                                     []string  `json:"ConditionalCompany,omitempty"`
	ConditionalDepartment                                  []string  `json:"ConditionalDepartment,omitempty"`
	AcceptMessagesOnlyFromSendersOrMembers                 []string  `json:"AcceptMessagesOnlyFromSendersOrMembers,omitempty"`
	BypassModerationFromSendersOrMembersWithDisplayNames   []string  `json:"BypassModerationFromSendersOrMembersWithDisplayNames,omitempty"`
	AcceptMessagesOnlyFromSendersOrMembersWithDisplayNames []string  `json:"AcceptMessagesOnlyFromSendersOrMembersWithDisplayNames,omitempty"`
	GrantSendOnBehalfToWithDisplayNames                    []string  `json:"GrantSendOnBehalfToWithDisplayNames,omitempty"`
	ModeratedByWithDisplayNames                            []string  `json:"ModeratedByWithDisplayNames,omitempty"`
	ManagedByWithDisplayName                               []string  `json:"ManagedByWithDisplayName,omitempty"`
	RequireSenderAuthenticationEnabled                     bool      `json:"RequireSenderAuthenticationEnabled,omitempty"`
	ModerationEnabled                                      bool      `json:"ModerationEnabled,omitempty"`
	ModeratedBy                                            []string  `json:"ModeratedBy,omitempty"`
	BypassModerationFromSendersOrMembers                   []string  `json:"BypassModerationFromSendersOrMembers,omitempty"`
	GrantSendOnBehalfTo                                    []string  `json:"GrantSendOnBehalfTo,omitempty"`
	SendModerationNotifications                            string    `json:"SendModerationNotifications,omitempty"`
	EmailAddresses                                         []string  `json:"EmailAddresses,omitempty"`
	SendAsPermissionList                                   []string  `json:"SendAsPermissionList,omitempty"`
	WhenChanged                                            string    `json:"WhenChanged,omitempty"`
	WhenCreated                                            string    `json:"WhenCreated,omitempty"`
	WhenChangedUTC                                         time.Time `json:"WhenChangedUTC,omitempty"`
	WhenCreatedUTC                                         time.Time `json:"WhenCreatedUTC,omitempty"`
	RecipientType                                          string    `json:"RecipientType,omitempty"`
	RecipientTypeDetails                                   string    `json:"RecipientTypeDetails,omitempty"`
	IsValid                                                bool      `json:"IsValid,omitempty"`
}

// EligibleDistributionGroup (OData ComplexType).
type EligibleDistributionGroup struct {
	Identity           string `json:"Identity,omitempty"`
	DisplayName        string `json:"DisplayName,omitempty"`
	PrimarySmtpAddress string `json:"PrimarySmtpAddress,omitempty"`
	LegacyExchangeDN   string `json:"LegacyExchangeDN,omitempty"`
}

// ExchangeManagementScope (OData ComplexType).
type ExchangeManagementScope struct {
	Id                         string               `json:"id,omitempty"`
	RecipientRoot              string               `json:"recipientRoot,omitempty"`
	Filter                     string               `json:"filter,omitempty"`
	RecipientFilter            string               `json:"recipientFilter,omitempty"`
	ServerFilter               string               `json:"serverFilter,omitempty"`
	DatabaseFilter             string               `json:"databaseFilter,omitempty"`
	TenantOrganizationFilter   string               `json:"tenantOrganizationFilter,omitempty"`
	ScopeRestrictionType       ScopeRestrictionType `json:"scopeRestrictionType,omitempty"`
	Exclusive                  bool                 `json:"exclusive,omitempty"`
	Name                       string               `json:"name,omitempty"`
	AdminDisplayName           string               `json:"adminDisplayName,omitempty"`
	Version                    string               `json:"version,omitempty"`
	DistinguishedName          string               `json:"distinguishedName,omitempty"`
	Guid                       string               `json:"guid,omitempty"`
	Identity                   string               `json:"identity,omitempty"`
	RecipientRestrictionFilter string               `json:"recipientRestrictionFilter,omitempty"`
}

// ExchangeRoleGroup (OData EntityType).
type ExchangeRoleGroup struct {
	Id                               string                     `json:"id,omitempty"`
	Identity                         string                     `json:"identity,omitempty"`
	Description                      string                     `json:"description,omitempty"`
	Name                             string                     `json:"name,omitempty"`
	ManagedBy                        []string                   `json:"managedBy,omitempty"`
	RoleAssignments                  []string                   `json:"roleAssignments,omitempty"`
	Roles                            []string                   `json:"roles,omitempty"`
	DisplayName                      string                     `json:"displayName,omitempty"`
	ExternalDirectoryObjectId        string                     `json:"externalDirectoryObjectId,omitempty"`
	SamAccountName                   string                     `json:"samAccountName,omitempty"`
	RoleGroupType                    string                     `json:"roleGroupType,omitempty"`
	LinkedGroup                      string                     `json:"linkedGroup,omitempty"`
	Capabilities                     []string                   `json:"capabilities,omitempty"`
	LinkedPartnerGroupId             string                     `json:"linkedPartnerGroupId,omitempty"`
	LinkedPartnerOrganizationId      string                     `json:"linkedPartnerOrganizationId,omitempty"`
	WellKnownObject                  string                     `json:"wellKnownObject,omitempty"`
	DistinguishedName                string                     `json:"distinguishedName,omitempty"`
	ObjectCategory                   string                     `json:"objectCategory,omitempty"`
	ObjectClass                      []string                   `json:"objectClass,omitempty"`
	IsValid                          bool                       `json:"isValid,omitempty"`
	WhenChangedDateTime              time.Time                  `json:"whenChangedDateTime,omitempty"`
	WhenCreatedDateTime              time.Time                  `json:"whenCreatedDateTime,omitempty"`
	WhenChangedUTCDateTime           time.Time                  `json:"whenChangedUTCDateTime,omitempty"`
	WhenCreatedUTCDateTime           time.Time                  `json:"whenCreatedUTCDateTime,omitempty"`
	ExchangeObjectId                 string                     `json:"exchangeObjectId,omitempty"`
	OrganizationId                   string                     `json:"organizationId,omitempty"`
	Guid                             string                     `json:"guid,omitempty"`
	OriginatingServer                string                     `json:"originatingServer,omitempty"`
	ObjectState                      string                     `json:"objectState,omitempty"`
	Version                          string                     `json:"version,omitempty"`
	CustomConfigWriteScope           string                     `json:"customConfigWriteScope,omitempty"`
	CustomRecipientWriteScope        string                     `json:"customRecipientWriteScope,omitempty"`
	RecipientOrganizationalUnitScope string                     `json:"recipientOrganizationalUnitScope,omitempty"`
	NewMembers                       []string                   `json:"newMembers,omitempty"`
	Members                          []*ExchangeRoleGroupMember `json:"members,omitempty"`
}

// ExchangeRoleGroupMember (OData ComplexType).
type ExchangeRoleGroupMember struct {
	ObjectKey                          string    `json:"objectKey,omitempty"`
	ExternalDirectoryObjectId          string    `json:"externalDirectoryObjectId,omitempty"`
	Identity                           string    `json:"identity,omitempty"`
	Alias                              string    `json:"alias,omitempty"`
	EmailAddresses                     []string  `json:"emailAddresses,omitempty"`
	DisplayName                        string    `json:"displayName,omitempty"`
	FirstName                          string    `json:"firstName,omitempty"`
	LastName                           string    `json:"lastName,omitempty"`
	Name                               string    `json:"name,omitempty"`
	ArchiveGuid                        string    `json:"archiveGuid,omitempty"`
	AuthenticationType                 string    `json:"authenticationType,omitempty"`
	City                               string    `json:"city,omitempty"`
	Notes                              string    `json:"notes,omitempty"`
	Company                            string    `json:"company,omitempty"`
	CountryOrRegion                    string    `json:"countryOrRegion,omitempty"`
	PostalCode                         string    `json:"postalCode,omitempty"`
	CustomAttribute1                   string    `json:"customAttribute1,omitempty"`
	CustomAttribute2                   string    `json:"customAttribute2,omitempty"`
	CustomAttribute3                   string    `json:"customAttribute3,omitempty"`
	CustomAttribute4                   string    `json:"customAttribute4,omitempty"`
	CustomAttribute5                   string    `json:"customAttribute5,omitempty"`
	CustomAttribute6                   string    `json:"customAttribute6,omitempty"`
	CustomAttribute7                   string    `json:"customAttribute7,omitempty"`
	CustomAttribute8                   string    `json:"customAttribute8,omitempty"`
	CustomAttribute9                   string    `json:"customAttribute9,omitempty"`
	CustomAttribute10                  string    `json:"customAttribute10,omitempty"`
	CustomAttribute11                  string    `json:"customAttribute11,omitempty"`
	CustomAttribute12                  string    `json:"customAttribute12,omitempty"`
	CustomAttribute13                  string    `json:"customAttribute13,omitempty"`
	CustomAttribute14                  string    `json:"customAttribute14,omitempty"`
	CustomAttribute15                  string    `json:"customAttribute15,omitempty"`
	ExtensionCustomAttribute1          []string  `json:"extensionCustomAttribute1,omitempty"`
	ExtensionCustomAttribute2          []string  `json:"extensionCustomAttribute2,omitempty"`
	ExtensionCustomAttribute3          []string  `json:"extensionCustomAttribute3,omitempty"`
	ExtensionCustomAttribute4          []string  `json:"extensionCustomAttribute4,omitempty"`
	ExtensionCustomAttribute5          []string  `json:"extensionCustomAttribute5,omitempty"`
	Database                           string    `json:"database,omitempty"`
	ArchiveDatabase                    string    `json:"archiveDatabase,omitempty"`
	DatabaseName                       string    `json:"databaseName,omitempty"`
	Department                         string    `json:"department,omitempty"`
	ManagedFolderMailboxPolicy         string    `json:"managedFolderMailboxPolicy,omitempty"`
	ExpansionServer                    string    `json:"expansionServer,omitempty"`
	ExternalEmailAddress               string    `json:"externalEmailAddress,omitempty"`
	HiddenFromAddressListsEnabled      bool      `json:"hiddenFromAddressListsEnabled,omitempty"`
	EmailAddressPolicyEnabled          bool      `json:"emailAddressPolicyEnabled,omitempty"`
	ResourceType                       string    `json:"resourceType,omitempty"`
	ManagedBy                          []string  `json:"managedBy,omitempty"`
	Manager                            string    `json:"manager,omitempty"`
	ActiveSyncMailboxPolicy            string    `json:"activeSyncMailboxPolicy,omitempty"`
	ActiveSyncMailboxPolicyIsDefaulted bool      `json:"activeSyncMailboxPolicyIsDefaulted,omitempty"`
	Office                             string    `json:"office,omitempty"`
	ObjectCategory                     string    `json:"objectCategory,omitempty"`
	OrganizationalUnit                 string    `json:"organizationalUnit,omitempty"`
	Phone                              string    `json:"phone,omitempty"`
	PoliciesIncluded                   []string  `json:"policiesIncluded,omitempty"`
	PoliciesExcluded                   []string  `json:"policiesExcluded,omitempty"`
	PrimarySmtpAddress                 string    `json:"primarySmtpAddress,omitempty"`
	RecipientType                      string    `json:"recipientType,omitempty"`
	RecipientTypeDetails               string    `json:"recipientTypeDetails,omitempty"`
	SamAccountName                     string    `json:"samAccountName,omitempty"`
	ServerLegacyDN                     string    `json:"serverLegacyDN,omitempty"`
	ServerName                         string    `json:"serverName,omitempty"`
	StateOrProvince                    string    `json:"stateOrProvince,omitempty"`
	StorageGroupName                   string    `json:"storageGroupName,omitempty"`
	Title                              string    `json:"title,omitempty"`
	UMMailboxPolicy                    string    `json:"uMMailboxPolicy,omitempty"`
	UMRecipientDialPlanId              string    `json:"uMRecipientDialPlanId,omitempty"`
	WindowsLiveID                      string    `json:"windowsLiveID,omitempty"`
	HasActiveSyncDevicePartnership     bool      `json:"hasActiveSyncDevicePartnership,omitempty"`
	AddressListMembership              []string  `json:"addressListMembership,omitempty"`
	OwaMailboxPolicy                   string    `json:"owaMailboxPolicy,omitempty"`
	AddressBookPolicy                  string    `json:"addressBookPolicy,omitempty"`
	SharingPolicy                      string    `json:"sharingPolicy,omitempty"`
	RetentionPolicy                    string    `json:"retentionPolicy,omitempty"`
	ShouldUseDefaultRetentionPolicy    bool      `json:"shouldUseDefaultRetentionPolicy,omitempty"`
	MailboxMoveTargetMDB               string    `json:"mailboxMoveTargetMDB,omitempty"`
	MailboxMoveSourceMDB               string    `json:"mailboxMoveSourceMDB,omitempty"`
	MailboxMoveFlags                   string    `json:"mailboxMoveFlags,omitempty"`
	MailboxMoveRemoteHostName          string    `json:"mailboxMoveRemoteHostName,omitempty"`
	MailboxMoveBatchName               string    `json:"mailboxMoveBatchName,omitempty"`
	MailboxMoveStatus                  string    `json:"mailboxMoveStatus,omitempty"`
	MailboxRelease                     string    `json:"mailboxRelease,omitempty"`
	ArchiveRelease                     string    `json:"archiveRelease,omitempty"`
	IsValidSecurityPrincipal           bool      `json:"isValidSecurityPrincipal,omitempty"`
	LitigationHoldEnabled              bool      `json:"litigationHoldEnabled,omitempty"`
	Capabilities                       []string  `json:"capabilities,omitempty"`
	ArchiveState                       string    `json:"archiveState,omitempty"`
	SKUAssigned                        bool      `json:"sKUAssigned,omitempty"`
	WhenMailboxCreated                 string    `json:"whenMailboxCreated,omitempty"`
	UsageLocation                      string    `json:"usageLocation,omitempty"`
	ExchangeGuid                       string    `json:"exchangeGuid,omitempty"`
	ArchiveStatus                      string    `json:"archiveStatus,omitempty"`
	SafeSendersHash                    []byte    `json:"safeSendersHash,omitempty"`
	SafeRecipientsHash                 []byte    `json:"safeRecipientsHash,omitempty"`
	BlockedSendersHash                 []byte    `json:"blockedSendersHash,omitempty"`
	WhenSoftDeleted                    string    `json:"whenSoftDeleted,omitempty"`
	Version                            string    `json:"version,omitempty"`
	DistinguishedName                  string    `json:"distinguishedName,omitempty"`
	ObjectClass                        []string  `json:"objectClass,omitempty"`
	WhenChangedDateTime                time.Time `json:"whenChangedDateTime,omitempty"`
	WhenCreatedDateTime                time.Time `json:"whenCreatedDateTime,omitempty"`
	WhenChangedUTCDateTime             time.Time `json:"whenChangedUTCDateTime,omitempty"`
	WhenCreatedUTCDateTime             time.Time `json:"whenCreatedUTCDateTime,omitempty"`
	ExchangeObjectId                   string    `json:"exchangeObjectId,omitempty"`
	OrganizationId                     string    `json:"organizationId,omitempty"`
	Id                                 string    `json:"id,omitempty"`
	Guid                               string    `json:"guid,omitempty"`
}

// ExoExchangeSecurityDescriptor (OData ComplexType).
type ExoExchangeSecurityDescriptor struct {
	ComplexEntry
	ControlFlags     string   `json:"ControlFlags,omitempty"`
	Owner            string   `json:"Owner,omitempty"`
	Group            string   `json:"Group,omitempty"`
	SystemAcl        []string `json:"SystemAcl,omitempty"`
	DiscretionaryAcl []string `json:"DiscretionaryAcl,omitempty"`
	BinaryLength     int32    `json:"BinaryLength,omitempty"`
}

// GenericHashTable (OData ComplexType).
type GenericHashTable struct {
	OpenComplexEntry
}

// GraphConnectorGroup (OData EntityType).
type GraphConnectorGroup struct {
	ExchangeObjectId  string                       `json:"exchangeObjectId,omitempty"`
	Guid              string                       `json:"guid,omitempty"`
	Id                string                       `json:"id,omitempty"`
	Name              string                       `json:"name,omitempty"`
	DisplayName       string                       `json:"displayName,omitempty"`
	Description       string                       `json:"description,omitempty"`
	CustomAttribute14 string                       `json:"customAttribute14,omitempty"`
	CustomAttribute15 string                       `json:"customAttribute15,omitempty"`
	DistinguishedName string                       `json:"distinguishedName,omitempty"`
	WhenChanged       time.Time                    `json:"whenChanged,omitempty"`
	WhenCreated       time.Time                    `json:"whenCreated,omitempty"`
	WhenChangedUTC    time.Time                    `json:"whenChangedUTC,omitempty"`
	WhenCreatedUTC    time.Time                    `json:"whenCreatedUTC,omitempty"`
	OrganizationId    string                       `json:"organizationId,omitempty"`
	OriginatingServer string                       `json:"originatingServer,omitempty"`
	Members           []*GraphConnectorGroupMember `json:"members,omitempty"`
}

// GraphConnectorGroupMember (OData ComplexType).
type GraphConnectorGroupMember struct {
	Id                        string `json:"id,omitempty"`
	Name                      string `json:"name,omitempty"`
	DisplayName               string `json:"displayName,omitempty"`
	ExternalDirectoryObjectId string `json:"externalDirectoryObjectId,omitempty"`
	ExchangeObjectId          string `json:"exchangeObjectId,omitempty"`
	Guid                      string `json:"guid,omitempty"`
	RecipientType             string `json:"recipientType,omitempty"`
	RecipientTypeDetails      string `json:"recipientTypeDetails,omitempty"`
	CustomAttribute14         string `json:"customAttribute14,omitempty"`
	CustomAttribute15         string `json:"customAttribute15,omitempty"`
}

// HistoricalSearch (OData ComplexType).
type HistoricalSearch struct {
	CompressFile            bool      `json:"CompressFile,omitempty"`
	NumOfBlocks             int32     `json:"NumOfBlocks,omitempty"`
	IsSaved                 bool      `json:"IsSaved,omitempty"`
	EncryptionTemplate      string    `json:"EncryptionTemplate,omitempty"`
	EncryptionType          string    `json:"EncryptionType,omitempty"`
	JobProgress             string    `json:"JobProgress,omitempty"`
	EstimatedCompletionTime time.Time `json:"EstimatedCompletionTime,omitempty"`
	CompletionDate          time.Time `json:"CompletionDate,omitempty"`
	Direction               string    `json:"Direction,omitempty"`
	DLPPolicy               []string  `json:"DLPPolicy,omitempty"`
	TransportRule           []string  `json:"TransportRule,omitempty"`
	MessageID               []string  `json:"MessageID,omitempty"`
	OriginalClientIP        string    `json:"OriginalClientIP,omitempty"`
	RecipientAddress        []string  `json:"RecipientAddress,omitempty"`
	SenderAddress           []string  `json:"SenderAddress,omitempty"`
	DeliveryStatus          string    `json:"DeliveryStatus,omitempty"`
	EndDate                 time.Time `json:"EndDate,omitempty"`
	StartDate               time.Time `json:"StartDate,omitempty"`
	NotifyAddress           []string  `json:"NotifyAddress,omitempty"`
	ReportType              string    `json:"ReportType,omitempty"`
	ReportStatusDescription string    `json:"ReportStatusDescription,omitempty"`
	FileUrl                 string    `json:"FileUrl,omitempty"`
	ErrorDescription        string    `json:"ErrorDescription,omitempty"`
	ErrorCode               string    `json:"ErrorCode,omitempty"`
	FileRows                int32     `json:"FileRows,omitempty"`
	Rows                    int32     `json:"Rows,omitempty"`
	Status                  string    `json:"Status,omitempty"`
	BlockStatus             string    `json:"BlockStatus,omitempty"`
	SmtpSecurityError       string    `json:"SmtpSecurityError,omitempty"`
	TLSUsed                 string    `json:"TLSUsed,omitempty"`
	ConnectorType           string    `json:"ConnectorType,omitempty"`
	ReportTitle             string    `json:"ReportTitle,omitempty"`
	SubmitDate              time.Time `json:"SubmitDate,omitempty"`
	Identity                string    `json:"Identity,omitempty"`
	JobId                   string    `json:"JobId,omitempty"`
}

// HostedContentFilterPolicyPresentation (OData ComplexType).
type HostedContentFilterPolicyPresentation struct {
	Identity                             string                      `json:"Identity,omitempty"`
	AdminDisplayName                     string                      `json:"AdminDisplayName,omitempty"`
	AddXHeaderValue                      string                      `json:"AddXHeaderValue,omitempty"`
	ModifySubjectValue                   string                      `json:"ModifySubjectValue,omitempty"`
	RedirectToRecipients                 []string                    `json:"RedirectToRecipients,omitempty"`
	TestModeBccToRecipients              []string                    `json:"TestModeBccToRecipients,omitempty"`
	FalsePositiveAdditionalRecipients    []string                    `json:"FalsePositiveAdditionalRecipients,omitempty"`
	QuarantineRetentionPeriod            int32                       `json:"QuarantineRetentionPeriod,omitempty"`
	TestModeAction                       SpamFilteringTestModeAction `json:"TestModeAction,omitempty"`
	IncreaseScoreWithImageLinks          SpamFilteringOption         `json:"IncreaseScoreWithImageLinks,omitempty"`
	IncreaseScoreWithNumericIps          SpamFilteringOption         `json:"IncreaseScoreWithNumericIps,omitempty"`
	IncreaseScoreWithRedirectToOtherPort SpamFilteringOption         `json:"IncreaseScoreWithRedirectToOtherPort,omitempty"`
	IncreaseScoreWithBizOrInfoUrls       SpamFilteringOption         `json:"IncreaseScoreWithBizOrInfoUrls,omitempty"`
	MarkAsSpamEmptyMessages              SpamFilteringOption         `json:"MarkAsSpamEmptyMessages,omitempty"`
	MarkAsSpamJavaScriptInHtml           SpamFilteringOption         `json:"MarkAsSpamJavaScriptInHtml,omitempty"`
	MarkAsSpamFramesInHtml               SpamFilteringOption         `json:"MarkAsSpamFramesInHtml,omitempty"`
	MarkAsSpamObjectTagsInHtml           SpamFilteringOption         `json:"MarkAsSpamObjectTagsInHtml,omitempty"`
	MarkAsSpamEmbedTagsInHtml            SpamFilteringOption         `json:"MarkAsSpamEmbedTagsInHtml,omitempty"`
	MarkAsSpamFormTagsInHtml             SpamFilteringOption         `json:"MarkAsSpamFormTagsInHtml,omitempty"`
	MarkAsSpamWebBugsInHtml              SpamFilteringOption         `json:"MarkAsSpamWebBugsInHtml,omitempty"`
	MarkAsSpamSensitiveWordList          SpamFilteringOption         `json:"MarkAsSpamSensitiveWordList,omitempty"`
	MarkAsSpamSpfRecordHardFail          SpamFilteringOption         `json:"MarkAsSpamSpfRecordHardFail,omitempty"`
	MarkAsSpamFromAddressAuthFail        SpamFilteringOption         `json:"MarkAsSpamFromAddressAuthFail,omitempty"`
	MarkAsSpamBulkMail                   SpamFilteringOption         `json:"MarkAsSpamBulkMail,omitempty"`
	MarkAsSpamNdrBackscatter             SpamFilteringOption         `json:"MarkAsSpamNdrBackscatter,omitempty"`
	IsDefault                            bool                        `json:"IsDefault,omitempty"`
	LanguageBlockList                    []string                    `json:"LanguageBlockList,omitempty"`
	RegionBlockList                      []string                    `json:"RegionBlockList,omitempty"`
	HighConfidenceSpamAction             SpamFilteringAction         `json:"HighConfidenceSpamAction,omitempty"`
	SpamAction                           SpamFilteringAction         `json:"SpamAction,omitempty"`
	DownloadLink                         bool                        `json:"DownloadLink,omitempty"`
	EnableRegionBlockList                bool                        `json:"EnableRegionBlockList,omitempty"`
	EnableLanguageBlockList              bool                        `json:"EnableLanguageBlockList,omitempty"`
	BulkThreshold                        int32                       `json:"BulkThreshold,omitempty"`
	ZapEnabled                           bool                        `json:"ZapEnabled,omitempty"`
	InlineSafetyTipsEnabled              bool                        `json:"InlineSafetyTipsEnabled,omitempty"`
	BulkMovesEnabled                     BulkMovesEnabled            `json:"BulkMovesEnabled,omitempty"`
	BulkSpamAction                       SpamFilteringAction         `json:"BulkSpamAction,omitempty"`
	PhishSpamAction                      SpamFilteringAction         `json:"PhishSpamAction,omitempty"`
	SpamZapEnabled                       bool                        `json:"SpamZapEnabled,omitempty"`
	PhishZapEnabled                      bool                        `json:"PhishZapEnabled,omitempty"`
	IntraOrgFilterState                  IntraOrgFilterState         `json:"IntraOrgFilterState,omitempty"`
	HighConfidencePhishAction            PhishFilteringAction        `json:"HighConfidencePhishAction,omitempty"`
	RecommendedPolicyType                RecommendedPolicyType       `json:"RecommendedPolicyType,omitempty"`
	Name                                 string                      `json:"Name,omitempty"`
	WhenChanged                          time.Time                   `json:"WhenChanged,omitempty"`
	WhenCreated                          time.Time                   `json:"WhenCreated,omitempty"`
	ExchangeObjectId                     string                      `json:"ExchangeObjectId,omitempty"`
	OrganizationId                       string                      `json:"OrganizationId,omitempty"`
	Guid                                 string                      `json:"Guid,omitempty"`
}

// HostedContentFilterRule (OData ComplexType).
type HostedContentFilterRule struct {
	Identity       string    `json:"Identity,omitempty"`
	State          string    `json:"State,omitempty"`
	Priority       int32     `json:"Priority,omitempty"`
	Comments       string    `json:"Comments,omitempty"`
	Guid           string    `json:"Guid,omitempty"`
	ImmutableId    string    `json:"ImmutableId,omitempty"`
	OrganizationId string    `json:"OrganizationId,omitempty"`
	Name           string    `json:"Name,omitempty"`
	WhenChanged    time.Time `json:"WhenChanged,omitempty"`
}

// InboundConnector (OData ComplexType).
type InboundConnector struct {
	ConnectorType                string   `json:"ConnectorType,omitempty"`
	ConnectorSource              string   `json:"ConnectorSource,omitempty"`
	Enabled                      bool     `json:"Enabled,omitempty"`
	Comment                      string   `json:"Comment,omitempty"`
	SenderIPAddresses            []string `json:"SenderIPAddresses,omitempty"`
	SenderDomains                []string `json:"SenderDomains,omitempty"`
	AssociatedAcceptedDomains    []string `json:"AssociatedAcceptedDomains,omitempty"`
	RequireTls                   bool     `json:"RequireTls,omitempty"`
	RestrictDomainsToIPAddresses bool     `json:"RestrictDomainsToIPAddresses,omitempty"`
	RestrictDomainsToCertificate bool     `json:"RestrictDomainsToCertificate,omitempty"`
	CloudServicesMailEnabled     bool     `json:"CloudServicesMailEnabled,omitempty"`
	TreatMessagesAsInternal      bool     `json:"TreatMessagesAsInternal,omitempty"`
	TlsSenderCertificateName     string   `json:"TlsSenderCertificateName,omitempty"`
	EFTestMode                   bool     `json:"EFTestMode,omitempty"`
	ScanAndDropRecipients        []string `json:"ScanAndDropRecipients,omitempty"`
	EFSkipLastIP                 bool     `json:"EFSkipLastIP,omitempty"`
	EFSkipIPs                    []string `json:"EFSkipIPs,omitempty"`
	EFSkipMailGateway            []string `json:"EFSkipMailGateway,omitempty"`
	EFUsers                      []string `json:"EFUsers,omitempty"`
	Name                         string   `json:"Name,omitempty"`
	Identity                     string   `json:"Identity,omitempty"`
	Guid                         string   `json:"Guid,omitempty"`
}

// MailDetailDlpPolicyReport (OData ComplexType).
type MailDetailDlpPolicyReport struct {
	ComplexEntry
	EventType                      string    `json:"EventType,omitempty"`
	SensitiveInformationConfidence int32     `json:"SensitiveInformationConfidence,omitempty"`
	SensitiveInformationCount      int32     `json:"SensitiveInformationCount,omitempty"`
	SensitiveInformationType       string    `json:"SensitiveInformationType,omitempty"`
	Justification                  string    `json:"Justification,omitempty"`
	UserAction                     string    `json:"UserAction,omitempty"`
	TransportRule                  string    `json:"TransportRule,omitempty"`
	DlpPolicy                      string    `json:"DlpPolicy,omitempty"`
	RecipientAddress               string    `json:"RecipientAddress,omitempty"`
	SenderAddress                  string    `json:"SenderAddress,omitempty"`
	Direction                      string    `json:"Direction,omitempty"`
	MessageSize                    int32     `json:"MessageSize,omitempty"`
	Subject                        string    `json:"Subject,omitempty"`
	Domain                         string    `json:"Domain,omitempty"`
	MessageId                      string    `json:"MessageId,omitempty"`
	Date                           time.Time `json:"Date,omitempty"`
	Organization                   string    `json:"Organization,omitempty"`
	Action                         string    `json:"Action,omitempty"`
	MessageTraceId                 string    `json:"MessageTraceId,omitempty"`
}

// MailDetailTransportRuleReport (OData ComplexType).
type MailDetailTransportRuleReport struct {
	ComplexEntry
	Action           string    `json:"Action,omitempty"`
	Date             time.Time `json:"Date,omitempty"`
	Direction        string    `json:"Direction,omitempty"`
	Domain           string    `json:"Domain,omitempty"`
	EventType        string    `json:"EventType,omitempty"`
	MessageId        string    `json:"MessageId,omitempty"`
	MessageSize      int32     `json:"MessageSize,omitempty"`
	Organization     string    `json:"Organization,omitempty"`
	RecipientAddress string    `json:"RecipientAddress,omitempty"`
	SenderAddress    string    `json:"SenderAddress,omitempty"`
	Subject          string    `json:"Subject,omitempty"`
	TransportRule    string    `json:"TransportRule,omitempty"`
	MessageTraceId   string    `json:"MessageTraceId,omitempty"`
}

// MailTrafficPolicyReport (OData ComplexType).
type MailTrafficPolicyReport struct {
	ComplexEntry
	Organization  string    `json:"Organization,omitempty"`
	Domain        string    `json:"Domain,omitempty"`
	Date          time.Time `json:"Date,omitempty"`
	DlpPolicy     string    `json:"DlpPolicy,omitempty"`
	TransportRule string    `json:"TransportRule,omitempty"`
	Action        string    `json:"Action,omitempty"`
	EventType     string    `json:"EventType,omitempty"`
	Direction     string    `json:"Direction,omitempty"`
	MessageCount  int32     `json:"MessageCount,omitempty"`
	SummarizeBy   string    `json:"SummarizeBy,omitempty"`
}

// Mailbox (OData EntityType).
type Mailbox struct {
	ObjectKey                                 string                         `json:"ObjectKey,omitempty"`
	ExternalDirectoryObjectId                 string                         `json:"ExternalDirectoryObjectId,omitempty"`
	Database                                  string                         `json:"Database,omitempty"`
	DatabaseGuid                              string                         `json:"DatabaseGuid,omitempty"`
	MailboxProvisioningConstraint             string                         `json:"MailboxProvisioningConstraint,omitempty"`
	IsMonitoringMailbox                       bool                           `json:"IsMonitoringMailbox,omitempty"`
	MailboxRegion                             string                         `json:"MailboxRegion,omitempty"`
	MailboxRegionLastUpdateTime               string                         `json:"MailboxRegionLastUpdateTime,omitempty"`
	MessageRecallProcessingEnabled            bool                           `json:"MessageRecallProcessingEnabled,omitempty"`
	MessageCopyForSMTPClientSubmissionEnabled bool                           `json:"MessageCopyForSMTPClientSubmissionEnabled,omitempty"`
	MessageCopyForSentAsEnabled               bool                           `json:"MessageCopyForSentAsEnabled,omitempty"`
	MessageCopyForSendOnBehalfEnabled         bool                           `json:"MessageCopyForSendOnBehalfEnabled,omitempty"`
	MailboxProvisioningPreferences            []string                       `json:"MailboxProvisioningPreferences,omitempty"`
	UseDatabaseRetentionDefaults              bool                           `json:"UseDatabaseRetentionDefaults,omitempty"`
	RetainDeletedItemsUntilBackup             bool                           `json:"RetainDeletedItemsUntilBackup,omitempty"`
	DeliverToMailboxAndForward                bool                           `json:"DeliverToMailboxAndForward,omitempty"`
	IsExcludedFromServingHierarchy            bool                           `json:"IsExcludedFromServingHierarchy,omitempty"`
	IsHierarchyReady                          bool                           `json:"IsHierarchyReady,omitempty"`
	IsHierarchySyncEnabled                    bool                           `json:"IsHierarchySyncEnabled,omitempty"`
	IsPublicFolderSystemMailbox               bool                           `json:"IsPublicFolderSystemMailbox,omitempty"`
	HasSnackyAppData                          bool                           `json:"HasSnackyAppData,omitempty"`
	LitigationHoldEnabled                     bool                           `json:"LitigationHoldEnabled,omitempty"`
	SingleItemRecoveryEnabled                 bool                           `json:"SingleItemRecoveryEnabled,omitempty"`
	RetentionHoldEnabled                      bool                           `json:"RetentionHoldEnabled,omitempty"`
	EndDateForRetentionHold                   string                         `json:"EndDateForRetentionHold,omitempty"`
	StartDateForRetentionHold                 string                         `json:"StartDateForRetentionHold,omitempty"`
	RetentionComment                          string                         `json:"RetentionComment,omitempty"`
	RetentionUrl                              string                         `json:"RetentionUrl,omitempty"`
	LitigationHoldDate                        string                         `json:"LitigationHoldDate,omitempty"`
	LitigationHoldOwner                       string                         `json:"LitigationHoldOwner,omitempty"`
	ElcProcessingDisabled                     bool                           `json:"ElcProcessingDisabled,omitempty"`
	ComplianceTagHoldApplied                  bool                           `json:"ComplianceTagHoldApplied,omitempty"`
	WasInactiveMailbox                        bool                           `json:"WasInactiveMailbox,omitempty"`
	DelayHoldApplied                          bool                           `json:"DelayHoldApplied,omitempty"`
	InactiveMailboxRetireTime                 string                         `json:"InactiveMailboxRetireTime,omitempty"`
	OrphanSoftDeleteTrackingTime              string                         `json:"OrphanSoftDeleteTrackingTime,omitempty"`
	LitigationHoldDuration                    string                         `json:"LitigationHoldDuration,omitempty"`
	PitrEnabled                               bool                           `json:"PitrEnabled,omitempty"`
	PitrCopyIntervalInSeconds                 int32                          `json:"PitrCopyIntervalInSeconds,omitempty"`
	PitrPaused                                bool                           `json:"PitrPaused,omitempty"`
	PitrPausedTimestamp                       string                         `json:"PitrPausedTimestamp,omitempty"`
	PitrOffboardedTimestamp                   string                         `json:"PitrOffboardedTimestamp,omitempty"`
	PitrState                                 string                         `json:"PitrState,omitempty"`
	M365BackupState                           string                         `json:"M365BackupState,omitempty"`
	ManagedFolderMailboxPolicy                string                         `json:"ManagedFolderMailboxPolicy,omitempty"`
	RetentionPolicy                           string                         `json:"RetentionPolicy,omitempty"`
	AddressBookPolicy                         string                         `json:"AddressBookPolicy,omitempty"`
	CalendarRepairDisabled                    bool                           `json:"CalendarRepairDisabled,omitempty"`
	ExchangeGuid                              string                         `json:"ExchangeGuid,omitempty"`
	MailboxContainerGuid                      string                         `json:"MailboxContainerGuid,omitempty"`
	UnifiedMailbox                            string                         `json:"UnifiedMailbox,omitempty"`
	MailboxLocations                          []string                       `json:"MailboxLocations,omitempty"`
	AggregatedMailboxGuids                    []string                       `json:"AggregatedMailboxGuids,omitempty"`
	ExchangeSecurityDescriptor                string                         `json:"ExchangeSecurityDescriptor,omitempty"`
	ExoExchangeSecurityDescriptor             *ExoExchangeSecurityDescriptor `json:"ExoExchangeSecurityDescriptor,omitempty"`
	ExchangeUserAccountControl                string                         `json:"ExchangeUserAccountControl,omitempty"`
	MessageTrackingReadStatusEnabled          bool                           `json:"MessageTrackingReadStatusEnabled,omitempty"`
	ExternalOofOptions                        string                         `json:"ExternalOofOptions,omitempty"`
	ForwardingAddress                         string                         `json:"ForwardingAddress,omitempty"`
	ForwardingAddressWithDisplayNames         []string                       `json:"ForwardingAddressWithDisplayNames,omitempty"`
	ForwardingSmtpAddress                     string                         `json:"ForwardingSmtpAddress,omitempty"`
	RetainDeletedItemsFor                     string                         `json:"RetainDeletedItemsFor,omitempty"`
	IsMailboxEnabled                          bool                           `json:"IsMailboxEnabled,omitempty"`
	Languages                                 []string                       `json:"Languages,omitempty"`
	OfflineAddressBook                        string                         `json:"OfflineAddressBook,omitempty"`
	ProhibitSendQuota                         string                         `json:"ProhibitSendQuota,omitempty"`
	ProhibitSendReceiveQuota                  string                         `json:"ProhibitSendReceiveQuota,omitempty"`
	RecoverableItemsQuota                     string                         `json:"RecoverableItemsQuota,omitempty"`
	RecoverableItemsWarningQuota              string                         `json:"RecoverableItemsWarningQuota,omitempty"`
	CalendarLoggingQuota                      string                         `json:"CalendarLoggingQuota,omitempty"`
	DowngradeHighPriorityMessagesEnabled      string                         `json:"DowngradeHighPriorityMessagesEnabled,omitempty"`
	ProtocolSettings                          []string                       `json:"ProtocolSettings,omitempty"`
	RecipientLimits                           string                         `json:"RecipientLimits,omitempty"`
	ImListMigrationCompleted                  bool                           `json:"ImListMigrationCompleted,omitempty"`
	SiloName                                  string                         `json:"SiloName,omitempty"`
	IsResource                                bool                           `json:"IsResource,omitempty"`
	IsLinked                                  bool                           `json:"IsLinked,omitempty"`
	IsShared                                  bool                           `json:"IsShared,omitempty"`
	IsRootPublicFolderMailbox                 bool                           `json:"IsRootPublicFolderMailbox,omitempty"`
	LinkedMasterAccount                       string                         `json:"LinkedMasterAccount,omitempty"`
	ResetPasswordOnNextLogon                  string                         `json:"ResetPasswordOnNextLogon,omitempty"`
	ResourceCapacity                          int32                          `json:"ResourceCapacity,omitempty"`
	ResourceCustom                            []string                       `json:"ResourceCustom,omitempty"`
	ResourceType                              string                         `json:"ResourceType,omitempty"`
	RoomMailboxAccountEnabled                 bool                           `json:"RoomMailboxAccountEnabled,omitempty"`
	SamAccountName                            string                         `json:"SamAccountName,omitempty"`
	SCLDeleteThreshold                        int32                          `json:"SCLDeleteThreshold,omitempty"`
	SCLDeleteEnabled                          bool                           `json:"SCLDeleteEnabled,omitempty"`
	SCLRejectThreshold                        int32                          `json:"SCLRejectThreshold,omitempty"`
	SCLRejectEnabled                          bool                           `json:"SCLRejectEnabled,omitempty"`
	SCLQuarantineThreshold                    int32                          `json:"SCLQuarantineThreshold,omitempty"`
	SCLQuarantineEnabled                      bool                           `json:"SCLQuarantineEnabled,omitempty"`
	SCLJunkThreshold                          int32                          `json:"SCLJunkThreshold,omitempty"`
	SCLJunkEnabled                            bool                           `json:"SCLJunkEnabled,omitempty"`
	AntispamBypassEnabled                     bool                           `json:"AntispamBypassEnabled,omitempty"`
	ServerLegacyDN                            string                         `json:"ServerLegacyDN,omitempty"`
	UseDatabaseQuotaDefaults                  bool                           `json:"UseDatabaseQuotaDefaults,omitempty"`
	IssueWarningQuota                         string                         `json:"IssueWarningQuota,omitempty"`
	RulesQuota                                string                         `json:"RulesQuota,omitempty"`
	Office                                    string                         `json:"Office,omitempty"`
	UserPrincipalName                         string                         `json:"UserPrincipalName,omitempty"`
	MaxSafeSenders                            int32                          `json:"MaxSafeSenders,omitempty"`
	MaxBlockedSenders                         int32                          `json:"MaxBlockedSenders,omitempty"`
	NetID                                     string                         `json:"NetID,omitempty"`
	ReconciliationId                          string                         `json:"ReconciliationId,omitempty"`
	WindowsLiveID                             string                         `json:"WindowsLiveID,omitempty"`
	MicrosoftOnlineServicesID                 string                         `json:"MicrosoftOnlineServicesID,omitempty"`
	ThrottlingPolicy                          string                         `json:"ThrottlingPolicy,omitempty"`
	RoleAssignmentPolicy                      string                         `json:"RoleAssignmentPolicy,omitempty"`
	DefaultPublicFolderMailbox                string                         `json:"DefaultPublicFolderMailbox,omitempty"`
	EffectivePublicFolderMailbox              string                         `json:"EffectivePublicFolderMailbox,omitempty"`
	SharingPolicy                             string                         `json:"SharingPolicy,omitempty"`
	RemoteAccountPolicy                       string                         `json:"RemoteAccountPolicy,omitempty"`
	MailboxPlan                               string                         `json:"MailboxPlan,omitempty"`
	ArchiveDatabase                           string                         `json:"ArchiveDatabase,omitempty"`
	ArchiveDatabaseGuid                       string                         `json:"ArchiveDatabaseGuid,omitempty"`
	ArchiveGuid                               string                         `json:"ArchiveGuid,omitempty"`
	ArchiveName                               []string                       `json:"ArchiveName,omitempty"`
	JournalArchiveAddress                     string                         `json:"JournalArchiveAddress,omitempty"`
	ArchiveQuota                              string                         `json:"ArchiveQuota,omitempty"`
	ArchiveWarningQuota                       string                         `json:"ArchiveWarningQuota,omitempty"`
	ArchiveDomain                             string                         `json:"ArchiveDomain,omitempty"`
	ArchiveStatus                             string                         `json:"ArchiveStatus,omitempty"`
	ArchiveState                              string                         `json:"ArchiveState,omitempty"`
	AutoExpandingArchiveEnabled               bool                           `json:"AutoExpandingArchiveEnabled,omitempty"`
	AutoArchivingEnabled                      bool                           `json:"AutoArchivingEnabled,omitempty"`
	DisabledMailboxLocations                  bool                           `json:"DisabledMailboxLocations,omitempty"`
	RemoteRecipientType                       string                         `json:"RemoteRecipientType,omitempty"`
	DisabledArchiveDatabase                   string                         `json:"DisabledArchiveDatabase,omitempty"`
	DisabledArchiveGuid                       string                         `json:"DisabledArchiveGuid,omitempty"`
	QueryBaseDN                               string                         `json:"QueryBaseDN,omitempty"`
	QueryBaseDNRestrictionEnabled             bool                           `json:"QueryBaseDNRestrictionEnabled,omitempty"`
	MailboxMoveTargetMDB                      string                         `json:"MailboxMoveTargetMDB,omitempty"`
	MailboxMoveSourceMDB                      string                         `json:"MailboxMoveSourceMDB,omitempty"`
	MailboxMoveFlags                          string                         `json:"MailboxMoveFlags,omitempty"`
	MailboxMoveRemoteHostName                 string                         `json:"MailboxMoveRemoteHostName,omitempty"`
	MailboxMoveBatchName                      string                         `json:"MailboxMoveBatchName,omitempty"`
	MailboxMoveStatus                         string                         `json:"MailboxMoveStatus,omitempty"`
	MailboxRelease                            string                         `json:"MailboxRelease,omitempty"`
	ArchiveRelease                            string                         `json:"ArchiveRelease,omitempty"`
	IsPersonToPersonTextMessagingEnabled      string                         `json:"IsPersonToPersonTextMessagingEnabled,omitempty"`
	IsMachineToPersonTextMessagingEnabled     string                         `json:"IsMachineToPersonTextMessagingEnabled,omitempty"`
	UserSMimeCertificate                      []*ByteArrayType               `json:"UserSMimeCertificate,omitempty"`
	UserCertificate                           []*ByteArrayType               `json:"UserCertificate,omitempty"`
	CalendarVersionStoreDisabled              bool                           `json:"CalendarVersionStoreDisabled,omitempty"`
	ImmutableId                               string                         `json:"ImmutableId,omitempty"`
	PersistedCapabilities                     []string                       `json:"PersistedCapabilities,omitempty"`
	SKUAssigned                               bool                           `json:"SKUAssigned,omitempty"`
	AuditEnabled                              bool                           `json:"AuditEnabled,omitempty"`
	AuditLogAgeLimit                          string                         `json:"AuditLogAgeLimit,omitempty"`
	AuditAdmin                                []string                       `json:"AuditAdmin,omitempty"`
	AuditDelegate                             []string                       `json:"AuditDelegate,omitempty"`
	AuditOwner                                []string                       `json:"AuditOwner,omitempty"`
	DefaultAuditSet                           []string                       `json:"DefaultAuditSet,omitempty"`
	WhenMailboxCreated                        string                         `json:"WhenMailboxCreated,omitempty"`
	SourceAnchor                              string                         `json:"SourceAnchor,omitempty"`
	UsageLocation                             string                         `json:"UsageLocation,omitempty"`
	IsSoftDeletedByRemove                     bool                           `json:"IsSoftDeletedByRemove,omitempty"`
	IsSoftDeletedByDisable                    bool                           `json:"IsSoftDeletedByDisable,omitempty"`
	IsInactiveMailbox                         bool                           `json:"IsInactiveMailbox,omitempty"`
	IncludeInGarbageCollection                bool                           `json:"IncludeInGarbageCollection,omitempty"`
	WhenSoftDeleted                           string                         `json:"WhenSoftDeleted,omitempty"`
	InPlaceHolds                              []string                       `json:"InPlaceHolds,omitempty"`
	GeneratedOfflineAddressBooks              []string                       `json:"GeneratedOfflineAddressBooks,omitempty"`
	AccountDisabled                           bool                           `json:"AccountDisabled,omitempty"`
	StsRefreshTokensValidFrom                 string                         `json:"StsRefreshTokensValidFrom,omitempty"`
	NonCompliantDevices                       []string                       `json:"NonCompliantDevices,omitempty"`
	DataEncryptionPolicy                      string                         `json:"DataEncryptionPolicy,omitempty"`
	HasPicture                                bool                           `json:"HasPicture,omitempty"`
	HasSpokenName                             bool                           `json:"HasSpokenName,omitempty"`
	IsDirSynced                               bool                           `json:"IsDirSynced,omitempty"`
	AcceptMessagesOnlyFrom                    []string                       `json:"AcceptMessagesOnlyFrom,omitempty"`
	AcceptMessagesOnlyFromDLMembers           []string                       `json:"AcceptMessagesOnlyFromDLMembers,omitempty"`
	AcceptMessagesOnlyFromSendersOrMembers    []string                       `json:"AcceptMessagesOnlyFromSendersOrMembers,omitempty"`
	AddressListMembership                     []string                       `json:"AddressListMembership,omitempty"`
	AdministrativeUnits                       []string                       `json:"AdministrativeUnits,omitempty"`
	Alias                                     string                         `json:"Alias,omitempty"`
	ArbitrationMailbox                        string                         `json:"ArbitrationMailbox,omitempty"`
	BypassModerationFromSendersOrMembers      []string                       `json:"BypassModerationFromSendersOrMembers,omitempty"`
	OrganizationalUnit                        string                         `json:"OrganizationalUnit,omitempty"`
	CustomAttribute1                          string                         `json:"CustomAttribute1,omitempty"`
	CustomAttribute10                         string                         `json:"CustomAttribute10,omitempty"`
	CustomAttribute11                         string                         `json:"CustomAttribute11,omitempty"`
	CustomAttribute12                         string                         `json:"CustomAttribute12,omitempty"`
	CustomAttribute13                         string                         `json:"CustomAttribute13,omitempty"`
	CustomAttribute14                         string                         `json:"CustomAttribute14,omitempty"`
	CustomAttribute15                         string                         `json:"CustomAttribute15,omitempty"`
	CustomAttribute2                          string                         `json:"CustomAttribute2,omitempty"`
	CustomAttribute3                          string                         `json:"CustomAttribute3,omitempty"`
	CustomAttribute4                          string                         `json:"CustomAttribute4,omitempty"`
	CustomAttribute5                          string                         `json:"CustomAttribute5,omitempty"`
	CustomAttribute6                          string                         `json:"CustomAttribute6,omitempty"`
	CustomAttribute7                          string                         `json:"CustomAttribute7,omitempty"`
	CustomAttribute8                          string                         `json:"CustomAttribute8,omitempty"`
	CustomAttribute9                          string                         `json:"CustomAttribute9,omitempty"`
	ExtensionCustomAttribute1                 []string                       `json:"ExtensionCustomAttribute1,omitempty"`
	ExtensionCustomAttribute2                 []string                       `json:"ExtensionCustomAttribute2,omitempty"`
	ExtensionCustomAttribute3                 []string                       `json:"ExtensionCustomAttribute3,omitempty"`
	ExtensionCustomAttribute4                 []string                       `json:"ExtensionCustomAttribute4,omitempty"`
	ExtensionCustomAttribute5                 []string                       `json:"ExtensionCustomAttribute5,omitempty"`
	DisplayName                               string                         `json:"DisplayName,omitempty"`
	EmailAddresses                            []string                       `json:"EmailAddresses,omitempty"`
	GrantSendOnBehalfTo                       []string                       `json:"GrantSendOnBehalfTo,omitempty"`
	HiddenFromAddressListsEnabled             bool                           `json:"HiddenFromAddressListsEnabled,omitempty"`
	LastExchangeChangedTime                   string                         `json:"LastExchangeChangedTime,omitempty"`
	LegacyExchangeDN                          string                         `json:"LegacyExchangeDN,omitempty"`
	MaxSendSize                               string                         `json:"MaxSendSize,omitempty"`
	MaxReceiveSize                            string                         `json:"MaxReceiveSize,omitempty"`
	ModeratedBy                               []string                       `json:"ModeratedBy,omitempty"`
	ModerationEnabled                         bool                           `json:"ModerationEnabled,omitempty"`
	PoliciesIncluded                          []string                       `json:"PoliciesIncluded,omitempty"`
	PoliciesExcluded                          []string                       `json:"PoliciesExcluded,omitempty"`
	EmailAddressPolicyEnabled                 bool                           `json:"EmailAddressPolicyEnabled,omitempty"`
	PrimarySmtpAddress                        string                         `json:"PrimarySmtpAddress,omitempty"`
	RecipientType                             string                         `json:"RecipientType,omitempty"`
	RecipientTypeDetails                      string                         `json:"RecipientTypeDetails,omitempty"`
	RejectMessagesFrom                        []string                       `json:"RejectMessagesFrom,omitempty"`
	RejectMessagesFromDLMembers               []string                       `json:"RejectMessagesFromDLMembers,omitempty"`
	RejectMessagesFromSendersOrMembers        []string                       `json:"RejectMessagesFromSendersOrMembers,omitempty"`
	RequireSenderAuthenticationEnabled        bool                           `json:"RequireSenderAuthenticationEnabled,omitempty"`
	SimpleDisplayName                         string                         `json:"SimpleDisplayName,omitempty"`
	SendModerationNotifications               string                         `json:"SendModerationNotifications,omitempty"`
	UMDtmfMap                                 []string                       `json:"UMDtmfMap,omitempty"`
	WindowsEmailAddress                       string                         `json:"WindowsEmailAddress,omitempty"`
	MailTip                                   string                         `json:"MailTip,omitempty"`
	MailTipTranslations                       []string                       `json:"MailTipTranslations,omitempty"`
	Identity                                  string                         `json:"Identity,omitempty"`
	Id                                        string                         `json:"Id,omitempty"`
	ExchangeVersion                           string                         `json:"ExchangeVersion,omitempty"`
	Name                                      string                         `json:"Name,omitempty"`
	DistinguishedName                         string                         `json:"DistinguishedName,omitempty"`
	ObjectCategory                            string                         `json:"ObjectCategory,omitempty"`
	ObjectClass                               []string                       `json:"ObjectClass,omitempty"`
	WhenChanged                               string                         `json:"WhenChanged,omitempty"`
	WhenCreated                               string                         `json:"WhenCreated,omitempty"`
	WhenChangedUTC                            string                         `json:"WhenChangedUTC,omitempty"`
	WhenCreatedUTC                            string                         `json:"WhenCreatedUTC,omitempty"`
	ExchangeObjectId                          string                         `json:"ExchangeObjectId,omitempty"`
	OrganizationId                            string                         `json:"OrganizationId,omitempty"`
	Guid                                      string                         `json:"Guid,omitempty"`
	IsExchangeCloudManaged                    bool                           `json:"IsExchangeCloudManaged,omitempty"`
	ApplyMandatoryProperties                  bool                           `json:"ApplyMandatoryProperties,omitempty"`
	InactiveMailbox                           bool                           `json:"InactiveMailbox,omitempty"`
	PublicFolder                              bool                           `json:"PublicFolder,omitempty"`
	RecalculateInactiveMailbox                bool                           `json:"RecalculateInactiveMailbox,omitempty"`
	RemoveDelayHoldApplied                    bool                           `json:"RemoveDelayHoldApplied,omitempty"`
	RemoveDelayReleaseHoldApplied             bool                           `json:"RemoveDelayReleaseHoldApplied,omitempty"`
	RemoveDisabledArchive                     bool                           `json:"RemoveDisabledArchive,omitempty"`
	RemoveMailboxProvisioningConstraint       bool                           `json:"RemoveMailboxProvisioningConstraint,omitempty"`
	SecondaryAddress                          string                         `json:"SecondaryAddress,omitempty"`
	SecondaryDialPlan                         string                         `json:"SecondaryDialPlan,omitempty"`
	Type                                      string                         `json:"Type,omitempty"`
	UpdateEnforcedTimestamp                   bool                           `json:"UpdateEnforcedTimestamp,omitempty"`
	EnforcedTimestamps                        string                         `json:"EnforcedTimestamps,omitempty"`
	CreateDTMFMap                             bool                           `json:"CreateDTMFMap,omitempty"`
	EnableRoomMailboxAccount                  bool                           `json:"EnableRoomMailboxAccount,omitempty"`
	ExcludeFromAllOrgHolds                    bool                           `json:"ExcludeFromAllOrgHolds,omitempty"`
	ExcludeFromOrgHolds                       bool                           `json:"ExcludeFromOrgHolds,omitempty"`
	RemoveOrphanedHolds                       []string                       `json:"RemoveOrphanedHolds,omitempty"`
	DeltaUpdates                              *GenericHashTable              `json:"DeltaUpdates,omitempty"`
	MailboxPermission                         []*MailboxPermission           `json:"MailboxPermission,omitempty"`
	MailboxFolder                             []*MailboxFolder               `json:"MailboxFolder,omitempty"`
	MobileDevice                              []*MobileDevice                `json:"MobileDevice,omitempty"`
	MailboxRecoverableItem                    []*MailboxRecoverableItem      `json:"MailboxRecoverableItem,omitempty"`
	MailboxRegionalConfiguration              *MailboxRegionalConfiguration  `json:"MailboxRegionalConfiguration,omitempty"`
}

// MailboxAutoReplyConfiguration (OData ComplexType).
type MailboxAutoReplyConfiguration struct {
	Identity                         string           `json:"Identity,omitempty"`
	AutoDeclineFutureRequestsWhenOOF bool             `json:"AutoDeclineFutureRequestsWhenOOF,omitempty"`
	AutoReplyState                   OofState         `json:"AutoReplyState,omitempty"`
	CreateOOFEvent                   bool             `json:"CreateOOFEvent,omitempty"`
	DeclineAllEventsForScheduledOOF  bool             `json:"DeclineAllEventsForScheduledOOF,omitempty"`
	DeclineEventsForScheduledOOF     bool             `json:"DeclineEventsForScheduledOOF,omitempty"`
	EventsToDeleteIDs                []string         `json:"EventsToDeleteIDs,omitempty"`
	EndTime                          time.Time        `json:"EndTime,omitempty"`
	ExternalAudience                 ExternalAudience `json:"ExternalAudience,omitempty"`
	ExternalMessage                  string           `json:"ExternalMessage,omitempty"`
	InternalMessage                  string           `json:"InternalMessage,omitempty"`
	DeclineMeetingMessage            string           `json:"DeclineMeetingMessage,omitempty"`
	OOFEventSubject                  string           `json:"OOFEventSubject,omitempty"`
	Recipients                       []string         `json:"Recipients,omitempty"`
	ReminderMinutesBeforeStart       int32            `json:"ReminderMinutesBeforeStart,omitempty"`
	ReminderMessage                  string           `json:"ReminderMessage,omitempty"`
	StartTime                        time.Time        `json:"StartTime,omitempty"`
}

// MailboxFolder (OData EntityType).
type MailboxFolder struct {
	Identity                string                     `json:"Identity,omitempty"`
	Name                    string                     `json:"Name,omitempty"`
	MailboxFolderPermission []*MailboxFolderPermission `json:"MailboxFolderPermission,omitempty"`
}

// MailboxFolderPermission (OData ComplexType).
type MailboxFolderPermission struct {
	Identity               string               `json:"Identity,omitempty"`
	MailboxFolderIdentity  string               `json:"MailboxFolderIdentity,omitempty"`
	FolderName             string               `json:"FolderName,omitempty"`
	User                   string               `json:"User,omitempty"`
	MailboxFolderUser      *MailboxFolderUserID `json:"MailboxFolderUser,omitempty"`
	AccessRights           []string             `json:"AccessRights,omitempty"`
	SharingPermissionFlags string               `json:"SharingPermissionFlags,omitempty"`
}

// MailboxFolderStatistics (OData ComplexType).
type MailboxFolderStatistics struct {
	ComplexEntry
	Name                              string   `json:"Name,omitempty"`
	SearchFolder                      bool     `json:"SearchFolder,omitempty"`
	CreationTime                      string   `json:"CreationTime,omitempty"`
	LastModifiedTime                  string   `json:"LastModifiedTime,omitempty"`
	FolderPath                        string   `json:"FolderPath,omitempty"`
	FolderId                          string   `json:"FolderId,omitempty"`
	ContentFolder                     bool     `json:"ContentFolder,omitempty"`
	ContentMailboxGuid                string   `json:"ContentMailboxGuid,omitempty"`
	Identity                          string   `json:"Identity,omitempty"`
	FolderType                        string   `json:"FolderType,omitempty"`
	RawContentMailboxGuid             string   `json:"RawContentMailboxGuid,omitempty"`
	Movable                           bool     `json:"Movable,omitempty"`
	RecoverableItemsFolder            bool     `json:"RecoverableItemsFolder,omitempty"`
	AssociatedIPMFolderPath           string   `json:"AssociatedIPMFolderPath,omitempty"`
	ContainerClass                    string   `json:"ContainerClass,omitempty"`
	Flags                             string   `json:"Flags,omitempty"`
	TargetQuota                       string   `json:"TargetQuota,omitempty"`
	StorageQuota                      int64    `json:"StorageQuota,omitempty"`
	StorageWarningQuota               int64    `json:"StorageWarningQuota,omitempty"`
	VisibleItemsInFolder              int32    `json:"VisibleItemsInFolder,omitempty"`
	HiddenItemsInFolder               int32    `json:"HiddenItemsInFolder,omitempty"`
	ItemsInFolder                     int32    `json:"ItemsInFolder,omitempty"`
	DeletedItemsInFolder              int32    `json:"DeletedItemsInFolder,omitempty"`
	FolderSize                        int64    `json:"FolderSize,omitempty"`
	ItemsInFolderAndSubfolders        int32    `json:"ItemsInFolderAndSubfolders,omitempty"`
	DeletedItemsInFolderAndSubfolders int32    `json:"DeletedItemsInFolderAndSubfolders,omitempty"`
	FolderAndSubfolderSize            int64    `json:"FolderAndSubfolderSize,omitempty"`
	CurrentSchemaVersion              string   `json:"CurrentSchemaVersion,omitempty"`
	OldestItemReceivedDate            string   `json:"OldestItemReceivedDate,omitempty"`
	NewestItemReceivedDate            string   `json:"NewestItemReceivedDate,omitempty"`
	OldestDeletedItemReceivedDate     string   `json:"OldestDeletedItemReceivedDate,omitempty"`
	NewestDeletedItemReceivedDate     string   `json:"NewestDeletedItemReceivedDate,omitempty"`
	OldestItemLastModifiedDate        string   `json:"OldestItemLastModifiedDate,omitempty"`
	NewestItemLastModifiedDate        string   `json:"NewestItemLastModifiedDate,omitempty"`
	OldestDeletedItemLastModifiedDate string   `json:"OldestDeletedItemLastModifiedDate,omitempty"`
	NewestDeletedItemLastModifiedDate string   `json:"NewestDeletedItemLastModifiedDate,omitempty"`
	ManagedFolder                     string   `json:"ManagedFolder,omitempty"`
	DeletePolicy                      string   `json:"DeletePolicy,omitempty"`
	ArchivePolicy                     string   `json:"ArchivePolicy,omitempty"`
	TopSubject                        string   `json:"TopSubject,omitempty"`
	TopSubjectSize                    int64    `json:"TopSubjectSize,omitempty"`
	TopSubjectCount                   int32    `json:"TopSubjectCount,omitempty"`
	TopSubjectClass                   string   `json:"TopSubjectClass,omitempty"`
	TopSubjectPath                    string   `json:"TopSubjectPath,omitempty"`
	TopSubjectReceivedTime            string   `json:"TopSubjectReceivedTime,omitempty"`
	TopSubjectFrom                    string   `json:"TopSubjectFrom,omitempty"`
	TopClientInfoForSubject           string   `json:"TopClientInfoForSubject,omitempty"`
	TopClientInfoCountForSubject      int32    `json:"TopClientInfoCountForSubject,omitempty"`
	SearchFolders                     []string `json:"SearchFolders,omitempty"`
	AuditAuxMailboxGuid               string   `json:"AuditAuxMailboxGuid,omitempty"`
	AuditFolderStubSize               string   `json:"AuditFolderStubSize,omitempty"`
	LastMovedTimeStamp                string   `json:"LastMovedTimeStamp,omitempty"`
	LowLatencyContainerId             string   `json:"LowLatencyContainerId,omitempty"`
	LowLatencyContainerFlags          string   `json:"LowLatencyContainerFlags,omitempty"`
	LowLatencyContainerQuota          int64    `json:"LowLatencyContainerQuota,omitempty"`
	Diagnostics                       string   `json:"Diagnostics,omitempty"`
	DiagnosticInfo                    string   `json:"DiagnosticInfo,omitempty"`
}

// MailboxFolderUserID (OData ComplexType).
type MailboxFolderUserID struct {
	ComplexEntry
	UserPrincipalName string `json:"UserPrincipalName,omitempty"`
	UserType          string `json:"UserType,omitempty"`
	DisplayName       string `json:"DisplayName,omitempty"`
}

// MailboxPermission (OData ComplexType).
type MailboxPermission struct {
	PermissionId    string                   `json:"PermissionId,omitempty"`
	MailboxIdentity string                   `json:"MailboxIdentity,omitempty"`
	User            string                   `json:"User,omitempty"`
	IsOwner         bool                     `json:"IsOwner,omitempty"`
	PermissionList  []*MailboxPermissionInfo `json:"PermissionList,omitempty"`
}

// MailboxPermissionInfo (OData ComplexType).
type MailboxPermissionInfo struct {
	ComplexEntry
	AccessRights    []string `json:"AccessRights,omitempty"`
	IsInherited     bool     `json:"IsInherited,omitempty"`
	Deny            bool     `json:"Deny,omitempty"`
	InheritanceType string   `json:"InheritanceType,omitempty"`
}

// MailboxPlan (OData ComplexType).
type MailboxPlan struct {
	Identity       string `json:"Identity,omitempty"`
	DisplayName    string `json:"DisplayName,omitempty"`
	MaxSendSize    int64  `json:"MaxSendSize,omitempty"`
	MaxReceiveSize int64  `json:"MaxReceiveSize,omitempty"`
}

// MailboxRecoverableItem (OData ComplexType).
type MailboxRecoverableItem struct {
	MailboxIdentity  string `json:"MailboxIdentity,omitempty"`
	ItemClass        string `json:"ItemClass,omitempty"`
	Subject          string `json:"Subject,omitempty"`
	EntryID          string `json:"EntryID,omitempty"`
	SourceFolder     string `json:"SourceFolder,omitempty"`
	LastModifiedTime string `json:"LastModifiedTime,omitempty"`
	LastParentPath   string `json:"LastParentPath,omitempty"`
}

// MailboxRecoverableItemsQuery (OData ComplexType).
type MailboxRecoverableItemsQuery struct {
	ComplexEntry
	Identity           []string  `json:"Identity,omitempty"`
	SubjectContains    string    `json:"SubjectContains,omitempty"`
	FilterItemType     string    `json:"FilterItemType,omitempty"`
	FilterStartTime    time.Time `json:"FilterStartTime,omitempty"`
	FilterEndTime      time.Time `json:"FilterEndTime,omitempty"`
	EntryID            string    `json:"EntryID,omitempty"`
	SourceFolder       string    `json:"SourceFolder,omitempty"`
	LastParentFolderID string    `json:"LastParentFolderID,omitempty"`
	ResultSize         int32     `json:"ResultSize,omitempty"`
}

// MailboxRegionalConfiguration (OData ComplexType).
type MailboxRegionalConfiguration struct {
	Identity   string `json:"Identity,omitempty"`
	DateFormat string `json:"DateFormat,omitempty"`
	Language   string `json:"Language,omitempty"`
	TimeFormat string `json:"TimeFormat,omitempty"`
	TimeZone   string `json:"TimeZone,omitempty"`
}

// MailboxStatistics (OData ComplexType).
type MailboxStatistics struct {
	ComplexEntry
	DisplayName                                string `json:"DisplayName,omitempty"`
	ExternalDirectoryOrganizationId            string `json:"ExternalDirectoryOrganizationId,omitempty"`
	LastLoggedOnUserAccount                    string `json:"LastLoggedOnUserAccount,omitempty"`
	LastLogoffTime                             string `json:"LastLogoffTime,omitempty"`
	LastLogonTime                              string `json:"LastLogonTime,omitempty"`
	StorageLimitStatus                         string `json:"StorageLimitStatus,omitempty"`
	SystemMessageSize                          int64  `json:"SystemMessageSize,omitempty"`
	SystemMessageSizeWarningQuota              int64  `json:"SystemMessageSizeWarningQuota,omitempty"`
	SystemMessageSizeShutoffQuota              int64  `json:"SystemMessageSizeShutoffQuota,omitempty"`
	SystemMessageCount                         int64  `json:"SystemMessageCount,omitempty"`
	DisconnectDate                             string `json:"DisconnectDate,omitempty"`
	DisconnectReason                           string `json:"DisconnectReason,omitempty"`
	LegacyDN                                   string `json:"LegacyDN,omitempty"`
	OwnerADGuid                                string `json:"OwnerADGuid,omitempty"`
	MailboxType                                string `json:"MailboxType,omitempty"`
	MailboxTypeDetail                          string `json:"MailboxTypeDetail,omitempty"`
	ServerName                                 string `json:"ServerName,omitempty"`
	DatabaseName                               string `json:"DatabaseName,omitempty"`
	IsDatabaseCopyActive                       bool   `json:"IsDatabaseCopyActive,omitempty"`
	IsAbandonedMoveDestination                 bool   `json:"IsAbandonedMoveDestination,omitempty"`
	MailboxGuid                                string `json:"MailboxGuid,omitempty"`
	IsArchiveMailbox                           bool   `json:"IsArchiveMailbox,omitempty"`
	IsMoveDestination                          bool   `json:"IsMoveDestination,omitempty"`
	ResourceUsageRollingAvgRop                 int32  `json:"ResourceUsageRollingAvgRop,omitempty"`
	ResourceUsageRollingAvgDatabaseReads       int32  `json:"ResourceUsageRollingAvgDatabaseReads,omitempty"`
	ResourceUsageRollingClientTypes            int64  `json:"ResourceUsageRollingClientTypes,omitempty"`
	IsHighDensityShard                         bool   `json:"IsHighDensityShard,omitempty"`
	NeedsToMove                                string `json:"NeedsToMove,omitempty"`
	AssociatedItemCount                        int64  `json:"AssociatedItemCount,omitempty"`
	DeletedItemCount                           int64  `json:"DeletedItemCount,omitempty"`
	ItemCount                                  int64  `json:"ItemCount,omitempty"`
	TotalDeletedItemSize                       int64  `json:"TotalDeletedItemSize,omitempty"`
	TotalItemSize                              int64  `json:"TotalItemSize,omitempty"`
	MessageTableTotalSize                      int64  `json:"MessageTableTotalSize,omitempty"`
	MessageTableAvailableSize                  int64  `json:"MessageTableAvailableSize,omitempty"`
	AttachmentTableTotalSize                   int64  `json:"AttachmentTableTotalSize,omitempty"`
	AttachmentTableAvailableSize               int64  `json:"AttachmentTableAvailableSize,omitempty"`
	OtherTablesTotalSize                       int64  `json:"OtherTablesTotalSize,omitempty"`
	OtherTablesAvailableSize                   int64  `json:"OtherTablesAvailableSize,omitempty"`
	TablesTotalSize                            int64  `json:"TablesTotalSize,omitempty"`
	TablesTotalAvailableSize                   int64  `json:"TablesTotalAvailableSize,omitempty"`
	MailboxMessagesPerFolderCountWarningQuota  int64  `json:"MailboxMessagesPerFolderCountWarningQuota,omitempty"`
	MailboxMessagesPerFolderCountReceiveQuota  int64  `json:"MailboxMessagesPerFolderCountReceiveQuota,omitempty"`
	DumpsterMessagesPerFolderCountWarningQuota int64  `json:"DumpsterMessagesPerFolderCountWarningQuota,omitempty"`
	DumpsterMessagesPerFolderCountReceiveQuota int64  `json:"DumpsterMessagesPerFolderCountReceiveQuota,omitempty"`
	FolderHierarchyChildrenCountWarningQuota   int64  `json:"FolderHierarchyChildrenCountWarningQuota,omitempty"`
	FolderHierarchyChildrenCountReceiveQuota   int64  `json:"FolderHierarchyChildrenCountReceiveQuota,omitempty"`
	FolderHierarchyDepthWarningQuota           int64  `json:"FolderHierarchyDepthWarningQuota,omitempty"`
	FolderHierarchyDepthReceiveQuota           int64  `json:"FolderHierarchyDepthReceiveQuota,omitempty"`
	FoldersCountWarningQuota                   int64  `json:"FoldersCountWarningQuota,omitempty"`
	FoldersCountReceiveQuota                   int64  `json:"FoldersCountReceiveQuota,omitempty"`
	NamedPropertiesCountQuota                  int64  `json:"NamedPropertiesCountQuota,omitempty"`
	DatabaseIssueWarningQuota                  int64  `json:"DatabaseIssueWarningQuota,omitempty"`
	DatabaseProhibitSendQuota                  int64  `json:"DatabaseProhibitSendQuota,omitempty"`
	DatabaseProhibitSendReceiveQuota           int64  `json:"DatabaseProhibitSendReceiveQuota,omitempty"`
	IsQuarantined                              bool   `json:"IsQuarantined,omitempty"`
	QuarantineClients                          string `json:"QuarantineClients,omitempty"`
	LastInteractionTime                        string `json:"LastInteractionTime,omitempty"`
	LastUserActionTime                         string `json:"LastUserActionTime,omitempty"`
}

// MalwareFilterPolicy (OData ComplexType).
type MalwareFilterPolicy struct {
	Identity                               string                  `json:"Identity,omitempty"`
	CustomAlertText                        string                  `json:"CustomAlertText,omitempty"`
	AdminDisplayName                       string                  `json:"AdminDisplayName,omitempty"`
	CustomInternalSubject                  string                  `json:"CustomInternalSubject,omitempty"`
	CustomInternalBody                     string                  `json:"CustomInternalBody,omitempty"`
	CustomExternalSubject                  string                  `json:"CustomExternalSubject,omitempty"`
	CustomExternalBody                     string                  `json:"CustomExternalBody,omitempty"`
	CustomFromName                         string                  `json:"CustomFromName,omitempty"`
	CustomFromAddress                      string                  `json:"CustomFromAddress,omitempty"`
	InternalSenderAdminAddress             string                  `json:"InternalSenderAdminAddress,omitempty"`
	ExternalSenderAdminAddress             string                  `json:"ExternalSenderAdminAddress,omitempty"`
	BypassInboundMessages                  bool                    `json:"BypassInboundMessages,omitempty"`
	BypassOutboundMessages                 bool                    `json:"BypassOutboundMessages,omitempty"`
	Action                                 MalwareFilteringAction  `json:"Action,omitempty"`
	FileTypeAction                         FileTypeFilteringAction `json:"FileTypeAction,omitempty"`
	IsDefault                              bool                    `json:"IsDefault,omitempty"`
	IsPolicyOverrideApplied                bool                    `json:"IsPolicyOverrideApplied,omitempty"`
	CustomNotifications                    bool                    `json:"CustomNotifications,omitempty"`
	EnableInternalSenderNotifications      bool                    `json:"EnableInternalSenderNotifications,omitempty"`
	EnableExternalSenderNotifications      bool                    `json:"EnableExternalSenderNotifications,omitempty"`
	EnableInternalSenderAdminNotifications bool                    `json:"EnableInternalSenderAdminNotifications,omitempty"`
	EnableExternalSenderAdminNotifications bool                    `json:"EnableExternalSenderAdminNotifications,omitempty"`
	EnableFileFilter                       bool                    `json:"EnableFileFilter,omitempty"`
	FileTypes                              []string                `json:"FileTypes,omitempty"`
	ZapEnabled                             bool                    `json:"ZapEnabled,omitempty"`
	RecommendedPolicyType                  RecommendedPolicyType   `json:"RecommendedPolicyType,omitempty"`
	Name                                   string                  `json:"Name,omitempty"`
	WhenChanged                            time.Time               `json:"WhenChanged,omitempty"`
	WhenCreated                            time.Time               `json:"WhenCreated,omitempty"`
	ExchangeObjectId                       string                  `json:"ExchangeObjectId,omitempty"`
	OrganizationId                         string                  `json:"OrganizationId,omitempty"`
	Guid                                   string                  `json:"Guid,omitempty"`
}

// MalwareFilterRule (OData ComplexType).
type MalwareFilterRule struct {
	Identity       string    `json:"Identity,omitempty"`
	State          string    `json:"State,omitempty"`
	Priority       int32     `json:"Priority,omitempty"`
	Comments       string    `json:"Comments,omitempty"`
	Guid           string    `json:"Guid,omitempty"`
	ImmutableId    string    `json:"ImmutableId,omitempty"`
	OrganizationId string    `json:"OrganizationId,omitempty"`
	Name           string    `json:"Name,omitempty"`
	WhenChanged    time.Time `json:"WhenChanged,omitempty"`
}

// MessageTrace (OData ComplexType).
type MessageTrace struct {
	ComplexEntry
	MessageId        string    `json:"MessageId,omitempty"`
	Received         time.Time `json:"Received,omitempty"`
	SenderAddress    string    `json:"SenderAddress,omitempty"`
	RecipientAddress string    `json:"RecipientAddress,omitempty"`
	Subject          string    `json:"Subject,omitempty"`
	Status           string    `json:"Status,omitempty"`
	ToIP             string    `json:"ToIP,omitempty"`
	FromIP           string    `json:"FromIP,omitempty"`
	Size             int32     `json:"Size,omitempty"`
	MessageTraceId   string    `json:"MessageTraceId,omitempty"`
}

// MessageTraceDetail (OData ComplexType).
type MessageTraceDetail struct {
	ComplexEntry
	MessageId      string    `json:"MessageId,omitempty"`
	Date           time.Time `json:"Date,omitempty"`
	Event          string    `json:"Event,omitempty"`
	Action         string    `json:"Action,omitempty"`
	Detail         string    `json:"Detail,omitempty"`
	Data           string    `json:"Data,omitempty"`
	MessageTraceId string    `json:"MessageTraceId,omitempty"`
}

// MobileDevice (OData ComplexType).
type MobileDevice struct {
	Identity string `json:"Identity,omitempty"`
}

// MobileDeviceMailboxPolicy (OData ComplexType).
type MobileDeviceMailboxPolicy struct {
	Identity                     string    `json:"Identity,omitempty"`
	Id                           string    `json:"Id,omitempty"`
	Guid                         string    `json:"Guid,omitempty"`
	IsDefault                    bool      `json:"IsDefault,omitempty"`
	Name                         string    `json:"Name,omitempty"`
	WhenChanged                  string    `json:"WhenChanged,omitempty"`
	WhenCreated                  string    `json:"WhenCreated,omitempty"`
	WhenChangedUTC               time.Time `json:"WhenChangedUTC,omitempty"`
	WhenCreatedUTC               time.Time `json:"WhenCreatedUTC,omitempty"`
	AllowNonProvisionableDevices bool      `json:"AllowNonProvisionableDevices,omitempty"`
	PasswordEnabled              bool      `json:"PasswordEnabled,omitempty"`
	AllowSimplePassword          bool      `json:"AllowSimplePassword,omitempty"`
	AlphanumericPasswordRequired bool      `json:"AlphanumericPasswordRequired,omitempty"`
	MinPasswordComplexCharacters int32     `json:"MinPasswordComplexCharacters,omitempty"`
	RequireDeviceEncryption      bool      `json:"RequireDeviceEncryption,omitempty"`
	MinPasswordLength            int32     `json:"MinPasswordLength,omitempty"`
	MaxPasswordFailedAttempts    int32     `json:"MaxPasswordFailedAttempts,omitempty"`
	MaxInactivityTimeLock        string    `json:"MaxInactivityTimeLock,omitempty"`
	PasswordExpiration           string    `json:"PasswordExpiration,omitempty"`
	PasswordHistory              int32     `json:"PasswordHistory,omitempty"`
}

// MobileDeviceStatistics (OData ComplexType).
type MobileDeviceStatistics struct {
	ComplexEntry
	AccountOnlyDeviceWipeAckTime       string `json:"AccountOnlyDeviceWipeAckTime,omitempty"`
	AccountOnlyDeviceWipeRequestTime   string `json:"AccountOnlyDeviceWipeRequestTime,omitempty"`
	AccountOnlyDeviceWipeSentTime      string `json:"AccountOnlyDeviceWipeSentTime,omitempty"`
	ClientType                         string `json:"ClientType,omitempty"`
	ClientVersion                      string `json:"ClientVersion,omitempty"`
	DeviceAccessControlRule            string `json:"DeviceAccessControlRule,omitempty"`
	DeviceAccessState                  string `json:"DeviceAccessState,omitempty"`
	DeviceAccessStateReason            string `json:"DeviceAccessStateReason,omitempty"`
	DeviceEnableOutboundSMS            bool   `json:"DeviceEnableOutboundSMS,omitempty"`
	DeviceFriendlyName                 string `json:"DeviceFriendlyName,omitempty"`
	DeviceID                           string `json:"DeviceID,omitempty"`
	DeviceImei                         string `json:"DeviceImei,omitempty"`
	DeviceMobileOperator               string `json:"DeviceMobileOperator,omitempty"`
	DeviceModel                        string `json:"DeviceModel,omitempty"`
	DeviceOS                           string `json:"DeviceOS,omitempty"`
	DeviceOSLanguage                   string `json:"DeviceOSLanguage,omitempty"`
	DevicePhoneNumber                  string `json:"DevicePhoneNumber,omitempty"`
	DevicePolicyApplicationStatus      string `json:"DevicePolicyApplicationStatus,omitempty"`
	DevicePolicyApplied                string `json:"DevicePolicyApplied,omitempty"`
	DeviceType                         string `json:"DeviceType,omitempty"`
	DeviceUserAgent                    string `json:"DeviceUserAgent,omitempty"`
	DeviceWipeAckTime                  string `json:"DeviceWipeAckTime,omitempty"`
	DeviceWipeRequestTime              string `json:"DeviceWipeRequestTime,omitempty"`
	DeviceWipeSentTime                 string `json:"DeviceWipeSentTime,omitempty"`
	FirstSyncTime                      string `json:"FirstSyncTime,omitempty"`
	Guid                               string `json:"Guid,omitempty"`
	Identity                           string `json:"Identity,omitempty"`
	IsRemoteWipeSupported              bool   `json:"IsRemoteWipeSupported,omitempty"`
	LastAccountOnlyDeviceWipeRequestor string `json:"LastAccountOnlyDeviceWipeRequestor,omitempty"`
	LastDeviceWipeRequestor            string `json:"LastDeviceWipeRequestor,omitempty"`
	LastPingHeartbeat                  int64  `json:"LastPingHeartbeat,omitempty"`
	LastPolicyUpdateTime               string `json:"LastPolicyUpdateTime,omitempty"`
	LastSuccessSync                    string `json:"LastSuccessSync,omitempty"`
	LastSyncAttemptTime                string `json:"LastSyncAttemptTime,omitempty"`
	MailboxLogReport                   string `json:"MailboxLogReport,omitempty"`
	NumberOfFoldersSynced              int32  `json:"NumberOfFoldersSynced,omitempty"`
	Status                             string `json:"Status,omitempty"`
	StatusNote                         string `json:"StatusNote,omitempty"`
	SyncStateUpgradeTime               string `json:"SyncStateUpgradeTime,omitempty"`
}

// OpenComplexEntry (OData ComplexType).
type OpenComplexEntry struct {
}

// OutboundConnector (OData ComplexType).
type OutboundConnector struct {
	ConnectorType                 string    `json:"ConnectorType,omitempty"`
	ConnectorSource               string    `json:"ConnectorSource,omitempty"`
	Enabled                       bool      `json:"Enabled,omitempty"`
	UseMXRecord                   bool      `json:"UseMXRecord,omitempty"`
	Comment                       string    `json:"Comment,omitempty"`
	SmartHosts                    []string  `json:"SmartHosts,omitempty"`
	RecipientDomains              []string  `json:"RecipientDomains,omitempty"`
	TlsDomain                     string    `json:"TlsDomain,omitempty"`
	TlsSettings                   string    `json:"TlsSettings,omitempty"`
	IsTransportRuleScoped         bool      `json:"IsTransportRuleScoped,omitempty"`
	RouteAllMessagesViaOnPremises bool      `json:"RouteAllMessagesViaOnPremises,omitempty"`
	CloudServicesMailEnabled      bool      `json:"CloudServicesMailEnabled,omitempty"`
	AllAcceptedDomains            bool      `json:"AllAcceptedDomains,omitempty"`
	TestMode                      bool      `json:"TestMode,omitempty"`
	LinkForModifiedConnector      string    `json:"LinkForModifiedConnector,omitempty"`
	ValidationRecipients          []string  `json:"ValidationRecipients,omitempty"`
	IsValidated                   bool      `json:"IsValidated,omitempty"`
	LastValidationTimestamp       time.Time `json:"LastValidationTimestamp,omitempty"`
	Name                          string    `json:"Name,omitempty"`
	Identity                      string    `json:"Identity,omitempty"`
	Guid                          string    `json:"Guid,omitempty"`
}

// OutboundConnectorValidationResult (OData ComplexType).
type OutboundConnectorValidationResult struct {
	ComplexEntry
	IsTaskSuccessful bool                                 `json:"IsTaskSuccessful,omitempty"`
	TaskName         string                               `json:"TaskName,omitempty"`
	TaskDetail       string                               `json:"TaskDetail,omitempty"`
	SubTaskResults   []*OutboundConnectorValidationResult `json:"SubTaskResults,omitempty"`
}

// Place (OData ComplexType).
type Place struct {
	Identity               string   `json:"Identity,omitempty"`
	IsWheelChairAccessible bool     `json:"IsWheelChairAccessible,omitempty"`
	Tags                   []string `json:"Tags,omitempty"`
	AudioDeviceName        string   `json:"AudioDeviceName,omitempty"`
	DisplayDeviceName      string   `json:"DisplayDeviceName,omitempty"`
	VideoDeviceName        string   `json:"VideoDeviceName,omitempty"`
}

// RbacResults (OData ComplexType).
type RbacResults struct {
	ComplexEntry
	Results []bool `json:"Results,omitempty"`
}

// Recipient (OData EntityType).
type Recipient struct {
	ObjectKey                          string                 `json:"ObjectKey,omitempty"`
	ExternalDirectoryObjectId          string                 `json:"ExternalDirectoryObjectId,omitempty"`
	Identity                           string                 `json:"Identity,omitempty"`
	Alias                              string                 `json:"Alias,omitempty"`
	EmailAddresses                     []string               `json:"EmailAddresses,omitempty"`
	DisplayName                        string                 `json:"DisplayName,omitempty"`
	FirstName                          string                 `json:"FirstName,omitempty"`
	LastName                           string                 `json:"LastName,omitempty"`
	Name                               string                 `json:"Name,omitempty"`
	ArchiveGuid                        string                 `json:"ArchiveGuid,omitempty"`
	AuthenticationType                 string                 `json:"AuthenticationType,omitempty"`
	City                               string                 `json:"City,omitempty"`
	Notes                              string                 `json:"Notes,omitempty"`
	Company                            string                 `json:"Company,omitempty"`
	CountryOrRegion                    string                 `json:"CountryOrRegion,omitempty"`
	PostalCode                         string                 `json:"PostalCode,omitempty"`
	CustomAttribute1                   string                 `json:"CustomAttribute1,omitempty"`
	CustomAttribute2                   string                 `json:"CustomAttribute2,omitempty"`
	CustomAttribute3                   string                 `json:"CustomAttribute3,omitempty"`
	CustomAttribute4                   string                 `json:"CustomAttribute4,omitempty"`
	CustomAttribute5                   string                 `json:"CustomAttribute5,omitempty"`
	CustomAttribute6                   string                 `json:"CustomAttribute6,omitempty"`
	CustomAttribute7                   string                 `json:"CustomAttribute7,omitempty"`
	CustomAttribute8                   string                 `json:"CustomAttribute8,omitempty"`
	CustomAttribute9                   string                 `json:"CustomAttribute9,omitempty"`
	CustomAttribute10                  string                 `json:"CustomAttribute10,omitempty"`
	CustomAttribute11                  string                 `json:"CustomAttribute11,omitempty"`
	CustomAttribute12                  string                 `json:"CustomAttribute12,omitempty"`
	CustomAttribute13                  string                 `json:"CustomAttribute13,omitempty"`
	CustomAttribute14                  string                 `json:"CustomAttribute14,omitempty"`
	CustomAttribute15                  string                 `json:"CustomAttribute15,omitempty"`
	ExtensionCustomAttribute1          []string               `json:"ExtensionCustomAttribute1,omitempty"`
	ExtensionCustomAttribute2          []string               `json:"ExtensionCustomAttribute2,omitempty"`
	ExtensionCustomAttribute3          []string               `json:"ExtensionCustomAttribute3,omitempty"`
	ExtensionCustomAttribute4          []string               `json:"ExtensionCustomAttribute4,omitempty"`
	ExtensionCustomAttribute5          []string               `json:"ExtensionCustomAttribute5,omitempty"`
	Database                           string                 `json:"Database,omitempty"`
	ArchiveDatabase                    string                 `json:"ArchiveDatabase,omitempty"`
	DatabaseName                       string                 `json:"DatabaseName,omitempty"`
	Department                         string                 `json:"Department,omitempty"`
	ManagedFolderMailboxPolicy         string                 `json:"ManagedFolderMailboxPolicy,omitempty"`
	ExpansionServer                    string                 `json:"ExpansionServer,omitempty"`
	ExternalEmailAddress               string                 `json:"ExternalEmailAddress,omitempty"`
	HiddenFromAddressListsEnabled      bool                   `json:"HiddenFromAddressListsEnabled,omitempty"`
	EmailAddressPolicyEnabled          bool                   `json:"EmailAddressPolicyEnabled,omitempty"`
	ResourceType                       string                 `json:"ResourceType,omitempty"`
	ManagedBy                          []string               `json:"ManagedBy,omitempty"`
	Manager                            string                 `json:"Manager,omitempty"`
	ActiveSyncMailboxPolicy            string                 `json:"ActiveSyncMailboxPolicy,omitempty"`
	ActiveSyncMailboxPolicyIsDefaulted bool                   `json:"ActiveSyncMailboxPolicyIsDefaulted,omitempty"`
	Office                             string                 `json:"Office,omitempty"`
	ObjectCategory                     string                 `json:"ObjectCategory,omitempty"`
	OrganizationalUnit                 string                 `json:"OrganizationalUnit,omitempty"`
	Phone                              string                 `json:"Phone,omitempty"`
	PoliciesIncluded                   []string               `json:"PoliciesIncluded,omitempty"`
	PoliciesExcluded                   []string               `json:"PoliciesExcluded,omitempty"`
	PrimarySmtpAddress                 string                 `json:"PrimarySmtpAddress,omitempty"`
	RecipientType                      string                 `json:"RecipientType,omitempty"`
	RecipientTypeDetails               string                 `json:"RecipientTypeDetails,omitempty"`
	SamAccountName                     string                 `json:"SamAccountName,omitempty"`
	ServerLegacyDN                     string                 `json:"ServerLegacyDN,omitempty"`
	ServerName                         string                 `json:"ServerName,omitempty"`
	StateOrProvince                    string                 `json:"StateOrProvince,omitempty"`
	StorageGroupName                   string                 `json:"StorageGroupName,omitempty"`
	Title                              string                 `json:"Title,omitempty"`
	UMMailboxPolicy                    string                 `json:"UMMailboxPolicy,omitempty"`
	UMRecipientDialPlanId              string                 `json:"UMRecipientDialPlanId,omitempty"`
	WindowsLiveID                      string                 `json:"WindowsLiveID,omitempty"`
	HasActiveSyncDevicePartnership     bool                   `json:"HasActiveSyncDevicePartnership,omitempty"`
	AddressListMembership              []string               `json:"AddressListMembership,omitempty"`
	OwaMailboxPolicy                   string                 `json:"OwaMailboxPolicy,omitempty"`
	AddressBookPolicy                  string                 `json:"AddressBookPolicy,omitempty"`
	SharingPolicy                      string                 `json:"SharingPolicy,omitempty"`
	RetentionPolicy                    string                 `json:"RetentionPolicy,omitempty"`
	ShouldUseDefaultRetentionPolicy    bool                   `json:"ShouldUseDefaultRetentionPolicy,omitempty"`
	MailboxMoveTargetMDB               string                 `json:"MailboxMoveTargetMDB,omitempty"`
	MailboxMoveSourceMDB               string                 `json:"MailboxMoveSourceMDB,omitempty"`
	MailboxMoveFlags                   string                 `json:"MailboxMoveFlags,omitempty"`
	MailboxMoveRemoteHostName          string                 `json:"MailboxMoveRemoteHostName,omitempty"`
	MailboxMoveBatchName               string                 `json:"MailboxMoveBatchName,omitempty"`
	MailboxMoveStatus                  string                 `json:"MailboxMoveStatus,omitempty"`
	MailboxRelease                     string                 `json:"MailboxRelease,omitempty"`
	ArchiveRelease                     string                 `json:"ArchiveRelease,omitempty"`
	IsValidSecurityPrincipal           bool                   `json:"IsValidSecurityPrincipal,omitempty"`
	LitigationHoldEnabled              bool                   `json:"LitigationHoldEnabled,omitempty"`
	Capabilities                       []string               `json:"Capabilities,omitempty"`
	ArchiveState                       string                 `json:"ArchiveState,omitempty"`
	SKUAssigned                        bool                   `json:"SKUAssigned,omitempty"`
	WhenMailboxCreated                 string                 `json:"WhenMailboxCreated,omitempty"`
	UsageLocation                      string                 `json:"UsageLocation,omitempty"`
	ExchangeGuid                       string                 `json:"ExchangeGuid,omitempty"`
	ArchiveStatus                      string                 `json:"ArchiveStatus,omitempty"`
	SafeSendersHash                    []byte                 `json:"SafeSendersHash,omitempty"`
	SafeRecipientsHash                 []byte                 `json:"SafeRecipientsHash,omitempty"`
	BlockedSendersHash                 []byte                 `json:"BlockedSendersHash,omitempty"`
	WhenSoftDeleted                    string                 `json:"WhenSoftDeleted,omitempty"`
	ExchangeVersion                    string                 `json:"ExchangeVersion,omitempty"`
	DistinguishedName                  string                 `json:"DistinguishedName,omitempty"`
	ObjectClass                        []string               `json:"ObjectClass,omitempty"`
	WhenChanged                        string                 `json:"WhenChanged,omitempty"`
	WhenCreated                        string                 `json:"WhenCreated,omitempty"`
	WhenChangedUTC                     string                 `json:"WhenChangedUTC,omitempty"`
	WhenCreatedUTC                     string                 `json:"WhenCreatedUTC,omitempty"`
	ExchangeObjectId                   string                 `json:"ExchangeObjectId,omitempty"`
	OrganizationId                     string                 `json:"OrganizationId,omitempty"`
	Id                                 string                 `json:"Id,omitempty"`
	Guid                               string                 `json:"Guid,omitempty"`
	Members                            string                 `json:"Members,omitempty"`
	RecipientPermission                []*RecipientPermission `json:"RecipientPermission,omitempty"`
}

// RecipientPermission (OData ComplexType).
type RecipientPermission struct {
	Identity          string                 `json:"Identity,omitempty"`
	Trustee           string                 `json:"Trustee,omitempty"`
	AccessControlType string                 `json:"AccessControlType,omitempty"`
	AccessRights      []RecipientAccessRight `json:"AccessRights,omitempty"`
	IsInherited       bool                   `json:"IsInherited,omitempty"`
	InheritanceType   string                 `json:"InheritanceType,omitempty"`
}

// ReportSchedule (OData ComplexType).
type ReportSchedule struct {
	StartDate               time.Time `json:"StartDate,omitempty"`
	ExpiryDate              time.Time `json:"ExpiryDate,omitempty"`
	RecipientAddress        string    `json:"RecipientAddress,omitempty"`
	SenderAddress           string    `json:"SenderAddress,omitempty"`
	Direction               string    `json:"Direction,omitempty"`
	EncryptionTemplate      string    `json:"EncryptionTemplate,omitempty"`
	EncryptionType          string    `json:"EncryptionType,omitempty"`
	NotificationEmail       []string  `json:"NotificationEmail,omitempty"`
	ReportFrequency         string    `json:"ReportFrequency,omitempty"`
	ReportType              string    `json:"ReportType,omitempty"`
	ScheduleName            string    `json:"ScheduleName,omitempty"`
	ReportStatusDescription string    `json:"ReportStatusDescription,omitempty"`
	Locale                  string    `json:"Locale,omitempty"`
	Identity                string    `json:"Identity,omitempty"`
	ScheduleId              string    `json:"ScheduleId,omitempty"`
	OrganizationId          string    `json:"OrganizationId,omitempty"`
}

// RetentionPolicy (OData ComplexType).
type RetentionPolicy struct {
	IsDefaultArbitrationMailbox bool      `json:"IsDefaultArbitrationMailbox,omitempty"`
	Identity                    string    `json:"Identity,omitempty"`
	AdminDisplayName            string    `json:"AdminDisplayName,omitempty"`
	IsDefault                   bool      `json:"IsDefault,omitempty"`
	IsValid                     bool      `json:"IsValid,omitempty"`
	Name                        string    `json:"Name,omitempty"`
	WhenChangedUTC              time.Time `json:"WhenChangedUTC,omitempty"`
	WhenCreatedUTC              time.Time `json:"WhenCreatedUTC,omitempty"`
}

// RoleAssignmentPolicy (OData ComplexType).
type RoleAssignmentPolicy struct {
	Description      string    `json:"Description,omitempty"`
	Identity         string    `json:"Identity,omitempty"`
	AdminDisplayName string    `json:"AdminDisplayName,omitempty"`
	IsDefault        bool      `json:"IsDefault,omitempty"`
	IsValid          bool      `json:"IsValid,omitempty"`
	Name             string    `json:"Name,omitempty"`
	WhenChangedUTC   time.Time `json:"WhenChangedUTC,omitempty"`
	WhenCreatedUTC   time.Time `json:"WhenCreatedUTC,omitempty"`
}

// RoleAssignments (OData EntityType).
type RoleAssignments struct {
	Id                               string                   `json:"id,omitempty"`
	Identity                         string                   `json:"identity,omitempty"`
	RunspaceId                       string                   `json:"runspaceId,omitempty"`
	DataObject                       string                   `json:"dataObject,omitempty"`
	User                             string                   `json:"user,omitempty"`
	AssignmentMethod                 string                   `json:"assignmentMethod,omitempty"`
	RoleAssigneeType                 string                   `json:"roleAssigneeType,omitempty"`
	PrincipalId                      string                   `json:"principalId,omitempty"`
	RoleDefinitionId                 string                   `json:"roleDefinitionId,omitempty"`
	Condition                        string                   `json:"condition,omitempty"`
	DirectoryScopeId                 string                   `json:"directoryScopeId,omitempty"`
	ResourceScope                    string                   `json:"resourceScope,omitempty"`
	RoleAssignmentDelegationType     string                   `json:"roleAssignmentDelegationType,omitempty"`
	CustomRecipientWriteScope        string                   `json:"customRecipientWriteScope,omitempty"`
	CustomConfigWriteScope           string                   `json:"customConfigWriteScope,omitempty"`
	RecipientReadScope               string                   `json:"recipientReadScope,omitempty"`
	ConfigReadScope                  string                   `json:"configReadScope,omitempty"`
	RecipientWriteScope              string                   `json:"recipientWriteScope,omitempty"`
	ConfigWriteScope                 string                   `json:"configWriteScope,omitempty"`
	Enabled                          bool                     `json:"enabled,omitempty"`
	RoleAssigneeName                 string                   `json:"roleAssigneeName,omitempty"`
	IsValid                          bool                     `json:"isValid,omitempty"`
	Version                          string                   `json:"version,omitempty"`
	DisplayName                      string                   `json:"displayName,omitempty"`
	DistinguishedName                string                   `json:"distinguishedName,omitempty"`
	ObjectCategory                   string                   `json:"objectCategory,omitempty"`
	ObjectClass                      []string                 `json:"objectClass,omitempty"`
	WhenChangedDateTime              time.Time                `json:"whenChangedDateTime,omitempty"`
	WhenCreatedDateTime              time.Time                `json:"whenCreatedDateTime,omitempty"`
	WhenChangedUTCDateTime           time.Time                `json:"whenChangedUTCDateTime,omitempty"`
	WhenCreatedUTCDateTime           time.Time                `json:"whenCreatedUTCDateTime,omitempty"`
	ExchangeObjectId                 string                   `json:"exchangeObjectId,omitempty"`
	OrganizationId                   string                   `json:"organizationId,omitempty"`
	Guid                             string                   `json:"guid,omitempty"`
	OriginatingServer                string                   `json:"originatingServer,omitempty"`
	AppScopeId                       string                   `json:"appScopeId,omitempty"`
	RecipientAdministrativeUnitScope string                   `json:"recipientAdministrativeUnitScope,omitempty"`
	RoleDefinition                   *RoleDefinitions         `json:"roleDefinition,omitempty"`
	AppScope                         *ExchangeManagementScope `json:"appScope,omitempty"`
}

// RoleDefinitions (OData ComplexType).
type RoleDefinitions struct {
	Id                          string            `json:"id,omitempty"`
	Description                 string            `json:"description,omitempty"`
	DisplayName                 string            `json:"displayName,omitempty"`
	Parent                      string            `json:"parent,omitempty"`
	RoleEntries                 []string          `json:"roleEntries,omitempty"`
	RolePermissions             []*RolePermission `json:"rolePermissions,omitempty"`
	ImplicitRecipientReadScope  string            `json:"implicitRecipientReadScope,omitempty"`
	ImplicitRecipientWriteScope string            `json:"implicitRecipientWriteScope,omitempty"`
	ImplicitConfigReadScope     string            `json:"implicitConfigReadScope,omitempty"`
	ImplicitConfigWriteScope    string            `json:"implicitConfigWriteScope,omitempty"`
	IsRootRole                  bool              `json:"isRootRole,omitempty"`
	IsEndUserRole               bool              `json:"isEndUserRole,omitempty"`
	IsEnabled                   bool              `json:"isEnabled,omitempty"`
	IsDeprecated                bool              `json:"isDeprecated,omitempty"`
	IsValid                     bool              `json:"isValid,omitempty"`
	Version                     string            `json:"version,omitempty"`
	RoleType                    string            `json:"roleType,omitempty"`
	EnabledCmdlets              []string          `json:"enabledCmdlets,omitempty"`
	Guid                        string            `json:"guid,omitempty"`
	Identity                    string            `json:"identity,omitempty"`
}

// RolePermission (OData ComplexType).
type RolePermission struct {
	ComplexEntry
	AllowedResourceActions  []string `json:"allowedResourceActions,omitempty"`
	ExcludedResourceActions []string `json:"excludedResourceActions,omitempty"`
	Condition               string   `json:"condition,omitempty"`
}

// SafeAttachmentPolicy (OData ComplexType).
type SafeAttachmentPolicy struct {
	Identity                                      string                      `json:"Identity,omitempty"`
	RedirectAddress                               string                      `json:"RedirectAddress,omitempty"`
	Redirect                                      bool                        `json:"Redirect,omitempty"`
	Action                                        SafeAttachmentAction        `json:"Action,omitempty"`
	ScanTimeout                                   int32                       `json:"ScanTimeout,omitempty"`
	ConfidenceLevelThreshold                      int32                       `json:"ConfidenceLevelThreshold,omitempty"`
	OperationMode                                 SafeAttachmentOperationMode `json:"OperationMode,omitempty"`
	Enable                                        bool                        `json:"Enable,omitempty"`
	ActionOnError                                 bool                        `json:"ActionOnError,omitempty"`
	RecommendedPolicyType                         RecommendedPolicyType       `json:"RecommendedPolicyType,omitempty"`
	IsDefault                                     bool                        `json:"IsDefault,omitempty"`
	AdminDisplayName                              string                      `json:"AdminDisplayName,omitempty"`
	EnableBlockingEncryptedAttachments            bool                        `json:"EnableBlockingEncryptedAttachments,omitempty"`
	EnableOrganizationBranding                    bool                        `json:"EnableOrganizationBranding,omitempty"`
	ExcludedTypesFromBlockingEncryptedAttachments []EncryptedFileTypes        `json:"ExcludedTypesFromBlockingEncryptedAttachments,omitempty"`
	Name                                          string                      `json:"Name,omitempty"`
	WhenChanged                                   time.Time                   `json:"WhenChanged,omitempty"`
	WhenCreated                                   time.Time                   `json:"WhenCreated,omitempty"`
	ExchangeObjectId                              string                      `json:"ExchangeObjectId,omitempty"`
	OrganizationId                                string                      `json:"OrganizationId,omitempty"`
	Guid                                          string                      `json:"Guid,omitempty"`
}

// SafeAttachmentRule (OData ComplexType).
type SafeAttachmentRule struct {
	Identity       string    `json:"Identity,omitempty"`
	State          string    `json:"State,omitempty"`
	Priority       int32     `json:"Priority,omitempty"`
	Comments       string    `json:"Comments,omitempty"`
	Guid           string    `json:"Guid,omitempty"`
	ImmutableId    string    `json:"ImmutableId,omitempty"`
	OrganizationId string    `json:"OrganizationId,omitempty"`
	Name           string    `json:"Name,omitempty"`
	WhenChanged    time.Time `json:"WhenChanged,omitempty"`
}

// SafeLinksPolicyPresentation (OData ComplexType).
type SafeLinksPolicyPresentation struct {
	Identity                   string                `json:"Identity,omitempty"`
	TrackClicks                bool                  `json:"TrackClicks,omitempty"`
	AllowClickThrough          bool                  `json:"AllowClickThrough,omitempty"`
	ScanUrls                   bool                  `json:"ScanUrls,omitempty"`
	EnableForInternalSenders   bool                  `json:"EnableForInternalSenders,omitempty"`
	DeliverMessageAfterScan    bool                  `json:"DeliverMessageAfterScan,omitempty"`
	DoNotRewriteUrls           []string              `json:"DoNotRewriteUrls,omitempty"`
	AdminDisplayName           string                `json:"AdminDisplayName,omitempty"`
	EnableSafeLinksForEmail    bool                  `json:"EnableSafeLinksForEmail,omitempty"`
	EnableSafeLinksForTeams    bool                  `json:"EnableSafeLinksForTeams,omitempty"`
	EnableSafeLinksForOffice   bool                  `json:"EnableSafeLinksForOffice,omitempty"`
	DisableUrlRewrite          bool                  `json:"DisableUrlRewrite,omitempty"`
	CustomNotificationText     string                `json:"CustomNotificationText,omitempty"`
	EnableOrganizationBranding bool                  `json:"EnableOrganizationBranding,omitempty"`
	RecommendedPolicyType      RecommendedPolicyType `json:"RecommendedPolicyType,omitempty"`
	Name                       string                `json:"Name,omitempty"`
	WhenChanged                time.Time             `json:"WhenChanged,omitempty"`
	WhenCreated                time.Time             `json:"WhenCreated,omitempty"`
	ExchangeObjectId           string                `json:"ExchangeObjectId,omitempty"`
	OrganizationId             string                `json:"OrganizationId,omitempty"`
	Guid                       string                `json:"Guid,omitempty"`
}

// SafeLinksRule (OData ComplexType).
type SafeLinksRule struct {
	Identity       string    `json:"Identity,omitempty"`
	State          string    `json:"State,omitempty"`
	Priority       int32     `json:"Priority,omitempty"`
	Comments       string    `json:"Comments,omitempty"`
	Guid           string    `json:"Guid,omitempty"`
	ImmutableId    string    `json:"ImmutableId,omitempty"`
	OrganizationId string    `json:"OrganizationId,omitempty"`
	Name           string    `json:"Name,omitempty"`
	WhenChanged    time.Time `json:"WhenChanged,omitempty"`
}

// SecurityPrincipal (OData ComplexType).
type SecurityPrincipal struct {
	Id                   string    `json:"Id,omitempty"`
	Identity             string    `json:"Identity,omitempty"`
	Guid                 string    `json:"Guid,omitempty"`
	ExchangeObjectId     string    `json:"ExchangeObjectId,omitempty"`
	DisplayName          string    `json:"DisplayName,omitempty"`
	RecipientTypeDetails string    `json:"RecipientTypeDetails,omitempty"`
	Type                 string    `json:"Type,omitempty"`
	ExchangeVersion      string    `json:"ExchangeVersion,omitempty"`
	Name                 string    `json:"Name,omitempty"`
	DistinguishedName    string    `json:"DistinguishedName,omitempty"`
	ObjectCategory       string    `json:"ObjectCategory,omitempty"`
	ObjectClass          []string  `json:"ObjectClass,omitempty"`
	WhenCreated          time.Time `json:"WhenCreated,omitempty"`
	WhenChangedUTC       time.Time `json:"WhenChangedUTC,omitempty"`
	WhenCreatedUTC       time.Time `json:"WhenCreatedUTC,omitempty"`
	OrganizationId       string    `json:"OrganizationId,omitempty"`
	IsValid              bool      `json:"IsValid,omitempty"`
	ObjectState          string    `json:"ObjectState,omitempty"`
}

// SharingPolicy (OData ComplexType).
type SharingPolicy struct {
	Identity         string    `json:"Identity,omitempty"`
	AdminDisplayName string    `json:"AdminDisplayName,omitempty"`
	Default          bool      `json:"Default,omitempty"`
	IsValid          bool      `json:"IsValid,omitempty"`
	Name             string    `json:"Name,omitempty"`
	WhenChangedUTC   time.Time `json:"WhenChangedUTC,omitempty"`
	WhenCreatedUTC   time.Time `json:"WhenCreatedUTC,omitempty"`
}

// StringFieldDeltaUpdateData (OData ComplexType).
type StringFieldDeltaUpdateData struct {
	ComplexEntry
	Add    []string `json:"Add,omitempty"`
	Remove []string `json:"Remove,omitempty"`
}

// TestMigrationServerAvailability (OData ComplexType).
type TestMigrationServerAvailability struct {
	Id                 string `json:"Id,omitempty"`
	Imap               bool   `json:"Imap,omitempty"`
	RemoteServer       string `json:"RemoteServer,omitempty"`
	Port               int32  `json:"Port,omitempty"`
	Security           string `json:"Security,omitempty"`
	Authentication     string `json:"Authentication,omitempty"`
	M365RcaTest        bool   `json:"M365RcaTest,omitempty"`
	TokenStore         string `json:"TokenStore,omitempty"`
	EmailAddress       string `json:"EmailAddress,omitempty"`
	Organization       string `json:"Organization,omitempty"`
	CredentialUsername string `json:"CredentialUsername,omitempty"`
	CredentialPassword string `json:"CredentialPassword,omitempty"`
	Result             string `json:"Result,omitempty"`
	Message            string `json:"Message,omitempty"`
	ErrorDetail        string `json:"ErrorDetail,omitempty"`
	SupportsCutover    bool   `json:"SupportsCutover,omitempty"`
}

// UnifiedGroup (OData ComplexType).
type UnifiedGroup struct {
	Guid                                                   string    `json:"Guid,omitempty"`
	Identity                                               string    `json:"Identity,omitempty"`
	GrantSendOnBehalfTo                                    []string  `json:"GrantSendOnBehalfTo,omitempty"`
	Language                                               string    `json:"Language,omitempty"`
	MigrationToUnifiedGroupInProgress                      bool      `json:"MigrationToUnifiedGroupInProgress,omitempty"`
	RejectMessagesFromSendersOrMembers                     []string  `json:"RejectMessagesFromSendersOrMembers,omitempty"`
	AcceptMessagesOnlyFromSendersOrMembers                 []string  `json:"AcceptMessagesOnlyFromSendersOrMembers,omitempty"`
	GrantSendOnBehalfToWithDisplayNames                    []string  `json:"GrantSendOnBehalfToWithDisplayNames,omitempty"`
	AcceptMessagesOnlyFromSendersOrMembersWithDisplayNames []string  `json:"AcceptMessagesOnlyFromSendersOrMembersWithDisplayNames,omitempty"`
	RejectMessagesFromSendersOrMembersWithDisplayNames     []string  `json:"RejectMessagesFromSendersOrMembersWithDisplayNames,omitempty"`
	SendAsPermissionList                                   []string  `json:"SendAsPermissionList,omitempty"`
	ObjectId                                               string    `json:"ObjectId,omitempty"`
	DistinguishedName                                      string    `json:"DistinguishedName,omitempty"`
	WindowsLiveID                                          string    `json:"WindowsLiveID,omitempty"`
	RecipientType                                          string    `json:"RecipientType,omitempty"`
	RecipientTypeDetails                                   string    `json:"RecipientTypeDetails,omitempty"`
	IsValid                                                bool      `json:"IsValid,omitempty"`
	WhenCreatedUTC                                         time.Time `json:"WhenCreatedUTC,omitempty"`
}

// UnifiedRbacManagementScope (OData ComplexType).
type UnifiedRbacManagementScope struct {
	Id                         string               `json:"id,omitempty"`
	Type                       ScopeRestrictionType `json:"type,omitempty"`
	DisplayName                string               `json:"displayName,omitempty"`
	CustomAttributes           *GenericHashTable    `json:"customAttributes,omitempty"`
	RecipientRoot              string               `json:"recipientRoot,omitempty"`
	Filter                     string               `json:"filter,omitempty"`
	RecipientFilter            string               `json:"recipientFilter,omitempty"`
	Exclusive                  bool                 `json:"exclusive,omitempty"`
	Name                       string               `json:"name,omitempty"`
	DistinguishedName          string               `json:"distinguishedName,omitempty"`
	Guid                       string               `json:"guid,omitempty"`
	Identity                   string               `json:"identity,omitempty"`
	RecipientRestrictionFilter string               `json:"recipientRestrictionFilter,omitempty"`
	ScopeRestrictionType       ScopeRestrictionType `json:"scopeRestrictionType,omitempty"`
}

// UnifiedRbacRoleAssignment (OData EntityType).
type UnifiedRbacRoleAssignment struct {
	Id                               string                      `json:"id,omitempty"`
	PrincipalId                      string                      `json:"principalId,omitempty"`
	RoleDefinitionId                 string                      `json:"roleDefinitionId,omitempty"`
	DirectoryScopeId                 string                      `json:"directoryScopeId,omitempty"`
	AppScopeId                       string                      `json:"appScopeId,omitempty"`
	Identity                         string                      `json:"identity,omitempty"`
	User                             string                      `json:"user,omitempty"`
	AssignmentMethod                 string                      `json:"assignmentMethod,omitempty"`
	RoleAssigneeType                 string                      `json:"roleAssigneeType,omitempty"`
	RoleAssignmentDelegationType     string                      `json:"roleAssignmentDelegationType,omitempty"`
	CustomRecipientWriteScope        string                      `json:"customRecipientWriteScope,omitempty"`
	CustomConfigWriteScope           string                      `json:"customConfigWriteScope,omitempty"`
	RecipientReadScope               string                      `json:"recipientReadScope,omitempty"`
	ConfigReadScope                  string                      `json:"configReadScope,omitempty"`
	RecipientWriteScope              string                      `json:"recipientWriteScope,omitempty"`
	ConfigWriteScope                 string                      `json:"configWriteScope,omitempty"`
	Enabled                          bool                        `json:"enabled,omitempty"`
	RoleAssigneeName                 string                      `json:"roleAssigneeName,omitempty"`
	IsValid                          bool                        `json:"isValid,omitempty"`
	DisplayName                      string                      `json:"displayName,omitempty"`
	DistinguishedName                string                      `json:"distinguishedName,omitempty"`
	Guid                             string                      `json:"guid,omitempty"`
	RecipientAdministrativeUnitScope string                      `json:"recipientAdministrativeUnitScope,omitempty"`
	RoleDefinition                   *UnifiedRbacRoleDefinition  `json:"roleDefinition,omitempty"`
	AppScope                         *UnifiedRbacManagementScope `json:"appScope,omitempty"`
}

// UnifiedRbacRoleDefinition (OData EntityType).
type UnifiedRbacRoleDefinition struct {
	Id                      string                     `json:"id,omitempty"`
	Description             string                     `json:"description,omitempty"`
	DisplayName             string                     `json:"displayName,omitempty"`
	RolePermissions         []*RolePermission          `json:"rolePermissions,omitempty"`
	IsEnabled               bool                       `json:"isEnabled,omitempty"`
	Version                 string                     `json:"version,omitempty"`
	IsBuiltIn               bool                       `json:"isBuiltIn,omitempty"`
	TemplateId              string                     `json:"templateId,omitempty"`
	AllowedPrincipalTypes   RolePrincipalTypes         `json:"allowedPrincipalTypes,omitempty"`
	Parent                  string                     `json:"parent,omitempty"`
	RoleEntries             []string                   `json:"roleEntries,omitempty"`
	IsRootRole              bool                       `json:"isRootRole,omitempty"`
	IsEndUserRole           bool                       `json:"isEndUserRole,omitempty"`
	IsDeprecated            bool                       `json:"isDeprecated,omitempty"`
	IsValid                 bool                       `json:"isValid,omitempty"`
	RoleType                string                     `json:"roleType,omitempty"`
	EnabledCmdlets          []string                   `json:"enabledCmdlets,omitempty"`
	Guid                    string                     `json:"guid,omitempty"`
	Identity                string                     `json:"identity,omitempty"`
	IsServicePrincipalRole  bool                       `json:"isServicePrincipalRole,omitempty"`
	InheritsPermissionsFrom *UnifiedRbacRoleDefinition `json:"inheritsPermissionsFrom,omitempty"`
}

// User (OData ComplexType).
type User struct {
	ExternalDirectoryObjectId           string   `json:"ExternalDirectoryObjectId,omitempty"`
	MailboxProvisioningConstraint       string   `json:"MailboxProvisioningConstraint,omitempty"`
	MailboxRegion                       string   `json:"MailboxRegion,omitempty"`
	MailboxRegionLastUpdateTime         int64    `json:"MailboxRegionLastUpdateTime,omitempty"`
	MailboxProvisioningPreferences      []string `json:"MailboxProvisioningPreferences,omitempty"`
	MailboxLocations                    []string `json:"MailboxLocations,omitempty"`
	IsLinked                            bool     `json:"IsLinked,omitempty"`
	LinkedMasterAccount                 string   `json:"LinkedMasterAccount,omitempty"`
	ResetPasswordOnNextLogon            bool     `json:"ResetPasswordOnNextLogon,omitempty"`
	RemotePowerShellEnabled             bool     `json:"RemotePowerShellEnabled,omitempty"`
	EXOModuleEnabled                    bool     `json:"EXOModuleEnabled,omitempty"`
	SamAccountName                      string   `json:"SamAccountName,omitempty"`
	Office                              string   `json:"Office,omitempty"`
	UserPrincipalName                   string   `json:"UserPrincipalName,omitempty"`
	NetID                               string   `json:"NetID,omitempty"`
	ConsumerNetID                       string   `json:"ConsumerNetID,omitempty"`
	WindowsLiveID                       string   `json:"WindowsLiveID,omitempty"`
	MicrosoftOnlineServicesID           string   `json:"MicrosoftOnlineServicesID,omitempty"`
	MailboxRelease                      string   `json:"MailboxRelease,omitempty"`
	ArchiveRelease                      string   `json:"ArchiveRelease,omitempty"`
	SKUAssigned                         bool     `json:"SKUAssigned,omitempty"`
	IsSoftDeletedByRemove               bool     `json:"IsSoftDeletedByRemove,omitempty"`
	IsSoftDeletedByDisable              bool     `json:"IsSoftDeletedByDisable,omitempty"`
	IsInactiveMailbox                   bool     `json:"IsInactiveMailbox,omitempty"`
	WhenSoftDeleted                     int64    `json:"WhenSoftDeleted,omitempty"`
	AccountDisabled                     bool     `json:"AccountDisabled,omitempty"`
	StsRefreshTokensValidFrom           string   `json:"StsRefreshTokensValidFrom,omitempty"`
	IsDirSynced                         bool     `json:"IsDirSynced,omitempty"`
	AdministrativeUnits                 []string `json:"AdministrativeUnits,omitempty"`
	OrganizationalUnit                  string   `json:"OrganizationalUnit,omitempty"`
	DisplayName                         string   `json:"DisplayName,omitempty"`
	LegacyExchangeDN                    string   `json:"LegacyExchangeDN,omitempty"`
	RecipientType                       string   `json:"RecipientType,omitempty"`
	RecipientTypeDetails                string   `json:"RecipientTypeDetails,omitempty"`
	SimpleDisplayName                   string   `json:"SimpleDisplayName,omitempty"`
	UMDtmfMap                           []string `json:"UMDtmfMap,omitempty"`
	WindowsEmailAddress                 string   `json:"WindowsEmailAddress,omitempty"`
	Identity                            string   `json:"Identity,omitempty"`
	Id                                  string   `json:"Id,omitempty"`
	ExchangeVersion                     string   `json:"ExchangeVersion,omitempty"`
	Name                                string   `json:"Name,omitempty"`
	DistinguishedName                   string   `json:"DistinguishedName,omitempty"`
	ObjectCategory                      string   `json:"ObjectCategory,omitempty"`
	ObjectClass                         []string `json:"ObjectClass,omitempty"`
	WhenChanged                         int64    `json:"WhenChanged,omitempty"`
	WhenCreated                         int64    `json:"WhenCreated,omitempty"`
	WhenChangedUTC                      int64    `json:"WhenChangedUTC,omitempty"`
	WhenCreatedUTC                      int64    `json:"WhenCreatedUTC,omitempty"`
	ExchangeObjectId                    string   `json:"ExchangeObjectId,omitempty"`
	OrganizationId                      string   `json:"OrganizationId,omitempty"`
	Guid                                string   `json:"Guid,omitempty"`
	UserAccountControl                  string   `json:"UserAccountControl,omitempty"`
	City                                string   `json:"City,omitempty"`
	Company                             string   `json:"Company,omitempty"`
	CountryOrRegion                     string   `json:"CountryOrRegion,omitempty"`
	Department                          string   `json:"Department,omitempty"`
	Fax                                 string   `json:"Fax,omitempty"`
	FirstName                           string   `json:"FirstName,omitempty"`
	LastName                            string   `json:"LastName,omitempty"`
	HomePhone                           string   `json:"HomePhone,omitempty"`
	OtherHomePhone                      string   `json:"OtherHomePhone,omitempty"`
	OtherTelephone                      string   `json:"OtherTelephone,omitempty"`
	MobilePhone                         string   `json:"MobilePhone,omitempty"`
	OtherFax                            string   `json:"OtherFax,omitempty"`
	Pager                               string   `json:"Pager,omitempty"`
	Phone                               string   `json:"Phone,omitempty"`
	TelephoneAssistant                  string   `json:"TelephoneAssistant,omitempty"`
	PhoneticDisplayName                 string   `json:"PhoneticDisplayName,omitempty"`
	PostalCode                          string   `json:"PostalCode,omitempty"`
	StateOrProvince                     string   `json:"StateOrProvince,omitempty"`
	StreetAddress                       string   `json:"StreetAddress,omitempty"`
	Title                               string   `json:"Title,omitempty"`
	WebPage                             string   `json:"WebPage,omitempty"`
	AssistantName                       string   `json:"AssistantName,omitempty"`
	SeniorityIndex                      string   `json:"SeniorityIndex,omitempty"`
	Notes                               string   `json:"Notes,omitempty"`
	Initials                            string   `json:"Initials,omitempty"`
	DirectReports                       []string `json:"DirectReports,omitempty"`
	DirectReportsWithDisplayName        []string `json:"DirectReportsWithDisplayName,omitempty"`
	Manager                             string   `json:"Manager,omitempty"`
	ManagerWithDisplayName              []string `json:"ManagerWithDisplayName,omitempty"`
	VoiceMailSettings                   []string `json:"VoiceMailSettings,omitempty"`
	PostOfficeBox                       []string `json:"PostOfficeBox,omitempty"`
	PreviousRecipientTypeDetails        string   `json:"PreviousRecipientTypeDetails,omitempty"`
	Sid                                 string   `json:"Sid,omitempty"`
	SidHistory                          []string `json:"SidHistory,omitempty"`
	GeoCoordinates                      string   `json:"GeoCoordinates,omitempty"`
	AuthenticationPolicy                string   `json:"AuthenticationPolicy,omitempty"`
	AllowUMCallsFromNonUsers            string   `json:"AllowUMCallsFromNonUsers,omitempty"`
	CertificateSubject                  []string `json:"CertificateSubject,omitempty"`
	CloudCacheProvider                  int32    `json:"CloudCacheProvider,omitempty"`
	CloudCacheScope                     int32    `json:"CloudCacheScope,omitempty"`
	CloudCacheRemoteEmailAddress        string   `json:"CloudCacheRemoteEmailAddress,omitempty"`
	CloudCacheUserName                  string   `json:"CloudCacheUserName,omitempty"`
	IsCloudCache                        bool     `json:"IsCloudCache,omitempty"`
	IsCloudCacheBlocked                 bool     `json:"IsCloudCacheBlocked,omitempty"`
	IsCloudCacheProvisioningComplete    bool     `json:"IsCloudCacheProvisioningComplete,omitempty"`
	InPlaceHoldsRaw                     []string `json:"InPlaceHoldsRaw,omitempty"`
	MailboxWorkloads                    string   `json:"MailboxWorkloads,omitempty"`
	DesiredMailboxWorkloads             string   `json:"DesiredMailboxWorkloads,omitempty"`
	DesiredMailboxWorkloadsGracePeriod  string   `json:"DesiredMailboxWorkloadsGracePeriod,omitempty"`
	DesiredMailboxWorkloadsModified     string   `json:"DesiredMailboxWorkloadsModified,omitempty"`
	IsSecurityPrincipal                 bool     `json:"IsSecurityPrincipal,omitempty"`
	LegalAgeGroup                       int32    `json:"LegalAgeGroup,omitempty"`
	UMDialPlan                          string   `json:"UMDialPlan,omitempty"`
	UMCallingLineIds                    []string `json:"UMCallingLineIds,omitempty"`
	UpgradeMessage                      string   `json:"UpgradeMessage,omitempty"`
	UpgradeDetails                      string   `json:"UpgradeDetails,omitempty"`
	UpgradeRequest                      string   `json:"UpgradeRequest,omitempty"`
	UpgradeStage                        string   `json:"UpgradeStage,omitempty"`
	UpgradeStageTimeStamp               int64    `json:"UpgradeStageTimeStamp,omitempty"`
	UpgradeStatus                       string   `json:"UpgradeStatus,omitempty"`
	CanHaveCloudCache                   bool     `json:"CanHaveCloudCache,omitempty"`
	PublicFolder                        bool     `json:"PublicFolder,omitempty"`
	CreateDTMFMap                       bool     `json:"CreateDTMFMap,omitempty"`
	BlockCloudCache                     bool     `json:"BlockCloudCache,omitempty"`
	PermanentlyClearPreviousMailboxInfo bool     `json:"PermanentlyClearPreviousMailboxInfo,omitempty"`
	RemoveMailboxProvisioningConstraint bool     `json:"RemoveMailboxProvisioningConstraint,omitempty"`
	Vip                                 bool     `json:"Vip,omitempty"`
}
