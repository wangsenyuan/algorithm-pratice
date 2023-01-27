package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, k int, items [][]int, expect float64) {
	tot, res := solve(k, items)

	if math.Abs(tot-expect) > 1e-7 {
		t.Fatalf("Sample result %.7f %v, not getting the best result %.7f", tot, res, expect)
	}
	var price float64
	for i := 0; i < k; i++ {
		if len(res[i]) == 0 {
			t.Fatalf("Sample result have no items in %d cart", i)
		}
		var min_stool_price = 1 << 30
		ok := false
		for _, j := range res[i] {
			j--
			ok = ok || items[j][1] == 1
			price += float64(items[j][0])
			if items[j][0] < min_stool_price {
				min_stool_price = items[j][0]
			}
		}
		if ok && min_stool_price < 1<<30 {
			price -= float64(min_stool_price) / 2
		}
	}
	if math.Abs(tot-price) > 1e-7 {
		t.Fatalf("Sample result %.7f %v, not getting the correct result %.7f", tot, res, price)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	items := [][]int{
		{2, 1},
		{3, 2},
		{3, 1},
	}
	expect := 5.5
	runSample(t, k, items, expect)
}

func TestSample2(t *testing.T) {
	k := 3
	items := [][]int{
		{4, 1},
		{1, 2},
		{2, 2},
		{3, 2},
	}
	expect := 8.0
	runSample(t, k, items, expect)
}
func TestSample3(t *testing.T) {
	k := 1
	items := [][]int{
		{1, 1},
	}
	expect := 0.5
	runSample(t, k, items, expect)
}

func TestSample4(t *testing.T) {
	k := 1
	items := [][]int{
		{28, 1},
		{1, 2},
		{1, 2},
		{1, 2},
		{15, 1},
		{16, 1},
		{22, 1},
		{20, 1},
		{1, 2},
		{1, 2},
	}
	expect := 105.5
	runSample(t, k, items, expect)
}
