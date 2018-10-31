package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertJsonAllTable(t *testing.T) {

	colors := []string{
		"blue",
		"white",
		"green",
		"yellowgreen",
		"yellow",
		"orange",
		"brown",
		"red",
	}
	widths := []float64{
		2, 3, 4,
	}
	for _, table := range orderAll {
		for _, color := range colors {
			for _, width := range widths {
				data := map[string]interface{}{
					"color": color,
					"width": width,
				}
				v, err := convertJsonWidhTable(data, table)
				assert.Nil(t, err)
				assert.NotNil(t, v)
				assert.NotNil(t, v["id"])

			}
		}
	}
}

func TestconvertJsonWidhTableSuccessCase1(t *testing.T) {

	SetConvertJsonToTestMode()
	data := map[string]interface{}{
		"color": "blue",
		"width": 2,
	}

	target := ConvertJson(data)

	v, ok := target["id"]
	assert.True(t, ok)
	assert.Equal(t, "object-ghost", v)
}

func TestconvertJsonWidhTableSuccessCase2(t *testing.T) {

	SetConvertJsonToTestMode()
	data := map[string]interface{}{
		"color": "blue",
		"width": "2",
	}

	target := ConvertJson(data)

	v, ok := target["id"]
	assert.True(t, ok)
	assert.Equal(t, "object-ghost", v)
}

func TestconvertJsonWidhTableFailCase1(t *testing.T) {

	SetConvertJsonToTestMode()
	data := map[string]interface{}{
		"abc":   "blue",
		"width": 2,
	}

	target := ConvertJson(data)

	v, ok := target["id"]
	assert.True(t, ok)
	assert.Equal(t, "filter-jump", v)
}

func TestconvertJsonWidhTableFailCase2(t *testing.T) {

	SetConvertJsonToTestMode()
	data := map[string]interface{}{
		"color": "blue",
		"width": 100,
	}

	target := ConvertJson(data)

	v, ok := target["id"]
	assert.True(t, ok)
	assert.Equal(t, "filter-jump", v)
}

func TestconvertJsonWidhTableFailCase3(t *testing.T) {

	SetConvertJsonToTestMode()
	data := map[string]interface{}{}

	target := ConvertJson(data)

	v, ok := target["id"]
	assert.True(t, ok)
	assert.Equal(t, "filter-jump", v)
}
