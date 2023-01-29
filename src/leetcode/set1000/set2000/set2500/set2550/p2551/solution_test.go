package p2551

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := monkeyMove(n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	runSample(t, 3, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 14)
}
