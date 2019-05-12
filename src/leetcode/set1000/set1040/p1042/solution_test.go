package p1042

import "testing"

func runSample(t *testing.T, N int, paths [][]int) {
	res := gardenNoAdj(N, paths)

	for i := 0; i < N; i++ {
		if res[i] < 1 || res[i] > 4 {
			t.Fatalf("Sample %d %v, answer %v wrong", N, paths, res)
		}
	}

	for _, path := range paths {
		u, v := path[0]-1, path[1]-1
		if res[u] == res[v] {
			t.Fatalf("Sample %d %v, answer %v wrong at %d %d", N, paths, res, u, v)
		}
	}
}

func TestSample1(t *testing.T) {
	N := 3
	paths := [][]int{{1, 2}, {1, 3}, {2, 3}}
	runSample(t, N, paths)
}

func TestSample2(t *testing.T) {
	N := 4
	paths := [][]int{{1, 2}, {3, 4}}
	runSample(t, N, paths)
}

func TestSample3(t *testing.T) {
	N := 4
	paths := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 3}, {2, 4}}
	runSample(t, N, paths)
}
