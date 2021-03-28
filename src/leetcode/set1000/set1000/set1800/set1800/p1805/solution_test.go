package p1805

import "testing"

func runSample(t *testing.T, word string, expect int) {
	res := numDifferentIntegers(word)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	word := "a123bc34d8ef34"
	expect := 3
	runSample(t, word, expect)
}

func TestSample2(t *testing.T) {
	word := "a1b01c001"
	expect := 1
	runSample(t, word, expect)
}
