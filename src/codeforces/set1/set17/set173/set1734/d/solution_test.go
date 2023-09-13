package main

import "testing"

func runSample(t *testing.T, arr []int, k int, expect bool) {
	res := solve(arr, k)
	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{-1, -2, -3, 6, -2, -3, -1}
	k := 4
	expect := true
	runSample(t, arr, k, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{232, -500, -700}
	k := 1
	expect := true
	runSample(t, arr, k, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{-1, -2, -4, 6, -2, -4, -1}
	k := 4
	expect := false
	runSample(t, arr, k, expect)
}

func TestSample4(t *testing.T) {
	arr := []int{-100, 10, -7, 6, -2, -3, 6, -10}
	k := 4
	expect := true
	runSample(t, arr, k, expect)
}

func TestSample5(t *testing.T) {
	arr := []int{-999, 0, -2, 3, 4, 5, 6, 7}
	k := 2
	expect := false
	runSample(t, arr, k, expect)
}

func TestSample6(t *testing.T) {
	arr := []int{7, 3, 3, 4, 2, 1, 1}
	k := 3
	expect := true
	runSample(t, arr, k, expect)
}
