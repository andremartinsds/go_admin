package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andremartinsds/go_admin/internal/controllers"

	"github.com/andremartinsds/go_admin/internal/infra/db/connection"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"

	"github.com/andremartinsds/go_admin/configs"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	err := configs.LoadEnvConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func bootstrap() (*chi.Mux, error) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db, err := connection.DatabaseStart()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	controllers.Initialize(db, r)

	return r, nil
}

func main() {
	r, err := bootstrap()

	err = http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("APP_PORT")), r)
	if err != nil {
		panic(err)
	}
}
