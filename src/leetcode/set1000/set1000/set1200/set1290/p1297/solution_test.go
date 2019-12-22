package p1297

import "testing"

func runSample(t *testing.T, s string, maxLetters int, minSize int, maxSize int, expect int) {
	res := maxFreq(s, maxLetters, minSize, maxSize)

	if res != expect {
		t.Errorf("Sample %s %d %d %d, expect %d, but got %d", s, maxLetters, minSize, maxSize, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aababcaab"
	maxLetters := 2
	minSize := 3
	maxSize := 4
	expect := 2
	runSample(t, s, maxLetters, minSize, maxSize, expect)
}

func TestSample2(t *testing.T) {
	s := "aaaa"
	maxLetters := 1
	minSize := 1
	maxSize := 4
	expect := 4
	runSample(t, s, maxLetters, minSize, maxSize, expect)
}

func TestSample3(t *testing.T) {
	s := "aabcabcab"
	maxLetters := 2
	minSize := 2
	maxSize := 3
	expect := 3
	runSample(t, s, maxLetters, minSize, maxSize, expect)
}

func TestSample4(t *testing.T) {
	s := "abcde"
	maxLetters := 2
	minSize := 3
	maxSize := 3
	expect := 0
	runSample(t, s, maxLetters, minSize, maxSize, expect)
}
