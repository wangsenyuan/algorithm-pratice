package p2110

import "testing"

func runSample(t *testing.T, prices []int, expect int) {
	res := int(getDescentPeriods(prices))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	prices := []int{3, 2, 1, 4}
	expect := 7
	runSample(t, prices, expect)
}
