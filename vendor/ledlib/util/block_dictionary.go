package util

import (
	"errors"
	"strconv"
)

type BlockSizeTable struct {
	Width int `json:"width"`
	Id    int `json:"id"`
}

// 2017.12 gadget fes

type widthToId map[int]string
type colorToTable map[string]widthToId

var orderTest = colorToTable{
	"blue":        {2: "object-ghost", 3: "filter-swaying", 4: "filter-bk-wave"},
	"green":       {2: "object-tree", 3: "filter-zoom", 4: "filter-bk-mountain"},
	"orange":      {2: "object-note", 3: "filter-rainbow", 4: "filter-bk-fireworks"},
	"white":       {2: "object-snowman", 3: "filter-rolldown", 4: "filter-bk-snows"},
	"red":         {2: "object-heart", 3: "filter-jump", 4: "filter-bk-sakura"},
	"brown":       {2: "object-socks", 3: "filter-spiral2", 4: "filter-bk-cloud"},
	"yellowgreen": {2: "object-yacht", 3: "filter-wakame", 4: "filter-bk-grass"},
	"yellow":      {2: "object-star", 3: "filter-skewed", 4: "filter-bk-stars"},
}

var order201712GadgetFes = colorToTable{
	"blue":        {2: "object-ghost", 3: "filter-swaying", 4: "filter-bk-wave"},
	"green":       {2: "object-tree", 3: "filter-zoom", 4: "filter-bk-mountain"},
	"orange":      {2: "object-note", 3: "filter-rainbow", 4: "filter-bk-fireworks"},
	"white":       {2: "object-snowman", 3: "filter-rolldown", 4: "filter-bk-snows"},
	"red":         {2: "object-heart", 3: "filter-jump", 4: "filter-bk-sakura"},
	"brown":       {2: "object-socks", 3: "filter-spiral2", 4: "filter-bk-cloud"},
	"yellowgreen": {2: "object-yacht", 3: "filter-wakame", 4: "filter-bk-grass"},
	"yellow":      {2: "object-star", 3: "filter-skewed", 4: "filter-bk-stars"},
}

// 2018.08 Maker Faire Tokyo
var order201808MakerFaireTokyo = colorToTable{
	"blue":        {2: "object-rocket", 3: "filter-swaying-ctrl", 4: "filter-bk-wave"},
	"white":       {2: "object-ghost", 3: "filter-rolling-ctrl", 4: "filter-bk-snows"},
	"green":       {2: "object-saboten", 3: "filter-jump-ctrl", 4: "filter-bk-mountain"},
	"yellowgreen": {2: "object-yacht", 3: "filter-wakame-ctrl", 4: "filter-bk-grass"},
	"yellow":      {2: "object-star", 3: "filter-zy-skewed-ctrl", 4: "filter-bk-stars"},
	"orange":      {2: "object-note", 3: "filter-rainbow-ctrl", 4: "filter-bk-rains"},
	"brown":       {2: "object-stickman", 3: "filter-zoom-ctrl", 4: "filter-bk-cloud"},
	"red":         {2: "object-tulip", 3: "filter-3d-explosion-ctrl", 4: "filter-bk-fireworks"},
}

// 2018.11&12 Yokohama Gadget Festival, Yahoo Japan Hack Day
var order201811YgfYhd = colorToTable{
	"blue":        {2: "object-rocket", 3: "filter-swaying-ctrl", 4: "filter-bk-wave"},
	"white":       {2: "object-ghost", 3: "filter-rolling-ctrl", 4: "filter-bk-snows"},
	"green":       {2: "object-saboten", 3: "filter-jump-ctrl", 4: "filter-bk-mountain"},
	"yellowgreen": {2: "object-yacht", 3: "filter-wakame-ctrl", 4: "filter-bk-grass"},
	"yellow":      {2: "object-star", 3: "filter-zy-skewed-ctrl", 4: "filter-bk-stars"},
	"orange":      {2: "object-note", 3: "filter-rainbow-ctrl", 4: "filter-bk-rains"},
	"brown":       {2: "object-stickman", 3: "filter-zoom-ctrl", 4: "filter-bk-cloud"},
	"red":         {2: "object-tulip", 3: "filter-3d-explosion-ctrl", 4: "filter-bk-fireworks"},
}

// this array is used for test to check whether all tables are correct.
var orderAll = []colorToTable{
	order201712GadgetFes,
	order201808MakerFaireTokyo,
	order201811YgfYhd,
}

var orderDefault colorToTable = order201811YgfYhd

func SetConvertJsonToTestMode() {
	orderDefault = orderTest
}

func convertJsonWidhTable(order map[string]interface{}, table colorToTable) (map[string]interface{}, error) {

	color, _ := order["color"].(string)
	var width int
	if widths, ok := order["width"].(string); !ok {
		width = int(order["width"].(float64))

	} else {
		width, _ = strconv.Atoi(widths)
	}
	if id, ok := table[color][width]; ok {
		order["id"] = id
		return order, nil
	} else {
		return nil, errors.New("unkown color or width")
	}
}

func ConvertJson(order map[string]interface{}) map[string]interface{} {
	if result, err := convertJsonWidhTable(order, orderDefault); err != nil {
		order["id"] = "filter-jump-ctrl"
		return order
	} else {
		return result
	}
}
