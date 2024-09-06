package main

import "testing"

func runSample(t *testing.T, weapons [][]int, armorSets [][]int, monsters [][]int, expect int) {
	res := solve(weapons, armorSets, monsters)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	weapons := [][]int{
		{2, 3},
		{4, 7},
	}
	armorSets := [][]int{
		{2, 4},
		{3, 2},
		{5, 11},
	}
	monsters := [][]int{
		{1, 2, 4},
		{2, 1, 6},
		{3, 4, 6},
	}
	expect := 1
	runSample(t, weapons, armorSets, monsters, expect)
}
