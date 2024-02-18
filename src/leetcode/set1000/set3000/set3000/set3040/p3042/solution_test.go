package p3042

import "testing"

func runSample(t *testing.T, words []string, expect int) {
	res := countPrefixSuffixPairs(words)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{"a", "aba", "ababa", "aa"}
	expect := 4
	runSample(t, words, expect)
}

func TestSample2(t *testing.T) {
	words := []string{"pa", "papa", "ma", "mama"}
	expect := 2
	runSample(t, words, expect)
}

func TestSample3(t *testing.T) {
	words := []string{"abab", "ab"}
	expect := 0
	runSample(t, words, expect)
}
