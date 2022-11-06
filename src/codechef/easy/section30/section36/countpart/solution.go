package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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
		if s[i] == '\n' || s[i] == '\r' {
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

const MOD = 998244353

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modSub(a, b int) int {
	return modAdd(a, MOD-b)
}

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
}

func solve(P []int) int {
	n := len(P)

	// left[i] = j 左边最小的j, P[j] < P[j+1] < ... < P[i]
	// right[i] = j 右边最小的j, P[i] < P[j]
	// 假设有一个区间是以 P[i] 作为最大值的 [l, r]
	// left[i] <= l <= i <= r < right[i]
	// dp[i] 是到i为止的有效计数
	// dp[i+1] = 要么从新开始一个part = dp[i]
	//         = 接到前面的part上面
	//             1. P[i+1] 是作为最大值的部分 sum(dp[left[i] - 1], ... dp[i])
	//             2. P[i+1] 是作为前面最大值尾部的部分， 假设，前面比它大的下标是k dp[k]
	//								这样的k有多个

	right := make([]int, n)

	stack := make([]int, n)
	var p int
	for i := 0; i < n; i++ {
		right[i] = n
		for p > 0 && P[stack[p-1]] < P[i] {
			right[stack[p-1]] = i
			p--
		}
		stack[p] = i
		p++
	}

	dp := make([]int, n)
	suf := make([]int, n+2)
	dp[n-1] = 1
	suf[n] = 1
	suf[n-1] = 2
	j := n - 1
	for i := n - 2; i >= 0; i-- {
		if P[i] > P[i+1] {
			j = right[i] - 1
		}
		dp[i] = modSub(suf[i+1], suf[j+2])
		suf[i] = modAdd(suf[i+1], dp[i])
	}

	return dp[0]
}
