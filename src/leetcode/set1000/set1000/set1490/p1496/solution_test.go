package p1496

import "testing"

func runSample(t *testing.T, arr []int, k int, expect bool) {
	res := canArrange(arr, k)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}
	expect := true

	runSample(t, arr, 5, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{-1, 1, -2, 2, -3, 3, -4, 4}
	expect := true

	runSample(t, arr, 3, expect)
}
