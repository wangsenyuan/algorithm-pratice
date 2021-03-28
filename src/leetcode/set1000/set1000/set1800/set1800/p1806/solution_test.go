package p1806

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := reinitializePermutation(n)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	expect := 1
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	expect := 2
	runSample(t, n, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	expect := 4
	runSample(t, n, expect)
}

func TestSample4(t *testing.T) {
	n := 10
	expect := 6
	runSample(t, n, expect)
}
