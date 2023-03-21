package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, traitors []int) {
	sort.Ints(traitors)
	mem := make(map[int]int)

	for _, i := range traitors {
		mem[i]++
	}

	res := solve(n, func(a int, b int, c int) int {
		x := mem[a] + mem[b] + mem[c]
		if x >= 2 {
			return 0
		}
		return 1
	})

	if !reflect.DeepEqual(res, traitors) {
		t.Fatalf("Sample expect %v, but got %v", traitors, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	traitors := []int{3, 4, 1, 2}
	runSample(t, n, traitors)
}

func TestSample2(t *testing.T) {
	n := 9
	traitors := []int{4, 2, 3, 6, 8}
	runSample(t, n, traitors)
}
