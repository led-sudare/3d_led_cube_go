package ledlib

import (
	"ledlib/util"
	"math/rand"
	"time"
)

type ObjectRain struct {
	timer   Timer
	x       int
	z       int
	y       float64
	gravity float64
}

func NewObjectRain() LedManagedObject {
	rain := ObjectRain{}
	rand.Seed(time.Now().UnixNano())
	rain.timer = NewTimer(15 * time.Millisecond)
	rain.x = rand.Intn(LedWidth)
	rain.z = rand.Intn(LedDepth)
	rain.y = 0
	rain.gravity = (rand.Float64() / 5) + 0.3

	return &rain
}

func (o *ObjectRain) Draw(cube util.Image3D) {
	if o.timer.IsPast() {
		o.y = o.y + o.gravity
	}

	if o.IsExpired() {
		return
	}
	for i := 0; i < 5; i++ {
		hsv := &util.HSV{0.6, 1, 1 / (float64(i + 1))}
		cube.SetAt(o.x, util.RoundToInt(o.y)-i, o.z, hsv.RGB())
	}
}
func (o *ObjectRain) IsExpired() bool {
	if o.y > LedHeight {
		return true
	}
	return false
}

//////////////////////////

type FilterBkRains struct {
	filterObjects *FilterObjects
	timer         Timer
}

func NewFilterBkRains(canvas LedCanvas) LedCanvas {
	filter := FilterBkRains{}
	filter.timer = NewTimer(400 * time.Millisecond)
	filter.filterObjects = NewFilterObjects(canvas)

	return &filter
}

func (f *FilterBkRains) Show(c util.Image3D, param LedCanvasParam) {
	cube := c.Copy()
	if f.timer.IsPast() {
		f.filterObjects.Append(NewObjectRain())
		f.filterObjects.Append(NewObjectRain())
	}
	f.filterObjects.Show(cube, param)
}
