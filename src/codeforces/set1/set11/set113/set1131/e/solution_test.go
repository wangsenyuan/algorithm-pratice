package main

import "testing"

func runSample(t *testing.T, p []string, expect int) {
	res := solve(p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []string{"a", "b", "a"}
	expect := 3
	runSample(t, p, expect)
}

func TestSample2(t *testing.T) {
	p := []string{"bnn", "a"}
	expect := 1
	runSample(t, p, expect)
}

func TestSample3(t *testing.T) {
	p := []string{"m", "g", "u", "t", "tttt"}
	expect := 9
	runSample(t, p, expect)
}

func TestSample4(t *testing.T) {
	p := []string{"r", "o", "w", "wwwwwwwwww"}
	expect := 21
	runSample(t, p, expect)
}


func TestSample5(t *testing.T) {
	p := []string{"r", "o", "w", "wwwwwwwwww"}
	expect := 21
	runSample(t, p, expect)
}