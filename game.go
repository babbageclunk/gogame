package main

import (
	"embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  int = 320
	screenHeight int = 240

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameCount  = 8
)

type Point struct {
	x, y int
}

func NewGame(img *ebiten.Image) *Game {
	return &Game{
		pos:       Point{30, 120},
		img:       img,
		direction: 1,
	}
}

type Game struct {
	ticks     int
	pos       Point
	direction int
	img       *ebiten.Image
}

func (g *Game) Update() error {
	g.ticks++
	g.pos.x += g.direction * 2

	if g.pos.x >= screenWidth {
		g.direction = -1
	}
	if g.pos.x <= 0 {
		g.direction = 1
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if g.direction < 0 {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(frameWidth), 0)
	}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(float64(g.pos.x), float64(g.pos.y))
	i := (g.ticks / 5) % frameCount
	sx, sy := frameOX+i*frameWidth, frameOY
	screen.DrawImage(g.img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (width, height int) {
	return screenWidth, screenHeight
}

//go:embed runner.png
var fs embed.FS

func readImage(name string) (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFileSystem(fs, name)
	if err != nil {
		return nil, err
	}
	return img, nil
}
