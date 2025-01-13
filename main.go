package main

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth = 640
	screenHeigh = 480
	gridSize    = 16
)

type Point struct {
	x, y int
}

type Game struct {
	snake []Point
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, p := range g.snake {
		vector.DrawFilledRect(
			screen,
			float32(p.x*gridSize),
			float32(p.y*gridSize),
			gridSize,
			gridSize,
			color.White,
			true,
		)
	}

}

func (g *Game) Layout(
	outsideWidth, outsideHeight int,
) (int, int) {
	return screenWidth, screenHeigh
}

func main() {
	g := &Game{
		snake: []Point{{
			x: screenWidth / gridSize / 2,
			y: screenWidth / gridSize / 2,
		}},
	}

	ebiten.SetWindowSize(screenWidth, screenHeigh)
	ebiten.SetWindowTitle("Snake Game")

	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}

}
