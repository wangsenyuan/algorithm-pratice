package p1574

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := findLengthOfShortestSubarray(arr)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 2, 3, 10, 4, 2, 3, 5}
	expect := 3
	runSample(t, arr, expect)
}
func TestSample2(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	expect := 4
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{1, 2, 3}
	expect := 0
	runSample(t, arr, expect)
}

func TestSample4(t *testing.T) {
	arr := []int{16, 10, 0, 3, 22, 1, 14, 7, 1, 12, 15}
	expect := 8
	runSample(t, arr, expect)
}
