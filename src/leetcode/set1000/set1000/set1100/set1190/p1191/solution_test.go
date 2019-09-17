package p1191

import "testing"

func runSample(t *testing.T, arr []int, k int, expect int) {
	res := kConcatenationMaxSum(arr, k)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", arr, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 2}
	k := 3
	runSample(t, arr, k, 9)
}

func TestSample2(t *testing.T) {
	arr := []int{1, -2, 1}
	k := 5
	runSample(t, arr, k, 2)
}

func TestSample3(t *testing.T) {
	arr := []int{-1, -2}
	k := 5
	runSample(t, arr, k, 0)
}

func TestSample4(t *testing.T) {
	arr := []int{-5, -2, 0, 0, 3, 9, -2, -5, 4}
	k := 5
	expect := 20
	runSample(t, arr, k, expect)
}
