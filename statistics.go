package main

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
	println("xMin: ", xMin," xMax: ",xMax, " xAvg: ", xAvg)
}
