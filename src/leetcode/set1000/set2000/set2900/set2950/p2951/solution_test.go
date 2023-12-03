package p2951

import "testing"

func runSample(t *testing.T, word string, k int, expect int) {
	res := countCompleteSubstrings(word, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	word := "igigee"
	k := 2
	expect := 3
	runSample(t, word, k, expect)
}

func TestSample2(t *testing.T) {
	word := "aaabbbccc"
	k := 3
	expect := 6
	runSample(t, word, k, expect)
}
