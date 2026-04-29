package model

import "time"

// User represents an Emby server user.
type User struct {
	ID                    string    `json:"Id"`
	Name                  string    `json:"Name"`
	Username              string    `json:"Username"`
	EmailAddress          string    `json:"Email,omitempty"`
	LoginUsername         string    `json:"LoginUsername,omitempty"`
	LoginPassword         string    `json:"LoginPassword,omitempty"`
	InvalidLoginAttemptCount int    `json:"InvalidLoginAttemptCount,omitempty"`
	LastLoginDate         time.Time `json:"LastLoginDate,omitempty"`
	LastActivityDate      time.Time `json:"LastActivityDate,omitempty"`
	AuthenticationProviderID string `json:"AuthenticationProviderId,omitempty"`
	PrimaryImageTag       string    `json:"PrimaryImageTag,omitempty"`
	MissingEULAVersion    string    `json:"MissingEulaVersion,omitempty"`
	SyncPlayAccess        string    `json:"SyncPlayAccess,omitempty"`
	HasConfiguredAutoLiveSettings bool `json:"HasConfiguredAutoLiveSettings,omitempty"`
	EnableAutoLogin       bool      `json:"EnableAutoLogin,omitempty"`
	DisableManageMediaBrowserServer bool `json:"DisableManageMediaBrowserServer,omitempty"`
	DisableAutoCommunityRating bool    `json:"DisableAutoCommunityRating,omitempty"`
	Policy                *UserPolicy `json:"Policy"`
}

// UserPolicy represents the access policy for a user.
type UserPolicy struct {
	IsAdministrator     bool     `json:"IsAdministrator"`
	IsHidden            bool     `json:"IsHidden"`
	IsHiddenResolved    bool     `json:"IsHiddenResolved"`
	IsDisabled          bool     `json:"IsDisabled"`
	BlockUnratedItems   []string `json:"BlockUnratedItems"`
	BlockChannelsWithoutEntry bool `json:"BlockChannelsWithoutEntry"`
	BlockedTags         []string `json:"BlockedTags"`
	EnabledChannels     []string `json:"EnabledChannels"`
	EnabledFolders      []string `json:"EnabledFolders"`
	MaxParentalRating   int      `json:"MaxParentalRating"`
	EnableRemoteControlOfOtherUsers bool `json:"EnableRemoteControlOfOtherUsers"`
	SharedDevices       []string `json:"SharedDevices"`
	EnabledDevices      []string `json:"EnabledDevices"`
	MaxStreamingBitrate int      `json:"MaxStreamingBitrate"`
	SimultaneousStreamCount int  `json:"SimultaneousStreamCount"`
	LiveTVAccess        string   `json:"LiveTvAccess"`
	Grouping            string   `json:"Grouping"`
	PlayAuthorization   []string `json:"PlayAuthorization"`
	ContentDownloading  string   `json:"ContentDownloading"`
	ManageLiveTv        bool     `json:"ManageLiveTv"`
	ManageMediaBrowserServer bool `json:"ManageMediaBrowserServer"`
	EnableCollectionManagement bool `json:"EnableCollectionManagement"`
	EnableMediaDeletion bool     `json:"EnableMediaDeletion"`
	EnableMediaConversion bool   `json:"EnableMediaConversion"`
	AllowCameraUpload   bool     `json:"AllowCameraUpload"`
	AllowSharingFolders []string `json:"AllowSharingFolders"`
	AllowAllDevicesForLiveTV string `json:"AllowAllDevicesForLiveTv"`
	RemoteClientBitrateOverride int `json:"RemoteClientBitrateOverride"`
}

// UserConfiguration represents user-specific settings.
type UserConfiguration struct {
	SubtitleMode                 string   `json:"SubtitleMode"`
	SubtitleFontSize             int      `json:"SubtitleFontSize"`
	SubtitleLanguage             string   `json:"SubtitleLanguage"`
	MissingEpisodes              bool     `json:"MissingEpisodes"`
	MissingEpisodesSection       int      `json:"MissingEpisodesSection"`
	HidePlayedInLatest           bool     `json:"HidePlayedInLatest"`
	NextEpisodeDays              int      `json:"NextEpisodeDays"`
	PlayNextEpisode              bool     `json:"PlayNextEpisode"`
	GroupedFolders               []string `json:"GroupedFolders"`
	SkipForwardDuration          int      `json:"SkipForwardDuration"`
	SkipBackwardDuration         int      `json:"SkipBackwardDuration"`
	TimedViewingData             string   `json:"TimedViewingData"`
	MaxAudioChannels             string   `json:"MaxAudioChannels"`
	EnableAutoStart              bool     `json:"EnableAutoStart"`
	MinRating                    int      `json:"MinRating"`
	PlayedIndicator              string   `json:"PlayedIndicator"`
	OrderItemsBy                 []string `json:"OrderItemsBy"`
	RememberAudioSelections      bool     `json:"RememberAudioSelections"`
	RememberSubtitleSelections   bool     `json:"RememberSubtitleSelections"`
	EnabledChannels              []interface{} `json:"EnabledChannels"`
	LatestItemsExcluded          []string `json:"LatestItemsExcluded"`
}
