package p926

import "testing"

func runSample(t *testing.T, S string, expect int) {
	res := minFlipsMonoIncr(S)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "00110", 1)
}

func TestSample2(t *testing.T) {
	runSample(t, "010110", 2)
}
func TestSample3(t *testing.T) {
	runSample(t, "00011000", 2)
}
