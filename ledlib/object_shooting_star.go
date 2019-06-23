package ledlib

import (
	"3d_led_cube_go/ledlib/util"
	"time"
)

type ObjectShootingStar struct {
	x, y, z, size int
	straitOffsets [][]int
	timer         Timer
}

func NewObjectShootingStar(x, y, z, size int) LedManagedObject {
	obj := ObjectShootingStar{}
	obj.x, obj.y, obj.z, obj.size = x, y, z, size

	obj.straitOffsets = make([][]int, 6)
	obj.straitOffsets[0] = []int{-1, 0, 0}
	obj.straitOffsets[1] = []int{-1, 0, 0}
	obj.straitOffsets[2] = []int{0, -1, 0}
	obj.straitOffsets[3] = []int{0, 1, 0}
	obj.straitOffsets[4] = []int{0, 0, 1}
	obj.straitOffsets[5] = []int{0, 0, -1}

	obj.timer = NewTimer(20 * time.Millisecond)

	return &obj
}

func (o *ObjectShootingStar) Draw(cube util.Image3D) {

	if o.timer.IsPast() {
		o.x--
		o.y++
	}

	for i := 0; i < 8; i++ {
		hsv := &util.HSV{0.2, 1, 1 / (float64(i + 1))}
		cube.SetAt(o.x+i, o.y-i, o.z, hsv.RGB())
	}
	o.drawStar(cube, 1, o.straitOffsets)

}

func (o *ObjectShootingStar) IsExpired() bool {
	return o.x < 0 || o.y < 0
}

func (o *ObjectShootingStar) drawStar(cube util.Image3D, width float64, offsets [][]int) {

	for i := 0; i <= util.RoundToInt(width); i++ {
		for _, offset := range offsets {
			offsetX, offsetY, offsetZ := offset[0], offset[1], offset[2]
			cube.SetAt(
				o.x+offsetX*i,
				o.y+offsetY*i,
				o.z+offsetZ*i,
				util.NewColorFromUint32(0xffff00))
		}
	}
}
