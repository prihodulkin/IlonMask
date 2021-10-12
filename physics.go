package main

import (
	"math"
	"math/rand"
)

const g float64 = 3.711
const width = 7000
const height = 3000

var anglesSin [181]float64
var anglesCos [181]float64

type Surface struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

type Ground []Surface

type point interface {
	X() float64
	Y() float64
}

type ShuttleState struct {
	x        float64
	y        float64
	hSpeed   float64
	vSpeed   float64
	fuel     float64
	rotate   int
	power    int
	routeInd int
}

func (state ShuttleState) SetPower(power int) {
	if power < 0 {
		state.power = 0
	} else if power > 4 {
		state.power = 4
	} else {
		state.power = power
	}
}

func (state ShuttleState) SetRotate(rotate int) {
	if rotate < -90 {
		state.rotate = -90
	} else if rotate > 90 {
		state.rotate = 90
	} else {
		state.rotate = rotate
	}
}
func (state ShuttleState) X() float64 {
	return state.x
}

func (state ShuttleState) Y() float64 {
	return state.y
}

func initAngles() {
	for ind := 0; ind < 181; ind++ {
		angle := float64(ind-90) / 360 * 2 * math.Pi
		anglesSin[ind] = math.Sin(angle)
		anglesCos[ind] = math.Cos(angle)
	}
}

//проверяет, x,y соответствуют пересечению поверхности
func isLandedOrCrashedOnSurface(surface Surface, x float64, y float64) bool {
	x1 := float64(surface.x1)
	x2 := float64(surface.x2)
	y1 := float64(surface.y1)
	y2 := float64(surface.y2)
	return (x-x1)/(x2-x1)*(y2-y1)+y1 >= y
}

func findFlatSurface(ground Ground) {
	for _, s := range ground {
		if isFlat(s) {
			xFlatMin = float64(s.x1)
			xFlatMax = float64(s.x2)
			break

		}
	}
}

func isFlat(s Surface) bool {
	return s.y2 == s.y1
}

func isLandedOrCrashed(surfaces []Surface, x float64, y float64) bool {
	for _, s := range surfaces {
		x1 := float64(s.x1)
		x2 := float64(s.x2)
		if x < x1 || x > x2 {
			continue
		}
		return isLandedOrCrashedOnSurface(s, x, y)
	}
	return false
}

func generateRandomPower(power int) int {
	min := power - 1
	if min < 0 {
		min = 0
	}
	max := power + 1
	if max > 4 {
		max = 4
	}
	step := rand.Intn(max - min + 1)
	return min + step
}

func generateRandomRotate(rotate int) int {
	min := rotate - 15
	if min < -90 {
		min = -90
	}
	max := rotate + 15
	if max > 90 {
		max = 90
	}
	step := rand.Intn(max - min + 1)
	return min + step

}

func move(s ShuttleState, time float64) ShuttleState {
	vA := float64(s.power)*anglesCos[s.rotate+90] - g
	hA := -float64(s.power) * anglesSin[s.rotate+90]
	hSpeed := s.hSpeed + hA*time
	vSpeed := s.vSpeed + vA*time
	x := s.x + (hSpeed+s.hSpeed)/2*time
	y := s.y + (vSpeed+s.vSpeed)/2*time
	fuel := s.fuel - float64(s.power)*time
	return ShuttleState{x, y, hSpeed, vSpeed, fuel, s.rotate, s.power, s.routeInd}
}