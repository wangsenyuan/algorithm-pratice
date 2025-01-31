package main

import "testing"

func runSample(t *testing.T, a int, b int, k int64, expect string) {
	res := solve(a, b, k)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, b := 2, 2
	var k int64 = 4
	expect := "baab"
	runSample(t, a, b, k, expect)
}

func TestSample2(t *testing.T) {
	a, b := 30, 30
	var k int64 = 118264581564861424
	expect := "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	runSample(t, a, b, k, expect)
}

func TestSample3(t *testing.T) {
	a, b := 2, 2
	var k int64 = 3
	expect := "abba"
	runSample(t, a, b, k, expect)
}

func TestSample4(t *testing.T) {
	a, b := 2, 2
	var k int64 = 2
	expect := "abab"
	runSample(t, a, b, k, expect)
}

func TestSample5(t *testing.T) {
	a, b := 2, 2
	var k int64 = 5
	expect := "baba"
	runSample(t, a, b, k, expect)
}

func TestSample6(t *testing.T) {
	a, b := 2, 2
	var k int64 = 1
	expect := "aabb"
	runSample(t, a, b, k, expect)
}
