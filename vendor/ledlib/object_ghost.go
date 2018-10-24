package ledlib

import (
	"ledlib/servicegateway"
	"ledlib/util"
	"time"
)

type ObjectGhost struct {
	cube  util.Image3D
	timer Timer
}

func NewObjectGhost() LedObject {
	paths := []string{
		"/asset/image/ghost/ghost1.png",
		"/asset/image/ghost/ghost2.png",
		"/asset/image/ghost/ghost3.png",
		"/asset/image/ghost/ghost4.png",
		"/asset/image/ghost/ghost4.png",
		"/asset/image/ghost/ghost3.png",
		"/asset/image/ghost/ghost2.png",
		"/asset/image/ghost/ghost5.png",
	}
	bmp := ObjectGhost{}
	bmp.cube = NewLedImage3DFromPaths(paths)
	bmp.timer = NewTimer(4 * time.Second)
	return &bmp
}

func (b *ObjectGhost) GetImage3D(param LedCanvasParam) util.ImmutableImage3D {
	if b.timer.IsPast() {
		servicegateway.GetAudigoSeriveGateway().Play("se_obake2.wav", false, false)
	}
	return b.cube
}
