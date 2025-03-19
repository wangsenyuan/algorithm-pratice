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
	for range tc {
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
	var cnt0 int
	for _, x := range a {
		if x == 0 {
			cnt0++
		}
	}
	if cnt0 == 0 {
		return n
	}
	res := n - cnt0

	var has_0 bool

	var arr []int
	for i := 0; i < n; i++ {
		if a[i] == 0 {
			if !has_0 {
				arr = append(arr, a[i])
				has_0 = true
			}
		} else {
			arr = append(arr, a[i])
		}
	}

	m := len(arr)
	dp := make([]int, m)
	for i := 0; i < m; i++ {
		dp[i] = arr[i]
		if i > 0 {
			dp[i] = min(dp[i], dp[i-1])
		}
	}
	var mex int
	cnt := make([]int, m+1)
	for i := m - 1; i > 0; i-- {
		if arr[i] < m {
			cnt[arr[i]]++
		}
		for cnt[mex] > 0 {
			mex++
		}
		if dp[i-1] < mex {
			return res
		}
	}
	return m
}
