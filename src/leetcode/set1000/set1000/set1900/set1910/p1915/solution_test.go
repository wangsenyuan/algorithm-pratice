package p1915

import "testing"

func runSample(t *testing.T, word string, expect int64) {
	res := wonderfulSubstrings(word)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	word := "aba"
	var expect int64 = 4
	runSample(t, word, expect)
}

func TestSample2(t *testing.T) {
	word := "aabb"
	var expect int64 = 9
	runSample(t, word, expect)
}

func TestSample3(t *testing.T) {
	word := "he"
	var expect int64 = 2
	runSample(t, word, expect)
}
