package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, balls []int, expect []int) {
	res := solve(balls)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	balls := []int{2, 3, 1, 4, 6, 5}
	expect := []int{3, 2, 2, 2, 2, 2}
	runSample(t, balls, expect)
}
