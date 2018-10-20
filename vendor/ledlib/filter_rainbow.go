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
}

func NewFilterRainbow(canvas LedCanvas) LedCanvas {
	f := FilterRainbow{}
	f.canvas = canvas
	f.timer = NewTimer(10 * time.Millisecond)
	servicegateway.GetAudigoSeriveGateway().Play("se_rainbow.wav", true, false)
	return &f
}

func (f *FilterRainbow) Show(c util.Image3D, param LedCanvasParam) {

	p := float64(f.timer.GetPastCount()) / 10.0
	c.ConcurrentForEach(func(x, y, z int, color util.Color32) {

		rainbow := (p + float64(-x-z-y)) / rainbowGrad
		h := (math.Sin(rainbow) + 1) / 2

		hsv := &util.HSV{h, 1, 1}

		c.SetAt(x, y, z, hsv.RGB())

	})
	f.canvas.Show(c, param)
}
