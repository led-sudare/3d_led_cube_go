package ledlib

import (
	"ledlib/util"
	"time"
)

type ObjectScrolledBitmap struct {
	timer  Timer
	offset int
	z      int
	loop   bool
	image  util.Image2D
}

func NewObjectScrolledBitmap(path string, z int, updateRate time.Duration, loop bool) LedManagedObject {
	obj := ObjectScrolledBitmap{}
	obj.timer = NewTimer(updateRate)
	obj.image = util.NewImage2DWithPath(path)
	obj.z = z
	obj.loop = loop
	return &obj
}

func (o *ObjectScrolledBitmap) Draw(cube util.Image3D) {

	o.offset = int(o.timer.GetPastCount()) - LedWidth + 1
	for x := 0; x < LedWidth; x++ {
		for y := 0; y < LedHeight; y++ {
			var c util.Color32
			if o.loop {
				c = o.image.GetAt(x+(o.offset%o.image.GetWidth()), y)
			} else {
				c = o.image.GetAt(x+o.offset, y)
			}
			if c != nil && !c.IsOff() {
				cube.SetAt(x, y, o.z, c)
			}
		}
	}
}
func (o *ObjectScrolledBitmap) IsExpired() bool {
	if o.loop {
		return false
	}
	if o.offset >= LedWidth {
		return true
	}
	return false
}
