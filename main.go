package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var images map[uint64]*ebiten.Image

const (
	SHIP = iota
	BULLET
)

const (
	W = 320
	H = 240

	SHIP_W = 16
	SHIP_H = 16

	BULLET_W = 2
	BULLET_H = 2
)

type Vec2 struct {
	x float64
	y float64
}

type Bullet struct {
	x    float64
	y    float64
	vx   float64
	vy   float64
	life uint64
}

var (
	ship    = Vec2{(W - SHIP_W) / 2, (H - SHIP_H) / 2}
	bullets []Bullet
)

// Game implements the ebiten Game interface
type Game struct{}

// Update proceeds the game state and is called every tick (1/60 s by default)
func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		ship.x--
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		ship.x++
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		ship.y--
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		ship.y++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		b := Bullet{ship.x, ship.y, 0, -1, 100}
		bullets = append(bullets, b)
	}
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errors.New("quit")
	}

	// Update bullets
	aliveBullets := make([]Bullet, 0, len(bullets))
	for i := 0; i < len(bullets); i++ {
		bullets[i].x += bullets[i].vx
		bullets[i].y += bullets[i].vy
		if bullets[i].life > 1 {
			bullets[i].life--
			aliveBullets = append(aliveBullets, bullets[i])
		}
	}
	bullets = aliveBullets

	return nil
}

// Draw is the render function and is called every frame (1/60s by default)
func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	//op.GeoM.Reset()
	op.GeoM.Translate(ship.x, ship.y)
	screen.DrawImage(images[SHIP], op)

	for i := 0; i < len(bullets); i++ {
		op := &ebiten.DrawImageOptions{}
		//op.GeoM.Reset()
		op.GeoM.Translate(bullets[i].x+SHIP_W/2-1, bullets[i].y)
		aliveRatio := float32(bullets[i].life) / float32(100.0)
		op.ColorScale.ScaleAlpha(aliveRatio)
		screen.DrawImage(images[BULLET], op)
	}

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return W, H
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

	err = loadImage("img/bullet.png", BULLET)
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
