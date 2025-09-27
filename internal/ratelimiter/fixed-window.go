package ratelimiter

import (
	"sync"
	"time"
)

type FixedWindowRateLimiter struct {
	sync.Mutex
	clients map[string]struct {
		count     int
		windowEnd time.Time
	}
	limit  int
	window time.Duration
}

func NewFixedWindowLimiter(limit int, window time.Duration) *FixedWindowRateLimiter {
	return &FixedWindowRateLimiter{
		clients: make(map[string]struct {
			count     int
			windowEnd time.Time
		}),
		limit:  limit,
		window: window,
	}
}

// Allow возвращает true, если запрос разрешен, и false + оставшееся время, если нет
func (rl *FixedWindowRateLimiter) Allow(ip string) (bool, time.Duration) {
	rl.Lock()
	defer rl.Unlock()

	now := time.Now()
	client, exists := rl.clients[ip]

	if !exists || now.After(client.windowEnd) {
		// новое окно для клиента
		rl.clients[ip] = struct {
			count     int
			windowEnd time.Time
		}{count: 1, windowEnd: now.Add(rl.window)}
		return true, 0
	}

	if client.count < rl.limit {
		// увеличиваем счетчик в текущем окне
		client.count++
		rl.clients[ip] = client
		return true, 0
	}

	// лимит достигнут, возвращаем оставшееся время
	retryAfter := client.windowEnd.Sub(now)
	return false, retryAfter
}
