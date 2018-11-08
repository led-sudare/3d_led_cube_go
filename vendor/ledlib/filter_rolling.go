package ledlib

import (
	"ledlib/servicegateway"
	"ledlib/util"
	"time"
)

const MaxAdd = 300

type FilterRolling struct {
	canvas LedCanvas
	add    int
	timer  Timer
	cube   util.Image3D
}

func NewFilterRolling(canvas LedCanvas) LedCanvas {
	return &FilterRolling{canvas, 0, NewTimer(100 * time.Millisecond), NewLedImage3D()}
}

func getDrawPoint(y, add int) int {

	if add >= 0 {
		return (y + add) % LedHeight
	}
	newadd := add % LedHeight
	return (y + LedHeight + newadd) % LedHeight
}

func (f *FilterRolling) Show(c util.ImmutableImage3D, param LedCanvasParam) {
	f.cube.Clear()
	if f.timer.IsPast() {
		f.add = (f.add - 1)
	}
	count := getDrawPoint(0, f.add)
	if count == LedHeight/2 {
		servicegateway.GetAudigoSeriveGateway().Play("se_rollup.wav", false, false)
	}
	c.ConcurrentForEach(func(x, y, z int, c util.Color32) {
		f.cube.SetAt(x, getDrawPoint(y, f.add), z, c)
	})
	f.canvas.Show(f.cube, param)
}
