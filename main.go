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

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	a := routes.App{
		DB:     db,
		Config: c,
		Logger: logger,
	}

	routes.NewRoutes(r, a)

	logger.Printf("Listening on port %d", c.Port)

	logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.Port), r))
}
