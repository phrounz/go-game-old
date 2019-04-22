package main

import (
	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
)

var img *ebiten.Image

func init() {

}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	//op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(50, 50)
	screen.DrawImage(img, nil) //op)
	ebitenutil.DrawRect(screen, 10, 10, 100, 100, color.RGBA{R: 100, G: 0, B: 0, A: 128})
	ebitenutil.DebugPrint(screen, "Hello, World!")
	return nil
}

func main() {
	var err error

	imgTmp, _, err := image.Decode(bytes.NewReader(images.Gophers_jpg))
	if err != nil {
		log.Fatal("main: " + err.Error())
	}
	img, _ = ebiten.NewImageFromImage(imgTmp, ebiten.FilterDefault)

	if err := ebiten.Run(update, 640, 480, 1.0, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
