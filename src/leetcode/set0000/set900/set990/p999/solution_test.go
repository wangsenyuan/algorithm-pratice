package p999

import "testing"

func runSample(t *testing.T, board []string, expect int) {
	tmp := make([][]byte, len(board))
	for i := 0; i < len(board); i++ {
		tmp[i] = []byte(board[i])
	}

	res := numRookCaptures(tmp)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", board, expect, res)
	}
}

func TestSample1(t *testing.T) {
	board := []string{
		"........",
		"...p....",
		"...R...p",
		"........",
		"........",
		"...p....",
	}
	runSample(t, board, 3)
}

func TestSample2(t *testing.T) {
	board := []string{
		"pppppppp",
		"pppBpppp",
		"ppBRBppp",
		"pppBpppp",
		"........",
		"...p....",
	}
	runSample(t, board, 0)
}
