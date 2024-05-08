package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	res := solve(s)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	var arr []int
	for _, cur := range res {
		for j, i := range cur {
			x := int(s[i-1] - '0')
			if x != j&1 {
				t.Fatalf("Sample result %v, not correct", res)
			}
			arr = append(arr, i)
		}
	}

	if len(arr) != len(s) {
		t.Fatalf("Sample result %v, not correct", res)
	}
	sort.Ints(arr)
	for i := 1; i < len(s); i++ {
		if arr[i] == arr[i-1] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := "0010100"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "111"
	expect := false
	runSample(t, s, expect)
}
