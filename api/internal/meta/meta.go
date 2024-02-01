package meta

import (
	"net/http"
)

func GetHealthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
