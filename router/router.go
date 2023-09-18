package router

import (
	"go-learn/controller/auth"
	"go-learn/repositories"
	"go-learn/service"
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()
	//set dependency
	repo := repositories.NewRepo()
	serv := service.NewService(repo)

	// call controllers auth
	controllerLogin := auth.NewControllerLogin(*serv)
	controllerRegister := auth.NewControllerRegister(*serv)
	controllerStatus := auth.NewControllerStatus(*serv)

	//login
	router.HandleFunc("/login", controllerLogin.HandleLogin).Methods("POST")
	router.HandleFunc("/register", controllerRegister.HandleRegister).Methods("POST")
	router.HandleFunc("/update-status", controllerStatus.Status).Methods("PUT")

	return router
}
