package p1589

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := sumOddLengthSubarrays(arr)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 4, 2, 5, 3}
	expect := 58
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{10, 11, 12}
	expect := 66
	runSample(t, arr, expect)
}
