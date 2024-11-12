package main

import "testing"

func runSample(t *testing.T, s string, x string, expect bool) {
	ok, res := solve(s, x)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t, %v", expect, ok, res)
	}
	if !expect {
		return
	}
	a := []byte(s)
	b := []byte(x)
	for _, cur := range res {
		i, j := cur[0], cur[1]
		i--
		j--
		a[i], b[j] = b[j], a[i]
	}

	if string(a) != string(b) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := "souse"
	x := "houhe"
	expect := true
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "cat"
	x := "dog"
	expect := false
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "abc"
	x := "bca"
	expect := true
	runSample(t, s, x, expect)
}
