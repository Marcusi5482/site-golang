package main

import (
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World!</h1>"))
	w.Write([]byte("<h1>Gigachad</h1>"))
	w.Write([]byte("<h1>https://github.com/Marcusi5482/site-golang/releases/tag/1</h1>"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
