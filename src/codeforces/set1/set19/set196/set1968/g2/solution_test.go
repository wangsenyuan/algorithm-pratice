package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, l int, r int, expect []int) {
	res := solve(s, l, r)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aba"
	l, r := 1, 3
	expect := []int{3, 1, 0}
	runSample(t, s, l, r, expect)
}

func TestSample2(t *testing.T) {
	s := "aaa"
	l, r := 2, 3
	expect := []int{1, 1}
	runSample(t, s, l, r, expect)
}

func TestSample3(t *testing.T) {
	s := "abacaba"
	l, r := 1, 5
	expect := []int{7, 3, 1, 1, 0}
	runSample(t, s, l, r, expect)
}

func TestSample4(t *testing.T) {
	s := "abababcab"
	l, r := 1, 6
	expect := []int{9, 2, 2, 2, 0, 0}
	runSample(t, s, l, r, expect)
}

func TestSample5(t *testing.T) {
	s := "kaaa"
	l, r := 1, 4
	expect := []int{4, 0, 0, 0}
	runSample(t, s, l, r, expect)
}
