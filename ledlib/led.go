package ledlib

/*
#cgo LDFLAGS: -lledlib
#include "./../../lib/led.h"
*/
/* ledlib  */
//import "C"
import (
	"log"
	"net"

	zmq "github.com/zeromq/goczmq"
)

const LedWidth = 16
const LedHeight = 32
const LedDepth = 8

const ledColor = 3
const ledRed = 0
const ledGreen = 1
const ledBulue = 2

type Led interface {
	SetUrl(url string)
	GetUrl() string
	SetLed(x, y, z int, rgb uint32)
	Clear()
	Show()
	Enable(enable bool)
	IsEnable() bool
	EnableSimulator(enable bool)
	//	SetPort(port uint16)
}

var sharedLedInstance Led

func GetLed() Led {
	if sharedLedInstance == nil {
		sharedLedInstance = newLed()
	}
	return sharedLedInstance
}

func newLed() *ledImpl {
	goImpl := newGoLed()
	/* ledlib */
	//	cImpl := newCLed()
	//return &ledImpl{goImpl, cImpl, goImpl, true, false}
	return &ledImpl{goImpl, goImpl, true, false}
}

/*
* ledImpl
 */
type ledImpl struct {
	goImpl *ledGoImpl
	/* ledlib */
	//	cImpl           *ledCImpl
	currentImpl     Led
	enable          bool
	enableSimulator bool
}

func (led *ledImpl) SetUrl(url string) {
	led.currentImpl.SetUrl(url)
}

func (led *ledImpl) GetUrl() string {
	return led.currentImpl.GetUrl()
}

func (led *ledImpl) SetLed(x, y, z int, rgb uint32) {
	led.currentImpl.SetLed(x, y, z, rgb)
}

func (led *ledImpl) Clear() {
	led.currentImpl.Clear()
}

func (led *ledImpl) Show() {
	if led.enable {
		led.currentImpl.Show()
	}
}

func (led *ledImpl) Enable(enable bool) {
	led.enable = enable
}
func (led *ledImpl) IsEnable() bool {
	return led.enable
}

func (led *ledImpl) EnableSimulator(enable bool) {
	/* -- ledlib --- */
	/*
		if enable {
			led.currentImpl = led.cImpl
		} else {
			led.currentImpl = led.goImpl
		}
		C.EnableSimulator(C.bool(enable))
	*/
}

/*
* Go Implimentation
 */
type ledGoImpl struct {
	ledUrl       string
	led565Buffer []byte
	sem          chan struct{}
	urlToIPmap   map[string]*net.UDPAddr
	zmqPubSock   *zmq.Sock
}

func newGoLed() *ledGoImpl {
	led := ledGoImpl{}
	led.led565Buffer = make([]byte, LedWidth*LedHeight*LedDepth*2)
	led.urlToIPmap = make(map[string]*net.UDPAddr)

	led.sem = make(chan struct{}, 1)
	return &led
}

func (led *ledGoImpl) SetUrl(url string) {
	led.ledUrl = url
	endpoint := "tcp://" + url
	led.zmqPubSock = zmq.NewSock(zmq.Pub)
	err := led.zmqPubSock.Connect(endpoint)
	if err != nil {
		panic(err)
	}

}
func (led *ledGoImpl) GetUrl() string {
	return led.ledUrl
}

func (led *ledGoImpl) SetLed(x, y, z int, rgb uint32) {
	if x < 0 || LedWidth <= x {
		log.Printf("invalid x : %d\n", x)
		return
	}
	if y < 0 || LedHeight <= y {
		log.Printf("invalid y : %d\n", y)
		return
	}
	if z < 0 || LedDepth <= z {
		log.Printf("invalid z : %d\n", z)
		return
	}

	r, g, b := byte(rgb>>16), byte(rgb>>8), byte(rgb>>0)

	index565 := z*2 + y*LedDepth*2 + x*LedHeight*LedDepth*2
	led.led565Buffer[index565+0] = r&0xF8 + g>>5
	led.led565Buffer[index565+1] = (g<<2)&0xe0 + b>>3
}

func (led *ledGoImpl) Clear() {
	for i, _ := range led.led565Buffer {
		led.led565Buffer[i] = 0
	}
}

func (led *ledGoImpl) Show() {
	tcpAddr, err := net.ResolveUDPAddr("udp", led.getUrl())
	if err != nil {
		if lastResolvedAddr, ok := led.urlToIPmap[led.getUrl()]; !ok {
			log.Printf("error: %s", err.Error())
			return
		} else {
			log.Println("cannot resolve ip address from hostname. use ip last connect.")
			tcpAddr = lastResolvedAddr
		}

	}
	led.urlToIPmap[led.getUrl()] = tcpAddr

	// ZMQ Endpoint
	if led.zmqPubSock != nil {
		led.zmqPubSock.SendFrame(led.led565Buffer, zmq.FlagNone)
	} else {
		log.Println("Warning.. zmqPubSock is not initialized.")
	}
}

func (led *ledGoImpl) Enable(enable bool) {
	// do nothing.
}

func (led *ledGoImpl) EnableSimulator(enable bool) {
	// do nothing.
}
func (led *ledGoImpl) IsEnable() bool {
	return true
}

func (led *ledGoImpl) getUrl() string {
	return led.ledUrl
}

/* -- ledlib --- */
/*
	type ledCImpl struct {
}

func newCLed() *ledCImpl {
	return &ledCImpl{}
}

func (led *ledCImpl) SetUrl(url string) {

	ipAndPort := strings.Split(url, ":")
	switch {
	case len(ipAndPort) == 2:
		C.SetUrl(C.CString(ipAndPort[0]))
		port, e := strconv.ParseInt(ipAndPort[1], 10, 16)
		if e != nil {
			log.Printf("invalid port number. %s\n", ipAndPort[1])
			return
		}
		C.SetPort(C.ushort(port))
	case len(ipAndPort) == 1:
		C.SetUrl(C.CString(ipAndPort[0]))
	case len(ipAndPort) == 0:
		log.Printf("invalid url %s\n", url)
		return
	}

}

func (led *ledCImpl) SetLed(x, y, z int, rgb uint32) {
	C.SetLed(C.int(x), C.int(y), C.int(z), C.int(rgb))
}

func (led *ledCImpl) Clear() {
	C.Clear()
}

func (led *ledCImpl) Show() {
	C.Show()
}

func (led *ledCImpl) Enable(enable bool) {
}
func (led *ledCImpl) IsEnable() bool {
	return true
}

func (led *ledCImpl) EnableSimulator(enable bool) {
	C.EnableSimulator(C.bool(enable))
}
*/
