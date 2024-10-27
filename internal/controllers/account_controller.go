package controllers

import (
	accountHandler "github.com/andremartinsds/go_admin/internal/handlers/account"
	"github.com/andremartinsds/go_admin/internal/infra/repositories"
	"github.com/go-chi/chi"
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
	accountController.Controller.C.Route("/accounts", func(r chi.Router) {
		accountRepository := repositories.AccountRepositoryInstancy(accountController.Controller.DB)
		h := accountHandler.AccountHandlerInstancy(accountRepository)
		r.Post("/", h.CreateAccount)
		r.Get("/", h.List)
		r.Get("/{accountId}", h.SelectAccount)
		r.Put("/{accountId}", h.UpdateAccount)
	})
}
