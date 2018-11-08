package ledlib

import (
	"ledlib/servicegateway"
	"ledlib/util"
	"math"
	"time"
)

type FilterElastic struct {
	canvas    LedCanvas
	timer     Timer
	cube      util.Image3D
	soundSide bool
}

func NewFilterElastic(canvas LedCanvas) LedCanvas {
	f := FilterElastic{}
	f.canvas = canvas
	f.timer = NewTimer(25 * time.Millisecond)
	f.cube = NewLedImage3D()
	return &f
}

func (f *FilterElastic) Show(c util.ImmutableImage3D, param LedCanvasParam) {
	f.cube.Clear()

	elasticSpeed := 0.08

	degree := float64(f.timer.GetPastCount()) * elasticSpeed
	zoom1 := ((math.Cos(degree) + 1) / 4) + 0.7
	zoom2 := 1 + (0.25 + 0.7) - zoom1

	if f.soundSide {
		if zoom1 > 1 {
			f.soundSide = false
			servicegateway.GetAudigoSeriveGateway().Play("se_zoom_extend1.wav", false, false)
		}
	} else {
		if zoom2 > 1 {
			f.soundSide = true
			servicegateway.GetAudigoSeriveGateway().Play("se_zoom_shrink1.wav", false, false)
		}
	}

	dx := float64(LedWidth / 2)
	dy := float64(LedHeight)
	dz := float64(LedDepth / 2)

	c.ConcurrentForEach(func(x, y, z int, c util.Color32) {

		xx := (float64(x)-dx)*zoom1 + dx
		yy := (float64(y)-dy)*zoom2 + dy
		zz := (float64(z)-dz)*zoom1 + dz

		f.cube.SetAt(int(xx), int(yy), int(zz), c)
		f.cube.SetAt(util.RoundToInt(xx), util.RoundToInt(yy), util.RoundToInt(zz), c)
	})

	f.canvas.Show(f.cube, param)
}
