package random

import (
	"math/rand"
	"time"
)

func NewRandomString(size int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	available := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

	res := make([]rune, size)
	for i := range res {
		res[i] = available[r.Intn(len(available))]
	}

	return string(res)
}
