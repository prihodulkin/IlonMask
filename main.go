package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	initAngles()
	ground := readGround()
	findFlatSurface(ground)
	s := readShuttleState()
	population := generateRoutesPopulation(s, ground)
	printRoutesInSVG(ground, population, "routes/routes0.svg")
	result := false
	for i := 1; !result&&i<1000; i++ {
		population, result = generateNextPopulation(population, ground)
		fmt.Printf("%d ",i)
		bestLastState:=population[0][len(population[0])-1]
		bestLastState.printShuttleLanding()
		//fileName := "routes/routes" + strconv.Itoa(i) + ".svg"
		//printRoutesInSVG(ground, population, fileName)
	}
}
