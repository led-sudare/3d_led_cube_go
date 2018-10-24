package ledlib

import (
	"ledlib/util"
)

type ObjectPainting struct {
	cube util.ImmutableImage3D
}

func NewObjectPainting() LedObject {
	obj := ObjectPainting{GetSharedLedImage3D(paintingSharedObjectID)}
	return &obj
}

func (b *ObjectPainting) GetImage3D(param LedCanvasParam) util.ImmutableImage3D {
	return b.cube
}
