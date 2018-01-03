package client

// Chop - Easily delete trailing characters from a string
func Chop(s string, i int) string {
	return s[0 : len(s)-i]
}
