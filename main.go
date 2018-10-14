package main

import (
	"bufio"
	"flag"
	"fmt"
	"ledlib"
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

	ledlib.SetUpWebAPIforCommon(renderer)
	ledlib.SetUpWebAPIforPainting(renderer)

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
