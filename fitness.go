package main

import (
	"fmt"
	"math"
)

var xFlatMin float64
var xFlatMax float64

type FitnessData struct {
	hSpeed       float64
	vSpeed       float64
	lastRotation int
	predRotation int
	x            float64
}

func Fitness(data FitnessData) float64 {
	return math.Pow(landingFitness(data.x, xFlatMin, xFlatMax), 3) +
		math.Pow(vSpeedFitness(data.vSpeed), 2) +
		math.Pow(hSpeedFitness(data.hSpeed), 2) +
		math.Pow(rotateFitness(data.lastRotation), 1) +
		math.Pow(rotateFitness(data.predRotation), 1)
}

func (data FitnessData) String() string {
	return fmt.Sprintf("x: %f, vSpeed: %f   hSpeed: %f , "+
		"pred dRotate: %d, last dRotate: %d , fitness:%f \n ",
		data.x, data.vSpeed, data.hSpeed, data.predRotation, data.lastRotation, Fitness(data))
}

func IsSolution(data FitnessData) bool {
	return landingFitness(data.x, xFlatMin, xFlatMax)+
		vSpeedFitness(data.vSpeed)+hSpeedFitness(data.hSpeed)+
		rotateFitness(data.lastRotation)+rotateFitness(data.predRotation) == 0
}

func fitnessState(state ShuttleData) float64 {
	return math.Pow(landingFitness(state.x, xFlatMin, xFlatMax), 3) +
		math.Pow(vSpeedFitness(state.vSpeed), 1) +
		math.Pow(hSpeedFitness(state.hSpeed), 1) +
		rotateFitness(state.rotate)
}

func isResultWithoutRotate(state ShuttleData) bool {
	return landingFitness(state.x, xFlatMin, xFlatMax)+
		vSpeedFitness(state.vSpeed)+
		hSpeedFitness(state.hSpeed) == 0
}

func isResult(state ShuttleData, prevRotate int) bool {
	return landingFitness(state.x, xFlatMin, xFlatMax)+
		vSpeedFitness(state.vSpeed)+
		hSpeedFitness(state.hSpeed)+
		rotateFitness(state.rotate)+
		rotateFitness(prevRotate) == 0
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
