package a

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := minCount(arr)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{4, 2, 1}, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{2, 3, 10}, 8)
}
