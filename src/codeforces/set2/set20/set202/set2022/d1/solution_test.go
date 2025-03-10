package main

import (
	"math/rand"
	"testing"
)

func runSample(t *testing.T, n int, players []int) {
	// 0 for knight, 1 for knave, 2 for imposter
	table := [][]int{
		{1, 0, 1},
		{0, 1, 0},
		{0, 1, 0},
	}
	var cnt int
	ask := func(a int, b int) int {
		cnt++
		if cnt > n+69 {
			t.Fatal("asked too much")
		}
		return table[players[a-1]][players[b-1]]
	}

	res := solve(n, ask)
	if players[res-1] != 2 {
		t.Fatalf("sample %d, %v wrong guess at %d", n, players, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	players := []int{1, 0, 1, 2, 1, 0}
	runSample(t, n, players)
}

func TestSample2(t *testing.T) {
	n := 7
	players := []int{1, 0, 1, 1, 1, 0, 2}
	runSample(t, n, players)
}

func TestSample3(t *testing.T) {
	n := 8
	players := []int{1, 0, 1, 1, 1, 0, 2, 1}
	runSample(t, n, players)
}

func TestSample4(t *testing.T) {
	players := []int{0, 1, 1, 0, 1, 1, 1, 2, 1, 1}
	runSample(t, len(players), players)
}

func TestSample5(t *testing.T) {
	for n := 3; n <= 1000; n += 11 {
		players := make([]int, n)
		for i := range n {
			players[i] = rand.Intn(2)
		}
		x := rand.Intn(n)
		players[x] = 2
		runSample(t, n, players)
	}
}

func TestSample6(t *testing.T) {
	players := []int{1, 0, 1, 1, 0, 1, 0, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0, 1, 1, 1, 1, 0, 0, 2}
	runSample(t, len(players), players)
}

func TestSample7(t *testing.T) {
	players := []int{0, 1, 1, 0, 0, 0, 0, 1, 0, 2, 0, 1, 0}
	runSample(t, len(players), players)
}

func TestSample8(t *testing.T) {
	players := []int{1, 1, 1, 0, 1, 2}
	runSample(t, len(players), players)
}

func TestSample9(t *testing.T) {
	players := []int{1, 1, 0, 1, 2}
	runSample(t, len(players), players)
}

func TestSample10(t *testing.T) {
	players := []int{0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 2, 0, 0}
	runSample(t, len(players), players)
}
