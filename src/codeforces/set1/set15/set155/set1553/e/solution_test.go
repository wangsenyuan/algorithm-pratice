package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, p []int, m int, expect []int) {
	res := solve(p, m)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{2, 3, 1 ,4}
	m := 1
	expect := []int{3}
	runSample(t, p, m, expect)
}
