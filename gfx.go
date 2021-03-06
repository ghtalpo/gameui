package ui

import (
	"image"
	"image/color"
	"image/draw"
	"io"
	"log"
	"os"
)

// hLine draws a horizontal line
func hLine(img draw.Image, x1, x2, y int, col color.Color) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// vLine draws a veritcal line
func vLine(img draw.Image, x, y1, y2 int, col color.Color) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

// DrawRect ...
func DrawRect(img draw.Image, r image.Rectangle, col color.Color) {
	vLine(img, r.Min.X, r.Min.Y, r.Max.Y, col) // left
	vLine(img, r.Max.X, r.Min.Y, r.Max.Y, col) // right

	hLine(img, r.Min.X+1, r.Max.X-1, r.Min.Y, col) // top
	hLine(img, r.Min.X+1, r.Max.X-1, r.Max.Y, col) // bottom
}

// OpenImage loads an image from file. based on Open from disintegration/imaging
// https://github.com/disintegration/imaging/blob/master/helpers.go#L68
func OpenImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, err := DecodeImage(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// RequireImage loads an image or exits if not successful
func RequireImage(filename string) image.Image {
	img, err := OpenImage(filename)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

// DecodeImage reads an image from r
func DecodeImage(r io.Reader) (image.Image, error) {
	img, _, err := image.Decode(r)
	return img, err
}
