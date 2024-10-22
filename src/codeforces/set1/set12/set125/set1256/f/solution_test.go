package main

import "testing"

func runSample(t *testing.T, s string, x string, expect bool) {
	res := solve(s, x)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abcd"
	x := "abdc"
	expect := false
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "ababa"
	x := "baaba"
	//经过一次len = 3的交换后， ababa, aabba (3)
	// 
	expect := true
	runSample(t, s, x, expect)
}
