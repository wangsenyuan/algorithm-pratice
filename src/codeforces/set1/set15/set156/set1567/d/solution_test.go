package main

import (
	"testing"
)

func runSample(t *testing.T, s int, n int, expect []int) {
	res := solve(s, n)

	x := calc(expect)
	y := calc(res)

	if x != y {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func calc(arr []int) int {
	var res int

	// cur is in base 11
	for i := 0; i < len(arr); i++ {
		res += getValueInBase11(arr[i])
	}

	return res
}

func getValueInBase11(num int) int {
	base := 1
	var res int
	for num > 0 {
		r := num % 10
		res = res*base + r
		base *= 11
		num /= 10
	}
	return res
}

func TestSample1(t *testing.T) {
	s, n := 97, 2
	expect := []int{70, 27}
	runSample(t, s, n, expect)
}

func TestSample2(t *testing.T) {
	s, n := 111, 4
	expect := []int{1, 1, 9, 100}
	runSample(t, s, n, expect)
}
