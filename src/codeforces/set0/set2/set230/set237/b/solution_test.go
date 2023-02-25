package main

import "testing"

func runSample(t *testing.T, A [][]int) {
	var a [][]int
	for _, row := range A {
		wo := make([]int, len(row))
		copy(wo, row)
		a = append(a, wo)
	}

	res := solve(A)

	for _, op := range res {
		x, y, u, v := op[0], op[1], op[2], op[3]
		a[x-1][y-1], a[u-1][v-1] = a[u-1][v-1], a[x-1][y-1]
	}

	for i := 1; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] <= a[i-1][j] {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}

	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a[i]); j++ {
			if a[i][j] <= a[i][j-1] {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	A := [][]int{
		{4, 3, 5},
		{6, 1},
		{2},
	}
	runSample(t, A)
}
