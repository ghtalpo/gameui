package ui

import (
	"image"
	"image/draw"
	"log"
	"math"
)

// LoadTiles ...
func LoadTiles(imgFile string, tileWidth, tileHeight int) (res []image.Image) {
	img, err := OpenImage(imgFile)
	if err != nil {
		log.Println("ERROR:", err)
		return nil
	}

	b := img.Bounds()
	imgWidth := b.Max.X
	imgHeight := b.Max.Y

	cols := float64(imgWidth) / float64(tileWidth)
	if cols != math.Floor(cols) {
		log.Fatalf("image width %d is not evenly divisable by tile width %d", imgWidth, tileWidth)
	}

	rows := float64(imgHeight) / float64(tileHeight)
	if rows != math.Floor(rows) {
		log.Fatalf("image height %d is not evenly divisable by tile height %d", imgHeight, tileHeight)
	}

	// slice up image into tiles
	for row := 0; row < int(rows); row++ {
		for col := 0; col < int(cols); col++ {
			x0 := col * tileWidth
			y0 := row * tileHeight
			x1 := (col + 1) * tileWidth
			y1 := (row + 1) * tileHeight

			sr := image.Rect(x0, y0, x1, y1)
			dst := image.NewRGBA(image.Rect(0, 0, tileWidth, tileHeight))
			r := sr.Sub(sr.Min).Add(image.Point{0, 0})
			draw.Draw(dst, r, img, sr.Min, draw.Src)

			res = append(res, dst)
		}
	}
	return
}
