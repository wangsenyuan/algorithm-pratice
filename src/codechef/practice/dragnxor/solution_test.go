package main

import "testing"

func TestDragnXor(t *testing.T) {
	ps := [][]int{
		[]int{3, 5, 4, 7},
		[]int{5, 0, 1, 16},
		[]int{4, 3, 7, 14},
	}

	for _, p := range ps {
		got := dragnxor(p[0], p[1], p[2])
		if got != p[3] {
			t.Errorf("test %v wrong answer %d\n", p, got)
		}
	}
}
