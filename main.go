package main

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var images map[uint64]*ebiten.Image

const (
	SHIP = iota
)

// Game implements the ebiten Game interface
type Game struct{}

// Update proceeds the game state and is called every tick (1/60 s by default)
func (g *Game) Update() error {
	// ...
	return nil
}

// Draw is the render function and is called every frame (1/60s by default)
func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(images[SHIP], nil)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func loadImage(filename string, imageID uint64) error {
	img, _, err := ebitenutil.NewImageFromFile(filename)
	if err != nil {
		return err
	}
	images[imageID] = img
	return nil
}

func main() {
	// Load resources
	images = make(map[uint64]*ebiten.Image, 0)

	err := loadImage("img/ship.png", SHIP)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	game := &Game{}

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("hi")

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
