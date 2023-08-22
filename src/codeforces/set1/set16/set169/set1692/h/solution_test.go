package main

import "testing"

func runSample(t *testing.T, arr []int, expect []int) {
	res := solve(arr)

	x := check(arr, expect)
	y := check(arr, res)
	if x != y {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func check(arr []int, ans []int) int {
	l, r := ans[1]-1, ans[2]-1
	var cnt int
	for i := l; i <= r; i++ {
		if arr[i] == ans[0] {
			cnt++
		} else {
			cnt--
		}
	}
	return cnt
}

func TestSample1(t *testing.T) {
	arr := []int{4, 4, 3, 4, 4}
	expect := []int{4, 1, 5}
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{11, 1, 11, 1, 11}
	expect := []int{1, 2, 2}
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{8, 8, 8, 9, 9, 6, 6, 9, 6, 6}
	expect := []int{6, 6, 10}
	runSample(t, arr, expect)
}
