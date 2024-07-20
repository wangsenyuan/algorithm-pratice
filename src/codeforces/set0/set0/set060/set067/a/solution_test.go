package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, s string, expect []int) {
	res := solve(n, s)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	s := "LRLR"
	expect := []int{2, 1, 2, 1, 2}
	runSample(t, n, s, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	s := "=RRR"
	expect := []int{1, 1, 2, 3, 4}
	runSample(t, n, s, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	s := "RLRL="
	expect := []int{1, 2, 1, 2, 1, 1}
	runSample(t, n, s, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	s := "R="
	expect := []int{1, 2, 2}
	runSample(t, n, s, expect)
}

func TestSample5(t *testing.T) {
	n := 7
	s := "=RL=L="
	expect := []int{1, 1, 3, 2, 2, 1, 1}
	runSample(t, n, s, expect)
}
