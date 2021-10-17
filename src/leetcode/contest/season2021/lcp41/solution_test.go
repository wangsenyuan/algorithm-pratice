package lcp41

import "testing"

func runSample(t *testing.T, chess []string, expect int) {
	res := flipChess(chess)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	chess := []string{
		".O.....",
		".O.....",
		"XOO.OOX",
		".OO.OO.",
		".XO.OX.",
		"..X.X..",
	}
	expect := 10
	runSample(t, chess, expect)
}
