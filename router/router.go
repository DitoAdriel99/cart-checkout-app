package router

import (
	"go-learn/controller/auth"
	"go-learn/controller/product"
	"go-learn/middleware"
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

	// middlewares
	tokenValidator := middleware.NewTokenValidator(*repo)

	//login
	router.HandleFunc("/login", controllerLogin.HandleLogin).Methods("POST")
	router.HandleFunc("/register", controllerRegister.HandleRegister).Methods("POST")
	router.HandleFunc("/update-status", controllerStatus.Status).Methods("PUT")

	// CRUAD Product
	controllerProduct := product.NewControllerProductCreate(*serv)

	//product for admin access
	adminRoutes := router.PathPrefix("").Subrouter()
	adminRoutes.Use(tokenValidator.ValidateTokenMiddleware("admin"))
	adminRoutes.HandleFunc("/product", controllerProduct.Create).Methods("POST")
	adminRoutes.HandleFunc("/product/{id}", controllerProduct.Update).Methods("PUT")
	adminRoutes.HandleFunc("/product/{id}", controllerProduct.Delete).Methods("DELETE")

	//product for admin access
	userRoutes := router.PathPrefix("").Subrouter()
	userRoutes.Use(tokenValidator.ValidateTokenMiddleware("admin", "user"))
	userRoutes.HandleFunc("/product", controllerProduct.Get).Methods("GET")
	userRoutes.HandleFunc("/product/{id}", controllerProduct.Detail).Methods("GET")

	userRoutes.HandleFunc("/cart/items", controllerProduct.AddToCart).Methods("POST")
	userRoutes.HandleFunc("/cart/items", controllerProduct.GetCart).Methods("GET")
	userRoutes.HandleFunc("/cart/items", controllerProduct.DeleteCart).Methods("DELETE")
	userRoutes.HandleFunc("/cart/items/checkout", controllerProduct.Checkout).Methods("POST")

	return router
}
