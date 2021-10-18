package main

type Point struct {
	X float64
	Y float64
}

type Path [] Point

func Line(p1 Point, p2 Point) (float64, float64, float64) {
	x1 := p1.X
	x2 := p2.X
	y1 := p1.Y
	y2 := p2.Y
	A := y2 - y1
	B := x1 - x2
	C := x2*y1 - x1*y2
	return A, B, C
}

func (l *Surface) IntersectPath(p1 Point, p2 Point) (Point, bool) {
	A1 := float64(l.A)
	B1 := float64(l.B)
	C1 := float64(l.C)
	A2, B2, C2 := Line(p1, p2)
	c := A1*B2 - A2*B1
	if c == 0 {
		return Point{}, false
	}
	a := B1*C2 - B2*C1
	b := A2*C1 - A1*C2
	x :=a / c
	y:=b/ c
	if float64(l.x1) <= x && x <= float64(l.x2) &&(p1.Y<=y&&y<=p2.Y||p2.Y<=y&&y<=p1.Y){
		return Point{x,y }, true
	}
	return Point{}, false
}

func IsLandedOrCrashed(ground Ground, p1 Point, p2 Point) (Point, bool) {
	for _, s := range ground {
		x1 := float64(s.x1)
		x2 := float64(s.x2)
		if p1.X >= x1 && p1.X <= x2 || p2.X >= x1 && p2.X <= x2 {
			p,res:= s.IntersectPath(p1, p2)
			if res{
				return p,res
			}
		}
	}
	return Point{}, false
}
