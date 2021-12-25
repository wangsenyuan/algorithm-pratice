package main

import (
	"math/rand"
	"testing"
)

func runSample(t *testing.T, n int, B []int) {
	ask := func(arr []int, x int) int {
		if len(arr) == 0 {
			return 1
		}
		freq := make(map[int]int)
		for _, i := range arr {
			freq[B[i-1]]++
		}
		freq[B[x-1]]++
		var res int
		for _, v := range freq {
			if res < v {
				res = v
			}
		}
		return res
	}

	A := solve(n, ask)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (B[i] == B[j]) != (A[i] == A[j]) {
				t.Fatalf("sampel result %v, not correct vs %v", A, B)
			}
			if (B[i] != B[j]) != (A[i] != A[j]) {
				t.Fatalf("sampel result %v, not correct vs %v", A, B)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	n := 3
	B := []int{10, 20, 10}
	runSample(t, n, B)
}

func genRandArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = int(rand.Int31n(20) + 1)
	}
	return arr
}

func TestSample2(t *testing.T) {
	for i := 2; i < 100; i++ {
		arr := genRandArray(i)
		runSample(t, i, arr)
	}
}

func TestSample3(t *testing.T) {
	B := []int{4, 17, 1, 4, 14, 18, 14, 2, 20, 14, 4, 12, 3}
	runSample(t, len(B), B)
}
