package controllers

type respLogin struct {
	Token string `json:"token"`
}

type respUserInfo struct {
	ID uint `json:"id"`

	Username string `json:"username"`
	FullName string `json:"fullname"`
	Role     string `json:"role"`

	CreatedAt int64 `json:"created_at"`
}

type respEnvironment struct {
	Environment string `json:"environment"`
	Count       uint   `json:"count"`
}

type respTagName struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Count uint   `json:"count"`
}

type respTagValue struct {
	Value string `json:"value"`
	Count uint   `json:"count"`
}

type respTag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type respTrack struct {
	Type        string `json:"type,omitempty"`
	Environment string `json:"environment"`

	Message  string `json:"message,omitempty"`
	Stack    string `json:"stack,omitempty"`
	Filename string `json:"filename,omitempty"`
	Lineno   int    `json:"lineno,omitempty"`
	Colno    int    `json:"colno,omitempty"`

	Meta map[string]string `json:"meta,omitempty"`

	Tags []respTag `json:"tags,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`
}

type respTracks map[string][]respTrack
