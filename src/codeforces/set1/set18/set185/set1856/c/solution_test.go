package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 3, 3}
	k := 4
	expect := 4
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 3, 4, 5, 1}
	k := 6
	expect := 7
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 1, 3, 179}
	k := 13
	expect := 179
	runSample(t, a, k, expect)
}

func TestSample4(t *testing.T) {
	a := []int{4, 3, 2, 2, 2}
	k := 3
	expect := 5
	runSample(t, a, k, expect)
}

func TestSample5(t *testing.T) {
	a := []int{6, 5, 4, 1, 5}
	//  6, 5, 4, 4, 4  (3)
	//  6, 5, 5, 4, 4, (1)
	//  6, 6, 5, 4, 4  (1)
	//  7...
	k := 6
	expect := 7
	runSample(t, a, k, expect)
}

func TestSample6(t *testing.T) {
	a := []int{2, 2, 2, 1, 1, 2, 2}
	k := 1
	expect := 3
	runSample(t, a, k, expect)
}
