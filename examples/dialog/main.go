// a dialog window, with a yes & no button

package main

import (
	"fmt"
	"log"
	"os"

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
	font20, _ = ui.NewFont(fontName, 20, 72, ui.White)
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
	exit := ui.NewText(font20).SetText("exit?")
	exit.Position = ui.Point{X: width/2 - exit.GetWidth()/2, Y: height / 3}
	gui.AddComponent(exit)

	btnYes := ui.NewButton(60, 20).SetText(font12, "YES")
	btnYes.Position = ui.Point{X: width/4 - btnYes.Dimension.Width/2, Y: height / 2}
	btnYes.OnClick = func() {
		fmt.Println("clicked", btnYes.Text.GetText())
		os.Exit(0)
	}
	gui.AddComponent(btnYes)

	btnNo := ui.NewButton(60, 20).SetText(font12, "NO")
	btnNo.Position = ui.Point{X: (width/4)*3 - btnYes.Dimension.Width/2, Y: height / 2}
	btnNo.OnClick = func() {
		fmt.Println("clicked", btnNo.Text.GetText())
	}
	gui.AddComponent(btnNo)

	gui.AddComponent(fps)

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
