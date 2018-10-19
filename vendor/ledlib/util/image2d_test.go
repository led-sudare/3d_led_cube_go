package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImage2d(t *testing.T) {
	buffer := NewImage2D(16, 32)
	buffer.SetAt(15, 31, NewColorFromUint32(0xffffff))

	actual := buffer.GetAt(15, 31)
	assert.Equal(t, uint32(0xffffff), actual.Uint32())

}
