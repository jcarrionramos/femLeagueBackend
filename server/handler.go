package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcarrionramos/femLeagueBackend/models"
	"github.com/jcarrionramos/femLeagueBackend/structures"
)

func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func positions(c *gin.Context) {
	teams, err := models.SelectAllTeams()

	if err != nil {
		c.JSON(http.StatusInternalServerError, structures.Response{
			Status: http.StatusInternalServerError,
			Meta:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, structures.Response{
		Status: http.StatusOK,
		Data:   teams,
	})
}

func nextMatchs(c *gin.Context) {

}

func maxGoalScore(c *gin.Context) {
	players, err := models.SelectTopScors()

	if err != nil {
		c.JSON(http.StatusInternalServerError, structures.Response{
			Status: http.StatusInternalServerError,
			Meta:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, structures.Response{
		Status: http.StatusOK,
		Data:   players,
	})
}
