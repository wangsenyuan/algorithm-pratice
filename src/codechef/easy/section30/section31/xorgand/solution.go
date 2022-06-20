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
		m := readNum(reader)
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 3)
		}
		res := solve(A, Q)
		for _, cur := range res {
			buf.WriteString(fmt.Sprintf("%d\n", cur))
		}
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
		if s[i] == '\n' {
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

const H = 32

func solve(A []int, Q [][]int) []int {
	n := len(A)

	cnt := make([][]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = make([]int, H+1)
		if i > 0 {
			copy(cnt[i], cnt[i-1])
		}
		h := H
		for h > 0 && (A[i]>>uint(h-1))&1 == 0 {
			h--
		}
		cnt[i][h]++
	}

	ans := make([]int, len(Q))
	for i, cur := range Q {
		l, r, x := cur[0], cur[1], cur[2]
		l--
		r--

		b := H
		for b > 0 && (x>>uint(b-1))&1 == 0 {
			b--
		}

		a := cnt[r][b]
		if l > 0 {
			a -= cnt[l-1][b]
		}
		ans[i] = r - l + 1 - a
	}

	return ans
}
