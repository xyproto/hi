package main

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ship struct {
	image *ebiten.Image
	w     float64
	h     float64
	x     float64
	y     float64
	vx    float64
	vy    float64
}

var shipImage *ebiten.Image

func LoadShipImage() error {
}

func NewShip() *Ship {
	var ship Ship
	var err error
	ship.image, _, err = ebitenutil.NewImageFromFile("img/ship.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	ship.w = 16
	ship.h = 16
	ship.x = (screenWidth - ship.w) / 2.0
	ship.y = (screenHeight - ship.h) / 2.0
	return &ship
}

func (s Ship) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ship.x, ship.y)
	screen.DrawImage(ship.image, op)
}
