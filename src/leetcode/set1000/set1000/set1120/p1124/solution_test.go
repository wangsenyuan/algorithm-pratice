package p1124

import "testing"

func runSample(t *testing.T, hours []int, expect int) {
	res := longestWPI(hours)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", hours, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{9, 9, 6, 0, 6, 6, 9}, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{9, 9, 6, 0, 6, 6, 9, 9, 9}, 9)
}
