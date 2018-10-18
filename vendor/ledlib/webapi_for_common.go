package ledlib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ledlib/servicegateway"
	"ledlib/util"
	"ledlib/webapi"
	"net/http"
)

type Status struct {
	Enable bool   `json:"enable"`
	Target string `json:"target"`
}

func SetUpWebAPIforCommon(renderer LedBlockRenderer) {

	// start http server
	// endpoins are below
	// POST /api/show       content json
	// POST /api/abort		no content
	// POST /api/target     text content

	http.Handle("/api/show", util.NewCORSHandler(
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "POST":
				bufbody := new(bytes.Buffer)
				bufbody.ReadFrom(r.Body)
				fmt.Fprintln(w, bufbody.String())
				renderer.Show(bufbody.String())
			default:
				http.Error(w, "Not implemented.", http.StatusNotFound)
			}
		}))
	http.Handle("/api/abort", util.NewCORSHandler(
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "POST":
				fmt.Fprintln(w, "abort")
				renderer.Abort()
			default:
				http.Error(w, "Not implemented.", http.StatusNotFound)
			}
		}))
	http.Handle("/api/config", util.NewCORSHandler(
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "POST":
				bufbody := new(bytes.Buffer)
				bufbody.ReadFrom(r.Body)
				config, err := webapi.UnmarshalConfigration(bufbody.Bytes())
				if err != nil {
					http.Error(w, "Invalid json body.", http.StatusNotFound)
				} else {
					GetLed().Enable(config.Enable)
					if config.Enable {
						servicegateway.GetAudigoSeriveGateway().SetVolume(1.0)
					} else {
						servicegateway.GetAudigoSeriveGateway().SetVolume(0)
					}

				}

			default:
				http.Error(w, "Not implemented.", http.StatusNotFound)
			}
		}))
	http.Handle("/api/hello", util.NewCORSHandler(
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				fmt.Fprintf(w, "Hello")
			default:
				http.Error(w, "Not implemented.", http.StatusNotFound)
			}
		}))
	http.Handle("/api/status", util.NewCORSHandler(
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				status := Status{GetLed().IsEnable(), GetLed().GetUrl()}
				jsoBytes, _ := json.Marshal(status)
				w.Write(jsoBytes)
			default:
				http.Error(w, "Not implemented.", http.StatusNotFound)
			}
		}))
}
