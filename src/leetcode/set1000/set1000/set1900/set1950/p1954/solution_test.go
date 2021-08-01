package p1954

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := countSpecialSubsequences(arr)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{0, 1, 2, 2}
	expect := 3
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{0, 1, 2, 0, 1, 2}
	expect := 7
	runSample(t, arr, expect)
}
