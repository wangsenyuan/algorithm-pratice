package p858

import "testing"

func runSample(t *testing.T, p, q int, expect int) {
	res := mirrorReflection(p, q)
	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", p, q, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 1, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 2, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, 3, 2)
}

func TestSample4(t *testing.T) {
	runSample(t, 5, 3, 1)
}
