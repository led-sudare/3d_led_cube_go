package ledlib

import (
	"ledlib/servicegateway"
	"ledlib/util"
	"math"
	"time"
)

type FilterExile struct {
	canvas LedCanvas
	timer  Timer
	cube   util.Image3D
}

func NewFilterExile(canvas LedCanvas) LedCanvas {
	f := FilterExile{}
	f.canvas = canvas
	f.timer = NewTimer(10 * time.Millisecond)
	f.cube = NewLedImage3D()
	servicegateway.GetAudigoSeriveGateway().Play("se_space.wav", true, false)
	return &f
}

func (f *FilterExile) Show(c util.Image3D, param LedCanvasParam) {
	f.cube.Clear()
	speed := 15.0
	buffer := util.NewImage2D(LedWidth, LedHeight)

	c.ConcurrentForEach(func(x, y, z int, color util.Color32) {

		if buffer.GetAt(x, y) == nil {
			buffer.SetAt(x, y, color)
		}
	})

	util.ConcurrentEnum(0, LedWidth, func(x int) {
		for y := 0; y < LedHeight; y++ {
			for z := 0; z < LedDepth; z++ {
				if newcolor := buffer.GetAt(x, y); newcolor != nil {

					step := 3
					p := float64(f.timer.GetPastCount())/speed - float64(z/step)
					sx := math.Sin(p) * 3
					sy := math.Cos(p) * 3

					if z >= 0 && z <= 2 {
					} else {
						h := (math.Sin(float64(f.timer.GetPastCount())/speed+float64(z/step)) + 1) / 2
						hsv := &util.HSV{h, 1, 1}
						rgb := hsv.RGB()
						newcolor = rgb
					}
					f.cube.SetAt(x+util.RoundToInt(sx), y+util.RoundToInt(sy), z, newcolor)
				}
			}
		}

	})

	f.canvas.Show(f.cube, param)
}
