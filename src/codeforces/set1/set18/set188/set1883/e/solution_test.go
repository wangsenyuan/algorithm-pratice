package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)
	if res != expect {
		t.Fatalf("Sample expect  %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 1}
	expect := 1
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 2, 1}
	expect := 3
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{7, 1, 5, 3}
	expect := 6
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{11, 2, 15, 7, 10}
	expect := 10
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{1, 8, 2, 16, 8, 16}
	expect := 3
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	expect := 66
	runSample(t, a, expect)
}

func TestSample7(t *testing.T) {
	a := []int{2, 1, 1, 3, 3, 11, 12, 22, 45, 777, 777, 1500}
	expect := 2
	runSample(t, a, expect)
}
