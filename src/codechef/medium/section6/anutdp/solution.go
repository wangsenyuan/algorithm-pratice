package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	q := readNum(scanner)
	L := make([]int, q)
	R := make([]int, q)
	N := make([]int, q)
	for i := 0; i < q; i++ {
		scanner.Scan()
		bs := scanner.Bytes()
		pos := readInt(bs, 0, &L[i])
		pos = readInt(bs, pos+1, &R[i])
		readInt(bs, pos+1, &N[i])
	}
	ans, res := solve(q, L, R, N)
	for i := 0; i < q; i++ {
		fmt.Printf("%d %s\n", ans[i], res[i])
	}
}

var parent map[Pair]int
var valid []int

const MAX_N = 100000000

func init() {
	valid = make([]int, 0, 100000)
	que := make([]int, 0, 100000)
	que = append(que, 0)
	que = append(que, 1)
	parent = make(map[Pair]int)
	parent[Pair{0, 1}] = -1
	var pos int
	for pos+2 <= len(que) {
		a, b := que[pos], que[pos+1]
		pos += 2

		if b == 0 {
			valid = append(valid, a)
		}
		c, d := b, a+b
		if _, found := parent[Pair{c, d}]; !found && d <= MAX_N {
			parent[Pair{c, d}] = a
			que = append(que, c)
			que = append(que, d)
		}
		d = 0
		if _, found := parent[Pair{c, d}]; !found {
			parent[Pair{c, d}] = a
			que = append(que, c)
			que = append(que, d)
		}
	}
	sort.Ints(valid)
}

type Pair struct {
	first, second int
}

type Tripe struct {
	Pair
	third int
	forth int
}

type Tripes []Tripe

func (this Tripes) Len() int {
	return len(this)
}

func (this Tripes) Less(i, j int) bool {
	return this[i].first < this[j].first
}

func (this Tripes) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func solve(q int, L []int, R []int, N []int) ([]int, []string) {
	ps := make([]Tripe, q)
	for i := 0; i < q; i++ {
		ps[i] = Tripe{Pair{L[i], R[i]}, N[i], i}
	}
	sort.Sort(Tripes(ps))
	ans := make([]int, q)
	ss := make([]string, q)
	var pos int
	for i := 0; i < q; i++ {
		ques := ps[i]
		l, r, n, j := ques.first, ques.second, ques.third, ques.forth
		for pos < len(valid) && valid[pos] < l {
			pos++
		}

		n += pos - 1

		if n < len(valid) {
			n = valid[n]
		} else {
			n = MAX_N + 10
		}

		if n > r {
			n = -1
		}
		ans[j] = n
		ss[j] = represent(n)
	}
	return ans, ss
}

var s_cache [100]byte

func represent(n int) string {
	if n < 0 {
		return ""
	}
	a, b, p := n, 0, 0
	for a != -1 {
		if b == 0 {
			s_cache[p] = '#'
		} else {
			s_cache[p] = '.'
		}
		p++
		a, b = parent[Pair{a, b}], a
	}
	for i, j := 0, p-1; i < j; i, j = i+1, j-1 {
		s_cache[i], s_cache[j] = s_cache[j], s_cache[i]
	}
	return string(s_cache[:p-1])
}
