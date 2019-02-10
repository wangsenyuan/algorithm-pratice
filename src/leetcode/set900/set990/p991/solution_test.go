package p991

import "testing"

func runSample(t *testing.T, x, y int, expect int) {
	res := brokenCalc(x, y)
	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", x, y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 8, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 10, 3)
}

func TestSample4(t *testing.T) {
	runSample(t, 1024, 1, 1023)
}
