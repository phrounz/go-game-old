package main

import (

	//"image/color"
	_ "image/jpeg"
	_ "image/png"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	//"github.com/hajimehoshi/ebiten/examples/resources/images"
)

// https://github.com/hajimehoshi/ebiten/issues/289
const FRAME_DURATION = 1.0 / 60.0
const (
	//RATIO_X_Y              = 1.7136
	RATIO_Y_Y_STAIRS_GREEN = 1.2
	RATIO_Y_Y_STAIRS_RED   = 1.0 + 1.2
	RATIO_Y_Y_STAIRS_BLUE  = 1.0 + 1.2
)

const SPEED = 150.0

const COLLISION_REQUEST_NONE = -1000

type Player struct {
	hasJustRotated   bool
	collisionRequest int
	pos              Pos
	imgLeftTop       *ebiten.Image
	imgLeftBottom    *ebiten.Image
	imgRightTop      *ebiten.Image
	imgRightBottom   *ebiten.Image
	imgToDraw        *ebiten.Image
}

func NewPlayer() *Player {
	var p = &Player{
		hasJustRotated:   true,
		collisionRequest: COLLISION_REQUEST_NONE,
		//pos:            Pos{x: 301.0, y: 664.0},
		pos:            Pos{x: 301.0, y: 675.0},
		imgLeftTop:     loadImageFromFile("data/misc/player_lefttop.png"),
		imgLeftBottom:  loadImageFromFile("data/misc/player_leftbottom.png"),
		imgRightTop:    loadImageFromFile("data/misc/player_righttop.png"),
		imgRightBottom: loadImageFromFile("data/misc/player_rightbottom.png")}
	p.imgToDraw = p.imgRightBottom
	return p
}

const (
	COLLISION_MODE_OBSTACLE             = iota
	COLLISION_MODE_ROTATECAMERA         = iota
	COLLISION_MODE_ROTATECAMERA_REVERSE = iota
	COLLISION_MODE_NORMAL               = iota
	COLLISION_MODE_STAIRS_GREEN         = iota
	COLLISION_MODE_STAIRS_RED           = iota
	COLLISION_MODE_STAIRS_BLUE          = iota
)

func getCollisionMode(collisionImage *ebiten.Image, pos Pos) int {
	var r, g, b, a = collisionImage.At(int(pos.x), int(pos.y)).RGBA()
	if r > 88*256 && r < 168*256 && g > 88*256 && g < 168*256 && b > 88*256 && b < 168*256 && a > 200*256 { // black
		return COLLISION_MODE_ROTATECAMERA
	} else if r > 128*256 && g < 150*256 && b > 128*256 && a > 200*256 {
		return COLLISION_MODE_ROTATECAMERA_REVERSE
	} else if r < 128*256 && g < 220*256 && b > 128*256 && a > 200*256 {
		return COLLISION_MODE_STAIRS_BLUE
	} else if r < 128*256 && g > 128*256 && b < 128*256 && a > 200*256 {
		return COLLISION_MODE_STAIRS_GREEN
	} else if r > 128*256 && g < 128*256 && b < 128*256 && a > 200*256 {
		return COLLISION_MODE_STAIRS_RED
	} else if r < 10*256 && g < 10*256 && b < 10*256 && a > 200*256 { // black
		return COLLISION_MODE_NORMAL
	} else if a < 10*256 {
		return COLLISION_MODE_OBSTACLE
	} else {
		//panic("getCollisionMode")
		return COLLISION_MODE_NORMAL
	}
}

func (p *Player) update(collisionImage *ebiten.Image) {

	var cm = getCollisionMode(collisionImage, p.pos)
	var nextPosSafe Pos
	var nextPos Pos
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		nextPosSafe = p.pos.Add(FRAME_DURATION*SPEED*RATIO_X_Y, FRAME_DURATION*SPEED)
		switch cm {
		case COLLISION_MODE_STAIRS_RED:
			nextPos = p.pos.Add(FRAME_DURATION*SPEED*RATIO_X_Y, FRAME_DURATION*SPEED*RATIO_Y_Y_STAIRS_RED)
		case COLLISION_MODE_STAIRS_GREEN:
			nextPos = p.pos.Add(FRAME_DURATION*SPEED*RATIO_X_Y, FRAME_DURATION*SPEED*(1.0-RATIO_Y_Y_STAIRS_GREEN))
		default:
			nextPos = nextPosSafe
		}
		p.imgToDraw = p.imgRightBottom
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		nextPosSafe = p.pos.Add(FRAME_DURATION*SPEED*RATIO_X_Y*-1.0, FRAME_DURATION*SPEED*-1.0)
		switch cm {
		case COLLISION_MODE_STAIRS_RED:
			nextPos = p.pos.Add(FRAME_DURATION*SPEED*RATIO_X_Y*-1.0, FRAME_DURATION*SPEED*RATIO_Y_Y_STAIRS_RED*-1.0)
		case COLLISION_MODE_STAIRS_GREEN:
			nextPos = p.pos.Add(FRAME_DURATION*SPEED*RATIO_X_Y*-1.0, FRAME_DURATION*SPEED*(-1.0+RATIO_Y_Y_STAIRS_GREEN))
		default:
			nextPos = nextPosSafe
		}
		p.imgToDraw = p.imgLeftTop
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		switch cm {
		case COLLISION_MODE_STAIRS_BLUE:
			nextPos = p.pos.Add(FRAME_DURATION*SPEED*RATIO_X_Y, FRAME_DURATION*SPEED*RATIO_Y_Y_STAIRS_BLUE*-1.0)
		default:
			nextPos = p.pos.Add(FRAME_DURATION*SPEED*RATIO_X_Y, FRAME_DURATION*SPEED*-1.0)
		}
		p.imgToDraw = p.imgRightTop
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		switch cm {
		case COLLISION_MODE_STAIRS_BLUE:
			nextPos = p.pos.Add(FRAME_DURATION*SPEED*RATIO_X_Y*-1.0, FRAME_DURATION*SPEED*RATIO_Y_Y_STAIRS_BLUE)
		default:
			nextPos = p.pos.Add(FRAME_DURATION*SPEED*RATIO_X_Y*-1.0, FRAME_DURATION*SPEED)
		}
		p.imgToDraw = p.imgLeftBottom
	}

	if (cm == COLLISION_MODE_ROTATECAMERA || cm == COLLISION_MODE_ROTATECAMERA_REVERSE) && !p.hasJustRotated {
		p.hasJustRotated = true
		p.collisionRequest = cm
	} else {
		p.hasJustRotated = false
		if getCollisionMode(collisionImage, nextPosSafe) == COLLISION_MODE_NORMAL {
			p.pos = nextPosSafe
		} else if getCollisionMode(collisionImage, nextPos) != COLLISION_MODE_OBSTACLE {
			p.pos = nextPos
		}
	}
}

func (p *Player) getCollisionRequestOnce() int {
	var cr = p.collisionRequest
	p.collisionRequest = COLLISION_REQUEST_NONE
	return cr
}

func (p *Player) getPos() Pos {
	return p.pos
}

func (p *Player) setPos(pos Pos) {
	p.pos = pos
}

func ff(value float64) string {
	return strconv.FormatFloat(value, 'g', 6, 64)
}

func (p *Player) draw(screen *ebiten.Image) {
	var width, height = p.imgToDraw.Size()
	drawAtPos(screen, p.imgToDraw, p.pos.AddPos(&Pos{float64(width) * -0.5, float64(height) * -0.5}))
}
