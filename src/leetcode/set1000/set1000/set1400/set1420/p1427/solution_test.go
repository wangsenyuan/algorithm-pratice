package p1427

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := maxDiff(n)

	if res != expect {
		t.Errorf("Sample %d expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 555, 888)
}

func TestSample2(t *testing.T) {
	runSample(t, 9, 8)
}

func TestSample3(t *testing.T) {
	runSample(t, 123456, 820000)
}

func TestSample4(t *testing.T) {
	runSample(t, 10000, 80000)
}

func TestSample5(t *testing.T) {
	runSample(t, 9288, 8700)
}

func TestSample6(t *testing.T) {
	runSample(t, 1101057, 8808050)
}
