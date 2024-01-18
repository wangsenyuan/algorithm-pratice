package main

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := solve(arr)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{-3, 1, 4, -1, 5, -9}
	expect := 9
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{998244353, 998244353, 998244353, 998244353, 998244353}
	expect := 2994733059
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{-2718}
	expect := -2718
	runSample(t, arr, expect)
}
