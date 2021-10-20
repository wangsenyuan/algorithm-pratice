package lcp47

import "testing"

func runSample(t *testing.T, capacities []int, k int, expect int) {
	res := securityCheck(capacities, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	capacities := []int{2, 2, 3}
	k := 2
	expect := 2
	runSample(t, capacities, k, expect)
}

func TestSample2(t *testing.T) {
	capacities := []int{4, 3, 2, 2}
	k := 6
	expect := 2
	runSample(t, capacities, k, expect)
}

func TestSample3(t *testing.T) {
	capacities := []int{3, 3}
	k := 3
	expect := 0
	runSample(t, capacities, k, expect)
}
