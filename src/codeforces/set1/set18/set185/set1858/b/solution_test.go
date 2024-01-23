package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, d int, s []int, expect []int) {
	res := solve(n, d, s)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, d := 6, 2
	s := []int{2, 5}
	expect := []int{3, 1}
	runSample(t, n, d, s, expect)
}

func TestSample2(t *testing.T) {
	n, d := 8, 2
	s := []int{3, 5, 8}
	expect := []int{4, 1}
	runSample(t, n, d, s, expect)
}

func TestSample3(t *testing.T) {
	n, d := 8, 9
	s := []int{2, 8, 9, 10}
	expect := []int{4, 4}
	runSample(t, n, d, s, expect)
}

func TestSample4(t *testing.T) {
	n, d := 30, 8
	s := []int{6, 8, 15, 24, 29}
	expect := []int{6, 4}
	runSample(t, n, d, s, expect)
}

func TestSample5(t *testing.T) {
	n, d := 30, 8
	s := []int{6, 8, 12, 20, 27}
	expect := []int{5, 2}
	runSample(t, n, d, s, expect)
}

func TestSample6(t *testing.T) {
	n, d := 8, 3
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expect := []int{7, 7}
	runSample(t, n, d, s, expect)
}

func TestSample7(t *testing.T) {
	n, d := 2, 2
	s := []int{1, 2}
	expect := []int{1, 1}
	runSample(t, n, d, s, expect)
}

func TestSample8(t *testing.T) {
	n, d := 1000000000, 20000000
	s := []int{57008429, 66778899, 837653445}
	expect := []int{51, 1}
	runSample(t, n, d, s, expect)
}
