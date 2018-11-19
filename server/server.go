package server

import (
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}

// New creates a new rest server
func New() *gin.Engine {
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/ping", pong)

	r.GET("/positions", positions)
	r.GET("/nextmatches", nextMatches)
	r.GET("/maxgoalscorers", topScorers)
	r.GET("/allplayers", allPlayers)
	r.GET("/allreferees", allReferees)

	r.GET("/newteam", newTeam)
	r.POST("/newplayer", newPlayer)
	r.POST("/newreferee", newReferee)

	r.GET("deleteteam", deleteTeam)
	r.GET("deleteplayer", deletePlayer)
	r.GET("deletereferee", deleteReferee)

	r.POST("newfixture", newFixture)

	return r
}
