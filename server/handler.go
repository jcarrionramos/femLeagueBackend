package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcarrionramos/femLeagueBackend/models"
	"github.com/jcarrionramos/femLeagueBackend/structures"
)

func pong(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func positions(ctx *gin.Context) {
	teams, err := models.SelectAllTeams()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, structures.Response{
			Status: http.StatusInternalServerError,
			Meta:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, structures.Response{
		Status: http.StatusOK,
		Data:   teams,
	})
}

func nextMatches(ctx *gin.Context) {
	matches, err := models.SelectNextsMatches()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, structures.Response{
			Status: http.StatusInternalServerError,
			Meta:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, structures.Response{
		Status: http.StatusOK,
		Data:   matches,
	})
}

func topScorers(ctx *gin.Context) {
	players, err := models.SelectTopScorers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, structures.Response{
			Status: http.StatusInternalServerError,
			Meta:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, structures.Response{
		Status: http.StatusOK,
		Data:   players,
	})
}
