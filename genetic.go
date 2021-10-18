package main

import (
	"math"
	"math/rand"
)





type Route []ShuttleData

func generateRoute(s ShuttleData, ground Ground) Route {
	result := make(Route, 1, routeCapacity)
	result[0] = s
	s = move(&s, dTime)
	result = append(result, s)
	moveResult := isLandedOrCrashed(ground, s.x, s.y)
	//p:=rand.Int()
	//k:=rand.Intn(3)
	for i := 1; !moveResult && s.x > 0 && s.x <= width && s.y <= height; i++ {
		//if p%c == 0 {
		//	//чтоб был больше разброс
		//	if k==2{
		//		result[i].rotate = generateRandomRotateWithBounds(result[i-1].rotate,0,15)
		//		result[i].power = generateRandomPower(result[i-1].power)
		//	} else if k==1{
		//		result[i].rotate = generateRandomRotateWithBounds(result[i-1].rotate,-15,0)
		//		result[i].power = generateRandomPower(result[i-1].power)
		//	}else{
		//		result[i].rotate = generateRandomRotate(result[i-1].rotate)
		//		result[i].power = generateRandomPower(result[i-1].power)
		//	}
		//} else{
		//	result[i].rotate = result[i-1].rotate
		//	result[i].power = result[i-1].power
		//}
		result[i].rotate = generateRandomRotate(result[i-1].rotate)
		result[i].power = generateRandomPower(result[i-1].power)
		s = move(&result[i], dTime)
		moveResult = isLandedOrCrashed(ground, s.x, s.y)
		result = append(result, s)
	}
	return result
}


func generateRoutesPopulation(s ShuttleData, ground Ground) []Route {
	population := make([]Route, populationSize)
	for i := 0; i < populationSize; i++ {
		population[i] = generateRoute(s, ground)
	}
	return population
}

func generateNextPopulation(population []Route, ground []Surface) ([]Route, bool) {
	By(FitnessCmp).Sort(population)
	bestLastState := population[0][len(population[0])-1]
	if isResult(bestLastState,population[0][len(population[0])-2].rotate) {
		return population, true
	}
	//c :=int( math.Ceil((-1 + math.Sqrt(1+8*populationSize)) / 2))
	result := make([]Route, 0, populationSize)
	for i:=0;i<parentsCount;i++{
		result = append(result, population[i])
	}
	for i := 0; i < populationSize-parentsCount; i++ {
		i1:=int(rand.Float64()*rand.Float64()*rand.Float64()* populationSize /5)
		p := rand.Float64()
		child := crossByPowerAndRotation(population[i1%parentsCount], population[i%parentsCount], p, ground)
		result = append(result, child)
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

//скрещивание с помощью Continuous Genetic Algorithm
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
		result[i+1] = move(&result[i], dTime)
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
		result = append(result, generateRoute(last, ground)...)
	}
	return result
}

//дозаполнение результата, если хромосома одного из родителей была короче
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
		result[i+1] = move(&result[i], dTime)
		moveResult := isLandedOrCrashed(ground, result[i+1].x, result[i+1].y)
		if moveResult {
			result = result[:i+2]
			return result
		}
	}
	return result
}
