package ui

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUI(t *testing.T) {
	w, h := 20, 10
	ui := New(w, h)

	btn := NewButton(w-10, h-4)
	btn.Position = image.Point{X: 5, Y: 4}
	ui.AddComponent(btn)

	txt := NewText("hello", 6)
	txt.Position = image.Point{X: 0, Y: 0}
	ui.AddComponent(txt)

	assert.Equal(t, 2, len(ui.components))

	// XXX render all components
	testCompareRender(t, []string{
		" XXX    ",
		"        ",
	}, renderAsText(ui.Render()))
}
