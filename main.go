package main

import (
	"go-crud/app/config"
	"go-crud/app/controller"
	"go-crud/app/repository"
	"go-crud/app/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// setup configuration
	configuration := config.New()
	database := config.NewMongoDatabase(*configuration)

	// setup repository
	userRepository := repository.NewUserRepository(database)

	// setup service
	userService := service.NewUserService(*userRepository)

	// setup controller
	userController := controller.NewUserController(*userService)

	// setup router
	r := chi.NewRouter()
	r.Post("/register", userController.Register)
	r.Post("/login", userController.Login)

	r.Group(func(r chi.Router) {
		r.Use(controller.JWTMiddleware)

		r.Get("/profile", userController.UserDetail)

		r.Get("/user", userController.ListAll)
		r.Patch("/user/{username}", userController.Update)
		r.Delete("/user/{username}", userController.Delete)
	})

	// start web server
	log.Println("Start web server on port :8080")
	http.ListenAndServe(":8080", r)
}
