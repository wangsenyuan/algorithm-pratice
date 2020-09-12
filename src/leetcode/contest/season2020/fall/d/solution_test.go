package d

import "testing"

func runSample(t *testing.T, target int, inc int, dec int, jump []int, cost []int, expect int) {
	res := busRapidTransit(target, inc, dec, jump, cost)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	target := 31
	inc := 5
	dec := 3
	jump := []int{6}
	cost := []int{10}
	expect := 33
	runSample(t, target, inc, dec, jump, cost, expect)
}
