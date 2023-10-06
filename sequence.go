package main

import (
	"fmt"
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

func getRandomKey(m map[int][]int) int {
    keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
    return keys[rand.Intn(len(keys))]
}

func getRandomElement(array []int) int {
    return rand.Intn(len(array))
}

// Creates a complex sequence
func createAutoComplexSeq(notes int, timeSignature int) [][]bool {
	result := make([][]bool, timeSignature)
	for i := range result {
		result[i] = make([]bool, rand.Intn(5)+1)
	}
	options := createComplexOptions(result)
    fmt.Println(options)

	for notes > 0 {
        x := getRandomKey(options)
        y := getRandomElement(options[x])
        choice := options[x][y]
		if result[x][choice] == true {
            fmt.Println(options)
            fmt.Println(result)
            fmt.Printf("%d, %d\n", x, y)
		}
		result[x][y] = true                                      // Set option to true
		options[x] = append(options[x][:y], options[x][y+1:]...) // Remove element from options
        if len(options[x]) == 0 {
            delete(options, x)
        }
		notes--
	}

	return result
}

// Createss options for the complex sequence.
func createComplexOptions(seq [][]bool) map[int][]int {
	result := make(map[int][]int)
	for i, arr := range seq {
        result[i] = createOptions(len(arr))
	}

	return result
}
