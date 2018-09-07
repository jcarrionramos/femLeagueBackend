package server

import (
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// New creates a new rest server
func New() *gin.Engine {
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/ping", pong)

	r.GET("/positions", positions)
	r.GET("/nextmatchs", nextMatchs)
	r.GET("/maxgoalscore", maxGoalScore)

	return r
}
