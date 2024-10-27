package p3333

import "testing"

func runSample(t *testing.T, word string, k int, expect int) {
	res := possibleStringCount(word, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	word := "aabbccdd"
	k := 7
	expect := 5
	runSample(t, word, k, expect)
}

func TestSample2(t *testing.T) {
	word := "aabbccdd"
	k := 8
	expect := 1
	runSample(t, word, k, expect)
}

func TestSample3(t *testing.T) {
	word := "aaabbb"
	// aaab aaabb aaabbb
	// aab aabb aabbb
	// abb abbb
	k := 3
	expect := 8
	runSample(t, word, k, expect)
}
