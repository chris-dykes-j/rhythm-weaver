package main

import (
	"errors"
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
		x, err := chooseNonEmptyArray(options)                        // Pick option array
        if (err != nil) {
            break
        }
		y := rand.Intn(len(options[x]))                          // Pick element in option array
		result[x][y] = true                                      // Set option to true
		options[x] = append(options[x][:y], options[x][y+1:]...) // Remove element from options
		notes--
	}

	return result
}

// Createss options for the complex sequence.
func createComplexOptions(seq [][]bool) [][]int {
	result := make([][]int, len(seq))
	for i := range result {
		result[i] = make([]int, len(seq[i]))
		for k := range result[i] {
			result[i][k] = k
		}
	}

	return result
}

func chooseNonEmptyArray(arrays [][]int) (int, error) {
    var r []int
    for k, array := range arrays {
        if len(array) == 0 {
            continue
        } else {
            r = append(r, k)
        }
    }
    if len(r) == 0 {
        return -1, errors.New("No more options")
    }

    return r[rand.Intn(len(r))], nil
}
