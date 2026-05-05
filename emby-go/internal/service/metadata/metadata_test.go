package metadata

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewTMDbProvider(t *testing.T) {
	p := NewTMDbProvider("test-api-key")
	if p == nil {
		t.Fatal("NewTMDbProvider returned nil")
	}
	if p.apiKey != "test-api-key" {
		t.Errorf("expected apiKey 'test-api-key', got '%s'", p.apiKey)
	}
	if p.httpClient == nil {
		t.Error("httpClient is nil")
	}
}

func TestTMDbProvider_GetImageURL(t *testing.T) {
	p := NewTMDbProvider("test-api-key")
	url := p.GetImageURL("/poster.jpg", "w500")
	expected := "https://image.tmdb.org/t/p/w500/poster.jpg"
	if url != expected {
		t.Errorf("expected '%s', got '%s'", expected, url)
	}
}

func TestNewTVDbProvider(t *testing.T) {
	p := NewTVDbProvider("test-api-key")
	if p == nil {
		t.Fatal("NewTVDbProvider returned nil")
	}
	if p.apiKey != "test-api-key" {
		t.Errorf("expected apiKey 'test-api-key', got '%s'", p.apiKey)
	}
}

func TestTVDbProvider_GetImageURL(t *testing.T) {
	p := NewTVDbProvider("test-api-key")
	url := p.GetImageURL("poster", "/banners/posters/123.jpg")
	expected := "https://artworks.thetvdb.com/banners//banners/posters/123.jpg"
	if url != expected {
		t.Errorf("expected '%s', got '%s'", expected, url)
	}
}

func TestTMDbProvider_SearchMovie_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	p := &TMDbProvider{
		apiKey: "test",
		httpClient: server.Client(),
	}

	_, err := p.SearchMovie("test")
	if err == nil {
		t.Error("expected error for HTTP 500")
	}
}

func TestTVDbProvider_SearchSeries_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	p := &TVDbProvider{
		apiKey: "test",
		httpClient: server.Client(),
	}

	_, err := p.SearchSeries("test")
	if err == nil {
		t.Error("expected error for HTTP 500")
	}
}

func TestTVDbProvider_GetSeries_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	p := &TVDbProvider{
		apiKey: "test",
		httpClient: server.Client(),
	}

	_, err := p.GetSeries(0)
	if err == nil {
		t.Error("expected error for HTTP 404")
	}
}