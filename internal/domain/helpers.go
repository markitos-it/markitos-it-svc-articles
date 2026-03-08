package domain

import (
	"crypto/rand"
	"math/big"
	"testing"
)

func HelperRandomAlphaPrefix(t *testing.T, length int) string {
	t.Helper()

	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	limit := big.NewInt(int64(len(alphabet)))

	for i := range b {
		n, err := rand.Int(rand.Reader, limit)
		if err != nil {
			t.Fatalf("failed generating random prefix: %v", err)
		}
		b[i] = alphabet[n.Int64()]
	}

	return string(b)
}
