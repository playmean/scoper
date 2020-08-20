package track

import "git.playmean.xyz/playmean/error-tracking/common"

type reportPacketSourcePosition struct {
	Lineno int `json:"lineno"`
	Colno  int `json:"colno"`
}

type reportPacketSource struct {
	Filename string                     `json:"filename"`
	Position reportPacketSourcePosition `json:"position"`
}

type reportPacket struct {
	Message string             `json:"message"`
	Stack   string             `json:"stack"`
	Source  reportPacketSource `json:"source"`
	Tags    interface{}        `json:"tags"`
}

type logPacket struct {
	Data interface{} `json:"data"`
	Tags interface{} `json:"tags"`
}

type response struct {
	common.Response

	Data map[string]string `json:"data,omitempty"`
}
