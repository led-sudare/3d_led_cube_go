package util

import (
	"image"
	"image/png"
	"log"
	"os"
	"os/user"
	"path"
	"time"
)

func getLogBaseDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(usr.HomeDir, ".3d_led_cube_go", "log")
}

func getLogDirForOrder() string {
	return path.Join(getLogBaseDir(), "order")
}

func getLogDirForPainting() string {
	return path.Join(getLogBaseDir(), "painting")
}

func mkdirs(path string) {
	_ = os.MkdirAll(path, 0777)
}

func getLogFileName(parent string, extension string) string {
	t := time.Now()
	mkdirs(parent)
	return path.Join(parent, t.Format("2006-01-02T15-04-05.000.")+extension)
}

func WriteImageToLog(img image.Image) {
	path := getLogFileName(getLogDirForPainting(), "png")
	if f, err := os.Create(path); err == nil {
		defer f.Close()
		if err := png.Encode(f, img); err != nil {
			panic(err)
		}
	} else {
		log.Println("faild to write log.", err)
	}
}

func WriteOrderToLog(order string) {
	path := getLogFileName(getLogDirForOrder(), "log")

	if f, err := os.Create(path); err == nil {
		defer f.Close()
		f.WriteString(order)
	} else {
		log.Println("faild to write log.", err)
	}
}
