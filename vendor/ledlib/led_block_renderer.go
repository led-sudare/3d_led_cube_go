package ledlib

import (
	"encoding/json"
	"errors"
	"ledlib/servicegateway"
	"ledlib/util"
	"log"
	"math"
	"time"
)

type LedBlockRenderer interface {
	Abort()
	Start()
	Terminate()
	Show(blocks string)
}

func NewLedBlockRenderer() LedBlockRenderer {
	return &ledBockRendererImpl{
		make(chan string),
		make(chan struct{})}
}

func GetJSONValue(m interface{}, key string) (interface{}, error) {
	value, _ := GetJSONValueOrDefault(m, key, nil)
	if value == nil {
		return nil, errors.New("invalid json format")
	} else {
		return value, nil
	}
}

func GetJSONValueOrDefault(m interface{}, key string, defaults interface{}) (interface{}, error) {
	if mm, ok := m.(map[string]interface{}); ok {
		if val, ok := mm[key]; ok {
			return val, nil
		}
	}
	return defaults, errors.New("invalid json format")
}

func getOrdersFromJson(rawJson string) ([]interface{}, error) {
	var ordersMap interface{}
	err := json.Unmarshal([]byte(rawJson), &ordersMap)
	if err != nil {
		return nil, err
	}

	if val, ok := ordersMap.(map[string]interface{}); ok {
		if val, ok := val["orders"]; ok {
			if val, ok := val.([]interface{}); ok {

				orders := make([]interface{}, 0)
				for _, v := range val {
					if order, ok := v.(map[string]interface{}); ok {
						if _, err := GetJSONValue(order, "color"); err == nil {
							order = util.ConvertJson(order)
						}

						orders = append(orders, order)
					}
				}

				return orders, nil
			}
		}
	}
	return nil, errors.New("invalid json format")
}

func getOrdersInLoop(orders []interface{}, start int) ([]interface{}, error) {
	ordersInLoop := make([]interface{}, 0)
	for i := start; i < len(orders); i++ {
		order := orders[i]
		mapOrder := order.(map[string]interface{})
		if val, ok := mapOrder["id"]; ok {
			if val.(string) == "ctrl-loop" {
				return ordersInLoop, nil
			}
			ordersInLoop = append(ordersInLoop, orders[i])
		} else {
			return nil, errors.New("invalid json format")
		}

	}
	return ordersInLoop, nil
}

func expands(orders []interface{}, count int) []interface{} {
	newOrders := make([]interface{}, 0)
	for i := 0; i < count; i++ {
		newOrders = append(newOrders, orders...)
	}
	return newOrders
}

func flattenOrders(orders []interface{}) ([]interface{}, error) {
	flatten := make([]interface{}, 0)

	for i := 0; i < len(orders); i++ {
		if val, err := GetJSONValueOrDefault(orders[i], "id", nil); err == nil {
			if val.(string) == "ctrl-loop" {
				ordersInLoop, err := getOrdersInLoop(orders, i+1)
				if err != nil {
					return nil, errors.New("invalid order format")
				}
				count, _ := GetJSONValueOrDefault(orders[i], "count", 3)
				flatten = append(flatten, expands(ordersInLoop, count.(int))...)
				i += len(ordersInLoop) + 1
			} else {
				flatten = append(flatten, orders[i])
			}
		} else {
			return nil, errors.New("invalid order array. key: id not found")
		}
	}
	return flatten, nil
}

type ledBockRendererImpl struct {
	orderCh chan string
	doneCh  chan struct{}
}

func (l *ledBockRendererImpl) Terminate() {
	close(l.orderCh)
	<-l.doneCh
}

func (l *ledBockRendererImpl) Abort() {
	l.orderCh <- `{"orders":[{"id":"object-blank", "lifetime":1}]}`
}

func (l *ledBockRendererImpl) Show(blocks string) {
	l.orderCh <- blocks
}

func (l *ledBockRendererImpl) getOrdersFromString(t string) []interface{} {
	if arrayOrders, err := getOrdersFromJson(t); err != nil {
		return nil
	} else if flattenOrders, err := flattenOrders(arrayOrders); err != nil {
		return nil
	} else {
		return flattenOrders
	}
}

func (l *ledBockRendererImpl) waitOrdersFromChanel() ([]interface{}, error) {
	// === get orders ===>
	t, ok := <-l.orderCh

	if !ok {
		return nil, errors.New("termination command recevied")
	}
	return l.getOrdersFromString(t), nil
}

func (l *ledBockRendererImpl) getOrdersFromChanel() ([]interface{}, error) {

	select {
	case t, ok := <-l.orderCh:
		if !ok {
			return nil, errors.New("termination command recevied")
		}
		return l.getOrdersFromString(t), nil
	default:
		return nil, nil
	}
}

func (l *ledBockRendererImpl) Start() {
	go func() {
		defer func() { close(l.doneCh) }()

		for {
			var filters LedCanvas = NewLedCanvas()
			param := NewLedCanvasParam()
			var lifetime float64 = 1
			var expiresDate int64

			orders, err := l.waitOrdersFromChanel()
			if err != nil {
				log.Println("terminated")
				return
			}
			if orders == nil {
				// order error
				continue
			}

			var object LedObject
			for {
				start := time.Now()
				newOrders, err := l.getOrdersFromChanel()
				if err != nil {
					log.Println("terminated")
					return
				}
				if newOrders != nil {
					servicegateway.GetAudigoSeriveGateway().Stop()
					filters = NewLedCanvas()
					param = NewLedCanvasParam()
					orders = newOrders
					lifetime = 1
					expiresDate = 0
				}

				// update filters
				if lifetime != 0 &&
					time.Now().Unix() > expiresDate { // if lifetime expired
					object, filters, lifetime, orders, param, err = GetFilterAndObject(orders, filters, param)
					if err != nil {
						servicegateway.GetAudigoSeriveGateway().Stop()
						break
					} else {
						expiresDate = time.Now().Unix() + int64(lifetime)
					}
				}

				ShowObject(filters, object, param)
				duration := time.Now().Sub(start)
				waittime := math.Max(0, float64(50*time.Millisecond-duration))

				/*
					log.Printf("rendering: %0.2f, wait: %0.2f\n",
						float64(duration)/float64(time.Millisecond),
						float64(waittime)/float64(time.Millisecond))
				*/

				time.Sleep(time.Duration(waittime))

			}

		}

	}()

}

func GetFilterAndObject(iOrders []interface{}, canvas LedCanvas, param LedCanvasParam) (LedObject, LedCanvas, float64, []interface{}, LedCanvasParam, error) {

	filter := canvas
	var object LedObject
	jsonOrders := iOrders
	for {
		if len(jsonOrders) == 0 {
			// invalid order
			return nil, nil, 0, iOrders, param, errors.New("invalid order format")
		}
		rawOrder := jsonOrders[0]
		jsonOrders = jsonOrders[1:]

		if jsonOrder, ok := rawOrder.(map[string]interface{}); ok {
			order, lifetime, err := CreateObject(jsonOrder, filter)
			if err != nil {
				return nil, nil, 0, iOrders, param, err
			}

			switch v := order.(type) {
			case LedObject:
				object = v
				return object, filter, lifetime, jsonOrders, param, nil
			case LedCanvas:
				value, _ := GetJSONValue(jsonOrder, "id")
				param.AppendsEffect(value.(string))
				filter = v
			default:
				return nil, nil, 0, iOrders, param, errors.New("invalid order format")
			}
		}
	}
}
