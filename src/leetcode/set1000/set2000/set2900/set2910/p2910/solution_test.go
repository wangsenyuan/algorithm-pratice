package p2910

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := minimumChanges(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abcac"
	k := 2
	expect := 1
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "abcdef"
	k := 2
	expect := 2
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "aabbaa"
	k := 3
	expect := 0
	runSample(t, s, k, expect)
}

func TestSample4(t *testing.T) {
	s := "acba"
	k := 2
	expect := 2
	runSample(t, s, k, expect)
}
