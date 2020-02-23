package p1361

import "testing"

func runSample(t *testing.T, n int, leftChild []int, rightChild []int, expect bool) {
	res := validateBinaryTreeNodes(n, leftChild, rightChild)
	if res != expect {
		t.Errorf("Sample fail")
	}
}

func TestSample1(t *testing.T) {
	n := 4
	left := []int{1, -1, 3, -1}
	right := []int{2, -1, -1, -1}
	expect := true
	runSample(t, n, left, right, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	left := []int{1, -1, 3, -1}
	right := []int{2, 3, -1, -1}
	expect := false
	runSample(t, n, left, right, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	left := []int{1, 0}
	right := []int{-1, -1}
	expect := false
	runSample(t, n, left, right, expect)
}

func TestSample4(t *testing.T) {
	n := 6
	left := []int{1, -1, -1, 4, -1, -1}
	right := []int{2, -1, -1, 5, -1, -1}
	expect := false
	runSample(t, n, left, right, expect)
}
