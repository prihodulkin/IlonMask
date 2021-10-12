package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	return ShuttleState{x, y, hSpeed, vSpeed, fuel, rotate, power}
}

func readFromFile(filePath string) InputData{
	fileHandle, _ := os.Open(filePath)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	fileScanner.Scan()
	var pointsN, _ = strconv.ParseInt(fileScanner.Text(), 10, 32)
	ground := make(Ground, pointsN-1)
	fileScanner.Scan()
	sarr := strings.Split(fileScanner.Text(), " ")
	var x, y int64
	x, _ = strconv.ParseInt(sarr[0], 10, 32)
	y, _ = strconv.ParseInt(sarr[1], 10, 32)
	for i := 1; i < int(pointsN); i++ {
		var x1, y1 int64
		fileScanner.Scan()
		sarr = strings.Split(fileScanner.Text(), " ")
		x1, _ = strconv.ParseInt(sarr[0], 10, 32)
		y1, _ = strconv.ParseInt(sarr[1], 10, 32)
		ground[i-1] = Surface{int(x), int(y), int(x1), int(y1)}
		x = x1
		y = y1
	}
	fileScanner.Scan()
	sarr = strings.Split(fileScanner.Text(), " ")
	var x0, _ = strconv.ParseFloat(sarr[0], 64)
	var y0, _ = strconv.ParseFloat(sarr[1], 64)
	var hSpeed, _ = strconv.ParseFloat(sarr[2], 64)
	var vSpeed, _ = strconv.ParseFloat(sarr[3], 64)
	var fuel, _ = strconv.ParseFloat(sarr[4], 64)
	var rotate, _ = strconv.ParseInt(sarr[5], 10,64)
	var power , _ = strconv.ParseInt(sarr[6], 10,64)
	state := ShuttleState{x0, y0, hSpeed, vSpeed, fuel, int(rotate), int(power)}
	return InputData{ground, state}
}

var input1 = InputData{
	Ground{
		{0, 100, 1000, 500},
		{1000, 500, 1500, 1500},
		{1500, 1500, 3000, 1000},
		{3000, 1000, 4000, 150},
		{4000, 150, 5500, 150},
		{5500, 150, 6999, 800},
	},
	ShuttleState{2500, 2700, 0, 0, 550, 0, 0}}
