package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, j1, j2 int) {
	reader := bufio.NewReader(strings.NewReader(s))

	p1, p2, friends, prices, pizzas := process(reader)



	check := func(x1 int, x2 int) []int {
		a := getState(pizzas[x1-1]) | getState(pizzas[x2-1])

		var cnt int
		for _, f := range friends {
			x := getState(f)
			if a&x == x {
				cnt++
			}
		}
		return []int{cnt, prices[x1-1] + prices[x2-1]}
	}

	a := check(j1, j2)
	b := check(p1, p2)

	if !reflect.DeepEqual(a, b) {
		t.Fatalf("Sample result not correct, expect %d %d, but got %d %d", j1, j2, p1, p2)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4
2 6 7
4 2 3 9 5
3 2 3 9
100 1 7
400 3 3 2 5
100 2 9 2
500 3 2 9 5
`
	j1, j2 := 2, 3
	runSample(t, s, j1, j2)
}

func TestSample2(t *testing.T) {
	s := `4 3
1 1
1 2
1 3
1 4
10 4 1 2 3 4
20 4 1 2 3 4
30 4 1 2 3 4
`
	j1, j2 := 1, 2
	runSample(t, s, j1, j2)
}

func TestSample3(t *testing.T) {
	s := `1 5
9 9 8 7 6 5 4 3 2 1
3 4 1 2 3 4
1 4 5 6 7 8
4 4 1 3 5 7
1 4 2 4 6 8
5 4 1 9 2 8
`
	j1, j2 := 2, 4
	runSample(t, s, j1, j2)
}

func TestSample4(t *testing.T) {
	s := `2 5
5 9 7 3 1 4
5 5 1 6 2 8
1 5 5 6 1 3 2
1 5 7 4 2 1 3
3 3 7 1 5
2 8 1 9 2 4 6 3 7 5
1 4 5 7 4 9
`
	j1, j2 := 1, 5
	runSample(t, s, j1, j2)
}
