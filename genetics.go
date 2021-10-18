package main

import (
	"math"
	"math/rand"
)

const populationSize = 100
const parentsCount = 100
const chromosomeSize = 200
const dTime float64 = 1
const deltaTime = 1
const routeCapacity = 50
const mutationProbability = 0.1

type Gene struct {
	power  int
	rotate int
}

type Chromosome []Gene

type Population []Chromosome

type ChromosomeData struct {
	Chromosome   Chromosome
	Path         Path
	FitnessData  FitnessData
	LandingPoint Point
}

type PopulationData []ChromosomeData

var population Population
var populationData PopulationData

func GenerateChromosome(data *ShuttleData) Chromosome {
	result := make(Chromosome, chromosomeSize)
	result[0].rotate=data.rotate
	result[0].power=data.power
	for i := 1; i < chromosomeSize; i++ {
		result[i].rotate=generateRandomRotate(result[i-1].rotate)
		result[i].power=generateRandomPower(result[i-1].power)
	}
	return result
}

func GeneratePopulation(data *ShuttleData)  {
	population = make(Population, populationSize)
	populationData=make(PopulationData, populationSize)
	for i := 0; i < populationSize; i++ {
		population[i] = GenerateChromosome(data)
	}
}

func GenerateNextPopulation() bool {
	//By(FitnessCmp).Sort(population)
	//bestLastState := population[0][len(population[0])-1]
	//if isResult(bestLastState, population[0][len(population[0])-2].rotate) {
	//	return population, true
	//}
	////c :=int( math.Ceil((-1 + math.Sqrt(1+8*populationSize)) / 2))
	//result := make([]Route, 0, populationSize)
	//for i := 0; i < parentsCount; i++ {
	//	result = append(result, population[i])
	//}
	//for i := 0; i < populationSize-parentsCount; i++ {
	//	i1 := int(rand.Float64() * rand.Float64() * rand.Float64() * populationSize / 5)
	//	p := rand.Float64()
	//	child := crossByPowerAndRotation(population[i1%parentsCount], population[i%parentsCount], p, ground)
	//	result = append(result, child)
	//}
	//
	//return result, false
	return false
}



//скрещивание с помощью Continuous Genetic Algorithm
func CrossByPowerAndRotation(first Route, second Route, p float64, ground Ground) Route {
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


