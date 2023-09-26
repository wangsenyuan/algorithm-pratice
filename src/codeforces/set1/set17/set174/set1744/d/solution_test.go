package main

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := solve(arr)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{2}
	expect := 0
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{3, 2}
	expect := 1
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{10, 6, 11}
	expect := 1
	runSample(t, arr, expect)
}

func TestSample4(t *testing.T) {
	arr := []int{13, 17, 1, 1}
	expect := -1
	runSample(t, arr, expect)
}

func TestSample5(t *testing.T) {
	arr := []int{1, 1, 12, 1, 1}
	expect := 2
	runSample(t, arr, expect)
}

func TestSample6(t *testing.T) {
	arr := []int{20, 7, 14, 18, 3, 5}
	expect := 1
	runSample(t, arr, expect)
}
