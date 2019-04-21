package p1032

import "testing"

func runSample(t *testing.T, words []string, queries []byte, expect []bool) {
	checker := Constructor(words)

	for i := 0; i < len(queries); i++ {
		res := checker.Query(queries[i])
		if res != expect[i] {
			t.Fatalf("Sample %v, expect %v, but fail at %d", words, expect, i)
		}
	}
}

func TestSample1(t *testing.T) {
	words := []string{"cd", "f", "kl"}
	queries := []byte("abcdefghijkl")
	expect := []bool{false, false, false, true, false, true, false, false, false, false, false, true}
	runSample(t, words, queries, expect)
}
