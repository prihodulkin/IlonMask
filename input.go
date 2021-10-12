package main

import "fmt"

func readGround() Ground {
	var pointsN int
	fmt.Scan(&pointsN)
	result := make(Ground, pointsN-1)
	var x, y int
	fmt.Scan(&x, &y)
	for i := 1; i < pointsN; i++ {
		var x1, y1 int
		fmt.Scan(&x1, &y1)
		result[i-1] = Surface{x, y, x1, y1}
		x = x1
		y = y1
	}
	return result
}

func readShuttleState() ShuttleState {
	var x float64
	var y float64
	var hSpeed float64
	var vSpeed float64
	var fuel float64
	var rotate int
	var power int
	fmt.Scan(&x, &y, &hSpeed, &vSpeed, &fuel, &rotate, &power)
	return ShuttleState{x, y, hSpeed, vSpeed, fuel, rotate, power,0}
}

