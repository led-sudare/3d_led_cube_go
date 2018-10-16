package ledlib

import (
	"ledlib/util"
	"math/rand"
	"time"
)

func red(ix int) uint8 {
	i := ix % 90
	switch {
	case i < 30:
		return uint8(i * 255 / 30)
	case i < 60:
		return uint8((60 - i) * 255 / 30)
	default:
		return 0
	}
}

func darken(c util.Color32) util.Color32 {
	r := ((c.Uint32() & 0xff0000) * 49 / 50) & 0xff0000
	g := ((c.Uint32() & 0xff00) * 49 / 50) & 0xff
	b := ((c.Uint32() & 0xff) * 49 / 50) & 0xff
	return util.NewColorFromUint32(uint32(r + g + b))
}

func rgb(ix float64) util.Color32 {
	n := util.FloorToInt(ix * 1 * 90)
	return util.NewColorFromRGB(red(n), red(n+30), red(n+60))
}

type ObjectFireworks struct {
	cube        util.Image3D
	poss        []util.PointC
	vs          []util.Point
	ix          int
	addTimer    *Timer
	updateTimer *Timer
}

func localNewObjectFireworks() *ObjectFireworks {
	obj := ObjectFireworks{}
	obj.cube = NewLedImage3D()
	obj.poss = make([]util.PointC, 0)
	obj.vs = make([]util.Point, 0)
	obj.ix = 0
	obj.addTimer = NewTimer(1100 * time.Millisecond)
	obj.updateTimer = NewTimer(80 * time.Millisecond)

	return &obj
}

func NewObjectFireworks() LedObject {
	return localNewObjectFireworks()
}
func NewManagedObjectFireworks() LedManagedObject {
	return localNewObjectFireworks()
}

func (b *ObjectFireworks) IsExpired() bool {
	return false
}
func (b *ObjectFireworks) Draw(cube util.Image3D) {
	if b.addTimer.IsPast() {
		cx := LedWidth * rand.Float64()
		cy := LedHeight * rand.Float64()
		cz := LedDepth * rand.Float64()

		for i := 0; i < 1000; i++ {
			sf := util.GetSphereFace()
			b.vs = append(b.vs, sf)
			b.poss = append(b.poss, util.NewPointC(cx, cy, cz, rgb(sf.Len())))
		}
	}

	dIdx := make([]int, 0)

	isPast := b.updateTimer.IsPast()

	for i, p := range b.poss {
		v := b.vs[i]
		if util.CanShow(p, LedWidth, LedHeight, LedDepth) {
			cube.SetAt(util.RoundToInt(p.X()),
				util.RoundToInt(p.Y()),
				util.RoundToInt(p.Z()),
				p.Color())
		} else {
			dIdx = append(dIdx, i)
		}
		if isPast {
			p.Add(v)
			p.SetColor(darken(p.Color()))
		}
	}
	for i := len(dIdx) - 1; i >= 0; i-- {
		b.vs = append(b.vs[:dIdx[i]], b.vs[dIdx[i]+1:]...)
		b.poss = append(b.poss[:dIdx[i]], b.poss[dIdx[i]+1:]...)
	}

}

func (b *ObjectFireworks) GetImage3D() util.Image3D {
	b.cube.Clear()
	b.Draw(b.cube)
	return b.cube
}
