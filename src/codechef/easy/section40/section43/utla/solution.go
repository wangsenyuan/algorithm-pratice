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
		P := readNNums(reader, n)
		res := solve(P)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

func solve(A []int) int {
	n := len(A)
	// A[0] + A[1] = A[2] + A[3] = A[3] + A[4] = ...
	// n <= 2000
	// A[i] <= 1e9
	// A[0] 和 A[1] 是可能被删除的
	// 如果给定了 A[i] + A[j] 是其中的一对，那么剩下的部分就很容易确定
	// 但是，O(n ** 3) 的复杂性
	// 可以divide-and-conquer吗？

	pairs := make(map[int][]Pair)
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			sum := A[i] + A[j]
			if len(pairs[sum]) == 0 {
				pairs[sum] = make([]Pair, 0, 1)
			}
			pairs[sum] = append(pairs[sum], Pair{i, j})
		}
	}
	var best int
	for _, v := range pairs {
		best = max(best, find(v))
	}

	return n - best
}

type Pair struct {
	first  int
	second int
}

func find(ps []Pair) int {
	// ps already sorted
	rx := -1
	var res int
	for _, p := range ps {
		l, r := p.first, p.second
		if l <= rx {
			continue
		}
		res++
		rx = r
	}
	return res * 2
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
