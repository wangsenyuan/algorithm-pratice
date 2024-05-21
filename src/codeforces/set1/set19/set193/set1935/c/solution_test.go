package main

import "testing"

func runSample(t *testing.T, l int, messages [][]int, expect int) {
	res := solve(l, messages)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	l := 8
	messages := [][]int{
		{4, 3},
		{1, 5},
		{2, 4},
		{4, 3},
		{2, 3},
	}
	expect := 3
	runSample(t, l, messages, expect)
}

func TestSample2(t *testing.T) {
	l := 6
	messages := [][]int{
		{4, 10},
	}
	expect := 1
	runSample(t, l, messages, expect)
}

func TestSample3(t *testing.T) {
	l := 12
	messages := [][]int{
		{4, 8},
		{2, 1},
		{2, 12},
	}
	expect := 2
	runSample(t, l, messages, expect)
}

func TestSample4(t *testing.T) {
	l := 26
	messages := [][]int{
		{24, 7},
		{8, 28},
		{30, 22},
		{3, 8},
		{17, 17},
	}
	expect := 1
	runSample(t, l, messages, expect)
}

func TestSample5(t *testing.T) {
	l := 14
	messages := [][]int{
		{15, 3},
		{1000000000, 998244353},
		{179, 239},
		{228, 1337},
		{993, 1007},
	}
	expect := 0
	runSample(t, l, messages, expect)
}
