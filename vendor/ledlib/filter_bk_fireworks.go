package ledlib

import (
	"ledlib/util"
	"time"
)

type FilterBkFireworks struct {
	canvas          LedCanvas
	objectFireworks LedManagedObject
	timer           Timer
}

func NewFilterBkFireworks(canvas LedCanvas) LedCanvas {
	filter := FilterBkFireworks{}
	filter.canvas = canvas
	filter.objectFireworks = NewManagedObjectFireworks()
	filter.timer = NewTimer(80 * time.Millisecond)

	return &filter
}

func (f *FilterBkFireworks) Show(c util.ImmutableImage3D, param LedCanvasParam) {
	filtered := c.Copy()
	f.objectFireworks.Draw(filtered)
	f.canvas.Show(filtered, param)
}
