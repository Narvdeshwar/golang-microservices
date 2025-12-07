package routes

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ReverseProxy(target string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		url, _ := url.Parse(target)
		proxy := httputil.NewSingleHostReverseProxy(url)
		ctx.Request.URL.Host = url.Host
		ctx.Request.URL.Scheme = url.Scheme
		ctx.Request.Host = url.Host
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
