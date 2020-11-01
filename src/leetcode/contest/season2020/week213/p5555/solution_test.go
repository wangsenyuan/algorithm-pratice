package p5555

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := countVowelStrings(n)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 15)
}

func TestSample3(t *testing.T) {
	runSample(t, 33, 66045)
}
