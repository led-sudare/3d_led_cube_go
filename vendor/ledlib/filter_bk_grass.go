package ledlib

import (
	"ledlib/util"
	"time"
)

type FilterBkGrass struct {
	filterObjects     *FilterObjects
	filterObjectsSnow *FilterObjects
}

func NewFilterBkGrass(canvas LedCanvas) LedCanvas {
	filter := FilterBkGrass{}
	filter.filterObjects = NewFilterObjects(canvas)

	duration := 100 * time.Millisecond
	filter.filterObjects.Append(NewObjectHorizontalScrolledBitmap(
		"/asset/image/grass/grass1.png", 0, duration, true))
	filter.filterObjects.Append(NewObjectHorizontalScrolledBitmap(
		"/asset/image/grass/grass2.png", 1, duration, true))
	filter.filterObjects.Append(NewObjectHorizontalScrolledBitmap(
		"/asset/image/grass/grass1.png", 2, duration, true))
	filter.filterObjects.Append(NewObjectHorizontalScrolledBitmap(
		"/asset/image/grass/grass2.png", 3, duration, true))
	filter.filterObjects.Append(NewObjectHorizontalScrolledBitmap(
		"/asset/image/grass/grass1.png", 4, duration, true))
	filter.filterObjects.Append(NewObjectHorizontalScrolledBitmap(
		"/asset/image/grass/grass3.png", 5, duration, true))

	filter.filterObjectsSnow.Append(NewObjectHorizontalScrolledBitmap(
		"/asset/image/grass/grass1-s.png", 0, duration, true))
	filter.filterObjectsSnow.Append(NewObjectHorizontalScrolledBitmap(
		"/asset/image/grass/grass2-s.png", 1, duration, true))
	filter.filterObjectsSnow.Append(NewObjectHorizontalScrolledBitmap(
		"/asset/image/grass/grass1-s.png", 2, duration, true))
	filter.filterObjectsSnow.Append(NewObjectHorizontalScrolledBitmap(
		"/asset/image/grass/grass2-s.png", 3, duration, true))
	filter.filterObjectsSnow.Append(NewObjectHorizontalScrolledBitmap(
		"/asset/image/grass/grass1-s.png", 4, duration, true))
	filter.filterObjectsSnow.Append(NewObjectHorizontalScrolledBitmap(
		"/asset/image/grass/grass3-s.png", 5, duration, true))

	return &filter
}

func (f *FilterBkGrass) Show(c util.ImmutableImage3D, param LedCanvasParam) {
	if param.HasEffect("filter-snows") {
		f.filterObjectsSnow.Show(c, param)
	} else {
		f.filterObjects.Show(c, param)
	}
}
