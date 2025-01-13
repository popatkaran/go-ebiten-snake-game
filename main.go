package main

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	// direction co-ordinates logics based on ebiten game engine
	dirUp    = Point{0, -1}
	dirDown  = Point{0, 1}
	dirLeft  = Point{-1, 0}
	dirRight = Point{1, 0}
)

const (
	gameSpeed   = time.Second / 6
	screenWidth = 640
	screenHeigh = 480
	gridSize    = 16
)

type Point struct {
	x, y int
}

type Game struct {
	snake      []Point
	direction  Point
	lastUpdate time.Time
	food       Point
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.direction = dirUp
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.direction = dirDown
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.direction = dirLeft
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.direction = dirRight
	}

	if time.Since(g.lastUpdate) < gameSpeed {
		return nil
	}
	g.lastUpdate = time.Now()
	g.updateSnake(&g.snake, g.direction)

	return nil
}

func (g *Game) updateSnake(snake *[]Point, direction Point) {
	head := (*snake)[0]

	newHead := Point{
		x: head.x + direction.x,
		y: head.y + direction.y,
	}

	if newHead == g.food {
		*snake = append(
			[]Point{newHead},
			*snake...,
		)
		g.spawnFood()
	} else {
		*snake = append(
			[]Point{newHead},
			(*snake)[:len(*snake)-1]...,
		)
	}
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

	vector.DrawFilledRect(
		screen,
		float32(g.food.x*gridSize),
		float32(g.food.y*gridSize),
		gridSize,
		gridSize,
		color.RGBA{255, 0, 0, 255},
		true,
	)
}

func (g *Game) Layout(
	outsideWidth, outsideHeight int,
) (int, int) {
	return screenWidth, screenHeigh
}

func (g *Game) spawnFood() {
	g.food = Point{rand.Intn(screenWidth / gridSize), rand.Intn(screenHeigh / gridSize)}
}

func main() {
	g := &Game{
		snake: []Point{{
			x: screenWidth / gridSize / 2,
			y: screenWidth / gridSize / 2,
		}},
		direction: Point{0, 0},
	}

	g.spawnFood()

	ebiten.SetWindowSize(screenWidth, screenHeigh)
	ebiten.SetWindowTitle("Snake Game")

	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}

}
