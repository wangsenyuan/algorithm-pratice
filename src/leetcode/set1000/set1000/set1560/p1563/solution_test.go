package p1563

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := stoneGameV(arr)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{6, 2, 3, 4, 5, 5}
	expect := 18
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{7, 7, 7, 7, 7, 7, 7}
	expect := 28
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{2, 1, 1}
	expect := 3
	runSample(t, arr, expect)
}
