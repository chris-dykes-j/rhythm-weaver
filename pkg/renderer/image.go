package renderer

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"
)

type ImageRenderer struct{}

func NewImageRenderer() *ImageRenderer {
	return &ImageRenderer{}
}

func (ir *ImageRenderer) CreateImage(sequence []bool, subDivision int, timeSignature int) (string, error) {
	imgs := ir.getImageNames(sequence, subDivision, timeSignature)
	notes, err := ir.getNotesImages(imgs)
	if err != nil {
		fmt.Errorf("Could not get images: %s", err)
	}

	height := notes[0].Bounds().Dy()
	length := 0
	for _, note := range notes {
		length = length + note.Bounds().Dx()
	}
	img := image.NewRGBA(image.Rect(0, 0, length, height))

	point := image.Point{0, 0}
	for _, note := range notes {
		r := note.Bounds()
		r = r.Add(point)
		draw.Draw(img, r, note, image.Point{0, 0}, draw.Over)
		point.X = point.X + note.Bounds().Dx()
	}

	var buffer bytes.Buffer
	err = jpeg.Encode(&buffer, img, nil)
	if err != nil {
		return "", err
	}
	imgBase64Str := base64.StdEncoding.EncodeToString(buffer.Bytes())

	return imgBase64Str, nil
}

func (ir *ImageRenderer) getNotesImages(images []string) ([]image.Image, error) {
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

func (ir *ImageRenderer) getImageNames(sequence []bool, subDivision int, timeSignature int) []string {
	var result []string
    switch timeSignature {
    case 4:
        result = append(result, "44.jpg")
    case 3:
        result = append(result, "34.jpg")
    }

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
