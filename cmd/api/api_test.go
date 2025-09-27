package main

import (
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Config struct {
	RequestsPerTimeFrame int
	TimeFrame            time.Duration
	Enabled              bool
}

type RateLimiter struct {
	cfg     Config
	clients map[string]*clientData
	mu      sync.Mutex
}

type clientData struct {
	count      int
	expiration time.Time
}

func New(cfg Config) *RateLimiter {
	return &RateLimiter{
		cfg:     cfg,
		clients: make(map[string]*clientData),
	}
}

// Middleware для HTTP
func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	if !rl.cfg.Enabled {
		return next
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getClientIP(r)

		rl.mu.Lock()
		defer rl.mu.Unlock()

		data, exists := rl.clients[ip]
		now := time.Now()

		if !exists || now.After(data.expiration) {
			// новое окно
			rl.clients[ip] = &clientData{
				count:      1,
				expiration: now.Add(rl.cfg.TimeFrame),
			}
		} else {
			data.count++
			if data.count > rl.cfg.RequestsPerTimeFrame {
				w.WriteHeader(http.StatusTooManyRequests)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

// Извлекаем IP из X-Forwarded-For или RemoteAddr
func getClientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return strings.Split(xff, ",")[0]
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
