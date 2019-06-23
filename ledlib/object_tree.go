package ledlib

import "3d_led_cube_go/ledlib/util"

type ObjectTree struct {
	winter util.Image3D
	summer util.Image3D
	spring util.Image3D
}

func NewObjectTree() LedObject {
	winter := []string{
		"/asset/image/tree/s_tree1.png",
		"/asset/image/tree/s_tree2.png",
		"/asset/image/tree/s_tree3.png",
		"/asset/image/tree/s_tree4.png",
		"/asset/image/tree/s_tree4.png",
		"/asset/image/tree/s_tree3.png",
		"/asset/image/tree/s_tree2.png",
		"/asset/image/tree/s_tree1.png",
	}

	summer := []string{
		"/asset/image/tree/tree1.png",
		"/asset/image/tree/tree2.png",
		"/asset/image/tree/tree3.png",
		"/asset/image/tree/tree4.png",
		"/asset/image/tree/tree4.png",
		"/asset/image/tree/tree3.png",
		"/asset/image/tree/tree2.png",
		"/asset/image/tree/tree1.png",
	}
	spring := []string{
		"/asset/image/tree/c_tree1.png",
		"/asset/image/tree/c_tree2.png",
		"/asset/image/tree/c_tree3.png",
		"/asset/image/tree/c_tree4.png",
		"/asset/image/tree/c_tree4.png",
		"/asset/image/tree/c_tree3.png",
		"/asset/image/tree/c_tree2.png",
		"/asset/image/tree/c_tree1.png",
	}
	bmp := ObjectTree{}
	bmp.winter = NewLedImage3DFromPaths(winter)
	bmp.summer = NewLedImage3DFromPaths(summer)
	bmp.spring = NewLedImage3DFromPaths(spring)
	return &bmp
}

func (b *ObjectTree) GetImage3D(param LedCanvasParam) util.ImmutableImage3D {
	switch {
	case param.HasEffect("filter-bk-snows"):
		return b.winter
	case param.HasEffect("filter-bk-sakura"):
		return b.spring
	default:
		return b.summer
	}
}
