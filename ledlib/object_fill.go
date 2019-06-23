package ledlib

import (
	"3d_led_cube_go/ledlib/util"
)

type ObjectFill struct {
	cube util.Image3D
}

func NewObjectFill(c util.Color32) LedObject {
	obj := ObjectFill{NewLedImage3D()}
	obj.cube.Fill(c)
	return &obj
}

func (b *ObjectFill) GetImage3D(param LedCanvasParam) util.ImmutableImage3D {
	return b.cube
}
