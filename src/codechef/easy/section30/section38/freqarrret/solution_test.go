package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, B []int, expect []int) {
	res := solve(B)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	B := []int{2, 3, 3, 3, 2}
	expect := []int{1, 2, 2, 2, 1}
	runSample(t, B, expect)
}

func TestSample2(t *testing.T) {
	B := []int{1, 1, 1, 1, 1}
	expect := []int{1, 2, 3, 4, 5}
	runSample(t, B, expect)
}

func TestSample3(t *testing.T) {
	B := []int{5, 5, 5, 5, 5}
	expect := []int{1, 1, 1, 1, 1}
	runSample(t, B, expect)
}

func TestSample4(t *testing.T) {
	B := []int{1, 2, 4}
	var expect []int
	runSample(t, B, expect)
}

func TestSample5(t *testing.T) {
	B := []int{1, 3, 2, 3, 2, 2, 2, 3}
	expect := []int{1, 2, 3, 2, 3, 4, 4, 2}
	runSample(t, B, expect)
}
