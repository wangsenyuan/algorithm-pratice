package main

import (
	"testing"
)

func validSolution(got, want []rune) bool {
	if len(got) != len(want) {
		return false
	}

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}

func TestCanPlayCases(t *testing.T) {
	tests := []struct {
		n, p, r, s int
		want       []rune
	}{
		{2, 1, 1, 0, []rune{'P', 'R'}},
		{4, 1, 1, 2, []rune{'P', 'S', 'R', 'S'}},
	}

	solution := NewSolution()

	for _, test := range tests {
		got, found := solution.play(test.n, test.p, test.r, test.s)
		if !found {
			t.Errorf(`should have solution for %d, %d, %d, %d`, test.n, test.p, test.r, test.s)
		} else if !validSolution(got, test.want) {
			t.Errorf(`got %v, but want %v for %d, %d, %d, %d`, got, test.want, test.n, test.p, test.r, test.s)
		}
	}
}


