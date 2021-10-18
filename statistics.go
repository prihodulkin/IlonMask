package main

import (
	"fmt"
	"math"
	"time"
)

func printPopulationStatistics(population []Route) {
	xMax := -1.0
	xMin := float64(width + 1)
	xAvg := 0.0
	for i := 0; i < len(population); i++ {
		x := population[i][len(population[i])-1].x
		if x > xMax {
			xMax = x
		} else if x < xMin {
			xMin = x
		}
		xAvg += x
	}
	xAvg /= float64(len(population))
	println("xMin: ", xMin, " xMax: ", xMax, " xAvg: ", xAvg)
}

func printTimeStatistics(attemptCount int, inputFilePath string) {
	initAngles()
	input := readFromFile(inputFilePath)
	s := input.shuttleData
	ground := input.ground
	var minTime = math.MaxInt64
	var maxTime = math.MinInt64
	var avgTime float64
	var minIter = math.MaxInt64
	var maxIter = math.MinInt64
	var avgIter float64

	for i := 0; i < attemptCount; i++ {
		result:= false
		findFlatSurface(ground)
		start := time.Now()
		population := generateRoutesPopulation(s, ground)
		iterCount := 1
		for ; !result; iterCount++ {
			population, result = generateNextPopulation(population, ground)
		}
		duration := int(time.Since(start).Milliseconds())
		avgTime += float64(duration)
		if duration < minTime {
			minTime = duration
		}
		if duration > maxTime {
			maxTime = duration
		}
		avgIter += float64(iterCount)
		if iterCount > maxIter {
			maxIter = iterCount
		}
		if iterCount < minIter {
			minIter = iterCount
		}
	}
	avgIter /= float64(attemptCount)
	avgTime /= float64(attemptCount)
	fmt.Printf("File: %s, attempts count: %d \n", inputFilePath, attemptCount)
	fmt.Printf("First population's count: %d, mutation probability: %f\n", populationSize, mutationProbability)
	fmt.Printf("Min iteration count: %d, Max iteration count: %d, Avg iteration count: %f\n", minIter, maxIter, avgIter)
	fmt.Printf("Min time: %d ms, Max time: %d ms, Avg time: %f ms", minTime, maxTime, avgTime)
}
