package main

import "testing"

func runSample(t *testing.T, n, k int, board []string, expect bool) {
	b := make([][]byte, n)
	for i := 0; i < n; i++ {
		b[i] = []byte(board[i])
	}

	res := solve(n, k, b)
	if res != expect {
		t.Errorf("sample %d %d %v, expect %t, but got %t", n, k, board, expect, res)
	}
}

func TestSample1(t *testing.T) {
	board := []string{
		"XOX",
		"O.O",
		"XOX",
	}
	runSample(t, 3, 3, board, true)
}

func TestSample2(t *testing.T) {
	board := []string{
		"...",
		"...",
		"...",
	}
	runSample(t, 3, 1, board, true)
}

func TestSample3(t *testing.T) {
	board := []string{
		"...",
		"...",
		"...",
	}
	runSample(t, 3, 2, board, false)
}

func TestSample4(t *testing.T) {
	board := []string{
		"XOXO",
		"OX..",
		"XO..",
		"OXOX",
	}
	runSample(t, 4, 4, board, true)
}

func TestSample5(t *testing.T) {
	board := []string{
		"X.X",
		"...",
		".O.",
	}
	runSample(t, 3, 3, board, true)
}

func TestSample6(t *testing.T) {
	board := []string{
		"X..",
		"...",
		"XO.",
	}
	runSample(t, 3, 3, board, true)
}

func TestSample7(t *testing.T) {
	board := []string{
		"X..",
		"...",
		".OX",
	}
	runSample(t, 3, 3, board, true)
}

func TestSample8(t *testing.T) {
	board := []string{
		"..X",
		"...",
		"XO.",
	}
	runSample(t, 3, 3, board, true)
}

func TestSample9(t *testing.T) {
	board := []string{
		"...",
		".X.",
		"XO.",
	}
	runSample(t, 3, 3, board, true)
}

func TestSample10(t *testing.T) {
	board := []string{
		"...",
		".X.",
		".O.",
	}
	runSample(t, 3, 3, board, false)
}

func TestSample11(t *testing.T) {
	board := []string{
		".O..",
		".XOX",
		".X..",
		".O..",
	}
	runSample(t, 4, 3, board, false)
}

func TestSample12(t *testing.T) {
	board := []string{
		".X..",
		".O.X",
		"....",
		"..X.",
	}
	runSample(t, 4, 3, board, false)
}
