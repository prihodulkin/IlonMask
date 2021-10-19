package main

type Surface struct {
	x1 int
	y1 int
	x2 int
	y2 int
	A int
	B int
	C int
}

type Ground []Surface

type ShuttleData struct {
	x      float64
	y      float64
	hSpeed float64
	vSpeed float64
	fuel   float64
	rotate int
	power  int
}

type InputData struct {
	ground      Ground
	shuttleData ShuttleData
}


