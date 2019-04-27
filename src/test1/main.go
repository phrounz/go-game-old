package main

import (

	//"image/color"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	//"github.com/hajimehoshi/ebiten/examples/resources/images"
)

var temple *Temple
var player *Player
var playerWin = false

var mplusNormalFont font.Face

func init() {
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.Fill(color.RGBA{R: 86, G: 86, B: 86, A: 255})

	if playerWin {
		var w, h = screen.Size()
		text.Draw(screen, "You win!", mplusNormalFont, w/2-40, h/2-10, color.White)
		return nil
	}

	temple.update(player.getAndResetRotationInfo())
	temple.draw(screen)

	player.update(temple.getImageCollision(), temple.getPrevOrientation(), temple.getOrientation())
	if !temple.isRotating() {
		player.draw(screen)
	}

	temple.drawFront(screen)

	//ebitenutil.DrawRect(screen, 10, 10, 100, 100, color.RGBA{R: 100, G: 0, B: 0, A: 128})
	//ebitenutil.DebugPrint(screen, "Hello, World!")
	return nil
}

func main() {

	temple = NewTemple()
	player = NewPlayer()

	//ebiten.SetFullscreen(true)

	if err := ebiten.Run(update, 1024, 1024, 1.0, "Hello, World!"); err != nil {
		log.Fatal(err)
	}

	//ebiten.ScreenSizeInFullscreen(1024, 1024)
}
