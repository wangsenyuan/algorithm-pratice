package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, coord [][]int, s []string, K int, expect float64) {
	res := solve(n, coord, s, K)
	if math.Abs(res-expect) > 0.000001 {
		t.Errorf("sample: %d %v %v %d, should give answer %f, but got %f", n, coord, s, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, K := 3, 6
	coord := [][]int{
		{100, 100},
		{-50, 100},
		{0, -100},
	}
	s := []string{
		"111110",
		"110011",
		"001111",
	}
	expect := float64(403.2247551123)
	runSample(t, n, coord, s, K, expect)
}

func TestSample2(t *testing.T) {
	n, K := 3, 6
	coord := [][]int{
		{100, 100},
		{-50, 100},
		{0, -100},
	}
	s := []string{
		"111110",
		"110011",
		"111111",
	}
	expect := float64(200.0000000000)
	runSample(t, n, coord, s, K, expect)
}

func TestSample3(t *testing.T) {
	n, K := 3, 6
	coord := [][]int{
		{100, 100},
		{-50, 100},
		{0, -210},
	}
	s := []string{
		"111110",
		"110011",
		"111111",
	}
	expect := float64(403.2247551123)
	runSample(t, n, coord, s, K, expect)
}
func TestSample4(t *testing.T) {
	n, K := 3, 6
	coord := [][]int{
		{100, 100},
		{-50, 100},
		{0, -210},
	}
	s := []string{
		"111010",
		"110011",
		"111011",
	}
	expect := float64(-1.0)
	runSample(t, n, coord, s, K, expect)
}

func TestSample5(t *testing.T) {
	n, K := 10, 10
	coord := [][]int{
		{763, 77},
		{512, 859},
		{494, 362},
		{-986, -885},
		{-866, 890},
		{128, 553},
		{-879, 981},
		{449, 310},
		{-406, 10},
		{-955, -475},
	}
	s := []string{
		"0000001000",
		"1000000000",
		"0000000100",
		"0001000000",
		"0000100000",
		"0010000000",
		"0000010000",
		"0000000001",
		"0100000000",
		"0000000010",
	}
	expect := float64(6652.7837723893)
	runSample(t, n, coord, s, K, expect)
}
