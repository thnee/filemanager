package main

import (
	"main/internal/api"
	"main/internal/meta"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux) {
	r.Route("/api", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/auth/user", api.AuthUser)
			r.Post("/auth/login", api.AuthLogin)
			r.Post("/auth/logout", api.AuthLogout)
		})

		r.Group(func(r chi.Router) {
			r.Get("/files/areas", api.FilesAreasList)
			r.Get("/files/file", api.FilesFile)
			r.Get("/files/list", api.FilesList)
			r.Get("/files/data", api.FilesData)
		})
	})

	r.Group(func(r chi.Router) {
		r.Get("/healthz", meta.GetHealthz)
	})
}
