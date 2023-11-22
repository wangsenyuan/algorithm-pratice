package main

import "testing"

func runSample(t *testing.T, l int, r int, expect int) {
	res := solve(l, r)

	if score(expect) != score(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
	if res < l || res > r {
		t.Fatalf("Sample result %d, out of range [%d - %d]", res, l, r)
	}
}

func score(num int) int {
	ds := toDigits(num)
	a, b := ds[0], ds[0]

	for i := 1; i < len(ds); i++ {
		a = max(a, ds[i])
		b = min(b, ds[i])
	}
	return a - b
}

func TestSample1(t *testing.T) {
	l, r := 59, 63
	expect := 63
	runSample(t, l, r, expect)
}

func TestSample2(t *testing.T) {
	l, r := 42, 49
	expect := 44
	runSample(t, l, r, expect)
}

func TestSample3(t *testing.T) {
	l, r := 48, 53
	expect := 53
	runSample(t, l, r, expect)
}

func TestSample4(t *testing.T) {
	l, r := 1, 100
	expect := 1
	runSample(t, l, r, expect)
}
