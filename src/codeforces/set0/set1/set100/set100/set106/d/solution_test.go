package main

import "testing"

func runSample(t *testing.T, mat []string, inst []string, expect string) {
	res := solve(mat, inst)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := []string{
		"##########",
		"#K#..#####",
		"#.#..##.##",
		"#..L.#...#",
		"###D###A.#",
		"##########",
	}
	inst := []string{
		"N 2",
		"S 1",
		"E 1",
		"W 2",
	}
	expect := "AD"
	runSample(t, mat, inst, expect)
}
