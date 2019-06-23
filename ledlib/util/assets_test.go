package util

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testAssetsFileExist(t *testing.T, path string) {
	var f http.File
	var e error
	f, e = Assets.Open(path)

	assert.NotNil(t, f, "path: "+path)
	assert.Nil(t, e, "path: "+path)
}

func TestAssets(t *testing.T) {

	datas := []string{
		"/asset/image/mountain/mountain1.png",
		"/asset/image/mountain/mountain2.png",
		"/asset/image/rocket/rocket1.png",
		"/asset/image/rocket/rocket2.png",
		"/asset/image/rocket/rocket3.png",
		"/asset/image/rocket/rocket4.png",
	}
	for _, d := range datas {
		testAssetsFileExist(t, d)
	}
}
