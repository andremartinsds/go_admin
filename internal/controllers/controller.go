package controllers

import (
	"github.com/go-chi/chi"
	"gorm.io/gorm"
	"sync"
)

type ControllerBase struct {
	DB *gorm.DB
	C  chi.Router
}

var (
	once sync.Once
	c    *ControllerBase
)

func Initialize(db *gorm.DB, route chi.Router) {
	once.Do(func() {
		c = &ControllerBase{
			DB: db,
			C:  route,
		}
	})
	registerControllers()
}

func registerControllers() {
	RegisterSellerController(c)
	RegisterAccountController(c)
	RegisterManagerController(c)
	RegisterUserController(c)
}
