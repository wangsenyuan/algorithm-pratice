package p1712

import "testing"

func runSample(t *testing.T, target []int, arr []int, expect int) {
	res := minOperations(target, arr)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	target := []int{5, 1, 3}
	arr := []int{9, 4, 2, 3, 4}
	expect := 2
	runSample(t, target, arr, expect)
}

func TestSample2(t *testing.T) {
	target := []int{6, 4, 8, 1, 3, 2}
	arr := []int{4, 7, 6, 2, 3, 8, 6, 1}
	expect := 3
	runSample(t, target, arr, expect)
}

func TestSample3(t *testing.T) {
	target := []int{2, 7, 4}
	arr := []int{4, 2, 3, 7, 2, 1, 4}
	expect := 0
	runSample(t, target, arr, expect)
}

func TestSample4(t *testing.T) {
	target := []int{16, 7, 20, 11, 15, 13, 10, 14, 6, 8}
	arr := []int{11, 14, 15, 7, 5, 5, 6, 10, 11, 6}
	expect := 6
	runSample(t, target, arr, expect)
}
