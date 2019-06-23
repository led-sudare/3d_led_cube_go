// +build !realsense

package ledlib

import "log"

const realsenseSharedObjectID = "realsense"

type serviceGatewayRealsense struct {
}

func InitSeriveGatewayRealsense(endpoint string) {
	log.Println(`[INFO] serivce gateway realsense run as dummy. if you want to build an application that support realsese. You sholud build with "-tags=realsense" opiton`)
}
