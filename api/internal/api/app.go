package api

import (
	"embed"
	"errors"
	"io"
	"mime"
	"net/http"
	"path"
	"path/filepath"
)

var WebFS embed.FS

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	err := tryRead(WebFS, "web", r.URL.Path, w)
	if err == nil {
		return
	}
	err = tryRead(WebFS, "web", "index.html", w)
	if err != nil {
		panic(err)
	}
}

func tryRead(fs embed.FS, prefix, requestedPath string, w http.ResponseWriter) error {
	f, err := fs.Open(path.Join(prefix, requestedPath))
	if err != nil {
		return err
	}
	defer f.Close()

	stat, _ := f.Stat()
	if stat.IsDir() {
		return errors.New("path is dir")
	}

	contentType := mime.TypeByExtension(filepath.Ext(requestedPath))
	w.Header().Set("Content-Type", contentType)
	_, err = io.Copy(w, f)
	return err
}
