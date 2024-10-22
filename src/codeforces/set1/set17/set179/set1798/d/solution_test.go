package main

import (
	"reflect"
	"slices"
	"sort"
	"testing"
)

func runSample(t *testing.T, a []int, expect bool) {
	res := solve(a)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	ma := slices.Max(a)
	mn := slices.Min(a)
	n := len(a)
	for i := 0; i < n; i++ {
		var sum int
		for j := i; j < n; j++ {
			sum += res[j]
			if abs(sum) >= ma-mn {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}

	sort.Ints(a)
	sort.Ints(res)

	if !reflect.DeepEqual(res, a) {
		t.Fatalf("Sample result %v, not the original %v", res, a)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 4, -2, -5}
	expect := true
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 2, 2, -3, -3}
	expect := true
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{-3, -3, 1, 1, 1, 1, 1, 1}
	expect := true
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{0, 1, -1}
	expect := true
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{-3, 4, 3, 4, -4, -4, 0}
	expect := true
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{-18, 13, -18, -17, 12, 15, 13}
	expect := true
	runSample(t, a, expect)
}
