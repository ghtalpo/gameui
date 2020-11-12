// press space to toggle hide / show grouop of child components

package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	ui "github.com/martinlindhe/gameui"
)

const (
	width, height = 320, 200
	scale         = 1.
	fontName      = "_resources/font/open_dyslexic/OpenDyslexic3-Regular.ttf"
)

var (
	gui       = ui.New(width, height)
	font12, _ = ui.NewFont(fontName, 12, 72, ui.White)
	fps       = ui.NewText(font12)
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

func main() {

	grp := ui.NewGroup(210, 16)
	grp.Position = ui.Point{
		X: (width / 2) - (grp.Dimension.Width / 2),
		Y: (height / 2) - (grp.Dimension.Height / 2)}
	gui.AddComponent(grp)

	txt := ui.NewText(font12).SetText("press space to toggle visible")
	txt.Position = ui.Point{X: 0, Y: 0}
	grp.AddChild(txt)
	gui.AddComponent(fps)

	gui.AddKeyFunc(ui.KeySpace, func() error {
		if grp.IsHidden() {
			grp.Show()
		} else {
			grp.Hide()
		}
		return nil
	})

	gui.AddKeyFunc(ui.KeyQ, func() error {
		fmt.Println("q - QUITTING")
		return ui.GracefulExitError{}
	})

	game := NewGame()

	// Specify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(width*scale, height*scale)
	ebiten.SetWindowTitle("Dialog (UI Demo)")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
