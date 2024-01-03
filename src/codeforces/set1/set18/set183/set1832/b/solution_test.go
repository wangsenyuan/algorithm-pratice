package main

import "testing"

func runSample(t *testing.T, b []int, k int, expect int) {
	res := solve(b, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	b := []int{2, 5, 1, 10, 6}
	expect := 21
	runSample(t, b, k, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	b := []int{2, 5, 1, 10, 6}
	expect := 11
	runSample(t, b, k, expect)
}

func TestSample3(t *testing.T) {
	k := 1
	b := []int{1, 2, 3}
	expect := 3
	runSample(t, b, k, expect)
}

func TestSample4(t *testing.T) {
	k := 1
	b := []int{15, 22, 12, 10, 13, 11}
	expect := 62
	runSample(t, b, k, expect)
}

func TestSample5(t *testing.T) {
	k := 2
	b := []int{15, 22, 12, 10, 13, 11}
	expect := 46
	runSample(t, b, k, expect)
}

func TestSample6(t *testing.T) {
	k := 1
	b := []int{999999996, 999999999, 999999997, 999999998, 999999995}
	expect := 3999999986
	runSample(t, b, k, expect)
}
