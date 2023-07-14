package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A, B []int, expect []int) {
	res := solve(A, B)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{0, 1, 1}
	B := []int{1, 1, 0}
	expect := []int{0, 1}
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{5, 5, 5}
	B := []int{5, 5, 5}
	expect := []int{0, 15}
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{479, 178, 29}
	B := []int{11, 145, 530}
	expect := []int{22, 334}
	runSample(t, A, B, expect)
}

func TestSample4(t *testing.T) {
	A := []int{0, 0, 3}
	B := []int{3, 0, 0}
	expect := []int{3, 3}
	runSample(t, A, B, expect)
}

func TestSample5(t *testing.T) {
	A := []int{10, 53, 256}
	B := []int{182, 103, 34}
	expect := []int{119, 226}
	runSample(t, A, B, expect)
}
