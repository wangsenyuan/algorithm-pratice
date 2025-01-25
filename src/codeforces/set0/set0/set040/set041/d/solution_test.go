package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	board, sum, col, moves := process(reader)

	if expect != sum {
		t.Fatalf("Sample expect %d, but got %d", expect, sum)
	}
	if expect < 0 {
		return
	}
	n := len(board)
	// m := len(board[0])
	col--
	sum = 0
	for i := 0; i < n-1; i++ {
		r := n - 1 - i
		sum += int(board[r][col] - '0')
		if moves[i] == 'L' {
			col--
		} else {
			col++
		}
	}

	sum += int(board[0][col] - '0')

	if sum != expect {
		t.Fatalf("Sample got the sum following the path %s", moves)
	}
}

func TestSample1(t *testing.T) {
	s := `3 3 1
123
456
789
`
	expect := 16
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3 0
123
456
789
`
	expect := 17
	runSample(t, s, expect)
}
func TestSample3(t *testing.T) {
	s := `2 2 10
98
75
`
	expect := -1
	runSample(t, s, expect)
}
