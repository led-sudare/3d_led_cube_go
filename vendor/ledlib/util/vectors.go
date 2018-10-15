package util

import "math/rand"

func GetSphereFace() Point {
	for {
		pos := NewPoint((rand.Float64()*2)-1,
			(rand.Float64()*2)-1,
			(rand.Float64()*2)-1)
		d2 := pos.X()*pos.X() + pos.Y()*pos.Y() + pos.Z()*pos.Z()
		if 0.1 < d2 && d2 < 1.0 {
			d := 0.5
			return NewPoint(pos.X()*d, pos.Y()*d, pos.Z()*d)
		}
	}
}

func CanShow(p Point, xmax, ymax, zmax float64) bool {
	return 0 <= p.X() && p.X() < xmax &&
		0 <= p.Y() && p.Y() < ymax &&
		0 <= p.Z() && p.Z() < zmax

}
