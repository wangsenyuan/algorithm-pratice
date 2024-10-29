package main

import "testing"

func runSample(t *testing.T, a int, ta int, b int, tb int, simion string, expect int) {
	res := solve(a, ta, b, tb, simion)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, ta := 10, 30
	b, tb := 10, 35
	simion := "05:20"
	expect := 5
	runSample(t, a, ta, b, tb, simion, expect)
}

func TestSample2(t *testing.T) {
	a, ta := 60, 120
	b, tb := 24, 100
	simion := "13:00"
	expect := 9
	runSample(t, a, ta, b, tb, simion, expect)
}

func TestSample3(t *testing.T) {
	a, ta := 30, 60
	b, tb := 10, 60
	simion := "23:30"
	expect := 8
	runSample(t, a, ta, b, tb, simion, expect)
}

func TestSample4(t *testing.T) {
	a, ta := 68, 34
	b, tb := 84, 78
	simion := "10:40"
	expect := 1
	runSample(t, a, ta, b, tb, simion, expect)
}
