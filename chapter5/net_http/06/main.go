package main

import (
	"net/http"
	"path"
)

func main() {
	fileserver := http.StripPrefix("/public/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if path.Ext(r.URL.Path) == ".xls" {
			w.Header().Set("Content-Type", "application/vnd.ms-excel")
		}
		fileserver.ServeHTTP(w, r)
	})
}
