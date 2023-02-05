package p2565

import "testing"

func runSample(t *testing.T, arr []int, k int, expect int) {
	res := minCapability(arr, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSampele(t *testing.T) {
	arr := []int{9, 6, 20, 21, 8}
	k := 3
	expect := 20
	runSample(t, arr, k, expect)
}

func TestSampele2(t *testing.T) {
	arr := []int{4, 22, 11, 14, 25}
	k := 3
	expect := 25
	runSample(t, arr, k, expect)
}
