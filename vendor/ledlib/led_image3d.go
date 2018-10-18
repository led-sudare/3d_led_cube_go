package ledlib

import (
	"image/png"
	"ledlib/util"
	"log"
)

// 内部のバッファサイズ
// マイナス座標の書き込みや、範囲外の書き込みも多少受け付けるため、大きめのバッファサイズにする
const LedInternalWidh = LedWidth * 3
const LedInternalHeight = LedHeight * 3
const LedInternalDepth = LedDepth * 3

const ledCubeOffsetX = LedWidth
const ledCubeOffsetY = LedHeight
const ledCubeOffsetZ = LedDepth

func NewLedData3D() util.Data3D {
	return util.NewData3D(
		LedInternalWidh, LedInternalHeight, LedInternalDepth,
		ledCubeOffsetX, ledCubeOffsetY, ledCubeOffsetZ)
}

func NewLedImage3D() util.Image3D {
	return util.NewImage3D(
		LedInternalWidh, LedInternalHeight, LedInternalDepth,
		ledCubeOffsetX, ledCubeOffsetY, ledCubeOffsetZ)
}

func NewLedImage3DFromPaths(paths []string) util.Image3D {
	image := NewLedImage3D()
	for z, path := range paths {
		if path == "" {
			continue
		}
		reader, err := util.Assets.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer reader.Close()
		m, err := png.Decode(reader)
		if err != nil {
			log.Fatal(err)
		}
		width, height := m.Bounds().Dx(), m.Bounds().Dy()
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				if m == nil {
					continue
				}
				image.SetAt(x, y, z, util.NewColorFromColor(m.At(x, y)))
			}
		}
	}
	return image
}
