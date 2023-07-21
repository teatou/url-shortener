package random

import (
	"math/rand"
	"time"
)

func NewRandomString(length int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	base := []rune("abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"1234567890",
	)

	ans := make([]rune, length)
	for i := 0; i < length; i++ {
		ans[i] = base[rnd.Intn(len(base))]
	}

	return string(ans)
}
