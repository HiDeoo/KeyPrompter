package ui

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"path"
)

//go:embed build/*
var Assets embed.FS

type fsHandler func(name string) (fs.File, error)

func (handler fsHandler) Open(name string) (fs.File, error) {
	return handler(name)
}

func AssetHandler() http.Handler {
	handler := fsHandler(func(name string) (fs.File, error) {
		assetPath := path.Join("build", name)

		file, err := Assets.Open(assetPath)

		if os.IsNotExist(err) {
			return Assets.Open("build/index.html")
		}

		return file, err
	})

	return http.StripPrefix("/", http.FileServer(http.FS(handler)))
}
