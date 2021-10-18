package main

import (
	"container/heap"
	"testing"
)

func TestLandingPositionHeap(t *testing.T) {
	xFlatMin = 500
	xFlatMax = 600
	h := &LandingPositionHeap{{100, 200, 0, 0, 100, 1,  0},
		{2000, 200, 0, 0, 100, 1, 233},
		{550, 200, 0, 0, 100, 1, 233},
		{499, 200, 0, 0, 100, 1, 233},
		{1000, 200, 0, 0, 100, 1, 233}}
	heap.Init(h)
	if !almostEqual(heap.Pop(h).(ShuttleData).x, 2000) {
		t.Error()
	}
	if !almostEqual(heap.Pop(h).(ShuttleData).x,1000){
		t.Error()
	}
	if !almostEqual(heap.Pop(h).(ShuttleData).x,100){
		t.Error()
	}
	if !almostEqual(heap.Pop(h).(ShuttleData).x,499){
		t.Error()
	}
	if !almostEqual(heap.Pop(h).(ShuttleData).x,550){
		t.Error()
	}
}


func TestVSpeedPositionHeap(t *testing.T) {
	h := &VSpeedHeap{{100, 200, 0, -45, 100, 1, 233},
		{2000, 200, 0, 10, 100, 1, 233},
		{550, 200, 0, 41, 100, 1, 233},
		{499, 200, 0, -60, 100, 1, 233},
		{1000, 200, 0, 90, 100, 1, 233}}
	heap.Init(h)
	if !almostEqual(heap.Pop(h).(ShuttleData).vSpeed, 90) {
		t.Error()
	}
	if !almostEqual(heap.Pop(h).(ShuttleData).vSpeed,-60){
		t.Error()
	}
	if !almostEqual(heap.Pop(h).(ShuttleData).vSpeed,-45){
		t.Error()
	}
	if !almostEqual(heap.Pop(h).(ShuttleData).vSpeed,41){
		t.Error()
	}
	if !almostEqual(heap.Pop(h).(ShuttleData).vSpeed,10){
		t.Error()
	}
}