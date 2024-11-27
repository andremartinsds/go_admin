package controllers

import (
	accountHandler "github.com/andremartinsds/go_admin/internal/handlers/account"
	"github.com/andremartinsds/go_admin/internal/infra/repositories"
	"github.com/go-chi/chi/v5"
)

type AccountController struct {
	Controller *ControllerBase
}

func RegisterAccountController(controller *ControllerBase) {
	accountController := &AccountController{
		Controller: controller,
	}
	accountController.Routes()
}

func (accountController *AccountController) Routes() {
	accountRepository := repositories.AccountRepositoryInstance(accountController.Controller.DB)
	h := accountHandler.AccountHandlerInstancy(accountRepository)

	accountController.Controller.Mux.Route("/accounts", func(r chi.Router) {
		r.Get("/{accountID}", h.SelectAccount)
		r.Put("/{accountID}", h.UpdateAccount)
		r.Post("/", h.CreateAccount)
		r.Get("/", h.List)
		r.Delete("/{accountID}", h.Delete)
	})
}
