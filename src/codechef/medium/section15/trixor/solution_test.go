package main

import (
	"testing"
)

func runSample(t *testing.T, A []int) {
	B := make([]int, len(A))
	copy(B, A)
	res := solve(B)

	cnt := make(map[int]int)

	for _, a := range A {
		cnt[a]++
	}

	for _, cur := range res {
		a, b, c := cur[0], cur[1], cur[2]
		if cnt[a] == 0 || cnt[b] == 0 || cnt[c] == 0 {
			t.Fatalf("cur step %v, not correct", cur)
		}
		cnt[a]--
		cnt[b]--
		cnt[c]--
		a, b, c = a^b, b^c, c^a
		cnt[a]++
		cnt[b]++
		cnt[c]++
	}

	if cnt[0] != len(A) {
		t.Errorf("Sample result not correct")
	}
}

func TestSample1(t *testing.T) {
	A := []int{0, 1, 2, 3, 2, 0}
	runSample(t, A)
}
