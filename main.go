package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth = 640
	screenHeigh = 480
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(
	outsideWidth, outsideHeight int,
) (int, int) {
	return screenWidth, screenHeigh
}

func main() {
	g := &Game{}

	ebiten.SetWindowSize(screenWidth, screenHeigh)
	ebiten.SetWindowTitle("Snake Game")

	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}

}
