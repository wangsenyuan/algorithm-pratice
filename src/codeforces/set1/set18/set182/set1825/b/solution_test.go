package main

import "testing"

func runSample(t *testing.T, n int, m int, arr []int, expect int) {
	res := solve(n, m, arr)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 2
	arr := []int{1, 3, 1, 4}
	expect := 9
	runSample(t, n, m, arr, expect)
}

func TestSample2(t *testing.T) {
	n, m := 2, 2
	arr := []int{-1, -1, -1, -1}
	expect := 0
	runSample(t, n, m, arr, expect)
}

func TestSample3(t *testing.T) {
	n, m := 2, 3
	arr := []int{7, 8, 9, -3, 10, 8}
	expect := 64
	runSample(t, n, m, arr, expect)
}

func TestSample4(t *testing.T) {
	n, m := 3, 2
	arr := []int{4, 8, -3, 0, -7, 1}
	expect := 71
	runSample(t, n, m, arr, expect)
}

func TestSample5(t *testing.T) {
	n, m := 4, 3
	arr := []int{-32030, 59554, 16854, -85927, 68060, -64460, -79547, 90932, 85063, 82703, -12001, 38762}
	expect := 1933711
	runSample(t, n, m, arr, expect)
}
