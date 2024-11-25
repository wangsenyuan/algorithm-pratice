package p3366

import "testing"

func runSample(t *testing.T, nums []int, k int, op1 int, op2 int, expect int) {
	res := minArraySum(nums, k, op1, op2)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num := []int{2, 8, 3, 19, 3}
	k := 3
	op1 := 1
	op2 := 1
	expect := 23
	runSample(t, num, k, op1, op2, expect)
}

func TestSample2(t *testing.T) {
	num := []int{2, 4, 3}
	k := 3
	op1 := 2
	op2 := 1
	expect := 3
	runSample(t, num, k, op1, op2, expect)
}
