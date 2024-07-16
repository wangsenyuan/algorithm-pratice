package main

import "testing"

func runSample(t *testing.T, n string, m int, expect int) {
	res := solve(n, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := "1912"
	m := 1
	expect := 5
	runSample(t, n, m, expect)
}

func TestSample2(t *testing.T) {
	n := "5"
	m := 6
	expect := 2
	runSample(t, n, m, expect)
}

func TestSample3(t *testing.T) {
	n := "999"
	m := 1
	expect := 6
	runSample(t, n, m, expect)
}

func TestSample4(t *testing.T) {
	n := "88"
	m := 2
	expect := 4
	runSample(t, n, m, expect)
}

func TestSample5(t *testing.T) {
	n := "12"
	m := 100
	expect := 2115
	runSample(t, n, m, expect)
}
