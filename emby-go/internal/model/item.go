package model

import "time"

// MediaItem represents a media item in the library.
type MediaItem struct {
	ID              string    `json:"Id"`
	Name            string    `json:"Name"`
	Overview        string    `json:"Overview,omitempty"`
	Tagline         string    `json:"Tagline,omitempty"`
	IndexNumber     int       `json:"IndexNumber,omitempty"`
	ParentIndex     int       `json:"ParentIndex,omitempty"`
	CommunityRating float32   `json:"CommunityRating,omitempty"`
	RunTimeTicks    int64     `json:"RunTimeTicks,omitempty"`
	ProductionYear  int       `json:"ProductionYear,omitempty"`
	OfficialRating  string    `json:"OfficialRating,omitempty"`
	ContentType     string    `json:"ContentType,omitempty"`
	MediaType       string    `json:"MediaType,omitempty"`
	Genres          []string  `json:"Genres"`
	Studios         []string  `json:"Studios"`
	people          []Person  `json:"People"`
	SeasonNumber    int       `json:"SeasonNumber,omitempty"`
	EpisodeNumber   int       `json:"EpisodeNumber,omitempty"`
	Album           string    `json:"Album,omitempty"`
	ArtistItems     []ItemRef `json:"ArtistItems"`
	Artists         []string  `json:"Artists"`
	ExtraType       string    `json:"ExtraType,omitempty"`
	ChannelImage    string    `json:"ChannelImage,omitempty"`
	ChannelNumber   int       `json:"ChannelNumber,omitempty"`
	StartDate       time.Time `json:"StartDate,omitempty"`
	EndDate         time.Time `json:"EndDate,omitempty"`
	IsLive          bool      `json:"IsLive,omitempty"`
	IsSeries        bool      `json:"IsSeries,omitempty"`
	IsMovie         bool      `json:"IsMovie,omitempty"`
	IsNews          bool      `json:"IsNews,omitempty"`
	IsSports        bool      `json:"IsSports,omitempty"`
	IsKids          bool      `json:"IsKids,omitempty"`
	IsPremiere      bool      `json:"IsPremiere,omitempty"`
	LocationType    string    `json:"LocationType"`
	Path            string    `json:"Path"`
	PrimaryImageURL string    `json:"PrimaryImageURL,omitempty"`
	BackdropImageURL []string  `json:"BackdropImageURL"`
	ParentBackdropImageItems []ItemRef `json:"ParentBackdropImageItems"`
	ParentID        string    `json:"ParentId,omitempty"`
	Width           int       `json:"Width,omitempty"`
	Height          int       `json:"Height,omitempty"`
	Video3DFormat   string    `json:"Video3DFormat,omitempty"`
	PostLiveFeedTime int64    `json:"PostLiveFeedTime,omitempty"`
	LiveMediaSourceID string  `json:"LiveMediaSourceId,omitempty"`
	MediaSources    []MediaSource `json:"MediaSources"`
	Channels        []Channel     `json:"Channels,omitempty"`
	CurrentProgram  *MediaItem    `json:"CurrentProgram,omitempty"`
	StartTimeTicks  int64         `json:"StartTimeTicks,omitempty"`
	EndTimeTicks    int64         `json:"EndTimeTicks,omitempty"`
	RemoteImageURL  string        `json:"RemoteImageURL,omitempty"`
	LocalTrailerCount int          `json:"LocalTrailerCount,omitempty"`
	UserData        *UserData     `json:"UserData,omitempty"`
	LockedFields    []Field       `json:"LockedFields"`
	LockData        bool          `json:"LockData,omitempty"`
	Disabled      bool          `json:"Disabled,omitempty"`
	EnableMediaSourceDisplay bool   `json:"EnableMediaSourceDisplay,omitempty"`
	ExtraIds        []string      `json:"ExtraIds"`
}

// MediaSource represents a media source for a media item.
type MediaSource struct {
	ID              string       `json:"Id"`
	Name            string       `json:"Name,omitempty"`
	Type            string       `json:"Type"`
	Container       string       `json:"Container"`
	Size            int64        `json:"Size,omitempty"`
	Path            string       `json:"Path"`
	Protocol        string       `json:"Protocol"`
	Encoder         bool         `json:"Encoder,omitempty"`
	VideoCodec      string       `json:"VideoCodec,omitempty"`
	AudioCodec      string       `json:"AudioCodec,omitempty"`
	Format          string       `json:"Format,omitempty"`
	Width           int          `json:"Width,omitempty"`
	Height          int          `json:"Height,omitempty"`
	RefFrames       int          `json:"RefFrames,omitempty"`
	VideoFramerate  string       `json:"VideoFramerate,omitempty"`
	VideoBitRate    int          `json:"VideoBitrate,omitempty"`
	AudioBitRate    int          `json:"AudioBitrate,omitempty"`
	AudioChannels   int          `json:"AudioChannels,omitempty"`
	AudioSampleRate string       `json:"AudioSampleRate,omitempty"`
	DefaultAudioStreamIndex int  `json:"DefaultAudioStreamIndex,omitempty"`
	SubtitleStreams []StreamInfo `json:"SubtitleStreams"`
	Formats         []string     `json:"Formats"`
	RequiredHttpHeaders map[string]string `json:"RequiredHttpHeaders"`
	SupportsTranscoding bool     `json:"SupportsTranscoding"`
	SupportsDirectStream bool    `json:"SupportsDirectStream"`
	SupportsDirectPlay bool      `json:"SupportsDirectPlay"`
	IsRemote          bool         `json:"IsRemote,omitempty"`
}

// StreamInfo represents a stream of media.
type StreamInfo struct {
	Index        int    `json:"Index"`
	Codec        string `json:"Codec"`
	Language     string `json:"Language,omitempty"`
	Default      bool   `json:"Default"`
	Forced       bool   `json:"Forced"`
	External     bool   `json:"External"`
	Path         string `json:"Path"`
	EncodingCodec string `json:"EncodingCodec,omitempty"`
}

// Channel represents a TV channel.
type Channel struct {
	ID          string `json:"Id"`
	Name        string `json:"Name"`
	Number      string `json:"Number"`
	ImageURL    string `json:"ImageURL,omitempty"`
	CurrentProgram *MediaItem `json:"CurrentProgram,omitempty"`
}

// Person represents a person associated with a media item.
type Person struct {
	ID     string `json:"Id"`
	Name   string `json:"Name"`
	Role   string `json:"Role,omitempty"`
	Type   string `json:"Type"`
	PrimaryImageURL string `json:"PrimaryImageURL,omitempty"`
}

// ItemRef is a reference to another item.
type ItemRef struct {
	ID   string `json:"Id"`
	Name string `json:"Name,omitempty"`
}

// UserData represents user-specific data for a media item.
type UserData struct {
	PlaybackPositionTicks int64 `json:"PlaybackPositionTicks,omitempty"`
	PlayCount            int   `json:"PlayCount,omitempty"`
	IsFavorite           bool  `json:"IsFavorite,omitempty"`
	Liked                bool  `json:"Liked,omitempty"`
	LastPlayedDate       string `json:"LastPlayedDate,omitempty"`
	Played               bool  `json:"Played,omitempty"`
	Key                  string `json:"Key,omitempty"`
	Rating               int   `json:"Rating,omitempty"`
}

// Field represents a metadata field that can be locked.
type Field string

const (
	FieldOverview      Field = "Overview"
	FieldGenres        Field = "Genres"
	FieldRuntime       Field = "Runtime"
	FieldProductionYear Field = "ProductionYear"
	FieldOfficialRating Field = "OfficialRating"
	FieldTagline       Field = "Tagline"
)
