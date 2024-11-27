package controllers

import (
	"github.com/andremartinsds/go_admin/internal/mod"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"sync"
)

type ControllerBase struct {
	DB  *gorm.DB
	Mux *chi.Mux
}

var (
	once sync.Once
	c    *ControllerBase
)

func Initialize(db *gorm.DB, mux *chi.Mux) {
	once.Do(func() {
		c = &ControllerBase{
			DB:  db,
			Mux: mux,
		}
	})
	mod.Mux = mux
	registerControllers()
}

func registerControllers() {
	RegisterLoginController(c)
	RegisterSellerController(c)
	RegisterAccountController(c)
	RegisterManagerController(c)
	RegisterUserController(c)
}
