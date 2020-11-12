package main

import (
	"fmt"
	"image/color"
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten"
	ui "github.com/martinlindhe/gameui"
)

const (
	width, height = 320, 200
	scale         = 2.
	fontName      = "_resources/font/open_dyslexic/OpenDyslexic3-Regular.ttf"
)

var (
	gui    *ui.UI
	fps    *ui.Text
	window = ui.NewWindow(200, 12*8, "Window Title")
	list   = ui.NewList(200, 12*8)

	windowBgColor     = color.RGBA{0x50, 0x50, 0x50, 192} // gray, 75% transparent
	windowTitleColor  = color.RGBA{0x50, 0x50, 0x50, 192} // gray
	windowBorderColor = color.RGBA{0x5c, 0x63, 0x69, 192} // gray
)

type line struct {
	text string
}

func (l line) Text() string {
	return l.text
}

// Game implements ebiten.Game interface.
type Game struct{}

// NewGame is
func NewGame() *Game {
	return &Game{}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	if err := gui.Update(); err != nil {
		return err
	}

	fps.SetText(fmt.Sprintf("%.1f", ebiten.CurrentFPS()))
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	frame := ebiten.NewImageFromImage(gui.Render()) //, ebiten.FilterNearest)
	screen.DrawImage(frame, &ebiten.DrawImageOptions{})
	return
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func init() {
	list.SetRowHeight(14)

	gui = ui.New(width, height)
	fnt, err := ui.NewFont(fontName, 12, 72, ui.White)
	if err != nil {
		log.Fatal(err)
	}
	greenFont, err := ui.NewFont(fontName, 12, 72, ui.Green)
	if err != nil {
		log.Fatal(err)
	}

	fps = ui.NewText(fnt)
	gui.AddComponent(fps)
	gui.AddKeyFunc(ui.KeyQ, func() error {
		fmt.Println("q - QUITTING")
		return ui.GracefulExitError{}
	})

	// center window on screen
	boxW := 70
	boxH := 50
	x0 := (width / 2) - boxW
	y0 := (height / 2) - boxH
	window.Dimension = ui.Dimension{Width: boxW * 2, Height: boxH * 2}
	window.Position = ui.Point{X: x0, Y: y0}
	window.SetBackgroundColor(windowBgColor)
	window.SetBorderColor(windowBorderColor)
	window.SetTitleColor(windowTitleColor)
	window.AddChild(list)

	list.AddLine(line{text: "entry one"}, greenFont, func() {
		log.Println("Entry one clicked")
	})

	list.AddLine(line{text: "entry two"}, greenFont, func() {
		log.Println("Entry two clicked")
	})

	gui.AddComponent(window)
}

func main() {
	game := NewGame()

	// Specify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(width*scale, height*scale)
	ebiten.SetWindowTitle("Tooltip (UI Demo)")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
