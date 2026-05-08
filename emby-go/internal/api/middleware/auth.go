package middleware

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
