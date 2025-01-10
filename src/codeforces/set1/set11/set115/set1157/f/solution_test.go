package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if len(res) != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	for i := 0; i < len(res); i++ {
		j := (i + 1) % len(res)
		if abs(res[i]-res[j]) > 1 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

	sort.Ints(a)
	sort.Ints(res)

	for i, j := 0, 0; i < len(res); i++ {
		for j < len(a) && a[j] != res[i] {
			j++
		}
		if j == len(a) {
			t.Fatalf("Sample result %v, not correct", res)
		}
		j++
	}
}

func abs(num int) int {
	return max(num, -num)
}

func TestSample1(t *testing.T) {
	a := []int{4, 3, 5, 1, 2, 2, 1}
	expect := 5
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 7, 5, 1, 5}
	expect := 2
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{5, 1, 4}
	expect := 2
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{2, 2, 3, 2, 1, 2, 2}
	expect := 7
	runSample(t, a, expect)
}
