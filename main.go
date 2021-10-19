package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

//func printInitialDistributions(filePath string) {
//	input := readFromFile(filePath)
//	ground := input.ground
//	s := input.shuttleData
//	for i := 1; i < 10; i++ {
//		population := generateRoutesPopulation(s, ground)
//		fileName := "populationData/routesDistribution" + strconv.Itoa(i) + ".svg"
//		printRoutesInSVG(ground, population, fileName)
//	}
//}
//
//func printLastStates(filePath string){
//	input:=readFromFile(filePath)
//	ground:=input.ground
//	findFlatSurface(ground)
//	s := input.shuttleData
//	population := generateRoutesPopulation(s, ground)
//	printRoutesInSVG(ground, population, "populationData/routes0.svg")
//	result := false
//	for i := 1; !result; i++ {
//		population, result = generateNextPopulation(population, ground)
//		fmt.Printf("%d ",i)
//		for i:=0;i<childrenCount;i++{
//			bestLastState:=population[i][len(population[i])-1]
//			fmt.Printf(bestLastState.String())
//		}
//		//fileName := "populationData/populationData" + strconv.Itoa(i) + ".svg"
//		//printRoutesInSVG(ground, population, fileName)
//	}
//}



func printSolution(input InputData) {
	ground := input.ground
	findFlatSurface(ground)
	shuttleData := &input.shuttleData
	GeneratePopulation(shuttleData)
	ApplyPopulation(shuttleData, ground)
	result := false
	PrintPathsInSVG(ground,"paths/paths0.svg")
	logger := log.New(os.Stderr, "", 0)
	for i := 1; !result; i++ {
		result = GenerateNextPopulation()
		ApplyPopulation(shuttleData, ground)
		bestLastState:=populationData[0]
		fileName := "paths/paths" + strconv.Itoa(i) + ".svg"
		PrintPathsInSVG(ground,fileName)
		logger.Printf("%d %s", i, bestLastState.FitnessData.String())
	}
}


func findSolution(input InputData) {
	ground := input.ground
	findFlatSurface(ground)
	shuttleData := &input.shuttleData
	GeneratePopulation(shuttleData)
	ApplyPopulation(shuttleData, ground)
	result := false
	PrintPathsInSVG(ground,"paths/paths0.svg")
	logger := log.New(os.Stderr, "", 0)
	for i := 1; !result; i++ {
		result = GenerateNextPopulation()
		ApplyPopulation(shuttleData, ground)
		bestLastState:=populationData[0]
		logger.Printf("%d %s", i, bestLastState.FitnessData.String())
	}
	PrintPathInSVG(ground,populationData[0].Path,"paths/solution.svg")
}
//поднимает http сервер localhost:3000, обновляя страницу запускаем новую итерацию
func httpServeSolution(input InputData) {
	ground := input.ground
	findFlatSurface(ground)
	s := input.shuttleData
	shuttleData := &input.shuttleData
	GeneratePopulation(shuttleData)
	ApplyPopulation(shuttleData, ground)
	GeneratePopulation(&s)
	ApplyPopulation(&s, ground)
	http.HandleFunc("/favicon.ico", func(writer http.ResponseWriter, request *http.Request) {})
	i := 0
	logger := log.New(os.Stderr, "", 0)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "image/svg+xml")
		WritePathsSVG(ground, writer)
		GenerateNextPopulation()
		ApplyPopulation(&s, ground)
		bestLastState := populationData[0]
		logger.Printf("%d %s", i, bestLastState.FitnessData.String())
		i++
	})
	log.Println(http.ListenAndServe("localhost:3000", nil))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	initAngles()
	PrintTimeStatistics(100, "input/input5.txt")
	//input := readFromFile("input/input5.txt")
	//findSolution(input)
	//printSolution(input)
}
