package track

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"
)

func marshal(body interface{}) string {
	res, _ := json.Marshal(&body)

	return string(res)
}

func hashID(number uint) string {
	hasher := sha1.New()
	hasher.Write([]byte(strconv.FormatUint(uint64(number), 10)))

	hash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return strings.ReplaceAll(hash, "=", "")
}
