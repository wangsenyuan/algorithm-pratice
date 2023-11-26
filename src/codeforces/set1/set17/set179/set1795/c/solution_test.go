package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, b []int, expect []int) {
	res := solve(a, b)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{10, 20, 15}
	b := []int{9, 8, 6}
	expect := []int{9, 9, 12}
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{5}
	b := []int{7}
	expect := []int{5}
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{13, 8, 5, 4}
	b := []int{3, 4, 2, 1}
	expect := []int{3, 8, 6, 4}
	runSample(t, a, b, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1000000000, 1000000000, 1000000000}
	b := []int{1, 1, 1000000000}
	expect := []int{1, 2, 2999999997}
	runSample(t, a, b, expect)
}

func TestSample5(t *testing.T) {
	a := []int{93, 40, 87, 66, 44, 45, 64, 55, 20, 82, 66, 77, 80, 32, 77, 30, 75, 72, 35, 18, 38, 21, 100, 16, 15, 97, 57, 89, 16, 28, 61, 85, 40, 40, 79, 17, 30, 35, 82, 68, 92, 11, 32, 62, 57, 29, 39, 22, 7, 47, 67, 90, 83, 22, 42, 13, 32, 48, 58}
	b := []int{28, 55, 100, 35, 54, 22, 41, 35, 30, 20, 86, 36, 33, 98, 17, 21, 9, 18, 28, 73, 80, 79, 87, 34, 7, 58, 56, 55, 31, 16, 34, 99, 11, 25, 5, 76, 92, 91, 39, 85, 28, 9, 87, 60, 15, 80, 65, 89, 90, 30, 87, 49, 50, 82, 56, 72, 33, 25, 95}
	expect := []int{28, 95, 97, 35, 75, 22, 64, 58, 40, 20, 128, 36, 66, 87, 17, 42, 27, 54, 96, 71, 38, 21, 87, 29, 7, 66, 95, 56, 47, 19, 46, 112, 11, 50, 14, 101, 30, 35, 39, 111, 28, 18, 89, 60, 17, 71, 39, 22, 7, 30, 84, 49, 91, 55, 42, 13, 32, 25, 81}
	runSample(t, a, b, expect)
}
