package p1013

import "testing"

func runSample(t *testing.T, time []int, expect int) {
	res := numPairsDivisibleBy60(time)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", time, expect, res)
	}
}

func TestSample1(t *testing.T) {
	time := []int{30, 20, 150, 100, 40}
	runSample(t, time, 3)
}

func TestSample2(t *testing.T) {
	time := []int{60, 60, 60}
	runSample(t, time, 3)
}
