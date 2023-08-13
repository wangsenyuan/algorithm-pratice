package main

import "testing"

func runSample(t *testing.T, n int, m int, commands string, expect []int) {
	r, c := solve(n, m, commands)

	expect_dist := move(n, m, commands, expect[0], expect[1])

	res_dist := move(n, m, commands, r, c)

	if expect_dist != res_dist {
		t.Fatalf("Sample result %d %d, not correct", r, c)
	}
}

func move(n int, m int, cmds string, r int, c int) int {
	r--
	c--
	for i := 0; i < len(cmds); i++ {
		if cmds[i] == 'L' {
			c--
		} else if cmds[i] == 'R' {
			c++
		} else if cmds[i] == 'U' {
			r--
		} else {
			r++
		}
		if r < 0 || r == n || c < 0 || c == m {
			return i
		}
	}
	return len(cmds)
}

func TestSample1(t *testing.T) {
	n, m := 1, 1
	cmds := "L"
	expect := []int{1, 1}
	runSample(t, n, m, cmds, expect)
}

func TestSample2(t *testing.T) {
	n, m := 1, 2
	cmds := "L"
	expect := []int{1, 2}
	runSample(t, n, m, cmds, expect)
}

func TestSample3(t *testing.T) {
	n, m := 3, 3
	cmds := "RRDLUU"
	expect := []int{2, 1}
	runSample(t, n, m, cmds, expect)
}

func TestSample4(t *testing.T) {
	n, m := 4, 3
	cmds := "LUURRDDLLLUU"
	expect := []int{3, 2}
	runSample(t, n, m, cmds, expect)
}
