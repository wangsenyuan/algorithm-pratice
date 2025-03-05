package main

import (
	"bufio"
	"math/bits"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	reader := bufio.NewReader(strings.NewReader(s))

	cards, res := process(reader)

	if len(res) != len(expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	if len(res) == 1 {
		// no solution
		return
	}
	m := len(cards[0])
	if res[1] != "There are no jokers." {
		jokers := []int{-1, -1}
		for i, row := range cards {
			for j, x := range row {
				if x == "J1" {
					jokers[0] = i*m + j
				}
				if x == "J2" {
					jokers[1] = i*m + j
				}
			}
		}
		if jokers[0] >= 0 {
			pos := strings.Index(res[1], "with") + 5
			i, j := jokers[0]/m, jokers[0]%m
			cards[i][j] = res[1][pos : pos+2]
			res[1] = strings.Replace(res[1], "with", "", 1)
		}
		if jokers[1] >= 0 {
			pos := strings.Index(res[1], "with") + 5
			i, j := jokers[1]/m, jokers[1]%m
			cards[i][j] = res[1][pos : pos+2]
		}
	}

	get := func(s string) []int {
		var r, c int

		i := strings.IndexByte(s, '(')
		pos := readInt([]byte(s), i+1, &r)
		for s[pos] < '0' || s[pos] > '9' {
			pos++
		}
		readInt([]byte(s), pos, &c)
		return []int{r, c}
	}
	first := get(res[2])
	second := get(res[3])

	check := func(r int, c int) bool {
		r--
		c--
		var rank int
		var suite int
		for i := r; i < r+3; i++ {
			for j := c; j < c+3; j++ {
				a, b := parse(cards[i][j])
				rank |= 1 << a
				suite |= 1 << b
			}
		}
		return bits.OnesCount(uint(rank)) == 9 || bits.OnesCount(uint(suite)) == 1
	}
	if !check(first[0], first[1]) || !check(second[0], second[1]) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 6
2S 3S 4S 7S 8S AS
5H 6H 7H 5S TC AC
8H 9H TH 7C 8C 9C
2D 2C 3C 4C 5C 6C
`
	expect := []string{
		"No solution.",
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 6
2S 3S 4S 7S 8S AS
5H 6H 7H J1 TC AC
8H 9H TH 7C 8C 9C
2D 2C 3C 4C 5C 6C
`
	expect := []string{
		"Solution exists.",
		"Replace J1 with 2H.",
		"Put the first square to (1, 1).",
		"Put the second square to (2, 4).",
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 6
2S 3S 4S 7S 8S AS
5H 6H 7H QC TC AC
8H 9H TH 7C 8C 9C
2D 2C 3C 4C 5C 6C
`
	expect := []string{
		"Solution exists.",
		"There are no jokers.",
		"Put the first square to (1, 1).",
		"Put the second square to (2, 4).",
	}
	runSample(t, s, expect)
}
