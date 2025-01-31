package main

import (
	"math/rand"
	"testing"
)

func runSample(t *testing.T, n int, q int) {
	rgs, play := solve(n)
	if len(rgs) > 50000 {
		t.Fatalf("Sample give too much rgs %d", len(rgs))
	}

	rand := rand.New(rand.NewSource(47))

	for q > 0 {
		// [0, n)
		l := rand.Intn(n)
		r := l
		if l < n-1 {
			r = l + rand.Intn(n-l)
		}
		if r < l || r >= n {
			t.Fatalf("get random query (%d %d) out of bounds %d", l, r, n)
		}
		l++
		r++
		res := play(l, r)

		if len(res) != 2 {
			t.Fatalf("Sample result %v, not correct", res)
		}
		a, b := res[0], res[1]
		a--
		b--
		if a < 0 || a >= len(rgs) || b < 0 || b >= len(rgs) {
			t.Fatalf("Sample result %v, not correct", res)
		}

		first := min(rgs[a].first, rgs[b].first)
		second := max(rgs[a].second, rgs[b].second)

		if first != l || second != r {
			t.Fatalf("Sample result %v, got range (%d %d), not expected (%d %d)", res, first, second, l, r)
		}

		q--
	}

}

func TestSample1(t *testing.T) {
	runSample(t, 10, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 4000, 10000)
}
