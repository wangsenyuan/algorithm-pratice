package main

import (
	"fmt"
	"testing"
)

func play1(t *testing.T, n int, p []int, pairs [][]int) {
	vis := make([]bool, 2*n+1)
	pick := func(x int) int {
		if vis[x] {
			panic("can't use a used one")
		}
		vis[x] = true

		for _, cur := range pairs {
			if cur[0] == x || cur[1] == x {
				y := x ^ cur[0] ^ cur[1]
				vis[y] = true
				return y
			}
		}

		for i := 1; i <= 2*n; i++ {
			if !vis[i] {
				vis[i] = true
				return i
			}
		}
		panic("asked too much")
	}

	solve1(n, p, pairs, pick)
}

func play2(t *testing.T, n int, p []int, pairs [][]int, first int) {
	vis := make([]bool, 2*n+1)
	vis[first] = true

	pick := func(x int) int {
		if vis[x] {
			panic(fmt.Sprintf("can't use a used one %d", x))
		}
		vis[x] = true

		for _, cur := range pairs {
			if cur[0] == x || cur[1] == x {
				y := x ^ cur[0] ^ cur[1]
				if vis[y] {
					// 都已经被使用了
					break
				}
				vis[y] = true
				return y
			}
		}

		for i := 1; i <= 2*n; i++ {
			if !vis[i] {
				vis[i] = true
				return i
			}
		}
		return 0
	}

	solve2(n, p, pairs, first, pick)
}

func TestSample1(t *testing.T) {
	n := 3
	p := []int{1, 2, 3, 4, 5, 6}
	pairs := [][]int{{2, 6}}
	play1(t, n, p, pairs)
}

func TestSample2(t *testing.T) {
	n := 3
	p := []int{1, 2, 3, 4, 5, 6}
	pairs := [][]int{{1, 5}}
	play2(t, n, p, pairs, 6)
}

func TestSample3(t *testing.T) {
	n := 3
	p := []int{1, 2, 3, 4, 5, 6}
	pairs := [][]int{{1, 5}}
	play2(t, n, p, pairs, 1)
}

func TestSample4(t *testing.T) {
	n := 20
	p := []int{594, 114, 289, 895, 991, 986, 870, 876, 403, 963, 593, 542, 703, 753, 37, 546, 787, 267, 835, 809, 257, 501, 824, 10, 216, 790, 901, 493, 706, 739, 462, 409, 150, 906, 68, 630, 79, 32, 955, 764}
	pairs := [][]int{
		{8, 29},
		{4, 12},
		{33, 9},
		{19, 5},
		{36, 37},
		{16, 1},
		{23, 3},
		{7, 15},
		{35, 39},
		{40, 11},
	}
	play2(t, n, p, pairs, 12)
}
