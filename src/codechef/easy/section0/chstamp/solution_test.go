package main

import "testing"

func TestSample(t *testing.T) {
	n, m := 3, 4
	stamps := []int{1, 3, 5}
	offers := [][]int{
		{1, 4, 3},
		{1, 1, 3},
		{2, 6, 22},
		{3, 5, 8},
	}
	expect := int64(16)
	res := solve(n, stamps, m, offers)
	if res != expect {
		t.Errorf("sample %d %v %d %v, should give %d, but got %d", n, stamps, m, offers, expect, res)
	}
}
