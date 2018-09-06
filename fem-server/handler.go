package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/jcarrionramos/femLeague-backend/fem-structures"
	structures "github.com/jcarrionramos/femLeague-backend/fem-structures"
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

}
