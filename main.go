package main

import (
	"context"
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

	// Places Config & DB into context()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), routes.ConfigKey, c)
			ctx = context.WithValue(ctx, routes.DBKey, db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	r.Group(routes.Protected)
	r.Group(routes.Public)

	log.Printf("Listening on %d", c.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.Port), r))
}
