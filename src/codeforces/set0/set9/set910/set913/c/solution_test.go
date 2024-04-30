package main

import "testing"

func runSample(t *testing.T, L int, cost []int, expect int) {
	res := solve(L, cost)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	L := 12
	cost := []int{20, 30, 70, 90}
	expect := 150
	runSample(t, L, cost, expect)
}

func TestSample2(t *testing.T) {
	L := 3
	cost := []int{10000, 1000, 100, 10}
	expect := 10
	runSample(t, L, cost, expect)
}

func TestSample3(t *testing.T) {
	L := 787787787
	cost := []int{123456789, 234567890, 345678901, 456789012, 987654321}
	expect := 44981600785557577
	runSample(t, L, cost, expect)
}
