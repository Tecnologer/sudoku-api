package router

import (
	"github.com/gorilla/mux"
	"github.com/tecnologer/sudoku/clients/sudoku-api/middleware"
)

//Router provides a new router instance
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/game", middleware.IsAuthorized(middleware.NewGame)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/game/set", middleware.IsAuthorized(middleware.SetValue)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/game/levels", middleware.IsAuthorized(middleware.GetLevels)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/game/validate", middleware.IsAuthorized(middleware.Validate)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/game/solve", middleware.IsAuthorized(middleware.Solve)).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/account/signup", middleware.SignUp).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/account/signin", middleware.SignIn).Methods("GET", "OPTIONS")

	router.HandleFunc("/user", middleware.IsAuthorized(middleware.UserIndex)).Methods("GET")

	return router
}
