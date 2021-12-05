package p2096

import "testing"

func runSample(t *testing.T, pairs [][]int) {
	res := validArrangement(pairs)

	if len(res) != len(pairs) || !valid(res) {
		t.Errorf("Sample result %v, not correct", res)
	}
}

func valid(res [][]int) bool {
	prev := res[0][1]
	for i := 1; i < len(res); i++ {
		if res[i][0] != prev {
			return false
		}
		prev = res[i][1]
	}

	return true
}

func TestSample1(t *testing.T) {
	pairs := [][]int{
		{17, 18}, {18, 10}, {10, 18},
	}
	runSample(t, pairs)
}

func TestSample2(t *testing.T) {
	pairs := [][]int{
		{8, 5}, {8, 7}, {0, 8}, {0, 5}, {7, 0}, {5, 0}, {0, 7}, {8, 0}, {7, 8},
	}
	runSample(t, pairs)
}

func TestSample3(t *testing.T) {
	pairs := [][]int{
		{17, 18}, {18, 10}, {10, 18},
	}
	runSample(t, pairs)
}
