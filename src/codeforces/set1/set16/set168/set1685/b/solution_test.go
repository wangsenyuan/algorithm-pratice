package main

import "testing"

func runSample(t *testing.T, s string, num []int, expect bool) {
	res := solve(s, num)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num := []int{0, 0, 1, 0}
	s := "AB"
	expect := true
	runSample(t, s, num, expect)
}

func TestSample2(t *testing.T) {
	num := []int{1, 1, 0, 1}
	s := "ABAB"
	expect := true
	runSample(t, s, num, expect)
}

func TestSample3(t *testing.T) {
	num := []int{1, 1, 2, 2}
	s := "BAABBABBAA"
	expect := true
	runSample(t, s, num, expect)
}

func TestSample4(t *testing.T) {
	num := []int{1, 1, 2, 3}
	s := "ABABABBAABAB"
	expect := true
	runSample(t, s, num, expect)
}

func TestSample5(t *testing.T) {
	num := []int{1, 3, 3, 10}
	s := "BBABABABABBBABABABABABABAABABA"
	expect := true
	runSample(t, s, num, expect)
}

func TestSample6(t *testing.T) {
	num := []int{2, 3, 5, 4}
	s := "AABAABBABAAABABBABBBABB"
	expect := false
	runSample(t, s, num, expect)
}
