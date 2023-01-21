package main

import "testing"

func runSample(t *testing.T, s, f string, L, R []int, expect bool) {
	res := solve(len(s), s, f, L, R)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "00000"
	f := "00111"
	L := []int{1, 1}
	R := []int{5, 3}
	expect := true
	runSample(t, s, f, L, R, expect)
}

func TestSample2(t *testing.T) {
	s := "00"
	f := "01"
	L := []int{1}
	R := []int{2}
	expect := false
	runSample(t, s, f, L, R, expect)
}

func TestSample3(t *testing.T) {
	s := "1111111111"
	f := "0110001110"
	L := []int{1, 5, 7, 1, 3, 6}
	R := []int{10, 9, 10, 7, 5, 10}
	expect := true
	runSample(t, s, f, L, R, expect)
}

func TestSample4(t *testing.T) {
	s := "10000"
	f := "11000"
	L := []int{2, 5}
	R := []int{1, 3}
	expect := false
	runSample(t, s, f, L, R, expect)
}
