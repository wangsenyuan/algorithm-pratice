package main

import "testing"

func runSample(t *testing.T, n int, c int, queries []int) {
	s := NewSolver(n, c)

	sheets := make([]int, n+1)

	for _, x := range queries {
		ok, i := s.Play(x)
		sheets[i] = x
		if ok {
			break
		}
	}

	for i := 1; i <= n; i++ {
		if sheets[i] == 0 || sheets[i] < sheets[i-1] {
			t.Fatalf("Sample result not correct")
		}
	}
}

func TestSample1(t *testing.T) {
	n, c := 2, 4
	queries := []int{2, 1, 3}
	runSample(t, n, c, queries)
}

func TestSample2(t *testing.T) {
	n, c := 2, 2
	queries := []int{1, 2}
	runSample(t, n, c, queries)
}
