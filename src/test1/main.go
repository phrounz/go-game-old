package main

import (

	//"image/color"
	"image/color"
	//"image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"
)

const (
	GAME_MODE_PRE_LOADING   = iota
	GAME_MODE_LOADING       = iota
	GAME_MODE_INTRODUCTION  = iota
	GAME_MODE_IN_GAME       = iota
	GAME_MODE_MISSING_CLUES = iota
	GAME_MODE_WIN_THE_GAME  = iota
)

var gameMode = GAME_MODE_PRE_LOADING

var imageLoading *ebiten.Image
var imageIntroduction *ebiten.Image
var imageMissingClues *ebiten.Image
var imageWinTheGame *ebiten.Image

var temple *Temple
var player *Player
var teleporter *Teleporter

var mplusNormalFont font.Face

//------------------------------------------------------------------------------

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

//------------------------------------------------------------------------------

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	screen.Fill(color.RGBA{R: 86, G: 86, B: 86, A: 255})

	if gameMode == GAME_MODE_PRE_LOADING {

		drawAtPos(screen, imageLoading, Pos{x: 0, y: 0})
		gameMode = GAME_MODE_LOADING

	} else if gameMode == GAME_MODE_LOADING {

		drawAtPos(screen, imageLoading, Pos{x: 0, y: 0})

		imageIntroduction = loadImageFromFile("data/misc/introduction.png")
		imageMissingClues = loadImageFromFile("data/misc/missing_clues.png")
		imageWinTheGame = loadImageFromFile("data/misc/win_the_game.png")

		loadTempleData()
		temple = NewTemple(imageIntroduction)
		player = NewPlayer()
		teleporter = NewTeleporter()

		gameMode = GAME_MODE_INTRODUCTION

	} else if gameMode == GAME_MODE_INTRODUCTION {

		drawAtPos(screen, imageIntroduction, Pos{x: 0, y: 0})
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			gameMode = GAME_MODE_IN_GAME
			temple.loadLevelCustomFading(1, 0, imageIntroduction)
			player.setPos(Pos{x: 301.0, y: 675.0})
			// temple.loadLevel(2, 0)
			// player.setPos(Pos{x: 650 /*688*/, y: 440 /*446*/})
		}

	} else if gameMode == GAME_MODE_MISSING_CLUES {

		drawAtPos(screen, imageMissingClues, Pos{x: 0, y: 0})
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			gameMode = GAME_MODE_IN_GAME
			temple.loadLevelCustomFading(3, 0, imageMissingClues)
			player.setPos(Pos{x: 282, y: 680})
		}

	} else if gameMode == GAME_MODE_WIN_THE_GAME {

		drawAtPos(screen, imageWinTheGame, Pos{x: 0, y: 0})
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			os.Exit(0)
		}

	} else if gameMode == GAME_MODE_IN_GAME {

		teleporter.update(player, temple)
		temple.update()
		temple.draw(screen)

		if !temple.isRotating() && !temple.isFading() {
			player.update(temple.getImageCollision())
			player.draw(screen)
		}

		temple.drawFront(screen)

		teleporter.drawClues(screen)

	} else {
		panic("unknown gameMode")
	}

	return nil
}

//------------------------------------------------------------------------------

func main() {
	imageLoading = loadImageFromFile("data/misc/loading.png")

	if len(os.Args) > 1 && os.Args[1] == "-fullscreen" {
		ebiten.SetFullscreen(true)
	}

	var ratio = 1.0
	var desktopHeight = getDesktopHeight()
	if desktopHeight < 1024 {
		ratio = float64(desktopHeight) / 1024.0 // 0.5
	}

	if err := ebiten.Run(update, 1024, 1024, ratio, "Ludum Dare 44"); err != nil {
		log.Fatal(err)
	}

}

//------------------------------------------------------------------------------
