package main

import (
	"math/rand"
	"testing"
)

func TestSample1(t *testing.T) {
	n := 5
	solver := NewSolver(n)

	solver.Update(2, 4, 4)
	solver.Update(3, 5, 5)
	solver.Update(1, 5, 5)

	res := solver.Query(5)

	if res != 1 {
		t.Errorf("Sample fail, expect 1 at pos 5, but got %d", res)
	}

	res = solver.Query(3)

	if res != 2 {
		t.Errorf("Sample fail, expect 2 at pos 3, but got %d", res)
	}
}

func TestSample2(t *testing.T) {
	n := 5
	solver := NewSolver(n)

	solver.Update(5, 5, 4)
	solver.Update(5, 5, 4)
	solver.Update(5, 5, 4)

	res := solver.Query(5)

	if res != 1 {
		t.Errorf("Sample fail, expect 1 at pos 5, but got %d", res)
	}
}

func TestSample3(t *testing.T) {
	n := 5
	solver := NewSolver(n)

	for i := 1; i <= 5; i++ {
		solver.Update(i, i, 1)
		solver.Update(i, i, 2)
	}

	for i := 1; i <= 5; i++ {
		r := solver.Query(i)

		if r != 2 {
			t.Errorf("Sample fail, expect 2 at pos %d, but got %d", i, r)
		}
	}
}

func TestSample4(t *testing.T) {
	n := 100
	sets := make([]map[int]bool, n+1)

	for i := 1; i <= n; i++ {
		sets[i] = make(map[int]bool)
	}

	solver := NewSolver(n)

	for k := 0; k < 200; k++ {
		tp := rand.Intn(2)
		if tp == 0 {
			l := rand.Intn(n) + 1
			r := rand.Intn(n-l) + l
			x := rand.Intn(n) + 1
			//t.Logf("will update range [%d %d] with value %d", l, r, x)
			for i := l; i <= r; i++ {
				sets[i][x] = true
			}
			solver.Update(l, r, x)
		} else {
			p := rand.Intn(n) + 1
			expect := len(sets[p])
			res := solver.Query(p)
			if expect != res {
				t.Errorf("Sample fail when query %d, expect %d, but got %d", p, expect, res)
			}
		}
	}
}
