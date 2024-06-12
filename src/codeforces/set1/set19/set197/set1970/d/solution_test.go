package main

import (
	"math/rand"
	"testing"
)

func runSample(t *testing.T, n int, q int) {
	var sets []string

	interact := func(f func(int) []int) {
		i := rand.Intn(n)
		j := rand.Intn(n)
		power := calc(sets[i] + sets[j])
		res := f(power)

		if len(res) == 0 || res[0] != i+1 || res[1] != j+1 {
			t.Fatalf("Sample (%d, %d): %s, %d result %v, not correct", i, j, sets[i]+sets[j], power, res)
		}
	}

	solve(n, func(s []string) {
		sets = s
	}, interact)
}

func calc(s string) int {
	cache := make(map[string]int)
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			cache[s[i:j+1]]++
		}
	}
	return len(cache)
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 9)
}

func TestSample3(t *testing.T) {
	runSample(t, 10, 100)
}

func TestSample4(t *testing.T) {
	runSample(t, 3, 100)
}
