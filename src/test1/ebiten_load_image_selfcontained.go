// +build USE_SELFCONTAINED_MODE

package main

import (
	"bytes"
	"image"
	"log"

	// IMPORTANT: path for local tests: replace to build in local with -tags USE_SELFCONTAINED_MODE
	//"./data_go"
	// IMPORTANT: path for JSGO:
	"github.com/phrounz/go-game/src/test1/data_go"

	"github.com/hajimehoshi/ebiten"
)

func loadImageFromFile(filepath string) *ebiten.Image {

	log.Print("Loading: " + filepath)
	var b = data_go.GetBytesFromFilename(filepath)

	// load image from data
	var imgTmp, _, err1 = image.Decode(bytes.NewReader(b))
	if err1 != nil {
		log.Fatal("main: " + err1.Error())
	}
	var img, err2 = ebiten.NewImageFromImage(imgTmp, ebiten.FilterDefault)
	if err2 != nil {
		log.Fatal("main: " + err2.Error())
	}
	return img
}
