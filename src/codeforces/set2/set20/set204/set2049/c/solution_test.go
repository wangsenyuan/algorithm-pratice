package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))

	n, x, y, a := process(reader)
	x--
	y--
	vis := make([]bool, 4)
	get := func(i int) int {
		clear(vis)
		vis[a[(i-1+n)%n]] = true
		vis[a[(i+1)%n]] = true
		if x == i {
			vis[a[y]] = true
		}
		if y == i {
			vis[a[x]] = true
		}
		var x int
		for vis[x] {
			x++
		}
		return x
	}

	for i := range n {
		x := get(i)
		if x != a[i] {
			t.Fatalf("Sample result %v, not correct at %d / %d", a, i+1, x)
		}
	}
}

func TestSample1(t *testing.T) {
	s := "5 1 3"
	runSample(t, s)
}
