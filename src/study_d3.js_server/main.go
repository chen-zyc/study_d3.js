package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	StaticServer(":9090", map[string]string{
		"/files": "files",
		"/html":  "html",
	})
}

func StaticServer(addr string, paths map[string]string) {
	servers := make(map[string]http.Handler, len(paths))
	for pattern, path := range paths {
		servers[pattern] = http.FileServer(http.Dir(path))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path
		for pattern, handler := range servers {
			if strings.HasPrefix(urlPath, pattern) {
				http.StripPrefix(pattern, handler).ServeHTTP(w, r)
				return
			}
		}
		http.NotFound(w, r)
	})
	log.Fatal(http.ListenAndServe(addr, nil))
}
