package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	craters, res := process(bufio.NewReader(strings.NewReader(s)))

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}

	check := func(i int, j int) bool {
		i--
		j--
		if craters[i][0] > craters[j][0] {
			i, j = j, i
		}
		if craters[i][0]+craters[i][1] <= craters[j][0]-craters[j][1] {
			return true
		}
		if craters[i][1] > craters[j][1] {
			i, j = j, i
		}
		// j 更大
		if craters[i][0]-craters[i][1] >= craters[j][0]-craters[j][1] &&
			craters[i][0]+craters[i][1] <= craters[j][0]+craters[j][1] {
			return true
		}
		return false
	}

	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if !check(res[i], res[j]) {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4
1 1
2 2
4 1
5 1
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1
2 3
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
2 1
4 1
6 1
8 1
5 3
`
	expect := 4
	runSample(t, s, expect)
}
