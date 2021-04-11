package p1824

import "testing"

func runSample(t *testing.T, obs []int, expect int) {
	res := minSideJumps(obs)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	obs := []int{0, 1, 2, 3, 0}
	expect := 2
	runSample(t, obs, expect)
}

func TestSample2(t *testing.T) {
	obs := []int{0, 1, 1, 3, 3, 0}
	expect := 0
	runSample(t, obs, expect)
}

func TestSample3(t *testing.T) {
	obs := []int{0, 2, 1, 0, 3, 0}
	expect := 2
	runSample(t, obs, expect)
}
