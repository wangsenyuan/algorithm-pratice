package p2035

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minimumDifference(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 9, 7, 3}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{-36, 36}
	expect := 72
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, -1, 0, 4, -2, -9}
	expect := 0
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{7772197, 4460211, -7641449, -8856364, 546755, -3673029, 527497, -9392076, 3130315, -5309187, -4781283, 5919119, 3093450, 1132720, 6380128, -3954678, -1651499, -7944388, -3056827, 1610628, 7711173, 6595873, 302974, 7656726, -2572679, 0, 2121026, -5743797, -8897395, -9699694}
	expect := 1
	runSample(t, nums, expect)
}
