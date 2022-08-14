package main

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := int(solve(arr))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{3, 4}
	expect := 6
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{4, 3, 4}
	expect := 16
	runSample(t, arr, expect)
}
