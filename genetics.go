package main

import (
	"math"
	"math/rand"
)

//все параметры подобраны методом научного тыка
const populationSize = 200
const childrenCount = 100
const chromosomeSize = 300
const dTime = 1
const mutationProbability = 0.05

type Gene struct {
	dPower  int
	dRotate int
}

type Chromosome []Gene


type ChromosomeData struct {
	Chromosome   Chromosome
	Path         Path
	FitnessData  FitnessData
	LandingPoint Point
}

type PopulationData []ChromosomeData

var populationData PopulationData

func GenerateRandomRotate() int {
	return -15 + rand.Intn(31)
}

func GenerateRandomPower() int {
	return -1 + rand.Intn(3)
}

func GenerateChromosome(data *ShuttleData, ind int)  {
	populationData[ind].Chromosome = make(Chromosome, chromosomeSize)
	result:=populationData[ind].Chromosome
	result[0].dRotate = data.rotate
	result[0].dPower = data.power
	for i := 1; i < chromosomeSize; i++ {
		result[i].dRotate = GenerateRandomRotate()
		result[i].dPower = GenerateRandomPower()
	}
}

func GeneratePopulation(data *ShuttleData) {
	populationData = make(PopulationData, populationSize)
	for i := 0; i < populationSize; i++ {
		 GenerateChromosome(data,i)
	}
}

func GenerateParentIndex() int {
	return int((rand.Float64() * rand.Float64()) * populationSize)
}

func GenerateNextPopulation() bool {
	By.Sort(FitnessCmp, populationData)
	if IsSolution(populationData[0].FitnessData) {
		return true
	}
	ApplyGenerationMethod1()
	return false
}

func ApplyGenerationMethod1() {
	j := populationSize - 1
	for i := 0; i < childrenCount; i++ {
		i1 := GenerateParentIndex()
		i2 := GenerateParentIndex()
		Cross(populationData[i1].Chromosome, populationData[i2].Chromosome, j)
		j--
	}
}

func ApplyGenerationMethod2() {
	parentsCount := int((-1 + math.Sqrt(1+8*float64(populationSize))) / 2)
	k := parentsCount
	for i := 0; i < parentsCount; i++ {
		for j := i + 1; j < parentsCount; j++ {
			Cross(populationData[i].Chromosome, populationData[j].Chromosome,k)
			k++
			if k==populationSize{
				break
			}
		}
	}
}

//чтоб не перевыделять память перепичываю поверх существующих плохих хромосом
func Cross(first Chromosome, second Chromosome, ind int) {
	res := populationData[ind].Chromosome
	p := rand.Float64()
	for i := 0; i < chromosomeSize; i++ {
		if rand.Float64() < mutationProbability {
			res[i].dPower = GenerateRandomPower()
			res[i].dRotate = GenerateRandomRotate()
		} else{
			res[i].dPower = int(math.Round(p*float64(first[i].dPower) + (1-p)*float64(second[i].dPower)))
			res[i].dRotate = int(math.Round(p*float64(first[i].dRotate) + (1-p)*float64(second[i].dRotate)))
		}
	}
}
