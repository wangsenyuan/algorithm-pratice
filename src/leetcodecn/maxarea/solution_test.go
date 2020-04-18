package maxarea

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := maxArea(arr)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expect := 49
	runSample(t, height, expect)
}

func TestSample2(t *testing.T) {
	height := []int{1, 2, 1}
	expect := 2
	runSample(t, height, expect)
}
