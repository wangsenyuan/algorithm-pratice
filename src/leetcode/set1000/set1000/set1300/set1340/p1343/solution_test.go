package p1343

import "testing"

func runSample(t *testing.T, arr []int, k int, threshold int, expect int) {
	res := numOfSubarrays(arr, k, threshold)

	if res != expect {
		t.Errorf("Sample %v %d %d, expect %d, but got %d", arr, k, threshold, expect, res)
	}

}

func TestSample1(t *testing.T) {
	arr := []int{2, 2, 2, 2, 5, 5, 5, 8}
	k := 3
	threshold := 4
	expect := 3
	runSample(t, arr, k, threshold, expect)
}
