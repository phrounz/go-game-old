package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

type Temple struct {
	images             []*ebiten.Image
	imagesIntermediate []*ebiten.Image
	imagesFront        []*ebiten.Image
	//imageTmpRotation    []byte
	//imageTmpRotationE   *ebiten.Image
	imagesCollision     []*ebiten.Image
	previousOrientation int
	orientation         int
	endOfRotationMs     int64
}

func (t *Temple) getPrevOrientation() int { return t.previousOrientation }
func (t *Temple) getOrientation() int     { return t.orientation }

const ROTATION_DURATION_MS = 300

//------------------------------------------------------------------------------

func NewTemple() *Temple {
	var t = Temple{
		previousOrientation: 0,
		orientation:         0,
		images:              make([]*ebiten.Image, 4, 4),
		imagesIntermediate:  make([]*ebiten.Image, 22, 22),
		imagesFront:         make([]*ebiten.Image, 4, 4),
		imagesCollision:     make([]*ebiten.Image, 4, 4),
		endOfRotationMs:     -1}
	t.images[0] = loadImageFromFile("data/temple-0.png")
	t.images[1] = loadImageFromFile("data/temple-1.png")
	t.images[2] = loadImageFromFile("data/temple-2.png")
	t.images[3] = loadImageFromFile("data/temple-3.png")

	t.imagesFront[0] = loadImageFromFile("data/temple-0_front.png")
	t.imagesFront[1] = nil
	t.imagesFront[2] = nil
	t.imagesFront[3] = loadImageFromFile("data/temple-3_front.png")

	t.imagesIntermediate[0] = loadImageFromFile("data/small/templem-15.png")
	t.imagesIntermediate[1] = loadImageFromFile("data/small/templem-30.png")
	t.imagesIntermediate[2] = loadImageFromFile("data/small/templem-45.png")
	t.imagesIntermediate[3] = loadImageFromFile("data/small/templem-60.png")
	t.imagesIntermediate[4] = loadImageFromFile("data/small/templem-75.png")

	t.imagesIntermediate[5] = loadImageFromFile("data/small/templem-90.png")
	t.imagesIntermediate[6] = loadImageFromFile("data/small/templem-105.png")
	t.imagesIntermediate[7] = loadImageFromFile("data/small/templem-120.png")
	t.imagesIntermediate[8] = loadImageFromFile("data/small/templem-135.png")
	t.imagesIntermediate[9] = loadImageFromFile("data/small/templem-150.png")

	t.imagesIntermediate[10] = loadImageFromFile("data/small/templem-165.png")
	t.imagesIntermediate[11] = loadImageFromFile("data/small/templem-180.png")
	t.imagesIntermediate[12] = loadImageFromFile("data/small/templep-165.png")
	t.imagesIntermediate[13] = loadImageFromFile("data/small/templep-150.png")
	t.imagesIntermediate[14] = loadImageFromFile("data/small/templep-135.png")
	t.imagesIntermediate[15] = loadImageFromFile("data/small/templep-120.png")

	t.imagesIntermediate[16] = loadImageFromFile("data/small/templep-105.png")
	t.imagesIntermediate[17] = loadImageFromFile("data/small/templep-90.png")
	t.imagesIntermediate[18] = loadImageFromFile("data/small/templep-75.png")
	t.imagesIntermediate[19] = loadImageFromFile("data/small/templep-45.png")
	t.imagesIntermediate[20] = loadImageFromFile("data/small/templep-30.png")
	t.imagesIntermediate[21] = loadImageFromFile("data/small/templep-15.png")

	t.imagesCollision[0] = loadImageFromFile("data/temple-0_collision.png")
	t.imagesCollision[1] = loadImageFromFile("data/temple-1_collision.png")
	t.imagesCollision[2] = loadImageFromFile("data/temple-2_collision.png")
	t.imagesCollision[3] = loadImageFromFile("data/temple-3_collision.png")
	return &t
}

//------------------------------------------------------------------------------

func (t *Temple) update(newOrientation int) {
	if newOrientation != NO_NEW_ORIENTATION {
		t.endOfRotationMs = nowMs() + ROTATION_DURATION_MS
		t.previousOrientation = t.orientation
		t.orientation = newOrientation
	}

	if t.endOfRotationMs != -1 && nowMs() >= t.endOfRotationMs {
		t.endOfRotationMs = -1
	}
	/*if ebiten.IsKeyPressed(ebiten.KeyE) && !t.isRotating() {
		t.endOfRotationMs = nowMs() + ROTATION_DURATION_MS
		t.previousOrientation = t.orientation
		t.orientation -= 1
		if t.orientation < 0 {
			t.orientation = 3
		}
	}*/
}

//------------------------------------------------------------------------------

func (t *Temple) isRotating() bool {
	return t.endOfRotationMs != -1
}

//------------------------------------------------------------------------------

func (t *Temple) draw(screen *ebiten.Image) {
	var posLeftTop = Pos{x: 0, y: 0}
	if t.endOfRotationMs != -1 {
		var ratioDone = 1.0 - (float64(t.endOfRotationMs-nowMs()))/float64(ROTATION_DURATION_MS)
		var nextOrientation = t.orientation
		if nextOrientation < t.previousOrientation-2 {
			nextOrientation += 4
		}
		if nextOrientation > t.previousOrientation+2 {
			nextOrientation -= 4
		}
		var index = int(math.Round(float64(t.previousOrientation)*5.0*(1.0-ratioDone) + float64(nextOrientation)*5.0*ratioDone + 2.5))
		index = index % 22
		if index < 0 {
			index += 22
		}
		var img = t.imagesIntermediate[index]
		var w, h = img.Size()
		drawAtPosScaled(screen, img, posLeftTop, 1024.0/float64(w), 1024.0/float64(h))
	} else {
		drawAtPos(screen, t.images[t.orientation], posLeftTop)
		//drawAtPos(screen, t.imagesCollision[t.orientation], posLeftTop)
	}
	return
}

//------------------------------------------------------------------------------

func (t *Temple) drawFront(screen *ebiten.Image) {
	if t.imagesFront[t.orientation] != nil && !t.isRotating() {
		var posLeftTop = Pos{x: 0, y: 0}
		drawAtPos(screen, t.imagesFront[t.orientation], posLeftTop)
	}
}

//------------------------------------------------------------------------------

func (t *Temple) getImageCollision() *ebiten.Image {
	return t.imagesCollision[t.orientation]
}

//------------------------------------------------------------------------------
