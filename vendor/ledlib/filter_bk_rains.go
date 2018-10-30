package ledlib

import (
	"ledlib/servicegateway"
	"ledlib/util"
	"time"
)

type FilterBkRains struct {
	filterObjects *FilterObjects
}

func NewFilterBkRains(canvas LedCanvas) LedCanvas {
	filter := FilterBkRains{}
	filter.filterObjects = NewFilterObjects(canvas)
	filter.filterObjects.Append(NewObjectVerticalScrolledBitmap(
		"/asset/image/rain/rain1.png", 3, 35*time.Millisecond, true))
	filter.filterObjects.Append(NewObjectVerticalScrolledBitmap(
		"/asset/image/rain/rain2.png", 7, 40*time.Millisecond, true))
	//	filter.filterObjects.Append(NewObjectVerticalScrolledBitmap(
	//		"/asset/image/rain/rain3.png", 7, 50*time.Millisecond, true))

	servicegateway.GetAudigoSeriveGateway().Play("se_rain.wav", true, false)

	return &filter
}

func (f *FilterBkRains) Show(c util.ImmutableImage3D, param LedCanvasParam) {
	f.filterObjects.Show(c, param)
}
