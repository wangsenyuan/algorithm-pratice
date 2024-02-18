package p3043

import "testing"

func runSample(t *testing.T, arr1 []int, arr2 []int, expect int) {
	res := longestCommonPrefix(arr1, arr2)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr1 := []int{1, 10, 100}
	arr2 := []int{1000}
	expect := 3
	runSample(t, arr1, arr2, expect)
}

func TestSample2(t *testing.T) {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4}
	expect := 0
	runSample(t, arr1, arr2, expect)
}
