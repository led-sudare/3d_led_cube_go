package util

import "math"

type Point interface {
	X() float64
	Y() float64
	Z() float64
	Len() float64
	Add(point Point)
}

type pointImpl struct {
	x, y, z float64
}

func NewPoint(x, y, z float64) Point {
	return &pointImpl{x, y, z}
}

func (p *pointImpl) X() float64 {
	return p.x
}

func (p *pointImpl) Y() float64 {
	return p.y
}

func (p *pointImpl) Z() float64 {
	return p.z
}

func (p *pointImpl) Len() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

func (p *pointImpl) Add(point Point) {
	p.x += point.X()
	p.y += point.Y()
	p.z += point.Z()
}

type PointC interface {
	Point
	Color() Color32
	SetColor(c Color32)
}

type pointCImpl struct {
	Point
	c Color32
}

func NewPointC(x, y, z float64, c Color32) PointC {
	return &pointCImpl{NewPoint(x, y, z), c}
}

func (p *pointCImpl) Color() Color32 {
	return p.c
}
func (p *pointCImpl) SetColor(c Color32) {
	p.c = c
}
