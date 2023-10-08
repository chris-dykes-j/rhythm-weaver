package main

import (
	"net/http"
    h "rhythm/web/handlers"
	g "rhythm/pkg/generator"
	r "rhythm/pkg/renderer"
)

func main() {
	app := h.NewApp(g.NewSequenceGenerator(), r.NewImageRenderer())
	http.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir("web/src"))))

	http.HandleFunc("/", app.IndexHandler)
	http.HandleFunc("/image", app.ImageHandler)
	http.ListenAndServe(":8080", nil)
}
