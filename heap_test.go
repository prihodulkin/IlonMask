package main

import (
	"container/heap"
	"testing"
)

func TestLandingPositionHeap(t *testing.T) {
	xFlatMin = 500
	xFlatMax = 600
	h := &LandingPositionHeap{{100, 200, 0, 0, 100, 1,  0, 0},
		{2000, 200, 0, 0, 100, 1, 233, 0},
		{550, 200, 0, 0, 100, 1, 233, 0},
		{499, 200, 0, 0, 100, 1, 233, 0},
		{1000, 200, 0, 0, 100, 1, 233, 0}}
	heap.Init(h)
	if !almostEqual(heap.Pop(h).(ShuttleState).x, 2000) {
		t.Error()
	}
	if !almostEqual(heap.Pop(h).(ShuttleState).x,1000){
		t.Error()
	}
	if !almostEqual(heap.Pop(h).(ShuttleState).x,100){
		t.Error()
	}
	if !almostEqual(heap.Pop(h).(ShuttleState).x,499){
		t.Error()
	}
	if !almostEqual(heap.Pop(h).(ShuttleState).x,550){
		t.Error()
	}
}


func TestVSpeedPositionHeap(t *testing.T) {
	h := &VSpeedHeap{{100, 200, 0, -45, 100, 1, 233, 0},
		{2000, 200, 0, 10, 100, 1, 233, 0},
		{550, 200, 0, 41, 100, 1, 233, 0},
		{499, 200, 0, -60, 100, 1, 233, 0},
		{1000, 200, 0, 90, 100, 1, 233, 0}}
	heap.Init(h)
	if !almostEqual(heap.Pop(h).(ShuttleState).vSpeed, 90) {
		t.Error()
	}
	if !almostEqual(heap.Pop(h).(ShuttleState).vSpeed,-60){
		t.Error()
	}
	if !almostEqual(heap.Pop(h).(ShuttleState).vSpeed,-45){
		t.Error()
	}
	if !almostEqual(heap.Pop(h).(ShuttleState).vSpeed,41){
		t.Error()
	}
	if !almostEqual(heap.Pop(h).(ShuttleState).vSpeed,10){
		t.Error()
	}
}