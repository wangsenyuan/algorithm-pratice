package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for i := 0; i < tc; i++ {
		buf.WriteString(fmt.Sprintf("Case #%d:\n", i+1))
		readString(reader)
		line_width := readNum(reader)
		n := readNum(reader)
		words := readNNums(reader, n)
		solver := NewSolver(words, line_width)

		for {
			s, _ := reader.ReadBytes('\n')
			if s[0] == 'E' {
				break
			}
			if s[0] == 'I' {
				var id int
				readInt(s, 2, &id)
				ans := solver.QueryLineNo(id)
				buf.WriteString(fmt.Sprintf("%d\n", ans))
			} else {
				var id, v int
				pos := readInt(s, 2, &id) + 1
				readInt(s, pos, &v)
				solver.ChangeWordLength(id, v)
			}
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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
	words      []int
	arr        []int
	n          int
	line_width int
}

func NewSolver(words []int, line_width int) *Solver {
	n := len(words)

	arr := make([]int, n+1)

	for i := 0; i < n; i++ {
		words[i]++
		set(arr, i+1, words[i])
	}

	line_width++

	return &Solver{words, arr, n, line_width}
}

func (solver *Solver) QueryLineNo(id int) int {
	cur := 1
	no := 1
	for {
		kk := get(solver.arr, cur-1)
		kk += solver.line_width
		cur = find(solver.arr, kk)
		if cur >= id {
			return no
		}
		cur++
		no++
	}
}

func (solver *Solver) ChangeWordLength(id int, width int) {
	width++
	set(solver.arr, id, width-solver.words[id-1])
	solver.words[id-1] = width
}

func set(arr []int, p int, v int) {
	for p < len(arr) {
		arr[p] += v
		p += p & -p
	}
}

func get(arr []int, p int) int {
	var res int
	for p > 0 {
		res += arr[p]
		p -= p & -p
	}
	return res
}

func find(arr []int, sum int) int {
	var x int
	var ans int
	for i := 16; i >= 0; i-- {
		if (x + (1 << uint(i))) < len(arr) {
			x |= 1 << uint(i)
			if ans+arr[x] > sum {
				x ^= 1 << uint(i)
			} else {
				ans += arr[x]
			}
		}
	}
	return x
}
