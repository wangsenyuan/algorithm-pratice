package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, expect []int) {
	res := solve(n)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	expect := []int{1, 0, 1}
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	expect := []int{3, 2, 1}
	runSample(t, n, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	expect := []int{12, 7, 1}
	runSample(t, n, expect)
}

func TestSample4(t *testing.T) {
	n := 8
	expect := []int{42, 27, 1}
	runSample(t, n, expect)
}

func TestSample5(t *testing.T) {
	n := 60
	expect := []int{341102826, 248150916, 1}
	runSample(t, n, expect)
}
