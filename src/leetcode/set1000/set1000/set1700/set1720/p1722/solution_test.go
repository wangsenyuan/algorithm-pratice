package p1722

import "testing"

func runSample(t *testing.T, jobs []int, k int, expect int) {
	res := minimumTimeRequired(jobs, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	jobs := []int{3, 2, 3}
	k := 3
	expect := 3
	runSample(t, jobs, k, expect)
}

func TestSample2(t *testing.T) {
	jobs := []int{1, 2, 4, 7, 8}
	k := 2
	expect := 11
	runSample(t, jobs, k, expect)
}

func TestSample3(t *testing.T) {
	jobs := []int{5, 15, 4, 9, 15, 8, 8, 9}
	k := 2
	expect := 37
	runSample(t, jobs, k, expect)
}
