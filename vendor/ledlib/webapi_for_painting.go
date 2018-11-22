package ledlib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
	"ledlib/util"
	"net/http"
	"strconv"
)

var paintingSharedObjectID = "painting"

func MakePaintingOrderWithFilters(data []byte) (string, error) {
	var jsonData interface{}
	if err := json.Unmarshal(data, &jsonData); err == nil {
		if j, ok := jsonData.(map[string]interface{}); ok {
			if f, ok := j["filters"].([]interface{}); ok {
				filters := append(f, interface{}(
					map[string]interface{}{
						"id":       "object-painting",
						"lifetime": 0}))

				order := make(map[string][]interface{})
				order["orders"] = filters
				b, _ := json.Marshal(order)
				return string(b), nil
			}
		}
		return "", errors.New("json parse error")
	} else {
		return "", err
	}
}

type LedAllPaintingData struct {
	Led [][]string `json:"led"`
}

type LedPointPaintingData struct {
	X     int    `json:"x"`
	Y     int    `json:"y"`
	Color string `json:"color"`
}

type LedPartOfPaintingData struct {
	Points []LedPointPaintingData `json:"points"`
}

func isPartOfUpdate(data []byte) bool {
	var jsonData interface{}
	if err := json.Unmarshal(data, &jsonData); err == nil {
		if j, ok := jsonData.(map[string]interface{}); ok {
			if _, ok := j["points"]; ok {
				return true
			}
		}
	}
	return false

}

func UpdatePartOfPaintingSharedObject(data []byte) error {

	var ledData LedPartOfPaintingData
	if err := json.Unmarshal(data, &ledData); err == nil {

		EditSharedLedImage3D(paintingSharedObjectID,
			func(editable util.Image3D) {

				for _, point := range ledData.Points {
					d := point.Color
					if c, err := strconv.ParseUint(d, 16, 32); err == nil {
						color := util.NewColorFromUint32(uint32(c))
						editable.SetAt(point.X, point.Y, 0, color)
						editable.SetAt(point.X, point.Y, 1, color)
					}

				}
			})
		return nil
	} else {
		return err
	}
	return errors.New("failed to paint")
}

func UpdateAllPaintingSharedObject(data []byte) error {
	var ledData LedAllPaintingData
	imgForLog := image.NewRGBA(image.Rect(0, 0, LedWidth, LedHeight))
	if err := json.Unmarshal(data, &ledData); err == nil {
		EditSharedLedImage3D(paintingSharedObjectID,
			func(editable util.Image3D) {
				for x := 0; x < LedWidth; x++ {
					for y := 0; y < LedHeight; y++ {
						d := ledData.Led[x][y]
						if i, err := strconv.ParseUint(d, 16, 32); err == nil {
							c := util.NewColorFromUint32(uint32(i))
							editable.SetAt(x, y, 0, c)
							editable.SetAt(x, y, 0, c)
							r, g, b := util.ToUint8s(uint32(i))
							imgForLog.Set(x, y, color.RGBA{r, g, b, 0xff})
						}
					}
				}
			})
		util.WriteImageToLog(imgForLog)
		return nil
	} else {
		return err
	}
	return errors.New("failed to paint")
}

func SetUpWebAPIforPainting(renderer LedBlockRenderer) {

	http.Handle("/api/filters", util.NewCORSHandler(
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "OPTIONS":
				return
			case "POST":
				bufbody := new(bytes.Buffer)
				bufbody.ReadFrom(r.Body)
				fmt.Fprintln(w, bufbody.String())

				// json に変換 filter keyのデータを取得する
				// filter keyのデータは配列、後ろにPainting objectを追加
				// orders key に上記配列を追加、文字列化してShow
				if j, e := MakePaintingOrderWithFilters(bufbody.Bytes()); e == nil {
					renderer.Show(j)
				} else {
					http.Error(w, "Invalid Parameter", http.StatusBadRequest)
				}
			default:
				http.Error(w, "Not implemented.", http.StatusNotFound)
			}
		}))
	http.Handle("/api/led", util.NewCORSHandler(
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "OPTIONS":
				return
			case "POST":
				bufbody := new(bytes.Buffer)
				bufbody.ReadFrom(r.Body)
				if isPartOfUpdate(bufbody.Bytes()) {
					UpdatePartOfPaintingSharedObject(bufbody.Bytes())
				} else {
					UpdateAllPaintingSharedObject(bufbody.Bytes())
				}

			default:
				http.Error(w, "Not implemented.", http.StatusNotFound)
			}
		}))
}
