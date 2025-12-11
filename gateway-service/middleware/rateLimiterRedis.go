package middleware

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/ulule/limiter/v3"
	store "github.com/ulule/limiter/v3/drivers/store/redis"
	ginlimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
)

func RedisRateLimiter() gin.HandlerFunc {
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  60,
	}

	client := redis.NewClient(&redis.Options{Addr: "redis:6379"})

	redisStore, err := store.NewStoreWithOptions(client, limiter.StoreOptions{})
	
	if err!=nil{
		panic(err)
	}
	limiterInstance:=limiter.New(redisStore,rate)
	return ginlimiter.NewMiddleware(limiterInstance)
}
