package ledlib

import (
	"ledlib/util"

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

	defer sock.Destroy()
	for {
		select {
		case <-c:
			done <- struct{}{}
			return
		default:
			sock.RecvFrame() // 読み捨て
			data, _, _ := sock.RecvFrame()

			ranges := []uint32{16, 32, 48, 64, 80, 96, 112, 128}
			img := GetSharedLedImage3D(realsenseSharedObjectID)
			img.Clear()

			util.ConcurrentEnum(0, LedHeight, func(y int) {
				for x := 0; x < LedWidth; x++ {
					idx := y*4 + LedHeight*4*x
					c := (uint32(data[idx]) << 16) +
						(uint32(data[idx+1]) << 8) +
						(uint32(data[idx+2]))
					color := util.NewColorFromUint32(c)
					depth := uint32(data[idx+3])

					for z := 0; z < LedDepth; z++ {
						if 1 < depth && depth < ranges[z] {
							img.SetAt(x, y, z, color)
						}
					}

				}
			})
		}
	}
}
