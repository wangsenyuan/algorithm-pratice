package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, students [][]int, expect []int) {
	res := solve(students)

	a := getExpectValue(students, expect)
	b := getExpectValue(students, res)

	if math.Abs(a-b) > 1e-7 {
		t.Fatalf("Sample expect %.8f, but got %.8f from %v", a, b, res)
	}
}

func getExpectValue(students [][]int, res []int) float64 {
	mem := make(map[int]bool)

	for _, k := range res {
		mem[k] = true
	}
	var tmp float64
	for _, st := range students {
		m, k := st[0], st[1]
		if mem[m] {
			if k >= len(res) {
				tmp++
			} else {
				tmp += float64(k) / float64(len(res))
			}
		}
	}

	return tmp
}

func TestSample1(t *testing.T) {
	students := [][]int{
		{13, 2},
		{42, 2},
		{37, 2},
	}
	expect := []int{42, 13, 37}
	runSample(t, students, expect)
}

func TestSample2(t *testing.T) {
	students := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
	}
	expect := []int{2, 3, 4}
	runSample(t, students, expect)
}

func TestSample3(t *testing.T) {
	students := [][]int{
		{10, 1},
		{5, 2},
		{10, 1},
	}
	expect := []int{10}
	runSample(t, students, expect)
}

func TestSample4(t *testing.T) {
	students := [][]int{
		{10, 1},
		{10, 2},
		{5, 2},
	}
	expect := []int{5, 10}
	runSample(t, students, expect)
}
