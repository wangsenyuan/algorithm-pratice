package p1363

import "testing"

func runSample(t *testing.T, nums []int, expect string) {
	res := largestMultipleOfThree(nums)

	if res != expect {
		t.Errorf("%v, expect %s, but got %s", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{8, 1, 9}, "981")
}

func TestSample2(t *testing.T) {
	runSample(t, []int{8, 6, 7, 1, 0}, "8760")
}

func TestSample3(t *testing.T) {
	runSample(t, []int{1}, "")
}

func TestSample4(t *testing.T) {
	runSample(t, []int{0, 0, 0, 0, 0, 0}, "0")
}

func TestSample5(t *testing.T) {
	runSample(t, []int{8, 6, 7, 2, 0}, "8760")
}
