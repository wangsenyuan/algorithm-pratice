package p1393

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := findLucky(arr)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{2, 2, 3, 4}, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{1, 2, 2, 3, 3, 3}, 3)
}

func TestSample4(t *testing.T) {
	runSample(t, []int{2, 2, 2, 3, 3}, -1)
}
