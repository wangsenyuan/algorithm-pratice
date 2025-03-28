package main

import (
	"math/rand"
	"testing"
	"time"
)

func runSample(t *testing.T, x int, y int, m int) {
	ask := func(n int) int {
		return (n | x) + (n | y)
	}

	res := solve(ask)

	u := res(m)
	v := ask(m)

	if u != v {
		t.Fatalf("Sample solution not correct, got %d, want %d", u, v)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 0, 0, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 7, 0, 1232)
}

func TestSample3(t *testing.T) {
	runSample(t, 15, 0, 16)
}

func TestSample4(t *testing.T) {
	runSample(t, 15, 19, 16)
}

func TestSample5(t *testing.T) {
	const limit = 1 << 30
	const times = 1000

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for range times {
		x := rnd.Intn(limit)
		y := rnd.Intn(limit)
		m := rnd.Intn(limit)

		// t.Logf("Testing with x = %d, y = %d, m = %d", x, y, m)
		runSample(t, x, y, m)
	}
}

func TestSample6(t *testing.T) {
	runSample(t, 91423622, 644999207, 215136919)
}