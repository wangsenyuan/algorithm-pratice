package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, m int, x []int, direction string, expect []int) {
	res := solve(m, x, direction)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 12
	x := []int{1, 2, 3, 4, 9, 10, 11}
	direction := "R R L L R R R"
	expect := []int{1, 1, 1, 1, 2, -1, 2}
	runSample(t, m, x, direction, expect)
}

func TestSample2(t *testing.T) {
	m := 10
	x := []int{1, 6}
	direction := "R R"
	expect := []int{-1, -1}
	runSample(t, m, x, direction, expect)
}

func TestSample3(t *testing.T) {
	m := 10
	x := []int{1, 3}
	direction := "L L"
	expect := []int{2, 2}
	runSample(t, m, x, direction, expect)
}

func TestSample4(t *testing.T) {
	m := 8
	x := []int{6, 1, 7, 2, 3, 5, 4}
	direction := "R L R L L L L"
	expect := []int{-1, 2, 7, 3, 2, 7, 3}
	runSample(t, m, x, direction, expect)
}
