package main

import "testing"

func runSample(t *testing.T, nums []int, m int, expect bool) {
	res := solve(nums, m)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 2
	a := []int{14, 2}
	expect := true
	runSample(t, a, m, expect)
}

func TestSample2(t *testing.T) {
	m := 5
	a := []int{9, 56, 1}
	expect := false
	runSample(t, a, m, expect)
}

func TestSample3(t *testing.T) {
	m := 10
	a := []int{1, 2007, 800, 1580}
	expect := false
	runSample(t, a, m, expect)
}

func TestSample4(t *testing.T) {
	m := 5
	a := []int{5000, 123, 30, 4}
	expect := true
	runSample(t, a, m, expect)
}

func TestSample5(t *testing.T) {
	m := 10
	a := []int{6, 4, 6, 2, 3, 1, 10, 9, 10, 7}
	expect := true
	runSample(t, a, m, expect)
}

func TestSample6(t *testing.T) {
	m := 5
	a := []int{10, 10, 10, 10}
	expect := true
	runSample(t, a, m, expect)
}
