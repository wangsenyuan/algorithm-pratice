package p1986

import "testing"

func runSample(t *testing.T, tasks []int, sessionTime int, expect int) {
	res := minSessions(tasks, sessionTime)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	tasks := []int{1, 2, 3}
	sessionTime := 3
	expect := 2
	runSample(t, tasks, sessionTime, expect)
}

func TestSample2(t *testing.T) {
	tasks := []int{3, 1, 3, 1, 1}
	sessionTime := 8
	expect := 2
	runSample(t, tasks, sessionTime, expect)
}

func TestSample3(t *testing.T) {
	tasks := []int{1, 2, 3, 4, 5}
	sessionTime := 15
	expect := 1
	runSample(t, tasks, sessionTime, expect)
}
