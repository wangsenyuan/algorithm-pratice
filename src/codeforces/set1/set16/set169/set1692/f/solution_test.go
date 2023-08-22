package main

import "testing"

func runSample(t *testing.T, arr []int, expect bool) {
	res := solve(arr)
	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{20, 22, 19, 84}
	expect := true
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{1, 11, 1, 2022}
	expect := true
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{1100, 1100, 1100, 1111}
	expect := false
	runSample(t, arr, expect)
}

func TestSample4(t *testing.T) {
	arr := []int{12, 34, 56, 78, 90}
	expect := false
	runSample(t, arr, expect)
}

func TestSample5(t *testing.T) {
	arr := []int{1, 9, 8, 4}
	expect := true
	runSample(t, arr, expect)
}

func TestSample6(t *testing.T) {
	arr := []int{16, 38, 94, 25, 18, 99}
	expect := true
	runSample(t, arr, expect)
}
