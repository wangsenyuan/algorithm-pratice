package main

import (
	"math/rand"
	"testing"
)

func runSample(t *testing.T, n int) {
	x, y := solve(n)

	if x+y != n {
		t.Fatalf("Sample result %d + %d != %d", x, y, n)
	}

	if abs(sumDigits(x)-sumDigits(y)) > 1 {
		t.Fatalf("Sample result %d %d, not correct", x, y)
	}
}

func sumDigits(num int) int {
	var res int
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func TestSample1(t *testing.T) {
	n := 1
	runSample(t, n)
}

func TestSample2(t *testing.T) {
	n := 161
	runSample(t, n)
}

func TestSample3(t *testing.T) {
	n := 67
	runSample(t, n)
}

func TestSample4(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := rand.Int()
		if n%2 == 0 {
			n++
		}
		runSample(t, n)
	}
}
