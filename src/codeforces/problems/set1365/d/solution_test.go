package main

import "testing"

func runSample(t *testing.T, n, m int, maze []string, expect bool) {
	res := solve(n, m, maze)

	if res != expect {
		t.Errorf("Sample %v, expect %t, but got %t", maze, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 1, 1
	maze := []string{
		".",
	}
	expect := true

	runSample(t, n, m, maze, expect)

}

func TestSample2(t *testing.T) {
	n, m := 1, 2
	maze := []string{
		"G.",
	}
	expect := true

	runSample(t, n, m, maze, expect)

}

func TestSample3(t *testing.T) {
	n, m := 2, 2
	maze := []string{
		"#B",
		"G.",
	}
	expect := false

	runSample(t, n, m, maze, expect)

}

func TestSample4(t *testing.T) {
	n, m := 2, 3
	maze := []string{
		"G.#",
		"B#.",
	}
	expect := false

	runSample(t, n, m, maze, expect)

}

func TestSample5(t *testing.T) {
	n, m := 3, 3
	maze := []string{
		"#B.",
		"#..",
		"GG.",
	}
	expect := true

	runSample(t, n, m, maze, expect)
}

func TestSample6(t *testing.T) {
	n, m := 2, 2
	maze := []string{
		"#B",
		"B.",
	}
	expect := true

	runSample(t, n, m, maze, expect)
}

//
//BBG
//##.

func TestSample7(t *testing.T) {
	n, m := 2, 3
	maze := []string{
		"BBG",
		"##.",
	}
	expect := false

	runSample(t, n, m, maze, expect)
}
