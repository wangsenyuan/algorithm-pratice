package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, m int, g []string, expect []int64) {
	G := make([][]byte, n)
	for i := 0; i < n; i++ {
		G[i] = []byte(g[i])
	}
	res := solve(n, m, G)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v, expect %v, but got %v", n, m, g, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 4
	g := []string{
		"0011",
		"0000",
		"0100",
	}
	expect := []int64{1, 0, 1, 1, 0}
	runSample(t, n, m, g, expect)
}
