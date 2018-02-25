package p789

import "testing"

func runSample(t *testing.T, ghosts [][]int, target []int, expect bool) {
	res := escapeGhosts(ghosts, target)

	if res != expect {
		t.Errorf("sample %v %v, expect %t, but got %t", ghosts, target, expect, res)
	}
}

func TestSample1(t *testing.T) {
	ghosts := [][]int{
		{1, 0}, {0, 3},
	}

	target := []int{
		0, 1,
	}

	runSample(t, ghosts, target, true)
}

func TestSample2(t *testing.T) {
	ghosts := [][]int{
		{1, 0},
	}

	target := []int{
		2, 0,
	}

	runSample(t, ghosts, target, false)
}

func TestSample3(t *testing.T) {
	ghosts := [][]int{
		{2, 0},
	}

	target := []int{
		1, 0,
	}

	runSample(t, ghosts, target, false)
}
