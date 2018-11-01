package servicegateway

import (
	"encoding/json"
	"ledlib/webapi"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

const (
	audigoApiUriBase = "audio/v1/"

	funcIDPlay                = "play"
	funcIDPause               = "pause"
	funcIDVolume              = "volume"
	funcIDStop                = "stop"
	funcIDResume              = "resume"
	funcIDServiceGatewayAbort = "abort"
)

type AudigoServiceGateway interface {
	Play(src string, loop bool, stop bool)
	Pause()
	Resume()
	Stop()
	SetVolume(volume float64)
	Terminate()
}

type AudigoServiceGatewayImpl struct {
	url       string
	contentID string
	order     chan *audigoOrder
	done      chan struct{}
}

type audigoOrderData struct {
	Src  string  `json:"src"`
	Loop bool    `json:"loop"`
	Stop bool    `json:"stop"`
	Vol  float64 `json:"vol"`
}

func NewAudigoOrderData() *audigoOrderData {
	o := &audigoOrderData{}
	o.Vol = 1.0
	o.Loop = false
	o.Stop = false
	return o
}

type audigoOrder struct {
	ContentID string
	Function  string
	Data      *audigoOrderData
}

func NewAudigoOrder(contentID, function string, data *audigoOrderData) *audigoOrder {
	o := &audigoOrder{}
	o.ContentID = contentID
	o.Function = function

	if data == nil {
		o.Data = NewAudigoOrderData()
	} else {
		o.Data = data
	}
	return o
}

func (a *audigoOrder) GetRestUri() string {
	return audigoApiUriBase + a.Function + "/" + a.ContentID
}

func (s *AudigoServiceGatewayImpl) newAudigoPlayOrder(src string, loop bool, stop bool) *audigoOrder {
	data := NewAudigoOrderData()
	data.Src = src
	data.Loop = loop
	data.Stop = stop

	return NewAudigoOrder(s.contentID, funcIDPlay, data)
}

func (s *AudigoServiceGatewayImpl) newAudigoPauseOrder() *audigoOrder {
	return NewAudigoOrder(s.contentID, funcIDPause, nil)
}

func (s *AudigoServiceGatewayImpl) newAudigoResumeOrder() *audigoOrder {
	return NewAudigoOrder(s.contentID, funcIDResume, nil)
}

func (s *AudigoServiceGatewayImpl) newAudigoStopOrder() *audigoOrder {
	return NewAudigoOrder(s.contentID, funcIDStop, nil)
}

func (s *AudigoServiceGatewayImpl) newAudigoVolumeOrder(volume float64) *audigoOrder {
	data := NewAudigoOrderData()
	data.Vol = volume

	return NewAudigoOrder(s.contentID, funcIDVolume, data)
}
func (s *AudigoServiceGatewayImpl) newAudigoServiceGatewayAbortOrder() *audigoOrder {
	return NewAudigoOrder(s.contentID, funcIDServiceGatewayAbort, nil)
}

var instance AudigoServiceGateway
var once sync.Once

func audigoServiceGatewayWorker(url string, c chan *audigoOrder, done chan struct{}) {

	for {
		order := <-c
		switch order.Function {
		case funcIDPlay:
			fallthrough
		case funcIDPause:
			fallthrough
		case funcIDResume:
			fallthrough
		case funcIDStop:
			fallthrough
		case funcIDVolume:
			if data, e := json.Marshal(order.Data); e == nil {
				log.Println(string(data))
				if e := webapi.HttpJsonPost(url+order.GetRestUri(), data); e != nil {
					log.Println(e)
				}
			}
		case funcIDServiceGatewayAbort:
			fallthrough
		default:
			// error
			done <- struct{}{}
		}

	}
}

func InitAudigoSeriveGateway(url string, contentID string) {
	rand.Seed(time.Now().UnixNano())
	impl := &AudigoServiceGatewayImpl{}

	impl.url = url + "/"
	if contentID == "" {
		impl.contentID = strconv.Itoa(rand.Int())
	} else {
		impl.contentID = contentID
	}
	impl.order = make(chan *audigoOrder)
	impl.done = make(chan struct{})

	instance = impl
	go audigoServiceGatewayWorker(impl.url, impl.order, impl.done)
}

func GetAudigoSeriveGateway() AudigoServiceGateway {
	return instance
}

func (s *AudigoServiceGatewayImpl) Play(
	src string, loop bool, stop bool) {
	s.order <- s.newAudigoPlayOrder(src, loop, stop)
}

func (s *AudigoServiceGatewayImpl) Pause() {
	s.order <- s.newAudigoPauseOrder()
}

func (s *AudigoServiceGatewayImpl) Resume() {
	s.order <- s.newAudigoResumeOrder()
}

func (s *AudigoServiceGatewayImpl) Stop() {
	s.order <- s.newAudigoStopOrder()
}

func (s *AudigoServiceGatewayImpl) SetVolume(volume float64) {
	s.order <- s.newAudigoVolumeOrder(volume)
}

func (s *AudigoServiceGatewayImpl) Terminate() {
	s.order <- s.newAudigoServiceGatewayAbortOrder()
	<-s.done
}
