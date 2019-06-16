package p866

import "testing"

func runSample(t *testing.T, N int, expect int) {
	res := primePalindrome(N)
	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 6, 7)
}

func TestSample2(t *testing.T) {
	runSample(t, 8, 11)
}

func TestSample3(t *testing.T) {
	runSample(t, 13, 101)
}
