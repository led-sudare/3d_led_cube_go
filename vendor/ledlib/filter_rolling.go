package ledlib

import (
	"ledlib/servicegateway"
	"ledlib/util"
	"time"
)

const MaxAdd = 300

type FilterRolling struct {
	canvas    LedCanvas
	add       int
	lastCount int
	timer     Timer
	cube      util.Image3D
}

func NewFilterRolling(canvas LedCanvas) LedCanvas {
	return &FilterRolling{canvas, 0, 0, NewTimer(100 * time.Millisecond), NewLedImage3D()}
}

func (f *FilterRolling) Show(c util.ImmutableImage3D, param LedCanvasParam) {
	f.cube.Clear()
	if f.timer.IsPast() {
		f.add = (f.add + 1)
	}
	count := f.add / LedHeight
	if count > f.lastCount {
		f.lastCount = count
		servicegateway.GetAudigoSeriveGateway().Play("se_rolldown.wav", false, false)
	}
	c.ConcurrentForEach(func(x, y, z int, c util.Color32) {
		f.cube.SetAt(x, (y+f.add)%LedHeight, z, c)
	})
	f.canvas.Show(f.cube, param)
}
