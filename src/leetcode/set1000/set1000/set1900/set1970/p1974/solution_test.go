package p1974

import "testing"

func runSample(t *testing.T, word string, expect int) {
	res := minTimeToType(word)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	word := "abc"
	expect := 5
	runSample(t, word, expect)
}

func TestSample2(t *testing.T) {
	word := "bza"
	expect := 7
	runSample(t, word, expect)
}

func TestSample3(t *testing.T) {
	word := "zjpc"
	expect := 34
	runSample(t, word, expect)
}