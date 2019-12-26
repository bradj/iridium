package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bradj/iridium/config"
	"github.com/bradj/iridium/persistence"
	"github.com/bradj/iridium/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const configLocation string = "config.toml"

func main() {
	c, err := config.NewConfig(configLocation)

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("Loaded config from %s", configLocation)

	db, err := persistence.NewDB(c)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	a := routes.App{
		DB:     db,
		Config: c,
	}

	routes.NewRoutes(r, a)

	log.Printf("Listening on %d", c.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.Port), r))
}
