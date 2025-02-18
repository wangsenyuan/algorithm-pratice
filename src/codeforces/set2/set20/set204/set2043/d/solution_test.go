package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, l int, r int, g int, expect []int) {
	res := solve(l, r, g)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	l, r, g := 3, 3, 1
	expect := []int{-1, -1}
	runSample(t, l, r, g, expect)
}


func TestSample2(t *testing.T) {
	l, r, g := 3, 3, 2
	expect := []int{-1, -1}
	runSample(t, l, r, g, expect)
}