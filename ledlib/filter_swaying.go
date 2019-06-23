package ledlib

import (
	"3d_led_cube_go/ledlib/servicegateway"
	"3d_led_cube_go/ledlib/util"
	"math"
	"time"
)

type FilterSwaying struct {
	canvas LedCanvas
	timer  Timer
	cube   util.Image3D
}

func NewFilterSwaying(canvas LedCanvas) LedCanvas {
	f := FilterSwaying{}
	f.canvas = canvas
	f.timer = NewTimer(25 * time.Millisecond)
	f.cube = NewLedImage3D()
	servicegateway.GetAudigoSeriveGateway().Play("se_wind.wav", true, false)
	return &f
}

func (f *FilterSwaying) Show(c util.ImmutableImage3D, param LedCanvasParam) {
	f.cube.Clear()

	swingSpeed := 0.08

	swaying := math.Cos(float64(f.timer.GetPastCount())*swingSpeed) * 0.020
	t := 4.0
	sin := math.Sin(swaying * math.Pi * t)
	cos := math.Cos(swaying * math.Pi * t)

	dx := float64(LedWidth / 2)
	dy := float64(LedHeight)

	c.ConcurrentForEach(func(x, y, z int, c util.Color32) {

		xx := ((float64(x)-dx)*cos + (float64(y)-dy)*sin) + dx
		yy := (-(float64(x)-dx)*sin + (float64(y)-dy)*cos) + dy

		f.cube.SetAt(util.RoundToInt(xx), util.RoundToInt(yy), z, c)
	})

	f.canvas.Show(f.cube, param)
}
