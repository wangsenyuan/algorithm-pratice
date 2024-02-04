package p3027

import "testing"

func runSample(t *testing.T, word string, k int, expect int) {
	res := minimumTimeToInitialState(word, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	word := "abacaba"
	k := 3
	expect := 2
	runSample(t, word, k, expect)
}

func TestSample2(t *testing.T) {
	word := "abacaba"
	k := 4
	expect := 1
	runSample(t, word, k, expect)
}
