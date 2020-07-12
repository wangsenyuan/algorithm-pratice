package p309

import "testing"

func runSample(t *testing.T, prices []int, expect int) {
	res := maxProfit(prices)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", prices, expect, res)
	}
}

func TestSample1(t *testing.T) {
	prices := []int{1, 2, 3, 0, 2}
	expect := 3
	runSample(t, prices, expect)
}

func TestSample2(t *testing.T) {
	prices := []int{1, 2, 4}
	expect := 3
	runSample(t, prices, expect)
}
