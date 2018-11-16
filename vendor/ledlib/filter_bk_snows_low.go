package ledlib

import (
	"ledlib/util"
	"time"
)

type FilterBkSnowsLow struct {
	filterObjects *FilterObjects
}

func NewFilterBkSnowsLow(canvas LedCanvas) LedCanvas {
	filter := FilterBkSnowsLow{}
	filter.filterObjects = NewFilterObjects(canvas)
	filter.filterObjects.Append(NewObjectVerticalScrolledBitmap(
		"/asset/image/snow/snow1.png", 3, 100*time.Millisecond, true))
	filter.filterObjects.Append(NewObjectVerticalScrolledBitmap(
		"/asset/image/snow/snow2.png", 7, 150*time.Millisecond, true))

	return &filter
}

func (f *FilterBkSnowsLow) Show(c util.ImmutableImage3D, param LedCanvasParam) {
	f.filterObjects.Show(c, param)
}
