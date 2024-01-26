package util

import (
	"math/rand"
	"strings"
)

// random  integer between  min and  max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Random String gen a random string of length n
func RandomString(n int) string {
	var sb strings.Builder

	alphabet := "abcdefghiklmnopqrstuvwxyz"
	len := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(len)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// random owner name
func RandomOwner() string {
	return RandomString(6)
}

// random money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// random currency
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	return currencies[rand.Intn(len(currencies))]
}
