package track

import "github.com/playmean/scoper/common"

type reportPacketSourcePosition struct {
	Lineno int `json:"lineno"`
	Colno  int `json:"colno"`
}

type reportPacketSource struct {
	Filename string                     `json:"filename"`
	Position reportPacketSourcePosition `json:"position"`
}

// ReportPacket for track
type ReportPacket struct {
	Message string             `json:"message"`
	Stack   string             `json:"stack"`
	Source  reportPacketSource `json:"source"`
	Tags    map[string]string  `json:"tags"`
}

// LogPacket for track
type LogPacket struct {
	Data interface{}       `json:"data"`
	Tags map[string]string `json:"tags"`
}

type response struct {
	common.Response

	Data map[string]string `json:"data,omitempty"`
}
