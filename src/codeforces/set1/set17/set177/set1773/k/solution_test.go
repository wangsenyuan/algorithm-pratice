package main

import "testing"

func runSample(t *testing.T, n int, k int, expect bool) {
	ok, res := solve(n, k)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t, %v", expect, ok, res)
	}

	if !expect {
		return
	}
	deg := make([]int, n)

	set := NewDSU(n)
	for _, e := range res {
		u, v := e[0]-1, e[1]-1
		deg[u]++
		deg[v]++
		set.Union(u, v)
	}

	ds := make(map[int]int)
	for _, v := range deg {
		ds[v]++
		if v == 0 {
			t.Fatalf("Sample result %v, not correct, it has disconnected node", res)
		}
	}

	if len(ds) != k {
		t.Fatalf("Sample result %v, its deg set is %v, not expected %d-size", res, ds, k)
	}

	for i := 1; i < n; i++ {
		if set.Find(0) != set.Find(1) {
			t.Fatalf("Sample result %v, not correct, is is not connected", res)
		}
	}
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 2, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 1, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 7, 3, true)
}

func TestSample4(t *testing.T) {
	runSample(t, 7, 4, true)
}

func TestSample5(t *testing.T) {
	runSample(t, 3, 2, true)
}
