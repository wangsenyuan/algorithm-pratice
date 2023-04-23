package lcp75

import "testing"

func runSample(t *testing.T, maze []string, expect int) {
	res := challengeOfTheKeeper(maze)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	maze := []string{
		".....", "##S..", "...#.", "T.#..", "###..",
	}
	expect := 7
	runSample(t, maze, expect)
}

func TestSample2(t *testing.T) {
	maze := []string{
		".#..", "..##", ".#S.", ".#.T",
	}
	expect := -1
	runSample(t, maze, expect)
}

func TestSample3(t *testing.T) {
	maze := []string{
		"S###.", "..###", "#..##", "##..#", "###.T",
	}
	expect := 5
	runSample(t, maze, expect)
}
