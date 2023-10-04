package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
)

func main() {
    // fmt.Println(createSequence(5, 2, 4))

    http.HandleFunc("/", indexHandler)
    http.ListenAndServe(":8080", nil)
}

func createSequence(notes int, subdivision int, timeSignature int) []bool {
    length := subdivision * timeSignature
    options := createOptions(length) 
    result := make([]bool, length)

    for notes > 0 {
        i := rand.Intn(len(options)) // Choose index for get option list.
        choice := options[i] // Choose option.
        result[choice] = true // Change the result's index.
        options = append(options[:i], options[i+1:]...) // Remove option.
        notes--
    }

    return result
}

// Creates a sequence of numbers from zero to length exclusive.
func createOptions(length int) []int {
    options := make([]int, length)
    for i := range options {
        options[i] = i
    }
    return options
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("views/index.gohtml"))
    err := tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: send image to client based off form data
}
