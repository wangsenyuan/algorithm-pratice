package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, c int, a []int, expect []int) {
	res := solve(c, a)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	c := 1
	a := []int{2, 0, 3}
	expect := []int{0, 1, 2}
	runSample(t, c, a, expect)
}

func TestSample2(t *testing.T) {
	c := 3
	a := []int{0, 10}
	expect := []int{1, 0}
	runSample(t, c, a, expect)
}

func TestSample3(t *testing.T) {
	c := 3
	a := []int{5, 4, 3, 2, 1}
	expect := []int{0, 1, 2, 3, 4}
	runSample(t, c, a, expect)
}

func TestSample4(t *testing.T) {
	c := 5
	a := []int{3, 10, 7, 1}
	expect := []int{1, 0, 2, 3}
	runSample(t, c, a, expect)
}

func TestSample5(t *testing.T) {
	c := 0
	a := []int{2, 2, 2, 3, 3, 3}
	expect := []int{1, 1, 2, 0, 4, 5}
	runSample(t, c, a, expect)
}
