package router

import (
	"github.com/gorilla/mux"
	"github.com/tecnologer/sudoku/clients/sudoku-api/middleware"
)

//Router provides a new router instance
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/game", middleware.NewGame).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/game/set", middleware.SetValue).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/game/levels", middleware.GetLevels).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/game/validate", middleware.Validate).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/game/solve", middleware.Solve).Methods("GET", "OPTIONS")

	return router
}
