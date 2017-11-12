package main

import "testing"

func TestNumberOfPatterns(t *testing.T) {
	tests := []struct {
		m, n, want int
	}{
		{1, 1, 9},
		{1, 2, 65},
		{1, 3, 385},
		{1, 9, 389497},
	}

	for _, test := range tests {
		got := numberOfPatterns(test.m, test.n)
		//fmt.Printf("got %d\n", got)
		if got != test.want {
			t.Errorf(`numberOfPatterns(%d, %d) = %d`, test.m, test.n, got)
		}
	}
}
