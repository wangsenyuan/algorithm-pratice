package main

import "testing"

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{")()())", 4},
		{")())", 2},
		{")()", 2},
		{"(()", 2},
		{"()()", 4},
		{"((())()(()()()()", 8},
	}

	for _, test := range tests {
		got := longestValidParentheses(test.s)
		if got != test.want {
			t.Errorf("%d != %d, wrooooog with %v!", got, test.want, test.s)
		}
	}

}
