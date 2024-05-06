package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, p []int, expect []int) {
	res := solve(n, p)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	p := []int{3, 2, 1}
	expect := []int{3, 2, 1, 5, 4}
	runSample(t, n, p, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	p := []int{2, 3, 1}
	runSample(t, n, p, nil)
}

func TestSample3(t *testing.T) {
	n := 5
	p := []int{3}
	expect := []int{3, 2, 1, 5, 4}
	runSample(t, n, p, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	p := []int{3, 4}
	runSample(t, n, p, nil)
}

func TestSample5(t *testing.T) {
	n := 5
	p := []int{3, 4}
	runSample(t, n, p, nil)
}

func TestSample6(t *testing.T) {
	n := 10
	p := []int{2, 1, 4, 3, 6, 5}
	expect := []int{2, 1, 4, 3, 6, 5, 10, 9, 8, 7}
	runSample(t, n, p, expect)
}
