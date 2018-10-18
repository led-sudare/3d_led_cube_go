package servicegateway

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMarshalAudigoOrder(t *testing.T) {

	target := &audigoOrder{}
	target.Function = "stop"

	json, err := json.Marshal(target)
	assert.Nil(t, err)
	assert.NotNil(t, json)
	t.Log(string(json))
}

func TestPlaySound(t *testing.T) {
	InitAudigoSeriveGateway("http://localhost:8082", "testtest")
	player := GetAudigoSeriveGateway()
	defer player.Terminate()

	player.Play("bgm_wave.wav", false, false)
	time.Sleep(3 * time.Second)
}
