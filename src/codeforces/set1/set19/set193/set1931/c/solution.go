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
	// 只能操作一次,

	n := len(a)
	res := n - 1

	// 三种情况，
	// 保留头部/尾部/头尾
	for i := 1; i <= n; i++ {
		if i == n || a[i] != a[0] {
			res = min(res, n-i)
			break
		}
	}
	if res == 0 {
		return res
	}
	for i := n - 2; i >= 0; i-- {
		if a[i] != a[n-1] {
			res = min(res, i+1)
			break
		}
	}
	if a[0] == a[n-1] {
		var l int
		for l < n && a[l] == a[0] {
			l++
		}
		r := n - 1
		for r >= l && a[r] == a[n-1] {
			r--
		}
		res = min(res, r-l+1)
	}

	return res
}
