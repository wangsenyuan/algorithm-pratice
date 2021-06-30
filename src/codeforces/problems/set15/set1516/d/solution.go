package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, q := readTwoNums(reader)
	A := readNNums(reader, n)

	Q := make([][]int, q)

	for i := 0; i < q; i++ {
		Q[i] = readNNums(reader, 2)
	}

	res := solve(n, A, Q)
	var buf bytes.Buffer

	for i := 0; i < q; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
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

const H = 20

func solve(n int, A []int, Q [][]int) []int {
	ma := A[0]
	for i := 1; i < n; i++ {
		ma = max(ma, A[i])
	}

	pf := make([]int, ma+1)
	for i := 2; i <= ma; i++ {
		if pf[i] == 0 {
			for j := i; j <= ma; j += i {
				pf[j] = i
			}
		}
	}

	pos := make([]int, ma+1)

	for i := 0; i <= ma; i++ {
		pos[i] = n
	}

	next := make([]int, n+1)
	next[n] = n

	for i := n - 1; i >= 0; i-- {
		next[i] = n
		cur := A[i]
		for cur > 1 {
			next[i] = min(next[i], pos[pf[cur]])
			cur /= pf[cur]
		}
		next[i] = min(next[i], next[i+1])
		cur = A[i]
		for cur > 1 {
			pos[pf[cur]] = i
			cur /= pf[cur]
		}
	}

	dp := make([][]int, n+1)
	for i := n; i >= 0; i-- {
		dp[i] = make([]int, H)
		dp[i][0] = next[i]
		for j := 1; j < H; j++ {
			dp[i][j] = dp[dp[i][j-1]][j-1]
		}
	}

	ans := make([]int, len(Q))

	for i, cur := range Q {
		l, r := cur[0], cur[1]
		l--
		r--

		for j := H - 1; j >= 0; j-- {
			if dp[l][j] <= r {
				ans[i] += 1 << uint(j)
				l = dp[l][j]
			}
		}
		ans[i]++
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
