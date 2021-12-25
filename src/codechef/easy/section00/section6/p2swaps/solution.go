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
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		P := readNNums(reader, n)
		q := readNum(reader)
		Q := make([][]int, q)
		for i := 0; i < q; i++ {
			s, _ := reader.ReadBytes('\n')
			if s[0] == '1' {
				Q[i] = []int{1}
			} else if s[0] == '2' {
				Q[i] = make([]int, 3)
				Q[i][0] = 2
				pos := readInt(s, 2, &Q[i][1]) + 1
				readInt(s, pos, &Q[i][2])
			} else {
				Q[i] = make([]int, 2)
				Q[i][0] = 3
				readInt(s, 2, &Q[i][1])
			}
		}
		res := solve(n, A, P, Q)
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d\n", res[i]))
		}
	}

	fmt.Print(buf.String())
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

const B = 20

func solve(n int, A []int, P []int, Q [][]int) []int {
	invp := make([]int, n)
	for i := 0; i < n; i++ {
		invp[P[i]-1] = i
	}

	next := make([][]int, B)
	for i := 0; i < B; i++ {
		next[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		next[0][i] = invp[i]
	}

	for b := 1; b < B; b++ {
		for i := 0; i < n; i++ {
			next[b][i] = next[b-1][next[b-1][i]]
		}
	}
	pos := func(u int, steps int) int {
		for b := B - 1; b >= 0; b-- {
			if (steps>>uint(b))&1 == 1 {
				u = next[b][u]
			}
		}
		return u
	}

	res := make([]int, 0, len(Q))

	var cnt int
	for _, q := range Q {
		t := q[0]
		if t == 1 {
			cnt++
		} else if t == 2 {
			x, y := q[1], q[2]
			x, y = pos(x-1, cnt), pos(y-1, cnt)
			A[x], A[y] = A[y], A[x]
		} else {
			x := q[1]
			res = append(res, A[pos(x-1, cnt)])
		}
	}

	return res
}
