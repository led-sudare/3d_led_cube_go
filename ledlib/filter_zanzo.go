package ledlib

import (
	"3d_led_cube_go/ledlib/util"
	"time"
)

type FilterZanzo struct {
	canvas LedCanvas
	cube   util.Image3D
	buffer util.Image3D
	timer  Timer
}

func NewFilterZanzo(canvas LedCanvas) LedCanvas {
	f := FilterZanzo{}
	f.canvas = canvas
	f.cube = NewLedImage3D()
	f.buffer = NewLedImage3D()
	f.timer = NewTimer(120 * time.Millisecond)
	return &f
}

func (f *FilterZanzo) Show(c util.ImmutableImage3D, param LedCanvasParam) {

	f.cube.Clear()
	//colorIndex := f.timer.GetPastCount()
	startIndex := int(f.timer.GetPastCount() % LedDepth)
	util.ConcurrentEnum(0, LedWidth, func(x int) {
		for y := 0; y < LedHeight; y++ {
			f.buffer.SetAt(x, y, startIndex, nil)
		}
	})

	c.ConcurrentForEach(func(x, y, z int, color util.Color32) {
		f.buffer.SetAt(x, y, startIndex, color)
	})
	f.buffer.ConcurrentForEach(func(x, y, z int, color util.Color32) {
		zz := ((startIndex + LedDepth - z) % LedDepth)
		if zz != 0 {
			color = util.GetRainbow(float64(z)/15 + 0.1)
		}
		f.cube.SetAt(x, y, zz, color)
	})

	f.canvas.Show(f.cube, param)
}
