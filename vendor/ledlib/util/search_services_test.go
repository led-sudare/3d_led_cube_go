package util

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAvahi(t *testing.T) {
	addr, err := net.LookupIP("pi-content-selector.local")
	if err != nil {
		t.Fail()
	} else {
		for _, ip := range addr {
			t.Log(ip.String())
		}
		assert.Equal(t, "192.168.1.19", addr[1].String())
	}
}
