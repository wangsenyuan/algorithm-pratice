package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	res := solve(s)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1"
	expect := []int{-1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "01"
	expect := []int{1, -1}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "010"
	expect := []int{0, 1, -1}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "10101"
	expect := []int{1, 3, -1, -1, -1}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "0000111"
	expect := []int{3, 6, 9, 12, -1, -1, -1}
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "001011000110"
	expect := []int{0, 2, 4, 6, 10, 15, 20, -1, -1, -1, -1, -1}
	runSample(t, s, expect)
}
