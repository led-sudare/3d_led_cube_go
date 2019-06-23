package ledlib

import (
	"image/color"
	"3d_led_cube_go/ledlib/util"
	"testing"
)

func TestColorToUint32(t *testing.T) {

	data := &color.RGBA{0xff, 0xff, 0xff, 0xff}
	result := util.NewColorFromColor(data).Uint32()
	if result != 0xffffff {
		t.Log(data)
		t.Fatalf("failed test result:%d", result)
	}
}
