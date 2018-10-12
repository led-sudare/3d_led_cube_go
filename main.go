package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"ledlib"
	"ledlib/webapi"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"strings"
	"time"
)

func getUnixNano() int64 {
	return time.Now().UnixNano()
}

type Status struct {
	Enable bool   `json:"enable"`
	Target string `json:"target"`
}

func main() {
	var (
		optDestination = flag.String("d", "localhost:9001", "Specify IP and port of Led Cube. if opt is not set, launch simulator.")
	)
	flag.Parse()
	if *optDestination == "" {
		runtime.LockOSThread()
		ledlib.GetLed().EnableSimulator(true)
	} else {
		ledlib.GetLed().EnableSimulator(false)
		ledlib.GetLed().SetUrl(*optDestination)
	}
	fmt.Println("udp target: " + *optDestination)
	go func() {
		//		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	/*
		setup renderer
	*/
	renderer := ledlib.NewLedBlockRenderer()
	renderer.Start()

	// start http server
	// endpoins are below
	// POST /api/show       content json
	// POST /api/abort		no content
	// POST /api/target     text content
	//
	http.HandleFunc("/api/show", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			bufbody := new(bytes.Buffer)
			bufbody.ReadFrom(r.Body)
			fmt.Fprintln(w, bufbody.String())
			renderer.Show(bufbody.String())
		default:
			http.Error(w, "Not implemented.", http.StatusNotFound)
		}
	})
	http.HandleFunc("/api/abort", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			fmt.Fprintln(w, "abort")
			renderer.Abort()
		default:
			http.Error(w, "Not implemented.", http.StatusNotFound)
		}
	})
	http.HandleFunc("/api/config", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			bufbody := new(bytes.Buffer)
			bufbody.ReadFrom(r.Body)
			config, err := webapi.UnmarshalConfigration(bufbody.Bytes())
			if err != nil {
				http.Error(w, "Invalid json body.", http.StatusNotFound)
			} else {
				ledlib.GetLed().Enable(config.Enable)
			}

		default:
			http.Error(w, "Not implemented.", http.StatusNotFound)
		}
	})
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "Hello")
		default:
			http.Error(w, "Not implemented.", http.StatusNotFound)
		}
	})
	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			status := Status{ledlib.GetLed().IsEnable(), *optDestination}
			jsoBytes, _ := json.Marshal(status)
			w.Write(jsoBytes)
		default:
			http.Error(w, "Not implemented.", http.StatusNotFound)
		}
	})
	fmt.Println("led framework is running ...  on port 8081")
	go func() {
		log.Fatal(http.ListenAndServe(":8081", nil))
	}()

	for {
		sc := bufio.NewScanner(os.Stdin)
		fmt.Print(">>")
		if sc.Scan() {
			input := sc.Text()
			fmt.Println("input:" + input)
			switch {
			case strings.HasPrefix(input, "show"):
				renderer.Show(strings.Replace(input, "show:", "", 1))
			case strings.HasPrefix(input, "abort"):
				renderer.Abort()
			}
		}
	}

	renderer.Terminate()

}
