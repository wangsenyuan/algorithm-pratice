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
	s := "(RaRbR)L)L("
	// (
	// (a
	// (ab
	// (ab)
	// (a))
	// (())
	expect := []int{-1, -1, -1, -1, -1, -1, 1, 1, -1, -1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "(R)R(R)Ra)c"
	expect := []int{-1, -1, 1, 1, -1, -1, 1, 1, 1, -1, 1}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "(R)R(R)Ra)c"
	expect := []int{-1, -1, 1, 1, -1, -1, 1, 1, 1, -1, 1}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := ")))aaa(R(RaR))R(R(a)))RRaLaa)LRaa(aR))(LaRR(a(a)LaR()R(RR)(RaRa(()aaLLLLL(a)(()R(()"
	expect := bruteForce(s)
	runSample(t, s, expect)
}
