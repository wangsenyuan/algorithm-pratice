package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, cost int, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))

	c1, res := process(reader)
	if cost != c1 || expect != res {
		t.Fatalf("Sample expect %d, %v, but got %d, %v", cost, expect, c1, res)
	}
}

func abs(num int) int {
	return max(num, -num)
}

func TestSample1(t *testing.T) {
	s := `6 5
898196
`
	cost := 4
	expect := "888188"
	runSample(t, s, cost, expect)
}

func TestSample2(t *testing.T) {
	s := `3 2
533
`
	cost := 0
	expect := "533"
	runSample(t, s, cost, expect)
}

func TestSample3(t *testing.T) {
	s := `10 6
0001112223
`
	cost := 3
	expect := "0000002223"
	runSample(t, s, cost, expect)
}

func TestSample4(t *testing.T) {
	s := `82 80
2119762952003918195325258677229419698255491250839396799769357665825441616335532825
`
	cost := 184
	expect := "5555555555005555555555555555555555555555555555555555555555555555555555555555555555"
	runSample(t, s, cost, expect)
}
