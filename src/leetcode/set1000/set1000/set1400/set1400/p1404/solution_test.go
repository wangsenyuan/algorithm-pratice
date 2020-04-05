package p1404

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := numSteps(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "1101", 6)
}

func TestSample2(t *testing.T) {
	runSample(t, "10", 1)
}

func TestSample3(t *testing.T) {
	runSample(t, "1", 0)
}

func TestSample4(t *testing.T) {
	runSample(t, "11", 3)
}

func TestSample5(t *testing.T) {
	runSample(t, "111", 4)
}
