package main

import (
	"net/http"
	g "rhythm/pkg/generator"
	r "rhythm/pkg/renderer"
	h "rhythm/web/handlers"
)

func main() {
	app := h.NewApp(g.NewSequenceGenerator(), r.NewImageRenderer())
	http.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir("web/src"))))

	http.HandleFunc("/", app.IndexHandler)
	http.HandleFunc("/image", app.SimpleHandler)
	http.HandleFunc("/surprise", app.UnhingedHandler)
	http.ListenAndServe(":8080", nil)
}
