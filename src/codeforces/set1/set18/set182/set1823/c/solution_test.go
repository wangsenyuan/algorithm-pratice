package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 6}
	expect := 1
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 4, 5}
	expect := 1
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 3}
	expect := 0
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{3, 10, 14}
	expect := 2
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{25, 30}
	expect := 2
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{1080}
	expect := 3
	runSample(t, a, expect)
}

func TestSample7(t *testing.T) {
	a := []int{3, 3, 3, 5, 5, 5, 7, 7, 7}
	expect := 4
	runSample(t, a, expect)
}

func TestSample8(t *testing.T) {
	a := []int{12, 15, 2, 2, 2, 2, 2, 3, 3, 3, 17, 21, 21, 21, 30, 6, 6, 33, 31, 39}
	expect := 15
	runSample(t, a, expect)
}
