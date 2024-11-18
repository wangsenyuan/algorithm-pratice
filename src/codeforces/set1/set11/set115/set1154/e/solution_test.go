package main

import "testing"

func runSample(t *testing.T, k int, a []int, expect string) {
	res := solve(a, k)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	a := []int{2, 4, 5, 3, 1}
	expect := "11111"
	runSample(t, k, a, expect)
}

func TestSample2(t *testing.T) {
	k := 1
	a := []int{2, 1, 3, 5, 4}
	expect := "22111"
	runSample(t, k, a, expect)
}

func TestSample3(t *testing.T) {
	k := 1
	a := []int{7, 2, 1, 3, 5, 4, 6}
	expect := "1121122"
	runSample(t, k, a, expect)
}
