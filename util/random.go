package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	source := rand.NewSource(int64(time.Now().UnixNano()))
	rand.New(source)
}

func RandomName() string {
	var sb strings.Builder
	abjad := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < 5; i++ {
		sb.WriteByte(abjad[rand.Intn(len(abjad))])
	}
	return sb.String()
}

func RandomSaldo() int64 {
	min := int64(10000)
	max := int64(1000000)
	return RandomInt(min, max)
}

func RandomInt(min, max int64) int64 {

	return min + (rand.Int63n(max - min + 1))
}
