package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, L int, A []int, expect float64) {
	res := solve(L, A)

	if math.Abs(res-expect) > 1e-7 {
		t.Fatalf("Sample expect %.8f, but got %.8f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	L := 10
	A := []int{1, 9}
	expect := 3.0
	runSample(t, L, A, expect)
}

func TestSample2(t *testing.T) {
	L := 10
	A := []int{1}
	expect := 3.666666666666667
	runSample(t, L, A, expect)
}

func TestSample3(t *testing.T) {
	L := 7
	A := []int{1, 2, 3, 4, 6}
	expect := 2.047619047619048
	runSample(t, L, A, expect)
}

func TestSample4(t *testing.T) {
	L := 1000000000
	A := []int{413470354, 982876160}
	expect := 329737645.750000000000000
	runSample(t, L, A, expect)
}

func TestSample5(t *testing.T) {
	L := 478
	A := []int{1, 10, 25, 33, 239, 445, 453, 468, 477}
	expect := 53.700000000000000
	runSample(t, L, A, expect)
}
