package p1523

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := numOfSubarrays(arr)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 3, 5}
	expect := 4
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{2, 4, 6}
	expect := 0
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	expect := 16
	runSample(t, arr, expect)
}

func TestSample4(t *testing.T) {
	arr := []int{100, 100, 99, 99}
	expect := 4
	runSample(t, arr, expect)
}
