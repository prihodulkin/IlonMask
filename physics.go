package main

import (
	"math"
)

const g float64 = 3.711
const width = 7000
const height = 3000

var anglesSin [181]float64
var anglesCos [181]float64

func initAngles() {
	for ind := 0; ind < 181; ind++ {
		angle := float64(ind) / 180 * math.Pi
		anglesSin[ind] = math.Sin(angle)
		anglesCos[ind] = math.Cos(angle)
	}
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


