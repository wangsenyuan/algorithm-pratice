package main

import "testing"

func runSample(t *testing.T, a int, b int, p []int, expect bool) {
	ok, res := solve(a, b, p)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t, %v", expect, ok, res)
	}

	if !expect {
		return
	}

	pos := make(map[int]int)

	for i, x := range p {
		pos[x] = i
	}

	for i := 0; i < len(p); i++ {
		if j, ok := pos[a-p[i]]; ok {
			if res[j] == 0 && res[i] == 1 {
				t.Fatalf("Sample result %v not correct, %d %d should be both in A, but not", res, p[i], a-p[i])
			}
		}
		if j, ok := pos[b-p[i]]; ok {
			if res[j] == 1 && res[i] == 0 {
				t.Fatalf("Sample result %v not correct, %d %d should be both in B, but not", res, p[i], b-p[i])
			}
		}
	}
}

func TestSample1(t *testing.T) {
	a, b := 5, 9
	p := []int{2, 3, 4, 5}
	expect := true
	runSample(t, a, b, p, expect)
}

func TestSample2(t *testing.T) {
	a, b := 3, 4
	p := []int{1, 2, 4}
	expect := false
	runSample(t, a, b, p, expect)
}
