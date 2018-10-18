package ledlib

import (
	"ledlib/util"
)

type ObjectBitmap struct {
	cube util.Image3D
}

func NewObjectBitmap(paths []string) LedObject {
	bmp := ObjectBitmap{}
	bmp.cube = NewLedImage3DFromPaths(paths)
	return &bmp
}

func (b *ObjectBitmap) GetImage3D(param LedCanvasParam) util.Image3D {
	return b.cube
}
