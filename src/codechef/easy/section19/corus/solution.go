package main

import (
	"bufio"
	"fmt"
	"os"
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
		n, q := readTwoNums(reader)
		s, _ := reader.ReadString('\n')
		solver := NewSolver(n, s)

		for q > 0 {
			q--
			c := readNum(reader)
			fmt.Println(solver.Query(c))
		}
	}
}

type Solver struct {
	n   int
	cnt []int
}

func NewSolver(n int, S string) Solver {
	cnt := make([]int, 26)
	for i := 0; i < n; i++ {
		cnt[int(S[i]-'a')]++
	}

	return Solver{n, cnt}
}

func (solver Solver) Query(c int) int {
	if c == 0 {
		// all in the pending queue
		return solver.n
	}

	var res int

	for i := 0; i < 26; i++ {
		res += max(0, solver.cnt[i]-c)
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
