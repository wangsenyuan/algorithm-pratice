package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	h := readNNums(reader, n)
	res := solve(h)
	fmt.Println(res)
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

func solve(h []int) int {
	n := len(h)
	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = -1
		b[i] = -1
		c[i] = -1
		d[i] = -1
	}

	st := make([]int, n)
	var p int

	for i := 0; i < n; i++ {
		for p > 0 && h[st[p-1]] > h[i] {
			p--
		}
		if p > 0 {
			a[i] = st[p-1]
		}
		st[p] = i
		p++
	}
	p = 0
	for i := 0; i < n; i++ {
		for p > 0 && h[st[p-1]] < h[i] {
			p--
		}
		if p > 0 {
			b[i] = st[p-1]
		}
		st[p] = i
		p++
	}
	p = 0
	for i := n - 1; i >= 0; i-- {
		for p > 0 && h[st[p-1]] > h[i] {
			p--
		}
		if p > 0 {
			c[i] = st[p-1]
		}
		st[p] = i
		p++
	}
	p = 0
	for i := n - 1; i >= 0; i-- {
		for p > 0 && h[st[p-1]] < h[i] {
			p--
		}
		if p > 0 {
			d[i] = st[p-1]
		}
		st[p] = i
		p++
	}

	jumps := make([][]int, n)
	for i := 0; i < n; i++ {
		jumps[i] = make([]int, 0, 1)
	}
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = i
		if a[i] >= 0 {
			jumps[a[i]] = append(jumps[a[i]], i)
		}
		if b[i] >= 0 {
			jumps[b[i]] = append(jumps[b[i]], i)
		}
		if c[i] >= 0 {
			jumps[i] = append(jumps[i], c[i])
		}
		if d[i] >= 0 {
			jumps[i] = append(jumps[i], d[i])
		}
	}

	for i := 0; i < n; i++ {
		for _, j := range jumps[i] {
			dp[j] = min(dp[j], dp[i]+1)
		}
	}

	return dp[n-1]
}

func min(arr ...int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res > arr[i] {
			res = arr[i]
		}
	}
	return res
}

func bruteForce(arr []int) int {
	n := len(arr)
	dp := make([]int, n)

	for i := n - 2; i >= 0; i-- {
		dp[i] = dp[i+1] + 1
		x, y := arr[i+1], arr[i+1]
		for j := i + 2; j < n; j++ {
			if max(arr[i], arr[j]) < x {
				dp[i] = min(dp[i], dp[j]+1)
			}
			if min(arr[i], arr[j]) > y {
				dp[i] = min(dp[i], dp[j]+1)
			}
			x = min(x, arr[j])
			y = max(y, arr[j])
		}
	}
	return dp[0]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
