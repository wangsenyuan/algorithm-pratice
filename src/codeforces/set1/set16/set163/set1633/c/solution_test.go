package main

import "testing"

func runSample(t *testing.T, player []int64, monster []int64, k int64, w int64, a int64, expect bool) {
	res := solve(player, monster, k, w, a)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	player := []int64{25, 4}
	monster := []int64{9, 20}
	var k, w, a int64 = 1, 1, 10
	expect := true
	runSample(t, player, monster, k, w, a, expect)
}

func TestSample2(t *testing.T) {
	player := []int64{25, 4}
	monster := []int64{12, 20}
	var k, w, a int64 = 1, 1, 10
	expect := false
	runSample(t, player, monster, k, w, a, expect)
}

func TestSample3(t *testing.T) {
	player := []int64{100, 1}
	monster := []int64{45, 2}
	var k, w, a int64 = 0, 4, 10
	expect := true
	runSample(t, player, monster, k, w, a, expect)
}

func TestSample4(t *testing.T) {
	player := []int64{9, 2}
	monster := []int64{69, 2}
	var k, w, a int64 = 4, 2, 7
	expect := true
	runSample(t, player, monster, k, w, a, expect)
}
