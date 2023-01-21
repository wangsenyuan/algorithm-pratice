package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a, b string, expect []string) {
	res := solve(a, b)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, b := "a8", "h1"
	expect := []string{"RD", "RD", "RD", "RD", "RD", "RD", "RD"}
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a, b := "a7", "h1"
	expect := []string{"RD", "RD", "RD", "RD", "RD", "RD", "R"}
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a, b := "h1", "a8"
	expect := []string{"LU", "LU", "LU", "LU", "LU", "LU", "LU"}
	runSample(t, a, b, expect)
}
