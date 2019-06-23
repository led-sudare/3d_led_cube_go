package ledlib

import (
	"3d_led_cube_go/ledlib/servicegateway"
	"3d_led_cube_go/ledlib/util"
	"math"
	"time"
)

type FilterWave struct {
	canvas LedCanvas
	cube   util.Image3D
	timer  Timer
}

func NewFilterWave(canvas LedCanvas) LedCanvas {
	f := FilterWave{}
	f.canvas = canvas
	f.cube = NewLedImage3D()
	f.timer = NewTimer(10 * time.Millisecond)
	servicegateway.GetAudigoSeriveGateway().Play("bgm_wave.wav", true, false)
	return &f
}

func (f *FilterWave) Show(c util.ImmutableImage3D, param LedCanvasParam) {

	f.cube.Clear()
	offset := float64(f.timer.GetPastCount()) / 200

	yWaveLen := 3.0 * math.Pi
	yWaveDepth := 1.5
	yDot := yWaveLen / float64(LedHeight)
	yStart := (offset * 10) + yDot

	xWaveLen := 3.0 * math.Pi
	xWaveDepth := 1.5
	xDot := xWaveLen / float64(LedWidth)
	xStart := (offset * 5) * xDot

	c.ConcurrentForEach(func(x, y, z int, color util.Color32) {

		zz := z + int(xWaveDepth+math.Sin(xDot*float64(x)+xStart)*xWaveDepth) +
			int(yWaveDepth+math.Sin(yDot*float64(y)+yStart)*yWaveDepth)

		f.cube.SetAt(x, y, zz, color)

	})
	f.canvas.Show(f.cube, param)
}
