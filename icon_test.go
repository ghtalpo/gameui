package ui

import (
	"image"
	"image/color"
	"testing"
)

func TestIconOnly(t *testing.T) {

	img := image.NewRGBA(image.Rect(0, 0, 3, 3))
	img.Set(0, 0, color.White)
	img.Set(2, 0, color.White)
	img.Set(1, 2, color.White)

	ico := NewIcon(img)

	// make sure same frame is delivered each time
	for i := 0; i < 10; i++ {
		im := ico.Draw(0, 0)
		testCompareRender(t, []string{
			"# #",
			"   ",
			" # ",
		}, renderAsText(im))
	}
}

func TestUIWithIconOnly(t *testing.T) {

	w, h := 9, 5
	ui := New(w, h)

	img := image.NewRGBA(image.Rect(0, 0, 3, 3))
	img.Set(0, 0, color.White)
	img.Set(2, 0, color.White)
	img.Set(1, 2, color.White)

	ico := NewIcon(img)
	ui.AddComponent(ico)

	// make sure same frame is delivered each time
	for i := 0; i < 10; i++ {
		im := ui.Render(0, 0)
		// XXX fails: blank output now
		testCompareRender(t, []string{
			"# #      ",
			"         ",
			" #       ",
			"         ",
			"         ",
		}, renderAsText(im))
	}
}