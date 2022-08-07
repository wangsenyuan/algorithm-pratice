package p2366

import "testing"

func runSample(t *testing.T, tasks []int, space int, expect int) {
	res := int(taskSchedulerII(tasks, space))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	tasks := []int{1, 2, 1, 2, 3, 1}
	space := 3
	expect := 9
	runSample(t, tasks, space, expect)
}

func TestSample2(t *testing.T) {
	tasks := []int{5, 8, 8, 5}
	space := 2
	expect := 6
	runSample(t, tasks, space, expect)
}
