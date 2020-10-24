package controllers

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/playmean/scoper/common"
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

func makeToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString(common.SigningKey)

	if err != nil {
		return "", err
	}

	return t, nil
}
