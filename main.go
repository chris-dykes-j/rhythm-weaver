package main

import (
	"fmt"
	"math/rand"
)

func main() {
    fmt.Println(createSequence(8, 2, 4))
    fmt.Println(createSequence(5, 2, 4))
    fmt.Println(createSequence(5, 2, 4))
    fmt.Println(createSequence(5, 2, 4))
    fmt.Println(createSequence(5, 2, 4))
    fmt.Println(createSequence(5, 2, 4))
    fmt.Println(createSequence(5, 2, 4))
    fmt.Println(createSequence(5, 2, 4))
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

