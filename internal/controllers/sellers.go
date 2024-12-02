package controllers

import (
	sellerHandler "github.com/andremartinsds/go_admin/internal/handlers/seller"
	"github.com/andremartinsds/go_admin/internal/infra/repositories"
	"github.com/go-chi/chi/v5"
)

type SellerController struct {
	Controller *ControllerBase
}

func RegisterSellerController(controller *ControllerBase) {
	sellerController := SellerController{
		Controller: controller,
	}
	sellerController.Routes()
}

func (s *SellerController) Routes() {
	sellerRepository := repositories.SellerRepositoryInstance(s.Controller.DB)
	accountRepository := repositories.AccountRepositoryInstance(s.Controller.DB)
	h := sellerHandler.Instance(sellerRepository, accountRepository)
	s.Controller.Mux.Route("/sellers", func(r chi.Router) {
		r.Post("/", h.CreateSeller)
		r.Get("/account/{accountID}/seller/{sellerID}", h.SelectSeller)
		r.Get("/account/{accountID}", h.ListSeller)
		r.Put("/{sellerID}", h.UpdateSeller)
		r.Delete("/{sellerID}", h.DesactiveSeller)
	})
}
