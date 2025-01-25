package main

import (
	"strings"
	"testing"
)

func runSample(t *testing.T, l []int, expect bool) {
	res := solve(l)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	// len(res) > 0
	for i := 0; i < len(l); i++ {
		if len(res[i]) != l[i] {
			t.Fatalf("Sample result %v not correct", res)
		}
		for j := 0; j < len(l); j++ {
			if i == j {
				continue
			}
			if strings.HasPrefix(res[j], res[i]) {
				t.Fatalf("Sample result %v not correct, %d is prefix of %d", res, i, j)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	l := []int{4, 4, 4, 4, 4, 4, 4, 4, 4, 4}
	expect := true
	/**
			0000
	0001
	0010
	0011
	0100
	0101
	0110
	0111
	1000
	1001
	*/
	runSample(t, l, expect)
}
