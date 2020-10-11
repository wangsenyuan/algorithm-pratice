package productsum

import "testing"

func runSample(t *testing.T, arr []int, m int, expect int) {
	res := ProductSum(arr, m)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", arr, m, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 2, 3}
	m := 2
	expect := 11
	runSample(t, arr, m, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{2, 3, 4, 5}
	m := 3
	expect := 154
	runSample(t, arr, m, expect)
}
