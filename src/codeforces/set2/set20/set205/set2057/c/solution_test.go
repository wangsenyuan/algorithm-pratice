package main

import "testing"

func runSample(t *testing.T, l int, r int, expect []int) {
	res := solve(l, r)

	get := func(arr []int) int {
		for _, x := range arr {
			if x < l || x > r {
				t.Fatalf("Sample result %v, exceeds the range [%d %d]", arr, l, r)
			}
		}
		a, b, c := arr[0], arr[1], arr[2]
		if a == b || b == c || c == a {
			t.Fatalf("Sample result %v, having duplicates", arr)
		}
		return (a ^ b) + (b ^ c) + (c ^ a)
	}

	x := get(expect)
	y := get(res)

	if x != y {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	l, r := 0, 2
	expect := []int{1, 2, 0}
	runSample(t, l, r, expect)
}

func TestSample2(t *testing.T) {
	l, r := 0, 8
	expect := []int{8, 7, 1}
	runSample(t, l, r, expect)
}

func TestSample3(t *testing.T) {
	l, r := 1, 3
	expect := []int{3, 2, 1}
	runSample(t, l, r, expect)
}

func TestSample4(t *testing.T) {
	l, r := 6, 22
	expect := []int{7, 16, 11}
	runSample(t, l, r, expect)
}

func TestSample5(t *testing.T) {
	l, r := 128, 137
	expect := []int{134, 132, 137}
	runSample(t, l, r, expect)
}

func TestSample6(t *testing.T) {
	l, r := 69, 98
	expect := []int{98, 85, 76}
	runSample(t, l, r, expect)
}

func TestSample7(t *testing.T) {
	l, r := 115, 127
	expect := []int{123, 121, 118}
	runSample(t, l, r, expect)
}

func TestSample8(t *testing.T) {
	l, r := 0, 1073741823
	expect := []int{965321865, 375544086, 12551794}
	runSample(t, l, r, expect)
}

func TestSample9(t *testing.T) {
	l, r := 3, 6
	expect := []int{3, 4, 5}
	runSample(t, l, r, expect)
}
