package p1442

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := countTriplets(arr)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{2, 3, 1, 6, 7}
	expect := 4
	runSample(t, arr, expect)
}
