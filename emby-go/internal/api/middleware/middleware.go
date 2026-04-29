package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

// CORSMiddleware returns CORS middleware.
func CORSMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, X-Emby-Token, X-Emby-Client, X-Emby-Device-Id, X-Emby-Device-Name, X-Emby-Client-Version")
			w.Header().Set("Access-Control-Max-Age", "86400")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// RequestLogger returns a request logging middleware.
func RequestLogger(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Wrap response writer to capture status code
			sw := &statusWriter{ResponseWriter: w, statusCode: http.StatusOK}

			next.ServeHTTP(sw, r)

			duration := time.Since(start)
			logger.Info("request",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Int("status", sw.statusCode),
				zap.Duration("duration", duration),
				zap.String("remote_addr", r.RemoteAddr),
				zap.String("user_agent", r.UserAgent()),
			)
		})
	}
}

// statusWriter wraps http.ResponseWriter to capture status code.
type statusWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code.
func (sw *statusWriter) WriteHeader(code int) {
	sw.statusCode = code
	sw.ResponseWriter.WriteHeader(code)
}

// LoggingMiddleware logs HTTP requests.
func LoggingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			path := r.URL.Path

			// Skip health check logs
			if path == "/health" || path == "/ping" {
				next.ServeHTTP(w, r)
				return
			}

			// Wrap response writer
			sw := &statusWriter{ResponseWriter: w, statusCode: http.StatusOK}

			next.ServeHTTP(sw, r)

			duration := time.Since(start)
			logger.Info("HTTP request",
				zap.String("method", r.Method),
				zap.String("path", path),
				zap.Int("status", sw.statusCode),
				zap.Duration("duration", duration),
			)
		})
	}
}

// AuthMiddleware returns authentication middleware.
func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check for API key in header or query parameter
			apiKey := r.Header.Get("X-Emby-Token")
			if apiKey == "" {
				apiKey = r.URL.Query().Get("api_key")
			}

			// For now, allow all requests (authentication will be implemented later)
			next.ServeHTTP(w, r)
		})
	}
}

// SessionMiddleware returns session middleware.
func SessionMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract session from cookie or header
			// For now, pass through (session management will be implemented later)
			next.ServeHTTP(w, r)
		})
	}
}

// CompressionMiddleware returns compression middleware.
func CompressionMiddleware() func(http.Handler) http.Handler {
	return middleware.Compress(8)
}

// RateLimitMiddleware returns rate limiting middleware.
func RateLimitMiddleware(maxRequests int, window time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Simple rate limiting (will be enhanced later)
			next.ServeHTTP(w, r)
		})
	}
}

// SecurityMiddleware returns security headers middleware.
func SecurityMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("X-Frame-Options", "DENY")
			w.Header().Set("X-XSS-Protection", "1; mode=block")
			w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")

			next.ServeHTTP(w, r)
		})
	}
}
