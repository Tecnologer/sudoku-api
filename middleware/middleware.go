package middleware

import (
	"encoding/json"
	"net/http"

	sudoku "github.com/tecnologer/sudoku/src"
)

var (
	game *sudoku.Game
)

//NewGame starts new sudoku game
func NewGame(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)

	levelStr := getParamString("level", "easy", r)

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
