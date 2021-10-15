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

type ShuttleState struct {
	x        float64
	y        float64
	hSpeed   float64
	vSpeed   float64
	fuel     float64
	rotate   int
	power    int
}

type InputData struct {
	ground Ground
	shuttleState ShuttleState
}

func (state ShuttleState) SetPower(power int) {
	if power < 0 {
		state.power = 0
	} else if power > 4 {
		state.power = 4
	} else {
		state.power = power
	}
}

func (state ShuttleState) SetRotate(rotate int) {
	if rotate < -90 {
		state.rotate = -90
	} else if rotate > 90 {
		state.rotate = 90
	} else {
		state.rotate = rotate
	}
}
func (state ShuttleState) X() float64 {
	return state.x
}

func (state ShuttleState) Y() float64 {
	return state.y
}
