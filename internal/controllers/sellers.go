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
	s.Controller.Mux.Route("/sellers", func(r chi.Router) {
		sellerRepository := repositories.SellerRepositoryInstancy(s.Controller.DB)
		accountRepository := repositories.AccountRepositoryInstance(s.Controller.DB)
		h := sellerHandler.Instance(sellerRepository, accountRepository)

		r.Post("/", h.CreateSeller)
		r.Get("/{sellerID}", h.SelectSeller)
		r.Get("/", h.ListSeller)
		r.Put("/{sellerID}", h.UpdateSeller)
		r.Delete("/{sellerID}", h.DesactiveSeller)

	})
}
