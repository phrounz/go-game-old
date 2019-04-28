package main

import (
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten"
)

type Temple struct {
	prevImageBeforeLevel *ebiten.Image
	level                int
	images               []*ebiten.Image
	imagesIntermediate   []*ebiten.Image
	imagesFront          []*ebiten.Image
	imagesCollision      []*ebiten.Image
	previousOrientation  int
	orientation          int
	endOfRotationMs      int64
	endOfFadingMs        int64
}

func (t *Temple) getPrevOrientation() int { return t.previousOrientation }
func (t *Temple) getOrientation() int     { return t.orientation }

const ROTATION_DURATION_MS = 300
const FADING_DURATION_IN_MS = 300
const FADING_DURATION_OUT_MS = 300

const RATIO_X_Y = 355.0 / 249.0
const MIDDLE_X = 512
const MIDDLE_Y = 675

var globalImages = make([][]*ebiten.Image, 3, 3)
var globalImagesFront = make([][]*ebiten.Image, 3, 3)
var globalImagesCollision = make([][]*ebiten.Image, 3, 3)
var globalImagesIntermediate = make([][]*ebiten.Image, 3, 3)

func loadTempleData() {

	for i := 0; i < 3; i++ {
		var level = i + 1
		var prefix = string("data/level" + strconv.Itoa(level))

		globalImages[i] = make([]*ebiten.Image, 4, 4)
		globalImagesFront[i] = make([]*ebiten.Image, 4, 4)
		globalImagesCollision[i] = make([]*ebiten.Image, 4, 4)
		globalImagesIntermediate[i] = make([]*ebiten.Image, 22, 22)
		for j := 0; j < 4; j++ {
			globalImages[i][j] = nil
			globalImagesFront[i][j] = nil
			globalImagesCollision[i][j] = nil
		}
		for j := 0; j < 22; j++ {
			globalImagesIntermediate[i][j] = nil
		}

		globalImages[i][0] = loadImageFromFile(prefix + "/0.png")
		if level != 3 {
			globalImages[i][1] = loadImageFromFile(prefix + "/1.png")
			globalImages[i][2] = loadImageFromFile(prefix + "/2.png")
			globalImages[i][3] = loadImageFromFile(prefix + "/3.png")
		}

		if level == 1 {
			globalImagesFront[i][0] = loadImageFromFile(prefix + "/0-front.png")
			globalImagesFront[i][3] = loadImageFromFile(prefix + "/3-front.png")
		}

		globalImagesCollision[i][0] = loadImageFromFile(prefix + "/0-collision.png")
		if level != 3 {
			globalImagesCollision[i][1] = loadImageFromFile(prefix + "/1-collision.png")
			globalImagesCollision[i][2] = loadImageFromFile(prefix + "/2-collision.png")
			globalImagesCollision[i][3] = loadImageFromFile(prefix + "/3-collision.png")
		}

		if level != 3 {
			globalImagesIntermediate[i][0] = loadImageFromFile(prefix + "/small/m-15.png")
			globalImagesIntermediate[i][1] = loadImageFromFile(prefix + "/small/m-30.png")
			globalImagesIntermediate[i][2] = loadImageFromFile(prefix + "/small/m-45.png")
			globalImagesIntermediate[i][3] = loadImageFromFile(prefix + "/small/m-60.png")
			globalImagesIntermediate[i][4] = loadImageFromFile(prefix + "/small/m-75.png")

			globalImagesIntermediate[i][5] = loadImageFromFile(prefix + "/small/m-90.png")
			globalImagesIntermediate[i][6] = loadImageFromFile(prefix + "/small/m-105.png")
			globalImagesIntermediate[i][7] = loadImageFromFile(prefix + "/small/m-120.png")
			globalImagesIntermediate[i][8] = loadImageFromFile(prefix + "/small/m-135.png")
			globalImagesIntermediate[i][9] = loadImageFromFile(prefix + "/small/m-150.png")

			globalImagesIntermediate[i][10] = loadImageFromFile(prefix + "/small/m-165.png")
			globalImagesIntermediate[i][11] = loadImageFromFile(prefix + "/small/m-180.png")
			globalImagesIntermediate[i][12] = loadImageFromFile(prefix + "/small/p-165.png")
			globalImagesIntermediate[i][13] = loadImageFromFile(prefix + "/small/p-150.png")
			globalImagesIntermediate[i][14] = loadImageFromFile(prefix + "/small/p-135.png")
			globalImagesIntermediate[i][15] = loadImageFromFile(prefix + "/small/p-120.png")

			globalImagesIntermediate[i][16] = loadImageFromFile(prefix + "/small/p-105.png")
			globalImagesIntermediate[i][17] = loadImageFromFile(prefix + "/small/p-90.png")
			globalImagesIntermediate[i][18] = loadImageFromFile(prefix + "/small/p-75.png")
			globalImagesIntermediate[i][19] = loadImageFromFile(prefix + "/small/p-45.png")
			globalImagesIntermediate[i][20] = loadImageFromFile(prefix + "/small/p-30.png")
			globalImagesIntermediate[i][21] = loadImageFromFile(prefix + "/small/p-15.png")
		}
	}

}

//------------------------------------------------------------------------------

func NewTemple(imageIntroduction *ebiten.Image) *Temple {
	var t = Temple{
		prevImageBeforeLevel: imageIntroduction,
		level:                -1,
		previousOrientation:  0,
		orientation:          0,
		images:               make([]*ebiten.Image, 4, 4),
		imagesFront:          make([]*ebiten.Image, 4, 4),
		imagesCollision:      make([]*ebiten.Image, 4, 4),
		imagesIntermediate:   make([]*ebiten.Image, 22, 22),
		endOfRotationMs:      -1,
		endOfFadingMs:        -1}
	return &t
}

func (t *Temple) getLevel() int {
	return t.level
}

func (t *Temple) loadLevel(level int, orientation int) {
	if t.level != -1 {
		t.loadLevelCustomFading(level, orientation, t.images[t.orientation])
	} else {
		t.loadLevelCustomFading(level, orientation, nil)
	}
}

func (t *Temple) loadLevelCustomFading(level int, orientation int, prevImageBeforeLevel *ebiten.Image) {

	t.prevImageBeforeLevel = prevImageBeforeLevel

	for i := 0; i < 4; i++ {
		t.images[i] = globalImages[level-1][i]
		t.imagesFront[i] = globalImagesFront[level-1][i]
		t.imagesCollision[i] = globalImagesCollision[level-1][i]
	}
	for i := 0; i < 22; i++ {
		t.imagesIntermediate[i] = globalImagesIntermediate[level-1][i]
	}

	t.level = level
	t.orientation = orientation
	t.endOfFadingMs = nowMs() + FADING_DURATION_OUT_MS + FADING_DURATION_IN_MS
}

//------------------------------------------------------------------------------

func (t *Temple) setOrientation(newOrientation int) {
	t.endOfRotationMs = nowMs() + ROTATION_DURATION_MS
	t.previousOrientation = t.orientation
	t.orientation = newOrientation
}

//------------------------------------------------------------------------------

func (t *Temple) update() {
	if t.endOfRotationMs != -1 && nowMs() >= t.endOfRotationMs {
		t.endOfRotationMs = -1
	}
	if t.endOfFadingMs != -1 && nowMs() >= t.endOfFadingMs {
		t.endOfFadingMs = -1
	}
}

//------------------------------------------------------------------------------

func (t *Temple) isRotating() bool {
	return t.endOfRotationMs != -1
}

func (t *Temple) isFading() bool {
	return t.endOfFadingMs != -1
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
		var accurate = float64(t.previousOrientation)*5.0*(1.0-ratioDone) + float64(nextOrientation)*5.0*ratioDone + 2.5
		var indexFloor = int(math.Floor(accurate)) % 22
		var indexCeil = int(math.Ceil(accurate)) % 22
		var indexRound = int(math.Round(accurate)) % 22
		if indexRound < 0 {
			indexRound += 22
		}
		if indexFloor < 0 {
			indexFloor += 22
		}
		if indexCeil < 0 {
			indexCeil += 22
		}
		if t.imagesIntermediate[indexRound] != nil {
			var w, h = t.imagesIntermediate[indexRound].Size()
			drawAtPosScaled(screen, t.imagesIntermediate[indexRound], posLeftTop, 1024.0/float64(w), 1024.0/float64(h))
		}
		if indexRound == indexFloor && accurate-float64(indexFloor) < 1.0 {
			if t.imagesIntermediate[indexCeil] != nil {
				var w, h = t.imagesIntermediate[indexCeil].Size()
				drawAtPosWithScaleAndOpacity(
					screen, t.imagesIntermediate[indexCeil], posLeftTop, 1024.0/float64(w), 1024.0/float64(h),
					(accurate - float64(indexFloor)))
			}
		} else if indexRound == indexCeil && float64(indexCeil)-accurate < 1.0 {
			if t.imagesIntermediate[indexFloor] != nil {
				var w, h = t.imagesIntermediate[indexFloor].Size()
				drawAtPosWithScaleAndOpacity(
					screen, t.imagesIntermediate[indexFloor], posLeftTop, 1024.0/float64(w), 1024.0/float64(h),
					(float64(indexCeil) - accurate))
			}
		}

		//------
	} else if t.endOfFadingMs != -1 {

		var timeRemaining = float64(t.endOfFadingMs - nowMs())
		if timeRemaining > float64(FADING_DURATION_IN_MS) {
			timeRemaining -= float64(FADING_DURATION_IN_MS)
			if t.prevImageBeforeLevel != nil {
				drawAtPosWithOpacity(
					screen, t.prevImageBeforeLevel,
					posLeftTop,
					timeRemaining/float64(FADING_DURATION_IN_MS))
			}
		} else {
			drawAtPosWithOpacity(
				screen, t.images[t.orientation],
				posLeftTop,
				1.0-(timeRemaining/float64(FADING_DURATION_OUT_MS)))
		}
		//------
	} else {

		drawAtPos(screen, t.images[t.orientation], posLeftTop)
		//drawAtPos(screen, t.imagesCollision[t.orientation], posLeftTop)
	}
	return
}

//------------------------------------------------------------------------------

func (t *Temple) drawFront(screen *ebiten.Image) {
	if t.imagesFront[t.orientation] != nil && !t.isRotating() && t.endOfFadingMs == -1 {
		var posLeftTop = Pos{x: 0, y: 0}
		drawAtPos(screen, t.imagesFront[t.orientation], posLeftTop)
	}
}

//------------------------------------------------------------------------------

func (t *Temple) getImageCollision() *ebiten.Image {
	return t.imagesCollision[t.orientation]
}

//------------------------------------------------------------------------------
