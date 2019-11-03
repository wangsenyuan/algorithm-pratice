package p1248

import "testing"

func runSample(t *testing.T, arr []int, k int, expect int) {
	res := numberOfSubarrays(arr, k)
	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", arr, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 1, 2, 1, 1}
	k := 3
	expect := 2
	runSample(t, arr, k, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{2, 4, 6}
	k := 1
	expect := 0
	runSample(t, arr, k, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}
	k := 2
	expect := 16
	runSample(t, arr, k, expect)
}

func TestSample4(t *testing.T) {
	arr := []int{1, 1, 1, 1, 1}
	k := 1
	expect := 5
	runSample(t, arr, k, expect)
}
