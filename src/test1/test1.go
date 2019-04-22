package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	ebitenutil.DrawRect(screen, 10, 10, 100, 100, color.RGBA{R: 100, G: 0, B: 0, A: 128})
	ebitenutil.DebugPrint(screen, "Hello, World!")
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
