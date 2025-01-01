package controllers

import (
	loginHandler "github.com/andremartinsds/go_admin/internal/handlers/login"
	"github.com/andremartinsds/go_admin/internal/infra/repositories"
	"github.com/go-chi/chi/v5"
)

type LoginController struct {
	Controller *ControllerBase
}

func RegisterLoginController(controller *ControllerBase) {
	loginController := &LoginController{
		Controller: controller,
	}
	loginController.Routes()
}

func (c *LoginController) Routes() {
	c.Controller.Mux.Route("/login", func(r chi.Router) {
		userRepository := repositories.UserRepositoryInstancy(c.Controller.DB)
		h := loginHandler.LoginHandlerInstance(userRepository)
		r.Post("/", h.Login)
	})
}
