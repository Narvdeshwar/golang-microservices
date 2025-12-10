package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mGin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimiter() gin.HandlerFunc {
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  60,
	}

	store := memory.NewStore() // Uses system RAM which may choke the system

	instance := limiter.New(store, rate)

	return mGin.NewMiddleware(instance)
}
