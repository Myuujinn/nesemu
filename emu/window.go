package emu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

const (
	ResolutionWidth  = 256
	ResolutionHeight = 240
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 256, 240
}

func Run(rom *string, scale int) {
	ebiten.SetWindowSize(ResolutionWidth*scale, ResolutionHeight*scale)
	ebiten.SetWindowTitle("nesemu")

	log.Printf("ROM: %s", *rom)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
