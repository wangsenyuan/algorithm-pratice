package p849

import "testing"

func runSample(t *testing.T, seats []int, expect int) {
	res := maxDistToClosest(seats)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", seats, expect, res)
	}
}

func TestSample1(t *testing.T) {
	seats := []int{1, 0, 0, 0, 1, 0, 1}
	expect := 2
	runSample(t, seats, expect)
}

func TestSample2(t *testing.T) {
	seats := []int{1, 0, 0, 0}
	expect := 3
	runSample(t, seats, expect)
}
