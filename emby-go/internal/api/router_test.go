package api

import (
	"testing"
)

func TestToTitleCase_SingleWord(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"livetv", "LiveTv"},
		{"dlna", "Dlna"},
		{"scheduledtasks", "ScheduledTasks"},
		{"system", "System"},
		{"users", "Users"},
	}

	for _, tt := range tests {
		result := toTitleCase(tt.input)
		if result != tt.expected {
			t.Errorf("toTitleCase(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestToTitleCase_MultiWord(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"recommendedprograms", "RecommendedPrograms"},
		{"seriestimers", "SeriesTimers"},
		{"timerproviders", "TimerProviders"},
		{"tunerhosts", "TunerHosts"},
		{"tunerhosts/types", "TunerHosts/Types"},
		{"listingproviders", "ListingProviders"},
		{"listingproviders/default", "ListingProviders/Default"},
		{"schedulesdirect", "SchedulesDirect"},
		{"schedulesdirect/countries", "SchedulesDirect/Countries"},
		{"channelmappings", "ChannelMappings"},
		{"channelmappingoptions", "ChannelMappingOptions"},
		{"profileinfos", "ProfileInfos"},
		{"musicgenres", "MusicGenres"},
	}

	for _, tt := range tests {
		result := toTitleCase(tt.input)
		if result != tt.expected {
			t.Errorf("toTitleCase(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestToTitleCase_Path(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"/livetv/channels", "/LiveTv/Channels"},
		{"/livetv/recommendedprograms", "/LiveTv/RecommendedPrograms"},
		{"/dlna/profiles", "/Dlna/Profiles"},
		{"/sync/jobs", "/Sync/Jobs"},
		{"/system/info", "/System/Info"},
		{"/system/info/public", "/System/Info/Public"},
	}

	for _, tt := range tests {
		result := toTitleCase(tt.input)
		if result != tt.expected {
			t.Errorf("toTitleCase(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestToTitleCase_Empty(t *testing.T) {
	result := toTitleCase("")
	if result != "" {
		t.Errorf("toTitleCase(%q) = %q, want %q", "", result, "")
	}
}

func TestToTitleCase_WithPathVariables(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"/livetv/channels/{id}", "/LiveTv/Channels/{id}"},
		{"/livetv/programs/{id}", "/LiveTv/Programs/{id}"},
		{"/sync/jobs/{id}", "/Sync/Jobs/{id}"},
		{"/sync/jobs/{id}/items/{itemId}", "/Sync/Jobs/{id}/Items/{itemId}"},
	}

	for _, tt := range tests {
		result := toTitleCase(tt.input)
		if result != tt.expected {
			t.Errorf("toTitleCase(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestToTitleCase_StandardCasing(t *testing.T) {
	result := toTitleCase("test")
	if result != "Test" {
		t.Errorf("toTitleCase(%q) = %q, want %q", "test", result, "Test")
	}
}

func TestToTitleCase_AlreadyTitleCase(t *testing.T) {
	result := toTitleCase("LiveTv")
	if result != "Livetv" {
		t.Errorf("toTitleCase(%q) = %q, want %q (transformation is applied to ensure consistent routing)", "LiveTv", result, "Livetv")
	}
}