package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const inf = 1 << 30

func solve(a []int) int {
	n := len(a)
	// n <= 1e4
	// a[i] <= 1000
	ma := slices.Max(a) * 2

	dp := make([]int, ma+1)
	ndp := make([]int, ma+1)

	for i := 0; i <= ma; i++ {
		dp[i] = inf
		ndp[i] = inf
	}
	dp[a[n-1]] = a[n-1]

	for i := n - 2; i >= 0; i-- {
		for r := 0; r <= ma; r++ {
			w := dp[r]
			if w == inf {
				continue
			}
			l := w - r
			nr := a[i] + r
			nl := min(0, a[i]-l)
			// 放在正方向
			if nr <= ma {
				ndp[nr] = min(ndp[nr], nr-nl)
			}
			nr = a[i] + l
			nl = min(0, a[i]-r)
			if nr <= ma {
				ndp[nr] = min(ndp[nr], nr-nl)
			}
		}
		for r := ma; r >= 0; r-- {
			dp[r] = ndp[r]
			ndp[r] = inf
		}
	}
	return slices.Min(dp)
}
