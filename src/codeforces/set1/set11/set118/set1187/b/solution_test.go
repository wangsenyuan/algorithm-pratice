package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, words []string, expect []int) {
	res := solve(s, words)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "arrayhead"
	words := []string{
		"arya",
		"harry",
		"ray",
		"r",
		"areahydra",
	}
	expect := []int{5, 6, 5, 2, 9}

	runSample(t, s, words, expect)
}
