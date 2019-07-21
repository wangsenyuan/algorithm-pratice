package p1131

import "testing"

func runSample(t *testing.T, arr1 []int, arr2 []int, expect int) {
	res := maxAbsValExpr(arr1, arr2)
	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", arr1, arr2, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{-1, 4, 5, 6}
	expect := 13
	runSample(t, arr1, arr2, expect)
}

func TestSample2(t *testing.T) {
	arr1 := []int{1, -2, -5, 0, 10}
	arr2 := []int{0, -2, -1, -7, -4}
	expect := 20
	runSample(t, arr1, arr2, expect)
}
