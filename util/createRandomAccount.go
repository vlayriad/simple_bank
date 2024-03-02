package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijqlmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(5)
}

func RandomBalance() int64 {
	return RandomInt(1, 1000)
}

func RandomCurrency() string {
	currencies := []string{"BRL", "USD", "EUR"}

	a := len(currencies)
	return currencies[rand.Intn(a)]
}

func RandomEmail() string {

	return fmt.Sprintf("%s@email.com", RandomString(6))
}
