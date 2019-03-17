package string

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	fake "github.com/brianvoe/gofakeit"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// StringWithCharset returns random string with given character set and length
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// String returns random string with given length
func String(length int) string {
	return StringWithCharset(length, charset)
}

//GenerateShortID returns random string that has no space inside. This is safe for create consumer in kong since kong needs no space in username
func GenerateShortID(prefix string) string {
	fake.Seed(time.Now().UnixNano())
	word := fmt.Sprintf("%s-%s", prefix, fake.Color())

	return strings.Replace(word, " ", "-", -1)
}
