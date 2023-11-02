package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, p int, expect []int) {
	res := solve(p)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := 3
	expect := []int{3, 0}
	runSample(t, p, expect)
}

func TestSample2(t *testing.T) {
	p := 4
	expect := []int{5, 6}
	runSample(t, p, expect)
}

func TestSample3(t *testing.T) {
	p := 12
	expect := []int{8, 16}
	runSample(t, p, expect)
}
