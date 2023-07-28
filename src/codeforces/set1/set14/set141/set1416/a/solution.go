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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)

		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(A []int) []int {
	n := len(A)
	last := make([]int, n+1)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		x := A[i-1]
		f[x] = max(f[x], i-last[x])
		last[x] = i
	}

	for x := 1; x <= n; x++ {
		f[x] = max(f[x], n-last[x]+1)
	}

	res := make([]int, n+1)
	for cur, k := 1, n; k > 0; k-- {
		// k 长度的子串一共有 n - k + 1个
		res[k] = -1
		for cur <= n && f[cur] > k {
			cur++
		}
		if cur <= n {
			res[k] = cur
		}

	}

	return res[1:]
}

func solve1(A []int) []int {
	n := len(A)
	last := make([]int, n+1)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		x := A[i-1]
		f[x] = max(f[x], i-last[x])
		last[x] = i
	}
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = -1
	}
	for x := 1; x <= n; x++ {
		f[x] = max(f[x], n-last[x]+1)
		for i := f[x]; i <= n && ans[i] < 0; i++ {
			ans[i] = x
		}
	}

	return ans[1:]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
