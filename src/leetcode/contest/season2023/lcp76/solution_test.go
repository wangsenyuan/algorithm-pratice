package lcp76

import "testing"

func runSample(t *testing.T, n int, m int, board []string, expect int64) {
	res := getSchemeCount(n, m, board)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample0(t *testing.T) {
	n, m := 1, 3
	board := []string{
		"..R",
	}
	var expect int64 = 1
	runSample(t, n, m, board, expect)
}

func TestSample1(t *testing.T) {
	n, m := 1, 3
	board := []string{
		"..?",
	}
	var expect int64 = 3
	runSample(t, n, m, board, expect)
}

func TestSample2(t *testing.T) {
	n, m := 1, 3
	board := []string{
		".B?",
	}
	var expect int64 = 3
	runSample(t, n, m, board, expect)
}

func TestSample3(t *testing.T) {
	n, m := 1, 3
	board := []string{
		"...",
	}
	var expect int64 = 1
	runSample(t, n, m, board, expect)
}

func TestSample4(t *testing.T) {
	n, m := 3, 3
	board := []string{
		"..R", "..B", "?R?",
	}
	var expect int64 = 5
	runSample(t, n, m, board, expect)
}

func TestSample5(t *testing.T) {
	n, m := 3, 3
	board := []string{
		"?R?", "B?B", "?R?",
	}
	var expect int64 = 105
	runSample(t, n, m, board, expect)
}

func TestSample6(t *testing.T) {
	n, m := 1, 30
	board := []string{
		"..BB..?????.B.?R.?R?..??R??.R.",
	}
	var expect int64 = 0
	runSample(t, n, m, board, expect)
}

func TestSample7(t *testing.T) {
	n, m := 2, 15
	board := []string{
		"????R??.R?.R?.?", "RR..???...??B..",
	}
	var expect int64 = 0
	runSample(t, n, m, board, expect)
}

func TestSample8(t *testing.T) {
	n, m := 2, 15
	board := []string{
		"???????????????", "???????????????",
	}
	//  4442889025
	// 17171219521
	var expect int64 = 17171219521
	runSample(t, n, m, board, expect)
}

func TestSample9(t *testing.T) {
	n, m := 2, 2
	board := []string{
		"??", "??",
	}
	var expect int64 = 3 * 3 * 3 * 3
	runSample(t, n, m, board, expect)
}

func TestSample10(t *testing.T) {
	n, m := 3, 3
	board := []string{
		"???", "???", "???",
	}
	var expect int64 = 8451
	runSample(t, n, m, board, expect)
}

func TestSample11(t *testing.T) {
	n, m := 1, 4
	board := []string{
		"????",
	}

	var expect int64 = 53
	runSample(t, n, m, board, expect)
}

func TestSample12(t *testing.T) {
	n, m := 4, 2
	board := []string{
		"?.",
		"?R",
		"?.",
		".?",
	}

	var expect int64 = 69
	runSample(t, n, m, board, expect)
}
