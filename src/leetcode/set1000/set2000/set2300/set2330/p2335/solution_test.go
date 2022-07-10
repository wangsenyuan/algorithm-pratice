package p2335

import "testing"

func runSample(t *testing.T, amount []int, expect int) {
	res := fillCups(amount)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	amount := []int{1, 4, 2}
	expect := 4
	runSample(t, amount, expect)
}

func TestSample2(t *testing.T) {
	amount := []int{5, 4, 4}
	expect := 7
	runSample(t, amount, expect)
}
