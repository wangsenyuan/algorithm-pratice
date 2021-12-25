package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		s, _ := reader.ReadString('\n')
		solver := NewSolver(s[:len(s)-1])

		n := readNum(reader)
		A := readNNums(reader, n)
		var buf bytes.Buffer
		for i := 0; i < n; i++ {
			res := solver.Query(A[i])
			buf.WriteString(strconv.Itoa(res))
			buf.WriteByte('\n')
		}
		fmt.Print(buf.String())
	}
}

type Solver struct {
	n     int
	pair  []int
	right []int
}

func NewSolver(s string) Solver {
	n := len(s)
	stack := make([]int, n)
	pair := make([]int, n)
	right := make([]int, n)
	var p int

	for i := 0; i < n; i++ {
		pair[i] = -2
		if s[i] == '(' {
			stack[p] = i
			p++
		} else if p > 0 {
			pair[stack[p-1]] = i
			p--
		}
	}
	for i := n - 1; i >= 0; i-- {
		if s[i] == '(' {
			right[i] = i
		} else if i < n-1 {
			right[i] = right[i+1]
		} else {
			right[i] = n
		}
	}

	return Solver{n, pair, right}
}

func (solver Solver) Query(i int) int {
	i--
	j := solver.right[i]
	if j == solver.n {
		return -1
	}
	k := solver.pair[j]

	return k + 1
}
