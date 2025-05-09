package api

import (
	"errors"
	"fmt"
	"log"
	"main/internal/g"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"syscall"
	"time"

	"github.com/spf13/viper"
)

func FilesAreasList(w http.ResponseWriter, r *http.Request) {
	areasData := []g.Data{}

	for _, area := range g.FileAreas {
		areasData = append(areasData, g.Data{
			"Name": area.Name,
		})
	}

	g.JSON(w, http.StatusOK, g.Data{"areas": areasData})
}

func FilesFile(w http.ResponseWriter, r *http.Request) {
	queryPath := r.URL.Query().Get("path")
	localPath, err := parseQueryPath(queryPath)
	if err != nil {
		g.JSON(w, http.StatusNotFound, "")
		return
	}

	fileInfo, err := os.Stat(localPath)
	if err != nil {
		g.Text(w, http.StatusNotFound, "")
		return
	}

	g.JSON(w, http.StatusOK, getFileDetails(queryPath, fileInfo))
}

func FilesList(w http.ResponseWriter, r *http.Request) {
	queryPath := r.URL.Query().Get("path")
	localPath, err := parseQueryPath(queryPath)
	if err != nil {
		g.JSON(w, http.StatusNotFound, "")
		return
	}

	dirEntries, _ := os.ReadDir(localPath)

	data := []g.Data{}

	sortDirEntries(dirEntries)

	for _, dirEntry := range dirEntries {
		fileInfo, err := dirEntry.Info()
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, getFileDetails(filepath.Join(queryPath, fileInfo.Name()), fileInfo))
	}

	g.JSON(w, http.StatusOK, g.Data{"Files": data})
}

func FilesData(w http.ResponseWriter, r *http.Request) {
	queryPath := r.URL.Query().Get("path")
	localPath, err := parseQueryPath(queryPath)
	if err != nil {
		g.JSON(w, http.StatusNotFound, "")
		return
	}

	fileInfo, err := os.Stat(localPath)
	if err != nil {
		g.Text(w, http.StatusNotFound, "")
		return
	}

	if r.URL.Query().Has("download") {
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileInfo.Name()))
	}
	http.ServeFile(w, r, localPath)
}

func parseQueryPath(path string) (localPath string, err error) {
	parts := strings.SplitN(path, string(os.PathSeparator), 2)
	areaName := parts[0]

	var area *g.FileArea

	for _, _area := range g.FileAreas {
		if _area.Name == areaName {
			area = &_area
			break
		}
	}

	if area == nil {
		return "", errors.New("area not valid")
	}

	localPath = area.Path

	if len(parts) == 2 {
		localPath = filepath.Join(localPath, parts[1])
	}

	return localPath, nil
}

func getFileDetails(path string, fileInfo os.FileInfo) (data g.Data) {
	ext := filepath.Ext(path)
	sysStat := fileInfo.Sys().(*syscall.Stat_t)

	atime := time.Unix(0, syscall.TimespecToNsec(sysStat.Atim))
	ctime := time.Unix(0, syscall.TimespecToNsec(sysStat.Ctim))
	mtime := time.Unix(0, syscall.TimespecToNsec(sysStat.Mtim))

	return g.Data{
		"IsDir":       fileInfo.IsDir(),
		"Name":        fileInfo.Name(),
		"NameStem":    strings.TrimSuffix(fileInfo.Name(), ext),
		"NameExt":     strings.TrimPrefix(ext, "."),
		"MimeType":    mime.TypeByExtension(ext),
		"Path":        path,
		"Mode":        fileInfo.Mode(),
		"Size":        fileInfo.Size(),
		"Atime":       atime.Format(time.RFC3339Nano),
		"Mtime":       mtime.Format(time.RFC3339Nano),
		"Ctime":       ctime.Format(time.RFC3339Nano),
		"DataURL":     viper.GetString("server.url") + "/api/files/data?path=" + path,
		"DownloadURL": viper.GetString("server.url") + "/api/files/data?path=" + path + "&download",
	}
}

func sortDirEntries(dirEntries []os.DirEntry) {
	sort.Slice(dirEntries, func(i, j int) bool {
		a := dirEntries[i]
		b := dirEntries[j]
		if a.IsDir() && !b.IsDir() {
			return true
		}
		if !a.IsDir() && b.IsDir() {
			return false
		}
		return a.Name() < b.Name()
	})
}
