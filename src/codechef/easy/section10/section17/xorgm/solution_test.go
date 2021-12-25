package main

import "testing"

func runSample(t *testing.T, n int, A, B []int, expect bool) {
	res := solve(n, A, B)

	if expect != (res >= 0) {
		t.Fatalf("sample %d %v %v, expect %t, but got %d", n, A, B, expect, res)
	}

	if expect {
		// A[i] ^ res = B[j]
		mem := make(map[int]int)

		for i := 0; i < n; i++ {
			mem[B[i]]++
		}

		for i := 0; i < n; i++ {
			x := A[i] ^ res
			if mem[x] < 0 {
				t.Fatalf("sample %d %v %v, expect %t, but got %d", n, A, B, expect, res)
			}
			mem[x]--
		}
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{3, 1, 2, 4, 5}
	B := []int{2, 4, 5, 1, 3}
	runSample(t, n, A, B, true)
}
