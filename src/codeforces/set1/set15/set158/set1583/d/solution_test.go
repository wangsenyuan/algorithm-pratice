package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func runSample(t *testing.T, p []int) {
	var cnt int

	n := len(p)

	s := make([]int, n)

	ask := func(a []int) int {
		cnt++
		if cnt > 2*n {
			panic("asked too much")
		}
		copy(s, a)
		vis := make(map[int]int)
		ans := -1
		for i := n - 1; i >= 0; i-- {
			s[i] += p[i]
			vis[s[i]]++
			if vis[s[i]] > 1 {
				ans = i
			}
		}
		return ans + 1
	}

	res := solve(n, ask)

	if !reflect.DeepEqual(res, p) {
		t.Fatalf("Sample expect %v, but got %v", p, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{3, 2, 1, 5, 4}
	runSample(t, p)
}

func TestSample2(t *testing.T) {
	p := []int{1, 2, 3, 4, 5, 6}
	runSample(t, p)
}

func TestSample3(t *testing.T) {
	p := []int{6, 5, 4, 3, 2, 1}
	runSample(t, p)
}

func TestSample4(t *testing.T) {
	n := 100
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i + 1
	}
	for j := 0; j < 10; j++ {
		rand.Shuffle(n, func(i, j int) {
			p[i], p[j] = p[j], p[i]
		})
		runSample(t, p)
	}
}
