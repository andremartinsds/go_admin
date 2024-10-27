package controllers

import (
	userHandler "github.com/andremartinsds/go_admin/internal/handlers/user"
	"github.com/andremartinsds/go_admin/internal/infra/repositories"
	"github.com/go-chi/chi"
)

type UserController struct {
	Controller *ControllerBase
}

func RegisterUserController(controller *ControllerBase) {
	userController := &UserController{
		Controller: controller,
	}
	userController.Routes()
}

func (userController *UserController) Routes() {
	userController.Controller.C.Route("/users", func(r chi.Router) {
		userRepository := repositories.UserRepositoryInstancy(userController.Controller.DB)
		accountHandler := userHandler.UserHandlerInstance(userRepository)
		r.Post("/", accountHandler.CreateUser)
		// r.Get("/", accountHandler.List)
		r.Get("/{userId}", accountHandler.SelectUser)
		r.Put("/{userId}", accountHandler.UpdateUser)
	})
}
