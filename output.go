package main

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"log"
	"math"
	"os"
)

func (state ShuttleState) printShuttleLanding() {
	fmt.Printf("x: %f, vSpeed: %f   hSpeed: %f , rotate: %d , fitness:%f \n ",
		state.x, state.vSpeed, state.hSpeed, state.rotate,  fitnessState(state))
}

func printGround(ground Ground, canvas *svg.SVG) {
	blackStroke := "stroke=\"black\""
	for _, s := range ground {
		canvas.Line(s.x1, height-s.y1, s.x2, height-s.y2, blackStroke)
	}
}

func printRoute(route Route, canvas *svg.SVG) {
	redStroke := "stroke=\"red\""
	if len(route) > 0 {
		x := int(math.Round(route[0].X()))
		y := int(math.Round(height - route[0].Y()))
		canvas.Circle(x, y, 1, redStroke)
		for i := 1; i < len(route); i++ {
			x1 := int(math.Round(route[i].X()))
			y1 := int(math.Round(height - route[i].Y()))
			canvas.Line(x, y, x1, y1, redStroke)
			canvas.Circle(x1, y1, 1, redStroke)
			x = x1
			y = y1
		}
	}
}

func printRoutesInSVG(ground Ground, routes []Route, filePath string) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	canvas := svg.New(f)
	canvas.Start(width, height)
	printGround(ground, canvas)
	for _, route := range routes {
		printRoute(route, canvas)
	}
	canvas.End()
}

func printRouteInSVG(ground Ground, route Route, filePath string) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	canvas := svg.New(f)
	canvas.Start(width, height)
	printGround(ground, canvas)
	printRoute(route, canvas)
	canvas.End()
}
