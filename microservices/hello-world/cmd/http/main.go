package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"microservices/hello-world/internal/handler"

	_ "microservices/hello-world/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Hello World API
// @version 1.0
// @description This is a sample server for Hello World microservice.
// @host localhost:8081
// @BasePath /
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	msg := os.Getenv("HELLO_MESSAGE")
	if msg == "" {
		msg = "Hello, world!"
	}

	r := gin.Default()

	// CORS対応
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/", handler.HelloHandler(msg))
	r.GET("/hello", handler.HelloAPIHandler(msg))

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + port)
}
