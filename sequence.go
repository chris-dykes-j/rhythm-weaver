package main

import (
	"math/rand"
)

// Creates simple sequence.
func createSequence(notes int, subdivision int, timeSignature int) []bool {
	length := subdivision * timeSignature
	options := createOptions(length)
	result := make([]bool, length)

	for notes > 0 {
		if len(options) == 0 {
			break
		}
		i := rand.Intn(len(options))                    // Choose index for get option list.
		choice := options[i]                            // Choose option.
		result[choice] = true                           // Change the result's index.
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

// Creates a complex sequence
func createAutoComplexSeq(notes int, timeSignature int) [][]bool {
	result := make([][]bool, timeSignature)
	for i := range result {
		result[i] = make([]bool, rand.Intn(5)+1)
	}
	options := createComplexOptions(result)

	for notes > 0 {
        if len(options) == 0 {
            break
        }
        choice1 := rand.Intn(len(options))
        x := options[choice1].index
        choice2 := rand.Intn(len(options[choice1].options))
        y := options[choice1].options[choice2]
        result[x][y] = true

        options[choice1].options = append(options[choice1].options[:choice2], options[choice1].options[choice2+1:]...)
        if len(options[choice1].options) == 0 {
            options = append(options[:choice1], options[choice1+1:]...)
        }

		notes--
	}

	return result
}

type Tuple struct {
    index int
    options []int
}

// Createss options for the complex sequence.
func createComplexOptions(seq [][]bool) []Tuple {
    result := make([]Tuple, len(seq))
    for i := range result {
        result[i].index = i
        for j := range seq[i] {
            result[i].options = append(result[i].options, j)
        }
    }
	return result
}
