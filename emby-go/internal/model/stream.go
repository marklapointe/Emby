package model

// MediaStream represents a stream within a media file.
type MediaStream struct {
	Index          int     `json:"Index"`
	Codecs         string  `json:"Codecs"`
	Language       string  `json:"Language,omitempty"`
	DisplayLanguage string `json:"DisplayLanguage,omitempty"`
	Width          int     `json:"Width,omitempty"`
	Height         int     `json:"Height,omitempty"`
	FrameRate      float64 `json:"FrameRate,omitempty"`
	BitRate        int     `json:"BitRate,omitempty"`
	IsDefault      bool    `json:"IsDefault"`
	IsForced       bool    `json:"IsForced"`
	IsExternal     bool    `json:"IsExternal"`
	Type           string  `json:"Type"`
	Codec          string  `json:"Codec,omitempty"`
	Profile        string  `json:"Profile,omitempty"`
	Level          int     `json:"Level,omitempty"`
	ChannelLayout  string  `json:"ChannelLayout,omitempty"`
	SampleRate     int     `json:"SampleRate,omitempty"`
}
