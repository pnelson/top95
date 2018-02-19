// Package top95 tests strings against a popular wordlist.
package top95

//go:generate top95 -o wordlist.go

// Include returns true if s is included in the wordlist.
func Include(s string) bool {
	_, ok := wordlist[s]
	return ok
}
