package main

import (
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
