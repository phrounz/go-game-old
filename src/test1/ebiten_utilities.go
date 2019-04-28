package main

import (
	"bytes"
	"image"
	//"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	//"github.com/hajimehoshi/ebiten/examples/resources/images"
)

func loadImageFromFile(filepath string) *ebiten.Image {
	var img, _, err = ebitenutil.NewImageFromFile(filepath, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func loadImageFromData(b []byte) *ebiten.Image {
	var imgTmp, _, err = image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal("main: " + err.Error())
	}
	var img, _ = ebiten.NewImageFromImage(imgTmp, ebiten.FilterDefault)
	return img
}

func drawAtPos(screen *ebiten.Image, image *ebiten.Image, pos Pos) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(pos.x, pos.y)
	screen.DrawImage(image, op)
}

func drawAtPosWithOpacity(screen *ebiten.Image, image *ebiten.Image, pos Pos, opacity float64) {
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1.0, 1.0, 1.0, opacity)
	// hue128 := 0
	// saturation128 := 128
	// value128 := 128 * opacity
	// hue := float64(hue128) * 2 * math.Pi / 128
	// saturation := float64(saturation128) / 128
	// value := float64(value128) / 128
	// op.ColorM.ChangeHSV(hue, saturation, value)
	op.GeoM.Translate(pos.x, pos.y)
	screen.DrawImage(image, op)
}

func drawAtPosWithScaleAndOpacity(screen *ebiten.Image, image *ebiten.Image, pos Pos, scaleW float64, scaleH float64, opacity float64) {
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1.0, 1.0, 1.0, opacity)
	op.GeoM.Translate(pos.x, pos.y)
	op.GeoM.Scale(scaleW, scaleH)
	screen.DrawImage(image, op)
}

func drawAtPosScaled(screen *ebiten.Image, image *ebiten.Image, pos Pos, scaleW float64, scaleH float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(pos.x, pos.y)
	op.GeoM.Scale(scaleW, scaleH)
	screen.DrawImage(image, op)
}
