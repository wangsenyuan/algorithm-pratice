package main

import "testing"

func runSample(t *testing.T, s string, days []int, medicines [][]string, expect int) {
	res := solve(len(s), s, days, medicines)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "10011"
	days := []int{3, 3, 3, 5}
	medicines := [][]string{
		{"10000", "00110"},
		{"00101", "00000"},
		{"01010", "00100"},
		{"11010", "00100"},
	}
	expect := 8
	runSample(t, s, days, medicines, expect)
}

func TestSample2(t *testing.T) {
	s := "0000"
	days := []int{10}
	medicines := [][]string{
		{"1011", "0100"},
	}
	expect := 0
	runSample(t, s, days, medicines, expect)
}
