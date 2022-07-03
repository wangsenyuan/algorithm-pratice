package p556

import "testing"

func runSample(t *testing.T, num int, expect int) {
	res := nextGreaterElement(num)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num := 2147483486
	expect := -1
	runSample(t, num, expect)
}
