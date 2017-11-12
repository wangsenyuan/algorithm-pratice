package main

import "testing"

func TestIsNumber(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{" 0.1", true},
		{" 0.1  ", true},
		{" 0.01a", false},
		{"1 a", false},
		{"2e10", true},
		{"-2e10", true},
		{"4.", true},
		{".", false},
		{"46.e3", true},
		{".2e81", true},
		{"92e1740e91", false},
	}

	for _, test := range tests {
		got := isNumber(test.s)
		if got != test.want {
			t.Errorf("isNumber(%s) = %v; but it should be %v\n", test.s, got, test.want)
		}
	}
}
