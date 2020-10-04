package p1553

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := minDays(n)

	if res != expect {
		t.Errorf("Sample %d expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 56, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 6, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 5, 4)
}

func TestSample5(t *testing.T) {
	runSample(t, 2000000001, 32)
}

func TestSample6(t *testing.T) {
	runSample(t, 10, 4)
}

func TestSample7(t *testing.T) {
	runSample(t, 16, 5)
}
