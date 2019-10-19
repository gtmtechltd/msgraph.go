// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EducationOneRosterAPIDataProvider undocumented
type EducationOneRosterAPIDataProvider struct {
	EducationSynchronizationDataProvider
	// ConnectionURL undocumented
	ConnectionURL *string `json:"connectionUrl,omitempty"`
	// ConnectionSettings undocumented
	ConnectionSettings *EducationSynchronizationConnectionSettings `json:"connectionSettings,omitempty"`
	// SchoolsIDs undocumented
	SchoolsIDs []string `json:"schoolsIds,omitempty"`
	// TermIDs undocumented
	TermIDs []string `json:"termIds,omitempty"`
	// ProviderName undocumented
	ProviderName *string `json:"providerName,omitempty"`
	// Customizations undocumented
	Customizations *EducationSynchronizationCustomizations `json:"customizations,omitempty"`
}