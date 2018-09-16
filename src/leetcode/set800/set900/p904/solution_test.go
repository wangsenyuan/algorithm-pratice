package p904

import "testing"

func runSample(t *testing.T, tree []int, expect int) {
	res := totalFruit(tree)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", tree, expect, res)
	}
}

func TestSample1(t *testing.T) {
	tree := []int{1, 2, 1}
	expect := 3
	runSample(t, tree, expect)
}

func TestSample2(t *testing.T) {
	tree := []int{0, 1, 2, 2}
	expect := 3
	runSample(t, tree, expect)
}

func TestSample3(t *testing.T) {
	tree := []int{1, 2, 3, 2, 2}
	expect := 4
	runSample(t, tree, expect)
}

func TestSample4(t *testing.T) {
	tree := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	expect := 5
	runSample(t, tree, expect)
}
