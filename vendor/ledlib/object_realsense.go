package ledlib

import (
	"ledlib/util"
)

type ObjectRealsense struct {
	cube util.Image3D
}

func NewObjectRealsense() LedObject {

	obj := ObjectRealsense{}
	return &obj
}

func (b *ObjectRealsense) GetImage3D(param LedCanvasParam) util.ImmutableImage3D {
	b.cube = GetSharedLedImage3D(realsenseSharedObjectID).Copy()
	return b.cube
}
