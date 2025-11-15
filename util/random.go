package util

import (
	"fmt"
	"math/rand"
	db "simplebank/db/sqlc"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn((k))]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomUsername() string {
	return RandomString(6)
}

func RandomCurrency() string {
	currencies := []string{EUR, USD, CAD, INR}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func randomUser(t *testing.T) (user db.User, password string) {
	password = RandomString(6)
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)

	user = db.User{
		Username:       RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       RandomOwner(),
		Email:          RandomEmail(),
	}
	return
}
