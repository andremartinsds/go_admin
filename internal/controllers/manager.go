package controllers

import (
	contactHandler "github.com/andremartinsds/go_admin/internal/handlers/contacts"
	"github.com/andremartinsds/go_admin/internal/infra/repositories"
	"github.com/go-chi/chi/v5"
)

type ManagerController struct {
	Controller *ControllerBase
}

func RegisterManagerController(controller *ControllerBase) {
	managerController := &ManagerController{
		Controller: controller,
	}
	managerController.Routes()
}

func (managerController *ManagerController) Routes() {
	managerController.Controller.Mux.Route("/manager", func(r chi.Router) {
		// contact repository instance
		contactRepository := repositories.ContactRepositoryInstancy(managerController.Controller.DB)
		// contact instance
		h := contactHandler.ContactHandlerInstancy(contactRepository)
		// contact routes
		r.Post("/contacts", h.CreateContact)
		r.Get("/contacts", h.List)
	})
}
