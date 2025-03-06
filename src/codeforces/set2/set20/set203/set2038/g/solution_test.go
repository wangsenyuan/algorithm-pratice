package main

import (
	"math/rand"
	"testing"
)

func runSample(t *testing.T, s string) {
	var times int
	ask := func(x string) int {
		if times == 3 {
			t.Fatalf("Asked too much")
		}
		times++
		var cnt int
		for i := 0; i+len(x) <= len(s); i++ {
			if x == s[i:i+len(x)] {
				cnt++
			}
		}
		return cnt
	}

	p, x := solve(len(s), ask)
	p--
	if s[p] != x {
		t.Fatalf("Sample %s, got wrong aswer [%d] %c", s, p, x)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "101")
}

func TestSample2(t *testing.T) {
	runSample(t, "0")
}
func TestSample(t *testing.T) {
	n := 50
	buf := make([]byte, n)
	xx := "01"
	for i := range n {
		buf[i] = xx[rand.Int()%2]
	}
	for range 10 {
		rand.Shuffle(n, func(i, j int) {
			buf[i], buf[j] = buf[j], buf[i]
		})
		runSample(t, string(buf))
	}
}
