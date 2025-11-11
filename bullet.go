package main

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Bullet struct {
	x    float64
	y    float64
	vx   float64
	vy   float64
	life uint64
}

var bulletImage *ebiten.Image
var bulletW = 2
var bulletH = 2

func init() {
	var err error
	bulletImage, _, err = ebitenutil.NewImageFromFile("img/bullet.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
