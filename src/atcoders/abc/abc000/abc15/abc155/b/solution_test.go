package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, a []int64, k int64, expect int64) {
	res := solve(k, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int64{3, 3, -4, -2}
	var k int64 = 3
	var expect int64 = -6
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int64{5, 4, 3, 2, -1, 0, 0, 0, 0, 0}
	var k int64 = 40
	var expect int64 = 6
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int64{3, 3, -4, -2}
	var k int64 = 5
	var expect int64 = 8
	runSample(t, a, k, expect)
}

func TestSample4(t *testing.T) {
	a := []int64{3, 3, -4, -2}
	var k int64 = 6
	var expect int64 = 9
	runSample(t, a, k, expect)
}

func TestSample5(t *testing.T) {
	a := []int64{-170202098, -268409015, 537203564, 983211703, 21608710, -443999067, -937727165, -97596546, -372334013, 398994917, -972141167, 798607104, -949068442, -959948616, 37909651, 0, 886627544, -20098238, 0, -948955241, 0, -214720580, 277222296, -18897162, 834475626, 0, -425610555, 110117526, 663621752, 0}
	var k int64 = 413

	var prods []int64

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			prods = append(prods, a[i]*a[j])
		}
	}
	sort.Slice(prods, func(i, j int) bool {
		return prods[i] < prods[j]
	})

	// var expect int64 = 448283280358331064
	expect := prods[k-1]

	runSample(t, a, k, expect)
}

func TestSample6(t *testing.T) {
	a := []int64{-2, -1, 0, 1, 2}
	// -4, -2, -2, -1, 0, 0, 0, 0, 2, 2
	var k int64 = 10
	var expect int64 = 2
	runSample(t, a, k, expect)
}
