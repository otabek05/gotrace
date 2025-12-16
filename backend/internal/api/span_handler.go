package api

import (
	"net/http"
	"os"
	"path/filepath"
)


func spaHandler(staticDir string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := filepath.Join(staticDir, r.URL.Path)

		if info, err := os.Stat(path); err == nil && !info.IsDir() {
			http.ServeFile(w, r, path)
			return
		}

		http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
	})
}