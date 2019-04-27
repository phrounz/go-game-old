package main

import "math"

const RATIO_X_Y = 355.0 / 249.0
const MIDDLE_X = 512
const MIDDLE_Y = 675
const TOP_Y = 312 //530

func rotatePointWithPerspective2(pz Pos, angleInDegree float64) Pos {
	pz.y = (pz.y-MIDDLE_Y)*RATIO_X_Y + MIDDLE_Y
	var pTransformed Pos
	var angleInRadians = math.Pi * angleInDegree / 180.0
	if pTransformed.y < TOP_Y {
		pTransformed.x = pz.x
	} else if pTransformed.y > TOP_Y && pTransformed.y < MIDDLE_Y {
		var interpolYRatio = (pTransformed.y - TOP_Y) / (MIDDLE_Y - TOP_Y)
		pTransformed.x =
			(1.0-interpolYRatio)*pz.x +
				interpolYRatio*(math.Cos(angleInRadians)*(pz.x-MIDDLE_X)-math.Sin(angleInRadians)*(pz.y-MIDDLE_Y)+MIDDLE_X)
	} else {
		pTransformed.x = math.Cos(angleInRadians)*(pz.x-MIDDLE_X) - math.Sin(angleInRadians)*(pz.y-MIDDLE_Y) + MIDDLE_X
	}
	pTransformed.y = math.Sin(angleInRadians)*(pz.x-MIDDLE_X) + math.Cos(angleInRadians)*(pz.y-MIDDLE_Y) + MIDDLE_Y
	pTransformed.y = (pTransformed.y-MIDDLE_Y)/RATIO_X_Y + MIDDLE_Y
	return pTransformed
}

func rotatePointWithPerspective(pz Pos, angleInDegree float64) Pos {
	pz.y = (pz.y-MIDDLE_Y)*RATIO_X_Y + MIDDLE_Y
	var pTransformed Pos
	var angleInRadians = math.Pi * angleInDegree / 180.0
	pTransformed.x = math.Cos(angleInRadians)*(pz.x-MIDDLE_X) - math.Sin(angleInRadians)*(pz.y-MIDDLE_Y) + MIDDLE_X
	pTransformed.y = math.Sin(angleInRadians)*(pz.x-MIDDLE_X) + math.Cos(angleInRadians)*(pz.y-MIDDLE_Y) + MIDDLE_Y
	pTransformed.y = (pTransformed.y-MIDDLE_Y)/RATIO_X_Y + MIDDLE_Y
	return pTransformed
}
