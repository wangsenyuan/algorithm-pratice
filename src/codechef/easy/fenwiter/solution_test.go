package main

import "testing"

func runSample(t *testing.T, l1, l2, l3 string, n int, expect int) {
	res := solve([]byte(l1), []byte(l2), []byte(l3), n)
	if res != expect {
		t.Errorf("sample %s %s %s %d, should give %d, but got %d", l1, l2, l3, n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	l1, l2, l3 := "001", "100", "011"
	n := 4
	expect := 6
	runSample(t, l1, l2, l3, n, expect)
}

func TestSample2(t *testing.T) {
	l1, l2, l3 := "1000", "1101", "100"
	n := 3
	expect := 12
	runSample(t, l1, l2, l3, n, expect)
}

func TestSample3(t *testing.T) {
	l1, l2, l3 := "1010", "001", "101"
	n := 4
	expect := 8
	runSample(t, l1, l2, l3, n, expect)
}

func TestSample4(t *testing.T) {
	l1, l2, l3 := "010", "101", "000"
	n := 4
	expect := 10
	runSample(t, l1, l2, l3, n, expect)
}
