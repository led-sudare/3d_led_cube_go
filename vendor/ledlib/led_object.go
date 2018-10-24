package ledlib

import "ledlib/util"

func ShowObject(canvas LedCanvas, obj LedObject, param LedCanvasParam) {
	canvas.Show(obj.GetImage3D(param), param)
}

type LedObject interface {
	GetImage3D(param LedCanvasParam) util.ImmutableImage3D
}
