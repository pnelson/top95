package top95

import (
	"math/rand"
	"testing"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func TestInclude(t *testing.T) {
	for _, s := range []string{"password", "pa55w0rd", "sysadmin"} {
		if !Include(s) {
			t.Errorf("%s should be in top95", s)
		}
	}
	if Include("woozy grumbly conductor opposite") {
		t.Fatal("should not be in top95")
	}
}

func BenchmarkInclude(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := rand.Intn(12)
		b := make([]byte, n)
		for i := range b {
			b[i] = charset[rand.Intn(len(charset))]
		}
		Include(string(b))
	}
}
