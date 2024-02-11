package main

import "testing"

func runSample(t *testing.T, a []int, b []int, k int, expect int) {
	res := solve(a, b, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	a := []int{1, 2}
	b := []int{3, 4}
	expect := 6
	runSample(t, a, b, k, expect)
}

func TestSample2(t *testing.T) {
	k := 10000
	a := []int{1}
	b := []int{2}
	expect := 1
	runSample(t, a, b, k, expect)
}

func TestSample3(t *testing.T) {
	k := 11037
	a := []int{1, 1, 4, 5}
	b := []int{1, 9, 1, 9, 8}
	expect := 19
	runSample(t, a, b, k, expect)
}

func TestSample4(t *testing.T) {
	k := 572637639
	a := []int{924242480, 924242480, 924242480, 924242480, 924242480}
	b := []int{606050564, 606050564, 7222523, 7222523}
	expect := 4621212400
	runSample(t, a, b, k, expect)
}
