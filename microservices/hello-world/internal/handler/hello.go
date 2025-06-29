package handler

import (
	"github.com/gin-gonic/gin"
)

// HelloHandler returns a plain text hello message
// @Summary Hello plain text
// @Description Returns a hello message in plain text
// @Tags hello
// @Produce plain
// @Success 200 {string} string "ok"
// @Router / [get]
func HelloHandler(message string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, message+"\n")
	}
}

// HelloAPIHandler returns a hello message as JSON
// @Summary Hello JSON
// @Description Returns a hello message in JSON format
// @Tags hello
// @Produce json
// @Success 200 {object} map[string]string
// @Router /hello [get]
func HelloAPIHandler(message string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": message})
	}
}
