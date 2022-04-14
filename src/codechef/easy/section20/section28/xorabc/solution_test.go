package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, X int, expect bool) {
	res := solve(X)

	if len(res) > 0 != expect {
		t.Errorf("Sample expect %t, but got %v", expect, res)
	} else if len(res) > 0 {
		sort.Ints(res)
		if res[0] == res[1] || res[1] == res[2] {
			t.Errorf("Sample result %v, not correct", res)
		} else {
			y := (res[0] ^ res[1]) + (res[1] ^ res[2]) + (res[0] ^ res[2])
			if X != y {
				t.Errorf("Sample result %v, xor sum %d, not correct", res, y)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, false)
}

func TestSample2(t *testing.T) {
	runSample(t, 6, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 20, true)
}
