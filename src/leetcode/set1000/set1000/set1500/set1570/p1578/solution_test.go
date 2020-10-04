package p1578

import "testing"

func runSample(t *testing.T, s string, cost []int, expect int) {
	res := minCost(s, cost)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abaac"
	cost := []int{1, 2, 3, 4, 5}
	expect := 3
	runSample(t, s, cost, expect)
}

func TestSample2(t *testing.T) {
	s := "abc"
	cost := []int{1, 2, 3}
	expect := 0
	runSample(t, s, cost, expect)
}

func TestSample3(t *testing.T) {
	s := "aabaa"
	cost := []int{1, 2, 3, 4, 1}
	expect := 2
	runSample(t, s, cost, expect)
}

func TestSample4(t *testing.T) {
	s := "cddcdcae"
	cost := []int{4, 8, 8, 4, 4, 5, 4, 2}
	expect := 8
	runSample(t, s, cost, expect)
}
