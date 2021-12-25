package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		S, _ := reader.ReadString('\n')
		n := readNum(reader)
		Q := make([][]int, n)
		for i := 0; i < n; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(len(S)-1, S, Q)

		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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
func solve(n int, S string, Q [][]int) []int {
	stack := make([]int, n+1)
	var p int
	push := func(v int) {
		stack[p] = v
		p++
	}

	pop := func() int {
		res := stack[p-1]
		p--
		return res
	}

	peek := func() int {
		return stack[p-1]
	}

	push(-1)
	st := make([]int, n)
	en := make([]int, n)
	pp := make([]int, n)
	negative := make([]bool, n)
	var neg bool
	var pos int

	for i := 0; i < n; i++ {
		if S[i] == '?' {
			st[pos] = i
			en[pos] = i
			pp[pos] = peek()
			negative[pos] = neg
			neg = false
			pos++
		} else if S[i] == '(' {
			st[pos] = i
			pp[pos] = peek()
			negative[pos] = neg
			neg = false
			push(pos)
			pos++
		} else if S[i] == ')' {
			en[pop()] = i
		} else if S[i] == '-' {
			neg = true
		}
	}

	adj := make([][]int, pos)

	for i := 0; i < pos; i++ {
		adj[i] = make([]int, 0, 2)
	}

	for i := 1; i < pos; i++ {
		adj[pp[i]] = append(adj[pp[i]], i)
	}

	mn := make([]int, pos)
	mx := make([]int, pos)

	for i := pos - 1; i >= 0; i-- {
		if st[i] == en[i] {
			mn[i] = 0
			mx[i] = 1
		} else {
			for _, v := range adj[i] {
				if negative[v] {
					mn[i] -= mx[v]
					mx[i] -= mn[v]
				} else {
					mn[i] += mn[v]
					mx[i] += mx[v]
				}
			}
		}
	}

	res := make([]int, len(Q))

	for i, cur := range Q {
		l, r := cur[0], cur[1]
		l--
		r--
		lo, hi := 0, pos-1
		for lo < hi {
			mid := (lo + hi) / 2
			if st[mid] >= l {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		res[i] = mx[hi]
	}

	return res
}
