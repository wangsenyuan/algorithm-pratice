package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	res := solve(s)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample %s expect %v, but got %v", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "><<"
	expect := []int{3, 6, 5}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "<<<<"
	expect := []int{1, 2, 3, 4}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "<><<<>"
	expect := []int{1, 4, 7, 10, 8, 1}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := ">><"
	expect := []int{5, 6, 3}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := ">><<"
	expect := []int{5, 10, 10, 5}
	runSample(t, s, expect)
}
