package model

import "time"

// SessionInfo represents an active client session.
type SessionInfo struct {
	ID                    string    `json:"Id"`
	AdditionalUsers       []SessionUser `json:"AdditionalUsers"`
	Actions               []Action  `json:"Actions"`
	CanSearch             bool      `json:"CanSearch"`
	CanNavigateMedia      bool      `json:"CanNavigateMedia"`
	CanAdjustMedia        bool      `json:"CanAdjustMedia"`
	CanMediaControl       bool      `json:"CanMediaControl"`
	Capabilities          Capabilities `json:"Capabilities"`
	Client                string    `json:"Client"`
	DeviceName            string    `json:"DeviceName"`
	DisplayName           string    `json:"DisplayName"`
	Endpoint              string    `json:"Endpoint"`
	IconURL               string    `json:"IconUrl,omitempty"`
	LastActivityTime      time.Time `json:"LastActivityTime"`
	LastPlaybackTime      time.Time `json:"LastPlaybackTime,omitempty"`
	LocalAddress          string    `json:"LocalAddress"`
	Location              Location  `json:"Location"`
	MachineID             string    `json:"MachineId"`
	PlaybackPositionTicks int64     `json:"PlaybackPositionTicks"`
	PlayState             PlayState `json:"PlayState"`
	PlayMethod            PlayMethod `json:"PlayMethod"`
	RemoteAddress         string    `json:"RemoteAddress"`
	SupportedCommands     []string  `json:"SupportedCommands"`
	NowPlayingItem        *MediaItem `json:"NowPlayingItem,omitempty"`
	TranscodingInfo       *TranscodingInfo `json:"TranscodingInfo,omitempty"`
	DeviceID              string    `json:"DeviceId"`
	StartTimeTicks        int64     `json:"StartTimeTicks"`
	SupportsMediaControl  bool      `json:"SupportsMediaControl"`
	SupportsPersistentIdentification bool `json:"SupportsPersistentIdentification"`
	SupportsSync          bool      `json:"SupportsSync"`
	IsInActiveSession     bool      `json:"IsInActiveSession"`
	IsTerminal            bool      `json:"IsTerminal"`
}

// SessionUser represents a user associated with a session.
type SessionUser struct {
	ID    string `json:"Id"`
	Name  string `json:"Name"`
	Thumb string `json:"Thumb,omitempty"`
}

// Action represents a command that can be sent to a session.
type Action struct {
	Type   string `json:"Type"`
	Target string `json:"Target,omitempty"`
}

// Capabilities represents the capabilities of a client.
type Capabilities struct {
	SupportedMediaTypes   []string `json:"SupportedMediaTypes"`
	PlayableMediaTypes    []string `json:"PlayableMediaTypes"`
	EnableContentDirs     bool     `json:"EnableContentDirs"`
	EnableContentDeletion bool     `json:"EnableContentDeletion"`
	EnableMediaDeletion   bool     `json:"EnableMediaDeletion"`
	EnableSharedCategories bool     `json:"EnableSharedCategories"`
	EnablePlaybackInfo    bool     `json:"EnablePlaybackInfo"`
	EnableMediaControl    bool     `json:"EnableMediaControl"`
	EnableRemoteControl   bool     `json:"EnableRemoteControl"`
	Id                    string   `json:"Id"`
	IconURL               string   `json:"IconUrl,omitempty"`
	TranscodingProfiles  []*TranscodingProfile `json:"TranscodingProfiles"`
	DirectPlayProfiles   []*DirectPlayProfile `json:"DirectPlayProfiles"`
	TimelineOffsetSeconds int      `json:"TimelineOffsetSeconds"`
	SubtitleProfiles     []SubtitleProfile `json:"SubtitleProfiles"`
}

// TranscodingProfile represents a transcoding profile.
type TranscodingProfile struct {
	Container       string   `json:"container"`
	Type            string   `json:"type"`
	AudioCodec      string   `json:"audioCodec"`
	VideoCodec      string   `json:"videoCodec"`
	MaxAudioChannels string  `json:"maxAudioChannels"`
	MinAudioChannels string  `json:"minAudioChannels"`
	Protocol        string   `json:"protocol"`
	MaxVideoBitrate int      `json:"maxVideoBitrate"`
	MaxAudioBitrate int      `json:"maxAudioBitrate"`
	VideoCodec2     string   `json:"videoCodec2,omitempty"`
	BreakOnNonKeyFrames bool `json:"breakOnNonKeyFrames"`
	MaxWidth        int      `json:"maxWidth"`
	MaxHeight       int      `json:"maxHeight"`
	MaxFrameRate    int      `json:"maxFrameRate"`
	AudioCodec2     string   `json:"audioCodec2,omitempty"`
	ConditionalAudioStreaming bool `json:"conditionalAudioStreaming"`
	Context         string   `json:"context"`
	MaxChannels     int      `json:"maxChannels"`
	Profile         string   `json:"profile,omitempty"`
	Level           string   `json:"level,omitempty"`
	CopyTimestamps  bool     `json:"copyTimestamps"`
	EnableMpegtsM2Ts bool    `json:"enableMpegtsM2Ts"`
	EnableSubtitlesInManifest bool `json:"enableSubtitlesInManifest"`
	MaxVideoResolution int   `json:"maxVideoResolution,omitempty"`
	RefFrames       int      `json:"refFrames,omitempty"`
	VideoBitRate    string   `json:"videoBitrate,omitempty"`
	VideoProfile    string   `json:"videoProfile,omitempty"`
	VideoLevel      string   `json:"videoLevel,omitempty"`
	EncodeH265      bool     `json:"encodeH265,omitempty"`
	NoAudioCodec    bool     `json:"noAudioCodec,omitempty"`
	IsDefault       bool     `json:"isDefault,omitempty"`
}

// DirectPlayProfile represents a direct play profile.
type DirectPlayProfile struct {
	Container  string   `json:"container"`
	Type       string   `json:"type"`
	AudioCodec string   `json:"audioCodec,omitempty"`
	VideoCodec string   `json:"videoCodec,omitempty"`
	VideoProfile string `json:"videoProfile,omitempty"`
	MaxChannels  int      `json:"maxChannels,omitempty"`
}

// SubtitleProfile represents a subtitle profile.
type SubtitleProfile struct {
	Method  SubtitleMethod `json:"method"`
	Format  string         `json:"format"`
	Language string        `json:"language,omitempty"`
}

// SubtitleMethod represents how subtitles are delivered.
type SubtitleMethod string

const (
	SubtitleMethodEmbed    SubtitleMethod = "Embed"
	SubtitleMethodExternal SubtitleMethod = "External"
	SubtitleMethodHls      SubtitleMethod = "Hls"
)

// Location represents the location of a session.
type Location struct {
	DeviceName string `json:"DeviceName"`
	IPAddress  string `json:"IPAddress"`
	Location   string `json:"Location,omitempty"`
}

// PlayState represents the playback state.
type PlayState struct {
	PositionTicks       int64 `json:"PositionTicks"`
	VolumePercent       int   `json:"VolumePercent"`
	IsMuted             bool  `json:"IsMuted"`
	IsPaused            bool  `json:"IsPaused"`
	AudioStreamIndex    int   `json:"AudioStreamIndex"`
	SubtitleStreamIndex int   `json:"SubtitleStreamIndex"`
	SubtitlesOffsetTicks int64 `json:"SubtitlesOffsetTicks"`
}

// PlayMethod represents how media is played.
type PlayMethod string

const (
	PlayMethodTranscode PlayMethod = "Transcode"
	PlayMethodDirectStream PlayMethod = "DirectStream"
	PlayMethodDirectPlay PlayMethod = "DirectPlay"
)

// TranscodingInfo represents transcoding details.
type TranscodingInfo struct {
	IsVideoDirect bool   `json:"isVideoDirect"`
	IsAudioDirect bool   `json:"isAudioDirect"`
	Container     string `json:"container"`
	Bitrate       int    `json:"bitrate"`
	Subtitles     string `json:"subtitles,omitempty"`
	AudioCodec    string `json:"audioCodec"`
	VideoCodec    string `json:"videoCodec"`
}
