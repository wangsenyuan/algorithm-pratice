package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		res := process(reader)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d ", x))
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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	queries := make([][]int, m)
	for i := range m {
		queries[i] = readNNums(reader, 2)
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	n := len(a)

	next := make([]int, n+1)
	next[n] = n
	for i := n - 1; i >= 0; i-- {
		next[i] = next[i+1]
		if i+1 < n && a[i] != a[i+1] {
			next[i] = i + 1
		}
	}

	k := bits.Len(uint(n))

	dp := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		dp[i] = make([]int, k)
		if i+1 < n {
			dp[i][0] = abs(a[i+1] - a[i])
			for j := 1; j < k; j++ {
				if i+1<<(j-1) < n {
					dp[i][j] = gcd(dp[i][j-1], dp[i+(1<<(j-1))][j-1])
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}

	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		l, r := cur[0], cur[1]
		l--
		r--
		if next[l] > r {
			ans[i] = 0
			continue
		}
		// next[l] <= r
		dist := r - l
		pos := l
		for j := k - 1; j >= 0 && pos < r; j-- {
			if (dist>>j)&1 == 1 {
				ans[i] = gcd(ans[i], dp[pos][j])
				pos += 1 << j
			}
		}
	}
	return ans
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func abs(num int) int {
	return max(num, -num)
}
