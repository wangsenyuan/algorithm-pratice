package main

import "testing"

func runSample(t *testing.T, a int64, k int, expect string) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var a int64 = 1990
	k := 1
	expect := "9190"
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	var a int64 = 1034
	k := 2
	expect := "3104"
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	var a int64 = 9090000078001234
	k := 6
	expect := "9907000008001234"
	runSample(t, a, k, expect)
}
