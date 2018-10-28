package ledlib

import (
	"ledlib/servicegateway"
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
	colors  []util.Color32
}

func NewObjectRain(colors []util.Color32) LedManagedObject {
	rain := ObjectRain{}
	rain.timer = NewTimer(50 * time.Millisecond)
	rain.x = rand.Intn(LedWidth)
	rain.z = rand.Intn(LedDepth)
	rain.y = 0
	rain.colors = colors

	return &rain
}

func (o *ObjectRain) Draw(cube util.Image3D) {
	o.y = o.y*0.8 + float64(o.timer.GetPastCount())/6

	if o.IsExpired() {
		return
	}
	for i, color := range o.colors {
		cube.SetAt(o.x, util.RoundToInt(o.y)-i, o.z, color)
	}
}

func (o *ObjectRain) IsExpired() bool {
	if int(o.y) > LedHeight+len(o.colors) {
		return true
	}
	return false
}

//////////////////////////

type FilterBkRains struct {
	filterObjects *FilterObjects
	timer         Timer
	colors        []util.Color32
}

func NewFilterBkRains(canvas LedCanvas) LedCanvas {
	rand.Seed(time.Now().UnixNano())
	filter := FilterBkRains{}
	filter.timer = NewTimer(400 * time.Millisecond)
	filter.filterObjects = NewFilterObjects(canvas)

	filter.colors = make([]util.Color32, 5)
	for i, _ := range filter.colors {
		hsv := &util.HSV{0.6, 1, 1 / (float64(i + 1))}
		filter.colors[i] = hsv.RGB()
	}

	servicegateway.GetAudigoSeriveGateway().Play("se_rain.wav", true, false)

	return &filter
}

func (f *FilterBkRains) Show(c util.ImmutableImage3D, param LedCanvasParam) {
	cube := c.Copy()
	if f.timer.IsPast() && f.filterObjects.Len() < 8 {
		f.filterObjects.Append(NewObjectRain(f.colors))
		//		f.filterObjects.Append(NewObjectRain())
	}
	f.filterObjects.Show(cube, param)
}
