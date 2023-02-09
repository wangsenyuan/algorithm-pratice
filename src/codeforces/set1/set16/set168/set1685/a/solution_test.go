package main

import "testing"

func runSample(t *testing.T, arr []int) {
	res := solve(arr)
	if len(res) == 0 {
		t.Fatalf("Sample result not correct")
	}
	n := len(res)
	x := res[0] < res[1]
	for i := 1; i+1 < n; i++ {
		if res[i] < res[i+1] == x {
			t.Fatalf("Sample result %v, not correct", res)
		}
		x = !x
	}
	if res[n-1] < res[0] == x {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 9, 8, 4}
	runSample(t, arr)
}

func TestSample2(t *testing.T) {
	arr := []int{1, 11, 1, 111, 1, 1111}
	runSample(t, arr)
}
