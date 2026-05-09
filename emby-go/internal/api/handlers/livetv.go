package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/emby/emby-go/internal/repository"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type LiveTVHandler struct {
	repo   *repository.ItemRepository
	logger *zap.Logger
}

func NewLiveTVHandler(repo *repository.ItemRepository, logger *zap.Logger) *LiveTVHandler {
	return &LiveTVHandler{
		repo:   repo,
		logger: logger,
	}
}

func (h *LiveTVHandler) GetLiveTvInfo(w http.ResponseWriter, r *http.Request) {
	info := map[string]interface{}{
		"EnableUsers":              true,
		"EnableProgramGuide":       true,
		"EnableRecordingScheduling": true,
		"EnableChannelRetrieval":   true,
		"EnableTunerDiscovery":     true,
		"EnabledMediaTypes":        []string{"Audio", "Video"},
		"SupportedServices":        []string{"m3u", "htsp"},
		"Version":                  "1.0.0.0",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func (h *LiveTVHandler) GetChannels(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")
	isFavorite := r.URL.Query().Get("IsFavorite")
	_ = isFavorite

	channels, err := h.repo.GetChannels(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(channels)
}

func (h *LiveTVHandler) GetChannel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	channel, err := h.repo.GetChannel(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if channel == nil {
		http.Error(w, "channel not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(channel)
}

func (h *LiveTVHandler) GetPrograms(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")
	isAiring := r.URL.Query().Get("IsAiring")
	isMovie := r.URL.Query().Get("IsMovie")
	isSports := r.URL.Query().Get("IsSports")
	isKids := r.URL.Query().Get("IsKids")
	isNews := r.URL.Query().Get("IsNews")
	isSeries := r.URL.Query().Get("IsSeries")
	startDate := r.URL.Query().Get("StartDate")
	endDate := r.URL.Query().Get("EndDate")
	imageTypeLimit, _ := strconv.Atoi(r.URL.Query().Get("ImageTypeLimit"))
	_ = imageTypeLimit

	_ = isAiring
	_ = isMovie
	_ = isSports
	_ = isKids
	_ = isNews
	_ = isSeries
	_ = startDate
	_ = endDate

	programs, err := h.repo.GetPrograms(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(programs)
}

func (h *LiveTVHandler) GetProgram(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	program, err := h.repo.GetProgram(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if program == nil {
		http.Error(w, "program not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(program)
}

func (h *LiveTVHandler) GetRecordings(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")

	recordings, err := h.repo.GetRecordings(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recordings)
}

func (h *LiveTVHandler) GetRecording(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	recording, err := h.repo.GetRecording(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if recording == nil {
		http.Error(w, "recording not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recording)
}

func (h *LiveTVHandler) DeleteRecording(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.repo.DeleteRecording(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *LiveTVHandler) GetRecordingSeries(w http.ResponseWriter, r *http.Request) {
	recordings, err := h.repo.GetRecordingSeries()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recordings)
}

func (h *LiveTVHandler) GetRecordingGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := h.repo.GetRecordingGroups()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}

func (h *LiveTVHandler) GetRecordingGroup(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	group, err := h.repo.GetRecordingGroup(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if group == nil {
		http.Error(w, "recording group not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}

func (h *LiveTVHandler) GetRecordingFolders(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")

	folders, err := h.repo.GetRecordingFolders(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folders)
}

func (h *LiveTVHandler) GetTimers(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")

	timers, err := h.repo.GetTimers(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timers)
}

func (h *LiveTVHandler) GetTimer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	timer, err := h.repo.GetTimer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if timer == nil {
		http.Error(w, "timer not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timer)
}

func (h *LiveTVHandler) GetDefaultTimer(w http.ResponseWriter, r *http.Request) {
	defaultTimer := map[string]interface{}{
		"PrePadding":       0,
		"PostPadding":      0,
		"RecordAnyTime":    false,
		"RecordAnyChannel": false,
		"RecordNewOnly":    true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(defaultTimer)
}

func (h *LiveTVHandler) CreateTimer(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ChannelID    string `json:"ChannelId"`
		ProgramID    string `json:"ProgramId"`
		StartDate    string `json:"StartDate"`
		EndDate      string `json:"EndDate"`
		PrePadding   int    `json:"PrePadding"`
		PostPadding  int    `json:"PostPadding"`
		Name         string `json:"Name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	timer, err := h.repo.CreateTimer(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(timer)
}

func (h *LiveTVHandler) UpdateTimer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req struct {
		ChannelID    string `json:"ChannelId"`
		ProgramID    string `json:"ProgramId"`
		StartDate    string `json:"StartDate"`
		EndDate      string `json:"EndDate"`
		PrePadding   int    `json:"PrePadding"`
		PostPadding  int    `json:"PostPadding"`
		Name         string `json:"Name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.repo.UpdateTimer(id, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *LiveTVHandler) DeleteTimer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.repo.DeleteTimer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *LiveTVHandler) GetSeriesTimers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *LiveTVHandler) GetSeriesTimer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	seriesTimer, err := h.repo.GetSeriesTimer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if seriesTimer == nil {
		http.Error(w, "series timer not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(seriesTimer)
}

func (h *LiveTVHandler) CreateSeriesTimer(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ChannelID        string `json:"ChannelId"`
		ProgramName     string `json:"ProgramName"`
		StartDate       string `json:"StartDate"`
		EndDate         string `json:"EndDate"`
		PrePadding      int    `json:"PrePadding"`
		PostPadding     int    `json:"PostPadding"`
		Days            []int  `json:"Days"`
		RecordAnyTime   bool   `json:"RecordAnyTime"`
		RecordAnyChannel bool   `json:"RecordAnyChannel"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	seriesTimer, err := h.repo.CreateSeriesTimer(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(seriesTimer)
}

func (h *LiveTVHandler) UpdateSeriesTimer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req struct {
		ChannelID        string `json:"ChannelId"`
		ProgramName     string `json:"ProgramName"`
		StartDate       string `json:"StartDate"`
		EndDate         string `json:"EndDate"`
		PrePadding      int    `json:"PrePadding"`
		PostPadding     int    `json:"PostPadding"`
		Days            []int  `json:"Days"`
		RecordAnyTime   bool   `json:"RecordAnyTime"`
		RecordAnyChannel bool   `json:"RecordAnyChannel"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.repo.UpdateSeriesTimer(id, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *LiveTVHandler) DeleteSeriesTimer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.repo.DeleteSeriesTimer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *LiveTVHandler) TunerReset(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	h.logger.Info("resetting tuner", zap.String("tunerId", id))
	w.WriteHeader(http.StatusNoContent)
}

func (h *LiveTVHandler) DiscoverTuners(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *LiveTVHandler) GetGuideInfo(w http.ResponseWriter, r *http.Request) {
	guideInfo := map[string]interface{}{
		"HasImage":       false,
		"MinProgramDate": "2026-04-29T00:00:00Z",
		"MaxProgramDate": "2026-05-06T00:00:00Z",
		"HasChannels":    true,
		"HasPrograms":    true,
		"HasRecordings":  true,
		"HasTimers":      true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(guideInfo)
}

func (h *LiveTVHandler) GetRecommendedPrograms(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")

	programs, err := h.repo.GetRecommendedPrograms(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(programs)
}

func (h *LiveTVHandler) GetChannelsWithImage(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")

	channels, err := h.repo.GetChannelsWithImage(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(channels)
}

func (h *LiveTVHandler) GetProgramWithImage(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")

	programs, err := h.repo.GetProgramsWithImage(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(programs)
}

func (h *LiveTVHandler) GetTimerProviders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *LiveTVHandler) GetTunerHosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *LiveTVHandler) GetTunerHost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"Id": id})
}

func (h *LiveTVHandler) CreateTunerHost(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Type     string `json:"Type"`
		Host     string `json:"Host"`
		Port     int    `json:"Port"`
		TunerIP  string `json:"TunerIp"`
		Friendly string `json:"FriendlyName"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tunerHost, err := h.repo.CreateTunerHost(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tunerHost)
}

func (h *LiveTVHandler) DeleteTunerHost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.repo.DeleteTunerHost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *LiveTVHandler) GetTunerHostTypes(w http.ResponseWriter, r *http.Request) {
	types := []map[string]string{
		{"Name": "M3U", "Type": "m3u", "Id": "m3u"},
		{"Name": "HTSP", "Type": "htsp", "Id": "htsp"},
		{"Name": "DVB", "Type": "dvb", "Id": "dvb"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(types)
}

func (h *LiveTVHandler) GetListingProviders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *LiveTVHandler) CreateListingProvider(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Type     string `json:"Type"`
		Username string `json:"Username"`
		Password string `json:"Password"`
		Country  string `json:"Country"`
		ZipCode  string `json:"ZipCode"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	provider, err := h.repo.CreateListingProvider(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(provider)
}

func (h *LiveTVHandler) DeleteListingProvider(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.repo.DeleteListingProvider(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *LiveTVHandler) GetDefaultListingProvider(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{})
}

func (h *LiveTVHandler) GetSchedulesDirectCountries(w http.ResponseWriter, r *http.Request) {
	countries := []map[string]interface{}{
		{"Name": "United States", "ShortCode": "USA"},
		{"Name": "Canada", "ShortCode": "CAN"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)
}

func (h *LiveTVHandler) GetChannelMappingOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{})
}

func (h *LiveTVHandler) CreateChannelMapping(w http.ResponseWriter, r *http.Request) {
	var req struct {
		TunerChannelNumber     string `json:"TunerChannelNumber"`
		ProviderChannelNumber string `json:"ProviderChannelNumber"`
		ProviderId           string `json:"ProviderId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.repo.CreateChannelMapping(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *LiveTVHandler) GetLiveStream(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	h.logger.Info("getting live stream", zap.String("streamId", id))

	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
}

func (h *LiveTVHandler) GetRecordingStream(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	h.logger.Info("getting recording stream", zap.String("recordingId", id))

	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
}