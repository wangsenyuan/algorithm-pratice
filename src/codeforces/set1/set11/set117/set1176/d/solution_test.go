package main

import (
	"reflect"
	"slices"
	"sort"
	"testing"
)

func runSample(t *testing.T, b []int) {
	a := solve(b)

	if len(a) != len(b)/2 {
		t.Fatalf("Sample result %v, not correct", a)
	}

	c := create(a, len(b), slices.Max(b))

	sort.Ints(b)
	sort.Ints(c)

	if !reflect.DeepEqual(b, c) {
		t.Fatalf("Sample result %v, not correct", a)
	}
}

func create(a []int, n int, x int) []int {
	lpf, primes := getPrimes(x)

	c := make([]int, len(a), n)
	copy(c, a)

	for _, x := range a {
		if lpf[x] == x {
			c = append(c, primes[x-1])
		} else {
			y := x / lpf[x]
			c = append(c, y)
		}
	}

	return c
}

func TestSample1(t *testing.T) {
	b := []int{3, 5, 2, 3, 2, 4}
	runSample(t, b)
}

func TestSample2(t *testing.T) {
	b := []int{2750131, 199999}
	runSample(t, b)
}

func TestSample3(t *testing.T) {
	b := []int{3, 6}
	runSample(t, b)
}

func TestSample4(t *testing.T) {
	a := []int{6359, 6047, 4157, 3037, 2677, 2659, 1931, 1667, 12273, 1433, 9040, 1117, 997, 937, 853, 5862, 5626, 5597, 5115, 4982, 4975, 4761, 4718, 4494, 4137, 4115, 3525, 3466, 3369, 3332, 3245, 2934, 2920, 2920, 419, 2872, 2721, 389, 2669, 2611, 2607, 2548, 2536, 2244, 2206, 2019, 1990, 1896, 1829, 1817, 1810, 1760, 1752, 1684, 1659, 1551, 1538, 1524, 1469, 1400, 1390, 1315, 1222, 1222, 1178, 1114, 1099, 1068, 1041, 998, 979, 973, 939, 939, 916, 879, 800, 784, 748, 746, 722, 718, 668, 621, 608, 591, 551, 484, 465, 412, 391, 325, 305, 238, 31, 111, 23, 23, 52, 49}
	b := create(a, len(a)*2, 2750131)
	runSample(t, b)
}
