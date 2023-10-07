package main

import (
	"fmt"
	"net/http"
    h "rhythm/web/handlers"
	g "rhythm/pkg/generator"
	i "rhythm/pkg/renderer"
)

func main() {
	sg := g.NewSequenceGenerator()
	ir := i.NewImageRenderer()
	app := h.NewApp(sg, ir)

	fmt.Println(sg.CreateAutoComplexSeq(5, 4))
	fmt.Println(sg.CreateAutoComplexSeq(10, 4))
	fmt.Println(sg.CreateAutoComplexSeq(15, 4))
	fmt.Println(sg.CreateAutoComplexSeq(20, 4))

	http.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir("web/src"))))

	http.HandleFunc("/", app.IndexHandler)
	http.HandleFunc("/image", app.ImageHandler)
	http.ListenAndServe(":8080", nil)
}
