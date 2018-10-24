// +build !realsense

package ledlib

import "fmt"

const realsenseSharedObjectID = "realsense"

type serviceGatewayRealsense struct {
}

func InitSeriveGatewayRealsense(endpoint string) {
	fmt.Println(`[INFO] serivce gateway realsense run as dummy. if you want to build an application that support realsese. You sholud build with "-tags=realsense" opiton`)
}
