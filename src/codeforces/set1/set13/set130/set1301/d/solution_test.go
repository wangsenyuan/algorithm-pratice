package main

import (
	"strconv"
	"testing"
)

func runSample(t *testing.T, n int, m int, k int, expect bool) {
	res := solve(n, m, k)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	var sum int
	for _, x := range res {
		var i int
		for x[i] != ' ' {
			i++
		}
		cnt, _ := strconv.Atoi(x[:i])
		mov := x[i+1 : len(x)-1]
		sum += cnt * len(mov)
	}

	if sum != k {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 3, 8, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 4, 16, true)
}