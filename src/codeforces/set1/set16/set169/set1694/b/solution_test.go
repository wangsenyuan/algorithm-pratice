package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1"
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "01"
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "100"
	expect := 4
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "1001"
	// 10
	// 01
	// 001
	// 1001 => 001   一个0可以把它前面的1都给替换掉，直到遇到一个0
	// 但是一个1，可以把所有前面的0都给你替换掉，直到遇到一个1
	expect := 8
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "11111"
	expect := 5
	runSample(t, s, expect)
}
