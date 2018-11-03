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
	EditSharedLedImage3D(realsenseSharedObjectID,
		func(editable util.Image3D) {
			b.cube = editable.Copy()
		})
	return b.cube
}
