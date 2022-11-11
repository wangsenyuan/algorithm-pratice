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

	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNInt64s(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

const H = 60
const MAX_NUM = int64(1000000000000000000)

func solve(A []int64) int {
	// 先建A[i], 进行开方操作, 使得它等于1
	// n := len(A)
	// 一个简单的dp dp[i][x] = a
	// 当第i项变成x时，使得前i项非递减，一共进行了多少次操作
	// x = 要么是开方，要么是平方，所以次数不会很多
	// 最多 H * H, 而且也不用记录前面的

	prv := []Pair{{1, 0}}

	for _, x := range A {
		costs := gen_all(x)
		var nxt []Pair
		var ptr int
		var mn int = 1e9
		for _, cost := range costs {
			y, d := cost.first, cost.second
			for ptr < len(prv) && prv[ptr].first <= y {
				mn = min(mn, prv[ptr].second)
				ptr++
			}
			nxt = append(nxt, Pair{y, d + mn})
		}
		prv = nxt
	}

	var ans int = 1e9

	for _, cur := range prv {
		ans = min(ans, cur.second)
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func gen_all(x int64) []Pair {
	que := make([]Pair, 0, 1)
	que = append(que, Pair{x, 0})
	vis := make(map[int64]bool)
	vis[x] = true

	for i := 0; i < len(que); i++ {
		u, d := que[i].first, que[i].second
		if u <= 1e9 && !vis[u*u] {
			que = append(que, Pair{u * u, d + 1})
			vis[u*u] = true
		}
		y := sqrt(u)
		if !vis[y] {
			que = append(que, Pair{y, d + 1})
			vis[y] = true
		}
	}
	sort.Slice(que, func(i, j int) bool {
		return que[i].first < que[j].first || que[i].first == que[j].first && que[i].second < que[j].second
	})
	return que
}

func sqrt(x int64) int64 {
	y := int64(math.Sqrt(float64(x)))
	for (y+1)*(y+1) <= x {
		y++
	}
	for y*y > x {
		y--
	}
	return y
}

type Pair struct {
	first  int64
	second int
}
