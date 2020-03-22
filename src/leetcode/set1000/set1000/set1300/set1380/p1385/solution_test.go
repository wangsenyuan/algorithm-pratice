package p1385

import "testing"

func runSample(t *testing.T, arr1, arr2 []int, d int, expect int) {
	res := findTheDistanceValue(arr1, arr2, d)

	if res != expect {
		t.Errorf("Sample %v %v %d, expect %d, but got %d", arr1, arr2, d, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr1 := []int{1, 4, 2, 3}
	arr2 := []int{-4, -3, 6, 10, 20, 30}
	d := 3
	expect := 2
	runSample(t, arr1, arr2, d, expect)
}
