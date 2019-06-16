package p885

import "testing"

func runSample(t *testing.T, people []int, limit int, expect int) {
	res := numRescueBoats(people, limit)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", people, limit, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 2}, 3, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{3, 2, 2, 1}, 3, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{3, 5, 3, 4}, 5, 4)
}
