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

func printInitialDistributions(filePath string) {
	input := readFromFile(filePath)
	ground := input.ground
	s := input.shuttleState
	for i := 1; i < 10; i++ {
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
			fmt.Printf(bestLastState.String())
		}
		//fileName := "routes/routes" + strconv.Itoa(i) + ".svg"
		//printRoutesInSVG(ground, population, fileName)
	}
}


func printSolution(input InputData) Route {
	ground := input.ground
	findFlatSurface(ground)
	s := input.shuttleState
	population := generateRoutesPopulation(s, ground)
	printRoutesInSVG(ground, population, "routes/routes0.svg")
	result := false
	// логгер - удобная обёртка над io.Writer для вывода
	// её полезно использовать, чтобы можно было легко перенаправить в любой io.Writer
	// *os.File (которыми являются потоки вывода os.Stdout и os.Stderr), *bytes.Buffer, http.ResponseWriter - реализуют io.Writer
	// io.Writer вообще очень важный интерфейс
	logger := log.New(os.Stderr, "", 0)
	for i := 1; !result; i++ {
		population, result = generateNextPopulation(population, ground)
		fmt.Printf("%d ",i)
		bestLastState:=population[0][len(population[0])-1]
		fmt.Printf(bestLastState.String())
		fileName := "routes/routes" + strconv.Itoa(i) + ".svg"
		printRoutesInSVG(ground, population, fileName)
		// Вместо того чтобы делать метод печати у состояния, лучше сделать так, чтобы он реализовывал интерфейс fmt.Stringer
		// т.е. имел метод String() string, тогда можно делать так
		logger.Printf("%d %s", i, bestLastState)
		//fileName := "routes/routes" + strconv.Itoa(i) + ".svg"
		//printRoutesInSVG(ground, population, fileName)

	}
	return population[0]
}

// поднимает http сервер localhost:3000, обновляя страницу запускаем новую итерацию
func httpServeSolution(input InputData) {
	ground := input.ground
	findFlatSurface(ground)
	s := input.shuttleState
	population := generateRoutesPopulation(s, ground)
	// Игнорируем запрос фавикона
	http.HandleFunc("/favicon.ico", func(writer http.ResponseWriter, request *http.Request) {})
	i := 0
	logger := log.New(os.Stderr, "", 0)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		population, _ = generateNextPopulation(population, ground)
		bestLastState := population[0][len(population[0])-1]
		logger.Printf("%d %s", i, bestLastState)
		i++
		writer.Header().Set("Content-Type", "image/svg+xml")
		WriteRoutesSVG(ground, population[:100], writer)
	})
	log.Println(http.ListenAndServe("localhost:3000", nil))
	//printRoutesInSVG(ground, population, "routes/routes0.svg")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	initAngles()

	//printSolution("input/input2.txt")
	//printInitialDistributions("input/input2.txt")
	printLastStates("input/input2.txt")

	input := readFromFile("input/input2.txt")
	//printRouteInSVG(input.ground, printSolution(input), "result.svg")
	//printTimeStatistics(100,"input/input1.txt")
	httpServeSolution(input)
}
