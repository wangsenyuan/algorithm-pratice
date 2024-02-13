package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	a, b := solve(s)
	res := []int{a, b}
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "10010"
	expect := []int{1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "111"
	expect := []int{2, 6}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "0101"
	expect := []int{0, 1}
	runSample(t, s, expect)
}
