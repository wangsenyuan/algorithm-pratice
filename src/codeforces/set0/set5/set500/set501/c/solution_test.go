package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, deg []int, s []int) {
	n := len(deg)

	u := make([]int, n)
	copy(u, deg)
	v := make([]int, n)
	copy(v, s)

	res := solve(u, v)

	d := make([]int, n)
	x := make([]int, n)

	for _, e := range res {
		u, v := e[0], e[1]
		d[u]++
		d[v]++
		x[u] ^= v
		x[v] ^= u
	}

	if !reflect.DeepEqual(d, deg) || !reflect.DeepEqual(x, s) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	deg := []int{2, 1, 1}
	s := []int{3, 0, 0}
	runSample(t, deg, s)
}

func TestSample2(t *testing.T) {
	deg := []int{1, 1}
	s := []int{1, 0}
	runSample(t, deg, s)
}

func TestSample3(t *testing.T) {
	data := `11
0 0
0 0
0 0
0 0
0 0
0 0
0 0
1 8
1 7
0 0
0 0
	`
	reader := bufio.NewReader(strings.NewReader(data))

	d, s := process(reader)
	runSample(t, d, s)
}
