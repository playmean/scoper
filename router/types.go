package router

type response struct {
	OK    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}
