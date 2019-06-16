package p974

import "testing"

func runSample(t *testing.T, nums []int, K int, expect int) {
	res := subarraysDivByK(nums, K)
	if res != expect {
		t.Errorf("Sample %v, %d, expect %d, but got %d", nums, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{4, 5, 0, -2, -3, 1}
	K := 5
	expect := 7
	runSample(t, nums, K, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{7, 4, -10}
	K := 5
	expect := 1
	runSample(t, nums, K, expect)
}
