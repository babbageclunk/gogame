package main

import (
	"embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed runner.png jab.wav
var fs embed.FS

func readImage(name string) (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFileSystem(fs, name)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func readWav(name string, sampleRate int) (*wav.Stream, error) {
	f, err := fs.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return wav.DecodeWithSampleRate(sampleRate, f)
}
