package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a, b uint, expect []uint) {
	res := solve(a, b)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var a, b uint = 142, 76
	expect := []uint{33, 109}
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	var a, b uint = 12608, 0
	expect := []uint{6304, 6304}
	runSample(t, a, b, expect)
}
