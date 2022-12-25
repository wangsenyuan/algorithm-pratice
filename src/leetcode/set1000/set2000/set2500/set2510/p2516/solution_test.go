package p2516

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := takeCharacters(s, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aabaaaacaabc"
	k := 2
	expect := 8
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "a"
	k := 1
	expect := -1
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "a"
	k := 0
	expect := 0
	runSample(t, s, k, expect)
}
