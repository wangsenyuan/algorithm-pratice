package p2187

import "testing"

func runSample(t *testing.T, time []int, trips int, expect int64) {
	res := minimumTime(time, trips)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	time := []int{1, 2, 3}
	trips := 5
	var expect int64 = 3
	runSample(t, time, trips, expect)
}

func TestSample2(t *testing.T) {
	time := []int{2}
	trips := 1
	var expect int64 = 2
	runSample(t, time, trips, expect)
}
