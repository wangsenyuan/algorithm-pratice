package main

import (
	"reflect"
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
	mark := make([]bool, len(a))
	for i := 0; i < len(res); i++ {
		mark[res[i]-1] = true
	}
	var r []int
	for i := 0; i < len(a); i++ {
		if !mark[i] {
			r = append(r, a[i])
		}
	}

	sort.Ints(res)
	sort.Ints(r)

	if !reflect.DeepEqual(res, r) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 4, 2, 2, 3}
	expect := true
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3}
	expect := false
	runSample(t, a, expect)
}
