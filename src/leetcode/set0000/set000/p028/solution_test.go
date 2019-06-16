package p028

import "testing"

func TestSample1(t *testing.T) {
	haystack := "hello"
	needle := "ll"
	res := strStr(haystack, needle)
	if res != 2 {
		t.Errorf("%s should at 2 in %s, but got %d", needle, haystack, res)
	}
}
