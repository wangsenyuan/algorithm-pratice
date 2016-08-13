package main

import "testing"

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		ransomNote string
		magazine   string
		want       bool
	}{
		{"fffbfg", "effjfggbffjdgbjjhhdegh", true},
		{"aa", "ab", false},
	}

	for _, test := range tests {
		got := canConstruct(test.ransomNote, test.magazine)
		if got != test.want {
			t.Errorf("%s should be (%v) capable of constructing from %s, but get %v\n", test.ransomNote, test.want, test.magazine, got)
		}
	}
}
