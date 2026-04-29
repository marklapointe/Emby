package middleware

import (
	"context"
	"net/http"
	"strings"
)

// ContextKey is the type for context keys.
type ContextKey string

const (
	// UserIDKey is the context key for user ID.
	UserIDKey ContextKey = "userID"
	// SessionIDKey is the context key for session ID.
	SessionIDKey ContextKey = "sessionID"
	// UserKey is the context key for user object.
	UserKey ContextKey = "user"
)

// AuthMiddleware handles API key and session authentication.
type AuthMiddleware struct{}

// NewAuthMiddleware creates a new authentication middleware.
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

// Handle authenticates requests via API key or session token.
func (m *AuthMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip auth for public endpoints
		if isPublicEndpoint(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		// Check for API key in header
		apiKey := r.Header.Get("X-Emby-Token")
		if apiKey != "" {
			// TODO: Validate API key against database
			ctx := context.WithValue(r.Context(), UserIDKey, "api-key-user")
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		// Check for session token in header or cookie
		sessionToken := r.Header.Get("X-Emby-Device-Id")
		if sessionToken == "" {
			sessionToken = r.Header.Get("X-Emby-Client")
		}
		if sessionToken != "" {
			// TODO: Validate session token
			ctx := context.WithValue(r.Context(), SessionIDKey, sessionToken)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		// No valid authentication
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

// isPublicEndpoint checks if a path is public (no auth required).
func isPublicEndpoint(path string) bool {
	publicPaths := []string{
		"/health",
		"/System/Info/Public",
		"/Users/Public",
		"/PluginInfo",
	}
	for _, p := range publicPaths {
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}
