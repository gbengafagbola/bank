package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init(){
	rand.Seed(time.Now().UnixNano()) // rand.Seed expect an int64 input so we have to convert time.Now with UnixNano before passing to the function
}

//  RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64{
	return min + rand.Int63n(max - min + 1) // this would return a random integer between 0 and (max - min)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

//RandomOwner generate a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomBalance generates a random amount of money/owner balance
func RandomBalance() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"NGN", "USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}