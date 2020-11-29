package p5557

import "testing"

func runSample(t *testing.T, seq string, word string, expect int) {
	res := maxRepeating(seq, word)

	if res != expect {
		t.Errorf("Sample %s %s, expect %d, but got %d", seq, word, expect, res)
	}
}

func TestSample1(t *testing.T) {
	sequence := "ababc"
	word := "ab"
	expect := 2
	runSample(t, sequence, word, expect)
}

func TestSample2(t *testing.T) {
	sequence := "ababab"
	word := "ab"
	expect := 3
	runSample(t, sequence, word, expect)
}

func TestSample3(t *testing.T) {
	sequence := "ababc"
	word := "ba"
	expect := 1
	runSample(t, sequence, word, expect)
}

func TestSample4(t *testing.T) {
	sequence := "ababc"
	word := "ac"
	expect := 0
	runSample(t, sequence, word, expect)
}
