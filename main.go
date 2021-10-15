package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func printInitialDistributions(filePath string){
	input:=readFromFile(filePath)
	ground:=input.ground
	s := input.shuttleState
	for i:=1;i<10;i++{
		population := generateRoutesPopulation(s, ground)
		fileName := "routes/routesDistribution" + strconv.Itoa(i) + ".svg"
		printRoutesInSVG(ground, population, fileName)
	}
}

func printLastStates(filePath string){
	input:=readFromFile(filePath)
	ground:=input.ground
	findFlatSurface(ground)
	s := input.shuttleState
	population := generateRoutesPopulation(s, ground)
	printRoutesInSVG(ground, population, "routes/routes0.svg")
	result := false
	for i := 1; !result; i++ {
		population, result = generateNextPopulation(population, ground)
		fmt.Printf("%d ",i)
		for i:=0;i<parentsCount;i++{
			bestLastState:=population[i][len(population[i])-1]
			bestLastState.printShuttleLanding()
		}

		//fileName := "routes/routes" + strconv.Itoa(i) + ".svg"
		//printRoutesInSVG(ground, population, fileName)
	}
}

func printSolution(filePath string){
	input:=readFromFile(filePath)
	ground:=input.ground
	findFlatSurface(ground)
	s := input.shuttleState
	population := generateRoutesPopulation(s, ground)
	printRoutesInSVG(ground, population, "routes/routes0.svg")
	result := false
	for i := 1; !result; i++ {
		population, result = generateNextPopulation(population, ground)
		fmt.Printf("%d ",i)
		bestLastState:=population[0][len(population[0])-1]
		bestLastState.printShuttleLanding()
		fileName := "routes/routes" + strconv.Itoa(i) + ".svg"
		printRoutesInSVG(ground, population, fileName)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	initAngles()
	//printSolution("input/input2.txt")
	//printInitialDistributions("input/input2.txt")
	printLastStates("input/input2.txt")
}
