package p2064

import "testing"

func runSample(t *testing.T, n int, quantities []int, expect int) {
	res := minimizedMaximum(n, quantities)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	quantities := []int{11, 6}
	expect := 3
	runSample(t, n, quantities, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	quantities := []int{15, 10, 10}
	expect := 5
	runSample(t, n, quantities, expect)
}
