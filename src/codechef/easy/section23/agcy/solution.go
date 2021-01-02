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
		n, q := readTwoNums(reader)
		Q := make([][]int, q)
		for i := 0; i < q; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(n, Q)

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

func solve(n int, Q [][]int) []int64 {
	//qs := make([]Query, len(Q))
	//
	//for i := 0; i < len(Q); i++ {
	//	qs[i] = Query{Q[i][0], Q[i][1]}
	//}
	//sort.Sort(Queries(qs))

	dp := make([]int64, n+1)
	for _, q := range Q {
		l, r := q[0], q[1]
		//dp[l]++
		//dp[r+1]--
		// fp[l] = 1, fp[l+1] = 2, fp[l+2] = 3... fp[r] = r - l + 1, fp[r+1] = 0
		l--
		dp[l]++
		dp[r]--
	}
	// 2, 3, 4
	for i := 1; i <= n; i++ {
		dp[i] += dp[i-1]
	}

	for _, q := range Q {
		l, r := q[0], q[1]
		l--
		dp[r] -= int64(r - l)
	}

	for i := 1; i <= n; i++ {
		dp[i] += dp[i-1]
	}

	return dp[:n]
}

type Query struct {
	left, right int
}

type Queries []Query

func (this Queries) Len() int {
	return len(this)
}

func (this Queries) Less(i, j int) bool {
	return this[i].left < this[j].left || this[i].left == this[j].left && this[i].right < this[j].right
}

func (this Queries) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
