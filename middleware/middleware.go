package middleware

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
	sudoku "github.com/tecnologer/sudoku/sudoku-lib/src"
)

var (
	game *sudoku.Game
)

//NewGame starts new sudoku game
func NewGame(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)

	levelStr := getParamString("level", "easy", r)

	logrus.WithFields(logrus.Fields{
		"level":  levelStr,
		"client": r.Header.Get("User-Agent"),
	}).Debug("new game")

	level := sudoku.StringToComplexity(levelStr)

	if level == sudoku.InvalidLevel {
		preconditionFailedf(&w, "the level \"%s\" is not valid", levelStr)
		return
	}

	game = sudoku.NewGame(level)
	gameBin, err := json.Marshal(game)
	if err != nil {
		internalErrorf(&w, "error creating the game: %v", err)
		return
	}

	ok(&w, gameBin)
}

//SetValue starts new sudoku game
func SetValue(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)

	xStr, provided := getParam("x", r)
	if !provided || len(xStr) == 0 {
		preconditionFailedf(&w, "the value for x is required")
		return
	}

	yStr, provided := getParam("y", r)
	if !provided || len(yStr) == 0 {
		preconditionFailedf(&w, "the value for y is required")
		return
	}

	nStr, provided := getParam("n", r)
	if !provided || len(nStr) == 0 {
		preconditionFailedf(&w, "the value for the cell is required")
		return
	}

	x, err := strconv.Atoi(xStr[0])
	if err != nil || x == -1 || x > 8 {
		msg := newMsgResf(http.StatusPreconditionFailed, "the value for x (%d) is not valid", x)
		if err != nil {
			msg = newMsgResf(http.StatusPreconditionFailed, "the value for x should be number between 1 and 9")
		}

		w.WriteHeader(http.StatusPreconditionFailed)
		w.Write(msg)
		return
	}

	y, err := strconv.Atoi(yStr[0])
	if err != nil || y == -1 || y > 8 {
		msg := newMsgResf(http.StatusPreconditionFailed, "the value for y (%d) is not valid", y)
		if err != nil {
			msg = newMsgResf(http.StatusPreconditionFailed, "the value for y should be number between 1 and 9")
		}

		w.WriteHeader(http.StatusPreconditionFailed)
		w.Write(msg)
		return
	}

	n, err := strconv.Atoi(nStr[0])
	if err != nil || n < 1 || n > 9 {
		msg := newMsgResf(http.StatusPreconditionFailed, "the value for cell (%d) is not valid", n)
		if err != nil {
			msg = newMsgResf(http.StatusPreconditionFailed, "the value for cell should be number between 1 and 9")
		}

		w.WriteHeader(http.StatusPreconditionFailed)
		w.Write(msg)
		return
	}

	logrus.WithFields(logrus.Fields{
		"x":      x,
		"y":      y,
		"n":      n,
		"client": r.Header.Get("User-Agent"),
	}).Debug("set value")

	if game.IsCoordinateLockedXY(x, y) {
		preconditionFailedf(&w, "the coordinate (%d, %d) cannot be modified", x, y)
		return
	}

	game.Set(x, y, n)
	gameBin, err := json.Marshal(game)
	if err != nil {
		internalErrorf(&w, "error updating the game: %v", err)
		return
	}

	ok(&w, gameBin)
}

//GetLevels returns the available levels
func GetLevels(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	logrus.WithField("client", r.Header.Get("User-Agent")).Debug("get levels")

	levels, err := json.Marshal(sudoku.GetComplexities())
	if err != nil {
		internalErrorf(&w, "error getting the levels")
		return
	}

	ok(&w, levels)
}

//Validate returns the available levels
func Validate(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	logrus.WithField("client", r.Header.Get("User-Agent")).Debug("validate")

	checkEmpties := getParamBool("empty", false, r)
	errs := game.Validate(checkEmpties)

	resp, err := json.Marshal(errs)
	if err != nil {
		internalErrorf(&w, "validating: parsing response to json. %v", err)
		return
	}
	ok(&w, resp)
}
