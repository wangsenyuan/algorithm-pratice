package main

import (
	"math/bits"
	"math/rand"
	"testing"
)

func runSample(t *testing.T, n int) {
	num := n

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("failed to handle %d", num)
		}
	}()

	var cnt int
	ask := func(x int) int {
		cnt++
		if x > n {
			t.Fatalf("%d > %d is not allowed", x, n)
			return -1
		}
		n -= x
		return bits.OnesCount(uint(n))
	}

	res := solve(bits.OnesCount(uint(n)), ask)

	if num != res || cnt > 30 {
		t.Fatalf("Sample result %d not correct, expect %d, or ask too much %d", res, num, cnt)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 3)
}

const inf = 1000000001

func TestSampleN(t *testing.T) {
	rg := rand.New(rand.NewSource(7))

	for i := 0; i < 100; i++ {
		n := rg.Intn(inf)
		for n == 0 {
			n = rg.Intn(inf)
		}
		runSample(t, n)
	}
}
