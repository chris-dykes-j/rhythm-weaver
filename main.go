package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func main() {
	http.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir("src"))))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/image", imageHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/index.gohtml"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
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

	seq := createSequence(int(notes), int(subdivision), timeSignature)
	img, _ := makeImage(seq, int(subdivision), timeSignature)

	// Send image to client
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<img src="data:image/jpg;base64,%s" alt="Generated Image">`, img)
}
