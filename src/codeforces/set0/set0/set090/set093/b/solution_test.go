package main

import (
	"math"
	"slices"
	"testing"
)

func runSample(t *testing.T, n int, w int, m int, expect bool) {
	res := solve(n, w, m)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	avg := float64(n*w) / float64(m)
	var arr []state
	for _, cur := range res {
		var sum float64
		for _, x := range cur {
			sum += x.vol
		}
		if math.Abs(sum-avg) > eps {
			t.Fatalf("Sample result %v, not correct", res)
		}
		arr = append(arr, cur...)
	}

	slices.SortFunc(arr, func(a, b state) int {
		return a.id - b.id
	})

	for i := 0; i < len(arr); {
		j := i
		var sum float64
		for i < len(arr) && arr[i].id == arr[j].id {
			sum += arr[i].vol
			i++
		}
		if arr[j].id > n {
			t.Fatalf("Sample result %v, not correct it has more bottles", res)
		}
		if math.Abs(sum-float64(w)) > eps {
			t.Fatalf("Sample result not correct, it lose some weight from %d", arr[j].id)
		}
		if i-j > 2 {
			t.Fatalf("Sample result %v not correct, it distributes %d more than 2 people", res, arr[j].id)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 500, 3, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 100, 5, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, 100, 7, false)
}

func TestSample4(t *testing.T) {
	runSample(t, 5, 500, 2, true)
}

func TestSample5(t *testing.T) {
	runSample(t, 4, 100, 8, true)
}

func TestSample6(t *testing.T) {
	runSample(t, 20, 1000, 30, true)
}

func TestSample7(t *testing.T) {
	runSample(t, 50, 1000, 49, true)
}

func TestSample8(t *testing.T) {
	runSample(t, 45, 1000, 50, true)
}

func TestSample9(t *testing.T) {
	runSample(t, 21, 1000, 27, false)
}
