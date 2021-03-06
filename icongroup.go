package ui

import (
	"image"
	"image/color"
	"image/draw"
	"log"
)

// IconGroup is a tile-hased grid display of object icons
type IconGroup struct {
	component
	columns, rows         int
	iconWidth, iconHeight int               // size of each icon
	objects               []IconGroupObject // holds the icons to display
	borderColor           color.Color
}

// IconGroupObject is something that is contained in the icon group
type IconGroupObject interface {
	Name() string
	Icon() image.Image
	Click()
	ID() uint64
}

const (
	iconBorderPad = 1
)

var (
	icongroupBorderColor = color.RGBA{0x50, 0x50, 0x50, 192} // grey, 75% transparent
)

// NewIconGroup ...
func NewIconGroup(columns, rows, iconWidth, iconHeight int) *IconGroup {
	pad := 2 // 1 px border, + 1 px cell padding
	componentWidth := (columns * iconWidth) + (pad * 2)
	componentHeight := (rows * iconHeight) + (pad * 2)
	grp := IconGroup{}
	grp.backgroundColor = Transparent
	grp.borderColor = icongroupBorderColor
	grp.columns = columns
	grp.rows = rows
	grp.Dimension.Width = componentWidth
	grp.Dimension.Height = componentHeight
	grp.iconWidth = iconWidth
	grp.iconHeight = iconHeight
	return &grp
}

// SetBorderColor sets the border color
func (grp *IconGroup) SetBorderColor(c color.Color) {
	grp.borderColor = c
}

// Draw redraws internal buffer
func (grp *IconGroup) Draw(mx, my int) *image.RGBA {
	if grp.isHidden {
		grp.isClean = true
		return nil
	}
	if grp.isClean {
		return grp.Image
	}
	grp.initImage()

	// draw outline
	outlineRect := image.Rect(0, 0, grp.Dimension.Width-1, grp.Dimension.Height-1)
	DrawRect(grp.Image, outlineRect, grp.borderColor)

	grp.drawIcons(mx, my)

	grp.isClean = true
	return grp.Image
}

// AddObject adds an object to display in the group
func (grp *IconGroup) AddObject(o IconGroupObject) {
	grp.objects = append(grp.objects, o)
	grp.isClean = false
}

// RemoveObjectByID ...
func (grp *IconGroup) RemoveObjectByID(id uint64) {
	for i, c := range grp.objects {
		if c.ID() == id {
			grp.objects = append(grp.objects[:i], grp.objects[i+1:]...)
			grp.isClean = false
			return
		}
	}
}

// RemoveAllObjects removes all displayed content
func (grp *IconGroup) RemoveAllObjects() {
	grp.objects = nil
	grp.isClean = false
}

func (grp *IconGroup) drawIcons(mx, my int) {
	x := iconBorderPad + 1
	y := iconBorderPad + 1
	col := 0
	row := 0

	for _, o := range grp.objects {

		img := o.Icon()
		if img == nil {
			log.Println("ERROR: UI IconGroup object", o.Name(), "lacks icon")
			continue
		}
		b := img.Bounds()
		w := b.Max.X
		h := b.Max.Y
		x1 := x + w
		y1 := y + h

		dr := image.Rect(x, y, x1, y1)
		draw.Draw(grp.Image, dr, img, image.ZP, draw.Over)
		x += w
		col++
		if col >= grp.columns {
			col = 0
			x = iconBorderPad + 1
			y += h
			row++
		}
		if row >= grp.rows {
			break
		}
	}
}

// Click pass click to child icon (click has happened)
func (grp *IconGroup) Click(mouse Point) bool {

	x := iconBorderPad + 1
	y := iconBorderPad + 1
	col := 0
	row := 0

	childPoint := Point{X: mouse.X - grp.Position.X, Y: mouse.Y - grp.Position.Y}

	for _, c := range grp.objects {
		b := c.Icon().Bounds()
		x1 := x + b.Max.X
		y1 := y + b.Max.Y
		r := image.Rect(x, y, x1, y1)
		if childPoint.In(r) {
			c.Click()
			// XXX mark click consumed so it dont re-trigger
			return true
		}

		x += b.Max.X
		col++
		if col >= grp.columns {
			col = 0
			x = iconBorderPad + 1
			y += b.Max.Y
			row++
		}
	}
	return false
}
