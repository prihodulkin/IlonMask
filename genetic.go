package main

import (
	"math"
	"math/rand"
)

const populationCount = 100
const dTime float64 = 1
const routeCapacity = 50
const mutationProbability = 00.1

type Route []ShuttleState

func generateRoute(s ShuttleState, ground Ground, routeInd int) Route {
	s.routeInd = routeInd
	result := make(Route, 1, routeCapacity)
	result[0] = s
	s = move(s, dTime)
	result = append(result, s)
	i := 1
	moveResult := isLandedOrCrashed(ground, s.x, s.y)
	for !moveResult && s.x > 0 && s.x <= width && s.y <= height {
		result[i].rotate = generateRandomRotate(result[i-1].rotate)
		power := generateRandomPower(result[i-1].power)
		if power <= int(s.fuel) {
			result[i].power = power
		} else {
			result[i].power = int(s.fuel)
		}
		s = move(result[i], dTime)
		moveResult = isLandedOrCrashed(ground, s.x, s.y)
		result = append(result, s)
		i++
	}
	return result
}

func generateRoutesPopulation(s ShuttleState, ground Ground) []Route {
	population := make([]Route, populationCount)
	for i := 0; i < populationCount; i++ {
		population[i] = generateRoute(s, ground, i)
	}
	return population
}

func generateNextPopulation(population []Route, ground []Surface) ([]Route, bool) {
	By(FitnessCmp).Sort(population)
	bestLastState:=population[0][len(population[0])-1]
	if isResultWithoutRotate(bestLastState) {
		return population, true
	}
	c := math.Ceil((-1 + math.Sqrt(1+8*populationCount)) / 2)
	parentsCount := int(c)
	result := make([]Route, populationCount)
	for i := 0; i < parentsCount; i++ {
		result[i] = population[i]
	}
	newLineInd := parentsCount
	for i := 0; i < parentsCount; i++ {
		newColInd := 0
		for j := i + 1; j < parentsCount; j++ {
			p := rand.Float64()
			ind := newLineInd + newColInd
			if ind > len(result)-1 {
				return result, false
			}
			child := crossByPowerAndRotation(population[i], population[j], p, ground)
			//if child[len(child)-1].x==0{
			//	print()
			//}
			result[newLineInd+newColInd] = child
			newColInd++
		}
		newLineInd += parentsCount - i - 1
	}
	return result, false
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func crossByPowerAndRotation(first Route, second Route, p float64, ground Ground) Route {
	l := Min(len(first), len(second))
	resultLen := Max(len(first), len(second))
	result := make(Route, resultLen)
	result[0] = first[0]
	result[1] = first[1]
	c := int(math.Round(1 / mutationProbability))
	for i := 1; i < l-1; i++ {
		if rand.Int()%c == 0 {
			result[i].rotate = generateRandomRotate(result[i-1].rotate)
			result[i].power = generateRandomPower(result[i-1].power)
		} else {
			dPower := int(math.Round(float64(first[i].power-first[i-1].power)*p +
				float64(second[i].power-second[i-1].power)*(1-p)))
			dRotate := int(math.Round(float64(first[i].rotate-first[i-1].rotate)*p +
				float64(second[i].rotate-second[i-1].rotate)*(1-p)))
			result[i].SetRotate(result[i-1].rotate + dRotate)
			result[i].SetPower(result[i-1].rotate + dPower)
		}
		result[i+1] = move(result[i], dTime)
		moveResult := isLandedOrCrashed(ground, result[i+1].x, result[i+1].y)
		if moveResult {
			result = result[:i+2]
			return result
		}
	}
	if len(first) > l {
		result = fillTail(first, result, l, c, ground)
	} else if len(second) > l {
		result = fillTail(second, result, l, c, ground)
	}
	last := result[len(result)-1]
	if !isLandedOrCrashed(ground, last.x, last.y) {
		result = append(result, generateRoute(last, ground, last.routeInd)...)
	}
	//if result[resultLen-1].x==0{
	//	print()
	//}
	return result
}

func fillTail(source Route, result Route, l int, c int, ground Ground) Route {
	for i := l - 1; i < len(source)-1; i++ {
		if rand.Int()%c == 0 {
			result[i].rotate = generateRandomRotate(result[i-1].rotate)
			result[i].power = generateRandomPower(result[i-1].power)
		} else {
			dPower := source[i].power - source[i-1].power
			dRotate := source[i].rotate - source[i-1].rotate
			result[i].SetRotate(result[i-1].rotate + dRotate)
			result[i].SetPower(result[i-1].power + dPower)
		}
		result[i+1] = move(result[i], dTime)
		moveResult := isLandedOrCrashed(ground, result[i+1].x, result[i+1].y)
		if moveResult {
			result = result[:i+2]
			return result
		}
	}
	return result
}
