package main

import "testing"

func runSample(t *testing.T, d int, k int, expect string) {
	res := solve(d, k)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	d, k := 2, 1
	expect := Utkarsh
	runSample(t, d, k, expect)
}

func TestSample2(t *testing.T) {
	d, k := 5, 2
	expect := Ashish
	runSample(t, d, k, expect)
}

func TestSample3(t *testing.T) {
	d, k := 10, 3
	expect := Utkarsh
	runSample(t, d, k, expect)
}

func TestSample4(t *testing.T) {
	d, k := 15441, 33
	expect := Ashish
	runSample(t, d, k, expect)
}
