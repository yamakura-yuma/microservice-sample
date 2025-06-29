package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Hello Swagger API
// @version 1.0
// @description This is a sample server for Swagger UI with gin.
// @host localhost:8080
// @BasePath /
func main() {
	// hello-worldのswagger.jsonのURLを指定
	url := ginSwagger.URL("http://hello-world:8080/swagger/doc.json")

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(":8080")
}
