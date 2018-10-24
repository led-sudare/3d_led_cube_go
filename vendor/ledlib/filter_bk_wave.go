package ledlib

import (
	"ledlib/util"
	"math"
	"time"
)

var waveDepth = 0.8
var colorWave = util.NewColorFromRGB(0, 0, 255)
var rangeWave = []int{27, 50}

type FilterBkWave struct {
	canvas LedCanvas
	timer  Timer
	cube   util.Image3D
}

func NewFilterBkWave(canvas LedCanvas) LedCanvas {
	f := &FilterBkWave{}

	f.canvas = canvas
	f.timer = NewTimer(10 * time.Millisecond)
	f.cube = NewLedImage3D()
	return f
}

func (f *FilterBkWave) Show(cube util.ImmutableImage3D, param LedCanvasParam) {

	f.cube = cube.Copy()

	zWaveLength := float64(2.0 * math.Pi)
	zWaveDepth := float64(waveDepth)
	zDot := zWaveLength / LedDepth
	zStart := float64(f.timer.GetPastCount()) / 30 * 2 * zDot

	xWaveLength := float64(3.0 * math.Pi)
	xWaveDepth := float64(waveDepth)
	xDot := xWaveLength / LedWidth
	xStart := float64(f.timer.GetPastCount()) / 30 * 2 * xDot

	cube.ConcurrentForEachAll(func(x, y, z int, c util.Color32) {
		if rangeWave[0] < y && y < rangeWave[1] {
			y0 := y + util.RoundToInt(
				(xWaveDepth+math.Sin(xDot*float64(x)+xStart)*xWaveDepth)+
					(zWaveDepth+math.Sin(zDot*float64(z)+zStart)+zWaveDepth))
			f.cube.SetAt(x, y0, z, colorWave)
		}
	})
	f.canvas.Show(f.cube, param)
}
