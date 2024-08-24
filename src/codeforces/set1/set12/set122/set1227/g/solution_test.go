package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, a []int) {
	res := solve(a)

	n := len(a)
	for j, set := range res {
		for i := 0; i < n; i++ {
			if set[i] == '1' {
				if a[i] == 0 {
					t.Fatalf("Sample result %v, try to decrease 0 value at %d %d", res, j, i)
				}
				a[i]--
			}
		}
	}

	for i := 0; i < n; i++ {
		if a[i] != 0 {
			t.Fatalf("Sample result %v, getting some non-zero value at %d", res, i)
		}
	}

	sort.Strings(res)

	for i := 1; i < len(res); i++ {
		if res[i] == res[i-1] {
			t.Fatalf("Sample result %v, having same set", res)
		}
	}
}

func TestSample1(t *testing.T) {
	a := []int{5, 5, 5, 5, 5}
	runSample(t, a)
}

func TestSample2(t *testing.T) {
	a := []int{5, 1, 1, 1, 1}
	runSample(t, a)
}

func TestSample3(t *testing.T) {
	a := []int{4, 1, 5, 3, 4}
	runSample(t, a)
}
