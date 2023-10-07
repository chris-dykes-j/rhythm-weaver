package handler 

import (
	"fmt"
	"html/template"
	"net/http"
	g "rhythm/pkg/generator"
	i "rhythm/pkg/renderer"
	"strconv"
)

type App struct {
	SequenceGenerator *g.SequenceGenerator
	ImageRenderer     *i.ImageRenderer
}

func NewApp(sg *g.SequenceGenerator, ir *i.ImageRenderer) *App {
	return &App{
		SequenceGenerator: sg,
		ImageRenderer:     ir,
	}
}

func (app *App) IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/views/index.gohtml"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *App) ImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	notes, err1 := strconv.ParseInt(r.FormValue("notes"), 10, 64)
	subdivision, err2 := strconv.ParseInt(r.FormValue("subdivision"), 10, 64)
	timeSignature := 4
	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid input parameters", http.StatusBadRequest)
		return
	}

	seq := app.SequenceGenerator.CreateSequence(int(notes), int(subdivision), timeSignature)
	img, _ := app.ImageRenderer.CreateImage(seq, int(subdivision), timeSignature)

	// Send image to client
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<img src="data:image/jpg;base64,%s" alt="Generated Image">`, img)
}
