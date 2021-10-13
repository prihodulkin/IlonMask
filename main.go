package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	initAngles()
	printTimeStatistics(100,"input/input1.txt")
	//input:=readFromFile("input/input1.txt")
	//ground:=input.ground
	//findFlatSurface(ground)
	//s := input.shuttleState
	//population := generateRoutesPopulation(s, ground)
	////printRoutesInSVG(ground, population, "routes/routes0.svg")
	//result := false
	//for i := 1; !result; i++ {
	//	population, result = generateNextPopulation(population, ground)
	//	fmt.Printf("%d ",i)
	//	bestLastState:=population[0][len(population[0])-1]
	//	bestLastState.printShuttleLanding()
	//	//fileName := "routes/routes" + strconv.Itoa(i) + ".svg"
	//	//printRoutesInSVG(ground, population, fileName)
	//}
}
