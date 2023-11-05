package p2924

import "testing"

func runSample(t *testing.T, edges [][]int, values []int, expect int) {
	res := maximumScoreAfterOperations(edges, values)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges := [][]int{
		{0, 1}, {0, 2}, {0, 3}, {2, 4}, {4, 5},
	}
	values := []int{5, 2, 5, 2, 1, 1}
	expect := 11
	runSample(t, edges, values, expect)
}

func TestSample2(t *testing.T) {
	edges := [][]int{
		{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6},
	}
	values := []int{20, 10, 9, 7, 4, 3, 5}
	expect := 40
	runSample(t, edges, values, expect)
}
