package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Obstacle struct {
	X, Y, Width, Height float32
}

type Player struct {
	Sprite    *ebiten.Image
	X, Y      float64
	Direction string
}

type Game struct {
	Player   Player
	Obstacle Obstacle
}

func (g *Game) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		g.Player.Direction = ""
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Player.Direction = "Right"
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Player.Direction = "Left"
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Player.Direction = "Up"
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Player.Direction = "Down"
	}

	switch g.Player.Direction {
	case "Right":
		g.Player.X += 2
	case "Left":
		g.Player.X -= 2
	case "Up":
		g.Player.Y -= 2
	case "Down":
		g.Player.Y += 2
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// bg
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// environment
	vector.DrawFilledRect(
		screen,
		g.Obstacle.X,
		g.Obstacle.Y,
		g.Obstacle.Width,
		g.Obstacle.Height,
		color.RGBA{255, 0, 0, 255},
		false,
	)

	// player
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.Player.X, g.Player.Y)

	screen.DrawImage(
		g.Player.Sprite.SubImage(
			image.Rect(0, 0, 13, 13),
		).(*ebiten.Image),
		&opts,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	playerImg, _, err := ebitenutil.NewImageFromFile("./assets/images/pacman/pacman.png")
	if err != nil {
		log.Fatal(err)
	}

	obstacle := Obstacle{X: 150, Y: 150, Width: 32, Height: 32}

	game := &Game{
		Player: Player{
			Sprite: playerImg,
			X:      100,
			Y:      100,
		},
		Obstacle: obstacle,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
