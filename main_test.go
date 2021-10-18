package main

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-5

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

//func TestLandedOrCrashedOnSurface(t *testing.T) {
//	var s = Surface{0, 0, 5, 0}
//
//	if !IsUnderground(s, 3, 0) {
//		t.Error()
//	}
//
//	s = Surface{0, 0, 4, 4}
//
//	if !IsUnderground(s, 2, 2) {
//		t.Error()
//	}
//
//	if IsUnderground(s, 2, 3) {
//		t.Error()
//	}
//
//	s = Surface{0, 0, 3, 2}
//
//	if !IsUnderground(s, 1, 2/3) {
//		t.Error()
//	}
//}

func testMovingCase(t *testing.T, res ShuttleData, expected ShuttleData) {
	vSpeedExpected := expected.vSpeed
	if !almostEqual(res.vSpeed, vSpeedExpected) {
		t.Error("vSpeed should be equals ", vSpeedExpected, " but is ", res.vSpeed)
	}

	hSpeedExpected := expected.hSpeed
	if !almostEqual(res.hSpeed, hSpeedExpected) {
		t.Error("hSpeed should be equals ", hSpeedExpected, " but is ", res.hSpeed)
	}

	xExpected := expected.x
	if !almostEqual(res.x, xExpected) {
		t.Error("x should be equals ", xExpected, " but is ", res.x)
	}

	yExpected := expected.y
	if !almostEqual(res.y, yExpected) {
		t.Error("y should be equals ", yExpected, " but is ", res.y)
	}

	fuelExpected := expected.fuel
	if !almostEqual(res.fuel, fuelExpected) {
		t.Error("fuel should be equals ", fuelExpected, " but is ", res.fuel)
	}
}

func TestMoving(t *testing.T) {
	initAngles()
	s := ShuttleData{100, 100, 0, 0, 100, 0, 4}
	s1 := move(&s, 1)
	expected := ShuttleData{100, 100.1445, 0, 0.289, 96, 0, 4}
	testMovingCase(t, s1, expected)

	s = ShuttleData{200, 100, -5, 10, 100, 30, 3}
	s1 = move(&s, 4)
	expected = ShuttleData{168, 131.096609, -11, 5.548304, 88, 30, 3}
	testMovingCase(t, s1, expected)
}
