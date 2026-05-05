package model

// StartupConfiguration represents the startup wizard configuration.
type StartupConfiguration struct {
	UICulture                   string `json:"UICulture"`
	MetadataCountryCode         string `json:"MetadataCountryCode"`
	PreferredMetadataLanguage    string `json:"PreferredMetadataLanguage"`
}

// StartupUser represents the startup wizard user info.
type StartupUser struct {
	Name            string `json:"Name"`
	ConnectUserName string `json:"ConnectUserName"`
}

// UpdateStartupUserResult represents the result of updating a startup user.
type UpdateStartupUserResult struct {
	UserLinkResult *UserLinkResult `json:"UserLinkResult,omitempty"`
}

// UserLinkResult represents the result of linking a user to Emby Connect.
type UserLinkResult struct {
	IsPending bool   `json:"IsPending"`
	Username  string `json:"Username,omitempty"`
	Message   string `json:"Message,omitempty"`
}

// StartupRemoteAccess represents the remote access configuration during wizard.
type StartupRemoteAccess struct {
	EnableRemoteAccess        bool `json:"EnableRemoteAccess"`
	EnableAutomaticPortMapping bool `json:"EnableAutomaticPortMapping"`
}

// StartupOptions represents startup options.
type StartupOptions struct {
	EnableUPnP bool `json:"EnableUPnP"`
	EnableDLNA bool `json:"EnableDLNA"`
}

// IsFirstRunResult represents the result of checking if this is first run.
type IsFirstRunResult struct {
	IsFirstRun   bool `json:"IsFirstRun"`
	HasPassword  bool `json:"HasPassword"`
	HasUsername  bool `json:"HasUsername"`
}
