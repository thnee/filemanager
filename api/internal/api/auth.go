package api

import (
	"encoding/json"
	"errors"
	"main/internal/g"
	"net/http"
)

func AuthUser(w http.ResponseWriter, r *http.Request) {
	username := g.Session.GetString(r.Context(), "username")

	user, err := getUser(username)
	if err != nil {
		g.JSON(w, http.StatusBadRequest, err)
		return
	}

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

	user, err := getUser(loginForm.Username)
	if err != nil {
		g.JSON(w, http.StatusBadRequest, err)
		return
	}

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

func getUser(username string) (g.User, error) {
	for _, _user := range g.Users {
		if _user.Username == username {
			return _user, nil
		}
	}

	return g.User{}, errors.New("user not found")
}
