package p1529

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := getLengthOfOptimalCompression(s, k)
	if res != expect {
		t.Errorf("Sample %s, %d, expect %d, but got %d", s, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aaabcccd"
	k := 2
	expect := 4
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "aabbaa"
	k := 2
	expect := 2
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "aaaaaaaaaaa"
	k := 0
	expect := 3
	runSample(t, s, k, expect)
}
