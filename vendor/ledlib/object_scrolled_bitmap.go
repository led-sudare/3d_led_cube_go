package ledlib

import (
	"ledlib/util"
	"time"
)

type ObjectScrolledBitmap struct {
	timer                  Timer
	z                      int
	loop                   bool
	image                  util.Image2D
	xDirection, yDirection int
	xOffset, yOffset       int
}

func NewObjectHorizontalScrolledBitmap(path string, z int, updateRate time.Duration, loop bool) LedManagedObject {
	obj := ObjectScrolledBitmap{}
	obj.timer = NewTimer(updateRate)
	obj.image = util.NewImage2DWithPath(path)
	obj.z = z
	obj.loop = loop

	obj.xDirection = 1
	obj.yDirection = 0

	return &obj
}
func NewObjectVerticalScrolledBitmap(path string, z int, updateRate time.Duration, loop bool) LedManagedObject {
	obj := ObjectScrolledBitmap{}
	obj.timer = NewTimer(updateRate)
	obj.image = util.NewImage2DWithPath(path)
	obj.z = z
	obj.loop = loop

	obj.xDirection = 0
	obj.yDirection = -1

	return &obj
}

func (o *ObjectScrolledBitmap) Draw(cube util.Image3D) {

	xOffset := int(o.timer.GetPastCount())*o.xDirection - (LedWidth+1)*o.xDirection
	yOffset := int(o.timer.GetPastCount())*o.yDirection - (LedHeight+1)*o.yDirection
	for x := 0; x < LedWidth; x++ {
		for y := 0; y < LedHeight; y++ {
			var c util.Color32
			if o.loop {
				//				log.Printf("yOffset:%d, image height:%d, y: %d, y': %d\n", yOffset, o.image.GetHeight(), y, (y+yOffset)%o.image.GetHeight())
				c = o.image.GetAt(util.AbsInt((x+xOffset)%o.image.GetWidth()),
					util.AbsInt((y+yOffset)%o.image.GetHeight()))
			} else {
				c = o.image.GetAt(util.AbsInt(x+xOffset), util.AbsInt(y+yOffset))
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
	if o.xOffset >= LedWidth {
		return true
	}
	if o.yOffset >= LedHeight {
		return true
	}
	return false
}
