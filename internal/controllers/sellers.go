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
		r.Get("/select", h.SelectSeller)
		r.Get("/list", h.ListSeller)
		r.Put("/", h.UpdateSeller)
		r.Delete("/{sellerID}", h.Inactive)
	})
}
