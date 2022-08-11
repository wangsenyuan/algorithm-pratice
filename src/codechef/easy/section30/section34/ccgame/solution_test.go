package main

import "testing"

func runSample(t *testing.T, arr []int, expect string) {
	res := solve(arr)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 0, 1, 0}
	expect := CHEF
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expect := CHEF
	runSample(t, arr, expect)
}
