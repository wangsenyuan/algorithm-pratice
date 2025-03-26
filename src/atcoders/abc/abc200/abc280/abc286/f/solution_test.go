package main

import "testing"

func runSample(t *testing.T, n int) {
	ask := func(arr []int) []int {
		m := len(arr)
		res := make([]int, len(arr))
		marked := make([]bool, m)

		for i := 0; i < m; i++ {
			arr[i]--
		}
		for i := 0; i < m; i++ {
			if marked[i] {
				continue
			}
			var cycle []int
			u := i
			for !marked[u] {
				marked[u] = true
				cycle = append(cycle, u)
				u = arr[u]
			}
			l := len(cycle)
			for j, v := range cycle {
				res[v] = cycle[(j+n%l)%l] + 1
			}
		}
		return res
	}

	res := solve(ask)

	if res != n {
		t.Fatalf("Sample expect %d, but got %d", n, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 100)
}

func TestSample3(t *testing.T) {
	runSample(t, 330000)
}