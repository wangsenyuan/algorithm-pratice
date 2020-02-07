package niuke0

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := numberOf1BeforeN(n)
	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 13, 6)
}

func TestSample4(t *testing.T) {
	runSample(t, 55, 16)
}
