package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, bulbs []int, expect []int) {
	sz, ways := solve(n, bulbs)

	res := []int{sz, ways}

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	bulbs := []int{2, 2, 1, 1}
	expect := []int{2, 4}
	runSample(t, n, bulbs, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	bulbs := []int{1, 2, 2, 1}
	expect := []int{1, 2}
	runSample(t, n, bulbs, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	bulbs := []int{1, 2, 1, 2}
	expect := []int{1, 4}
	runSample(t, n, bulbs, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	bulbs := []int{3, 4, 4, 5, 3, 1, 1, 5, 2, 2}
	expect := []int{2, 8}
	runSample(t, n, bulbs, expect)
}

// 4 1 2 3 4 2 1 3 5 5
func TestSample5(t *testing.T) {
	n := 5
	bulbs := []int{4, 1, 2, 3, 4, 2, 1, 3, 5, 5}
	expect := []int{2, 16}
	runSample(t, n, bulbs, expect)
}
