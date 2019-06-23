package ledlib

import (
	"3d_led_cube_go/ledlib/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateObject1(t *testing.T) {
	util.SetConvertJsonToTestMode()
	data := map[string]interface{}{
		"id": "object-rocket",
	}

	obj, _, _ := CreateObject(data, nil)
	assert.NotNil(t, obj.(*ObjectBitmap))
}
