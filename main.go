package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	img, err := readImage("runner.png")
	check(err)
	sound, err := readWav("jab.wav", sampleRate)
	check(err)
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("render image")
	game, err := NewGame(img, sound)
	check(err)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
