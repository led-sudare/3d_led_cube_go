package ledlib

import (
	"ledlib/util"
	"math/rand"
	"time"
)

type FilterBkClouds struct {
	filterObjects *FilterObjects
	timer         Timer
}

func NewFilterBkClouds(canvas LedCanvas) LedCanvas {
	filter := FilterBkClouds{}
	filter.filterObjects = NewFilterObjects(canvas)
	filter.timer = NewTimer(1000 * time.Millisecond)
	rand.Seed(time.Now().UnixNano())

	filter.filterObjects.Append(NewObjectCloud(3, 200*time.Millisecond))

	return &filter
}

func (f *FilterBkClouds) Show(c util.ImmutableImage3D, param LedCanvasParam) {

	if f.timer.IsPast() && rand.Intn(4) == 0 {
		z := rand.Intn(LedDepth-4) + 2
		updateRate := rand.Intn(200) + 200

		f.filterObjects.Append(NewObjectCloud(z, time.Duration(updateRate)*time.Millisecond))
	}

	f.filterObjects.Show(c, param)
}
