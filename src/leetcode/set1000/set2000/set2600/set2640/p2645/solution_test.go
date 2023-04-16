package p2645

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := addMinimum(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	word := "aaa"
	expect := 6
	runSample(t, word, expect)
}

func TestSample2(t *testing.T) {
	word := "abc"
	expect := 0
	runSample(t, word, expect)
}
