package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		line := readNNums(reader, 4)
		K, N, L, R := line[0], line[1], line[2], line[3]
		arrs := make([][]int, K)
		for i := 0; i < K; i++ {
			arrs[i] = readNNums(reader, N)
		}
		res := solve(K, N, L, R, arrs)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

const INF = 1 << 60
const N_INF = -INF

func solve(K int, N int, L int, R int, arrs [][]int) int {
	sum := make([][]int64, K)
	for i := 0; i < K; i++ {
		sum[i] = make([]int64, N+1)
		sort.Ints(arrs[i])
		for j := 0; j < N; j++ {
			sum[i][j+1] = sum[i][j] + int64(arrs[i][j])
		}
	}

	check := func(x int) bool {
		// can turn them equal in at most x steps?
		var a, b int64 = 0, INF
		// a is min, and b is max
		for i := 0; i < K; i++ {
			//take x max elements, make them to L, to get a min
			c := sum[i][N-x] + int64(x)*int64(L)
			// take x min elements, make them to R, to gget a max
			d := sum[i][N] - sum[i][x] + int64(x)*int64(R)
			// 1, 2, 3
			a = max(a, c)
			b = min(b, d)
		}
		return a <= b
	}

	l, r := 0, N
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	return a + b - min(a, b)
}
