package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	W := readNNums(reader, n)
	m := readNum(reader)
	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNNums(reader, 2)
	}
	res := solve(W, Q)
	var buf bytes.Buffer

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

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

func solve(W []int, Q [][]int) []int64 {
	// （a, b) pair
	// 对于未知 i (i >= a), 如果 i % b == a % b, 那么需要计算在内
	// 如果迭代Q，那么复杂性将是 n * m
	// 迭代Q， b1, b2, 如果 b2 整除 b1, 那么b2（不考虑a的情况下）， 包含了b1的结果
	qs := make([]Query, len(Q))

	for i, cur := range Q {
		qs[i] = Query{cur[0], cur[1], i}
	}

	sort.Slice(qs, func(i, j int) bool {
		return qs[i].Less(qs[j])
	})

	n := len(W)

	x := int(math.Sqrt(float64(n)))

	ans := make([]int64, len(Q))

	dp := make([]int64, n+1)

	// find2 := func(a, b int) int64 {
	// 	var ans int64
	// 	for i := a; i <= n; i += b {
	// 		ans += int64(W[i-1])
	// 	}
	// 	return ans
	// }

	for i := 0; i < len(qs); {
		if qs[i].b <= x {
			for j := n; j >= 1; j-- {
				if j+qs[i].b > n {
					dp[j] = int64(W[j-1])
				} else {
					dp[j] = dp[j+qs[i].b] + int64(W[j-1])
				}
			}
			b := qs[i].b
			for i < len(qs) && qs[i].b == b {
				ans[qs[i].id] = dp[qs[i].a]
				i++
			}
			continue
		}

		for j := qs[i].a; j <= n; j += qs[i].b {
			ans[qs[i].id] += int64(W[j-1])
		}
		i++
	}
	return ans
}

type Query struct {
	a  int
	b  int
	id int
}

func (this Query) Less(that Query) bool {
	return this.b < that.b
}
