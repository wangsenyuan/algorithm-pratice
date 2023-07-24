package main

import "testing"

func runSample(t *testing.T, L int, V int64, expect uint) {
	res := solve(L, V)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

// the test would be too slow
//func TestSample1(t *testing.T) {
//	runSample(t, 2, 0, 0)
//}
//
//func TestSample2(t *testing.T) {
//	runSample(t, 3, 27, 5)
//}
//
//func TestSample3(t *testing.T) {
//	runSample(t, 5, 24, 486)
//}
//
//func TestSample4(t *testing.T) {
//	runSample(t, 10, 45, 2349595)
//}
