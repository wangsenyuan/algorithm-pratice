package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, x string, expect []int) {
	a, b := solve(s, x)
	res := []int{a, b}
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abababacababa"
	x := "aba"
	expect := []int{2, 2}
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "ddddddd"
	x := "dddd"
	expect := []int{1, 4}
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "xyzxyz"
	x := "xyz"
	expect := []int{2, 1}
	runSample(t, s, x, expect)
}

func TestSample4(t *testing.T) {
	s := "aaaaaaaa"
	x := "aa"
	expect := []int{3, 6}
	runSample(t, s, x, expect)
}

func TestSample5(t *testing.T) {
	s := "abababacabababababababacababababacacababababababababacababababababacabacabababababababacabababababababababacabababacabababababababacababababababacabacababababacababababacabacababababababacabababababababacababababacabababacabababacabacababababacacabababababababababacababababababacababababababacababababababababacacababababacabababababababacabababababababacababababababacabababababababababacacababacababacabababacababacababacababababababacababababacababaca"
	x := "aba"
	expect := []int{75, 242775700}
	runSample(t, s, x, expect)
}
