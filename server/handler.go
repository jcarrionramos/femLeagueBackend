package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jcarrionramos/femLeagueBackend/models"
	"github.com/jcarrionramos/femLeagueBackend/structures"
)

func pong(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func positions(ctx *gin.Context) {
	active := ctx.Query("active")
	teams, err := models.SelectActivesTeams(active)

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
		ctx.JSON(500, structures.Response{
			Status: 500,
			Meta:   err.Error(),
		})
		return
	}

	ctx.JSON(200, structures.Response{
		Status: 200,
		Data:   players,
	})
}

func allPlayers(ctx *gin.Context) {
	players, err := models.SelectAllPlayers()

	if err != nil {
		ctx.JSON(500, structures.Response{
			Status: 500,
			Meta:   err.Error(),
		})
		return
	}

	ctx.JSON(200, structures.Response{
		Status: 200,
		Data:   players,
	})
}

func allReferees(ctx *gin.Context) {
	referees, err := models.SelectAllReferees()

	if err != nil {
		ctx.JSON(500, structures.Response{
			Status: 500,
			Meta:   err.Error(),
		})
		return
	}

	ctx.JSON(200, structures.Response{
		Status: 200,
		Data:   referees,
	})
}

func newTeam(ctx *gin.Context) {
	team := structures.Team{
		Name:  ctx.Query("name"),
		Win:   0,
		Draw:  0,
		Loss:  0,
		Total: 0,
	}

	err := models.InsertTeam(team)

	if err != nil {
		ctx.JSON(500, structures.Response{
			Status: 500,
			Meta:   err.Error(),
		})
		return
	}

	ctx.JSON(200, structures.Response{
		Status: 200,
		Data:   "succes",
	})
}

func newPlayer(ctx *gin.Context) {
	var player structures.Player

	err := ctx.ShouldBindJSON(&player)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, structures.Response{
			Status: http.StatusBadRequest,
			Meta:   err.Error(),
		})
		return
	}

	err = models.InsertPlayer(player)

	if err != nil {
		ctx.JSON(500, structures.Response{
			Status: 500,
			Meta:   err.Error(),
		})
		return
	}

	ctx.JSON(200, structures.Response{
		Status: 200,
		Data:   "succes",
	})
}

func newReferee(ctx *gin.Context) {
	var referee structures.Referee

	err := ctx.ShouldBindJSON(&referee)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, structures.Response{
			Status: http.StatusBadRequest,
			Meta:   err.Error(),
		})
		return
	}

	err = models.InsertReferee(referee)

	if err != nil {
		ctx.JSON(400, structures.Response{
			Status: 400,
			Meta:   err.Error(),
		})
		return
	}

	ctx.JSON(200, structures.Response{
		Status: 200,
		Data:   "succes",
	})
}

func deleteTeam(ctx *gin.Context) {
	name := ctx.Query("name")

	err := models.DeleteTeam(name)

	if err != nil {
		ctx.JSON(400, structures.Response{
			Status: 400,
			Meta:   err.Error(),
		})
		return
	}

	ctx.JSON(200, structures.Response{
		Status: 200,
		Data:   "succes",
	})
}

func deletePlayer(ctx *gin.Context) {
	rut, _ := strconv.Atoi(ctx.Query("rut"))

	err := models.DeletePlayer(rut)

	if err != nil {
		ctx.JSON(400, structures.Response{
			Status: 400,
			Meta:   err.Error(),
		})
		return
	}

	ctx.JSON(200, structures.Response{
		Status: 200,
		Data:   "succes",
	})
}

func deleteReferee(ctx *gin.Context) {
	rut, _ := strconv.Atoi(ctx.Query("rut"))

	err := models.DeleteReferee(rut)

	if err != nil {
		ctx.JSON(400, structures.Response{
			Status: 400,
			Meta:   err.Error(),
		})
		return
	}

	ctx.JSON(200, structures.Response{
		Status: 200,
		Data:   "succes",
	})
}

func newFixture(ctx *gin.Context) {
	d := struct {
		TeamsName []string `json:"TeamsName"`
		Season    string   `json:"Season"`
	}{}

	err := ctx.ShouldBindJSON(&d)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, structures.Response{
			Status: http.StatusBadRequest,
			Meta:   err.Error(),
		})
		return
	}

	//FIRST ROUND
	_ = models.InsertMatch(d.TeamsName[0], d.TeamsName[5], d.Season, 1)
	_ = models.InsertMatch(d.TeamsName[1], d.TeamsName[4], d.Season, 1)
	_ = models.InsertMatch(d.TeamsName[2], d.TeamsName[3], d.Season, 1)

	//SECOND ROUND
	_ = models.InsertMatch(d.TeamsName[0], d.TeamsName[1], d.Season, 2)
	_ = models.InsertMatch(d.TeamsName[4], d.TeamsName[2], d.Season, 2)
	_ = models.InsertMatch(d.TeamsName[5], d.TeamsName[3], d.Season, 2)

	//THIRD ROUND
	_ = models.InsertMatch(d.TeamsName[2], d.TeamsName[0], d.Season, 3)
	_ = models.InsertMatch(d.TeamsName[3], d.TeamsName[4], d.Season, 3)
	_ = models.InsertMatch(d.TeamsName[1], d.TeamsName[5], d.Season, 3)

	//FOURTH ROUND
	_ = models.InsertMatch(d.TeamsName[0], d.TeamsName[3], d.Season, 4)
	_ = models.InsertMatch(d.TeamsName[1], d.TeamsName[2], d.Season, 4)
	_ = models.InsertMatch(d.TeamsName[5], d.TeamsName[4], d.Season, 4)

	//FIFTH ROUND
	_ = models.InsertMatch(d.TeamsName[4], d.TeamsName[0], d.Season, 5)
	_ = models.InsertMatch(d.TeamsName[3], d.TeamsName[1], d.Season, 5)
	_ = models.InsertMatch(d.TeamsName[2], d.TeamsName[5], d.Season, 5)

	_ = models.UpdateActivesTeams(d.TeamsName)

	ctx.JSON(200, structures.Response{
		Status: 200,
		Data:   "succes",
	})
}

func updateMatch(ctx *gin.Context) {
	var m structures.Match

	err := ctx.ShouldBindJSON(&m)

	if err != nil {
		ctx.JSON(400, structures.Response{
			Status: 400,
			Meta:   err.Error(),
		})
	}

	err = models.UpdateFixture(m)

	if err != nil {
		ctx.JSON(400, structures.Response{
			Status: 400,
			Meta:   err.Error(),
		})
	}

	ctx.JSON(200, structures.Response{
		Status: 200,
		Meta:   "success",
	})
}
