package generator

import (
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(newSeed()))

var charset = "qwertyuiopasdfghjklzxcvbnm1234567890-_!@#4%"

func newSeed() int64 {
	return time.Now().UnixNano()
}

// Password string
func Password(length uint) string {
	random.Seed(newSeed())

	b := make([]byte, length)

	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}

	return string(b)
}
