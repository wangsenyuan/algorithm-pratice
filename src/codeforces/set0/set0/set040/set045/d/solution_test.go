package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))

	tasks, res := process(reader)

	occ := make(map[int]int)

	for i := range len(tasks) {
		l, r := tasks[i][0], tasks[i][1]
		d := res[i]
		if d < l || d > r {
			t.Fatalf("wrong day %d for task %v", d, tasks[i])
		}
		occ[d]++
		if occ[d] > 1 {
			t.Fatalf("Sample result %v, not correct, it has assigned same day to different tasks", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 2
2 3
3 4
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `2
1 3
1 3
`
	runSample(t, s)
}