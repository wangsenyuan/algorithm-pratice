package main

import "testing"

func runSample(t *testing.T, arr []int, expect []int) {
	res := solve(arr)

	var a int
	var b int
	for i := 0; i < len(arr); i++ {
		a |= expect[i]
		b |= res[i]
		if a != b {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 2, 4, 8}
	expect := []int{8, 4, 2, 1}
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{5, 1, 2, 3, 4, 5, 5}
	expect := []int{5, 2, 1, 3, 4, 5, 5}
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{1, 101}
	expect := []int{101, 1}
	runSample(t, arr, expect)
}

func TestSample4(t *testing.T) {
	arr := []int{2, 3, 4, 2, 3, 4}
	expect := []int{4, 3, 2, 2, 3, 4}
	runSample(t, arr, expect)
}

func TestSample5(t *testing.T) {
	arr := []int{1, 4, 2, 3, 4, 5, 7, 1}
	expect := []int{7, 1, 4, 2, 3, 4, 5, 1}
	runSample(t, arr, expect)
}
