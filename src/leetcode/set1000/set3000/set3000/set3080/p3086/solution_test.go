package p3086

import "testing"

func runSample(t *testing.T, word string, k int, expect int) {
	res := minimumDeletions(word, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	word := "dabdcbdcdcd"
	k := 2
	expect := 2
	runSample(t, word, k, expect)
}
