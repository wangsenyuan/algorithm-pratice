package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, fake []int) {
	mem := make(map[int]bool)
	for _, x := range fake {
		mem[x] = true
	}
	var cnt int
	ask := func(x int, y int) int {
		cnt++
		// 0 for same 1 for different
		if mem[x] == mem[y] {
			return 0
		}
		return 1
	}

	res := solve(ask)

	if cnt > 950 {
		t.Fatalf("solution took too much times %d", cnt)
	}
	sort.Ints(fake)

	if !reflect.DeepEqual(res, fake) {
		t.Fatalf("Sample result %v, not correct, expect %v", res, fake)
	}
}

func TestSample1(t *testing.T) {
	fake := make([]int, 10)
	for i := 0; i < 10; i++ {
		fake[i] = 982 + i
	}
	runSample(t, fake)
}

func TestSample2(t *testing.T) {

	arr := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i + 1
	}

	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	fake := arr[:10]

	runSample(t, fake)
}

func TestSample3(t *testing.T) {
	fake := make([]int, 10)
	for i := 0; i < 10; i++ {
		fake[i] = 43 + i
	}
	runSample(t, fake)
}

func TestSample4(t *testing.T) {
	fake := make([]int, 10)
	for i := 0; i < 10; i++ {
		fake[i] = 1 + i
	}
	runSample(t, fake)
}
