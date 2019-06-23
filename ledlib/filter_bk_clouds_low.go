package ledlib

import (
	"3d_led_cube_go/ledlib/util"
	"time"
)

type FilterBkCloudsLow struct {
	filterObjects *FilterObjects
}

func NewFilterBkCloudsLow(canvas LedCanvas) LedCanvas {
	filter := FilterBkCloudsLow{}
	filter.filterObjects = NewFilterObjects(canvas)
	filter.filterObjects.Append(NewObjectHorizontalScrolledBitmap(
		"/asset/image/cloud/cloud1_low.png", 7, 200*time.Millisecond, true))
	filter.filterObjects.Append(NewObjectHorizontalScrolledBitmap(
		"/asset/image/cloud/cloud2_low.png", 6, 300*time.Millisecond, true))

	return &filter
}

func (f *FilterBkCloudsLow) Show(c util.ImmutableImage3D, param LedCanvasParam) {
	f.filterObjects.Show(c, param)
}
