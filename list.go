package ui

import (
	"image"
)

// List holds a number of rows of text, each is clickable (UI component)
type List struct {
	component
	rowHeight int
}

// Line defines the interface for lines of text usable with the List object
type Line interface {
	Text() string
}

// NewList ...
func NewList(width, height int) *List {
	lst := List{}
	lst.backgroundColor = Transparent
	lst.Dimension = Dimension{Width: width, Height: height}
	lst.rowHeight = 12
	return &lst
}

// SetRowHeight sets the list row height
func (lst *List) SetRowHeight(n int) {
	lst.rowHeight = n
}

// addChild ...
func (lst *List) addChild(c Component) {
	lst.children = append(lst.children, c)
	lst.isClean = false
}

// AddLine ...
func (lst *List) AddLine(l Line, fnt *Font, fnc func()) {
	h := NewText(fnt)
	h.OnClick = fnc
	h.SetText(l.Text())
	h.Position = Point{X: 0, Y: len(lst.children) * lst.rowHeight}
	h.Dimension = Dimension{Width: lst.Dimension.Width, Height: lst.rowHeight}
	lst.addChild(h)
	lst.isClean = false
}

// Draw redraws internal buffer
func (lst *List) Draw(mx, my int) *image.RGBA {
	if lst.isHidden {
		lst.isClean = true
		return nil
	}
	if lst.isClean {
		if lst.isChildrenClean() {
			return lst.Image
		}
		lst.isClean = false
	}
	lst.initImage()

	// draw background color
	lst.drawChildren(mx, my)

	lst.isClean = true
	return lst.Image
}

// Click pass click to window child components
func (lst *List) Click(mouse Point) bool {
	childPoint := Point{X: mouse.X - lst.Position.X, Y: mouse.Y - lst.Position.Y}
	for _, c := range lst.children {
		if childPoint.In(c.GetBounds()) {
			c.Click(childPoint)
			return true
		}
	}
	return false
}
