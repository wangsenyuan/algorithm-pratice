package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, k int, expect []string) {
	res := solve(n, k)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 2
	expect := []string{
		"R..",
		"...",
		"..R",
	}
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 40, 33

	runSample(t, n, k, nil)
}
