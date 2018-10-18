package ledlib

import (
	"ledlib/util"
)

type ObjectPainting struct {
	cube util.Image3D
}

func NewObjectPainting() LedObject {
	obj := ObjectPainting{GetSharedLedImage3D(paintingSharedObjectID)}
	return &obj
}

func (b *ObjectPainting) GetImage3D(param LedCanvasParam) util.Image3D {
	return b.cube
}
