package p2244

import "testing"

func runSample(t *testing.T, tasks []int, expect int) {
	res := minimumRounds(tasks)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	tasks := []int{5, 5, 5, 5}
	expect := 2
	runSample(t, tasks, expect)
}

func TestSample2(t *testing.T) {
	tasks := []int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}
	expect := 4
	runSample(t, tasks, expect)
}
