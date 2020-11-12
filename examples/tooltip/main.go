package main

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	ui "github.com/martinlindhe/gameui"
)

const (
	width, height = 413, 240
	scale         = 2.
	fontName      = "_resources/font/open_dyslexic/OpenDyslexic3-Regular.ttf"
)

var (
	gui        *ui.UI
	fps        *ui.Text
	mana       *ui.Bar
	bar1       *ui.Bar
	bar2       *ui.Bar
	bar3       *ui.Bar
	lastInc    time.Time
	background image.Image
)

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

	if lastInc.Before(time.Now().Add(-1 * time.Second)) {
		lastInc = time.Now()
		mana.IncValue(2)
		bar1.IncValue(5)
		bar2.IncValue(3)
		bar3.IncValue(1)
		mana.SetTooltip(fmt.Sprintf("mana = %d", mana.GetValue()))
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	rect := image.Rect(0, 0, gui.Dimension.Width, gui.Dimension.Height)
	frame := image.NewRGBA(rect)
	draw.Draw(frame, rect, background, image.ZP, draw.Src)
	draw.Draw(frame, rect, gui.Render(), image.ZP, draw.Over)

	eframe := ebiten.NewImageFromImage(frame) //, ebiten.FilterNearest)
	screen.DrawImage(eframe, &ebiten.DrawImageOptions{})
	return
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func init() {
	var err error
	background, err = ui.OpenImage("_resources/gfx/fiveyears.jpg")
	if err != nil {
		log.Fatal(err)
	}

	gui = ui.New(width, height)
	fnt, err := ui.NewFont(fontName, 12, 72, ui.White)
	if err != nil {
		log.Fatal(err)
	}
	fps = ui.NewText(fnt)
	gui.AddComponent(fps)
	gui.AddKeyFunc(ui.KeyQ, func() error {
		fmt.Println("q - QUITTING")
		return ui.GracefulExitError{}
	})

	mana = ui.NewBar(width-2, 16)
	mana.SetValue(25)
	mana.SetFillColor(ui.Blue)
	mana.Position = ui.Point{X: 0, Y: (height / 2) + (height / 4)}
	mana.SetTooltip(fmt.Sprintf("mana = %d", mana.GetValue()))
	gui.AddComponent(mana)

	grpWidth := 40
	grp := ui.NewGroup(grpWidth, 100)
	grp.Position = ui.Point{X: (width / 2) - (grpWidth / 2), Y: 10}
	bar1 = ui.NewBar(grpWidth, 10)
	bar1.Position = ui.Point{X: 0, Y: 0}
	bar1.SetValue(10)
	bar1.SetFillColor(ui.Red)
	bar1.SetTooltip("bar 1")
	grp.AddChild(bar1)

	bar2 = ui.NewBar(grpWidth, 10)
	bar2.Position = ui.Point{X: 0, Y: 15}
	bar2.SetValue(20)
	bar2.SetFillColor(ui.Green)
	bar2.SetTooltip("bar 2")
	grp.AddChild(bar2)

	bar3 = ui.NewBar(grpWidth, 10)
	bar3.Position = ui.Point{X: 0, Y: 30}
	bar3.SetValue(40)
	bar3.SetFillColor(ui.Yellow)
	bar3.SetTooltip("bar 3")
	grp.AddChild(bar3)

	gui.AddComponent(grp)

	lastInc = time.Now()
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
