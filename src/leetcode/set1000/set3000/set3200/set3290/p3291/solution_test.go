package p3291

import "testing"

func runSample(t *testing.T, words []string, target string, expect int) {
	res := minValidStrings(words, target)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{"abc", "aaaaa", "bcdef"}
	target := "aabcdabc"
	expect := 3
	runSample(t, words, target, expect)
}

func TestSample2(t *testing.T) {
	words := []string{"abababab", "ab"}
	target := "ababaababa"
	expect := 2
	runSample(t, words, target, expect)
}

func TestSample3(t *testing.T) {
	words := []string{"abcdef", "ab"}
	target := "xyz"
	expect := -1
	runSample(t, words, target, expect)
}
