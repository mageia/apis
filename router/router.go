package router

import (
	"strings"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(func(c *gin.Context) {
		if c.Request.ProtoMajor == 2 && strings.HasPrefix(c.GetHeader("Content-Type"), "application/grpc") {
			grpc.NewServer().ServeHTTP(c.Writer, c.Request)
			c.Abort()
			return
		}
	})
	r.Use(func(c *gin.Context) {
		if len(c.Errors) != 0 {
			c.AbortWithStatusJSON(400, gin.H{"message": c.Errors.String()})
			return
		}
		c.Next()
	})

	r.GET("/healthz", func(c *gin.Context) { c.String(200, "ok") })

	api := r.Group("/api/v1")

	api.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "test"})
	})

	return r
}
