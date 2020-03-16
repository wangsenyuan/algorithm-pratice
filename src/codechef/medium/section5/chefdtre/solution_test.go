package main

import "testing"

func TestSample1(t *testing.T) {
	solver := NewSolver(7)

	solver.Union(1, 2)
	solver.Union(8, 3)

	res := solver.Get(4, 1)

	if res != 4 {
		t.Fatalf("Sample expect 4 at step 1, but got %d", res)
	}

	solver.Union(4, 5)

	res = solver.Get(10, 2)

	if res != 5 {
		t.Fatalf("Sample expect 5 at step 2, but got %d", res)
	}

	res = solver.Get(9, 3)
	if res != 3 {
		t.Fatalf("Sample expect 3 at step 3, but got %d", res)
	}

	res = solver.Get(9, 2)
	if res != 2 {
		t.Fatalf("Sample expect 2 at step 4, but got %d", res)
	}

	res = solver.Get(9, 1)
	if res != 1 {
		t.Fatalf("Sample expect 1 at step 5, but got %d", res)
	}
}

func TestSample2(t *testing.T) {
	solver := NewSolver2(7)

	solver.Union(1, 2)
	solver.Union(8, 3)

	res := solver.Get(4, 1)

	if res != 4 {
		t.Fatalf("Sample expect 4 at step 1, but got %d", res)
	}

	solver.Union(4, 5)

	res = solver.Get(10, 2)

	if res != 5 {
		t.Fatalf("Sample expect 5 at step 2, but got %d", res)
	}

	res = solver.Get(9, 3)
	if res != 3 {
		t.Fatalf("Sample expect 3 at step 3, but got %d", res)
	}

	res = solver.Get(9, 2)
	if res != 2 {
		t.Fatalf("Sample expect 2 at step 4, but got %d", res)
	}

	res = solver.Get(9, 1)
	if res != 1 {
		t.Fatalf("Sample expect 1 at step 5, but got %d", res)
	}
}
