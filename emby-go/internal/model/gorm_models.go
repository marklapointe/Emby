package model

import (
	"time"
)

// GORMItem is the GORM model for Items table
type GORMItem struct {
	Id                           string    `gorm:"column:Id;primaryKey"`
	Name                         string    `gorm:"column:Name;not null"`
	Overview                     string    `gorm:"column:Overview"`
	Tagline                      string    `gorm:"column:Tagline"`
	IndexNumber                  int       `gorm:"column:IndexNumber"`
	ParentIndex                  int       `gorm:"column:ParentIndex"`
	CommunityRating              float32   `gorm:"column:CommunityRating"`
	RunTimeTicks                 int64     `gorm:"column:RunTimeTicks"`
	ProductionYear               int       `gorm:"column:ProductionYear"`
	OfficialRating               string    `gorm:"column:OfficialRating"`
	ContentType                  string    `gorm:"column:ContentType"`
	MediaType                    string    `gorm:"column:MediaType"`
	Genres                       string    `gorm:"column:Genres"`
	Studios                      string    `gorm:"column:Studios"`
	SeasonNumber                 int       `gorm:"column:SeasonNumber"`
	EpisodeNumber                int       `gorm:"column:EpisodeNumber"`
	Album                        string    `gorm:"column:Album"`
	Artists                      string    `gorm:"column:Artists"`
	ExtraType                    string    `gorm:"column:ExtraType"`
	ChannelNumber                int       `gorm:"column:ChannelNumber"`
	StartDate                    time.Time `gorm:"column:StartDate"`
	EndDate                      time.Time `gorm:"column:EndDate"`
	LocationType                 string    `gorm:"column:LocationType"`
	Path                         string    `gorm:"column:Path"`
	PrimaryImageURL              string    `gorm:"column:PrimaryImageURL"`
	BackdropImageURL              string    `gorm:"column:BackdropImageURL"`
	ParentID                     string    `gorm:"column:ParentID"`
	Width                        int       `gorm:"column:Width"`
	Height                       int       `gorm:"column:Height"`
	Video3DFormat                string    `gorm:"column:Video3DFormat"`
	PostLiveFeedTime             int64     `gorm:"column:PostLiveFeedTime"`
	LiveMediaSourceID             string    `gorm:"column:LiveMediaSourceID"`
	StartTimeTicks               int64     `gorm:"column:StartTimeTicks"`
	EndTimeTicks                 int64     `gorm:"column:EndTimeTicks"`
	RemoteImageURL                string    `gorm:"column:RemoteImageURL"`
	LocalTrailerCount            int       `gorm:"column:LocalTrailerCount"`
	LockedFields                 string    `gorm:"column:LockedFields"`
	LockData                     bool      `gorm:"column:LockData"`
	Disabled                      bool      `gorm:"column:Disabled"`
	EnableMediaSourceDisplay      bool      `gorm:"column:EnableMediaSourceDisplay"`
	ExtraIds                     string    `gorm:"column:ExtraIds"`
	CreatedDate                   time.Time `gorm:"column:CreatedDate"`
	ModifiedDate                  time.Time `gorm:"column:ModifiedDate"`
}

func (GORMItem) TableName() string {
	return "Items"
}

// GORMItemMediaType represents the media types an item belongs to (Movie, Series, Live, News, Sports, Kids, Premiere)
type GORMItemMediaType struct {
	Id        uint   `gorm:"column:Id;primaryKey;autoIncrement"`
	ItemId    string `gorm:"column:ItemId;not null;index"`
	MediaType string `gorm:"column:MediaType;not null"`
}

func (GORMItemMediaType) TableName() string {
	return "ItemMediaTypes"
}

// GORMMediaSource is the GORM model for MediaSources table
type GORMMediaSource struct {
	Id                           string `gorm:"column:Id;primaryKey"`
	ItemId                       string `gorm:"column:ItemId;not null"`
	Name                         string `gorm:"column:Name"`
	Type                         string `gorm:"column:Type"`
	Container                    string `gorm:"column:Container"`
	Size                         int64  `gorm:"column:Size"`
	Path                         string `gorm:"column:Path"`
	Protocol                     string `gorm:"column:Protocol"`
	Encoder                      bool   `gorm:"column:Encoder"`
	VideoCodec                   string `gorm:"column:VideoCodec"`
	AudioCodec                   string `gorm:"column:AudioCodec"`
	Format                       string `gorm:"column:Format"`
	Width                        int    `gorm:"column:Width"`
	Height                       int    `gorm:"column:Height"`
	RefFrames                    int    `gorm:"column:RefFrames"`
	VideoFramerate               string `gorm:"column:VideoFramerate"`
	VideoBitRate                 int    `gorm:"column:VideoBitRate"`
	AudioBitRate                 int    `gorm:"column:AudioBitRate"`
	AudioChannels                int    `gorm:"column:AudioChannels"`
	AudioSampleRate              string `gorm:"column:AudioSampleRate"`
	DefaultAudioStreamIndex       int    `gorm:"column:DefaultAudioStreamIndex"`
	SupportsTranscoding           bool   `gorm:"column:SupportsTranscoding"`
	SupportsDirectStream          bool   `gorm:"column:SupportsDirectStream"`
	SupportsDirectPlay            bool   `gorm:"column:SupportsDirectPlay"`
	IsRemote                     bool   `gorm:"column:IsRemote"`
}

func (GORMMediaSource) TableName() string {
	return "MediaSources"
}

// GORMUser is the GORM model for Users table
type GORMUser struct {
	Id                           string    `gorm:"column:Id;primaryKey"`
	Name                         string    `gorm:"column:Name;not null"`
	Username                     string    `gorm:"column:Username"`
	EmailAddress                 string    `gorm:"column:EmailAddress"`
	LoginUsername                string    `gorm:"column:LoginUsername"`
	LoginPassword                string    `gorm:"column:LoginPassword"`
	InvalidLoginAttemptCount     int       `gorm:"column:InvalidLoginAttemptCount"`
	LastLoginDate                time.Time `gorm:"column:LastLoginDate"`
	LastActivityDate             time.Time `gorm:"column:LastActivityDate"`
	AuthenticationProviderID     string    `gorm:"column:AuthenticationProviderID"`
	PrimaryImageTag              string    `gorm:"column:PrimaryImageTag"`
	Policy                       string    `gorm:"column:Policy"`
	HasConfiguredPassword       bool      `gorm:"-"`
	HasConfiguredEasyPassword    bool      `gorm:"-"`
	ConnectUserName              string    `gorm:"-"`
}

func (GORMUser) TableName() string {
	return "Users"
}

// GORMUserItem is the GORM model for UserItems table
type GORMUserItem struct {
	Id                       string `gorm:"column:Id;primaryKey"`
	UserId                    string `gorm:"column:UserId;not null"`
	ItemID                    string `gorm:"column:ItemID;not null"`
	PlaybackPositionTicks     int64  `gorm:"column:PlaybackPositionTicks"`
	PlayCount                 int    `gorm:"column:PlayCount"`
	IsFavorite                bool   `gorm:"column:IsFavorite"`
	Liked                     bool   `gorm:"column:Liked"`
	LastPlayedDate            string `gorm:"column:LastPlayedDate"`
	Played                    bool   `gorm:"column:Played"`
	Rating                    int    `gorm:"column:Rating"`
}

func (GORMUserItem) TableName() string {
	return "UserItems"
}

//GORMSession is the GORM model for Sessions table
type GORMSession struct {
	Id                             string    `gorm:"column:Id;primaryKey"`
	Client                         string    `gorm:"column:Client"`
	DeviceName                     string    `gorm:"column:DeviceName"`
	DisplayName                    string    `gorm:"column:DisplayName"`
	Endpoint                       string    `gorm:"column:Endpoint"`
	LocalAddress                   string    `gorm:"column:LocalAddress"`
	RemoteAddress                  string    `gorm:"column:RemoteAddress"`
	MachineId                      string    `gorm:"column:MachineId"`
	LastActivityTime               time.Time `gorm:"column:LastActivityTime"`
	LastPlaybackTime               time.Time `gorm:"column:LastPlaybackTime"`
	PlaybackPositionTicks          int64     `gorm:"column:PlaybackPositionTicks"`
	PlayMethod                    string    `gorm:"column:PlayMethod"`
	SupportsMediaControl          bool      `gorm:"column:SupportsMediaControl"`
	SupportsPersistentIdentification bool   `gorm:"column:SupportsPersistentIdentification"`
	SupportsSync                   bool     `gorm:"column:SupportsSync"`
	IsInActiveSession             bool      `gorm:"column:IsInActiveSession"`
	IsTerminal                    bool      `gorm:"column:IsTerminal"`
	StartTimeTicks                int64     `gorm:"column:StartTimeTicks"`
}

func (GORMSession) TableName() string {
	return "Sessions"
}

// GORMTimer represents a TV recording timer.
type GORMTimer struct {
	Id          string    `gorm:"column:Id;primaryKey"`
	ChannelId   string    `gorm:"column:ChannelId"`
	ProgramId   string    `gorm:"column:ProgramId"`
	Name        string    `gorm:"column:Name"`
	StartTime   time.Time `gorm:"column:StartTime"`
	EndTime     time.Time `gorm:"column:EndTime"`
	PrePadding  int       `gorm:"column:PrePadding"`
	PostPadding int       `gorm:"column:PostPadding"`
	Status     string    `gorm:"column:Status"`
	RecordingId string    `gorm:"column:RecordingId"`
	CreatedDate time.Time `gorm:"column:CreatedDate"`
}

func (GORMTimer) TableName() string {
	return "Timers"
}

// GORMSeriesTimer represents a series recording timer.
type GORMSeriesTimer struct {
	Id                string    `gorm:"column:Id;primaryKey"`
	ChannelId         string    `gorm:"column:ChannelId"`
	ProgramName       string    `gorm:"column:ProgramName"`
	StartTime         time.Time `gorm:"column:StartTime"`
	EndTime           time.Time `gorm:"column:EndTime"`
	PrePadding       int       `gorm:"column:PrePadding"`
	PostPadding       int       `gorm:"column:PostPadding"`
	Days             string    `gorm:"column:Days"`
	RecordAnyTime    bool      `gorm:"column:RecordAnyTime"`
	RecordAnyChannel bool      `gorm:"column:RecordAnyChannel"`
	RecordNewOnly    bool      `gorm:"column:RecordNewOnly"`
	Status           string    `gorm:"column:Status"`
	CreatedDate      time.Time `gorm:"column:CreatedDate"`
}

func (GORMSeriesTimer) TableName() string {
	return "SeriesTimers"
}

// GORMRecording represents a TV recording.
type GORMRecording struct {
	Id            string    `gorm:"column:Id;primaryKey"`
	Name          string    `gorm:"column:Name"`
	ChannelId     string    `gorm:"column:ChannelId"`
	ProgramId     string    `gorm:"column:ProgramId"`
	StartTime     time.Time `gorm:"column:StartTime"`
	EndTime       time.Time `gorm:"column:EndTime"`
	Status        string    `gorm:"column:Status"`
	Format        string    `gorm:"column:Format"`
	FileSize      int64     `gorm:"column:FileSize"`
	ItemId        string    `gorm:"column:ItemId"`
	SeriesTimerId string    `gorm:"column:SeriesTimerId"`
	CreatedDate   time.Time `gorm:"column:CreatedDate"`
}

func (GORMRecording) TableName() string {
	return "Recordings"
}

// GORMTunerHost represents a TV tuner host.
type GORMTunerHost struct {
	Id          string `gorm:"column:Id;primaryKey"`
	Type        string `gorm:"column:Type"`
	Host        string `gorm:"column:Host"`
	Port        int    `gorm:"column:Port"`
	TunerIp     string `gorm:"column:TunerIp"`
	FriendlyName string `gorm:"column:FriendlyName"`
	Enabled     bool   `gorm:"column:Enabled"`
}

func (GORMTunerHost) TableName() string {
	return "TunerHosts"
}

// GORMListingProvider represents an EPG listing provider.
type GORMListingProvider struct {
	Id       string `gorm:"column:Id;primaryKey"`
	Type     string `gorm:"column:Type"`
	Username string `gorm:"column:Username"`
	Password string `gorm:"column:Password"`
	Country  string `gorm:"column:Country"`
	ZipCode  string `gorm:"column:ZipCode"`
	Enabled  bool   `gorm:"column:Enabled"`
}

func (GORMListingProvider) TableName() string {
	return "ListingProviders"
}

// AuthenticationToken represents an API key/token stored in the local database
// Mirrors the C# AuthenticationInfo class from AuthenticationRepository.cs
type AuthenticationToken struct {
	AccessToken       string    `gorm:"column:AccessToken;primaryKey"`
	DeviceID          string    `gorm:"column:DeviceId"`
	AppName           string    `gorm:"column:AppName"`
	AppVersion        string    `gorm:"column:AppVersion"`
	DeviceName        string    `gorm:"column:DeviceName"`
	UserID           string    `gorm:"column:UserId"`
	UserName         string    `gorm:"column:UserName"`
	DateCreated      time.Time `gorm:"column:DateCreated"`
	DateLastActivity time.Time `gorm:"column:DateLastActivity"`
}

func (AuthenticationToken) TableName() string {
	return "AuthenticationTokens"
}