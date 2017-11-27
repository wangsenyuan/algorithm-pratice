package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		var n int
		scanner.Scan()
		readInt(scanner.Bytes(), 0, &n)
		hints := make([][]byte, n)

		for j := 0; j < n; j++ {
			scanner.Scan()
			hints[j] = scanner.Bytes()
		}
		res := solve(n, hints)
		fmt.Println(res)
	}
}

const MAX = 1000000000

func solve(n int, ss [][]byte) int {
	m := 2*n + 1
	count := make([]int, 2*m)
	touched := make([]bool, 2*m)
	var increment func(left, right, x, y, pos int)
	increment = func(left, right, x, y, pos int) {
		if x > right || y < left {
			return
		}
		touched[pos] = true
		if left <= x && y <= right {
			count[pos]++
			return
		}
		mid := (x + y) / 2
		increment(left, right, x, mid, 2*pos+1)
		increment(left, right, mid+1, y, 2*pos+2)
	}

	var maxRange func(x, y, pos int) int
	maxRange = func(x, y, pos int) int {
		if !touched[pos] {
			return 0
		}
		ans := count[pos]
		if x == y {
			return ans
		}
		mid := (x + y) / 2
		left := maxRange(x, mid, 2*pos+1)
		right := maxRange(mid+1, y, 2*pos+2)
		return ans + max(left, right)
	}

	hints := make([]Hint, n)
	for i := 0; i < n; i++ {
		hints[i] = newHint(ss[i])
	}

	sort.Sort(Hints(hints))

	current, last := 0, 1

	for _, hint := range hints {
		if hint.yes && ((hint.value == 1 && hint.op == '<') || hint.value == MAX && hint.op == '>') {
			continue
		}

		if hint.value > last {
			if hint.value-last > 1 {
				current++
			}
			current++
			last = hint.value
		}
		value := current
		switch hint.op {
		case '<':
			if hint.yes {
				increment(0, value-1, 0, m, 0)
			} else {
				increment(value, m, 0, m, 0)
			}
		case '=':
			if hint.yes {
				increment(value, value, 0, m, 0)
			} else {
				increment(0, value-1, 0, m, 0)
				increment(value+1, m, 0, m, 0)
			}
		case '>':
			if hint.yes {
				increment(value+1, m, 0, m, 0)
			} else {
				increment(0, value, 0, m, 0)
			}
		}
	}
	return n - maxRange(0, m, 0)
}

type Hint struct {
	value int
	op    byte
	yes   bool
}

func newHint(s []byte) Hint {
	op := s[0]
	var value int
	x := readInt(s, 2, &value)
	yes := s[x+1] == 'Y'
	return Hint{value, op, yes}
}

type Hints []Hint

func (hints Hints) Len() int {
	return len(hints)
}
func (hints Hints) Less(i, j int) bool {
	return hints[i].value < hints[j].value
}
func (hints Hints) Swap(i, j int) {
	hints[i], hints[j] = hints[j], hints[i]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
