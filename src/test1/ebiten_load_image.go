// +build !USE_SELFCONTAINED_MODE

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func loadImageFromFile(filepath string) *ebiten.Image {
	var img, _, err = ebitenutil.NewImageFromFile(filepath, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
