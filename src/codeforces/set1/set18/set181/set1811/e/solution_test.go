package main

import "testing"

func runSample(t *testing.T, num int64, expect string) {
	res := solve(num)
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var num int64 = 5
	expect := "6"
	runSample(t, num, expect)
}

func TestSample2(t *testing.T) {
	var num int64 = 22
	expect := "25"
	runSample(t, num, expect)
}

func TestSample3(t *testing.T) {
	var num int64 = 100
	expect := "121"
	runSample(t, num, expect)
}

func TestSample4(t *testing.T) {
	var num int64 = 12345
	expect := "18937"
	runSample(t, num, expect)
}

func TestSample5(t *testing.T) {
	var num int64 = 827264634912
	expect := "2932285320890"
	runSample(t, num, expect)
}
