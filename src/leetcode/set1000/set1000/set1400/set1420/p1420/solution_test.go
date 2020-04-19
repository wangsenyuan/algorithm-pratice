package p1420

import "testing"

func runSample(t *testing.T, n, m, k int, expect int) {
	res := numOfArrays(n, m, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, 1, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 2, 3, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, 9, 1, 1, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 50, 100, 25, 34549172)
}

func TestSample5(t *testing.T) {
	runSample(t, 37, 17, 7, 418930126)
}
