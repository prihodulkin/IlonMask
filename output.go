package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"

	"github.com/ajstarks/svgo"
)

func (state ShuttleData) String() string {
	return fmt.Sprintf("x: %f, vSpeed: %f   hSpeed: %f , dRotate: %d , fitness:%f \n ",
		state.x, state.vSpeed, state.hSpeed, state.rotate, fitnessState(state))
}

func PrintGround(ground Ground, canvas *svg.SVG) {
	blackStroke := `stroke="black"`
	for _, s := range ground {
		canvas.Line(s.x1, height-s.y1, s.x2, height-s.y2, blackStroke)
	}
}

func PrintPath(path Path, canvas *svg.SVG) {
	redStroke := "stroke=\"red\" title=\"path\""
	if len(path) > 0 {
		x := int(math.Round(path[0].X))
		y := int(math.Round(height - path[0].Y))
		canvas.Circle(x, y, 1, redStroke)
		for i := 1; i < len(path); i++ {
			x1 := int(math.Round(path[i].X))
			y1 := int(math.Round(height - path[i].Y))
			canvas.Line(x, y, x1, y1, redStroke)
			canvas.Circle(x1, y1, 1, redStroke)
			x = x1
			y = y1
		}
	}
}

func PrintPathsInSVG(ground Ground, filePath string) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	WritePathsSVG(ground, f)
}

func WritePathsSVG(ground Ground, w io.Writer) {
	canvas := svg.New(w)
	canvas.Start(width, height)
	PrintGround(ground, canvas)
	// из-за того что браузер очень тупит при отрисовке больших svg, отрисовываю только самые интересные
	for _, data := range populationData[:50] {
		PrintPath(data.Path, canvas)
	}
	canvas.End()
}

func PrintPathInSVG(ground Ground, path Path, filePath string) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	canvas := svg.New(f)
	canvas.Start(width, height)
	PrintGround(ground, canvas)
	PrintPath(path, canvas)
	canvas.End()
}

func PrintInitialDistributions(filePath string) {
	input := readFromFile(filePath)
	ground := input.ground
	shuttleData := &input.shuttleData
	GeneratePopulation(shuttleData)
	ApplyPopulation(shuttleData, ground)
	fileName := "paths/routesDistribution" + ".svg"
	PrintPathsInSVG(ground, fileName)
}
