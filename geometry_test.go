package main

import "testing"

func TestIntersection(t *testing.T) {
	p1 := Point{0.0, 0.0}
	p2 := Point{100, 100}
	A, B, C := Line(p1, p2)
	surface := Surface{int(p1.X),
		int(p1.Y), int(p2.X),
		int(p2.Y), int(A), int(B),
		int(C)}
	p3 := Point{50, 100}
	p4 := Point{50, 0}
	_, res := surface.IntersectPath(p3, p4)
	if !res {
		t.Error()
	}
	p3 = Point{50, 100}
	p4 = Point{60, 50}
	_, res = surface.IntersectPath(p3, p4)
	if !res {
		t.Error()
	}
}
