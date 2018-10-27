package ledlib

import (
	"ledlib/util"
	"math/rand"
	"time"
)

type FilterBkStars struct {
	filterObjects *FilterObjects
	timer         Timer
}

func NewFilterBkStars(canvas LedCanvas) LedCanvas {
	filter := FilterBkStars{}
	filter.filterObjects = NewFilterObjects(canvas)
	filter.timer = NewTimer(3000 * time.Millisecond)
	rand.Seed(time.Now().UnixNano())

	filter.addTwinkleStars()

	return &filter
}

func (f *FilterBkStars) addTwinkleStars() {
	f.filterObjects.Append(createTwinkleStar())
	f.filterObjects.Append(createTwinkleStar())
	f.filterObjects.Append(createTwinkleStar())
	f.filterObjects.Append(createTwinkleStar())
	f.filterObjects.Append(createTwinkleStar())
	f.filterObjects.Append(createTwinkleStar())
	f.filterObjects.Append(createTwinkleStar())
	f.filterObjects.Append(createShootingStar())
}

func createTwinkleStar() LedManagedObject {

	x, y, z := rand.Intn(LedWidth-1), rand.Intn(LedHeight-1)/2, rand.Intn(LedDepth-1)
	size := 1
	return NewObjectTwinkleStar(x, y, z, size, 2800*time.Millisecond)
}
func createShootingStar() LedManagedObject {

	y, z := rand.Intn(LedHeight-1)/2, rand.Intn(LedDepth-1)
	size := 1
	return NewObjectShootingStar(LedWidth, y, z, size)
}

func (f *FilterBkStars) Show(c util.ImmutableImage3D, param LedCanvasParam) {

	if f.timer.IsPast() {
		f.addTwinkleStars()
	}

	f.filterObjects.Show(c, param)
}
