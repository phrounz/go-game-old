package main

import "github.com/hajimehoshi/ebiten"

//------------------------------------------------------------------------------

type Teleporter struct {
	nbCluesFound    int   // 0,1,2,or 3
	clueFound       []int // -1,0,1,2
	isInClue        []bool
	imageParchments []*ebiten.Image
	cluePos         []Pos
}

//------------------------------------------------------------------------------

func NewTeleporter() *Teleporter {
	return &Teleporter{
		nbCluesFound: 0,
		clueFound:    []int{-1, -1, -1},
		cluePos:      []Pos{{x: 506, y: 728}, {x: 445, y: 671}, {x: 508, y: 622}},
		isInClue:     []bool{false, false, false},
		imageParchments: []*ebiten.Image{
			loadImageFromFile("data/misc/parchment0.png"),
			loadImageFromFile("data/misc/parchment1.png"),
			loadImageFromFile("data/misc/parchment2.png")}}
}

//------------------------------------------------------------------------------

func (te *Teleporter) update(player *Player, temple *Temple) {

	var collisionRequest = player.getCollisionRequestOnce()
	if collisionRequest != COLLISION_REQUEST_NONE {

		var prevOrientation = temple.getPrevOrientation()
		var orientation = temple.getOrientation()
		var pos = player.getPos()
		var level = temple.getLevel()

		if temple.getLevel() == 1 && collisionRequest == COLLISION_MODE_ROTATECAMERA && orientation == 0 && pos.y < 532 {
			temple.loadLevel(2, 0)
			player.setPos(Pos{x: 640, y: 430}) // 688 446

		} else if temple.getLevel() == 2 && collisionRequest == COLLISION_MODE_ROTATECAMERA && orientation == 0 && pos.y < 532 {
			temple.loadLevel(1, 0)
			player.setPos(Pos{x: 400, y: 240})

		} else if temple.getLevel() == 2 && collisionRequest == COLLISION_MODE_ROTATECAMERA_REVERSE && orientation == 2 {
			temple.loadLevel(3, 0)
			player.setPos(Pos{x: 240, y: 654})

		} else if temple.getLevel() == 3 && collisionRequest == COLLISION_MODE_ROTATECAMERA_REVERSE {
			temple.loadLevel(2, 2)
			player.setPos(Pos{x: 512, y: 843})

		} else if temple.getLevel() == 3 && collisionRequest == COLLISION_MODE_ROTATECAMERA {
			if te.nbCluesFound >= 3 {
				gameMode = GAME_MODE_WIN_THE_GAME
			} else {
				gameMode = GAME_MODE_MISSING_CLUES
			}

		} else {

			var newOrientation = -1000

			if collisionRequest == COLLISION_MODE_ROTATECAMERA_REVERSE {
				newOrientation = orientation - 1
				if newOrientation < 0 {
					newOrientation = 3
				}
			} else {
				newOrientation = orientation + 1
				if newOrientation > 3 {
					newOrientation = 0
				}
			}

			var newPos Pos
			if level == 1 {
				newPos = te.getPosAfterRotationLevel1(pos, prevOrientation, orientation, collisionRequest == COLLISION_MODE_ROTATECAMERA_REVERSE)
			} else {
				newPos = te.getPosAfterRotationLevel2(pos, prevOrientation, orientation, collisionRequest == COLLISION_MODE_ROTATECAMERA_REVERSE)
			}
			player.setPos(newPos)

			temple.setOrientation(newOrientation)
		}

	} else if !temple.isRotating() {
		te.isInClue[0] = false
		te.isInClue[1] = false
		te.isInClue[2] = false
		var orientation = temple.getOrientation()
		var level = temple.getLevel()
		var x = player.getPos().x
		var y = player.getPos().y
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			if level == 1 && orientation == 2 && x >= 480 && y >= 685 && x < 540 && y < 760 { //x >= 444 && y >= 629 && x < 566 && y < 652 {
				if te.clueFound[0] == -1 {
					te.clueFound[0] = te.nbCluesFound
					te.nbCluesFound++
				}
				te.isInClue[0] = true
			} else if level == 1 && orientation == 3 && x >= 419 && y >= 635 && x < 554 && y < 728 {
				if te.clueFound[1] == -1 {
					te.clueFound[1] = te.nbCluesFound
					te.nbCluesFound++
				}
				te.isInClue[1] = true
			} else if level == 2 && orientation == 3 && x >= 460 && y >= 575 && x < 582 && y < 675 {
				if te.clueFound[2] == -1 {
					te.clueFound[2] = te.nbCluesFound
					te.nbCluesFound++
				}
				te.isInClue[2] = true
			}
		}

	}
}

//------------------------------------------------------------------------------

func (te *Teleporter) drawClues(screen *ebiten.Image) {
	for i := 0; i < 3; i++ {
		if te.isInClue[i] {
			drawAtPos(screen, te.imageParchments[te.clueFound[i]], Pos{x: 0, y: 0})
		}
	}
}

//------------------------------------------------------------------------------

func (te *Teleporter) getPosAfterRotationLevel2(pos Pos, prevOrientation int, nextOrientation int, isReverseRotation bool) Pos {
	if isReverseRotation {
		switch nextOrientation {
		case 0:
			return Pos{x: 495, y: 463}
		case 1:
			return Pos{x: 420 * 0.1, y: 515 * 0.1}
		case 2:
			return Pos{x: 420 * 0.1, y: 515 * 0.1}
		case 3:
			return Pos{x: 440, y: 530} //{x: 495, y: 463}
		default:
		}
	} else {
		switch nextOrientation {
		case 0:
			return Pos{x: 0, y: 0}
		case 1:
			return Pos{x: 0, y: 0}
		case 2:
			return Pos{x: 512, y: 468}
		case 3:
			return Pos{x: 574, y: 454}
		default:
		}
	}
	return Pos{x: 0, y: 0}
}

//------------------------------------------------------------------------------

func (te *Teleporter) getPosAfterRotationLevel1(pos Pos, prevOrientation int, nextOrientation int, isReverseRotation bool) Pos {
	if isReverseRotation {
		switch nextOrientation {
		case 0:
			return Pos{x: 660 /*655*/, y: 545 /*562*/} //Pos{x: 482, y: 818}
		case 1:
			return Pos{x: 460, y: 795}
		case 2:
			return Pos{x: 475, y: 806} //{x: 460, y: 795} //Pos{x: 604, y: 694}
		case 3: //p.pos.y<532
			return Pos{x: 480, y: 642} //{x: 516, y: 583}
		default:
		}
	} else {
		switch nextOrientation {
		case 0:
			return Pos{x: 392, y: 746} //if pos.y < 532 {
		case 1:
			return Pos{x: 434, y: 654} //{x: 313, y: 664}
		case 2:
			return Pos{x: 604, y: 694}
		case 3:
			return Pos{x: 525 /*516*/, y: 570 /*583*/} //p.pos.y<532
		default:
		}
	}
	return Pos{x: 0, y: 0}
}

//------------------------------------------------------------------------------
