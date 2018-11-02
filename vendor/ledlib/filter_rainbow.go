package ledlib

import (
	"ledlib/servicegateway"
	"ledlib/util"
	"math"
	"time"
)

const (
	rainbowGrad = 5.0
)

type FilterRainbow struct {
	canvas LedCanvas
	timer  Timer
	cube   util.Image3D
}

func NewFilterRainbow(canvas LedCanvas) LedCanvas {
	f := FilterRainbow{}
	f.canvas = canvas
	f.timer = NewTimer(10 * time.Millisecond)
	f.cube = NewLedImage3D()
	servicegateway.GetAudigoSeriveGateway().Play("se_rainbow.wav", true, false)
	return &f
}

func (f *FilterRainbow) Show(c util.ImmutableImage3D, param LedCanvasParam) {

	f.cube = c.Copy()

	p := float64(f.timer.GetPastCount()) / 10.0
	c.ConcurrentForEach(func(x, y, z int, color util.Color32) {

		rainbow := (p + float64(-x-z-y)) / rainbowGrad
		h := (math.Sin(rainbow) + 1) / 2
		f.cube.SetAt(x, y, z, util.GetRainbow(h))

	})
	f.canvas.Show(f.cube, param)
}
