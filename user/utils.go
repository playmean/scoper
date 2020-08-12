package user

import (
	"crypto/sha1"
	"encoding/base64"
)

func hashPassword(password string) string {
	hasher := sha1.New()
	hasher.Write([]byte(password))

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
