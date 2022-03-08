package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := solve(A)

	n := len(A)

	if expect != (res > n) {
		t.Fatalf("Sample expect %t, but got %d", expect, res)
	} else if expect {
		mem := make(map[int]bool)
		for _, a := range A {
			x := a % res
			if x == 0 || x > n {
				t.Fatalf("Sample result %d not correct", res)
			}
			mem[x] = true
		}
		if len(mem) != n {
			t.Fatalf("Sample result %d not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 7, 1}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 2}
	expect := false
	runSample(t, A, expect)
}
