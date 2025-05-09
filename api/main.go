package main

import (
	"embed"
	"main/internal/api"
	"main/internal/cmd"
)

//go:embed web/*
var webFS embed.FS

func main() {
	api.WebFS = webFS
	cmd.Execute()
}
