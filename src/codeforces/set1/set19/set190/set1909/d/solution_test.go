package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 4}
	k := 1
	expect := 3
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{7, 11}
	k := 3
	expect := 1
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int{100, 40, 100}
	k := 10
	expect := 4
	runSample(t, a, k, expect)
}

func TestSample4(t *testing.T) {
	a := []int{321037808624, 222034505841, 214063039282, 441536506916, 464097941819}
	k := 616753575719
	expect := 28999960732
	runSample(t, a, k, expect)
}
