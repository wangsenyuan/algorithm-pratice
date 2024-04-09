package main

import "testing"

func runSample(t *testing.T, instructions []string) {
	res := solve(instructions)

	for i := 0; i < H; i++ {
		for j := 0; j < 2; j++ {
			x := process(instructions, i, j)
			y := process(res, i, j)
			if x != y {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
}

func process(instructions []string, d int, x int) int {

	for _, cur := range instructions {
		var v int
		readInt([]byte(cur), 2, &v)
		v = (v >> d) & 1
		if cur[0] == '|' {
			x |= v
		} else if cur[0] == '&' {
			x &= v
		} else {
			x ^= v
		}
	}

	return x
}

func TestSample1(t *testing.T) {
	cmds := []string{
		"| 3",
		"^ 2",
		"| 1",
	}
	runSample(t, cmds)
}

func TestSample2(t *testing.T) {
	cmds := []string{
		"& 1",
		"& 3",
		"& 5",
	}
	runSample(t, cmds)
}

func TestSample3(t *testing.T) {
	cmds := []string{
		"^ 1",
		"^ 3",
		"^ 5",
	}
	runSample(t, cmds)
}
