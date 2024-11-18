package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andremartinsds/go_admin/internal/controllers"
	"github.com/andremartinsds/go_admin/internal/infra/db/connection"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"

	"github.com/andremartinsds/go_admin/configs"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := configs.LoadEnvConfig()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	//middlewares.AdminGroup(r)
	r.Use(middleware.WithValue("TokenAuth", configs.Config.TokenAuth))
	r.Use(middleware.WithValue("JWTExpiresIn", configs.Config.JWTExpiresIn))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	//r.Use(jwtauth.Verifier(configs.Config.TokenAuth))
	//r.Use(jwtauth.Authenticator(configs.Config.TokenAuth))
	//r.Use(middlewares.AuthPermissions)

	db, err := connection.DatabaseStart()
	if err != nil {
		log.Fatal(err)
	}

	controllers.Initialize(db, r)

	err = http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("APP_PORT")), r)
	if err != nil {
		panic(err)
	}
}
