package middleware

import (
	"clean-arsitektur/pkg"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

var visitors = make(map[string]*rate.Limiter)
var mu sync.Mutex

func limiter(ip string, rps, burst int) *rate.Limiter {
	var limit = time.Duration(rps) * time.Second

	mu.Lock()
	defer mu.Unlock()

	limiter, exists := visitors[ip]
	if !exists {
		limiter = rate.NewLimiter(rate.Every(limit), burst)
		visitors[ip] = limiter
	}

	return limiter
}

func RateLimiter(sec, burst int, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := pkg.ClientIP(r)
		limiter := limiter(ip, sec, burst)

		if !limiter.Allow() {
			return
		}

		next(w, r)
	}
}
