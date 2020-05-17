//go:generate sqlboiler --wipe psql

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bradj/iridium/config"
	"github.com/bradj/iridium/persistence"
	"github.com/bradj/iridium/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

const configLocation string = "config.toml"

func init() {
	// overrides default Chi logger to log in UTC
	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags+log.LUTC), NoColor: false})
}

func main() {
	c, err := config.NewConfig(configLocation)

	logger := log.New(os.Stdout, "", log.LstdFlags+log.LUTC)

	if err != nil {
		log.Fatal(err)
		return
	}

	logger.Printf("Loaded config from '%s'", configLocation)

	db, err := persistence.NewDB(c)

	if err != nil {
		logger.Fatal(err)
		return
	}

	defer db.Close()

	r := chi.NewRouter()

	xs := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(xs.Handler)

	a := routes.App{
		DB:     db,
		Config: c,
	}

	routes.NewRoutes(r, a)

	logger.Printf("Listening on port %d", c.Port)

	logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.Port), r))
}
