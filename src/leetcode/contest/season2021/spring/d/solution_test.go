package d

import "testing"

func runSample(t *testing.T, maze [][]string, expect bool) {
	res := escapeMaze(maze)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	maze := [][]string{
		{".#.", "#.."}, {"...", ".#."}, {".##", ".#."}, {"..#", ".#."},
	}
	expect := true
	runSample(t, maze, expect)
}

func TestSample2(t *testing.T) {
	maze := [][]string{
		{".#.", "..."}, {"...", "..."},
	}
	expect := false
	runSample(t, maze, expect)
}

func TestSample3(t *testing.T) {
	maze := [][]string{
		{"...", "...", "..."}, {".##", "###", "##."}, {".##", "###", "##."}, {".##", "###", "##."}, {".##", "###", "##."}, {".##", "###", "##."}, {".##", "###", "##."},
	}
	expect := false
	runSample(t, maze, expect)
}
