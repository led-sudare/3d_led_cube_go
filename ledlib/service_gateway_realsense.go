// +build realsense

package ledlib

import (
	"3d_led_cube_go/ledlib/util"
	"time"

	zmq "github.com/zeromq/goczmq"
)

type serviceGatewayRealsense struct {
	sock  *zmq.Sock
	order chan string
	done  chan struct{}
}

const realsenseSharedObjectID = "realsense"

var gServiceGatewayRealsense *serviceGatewayRealsense

func InitSeriveGatewayRealsense(endpoint string) {
	gServiceGatewayRealsense = &serviceGatewayRealsense{}
	var err error
	gServiceGatewayRealsense.sock, err = zmq.NewSub(endpoint, "")
	if err != nil {
		panic(err)
	}
	gServiceGatewayRealsense.sock.Connect(endpoint)
	gServiceGatewayRealsense.order = make(chan string)
	gServiceGatewayRealsense.done = make(chan struct{})

	go serviceGatewayRealsenseWorker(
		gServiceGatewayRealsense.sock,
		gServiceGatewayRealsense.order,
		gServiceGatewayRealsense.done)
}

func serviceGatewayRealsenseWorker(sock *zmq.Sock, c chan string, done chan struct{}) {
	ranges := []uint32{0, 16, 32, 48, 64, 80, 96, 112}
	timer := NewTimer(50 * time.Millisecond)

	defer sock.Destroy()
	for {
		select {
		case <-c:
			done <- struct{}{}
			return
		default:
			sock.RecvFrame() // 読み捨て
			data, _, _ := sock.RecvFrame()

			if !timer.IsPast() {
				continue
			}

			EditSharedLedImage3D(realsenseSharedObjectID,
				func(editable util.Image3D) {
					editable.Clear()

					util.ConcurrentEnum(0, LedHeight, func(y int) {
						for x := 0; x < LedWidth; x++ {
							idx := y*4 + LedHeight*4*x
							c := (uint32(data[idx+0] << 0)) +
								(uint32(data[idx+1]) << 8) +
								(uint32(data[idx+2]) << 16)
							color := util.NewColorFromUint32(c)
							depth := uint32(data[idx+3])

							if depth == 0 {
								continue
							}
							for z := LedDepth - 1; z >= 0; z-- {
								if depth < ranges[z] {
									editable.SetAt(x, y, z, color)
								} else {
									break
								}
							}
						}
					})
				})
		}
	}
}
