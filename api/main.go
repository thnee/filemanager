package main

import (
	"flag"
	"log"
	"main/internal/config"
	"main/internal/g"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

var customCors = cors.New(cors.Options{
	AllowedOrigins:   []string{"http://127.0.0.1:5173"},
	AllowedMethods:   []string{"POST", "GET", "PUT", "OPTIONS"},
	AllowedHeaders:   []string{"Content-Type"},
	AllowCredentials: true,

	Debug: false,
})

func main() {
	var configFile string

	flag.StringVar(&configFile, "config", "config.yml", "Config file path")
	flag.Parse()

	err := config.ReadConfig(configFile)
	if err != nil {
		log.Fatal("Failed reading config file: ", err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(customCors.Handler)
	r.Use(g.Session.LoadAndSave)

	RegisterRoutes(r)

	r.NotFound(NotFoundHandler)

	http.ListenAndServe("127.0.0.1:4000", r)
}
