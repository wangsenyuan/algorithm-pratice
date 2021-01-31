package p1743

import "testing"

func runSample(t *testing.T, pairs [][]int) {
	res := restoreArray(pairs)

	seen := make(map[[2]int]bool)

	for _, p := range pairs {
		x := [...]int{p[0], p[1]}
		seen[x] = true
	}

	for i := 1; i < len(res); i++ {
		x := [...]int{res[i-1], res[i]}
		if seen[x] {
			continue
		}
		x = [...]int{res[i], res[i-1]}
		if seen[x] {
			continue
		}
		t.Fatalf("result %v, not correct, %v, not valid", res, x)
	}
}

func TestSample1(t *testing.T) {
	pairs := [][]int{{2, 1}, {3, 4}, {3, 2}}
	runSample(t, pairs)
}

func TestSample2(t *testing.T) {
	pairs := [][]int{{4, -2}, {1, 4}, {-3, 1}}
	runSample(t, pairs)
}

func TestSample3(t *testing.T) {
	pairs := [][]int{{100000, -100000}}
	runSample(t, pairs)
}
