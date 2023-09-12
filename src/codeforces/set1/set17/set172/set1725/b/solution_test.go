package main

import "testing"

func runSample(t *testing.T, d int, players []int, expect int) {
	res := solve(d, players)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	d := 180
	players := []int{90, 80, 70, 60, 50, 100}
	expect := 2
	runSample(t, d, players, expect)
}

func TestSample2(t *testing.T) {
	d := 10
	players := []int{10, 10, 10}
	expect := 1
	runSample(t, d, players, expect)
}

func TestSample3(t *testing.T) {
	d := 4
	players := []int{1, 2}
	expect := 0
	runSample(t, d, players, expect)
}
