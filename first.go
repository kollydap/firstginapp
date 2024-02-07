package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		duration := time.Since(start)
		log.Printf("Request - Method: %s | Status: %d | Duration: %v", ctx.Request.Method, ctx.Writer.Status(), duration)
	}
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("X-API-Key")
		if apiKey == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorizes"})
			return
		}
		ctx.Next()
	}
}
func main() {
	// Your Gin application code goes here
	router := gin.Default()
	router.Use(LoggerMiddleWare())
	authGroup := router.Group("/api")

	authGroup.Use(AuthMiddleWare())
	{
		authGroup.GET("/data", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Authenticated and authorized!"})
		})
		authGroup.GET("/user/:id", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": ctx.Param("id")})
		})
		authGroup.GET("/search", func(ctx *gin.Context) {
			query := ctx.DefaultQuery("name", "Default-value")
			ctx.String(200, "Search query: "+query)
		})
	}
	public := router.Group("/public")
	{
		public.GET("/info", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "public Info"})
		})
	}

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello World!")
	})
	router.GET("/bye", func(ctx *gin.Context) {
		ctx.String(200, "boss")
	})

	router.Run(":8080")
}
