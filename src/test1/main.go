package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var img *ebiten.Image

func init() {
	/*	var err error
		img, _, err = ebitenutil.NewImageFromFile("gopher.png", ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}*/
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	/*op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(50, 50)
	screen.DrawImage(img, op)*/
	ebitenutil.DrawRect(screen, 10, 10, 100, 100, color.RGBA{R: 100, G: 0, B: 0, A: 128})
	ebitenutil.DebugPrint(screen, "Hello, World!")
	return nil
}

func main() {
	if err := ebiten.Run(update, 640, 480, 1.0, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
