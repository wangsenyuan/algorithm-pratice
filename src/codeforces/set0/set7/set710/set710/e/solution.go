package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, x, y := readThreeNums(reader)

	res := solve(x, y, n)

	fmt.Println(res)
}

const inf = 1 << 60

func solve(x int, y int, n int) int {

	m := n + 1

	arr := make([]int, 2*m)
	for i := 0; i < len(arr); i++ {
		arr[i] = inf
	}

	set := func(p int, v int) {
		p += m
		arr[p] = min(arr[p], v)
		for p > 1 {
			arr[p>>1] = min(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l int, r int) int {
		l += m
		r += m
		res := inf
		for l < r {
			if l&1 == 1 {
				res = min(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = min(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}
	// 输入一个a
	dp := make([]int, n+1)
	dp[1] = x

	set(1, dp[1]+2*x)

	for i := 2; i <= n; i++ {
		// 前一个
		tmp := dp[i-1] + x
		if i%2 == 0 {
			// 复制一份
			tmp = min(tmp, dp[i/2]+y)
		}
		// i = 4, 从3开始
		// i = 5, 从3开始
		j := i/2 + 1
		tmp2 := get(j, i) + y - i*x

		dp[i] = min(tmp, tmp2)

		set(i, dp[i]+2*i*x)
	}

	return dp[n]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
