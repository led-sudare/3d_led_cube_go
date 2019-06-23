package ledlib

import (
	"3d_led_cube_go/ledlib/servicegateway"
	"3d_led_cube_go/ledlib/util"
	"math"
	"time"
)

type FilterWakame struct {
	canvas LedCanvas
	timer  Timer
	cube   util.Image3D
}

func NewFilterWakame(canvas LedCanvas) LedCanvas {
	f := FilterWakame{}
	f.canvas = canvas
	f.timer = NewTimer(10 * time.Millisecond)
	f.cube = NewLedImage3D()
	servicegateway.GetAudigoSeriveGateway().Play("bgm_wave.wav", true, false)
	return &f
}

func (f *FilterWakame) Show(c util.ImmutableImage3D, param LedCanvasParam) {
	f.cube.Clear()

	yWaveLength := float64(3.0 * math.Pi)
	yWaveDepth := 1.5
	yDot := yWaveLength / LedDepth
	yStart := float64(f.timer.GetPastCount()) / 30 * 2 * yDot

	xWaveLength := float64(3.0 * math.Pi)
	xWaveDepth := 1.5
	xDot := xWaveLength / LedWidth
	xStart := float64(f.timer.GetPastCount()) / 30 * 2 * xDot

	c.ConcurrentForEach(func(x, y, z int, c util.Color32) {
		z0 := z + util.RoundToInt(
			(math.Sin(xDot*float64(x)+xStart)*xWaveDepth)+
				(math.Sin(yDot*float64(y)+yStart)+yWaveDepth))
		f.cube.SetAt(x, y, z0, c)
	})
	f.canvas.Show(f.cube, param)
}
