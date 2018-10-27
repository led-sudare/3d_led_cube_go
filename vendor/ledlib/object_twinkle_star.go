package ledlib

import (
	"ledlib/util"
	"math"
	"time"
)

type ObjectTwinkleStar struct {
	x, y, z, size int
	timeout       time.Duration
	bornAt        time.Time
	straitOffsets [][]int
	tiltOffsets   [][]int
	twinkleSpeed  float64
	first0        bool
	timer         Timer
}

func NewObjectTwinkleStar(x, y, z, size int, timeout time.Duration) LedManagedObject {
	obj := ObjectTwinkleStar{}
	obj.x, obj.y, obj.z, obj.size = x, y, z, size
	obj.bornAt = time.Now()
	obj.timeout = timeout

	obj.straitOffsets = make([][]int, 6)
	obj.straitOffsets[0] = []int{-1, 0, 0}
	obj.straitOffsets[1] = []int{1, 0, 0}
	obj.straitOffsets[2] = []int{0, -1, 0}
	obj.straitOffsets[3] = []int{0, 1, 0}
	obj.straitOffsets[4] = []int{0, 0, 1}
	obj.straitOffsets[5] = []int{0, 0, -1}

	obj.tiltOffsets = make([][]int, 8)
	obj.tiltOffsets[0] = []int{1, 1, 1}    //
	obj.tiltOffsets[1] = []int{-1, -1, -1} //
	obj.tiltOffsets[2] = []int{1, 1, -1}
	obj.tiltOffsets[3] = []int{-1, -1, 1}
	obj.tiltOffsets[4] = []int{-1, 1, 1}
	obj.tiltOffsets[5] = []int{1, -1, -1}
	obj.tiltOffsets[6] = []int{1, -1, 1}
	obj.tiltOffsets[7] = []int{-1, 1, -1}

	obj.first0 = false
	obj.timer = NewTimer(20 * time.Millisecond)
	//	obj.twinkleSpeed = rand.Float64() / 10
	obj.twinkleSpeed = 0.05

	return &obj
}

func (o *ObjectTwinkleStar) Draw(cube util.Image3D) {

	width := math.Cos(float64(o.timer.GetPastCount())*o.twinkleSpeed) * float64(o.size)

	if !o.first0 {
		if width > 0 && width < 0.1 {
			o.first0 = true
		} else {
			return
		}
	}
	sign := util.GetSign(width)
	if sign < 0 {
		o.drawStar(cube, -width, o.straitOffsets)
	} else {
		o.drawStar(cube, width, o.straitOffsets)
	}
}

func (o *ObjectTwinkleStar) IsExpired() bool {
	return time.Now().Sub(o.bornAt) > o.timeout
}

func (o *ObjectTwinkleStar) drawStar(cube util.Image3D, width float64, offsets [][]int) {

	if width < 0.3 {
		return
	}

	for i := 0; i <= util.RoundToInt(width); i++ {
		for _, offset := range offsets {
			offsetX, offsetY, offsetZ := offset[0], offset[1], offset[2]
			hsv := &util.HSV{0.2, 1, 1 / (float64(i + 1))}
			cube.SetAt(
				o.x+offsetX*i,
				o.y+offsetY*i,
				o.z+offsetZ*i,
				//				util.DarkenWithRatio(o.starColor, 100-uint32(i)*30))
				hsv.RGB())
		}
	}
}
