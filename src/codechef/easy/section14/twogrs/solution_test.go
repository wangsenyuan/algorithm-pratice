package main

import "testing"

func runSample(t *testing.T, arr []int, expect bool) {
	res := solve(arr)

	if res != expect {
		t.Errorf("Sample %v, expect %v, but got %v", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {

	arr := []int{3, 3, 3}
	expect := true
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {

	arr := []int{3, 2, 3}
	expect := true
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {

	arr := []int{1, 1, 2}
	expect := false
	runSample(t, arr, expect)
}

func TestSample4(t *testing.T) {

	arr := []int{1, 0, 1}
	expect := false
	runSample(t, arr, expect)
}

func TestSample5(t *testing.T) {

	arr := []int{1, 1, 1}
	expect := true
	runSample(t, arr, expect)
}
