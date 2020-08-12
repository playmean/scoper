package track

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
}

type response struct {
	OK    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}
