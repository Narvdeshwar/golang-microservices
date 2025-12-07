package routes

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ReverseProxy(target string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		remote, _ := url.Parse(target)
		proxy := httputil.NewSingleHostReverseProxy(remote)
		ctx.Request.URL.Path = ctx.Param("any")
		ctx.Request.Host = remote.Host
		ctx.Request.URL.Host = remote.Host
		ctx.Request.URL.Scheme = remote.Scheme
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
