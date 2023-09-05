package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, expect []int64) {
	res := solve(s)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "LLR"
	expect := []int64{3, 5, 5}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "LRRLL"
	expect := []int64{16, 16, 16, 16, 16}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "LRRRLLLRLLRL"
	expect := []int64{86, 95, 98, 101, 102, 102, 102, 102, 102, 102, 102, 102}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "LLLLLRRRRR"
	expect := []int64{29, 38, 45, 52, 57, 62, 65, 68, 69, 70}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "LRLRLRLRL"
	expect := []int64{44, 50, 54, 56, 56, 56, 56, 56, 56}
	runSample(t, s, expect)
}
