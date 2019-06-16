package p928

import "testing"

func runSample(t *testing.T, graph [][]int, initial []int, expect int) {
	res := minMalwareSpread(graph, initial)
	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", graph, initial, expect, res)
	}
}

func TestSample1(t *testing.T) {
	graph := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	initial := []int{0, 1}
	expect := 0
	runSample(t, graph, initial, expect)
}

func TestSample2(t *testing.T) {
	graph := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	initial := []int{0, 2}
	expect := 0
	runSample(t, graph, initial, expect)
}

func TestSample3(t *testing.T) {
	graph := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	initial := []int{1, 2}
	expect := 1
	runSample(t, graph, initial, expect)
}

func TestSample4(t *testing.T) {
	graph := [][]int{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 1}, {0, 0, 1, 1}}
	initial := []int{1, 3}
	expect := 3
	runSample(t, graph, initial, expect)
}

func TestSample5(t *testing.T) {
	graph := [][]int{{1, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 1, 1}, {0, 0, 1, 1}}
	initial := []int{0, 1}
	expect := 1
	runSample(t, graph, initial, expect)
}
