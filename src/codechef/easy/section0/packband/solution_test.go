package main

import "testing"

func TestSample(t *testing.T) {
	ls := []int{10, 20, 34, 55}
	bands := [][]int{
		{7, 14},
		{7, 21},
		{14, 21},
		{7, 35},
	}

	res := solve(len(ls), ls, len(bands), bands)

	if res != 2 {
		t.Errorf("this sample should give answer 2, but got %d\n", res)
	}
}
