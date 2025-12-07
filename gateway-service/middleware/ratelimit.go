package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mGin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimiter() gin.HandlerFunc {
	// 60 requests per minute per IP
	rate := limiter.Rate{
		Period: 60 * 1e9, // 60 seconds
		Limit:  60,
	}

	// In-memory storage
	store := memory.NewStore()

	// Create limiter
	instance := limiter.New(store, rate)

	// Gin middleware
	return mGin.NewMiddleware(instance)
}
