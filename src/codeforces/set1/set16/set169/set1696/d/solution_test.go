package main

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := solve(arr)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1}
	expect := 0
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{1, 2}
	expect := 1
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{1, 4, 2, 3, 5}
	expect := 1
	runSample(t, arr, expect)
}

func TestSample4(t *testing.T) {
	arr := []int{2, 1, 5, 3, 4}
	expect := 4
	runSample(t, arr, expect)
}

func TestSample5(t *testing.T) {
	arr := []int{7, 4, 8, 1, 6, 10, 3, 5, 2, 9}
	expect := 6
	runSample(t, arr, expect)
}
