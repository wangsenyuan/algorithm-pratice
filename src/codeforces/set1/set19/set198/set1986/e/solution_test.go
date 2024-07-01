package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{624323799, 708290323}
	k := 1
	expect := 83966524
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 2, 1}
	k := 1
	expect := 1
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int{7, 1, 5, 3}
	k := 1
	expect := 4
	runSample(t, a, k, expect)
}

func TestSample4(t *testing.T) {
	a := []int{11, 2, 15, 7, 10}
	k := 1
	expect := 6
	runSample(t, a, k, expect)
}

func TestSample5(t *testing.T) {
	a := []int{1, 8, 2, 16, 8, 16, 31}
	k := 1
	expect := 1
	runSample(t, a, k, expect)
}

func TestSample6(t *testing.T) {
	a := []int{2, 1, 1, 3, 3, 11, 12, 22, 45, 777, 777, 1500, 74}
	k := 1
	expect := 48
	runSample(t, a, k, expect)
}

func TestSample7(t *testing.T) {
	a := []int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2}
	k := 2
	expect := -1
	runSample(t, a, k, expect)
}

func TestSample8(t *testing.T) {
	a := []int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1}
	k := 2
	expect := 0
	runSample(t, a, k, expect)
}
