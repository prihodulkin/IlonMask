package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

//на каждой итерации выводит в консоль  данные о лучшем пути и отриовывает все пути в файле paths/path{i}.svg
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

//выводит в консоль данные о лучшем пути на каждой итерации и рисует результат в paths/solution.svg
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
	for i:=1;i<=5;i++{
		filePath:="input/input"+strconv.Itoa(i)+".txt"
		PrintTimeStatistics(10, filePath)
		fmt.Println("\n\n ")
	}
	//при желании можно посмотреть по шагам на сервере, вывести путь решения в svg-файл
	//или вывести в svg-файлы все итерации
	//input := readFromFile("input/input5.txt")
	//httpServeSolution(input)
	//findSolution(input)
	//printSolution(input)
}
