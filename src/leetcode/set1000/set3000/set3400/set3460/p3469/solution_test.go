package p3469

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minCost(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	a := []int{6, 2, 8, 4}
	expect := 12
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 1, 3, 3}
	expect := 5
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{9, 1, 5}
	expect := 10
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{8, 5, 7, 9, 5}
	expect := 21
	runSample(t, a, expect)
}
