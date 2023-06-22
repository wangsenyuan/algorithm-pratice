package main

import (
	"testing"
)

func runSample(t *testing.T, cmds string, expect bool) {
	res := solve(cmds)

	if res[0] == 0 && res[1] == 0 && expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	x0, y0 := res[0], res[1]

	var x, y int

	for i := 0; i < len(cmds); i++ {
		if cmds[i] == 'L' && (y != y0 || (x-1 > x0 || x < x0)) {
			x--
		} else if cmds[i] == 'R' && (y != y0 || (x+1 < x0 || x > x0)) {
			x++
		} else if cmds[i] == 'D' && (x != x0 || (y-1 > y0 || y < y0)) {
			y--
		} else if cmds[i] == 'U' && (x != x0 || (y+1 < y0 || y > y0)) {
			y++
		}
	}

	if x != 0 || y != 0 {
		t.Fatalf("Sample result %v giving wrong final result %d %d", res, x, y)
	}
}

func TestSample1(t *testing.T) {
	cmds := "L"
	runSample(t, cmds, true)
}

func TestSample2(t *testing.T) {
	cmds := "RUUDL"
	runSample(t, cmds, true)
}

func TestSample3(t *testing.T) {
	cmds := "LLUU"
	runSample(t, cmds, false)
}
func TestSample4(t *testing.T) {
	cmds := "DDDUUUUU"
	runSample(t, cmds, true)
}
