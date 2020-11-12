package main

import (
	"fmt"
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	ui "github.com/martinlindhe/gameui"
)

const (
	width, height = 320, 200
	scale         = 2.
	fontName      = "_resources/font/open_dyslexic/OpenDyslexic3-Regular.ttf"
)

var (
	gui     *ui.UI
	fps     *ui.Text
	hp      *ui.Bar
	hp2     *ui.Bar
	mana    *ui.Bar
	lastInc time.Time
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

	expired := time.Now().Add(-1 * time.Second)
	if lastInc.Before(expired) {
		lastInc = time.Now()

		hp.IncValue(1)
		hp2.IncValue(1)
		mana.IncValue(2)
	}
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

	heart, err := ui.OpenImage("_resources/tile/7x7_heart.png")
	if err != nil {
		log.Fatal(err)
	}

	hp = ui.NewBar(width-2, 7+2)
	hp.SetValue(0)
	hp.SetFillImage(heart)
	hp.Position = ui.Point{X: width/2 - hp.Dimension.Width/2, Y: 20}
	gui.AddComponent(hp)

	hp2 = ui.NewBar(width-2, 7)
	hp2.SetValue(0)
	hp2.SetFillColor(ui.Red)
	hp2.Position = ui.Point{X: width/2 - hp.Dimension.Width/2, Y: 30}
	gui.AddComponent(hp2)

	mana = ui.NewBar(width-2, 16)
	mana.SetValue(25)
	mana.SetFillColor(ui.Blue)
	mana.Position = ui.Point{X: width/2 - hp.Dimension.Width/2, Y: height - 16 - 20}
	gui.AddComponent(mana)

	lastInc = time.Now()
}

func main() {
	game := NewGame()

	// Specify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(width*scale, height*scale)
	ebiten.SetWindowTitle("Bars (UI Demo)")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
