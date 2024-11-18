package controllers

import (
	userHandler "github.com/andremartinsds/go_admin/internal/handlers/user"
	"github.com/andremartinsds/go_admin/internal/infra/repositories"
	"github.com/go-chi/chi/v5"
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

func (controller *UserController) Routes() {
	controller.Controller.Mux.Route("/users", func(r chi.Router) {
		userRepository := repositories.UserRepositoryInstancy(controller.Controller.DB)
		accountRepository := repositories.AccountRepositoryInstance(controller.Controller.DB)
		sellerRepository := repositories.SellerRepositoryInstancy(controller.Controller.DB)
		accountHandler := userHandler.UserHandlerInstance(userRepository, accountRepository, sellerRepository)
		r.Post("/", accountHandler.CreateUser)
		// r.Get("/", accountHandler.List)
		r.Get("/{userId}", accountHandler.SelectUser)
		r.Put("/{userId}", accountHandler.UpdateUser)
	})
}
