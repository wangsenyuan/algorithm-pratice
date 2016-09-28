package main

import "testing"

func TestValidTree(t *testing.T) {
	edges := [][]int{[]int{0, 1}, []int{0, 2}, []int{0, 3}, []int{1, 4}}
	if !validTree(5, edges) {
		t.Errorf("%v should be a valid tree", edges)
	}
}

func TestInValidTree(t *testing.T) {
	edges := [][]int{[]int{0, 1}, []int{1, 2}, []int{2, 3}, []int{1, 3}, []int{1, 4}}
	if validTree(5, edges) {
		t.Errorf("%v should not be a valid tree", edges)
	}
}
