package p1774

import "testing"

func runSample(t *testing.T, base []int, topping []int, target int, expect int) {
	res := closestCost(base, topping, target)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	base := []int{1, 7}
	topping := []int{3, 4}
	target := 10
	expect := 10
	runSample(t, base, topping, target, expect)
}

func TestSample2(t *testing.T) {
	base := []int{3, 10}
	topping := []int{2, 5}
	target := 9
	expect := 8
	runSample(t, base, topping, target, expect)
}

func TestSample3(t *testing.T) {
	base := []int{2, 3}
	topping := []int{4, 5, 100}
	target := 18
	expect := 17
	runSample(t, base, topping, target, expect)
}
