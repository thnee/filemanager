package api

import (
	"main/internal/g"
	"net/http"
)

func AuthUser(w http.ResponseWriter, r *http.Request) {
	userID := g.Session.Get(r.Context(), "userID")
	var data g.Data
	if userID == nil {
		data = g.Data{}
	} else {
		data = g.Data{
			"ID":       123,
			"Username": "matte",
			"Email":    "matte@gmail",
		}
	}

	g.JSON(w, http.StatusOK, data)
}

func AuthLogin(w http.ResponseWriter, r *http.Request) {
	err := g.Session.RenewToken(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	g.Session.Put(r.Context(), "userID", 123)
}

func AuthLogout(w http.ResponseWriter, r *http.Request) {
	g.Session.Destroy(r.Context())

	g.Text(w, http.StatusOK, "")
}
