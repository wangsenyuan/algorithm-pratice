package p1187

import "testing"

func runSample(t *testing.T, arr1 []int, arr2 []int, expect int) {
	res := makeArrayIncreasing(arr1, arr2)
	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", arr1, arr2, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr1 := []int{1, 5, 3, 6, 7}
	arr2 := []int{1, 3, 2, 4}
	expect := 1
	runSample(t, arr1, arr2, expect)
}

func TestSample2(t *testing.T) {
	arr1 := []int{1, 5, 3, 6, 7}
	arr2 := []int{1, 3, 4}
	expect := 2
	runSample(t, arr1, arr2, expect)
}

func TestSample3(t *testing.T) {
	arr1 := []int{1, 5, 3, 6, 7}
	arr2 := []int{1, 3, 6}
	expect := -1
	runSample(t, arr1, arr2, expect)
}

func TestSample4(t *testing.T) {
	arr1 := []int{5, 16, 19, 2, 1, 12, 7, 14, 5, 16}
	arr2 := []int{6, 17, 4, 3, 6, 13, 4, 3, 18, 17, 16, 7, 14, 1, 16}
	expect := 8
	runSample(t, arr1, arr2, expect)
}
