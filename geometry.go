package main

func line(x1 int, x2 int, y1 int, y2 int) (int, int, int) {
	A := y1 - y2
	B := x2 - x1
	C := x1*y2 - x2*y1
	return A, B, C
}
