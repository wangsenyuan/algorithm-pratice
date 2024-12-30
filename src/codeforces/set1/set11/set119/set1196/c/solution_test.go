package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res, robots := process(reader)

	if len(expect) == 0 {
		if len(res) != 0 {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
		return
	}
	if len(res) == 0 {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	x0, y0 := res[0], res[1]
	check := func(robot []int) {
		x, y := robot[0], robot[1]
		if x > x0 && robot[2] == 0 {
			t.Fatalf("Sample result not correct, robot %v, not able to reach %v", robot, res)
		}
		if x < x0 && robot[4] == 0 {
			t.Fatalf("Sample result not correct, robot %v, not able to reach %v", robot, res)
		}
		if y > y0 && robot[5] == 0 {
			t.Fatalf("Sample result not correct, robot %v, not able to reach %v", robot, res)
		}
		if y < y0 && robot[3] == 0 {
			t.Fatalf("Sample result not correct, robot %v, not able to reach %v", robot, res)
		}
	}

	for _, cur := range robots {
		check(cur)
	}
}

func TestSample1(t *testing.T) {
	s := `2
-1 -2 0 0 0 0
-1 -2 0 0 0 0`
	expect := []int{-1, -2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 5 1 1 1 1
2 5 0 1 0 1
3 5 1 0 0 0`
	expect := []int{2, 5}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
1337 1337 0 1 1 1
1336 1337 1 1 0 1`
	expect := []int{}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1
3 5 1 1 1 1`
	expect := []int{3, 5}
	runSample(t, s, expect)
}
