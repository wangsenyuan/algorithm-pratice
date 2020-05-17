package p1449

import "testing"

func runSample(t *testing.T, starting, ending []int, queryAt int, expect int) {
	res := busyStudent(starting, ending, queryAt)

	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", starting, ending, expect, res)
	}
}

func TestSample1(t *testing.T) {
	start := []int{1, 2, 3}
	end := []int{3, 2, 7}
	queryAt := 4
	expect := 1
	runSample(t, start, end, queryAt, expect)
}
