package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a, b, c int, expect []int) {
	res := solve(a, b, c)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, b, c := 931480234, -1767614767, -320146190
	expect := []int{-98880374013340920, -52107006370101410}
	runSample(t, a, b, c, expect)
}
