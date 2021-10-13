package main

import "math"

var xFlatMin float64
var xFlatMax float64

func fitness(route Route) float64 {
	l := len(route)
	state := route[l-1]
	return fitnessState(state)// + routeRotateFitness(route[len(route)-10:])

}

func fitnessState(state ShuttleState) float64 {
	return math.Pow(landingFitness(state.x, xFlatMin, xFlatMax),2 ) +
		math.Pow(vSpeedFitness(state.vSpeed), 2) +
		math.Pow(hSpeedFitness(state.hSpeed), 2)
	//+ math.Abs(float64(state.rotate))
}

func isResultWithoutRotate(state ShuttleState) bool {
	return landingFitness(state.x, xFlatMin, xFlatMax)+
		vSpeedFitness(state.vSpeed)+
		hSpeedFitness(state.hSpeed) == 0
}

func isResult(state ShuttleState) bool {
	return landingFitness(state.x, xFlatMin, xFlatMax)+
		vSpeedFitness(state.vSpeed)+
		hSpeedFitness(state.hSpeed)/*+math.Abs(float64(state.rotate))*/ == 0
}

func landingFitness(x float64, x1 float64, x2 float64) float64 {
	if x < x1 {
		return x1 - x
	} else if x > x2 {
		return x - x2
	} else {
		return 0
	}
}

func vSpeedFitness(vSpeed float64) float64 {
	vSpeedAbs := math.Abs(vSpeed)
	if vSpeedAbs <= 40 {
		return 0
	} else {
		return vSpeedAbs - 40
	}
}

func hSpeedFitness(hSpeed float64) float64 {
	hSpeedAbs := math.Abs(hSpeed)
	if hSpeedAbs <= 20 {
		return 0
	} else {
		return hSpeedAbs - 20
	}
}

func rotateFitness(rotate int) float64 {
	rotateAbs := math.Round(math.Abs(float64(rotate)))
	if rotateAbs <= 15 {
		return 0
	} else {
		return rotateAbs - 15
	}
}

func routeRotateFitness(route Route) float64 {
	result := 0.0
	for i, r := range route {
		result += rotateFitness(r.rotate)*float64(i+1)
	}
	return result
}
