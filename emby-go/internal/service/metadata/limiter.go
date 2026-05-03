package metadata

import (
	"fmt"
	"net/http"
	"time"
)

// RateLimiter controls API request rates.
type RateLimiter struct {
	maxRequests int
	window      time.Duration
	requests    map[string][]time.Time
	mu          interface{}
}

// NewRateLimiter creates a new rate limiter.
func NewRateLimiter(maxRequests int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		maxRequests: maxRequests,
		window:      window,
		requests:    make(map[string][]time.Time),
	}
}

// Allow checks if a request is allowed.
func (rl *RateLimiter) Allow(key string) bool {
	now := time.Now()
	windowStart := now.Add(-rl.window)

	// Clean old requests
	if _, exists := rl.requests[key]; !exists {
		rl.requests[key] = []time.Time{}
	}

	var validRequests []time.Time
	for _, t := range rl.requests[key] {
		if t.After(windowStart) {
			validRequests = append(validRequests, t)
		}
	}

	if len(validRequests) >= rl.maxRequests {
		return false
	}

	rl.requests[key] = append(validRequests, now)
	return true
}

// WaitTime returns the time to wait before the next request.
func (rl *RateLimiter) WaitTime(key string) time.Duration {
	if len(rl.requests[key]) == 0 {
		return 0
	}

	windowStart := time.Now().Add(-rl.window)
	for _, t := range rl.requests[key] {
		if t.After(windowStart) {
			return rl.window - time.Since(t)
		}
	}

	return 0
}

// Clear clears all rate limiter data.
func (rl *RateLimiter) Clear() {
	rl.requests = make(map[string][]time.Time)
}

// HTTPClient wraps http.Client with rate limiting.
type HTTPClient struct {
	client    *http.Client
	limiter   *RateLimiter
	key       string
}

// NewHTTPClient creates a new rate-limited HTTP client.
func NewHTTPClient(limiter *RateLimiter, key string) *HTTPClient {
	return &HTTPClient{
		client:  &http.Client{Timeout: 30 * time.Second},
		limiter: limiter,
		key:     key,
	}
}

// Get performs a GET request with rate limiting.
func (hc *HTTPClient) Get(url string) (*http.Response, error) {
	if !hc.limiter.Allow(hc.key) {
		return nil, fmt.Errorf("rate limit exceeded")
	}

	return hc.client.Get(url)
}
