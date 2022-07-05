package main

import "testing"

func TestSample1(t *testing.T) {
	solver := NewSolver(9)
	solver.Add(1)
	solver.Add(7)

	res := solver.Report()

	if res != -1 {
		t.Fatalf("Sample expect -1, but got %d", res)
	}

	solver.Add(9)
	solver.Add(21)
	solver.Add(8)
	solver.Add(5)

	res = solver.Report()

	if res != 9 {
		t.Fatalf("Sample expect 9, but got %d", res)
	}

	solver.Add(9)

	res = solver.Report()

	if res != 9 {
		t.Fatalf("Sample expect 9, but got %d", res)
	}

}

func TestSample2(t *testing.T) {
	arr := []int{2, 8, 3, 1, 6, 4, 5, 7}

	solver := NewSolver(len(arr))
	// 2 8 3 1 6 4 5 7
	for _, num := range arr {
		solver.Add(num)
	}

	res := solver.Report()

	if res != 7 {
		t.Fatalf("Sample expect 7, but got %d", res)
	}

}

func TestSample3(t *testing.T) {
	arr := []int{2, 8, 3, 1, 6, 4, 5, 7, 9}

	solver := NewSolver(len(arr))
	// 2 8 3 1 6 4 5 7
	for _, num := range arr {
		solver.Add(num)
	}

	res := solver.Report()

	if res != 7 {
		t.Fatalf("Sample expect 7, but got %d", res)
	}

}
