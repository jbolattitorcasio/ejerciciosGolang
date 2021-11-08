package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := muxDefault()

	pathToUrl := map[string]string{
		"/pepe":  "https://es.wikipedia.org/wiki/Pepe_(futbolista)",
		"/jorge": "https://es.wikipedia.org/wiki/Jorge",
	}
	mapHandler := MapHandler(pathToUrl, mux)

	http.ListenAndServe(":8080", mapHandler)
}

func muxDefault() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", inicio)
	return mux
}

func inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Bienvenidos al Ejecicio dos de <span>Gophercise</span></h1>")
}

func MapHandler(pathToUrl map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathToUrl[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}
