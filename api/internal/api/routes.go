package api

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux) {
	r.Route("/api", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/auth/user", AuthUser)
			r.Post("/auth/login", AuthLogin)
			r.Post("/auth/logout", AuthLogout)
		})

		r.Group(func(r chi.Router) {
			r.Get("/files/areas", FilesAreasList)
			r.Get("/files/file", FilesFile)
			r.Get("/files/list", FilesList)
			r.Get("/files/data", FilesData)
		})
	})

	r.Group(func(r chi.Router) {
		r.Get("/healthz", GetHealthz)
	})
}
