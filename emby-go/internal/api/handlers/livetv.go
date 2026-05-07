package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/emby/emby-go/internal/repository"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// LiveTVHandler handles Live TV-related API endpoints.
type LiveTVHandler struct {
	repo   *repository.ItemRepository
	logger *zap.Logger
}

// NewLiveTVHandler creates a new Live TV handler.
func NewLiveTVHandler(repo *repository.ItemRepository, logger *zap.Logger) *LiveTVHandler {
	return &LiveTVHandler{
		repo:   repo,
		logger: logger,
	}
}

// GetChannels handles GET /LiveTv/Channels
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

// GetPrograms handles GET /LiveTv/Programs
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

// GetProgram handles GET /LiveTv/Programs/{id}
func (h *LiveTVHandler) GetProgram(w http.ResponseWriter, r *http.Request) {
	programId := chi.URLParam(r, "id")

	program, err := h.repo.GetProgram(programId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(program)
}

// GetRecordings handles GET /LiveTv/Recordings
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

// GetRecording handles GET /LiveTv/Recordings/{id}
func (h *LiveTVHandler) GetRecording(w http.ResponseWriter, r *http.Request) {
	recordingId := chi.URLParam(r, "id")

	recording, err := h.repo.GetRecording(recordingId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recording)
}

// GetTimers handles GET /LiveTv/Timers
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

// CreateTimer handles POST /LiveTv/Timers
func (h *LiveTVHandler) CreateTimer(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = req

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "timer created"})
}

// GetGuideInfo handles GET /LiveTv/GuideInfo
func (h *LiveTVHandler) GetGuideInfo(w http.ResponseWriter, r *http.Request) {
	guideInfo := map[string]interface{}{
		"HasImage":      false,
		"MinProgramDate": "2026-04-29T00:00:00Z",
		"MaxProgramDate": "2026-05-06T00:00:00Z",
		"HasChannels":   true,
		"HasPrograms":   true,
		"HasRecordings": true,
		"HasTimers":     true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(guideInfo)
}

// GetChannelsWithImage handles GET /LiveTv/ChannelsWithImage
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

// GetProgramWithImage handles GET /LiveTv/ProgramsWithImage
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

// GetRecordingFolders handles GET /LiveTv/RecordingFolders
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

// GetRecommendedPrograms handles GET /LiveTv/RecommendedPrograms
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

func (h *LiveTVHandler) GetSeriesTimers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{})
}

func (h *LiveTVHandler) DeleteTunerHost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (h *LiveTVHandler) GetTunerHostTypes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *LiveTVHandler) GetListingProviders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *LiveTVHandler) CreateListingProvider(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{})
}

func (h *LiveTVHandler) GetDefaultListingProvider(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{})
}

func (h *LiveTVHandler) GetSchedulesDirectCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *LiveTVHandler) CreateChannelMapping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{})
}

func (h *LiveTVHandler) GetChannelMappingOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{})
}
