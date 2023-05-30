package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	sampleRate = 48000

	screenWidth  int = 320
	screenHeight int = 240

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameCount  = 8
)

var redImage = ebiten.NewImage(1, 1)

func init() {
	redImage.Fill(color.RGBA{0xff, 0, 0, 0xff})
}

type Point struct {
	x, y int
}

func PosPoint(x, y int) Point {
	return Point{x, y}
}

type BBox struct {
	topleft, bottomright Point
}

func (b BBox) Contains(p Point) bool {
	return (b.topleft.x <= p.x &&
		p.x <= b.bottomright.x &&
		b.topleft.y <= p.y &&
		p.y <= b.bottomright.y)
}

func (b BBox) Width() int {
	return b.bottomright.x - b.topleft.x
}

func (b BBox) Height() int {
	return b.bottomright.y - b.topleft.y
}

func NewGame(img *ebiten.Image, sound *wav.Stream) (*Game, error) {
	audioCtx := audio.NewContext(sampleRate)
	audioPlayer, err := audioCtx.NewPlayer(sound)
	if err != nil {
		return nil, err
	}
	return &Game{
		audioPlayer: audioPlayer,
		pos:         Point{30, 120},
		img:         img,
		direction:   1,
	}, nil
}

type Game struct {
	ticks       int
	pos         Point
	direction   int
	img         *ebiten.Image
	audioPlayer *audio.Player
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
	if g.guyClicked() {
		g.audioPlayer.Rewind()
		g.audioPlayer.Play()
	}
	return nil
}

func (g *Game) guyClicked() bool {
	return g.bbox().Contains(PosPoint(ebiten.CursorPosition())) &&
		inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}

func (g *Game) bbox() BBox {
	return BBox{
		topleft:     Point{g.pos.x - frameWidth/2 + 8, g.pos.y - frameHeight/2},
		bottomright: Point{g.pos.x + frameWidth/2 - 8, g.pos.y + frameHeight/2},
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw a red box around the guy
	var boxOp ebiten.DrawImageOptions
	bbox := g.bbox()
	boxOp.GeoM.Scale(float64(bbox.Width()), float64(bbox.Height()))
	boxOp.GeoM.Translate(-float64(bbox.Width())/2, -float64(bbox.Height())/2)
	boxOp.GeoM.Translate(float64(g.pos.x), float64(g.pos.y))
	screen.DrawImage(redImage, &boxOp)

	var op ebiten.DrawImageOptions
	if g.direction < 0 {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(frameWidth), 0)
	}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(float64(g.pos.x), float64(g.pos.y))
	i := (g.ticks / 5) % frameCount
	sx, sy := frameOX+i*frameWidth, frameOY
	screen.DrawImage(g.img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), &op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (width, height int) {
	return screenWidth, screenHeight
}
