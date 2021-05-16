package p1859

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := sumOfFlooredPairs(arr)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{2, 5, 9}
	expect := 10
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{7, 7, 7, 7, 7, 7, 7}
	expect := 49
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	expect := 27
	runSample(t, arr, expect)
}
