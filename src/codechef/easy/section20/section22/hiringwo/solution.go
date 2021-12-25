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
		k, x := readTwoNums(reader)
		res := solve(k, x)
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

func solve(k int, x int) int {
	values := make([]int, 0, 7)

	for i := 2; i <= x; i++ {
		if x%i == 0 {
			var c int = 1
			for x%i == 0 {
				c *= i
				x /= i
			}
			values = append(values, c)
		}
	}

	if len(values) <= k {
		// ans for 1
		var ans = k - len(values)
		for _, v := range values {
			ans += v
		}
		return ans
	}

	// k < len(values)
	n := len(values)
	tot := 1 << uint(n)

	prod := make([]int, tot)

	for state := 0; state < tot; state++ {
		res := 1
		for i := 0; i < n; i++ {
			if state&(1<<uint(i)) > 0 {
				res *= values[i]
			}
		}
		prod[state] = res
	}
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, tot)
	}

	dp[1] = prod

	for layer := 2; layer <= k; layer++ {
		for j := 0; j < tot; j++ {
			dp[layer][j] = 100000000
		}
		for mask := 0; mask < tot; mask++ {
			for sub := mask - 1; sub > 0; sub = (sub - 1) & mask {
				dp[layer][mask] = min(dp[layer][mask], dp[layer-1][sub]+prod[mask^sub])
			}
		}
	}

	return dp[k][tot-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
