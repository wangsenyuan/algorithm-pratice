package main

import "testing"

func runSample(t *testing.T, a []int, cold []int, hot []int, expect int) {
	res := solve(a, cold, hot)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 2}
	cold := []int{3, 2}
	hot := []int{2, 1}
	expect := 6
	runSample(t, a, cold, hot, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 1, 2}
	cold := []int{5, 3}
	hot := []int{2, 1}
	expect := 11
	runSample(t, a, cold, hot, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 3, 1}
	cold := []int{100, 100, 100}
	hot := []int{1, 1, 1}
	expect := 301
	runSample(t, a, cold, hot, expect)
}

func TestSample4(t *testing.T) {
	a := []int{2, 1, 2, 1, 1}
	cold := []int{65, 45}
	hot := []int{54, 7}
	expect := 225
	runSample(t, a, cold, hot, expect)
}
