package controllers

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"strconv"
)

func marshal(body interface{}) string {
	res, _ := json.Marshal(&body)

	return string(res)
}

func hashID(number uint) string {
	hasher := sha1.New()
	hasher.Write([]byte(strconv.FormatUint(uint64(number), 10)))

	return hex.EncodeToString(hasher.Sum(nil))
}
