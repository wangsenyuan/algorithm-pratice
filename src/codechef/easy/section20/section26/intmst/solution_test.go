package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestSample1(t *testing.T) {
	n, m := 3, 3
	edges := [][]int{
		{0, 1},
		{1, 2},
		{0, 2},
	}
	P := []int{2, 0, 1}
	expect := true
	runSample(t, n, m, edges, P, expect)
}

func runSample(t *testing.T, n int, m int, edges [][]int, P []int, expect bool) {
	query := func(n int, w []int) []int {
		es := mst(n, w, edges)
		res := make([]int, n-1)
		for i, e := range es {
			res[i] = P[e]
		}
		return res
	}

	res := solve(n, m, query)

	if expect != (res != nil) {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
		return
	}

	if !expect {
		return
	}

	if !reflect.DeepEqual(res, P) {
		t.Errorf("Sample result %v, not correct, should be %v", res, P)
	}
}

func mst(n int, weight []int, edges [][]int) []int {
	E := make([]Edge, len(weight))

	for i, w := range weight {
		E[i] = Edge{i, w}
	}

	sort.Slice(E, func(i, j int) bool {
		return E[i].weight < E[j].weight
	})

	uf := NewUF(n)

	res := make([]int, 0, n-1)
	for _, e := range E {
		i := e.index
		u, v := edges[i][0], edges[i][1]
		if uf.Union(u, v) {
			res = append(res, i)
		}
	}
	return res
}

type Edge struct {
	index  int
	weight int
}

type UF struct {
	arr []int
	cnt []int
}

func NewUF(n int) UF {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i]++
	}
	return UF{arr, cnt}
}

func (uf *UF) Find(x int) int {
	if uf.arr[x] != x {
		uf.arr[x] = uf.Find(uf.arr[x])
	}
	return uf.arr[x]
}

func (uf *UF) Union(a, b int) bool {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa == pb {
		return false
	}
	if uf.cnt[pa] < uf.cnt[pb] {
		pa, pb = pb, pa
	}
	uf.cnt[pa] += uf.cnt[pb]
	uf.arr[pb] = pa
	return true
}
