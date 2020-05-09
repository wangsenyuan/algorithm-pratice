package main

import "testing"

func runSample(t *testing.T, arr []int, expect bool) {
	res := solve(len(arr), arr)

	if res != expect {
		t.Errorf("Sample %v, expect %t, but got %t", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{14}
	runSample(t, arr, true)
}

func TestSample2(t *testing.T) {
	arr := []int{1, -1}
	runSample(t, arr, true)
}

func TestSample3(t *testing.T) {
	arr := []int{5, 5, 5, 1}
	runSample(t, arr, true)
}

func TestSample4(t *testing.T) {
	arr := []int{3, 2, 1}
	runSample(t, arr, false)
}
func TestSample5(t *testing.T) {
	arr := []int{0, 1}
	runSample(t, arr, false)
}

func TestSample6(t *testing.T) {
	arr := []int{-239, -2, -100, -3, -11}
	runSample(t, arr, true)
}
