package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"
)

// For now, let's save the image to a directory. Long term, we may just want to generate them and send them to the client if it is possible.
// Weigh pros and cons later.
func makeImage(sequence []bool, subDivision int, timeSignature int) {
	imgs := getImageNames(sequence, subDivision)
	notes, err := getNotesImages(imgs)
	if err != nil {
		fmt.Errorf("Could not get images: %s", err)
	}

	height := notes[0].Bounds().Dy()
	length := 0
	for _, note := range notes {
		length = length + note.Bounds().Dx()
	}

	img := image.NewRGBA(image.Rect(0, 0, length, height))

	// Add images to canvas
	point := image.Point{0, 0}
	for _, note := range notes {
		r := note.Bounds()
		r = r.Add(point)
		draw.Draw(img, r, note, image.Point{0, 0}, draw.Over)
		point.X = point.X + note.Bounds().Dx()
	}

	// Save image
	path := filepath.Join("./images/result", createImageName(sequence))
	file, err := os.Create(path)
	if err != nil {
		fmt.Errorf("Could not create image file: %s", err)
	}
	defer file.Close()

	opt := jpeg.Options{
		Quality: 100,
	}

	err = jpeg.Encode(file, img, &opt)
	if err != nil {
		fmt.Errorf("Could not create image file: %s", err)
	}
}

func getNotesImages(images []string) ([]image.Image, error) {
	var result []image.Image
	for _, img := range images {
		path := filepath.Join(os.Getenv("MUSIC_NOTE_DIR"), img)
		file, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("Failed to open: %s", err)
		}
		defer file.Close()

		note, err := jpeg.Decode(file)
		if err != nil {
			return nil, fmt.Errorf("Failed to decode: %s", err)
		}
		result = append(result, note)
	}
	return result, nil
}

func getImageNames(sequence []bool, subDivision int) []string {
	var result []string
	var sb strings.Builder
	for i, k := range sequence {
		if k {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		if (i+1)%subDivision == 0 {
			sb.WriteString(".jpg")
			result = append(result, sb.String())
			sb.Reset()
		}
	}
	return result
}

func createImageName(sequence []bool) string {
	var sb strings.Builder
	for _, k := range sequence {
		if k {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
	}
	sb.WriteString(".jpg")
	return sb.String()
}
