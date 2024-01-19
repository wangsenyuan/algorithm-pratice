package main

import "testing"

func runSample(t *testing.T, s string, m int, l string, r string, expect bool) {
	res := solve(s, m, l, r)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "88005553535123456"
	m := 2
	l := "50"
	r := "56"
	expect := true
	runSample(t, s, m, l, r, expect)
}

func TestSample2(t *testing.T) {
	s := "880055535351234560"
	m := 2
	l := "50"
	r := "56"
	expect := false
	runSample(t, s, m, l, r, expect)
}

func TestSample3(t *testing.T) {
	s := "123412341234"
	m := 3
	l := "111"
	r := "444"
	expect := false
	runSample(t, s, m, l, r, expect)
}

func TestSample4(t *testing.T) {
	s := "1234"
	m := 4
	l := "4321"
	r := "4321"
	expect := true
	runSample(t, s, m, l, r, expect)
}

func TestSample5(t *testing.T) {
	s := "00010"
	m := 2
	l := "10"
	r := "11"
	expect := true
	runSample(t, s, m, l, r, expect)
}
