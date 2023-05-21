package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	img, err := readImage("runner.png")
	check(err)
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("render image")
	if err := ebiten.RunGame(NewGame(img)); err != nil {
		log.Fatal(err)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
