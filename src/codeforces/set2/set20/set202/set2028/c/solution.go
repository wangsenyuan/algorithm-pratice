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
		res := process(reader)
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

func process(reader *bufio.Reader) int {
	n, m, v := readThreeNums(reader)
	a := readNNums(reader, n)
	return solve(a, m, v)
}

func solve(a []int, m int, v int) int {
	n := len(a)
	// dp[i]表示在前缀i中的，满足每段和>= v的数量
	dp := make([]int, n+1)
	var sum int
	var tot int
	for i := 0; i < n; i++ {
		dp[i+1] = dp[i]
		sum += a[i]
		if sum >= v {
			dp[i+1]++
			sum = 0
		}
		tot += a[i]
	}
	if dp[n] < m {
		return -1
	}
	fp := make([]int, n+1)
	sum = 0
	for i := n - 1; i >= 0; i-- {
		fp[i] = fp[i+1]
		sum += a[i]
		if sum >= v {
			fp[i]++
			sum = 0
		}
	}

	// 能否得到至少为exp的一段满足条件
	check := func(exp int) bool {
		var sum int
		for l, r := 0, 0; r < n; r++ {
			sum += a[r]
			for l <= r && sum-a[l] >= exp {
				sum -= a[l]
				l++
			}
			if sum >= exp && fp[r+1]+dp[l] >= m {
				return true
			}
		}
		return false
	}

	l, r := 0, tot-v*m+1
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return r - 1
}
