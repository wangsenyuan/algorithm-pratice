package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, mat []string, s []int, expect []int) {
	res := solve(mat, s)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := []string{
		"1..",
		"...",
		"..2",
	}
	s := []int{1, 1}
	expect := []int{6, 3}
	runSample(t, mat, s, expect)
}

func TestSample2(t *testing.T) {
	mat := []string{
		"....",
		"#...",
		"1234",
	}
	s := []int{1, 1, 1, 1}
	expect := []int{1, 4, 3, 3}
	runSample(t, mat, s, expect)
}

func TestSample3(t *testing.T) {
	mat := []string{
		"....",
		"....",
		"..2.",
		"...1",
	}
	s := []int{1, 1000000000}
	expect := []int{3, 13}
	runSample(t, mat, s, expect)
}
