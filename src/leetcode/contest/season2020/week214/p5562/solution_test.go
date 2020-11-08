package p5562

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := minDeletions(s)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aab"
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "aaabbbcc"
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "ceabaacb"
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "accdcdadddbaadbc"
	expect := 1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "bbcebab"
	expect := 2
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "beaddedbacdcd"
	expect := 5
	runSample(t, s, expect)
}
