package p796

import "strings"

func rotateString(A string, B string) bool {
	return strings.Contains(A+A, B)
}
