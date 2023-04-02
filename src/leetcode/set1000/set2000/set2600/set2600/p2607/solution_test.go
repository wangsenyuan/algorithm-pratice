package p2607

import "testing"

func runSample(t *testing.T, arr []int, k int, expect int64) {
	res := makeSubKSumEqual(arr, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 4, 1, 3}
	k := 2
	var expect int64 = 1
	runSample(t, arr, k, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{2, 5, 5, 7}
	k := 3
	var expect int64 = 5
	runSample(t, arr, k, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{6, 2, 8, 5, 7, 10}
	k := 4
	var expect int64 = 10
	runSample(t, arr, k, expect)
}
