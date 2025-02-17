package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, p []int) {
	n := len(p) + 1
	var cnt int
	ask := func(a int, b int) int {
		if a == b {
			t.Fatalf("Sample asked same value [a = b] %d", a)
		}
		cnt++
		if a > b {
			a, b = b, a
		}
		if a == 0 {
			t.Fatal("0 is not allowed")
		}
		// a < b
		for b != a && b != 0 {
			b = p[b-1]
		}
		if b == 0 {
			return 1
		}
		return 0
	}

	res := solve(n, ask)

	if cnt > 2*n-6 {
		t.Fatalf("Sample asked too much %d", cnt)
	}

	if !reflect.DeepEqual(res, p) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{0, 0, 1}
	runSample(t, p)
}

func TestSample2(t *testing.T) {
	p := []int{0, 0, 1, 2}
	runSample(t, p)
}

func TestSample3(t *testing.T) {
	p := []int{0, 0, 0, 1, 3, 5, 6, 7}
	runSample(t, p)
}

func TestSample4(t *testing.T) {
	p := []int{0, 0, 0, 1}
	runSample(t, p)
}

func TestSample5(t *testing.T) {
	p := []int{0, 1, 2}
	runSample(t, p)
}
