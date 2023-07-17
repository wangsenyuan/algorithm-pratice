package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, expect int) {
	cnt, res := solve(n)

	if cnt != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, cnt)
	}
	fs := factorize(n)
	arr := make([]int, 0, len(fs))
	for f := range fs {
		arr = append(arr, f)
	}

	for i := 0; i < len(res); i++ {
		a := res[i]
		b := res[(i+1)%len(res)]
		g := gcd(a, b)
		if g == 1 {
			if cnt > 0 {
				cnt--
			} else {
				t.Fatalf("Sample result %v not correct", res)
			}
		}
	}
	sort.Ints(arr)
	sort.Ints(res)

	if !reflect.DeepEqual(res, arr) {
		t.Fatalf("Sample result %v, not correct factors of %d", res, n)
	}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func TestSample1(t *testing.T) {
	runSample(t, 6, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, 30, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, 10, 1)
}

func TestSample5(t *testing.T) {
	runSample(t, 12, 0)
}
