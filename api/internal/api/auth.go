package api

import (
	"encoding/json"
	"main/internal/config"
	"main/internal/g"
	"net/http"
)

func AuthUser(w http.ResponseWriter, r *http.Request) {
	username := g.Session.GetString(r.Context(), "username")

	user := getUser(username)

	var data g.Data

	if user.Username != "" {
		data = g.Data{
			"Username": user.Username,
		}

	} else {
		data = g.Data{}
	}

	g.JSON(w, http.StatusOK, data)
}

func AuthLogin(w http.ResponseWriter, r *http.Request) {
	var loginForm struct {
		Username string
		Password string
	}

	err := json.NewDecoder(r.Body).Decode(&loginForm)
	if err != nil {
		return
	}

	user := getUser(loginForm.Username)

	if user.Username != "" && loginForm.Password == user.PasswordClear {
		err := g.Session.RenewToken(r.Context())
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		g.Session.Put(r.Context(), "username", loginForm.Username)
		g.Text(w, http.StatusOK, "")

	} else {
		g.Session.Destroy(r.Context())
		g.JSON(w, http.StatusBadRequest, g.Data{"errors": []string{"Invalid credentials."}})
	}
}

func AuthLogout(w http.ResponseWriter, r *http.Request) {
	g.Session.Destroy(r.Context())

	g.Text(w, http.StatusOK, "")
}

func getUser(username string) config.User {
	for _, _user := range config.Config.Users {
		if _user.Username == username {
			return _user
		}
	}

	return config.User{}
}
