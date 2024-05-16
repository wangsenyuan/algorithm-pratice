package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, mat []string, queries [][]int, expect []int) {
	res := solve(n, m, mat, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 9, 9
	mat := []string{
		".........",
		".........",
		"....###..",
		"...v#....",
		"..###....",
		"...##...v",
		"...##....",
		".........",
		"v........",
	}
	queries := [][]int{
		{1, 1},
		{9, 1},
		{5, 7},
	}
	expect := []int{3, 0, 3}
	runSample(t, n, m, mat, queries, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 3
	mat := []string{
		"..v",
		".#.",
		"...",
	}
	queries := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
		{2, 1},
		{3, 2},
	}
	expect := []int{0, 0, 0, 0, 0}
	runSample(t, n, m, mat, queries, expect)
}

func TestSample3(t *testing.T) {
	n, m := 14, 13
	mat := []string{
		".............",
		".............",
		".............",
		"...vvvvvvv...",
		"...v.....v...",
		"...v.###.v...",
		"...v.#.#.v...",
		"...v..v..v...",
		"...v..v..v...",
		"....v...v....",
		".....vvv.....",
		".............",
		".............",
		".............",
	}
	queries := [][]int{
		{1, 1},
		{7, 7},
		{5, 6},
		{4, 10},
		{13, 6},
	}
	expect := []int{3, 0, 1, 0, 2}
	runSample(t, n, m, mat, queries, expect)
}

func TestSample4(t *testing.T) {
	n, m := 10, 11
	mat := []string{
		"...........",
		"..#######..",
		"..#..#..#..",
		"..#.....#..",
		"..#..v..#..",
		"..#.###.#..",
		"..#.#.#.#..",
		"..#...#.#..",
		"..#####.#..",
		"...........",
	}
	queries := [][]int{
		{7, 6},
		{3, 7},
		{6, 8},
		{1, 1},
	}
	expect := []int{1, 2, 3, 4}
	runSample(t, n, m, mat, queries, expect)
}
