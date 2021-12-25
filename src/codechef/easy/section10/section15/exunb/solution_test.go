package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, expect []string) {
	res := solve(n)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d, expect %v, but got %v", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	expect := []string{
		"010",
		"001",
		"100",
	}
	runSample(t, n, expect)
}
