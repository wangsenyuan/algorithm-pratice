package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, b []int) {
	a := solve(n, b)
	var c []int
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			c = append(c, min(a[i], a[j]))
		}
	}

	sort.Ints(b)
	sort.Ints(c)

	if !reflect.DeepEqual(b, c) {
		t.Fatalf("Sample result %v, not correct", a)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	b := []int{1, 3, 1}
	runSample(t, n, b)
}

func TestSample2(t *testing.T) {
	n := 2
	b := []int{10}
	runSample(t, n, b)
}

func TestSample3(t *testing.T) {
	n := 4
	b := []int{7, 5, 3, 5, 3, 3}
	runSample(t, n, b)
}

func TestSample4(t *testing.T) {
	n := 5
	b := []int{3, 0, 0, -2, 0, -2, 0, 0, -2, -2}
	runSample(t, n, b)
}
