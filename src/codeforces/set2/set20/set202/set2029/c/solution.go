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
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
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
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

func solve(a []int) int {
	n := len(a)

	f := func(a int, x int) int {
		if a < x {
			a++
		} else if a > x {
			a--
		}
		return a
	}
	dp := []int{0, -n, -n}
	for _, x := range a {
		dp[2] = max(f(dp[1], x), f(dp[2], x))
		dp[1] = max(dp[1], dp[0])
		dp[0] = f(dp[0], x)
	}
	return max(dp[1], dp[2])
}

func solve1(a []int) int {
	n := len(a)
	if n == 1 {
		return 0
	}
	f := make([]int, n)
	var x int
	for i := 0; i < n; i++ {
		if a[i] > x {
			x++
		} else if a[i] < x {
			x--
		}
		f[i] = x
		if i > 0 {
			f[i] = max(f[i], f[i-1])
		}
	}
	check := func(k int) bool {
		g := k
		for i := n - 1; i > 0; i-- {
			if f[i-1] >= g {
				return true
			}
			if a[i] >= g {
				// 如果要达到k，那么在i+1的时候，需要至少是g (为啥是至少）
				// 如果a[i] >= g, 在进入i之前x = g - 1, 这里可以增加到g
				// 如果进入i前是g, 那么这里也可以保持不变
				g--
			} else {
				g++
			}
		}
		return false
	}

	l, r := 0, n
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
