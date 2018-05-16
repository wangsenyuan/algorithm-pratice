package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, K int, L int64, expect []int) {
	res := solve(n, K, L)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %d, expect %v, but got %v", n, K, L, expect, res)
	}
}

func TestSample1(t *testing.T) {
	expect := []int{1, 1, 1, 1}
	runSample(t, 4, 3, 1, expect)
}

func TestSample2(t *testing.T) {
	expect := []int{1, 1, 1, 2}
	runSample(t, 4, 3, 2, expect)
}

func TestSample3(t *testing.T) {
	expect := []int{1, 1, 1, 3}
	runSample(t, 4, 3, 3, expect)
}

func TestSample4(t *testing.T) {
	expect := []int{1, 1, 2, 1}
	runSample(t, 4, 3, 4, expect)
}

func TestSample5(t *testing.T) {
	expect := []int{1, 1, 2, 2}
	runSample(t, 4, 3, 5, expect)
}

func TestSample6(t *testing.T) {
	i := int64(1)
	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			for c := 1; c <= 3; c++ {
				for d := 1; d <= 3; d++ {
					runSample(t, 4, 3, i, []int{a, b, c, d})
					i++
				}
			}
		}
	}
}

func TestSample7(t *testing.T) {
	i := int64(1)
	for a := 1; a <= 4; a++ {
		for b := 1; b <= 4; b++ {
			for c := 1; c <= 4; c++ {
				for d := 1; d <= 4; d++ {
					runSample(t, 4, 4, i, []int{a, b, c, d})
					i++
				}
			}
		}
	}
}
