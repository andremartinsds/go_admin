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
		sellerRepository := repositories.SellerRepositoryInstance(controller.Controller.DB)
		handler := userHandler.UserHandlerInstance(userRepository, accountRepository, sellerRepository)
		r.Post("/", handler.CreateUser)
		// r.Get("/", accountHandler.List)
		r.Get("/{userId}", handler.SelectUser)
		r.Put("/{userId}", handler.UpdateUser)
	})
}
