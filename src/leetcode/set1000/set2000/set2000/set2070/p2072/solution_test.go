package p2072

import "testing"

func runSample(t *testing.T, tickes []int, k int, expect int) {
	res := timeRequiredToBuy(tickes, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	tickets := []int{2, 3, 2}
	k := 2
	expect := 6
	runSample(t, tickets, k, expect)
}

func TestSample2(t *testing.T) {
	tickets := []int{5, 1, 1, 1}
	k := 0
	expect := 8
	runSample(t, tickets, k, expect)
}

func TestSample3(t *testing.T) {
	tickets := []int{84, 49, 5, 24, 70, 77, 87, 8}
	k := 3
	expect := 154
	runSample(t, tickets, k, expect)
}
