package ledlib

import (
	"ledlib/util"
)

type SharedLedImage3D interface {
	GetImage(id string) util.Image3D
	ClearImage(id string)
}

var sharedLedImage3D SharedLedImage3D

func GetSharedLedImage3D(id string) util.Image3D {
	return getSharedLedImage3DInstance().GetImage(id)
}

func ClearSharedLedImage3D(id string) {
	getSharedLedImage3DInstance().ClearImage(id)
}

func getSharedLedImage3DInstance() SharedLedImage3D {
	if sharedLedImage3D == nil {
		sharedLedImage3D = newSharedLedImage3D()
	}
	return sharedLedImage3D
}

func newSharedLedImage3D() SharedLedImage3D {
	return &sharedLedImage3DImpl{make(map[string]util.Image3D)}
}

type sharedLedImage3DImpl struct {
	images map[string]util.Image3D
}

func (o *sharedLedImage3DImpl) GetImage(id string) util.Image3D {
	if i, ok := o.images[id]; !ok {
		o.images[id] = NewLedImage3D()
		return o.images[id]
	} else {
		return i
	}
}

func (o *sharedLedImage3DImpl) ClearImage(id string) {
	delete(o.images, id)
}
