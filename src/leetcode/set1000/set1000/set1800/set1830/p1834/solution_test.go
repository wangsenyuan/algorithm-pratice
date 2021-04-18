package p1834

import "testing"

func runSample(t *testing.T, arr1, arr2 []int, expect int) {
	res := getXORSum(arr1, arr2)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{6, 5}
	expect := 0
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{12}
	b := []int{4}
	expect := 4
	runSample(t, a, b, expect)
}
