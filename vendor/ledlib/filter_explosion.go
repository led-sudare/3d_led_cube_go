package ledlib

import (
	"ledlib/servicegateway"
	"ledlib/util"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

type FilterExplosion struct {
	canvas    LedCanvas
	timer     Timer
	speeds    util.Data3D
	centers   util.Data3D
	cube      util.Image3D
	dimension int
	t         float64
	sin       float64
	preSign   int
}

func getSign(value float64) int {
	if value >= 0 {
		return 1
	} else {
		return -1
	}
}

func NewFilterExplosion(canvas LedCanvas, dimension int) LedCanvas {
	f := &FilterExplosion{}

	f.canvas = canvas
	f.timer = NewTimer(50 * time.Millisecond)
	f.speeds = NewLedData3D()
	f.centers = NewLedData3D()
	f.cube = NewLedImage3D()
	f.preSign = -1

	f.dimension = dimension
	rand.Seed(time.Now().UnixNano())
	return f
}

func (f *FilterExplosion) getSpeed(x, y, z int) []float64 {
	//	ctrl := 1
	if f.speeds.GetAt(x, y, z) == nil {
		f.speeds.SetAt(x, y, z,
			[]float64{
				rand.Float64()*1.5 + 1.5,
				rand.Float64()*2 + 2})
	}
	return f.speeds.GetAt(x, y, z).([]float64)
}

func (f *FilterExplosion) getCenter(x, y, z int) util.Point {
	//	ctrl := 1
	if f.centers.GetAt(x, y, z) == nil {
		if f.dimension == 3 {
			f.centers.SetAt(x, y, z,
				util.NewPoint(
					LedWidth/2.0+rand.Float64()*2-1,
					LedHeight/4.0*3+rand.Float64()*2-1,
					LedDepth/2.0+rand.Float64()*2-1))
		} else {
			f.centers.SetAt(x, y, z,
				util.NewPoint(
					LedWidth/2.0+rand.Float64()*2-1,
					LedHeight/4.0*3+rand.Float64()*2-1,
					rand.Float64()*2-2))
		}
	}
	return f.centers.GetAt(x, y, z).(util.Point)
}

func (f *FilterExplosion) Show(c util.ImmutableImage3D, param LedCanvasParam) {
	f.cube.Clear()
	if f.timer.IsPast() {
		p1 := 0.5
		f.t += 0.01 + p1*0.28
		f.sin = math.Sin(f.t)
		sign := getSign(f.sin)

		if f.preSign < 0 && sign > 0 {
			servicegateway.GetAudigoSeriveGateway().Play("se_explosion.wav", true, false)
		}
		f.preSign = sign
		//		f.add = (f.add + 1)
	}
	c.ConcurrentForEach(func(x, y, z int, c util.Color32) {
		cent := f.getCenter(x, y, z)
		if f.sin > 0 {
			for _, speed := range f.getSpeed(x, y, z) {
				pt := mat.NewDense(3, 1, []float64{
					float64(x) - cent.X(),
					float64(y) - cent.Y(),
					float64(z) - cent.Z()})
				rate := mat.NewDense(1, 1, []float64{
					f.sin * 3.5 * speed})
				var result mat.Dense
				result.Mul(pt, rate)
				f.cube.SetAt(
					util.RoundToInt(result.At(0, 0)+cent.X()),
					util.RoundToInt(result.At(1, 0)+cent.Y()),
					util.RoundToInt(result.At(2, 0)+cent.Z()), c)

			}
		} else {
			f.cube.SetAt(x, y, z, c)
		}
	})
	f.canvas.Show(f.cube, param)
}
