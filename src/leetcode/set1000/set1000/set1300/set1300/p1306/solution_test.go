package p1306

import "testing"

func runSample(t *testing.T, arr []int, start int, expect bool) {
	res := canReach(arr, start)

	if res != expect {
		t.Errorf("Sample %v %d, expect %t, but got %t", arr, start, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{4, 2, 3, 0, 3, 1, 2}
	start := 5
	expect := true
	runSample(t, arr, start, expect)
}
