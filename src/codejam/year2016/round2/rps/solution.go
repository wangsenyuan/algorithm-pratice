package main

import (
	"math"
	"sort"
	"fmt"
	"bytes"
)

type rps struct {
	n, p, r, s int
}

func lose(win rune) rune {
	if win == 'P' {
		return 'R'
	}
	if win == 'R' {
		return 'S'
	}

	return 'P'
}

func concat(a, b []rune) []rune {
	c := make([]rune, len(a), len(a) + len(b))
	copy(c, a)
	return append(c, b...)
}

func constructLine(n int, win rune) []rune {
	if n == 1 {
		return []rune{win}
	}

	return concat(constructLine(n / 2, win), constructLine(n / 2, lose(win)))
}

func getRpsKey(line []rune) rps {
	m := make(map[rune]int)
	for _, c := range line {
		m[c]++
	}
	p := m['P']
	r := m['R']
	s := m['S']
	return rps{p + r + s, p, r, s}
}

func sortLine(line []rune) {
	if len(line) == 1 {
		return
	}

	mid := len(line) / 2

	sortLine(line[0:mid])
	sortLine(line[mid:])

	less := true
	for i := 0; less && i < mid; i++ {
		if line[i] > line[mid + i] {
			less = false
		}
	}

	if less {
		return
	}

	for i := 0; i < mid; i++ {
		line[i], line[mid + i] = line[mid + i], line[i]
	}
}

func constructAnswers() map[rps][]rune {
	answers := make(map[rps][]rune)
	cs := []rune{'P', 'R', 'S'}
	for i := 1; i <= 12; i++ {
		n := int(math.Pow(2, float64(i)))
		for _, c := range cs {
			line := constructLine(n, c)
			rps := getRpsKey(line)
			sortLine(line)
			answers[rps] = line
		}
	}

	return answers
}

type plan [][]rune

func (p plan) Less(i, j int) bool {
	a, b := p[i], p[j]
	for i := 0; i < len(a); i++ {
		if a[i] < b[i] {
			return true
		}

		if a[i] > b[i] {
			return false
		}
	}
	return true
}

func (p plan) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p plan) Len() int {
	return len(p)
}
func (solution *Solution) play(n, p, r, s int) ([]rune, bool) {
	lines := make([][]rune, 0, 10)
	for rps, line := range solution.answers {
		if rps.n == n && rps.p == p && rps.r == r && rps.s == s {
			lines = append(lines, line)
		}
	}

	if len(lines) == 0 {
		return nil, false
	}

	sort.Sort(plan(lines))

	return lines[0], true
}


//Solution is a holder
type Solution struct {
	answers map[rps][]rune
}

func NewSolution() *Solution {
	answers := constructAnswers()
	return &Solution{answers}
}


func mkString(line []rune) string {
	var buffer bytes.Buffer
	for _, x := range line {
		buffer.WriteString(string(x))
	}

	return buffer.String()
}

func main() {
	solution := NewSolution()
	var t = 0
	fmt.Scanf("%d", &t)
	for i := 1; i <= t; i++ {
		n, p, r, s := 0, 0, 0,  0
		fmt.Scanf("%d %d %d %d", &n, &r, &p, &s)
		result, found := solution.play(r + p + s, p, r, s)
		if !found {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", i)
		} else {
			fmt.Printf("Case #%d: %s\n", i, mkString(result))
		}

	}
}
