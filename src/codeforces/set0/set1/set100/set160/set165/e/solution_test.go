package main

import (
	"testing"
)

func runSample(t *testing.T, a []int, expect []int) {
	res := solve(a)

	for i, x := range res {
		if x < 0 {
			if expect[i] > 0 {
				t.Fatalf("Sample expect %v, but got %v", expect, res)
			}
			continue
		}
		if x&int32(a[i]) != 0 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	a := []int{90, 36}
	expect := []int{36, 90}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 6, 3, 6}
	expect := []int{-1, -1, -1, -1}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{10, 6, 9, 8, 2}
	expect := []int{-1, 8, 2, 2, 8}
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{4000000}
	expect := []int{-1}
	runSample(t, a, expect)
}

// func TestSample5(t *testing.T) {
// 	n := 10000
// 	x := 4_000_000
// 	arr := make([]int, n)
// 	rand := rand.New(rand.NewSource(47))

// 	for i := 0; i < n; i++ {
// 		arr[i] = rand.Intn(x) + 1
// 	}
// 	expect := []int{-1}
// 	runSample(t, arr, expect)
// }
