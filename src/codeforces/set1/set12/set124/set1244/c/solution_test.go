package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	nums, res := process(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	n, p, w, d := nums[0], nums[1], nums[2], nums[3]
	x, y, z := res[0], res[1], res[2]
	if x < 0 || y < 0 || z < 0 || x+y+z != n {
		t.Fatalf("Sample result %v,not correct", res)
	}
	if w*x+d*y != p {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `30 60 3 1`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 51 5 4`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `20 0 15 5`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1000000000000 99999999999999999 100000 99999`
	expect := true
	runSample(t, s, expect)
}
