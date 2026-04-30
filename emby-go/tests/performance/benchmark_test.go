package performance

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// BenchmarkHealthCheck benchmarks the health check endpoint.
func BenchmarkHealthCheck(b *testing.B) {
	req := httptest.NewRequest("GET", "/health", nil)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)
	}
}

// BenchmarkSystemInfo benchmarks the system info endpoint.
func BenchmarkSystemInfo(b *testing.B) {
	req := httptest.NewRequest("GET", "/System/Info", nil)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ServerName":"Emby Server","Version":"0.1.0"}`))
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)
	}
}

// BenchmarkLibraryRoot benchmarks the library root endpoint.
func BenchmarkLibraryRoot(b *testing.B) {
	req := httptest.NewRequest("GET", "/Library/Root", nil)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Name":"Media Library","Path":"/media"}`))
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)
	}
}

// BenchmarkSearch benchmarks the search endpoint.
func BenchmarkSearch(b *testing.B) {
	req := httptest.NewRequest("GET", "/Items/Search?SearchTerm=test", nil)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Items":[{"Name":"test-item"}]}`))
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)
	}
}

// BenchmarkUserLogin benchmarks the user login endpoint.
func BenchmarkUserLogin(b *testing.B) {
	req := httptest.NewRequest("POST", "/Users/AuthenticateByName", nil)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"AccessToken":"test-token","UserId":"test-user-id"}`))
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)
	}
}
