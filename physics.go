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

func generatePower(power int, variabilityCoefficient int) int {
	if power < 0 || power > 4 {
		panic("Incorrect power!")
	}
	if rand.Intn(10) > variabilityCoefficient {
		if power == 4 {
			return 3
		} else if power == 0{
			return 1
		} else{
			return power+rand.Intn(2)*2-1
		}
	} else{
		return power
	}
}

func generateRandomPower(power int) int {
	if power < 0 || power > 4 {
		print("Incorrect power");
		//panic("Incorrect power!")
	}
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

func generateRotate(rotate int, variabilityCoefficient int) int {
	if rotate < -90 || rotate > 90 {
		panic("Incorrect rotate!")
	}
	if rand.Intn(10) > variabilityCoefficient {
		return generateRandomRotate(rotate)
	} else{
		return rotate
	}
}

func generateRandomRotate(rotate int) int {
	if rotate < -90 || rotate > 90 {
		panic("Incorrect rotate!")
	}
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
	return ShuttleState{x, y, hSpeed, vSpeed, fuel, s.rotate, s.power}
}
