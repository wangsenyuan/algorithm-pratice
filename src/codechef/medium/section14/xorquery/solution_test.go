package main

import "testing"

func TestSample1(t *testing.T) {
	n := 4
	solver := NewSolver(n)

	res := solver.Update(1, 1, 10)

	if res {
		t.Errorf("Sample fail, should get WA, but got AC")
	}

	res = solver.Update(1, 2, 2)

	if !res {
		t.Errorf("Sample fail, should get AC, but got WA")
	}

	res = solver.Update(2, 3, 4)

	if !res {
		t.Fatalf("sample fail, should get AC, but got WA")
	}

	res = solver.Update(1, 3, 7)
	if res {
		t.Fatalf("Sample fail, should get WA, but got AC")
	}

	ans := solver.Query(1, 3)

	if ans != 6 {
		t.Errorf("Sample fail, should get 6, but got %d", ans)
	}

	ans = solver.Query(3, 4)

	if ans != -1 {
		t.Fatalf("Sample fail, should get -1, but got %d", ans)
	}
}
