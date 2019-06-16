package p794

import "testing"

func runSample(t *testing.T, board []string, expect bool) {
	res := validTicTacToe(board)

	if res != expect {
		t.Errorf("Sample %v, expect %t, but got %t", board, expect, res)
	}
}

func TestSample1(t *testing.T) {
	board := []string{"O  ", "   ", "   "}
	runSample(t, board, false)
}

func TestSample2(t *testing.T) {
	board := []string{"XOX", " X ", "   "}
	runSample(t, board, false)
}

func TestSample3(t *testing.T) {
	board := []string{"XXX", "   ", "OOO"}
	runSample(t, board, false)
}

func TestSample4(t *testing.T) {
	board := []string{"XOX", "O O", "XOX"}
	runSample(t, board, true)
}

func TestSample5(t *testing.T) {
	board := []string{"   ", "   ", "   "}
	runSample(t, board, true)
}

func TestSample6(t *testing.T) {
	board := []string{"XXX", "XOO", "OO "}
	runSample(t, board, false)
}
