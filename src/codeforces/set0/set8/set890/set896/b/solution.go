package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, c := readThreeNums(reader)

	solver := NewSolver(n, c)

	for i := 0; i < m; i++ {
		x := readNum(reader)
		ok, y := solver.Play(x)
		fmt.Println(y)
		if ok {
			break
		}
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

type Solver struct {
	n      int
	c      int
	first  []int
	second []int
}

func NewSolver(n int, c int) *Solver {
	first := make([]int, 0, n/2)
	second := make([]int, 0, n/2)
	return &Solver{n, c, first, second}
}

func (s *Solver) Play(x int) (bool, int) {
	var pos int
	if x <= s.c/2 {
		// put it in the front
		pos = sort.Search(len(s.first), func(i int) bool {
			return s.first[i] > x
		})
		if pos == len(s.first) {
			s.first = append(s.first, x)
		} else {
			s.first[pos] = x
		}
	} else {
		// s.second 是倒序的
		tmp := sort.Search(len(s.second), func(i int) bool {
			return s.second[i] < x
		})
		if tmp == len(s.second) {
			pos = s.n - 1 - len(s.second)
			s.second = append(s.second, x)
		} else {
			pos = s.n - 1 - tmp
			s.second[tmp] = x
		}
	}
	ok := s.n == len(s.first)+len(s.second)
	return ok, pos + 1
}
