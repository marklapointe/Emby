package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emby/emby-go/internal/repository"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func TestNewLiveTVHandler(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	itemRepo := &repository.ItemRepository{}
	h := NewLiveTVHandler(itemRepo, logger)
	if h == nil {
		t.Fatal("NewLiveTVHandler returned nil")
	}
	if h.repo == nil {
		t.Error("repo should not be nil")
	}
	if h.logger == nil {
		t.Error("logger should not be nil")
	}
}

func TestGetSeriesTimers(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	itemRepo := &repository.ItemRepository{}
	h := NewLiveTVHandler(itemRepo, logger)

	req := httptest.NewRequest("GET", "/LiveTv/SeriesTimers", nil)
	w := httptest.NewRecorder()

	h.GetSeriesTimers(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var result interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
}

func TestGetTimerProviders(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	itemRepo := &repository.ItemRepository{}
	h := NewLiveTVHandler(itemRepo, logger)

	req := httptest.NewRequest("GET", "/LiveTv/TimerProviders", nil)
	w := httptest.NewRecorder()

	h.GetTimerProviders(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestGetTunerHosts(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	itemRepo := &repository.ItemRepository{}
	h := NewLiveTVHandler(itemRepo, logger)

	req := httptest.NewRequest("GET", "/LiveTv/TunerHosts", nil)
	w := httptest.NewRecorder()

	h.GetTunerHosts(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestGetTunerHost_WithChi(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	itemRepo := &repository.ItemRepository{}
	h := NewLiveTVHandler(itemRepo, logger)

	r := chi.NewRouter()
	r.Get("/LiveTv/TunerHosts/{id}", h.GetTunerHost)

	req := httptest.NewRequest("GET", "/LiveTv/TunerHosts/tuner-123", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if result["Id"] != "tuner-123" {
		t.Errorf("expected Id 'tuner-123', got %v", result["Id"])
	}
}

func TestCreateTunerHost(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	itemRepo := &repository.ItemRepository{}
	h := NewLiveTVHandler(itemRepo, logger)

	req := httptest.NewRequest("POST", "/LiveTv/TunerHosts", nil)
	w := httptest.NewRecorder()

	h.CreateTunerHost(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestDeleteTunerHost_WithChi(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	itemRepo := &repository.ItemRepository{}
	h := NewLiveTVHandler(itemRepo, logger)

	r := chi.NewRouter()
	r.Delete("/LiveTv/TunerHosts/{id}", h.DeleteTunerHost)

	req := httptest.NewRequest("DELETE", "/LiveTv/TunerHosts/tuner-123", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("expected status 204, got %d", w.Code)
	}
}

func TestGetTunerHostTypes(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	itemRepo := &repository.ItemRepository{}
	h := NewLiveTVHandler(itemRepo, logger)

	req := httptest.NewRequest("GET", "/LiveTv/TunerHosts/Types", nil)
	w := httptest.NewRecorder()

	h.GetTunerHostTypes(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestGetListingProviders(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	itemRepo := &repository.ItemRepository{}
	h := NewLiveTVHandler(itemRepo, logger)

	req := httptest.NewRequest("GET", "/LiveTv/ListingProviders", nil)
	w := httptest.NewRecorder()

	h.GetListingProviders(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestCreateListingProvider(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	itemRepo := &repository.ItemRepository{}
	h := NewLiveTVHandler(itemRepo, logger)

	req := httptest.NewRequest("POST", "/LiveTv/ListingProviders", nil)
	w := httptest.NewRecorder()

	h.CreateListingProvider(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestGetDefaultListingProvider(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	itemRepo := &repository.ItemRepository{}
	h := NewLiveTVHandler(itemRepo, logger)

	req := httptest.NewRequest("GET", "/LiveTv/ListingProviders/Default", nil)
	w := httptest.NewRecorder()

	h.GetDefaultListingProvider(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestGetSchedulesDirectCountries(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	itemRepo := &repository.ItemRepository{}
	h := NewLiveTVHandler(itemRepo, logger)

	req := httptest.NewRequest("GET", "/LiveTv/ListingProviders/SchedulesDirect/Countries", nil)
	w := httptest.NewRecorder()

	h.GetSchedulesDirectCountries(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestCreateChannelMapping(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	itemRepo := &repository.ItemRepository{}
	h := NewLiveTVHandler(itemRepo, logger)

	req := httptest.NewRequest("POST", "/LiveTv/ChannelMappings", nil)
	w := httptest.NewRecorder()

	h.CreateChannelMapping(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestGetChannelMappingOptions(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	itemRepo := &repository.ItemRepository{}
	h := NewLiveTVHandler(itemRepo, logger)

	req := httptest.NewRequest("GET", "/LiveTv/ChannelMappingOptions", nil)
	w := httptest.NewRecorder()

	h.GetChannelMappingOptions(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}