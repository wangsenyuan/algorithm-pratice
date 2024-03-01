package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 9, 7, 15, 1}
	k := 2
	expect := 1
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{779547116602436426, 578223540024979445, 335408917861648772, 74859962623690081, 252509054433933447, 760713016476190629, 919845426262703497, 585335723211047202, 522842184971407775}
	k := 2
	expect := 899214809992477
	runSample(t, a, k, expect)
}
